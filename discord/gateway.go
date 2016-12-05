package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Gateway struct {
	url      string
	interval time.Duration
	sequence int
	ws       *websocket.Conn
	mu       sync.RWMutex
	//	wch      chan *payload
	readCh   chan *payload
	errorCh  chan error
	handleCh chan bool
	handlers []interface{}
}

func NewGateway() (*Gateway, error) {
	resp, err := http.Get("https://discordapp.com/api/gateway")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := struct {
		URL string `json:"url"`
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	v := url.Values{
		"v":        {"5"},
		"encoding": {"json"},
	}
	url := fmt.Sprint(data.URL, "?", v.Encode())

	return &Gateway{
		url: url,
		//		wch: make(chan *payload),
		readCh:   make(chan *payload),
		errorCh:  make(chan error),
		handleCh: make(chan bool),
		handlers: make([]interface{}, 0),
	}, nil
}

func (g *Gateway) Start(token string) error {
	ws, err := websocket.Dial(g.url, "", "https://localhost/")
	if err != nil {
		return err
	}
	g.ws = ws

	pl, err := g.receive()
	if err != nil {
		return err
	}
	_, err = pl.decode()
	if err != nil {
		return err
	}
	var hello *payloadHello
	var ok bool
	if hello, ok = pl.Data.(*payloadHello); !ok {
		panic("first recieve is not hello. wtf?")
	}
	g.interval = time.Millisecond * time.Duration(hello.HeartbeatInterval)
	go g.heart()

	pli := payloadIdentify{
		Token: token,
		Properties: map[string]string{
			"$os":               runtime.GOOS,
			"$browser":          "go-discordapp",
			"$device":           "go-discordapp",
			"$referrer":         "",
			"$referring_domain": "",
		},
		Compress:        false,
		LargeThreashold: 250,
		Shard:           [2]int{0, 1},
	}

	if err := g.send(pli.encode()); err != nil {
		return err
	}

	go g.receiver()

	for {
		select {
		case err := <-g.errorCh:
			if err != io.EOF {
				return nil
			}
			panic(err)
		case pl := <-g.readCh:
			if data, ok := pl.Data.(*payloadDispatch); ok {
				g.handle(data.Event)
			}
		}
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
	g.sequence = pl.Sequence
	if _, err := pl.decode(); err != nil {
		return nil, err
	}
	return pl, nil
}

func (g *Gateway) receiver() {
	for {
		pl, err := g.receive()
		if err != nil {
			g.errorCh <- err
			continue
		}
		g.readCh <- pl
	}
}

var gwCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		msg, err := json.Marshal(v)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(string(msg))
		return msg, websocket.TextFrame, err
	},
	Unmarshal: func(msg []byte, payloadType byte, v interface{}) error {
		err := json.Unmarshal(msg, v)
		if err != nil {
			return err
		}
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println(string(msg))
		return nil
	},
}
