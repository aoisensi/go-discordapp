package discord

import (
	"fmt"
	"net/http"
)

type GuildsService struct {
	client *Client
}

type Guild struct {
	ID                Snowflake `json:"id"`
	Name              string    `json:"name"`
	Icon              string    `json:"icon"`
	Splash            string    `json:"Splash"`
	OwnerID           string    `json:"owner_id"`
	Region            string    `json:"region"`
	AFKChannelId      string    `json:"afk_channel_id"`
	AFKTimeout        int       `json:"afk_timeout"`
	EmbedEnabled      bool      `json:"embed_enabled"`
	EmbedChannelID    string    `json:"embed_channel_id"`
	VerificationLevel int       `json:"verification_level"`
	//VoiceStates
	//Roles
	//Emojis
	//Features
}

type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

type Emoji struct {
	ID            Snowflake   `json:"id"`
	Name          string      `json:"name"`
	Roles         []Snowflake `json:"roles"`
	RequireColons bool        `json:"require_colons"`
	Managed       bool        `json:"managed"`
}

func (s *GuildsService) Get(id string) (*Guild, *http.Response, error) {
	u := fmt.Sprintf("guilds/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	guild := new(Guild)
	resp, err := s.client.Do(req, guild)
	if err != nil {
		return nil, resp, err
	}

	return guild, resp, err
}
