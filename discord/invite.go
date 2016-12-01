package discord

import "time"

type Invite struct {
	Code    string        `json:"code"`
	Guild   InviteGuild   `json:"guild"`
	Channel InviteChannel `json:"channel"`
}

type InviteMetadata struct {
	Inviter   User      `json:"inviter"`
	Uses      int       `json:"uses"`
	MaxUses   int       `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
	Revoked   bool      `json:"revoked"`
}

type InviteGuild struct {
	ID         Snowflake `json:"id"`
	Name       string    `json:"name"`
	SplashHash string    `json:"splash_hash"`
}

type InviteChannel struct {
	ID   Snowflake   `json:"id"`
	Name string      `json:"name"`
	Type ChannelType `json:"type"`
}
