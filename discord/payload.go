package discord

import (
	"encoding/json"
	"fmt"
)

type opcode int

const (
	opcodeDispatch opcode = iota
	opcodeHeartbeat
	opcodeIdentify
	opcodeStatusUpdate
	opcodeVoiceStatusUpdate
	opcodeVoiceServerPing
	opcodeResume
	opcodeReconnect
	opcodeRequestGuildMembers
	opcodeInvalidSession
	opcodeHello
	opcodeHeartbackACK
)

type payload struct {
	Opcode   opcode           `json:"op"`
	Raw      *json.RawMessage `json:"d"`
	Sequence int              `json:"s,omitempty"`
	Name     EventName        `json:"t,omitempty"`
	Data     payloadData      `json:"-"`
}

type payloadData interface {
	encode() *payload
}

type payloadDispatch struct {
	Raw   *json.RawMessage
	Event Event
}

func (p *payloadDispatch) encode() *payload {
	return &payload{
		Opcode:   opcodeDispatch,
		Data:     p,
		Sequence: 0,
		Name:     p.Event.EventName(),
	}
}

type payloadHeartbeat int

func (p *payloadHeartbeat) encode() *payload {
	return &payload{
		Opcode: opcodeHeartbeat,
		Data:   p,
	}
}

type payloadIdentify struct {
	Token           string            `json:"token"`
	Properties      map[string]string `json:"properties"`
	Compress        bool              `json:"compress"`
	LargeThreashold int               `json:"large_threashold"`
	Shard           [2]int            `json:"shard"`
}

func (p *payloadIdentify) encode() *payload {
	return &payload{
		Opcode: opcodeIdentify,
		Data:   p,
	}
}

type payloadStatusUpdate struct {
	IdleSince *int  `json:"idle_since"`
	Game      *Game `json:"game"`
}

func (p *payloadStatusUpdate) encode() *payload {
	return &payload{
		Opcode: opcodeStatusUpdate,
		Data:   p,
	}
}

type payloadVoiceStateUpdate struct {
	GuildID   Snowflake  `json:"guild_id"`
	ChannelID *Snowflake `json:"channel_id"`
	SelfMute  bool       `json:"self_mute"`
	SelfDeaf  bool       `json:"self_deaf"`
}

func (p *payloadVoiceStateUpdate) encode() *payload {
	return &payload{
		Opcode: opcodeVoiceStatusUpdate,
		Data:   p,
	}
}

type payloadVoiceServerPing struct{}

func (p *payloadVoiceServerPing) encode() *payload {
	return &payload{
		Opcode: opcodeVoiceServerPing,
	}
}

type payloadHello struct {
	HeartbeatInterval int      `json:"heartbeat_interval"`
	Trace             []string `json:"_trace"`
}

func (p *payloadHello) encode() *payload {
	return &payload{
		Opcode: opcodeHello,
		Data:   p,
	}
}

type payloadHeartbackACK struct{}

func (p *payloadHeartbackACK) encode() *payload {
	return &payload{
		Opcode: opcodeHeartbackACK,
	}
}

func (pl *payload) encodeData() error {
	if pl.Raw != nil {
		return nil
	}
	var data interface{}
	if d, ok := pl.Data.(*payloadDispatch); ok {
		data = d.Event
	} else {
		data = pl.Data
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	raw := json.RawMessage(body)
	pl.Raw = &raw
	return nil
}

func (pl *payload) decode() (payloadData, error) {
	unmarshal := false
	switch pl.Opcode {
	case opcodeDispatch:
		data := new(payloadDispatch)
		data.Raw = pl.Raw
		if err := data.decode(pl.Name); err != nil {
			return nil, err
		}
		pl.Data = data
	case opcodeHello:
		pl.Data = new(payloadHello)
		unmarshal = true
	case opcodeHeartbackACK:
		pl.Data = new(payloadHeartbackACK)

	default:
		panic(fmt.Sprintf("tried unknown opcode: %v", pl.Opcode))
	}
	fmt.Println("FDSFDSF")
	var err error
	if unmarshal {
		err = json.Unmarshal([]byte(*pl.Raw), pl.Data)
	}
	return pl.Data, err
}
