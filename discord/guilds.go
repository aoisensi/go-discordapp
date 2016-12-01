package discord

import "net/http"

type GuildsService struct {
	client  *Client
	GuildID Snowflake
}

func (s *GuildsService) baseURL() string {
	return "guilds/" + s.GuildID.String()
}

type Guild struct {
	ID                          Snowflake      `json:"id,string"`
	Name                        string         `json:"name"`
	Icon                        string         `json:"icon"`
	Splash                      string         `json:"Splash"`
	OwnerID                     Snowflake      `json:"owner_id,string"`
	Region                      string         `json:"region"`
	AFKChannelID                Snowflake      `json:"afk_channel_id,string"`
	AFKTimeout                  int            `json:"afk_timeout"`
	EmbedEnabled                bool           `json:"embed_enabled"`
	EmbedChannelID              string         `json:"embed_channel_id,string"`
	VerificationLevel           int            `json:"verification_level"`
	DefaultMessageNotifications int            `json:"default_message_notifications"`
	Roles                       []Role         `json:"roles"`
	Emojis                      []Emoji        `json:"emojis"`
	MFALevel                    int            `json:"mfa_level"`
	Large                       *bool          `json:"large,omitempty"`
	Unavailable                 *bool          `json:"unavailable,omitempty"`
	MemberCount                 int            `json:"member_count"`
	VoiceStates                 *[]VoiceState  `json:"voice_states,omitempty"`
	Members                     *[]GuildMember `json:"members,omitempty"`
	//TODO
	//Channels
	//Features
	//JoindAt
	//Presences
}

type UnavailableGuild struct {
	ID          Snowflake `json:"id,string"`
	Unavailable bool      `json:"unavailable"`
}

type GuildEmbed struct {
	Enabled   bool      `json:"enabled"`
	ChannelID Snowflake `json:"channel_id"`
}

type GuildMember struct {
	User  User        `json:"user"`
	Nick  *string     `json:"nick"`
	Roles []Snowflake `json:"roles"`
	Deaf  bool        `json:"deaf"`
	Mute  bool        `json:"mute"`
	//TODO
	//JoinedAt
}

type Integration struct {
	ID                Snowflake          `json:"id,string"`
	Name              string             `json:"name"`
	Type              string             `json:"type"`
	Enabled           bool               `json:"enabled"`
	Syncing           bool               `json:"syncing"`
	RoleID            Snowflake          `json:"role_id"`
	ExpireBehavior    int                `json:"expire_behavior"`
	ExpireGracePeriod int                `json:"expire_grace_period"`
	User              User               `json:"user"`
	Account           IntegrationAccount `json:"account"`
	SyncedAt          Timestamp          `json:"timestamp"`
}

type IntegrationAccount struct {
	ID   string `json:"id`
	Name string `json:"name"`
}

type Emoji struct {
	ID            Snowflake   `json:"id,string"`
	Name          string      `json:"name"`
	Roles         []Snowflake `json:"roles,string"`
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
