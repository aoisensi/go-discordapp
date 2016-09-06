package discord

/*
import "golang.org/x/oauth2"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://discordapp.com/api/oauth2/authorize",
	TokenURL: "https://discordapp.com/api/oauth2/token",
}
*/

type Scope string

const (
	ScopeIdentify    = Scope("identify")
	ScopeEmail       = Scope("email")
	ScopeConnections = Scope("connections")
	ScopeGuilds      = Scope("guilds")
	ScopeGuildsJoin  = Scope("guilds.join")
	ScopeBot         = Scope("bot")
)
