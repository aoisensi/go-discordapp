package discord

type VoiceService struct {
	client *Client
}

//Used to represent a user's voice connection status.
type VoiceState struct {
	GuildID   *Snowflake `json:"guild_id,string"`
	ChannelID Snowflake  `json:"chahhel_id,string"`
	UserID    Snowflake  `json:"user_id,string"`
	SessionID string     `json:"session_id"`
	Deaf      bool       `json:"deaf"`
	Mute      bool       `json:"mute"`
	SelfDeaf  bool       `json:"self_deaf"`
	SelfMute  bool       `json:"self_mute"`
	Suppress  bool       `json:"suppress"`
}

type VoiceRegion struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SampleHostname string `json:"sample_hostname"`
	SamplePort     string `json:"sample_port"`
	VIP            bool   `json:"vip"`
	Optimal        bool   `json:"optimal"`
	Deprecated     bool   `json:"deprecated"`
	Custom         bool   `json:"custom"`
}

func (c *Client) ListVoiceRegions() ([]*VoiceRegion, error) {
	req, err := c.NewRequest("GET", "voice/regions", nil)
	if err != nil {
		return nil, err
	}
	var r []*VoiceRegion
	if _, err := c.Do(req, &r); err != nil {
		return nil, err
	}
	return r, nil
}
