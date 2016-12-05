package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"time"

	"net/http"
	"net/url"

	"golang.org/x/net/websocket"
)

type Gateway struct {
	token     string
	url       string
	interval  time.Duration
	sequence  int
	ws        *websocket.Conn
	receiveCh chan *payload
	errorCh   chan error
	handleCh  chan bool
	handlers  []interface{}
}

func (c *Client) Gateway() (*Gateway, *http.Response, error) {
	u := "https://discordapp.com/api/gateway"
	if c.bot {
		u += "/bot"
	}
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var data struct {
		URL    string `json:"url"`
		Shards int    `json:"shards"`
	}
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	v := url.Values{"v": {"5"}, "encoding": {"json"}}
	url := fmt.Sprint(data.URL, "?", v.Encode())

	return &Gateway{
		url:       url,
		token:     c.token,
		receiveCh: make(chan *payload),
		errorCh:   make(chan error),
		handleCh:  make(chan bool),
	}, resp, nil
}

func (g *Gateway) Start() error {
	ws, err := websocket.Dial(g.url, "", "https://localhost/")
	defer ws.Close()
	if err != nil {
		return err
	}
	g.ws = ws

	go g.receiver()

	for {
		select {
		case err := <-g.errorCh:
			panic(err)
		case pl := <-g.receiveCh:
			g.parse(pl)
		}
	}
}

func (g *Gateway) parse(pl *payload) {
	switch data := pl.Data.(type) {
	case *payloadDispatch:
		g.handle(data.Event)
	case *payloadHeartbackACK:
		break
	case *payloadHello:
		g.interval = time.Millisecond * time.Duration(data.HeartbeatInterval)
		go g.heart()

		pli := payloadIdentify{
			Token: g.token,
			Properties: map[string]string{
				"os":               runtime.GOOS,
				"browser":          "go-discordapp",
				"device":           "go-discordapp",
				"referrer":         "",
				"referring_domain": "",
			},
			Compress:        false,
			LargeThreashold: 250,
			Shard:           [2]int{0, 1},
		}

		if err := g.send(pli.encode()); err != nil {
			g.errorCh <- err
		}
	default:
		panic("unknown payload")
	}
}

func (g *Gateway) Close() error {
	if g.ws == nil {
		return nil
	}
	return g.ws.Close()
}

func (g *Gateway) AddHandler(handler interface{}) {
	g.handlers = append(g.handlers, handler)
}

func (g *Gateway) StatusUpdate(idle int, game *Game) error {
	pl := &payloadStatusUpdate{
		Game: game,
	}
	if idle > 0 {
		pl.IdleSince = &idle
	}
	return g.send(pl.encode())
}

func (g *Gateway) handle(event Event) {
	for _, handler := range g.handlers {
		hType := reflect.TypeOf(handler)
		if hType.NumIn() != 1 {
			continue
		}
		if hType.In(0) == reflect.TypeOf(event) {
			reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(event)})
		}
	}
}

func (g *Gateway) heart() {
	ticker := time.NewTicker(g.interval)
	for {
		pl := payloadHeartbeat(g.sequence)
		err := g.send(pl.encode())
		if err != nil && err != io.EOF {
			g.errorCh <- err
		}
		time.Sleep(g.interval)
		select {
		case <-ticker.C:
			continue
		case <-g.handleCh:
			return
		}
	}
}

func (g *Gateway) send(pl *payload) error {
	if err := pl.encodeData(); err != nil {
		return err
	}
	err := gwCodec.Send(g.ws, pl)
	return err
}

func (g *Gateway) receive() (*payload, error) {
	pl := new(payload)
	err := gwCodec.Receive(g.ws, pl)

	if err != nil {
		return nil, err
	}
	if pl.Sequence != nil {
		g.sequence = *pl.Sequence
	}
	if _, err := pl.decode(); err != nil {
		return nil, err
	}
	return pl, nil
}

func (g *Gateway) receiver() {
	for {
		pl, err := g.receive()
		if err != nil {
			if err != io.EOF {
				g.errorCh <- err
			}
			continue
		}
		g.receiveCh <- pl
	}
}

var gwCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		msg, err := json.Marshal(v)
		if err != nil {
			return nil, websocket.UnknownFrame, err
		}
		//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		//fmt.Println(string(msg))
		return msg, websocket.TextFrame, err
	},
	Unmarshal: func(msg []byte, payloadType byte, v interface{}) error {
		err := json.Unmarshal(msg, v)
		if err != nil {
			return err
		}
		//fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		//fmt.Println(string(msg))
		return nil
	},
}
