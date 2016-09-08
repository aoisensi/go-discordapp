package discord

import "net/http"

type GuildsService struct {
	client  *Client
	GuildID Snowflake
}

func (s *GuildsService) baseURL() string {
	return "guilds/" + string(s.GuildID)
}

type Guild struct {
	ID                Snowflake `json:"id"`
	Name              string    `json:"name"`
	Icon              string    `json:"icon"`
	Splash            string    `json:"Splash"`
	OwnerID           string    `json:"owner_id"`
	Region            string    `json:"region"`
	AFKChannelID      string    `json:"afk_channel_id"`
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

//Returns the new guild object for the given id.
func (s *GuildsService) GetGuild() (*Guild, *http.Response, error) {
	req, err := s.client.NewRequest("GET", s.baseURL(), nil)
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

//Returns a list of guild channel objects.
func (s *GuildsService) GetGuildChannel() ([]GuildChannel, *http.Response, error) {
	req, err := s.client.NewRequest("GET", s.baseURL()+"/channels", nil)
	if err != nil {
		return nil, nil, err
	}

	var channels []GuildChannel
	resp, err := s.client.Do(req, &channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, err
}

//Create a new channel object for the guild.
//Requires the 'MANAGE_CHANNELS' permission.
//Returns the new channel object on success.
//Fires a Channel Create Gateway event.
/*
func (s *GuildsService) CreateGuildChannel(name string, t ChannelType, bitrate int, userLimit int) ([]*GuildChannel, *http.Response, error) {
	var type
	req, err := s.client.NewRequest("POST", s.baseURL(), nil)
	if err != nil {
		return nil, nil, err
	}

	channels := make([]*GuildChannel, 10)
	resp, err := s.client.Do(req, channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, err
}
*/
