package discord

type Webhook struct {
	ID        Snowflake  `json:"id,string"`
	GuildID   *Snowflake `json:"guild_id,string"`
	ChannelID *Snowflake `json:"channel_id,string"`
	User      *User      `json:"user"`
	Name      *string    `json:"name"`
	Avatar    *string    `json:"avatar"`
	Token     string     `json:"token"`
}
