package discord

import "encoding/json"

type EventName string

const (
	EventNameReady                   EventName = "READY"
	EventNameResumed                 EventName = "RESUMED"
	EventNameChannelCreate           EventName = "CHANNEL_CREATE"
	EventNameChannelUpdate           EventName = "CHANNEL_UPDATE"
	EventNameChannelDelete           EventName = "CHANNEL_DELETE"
	EventNameGuildBanAdd             EventName = "GUILD_BAN_ADD"
	EventNameGuildbanRemove          EventName = "GUILDBAN_REMOVE"
	EventNameGuildCreate             EventName = "GUILD_CREATE"
	EventNameGuildUpdate             EventName = "GUILD_UPDATE"
	EventNameGuildEmojiUpdate        EventName = "GUILD_EMOJI_UPDATE"
	EventNameGuildDelete             EventName = "GUILD_DELETE"
	EventNameGuildIntegrationsUpdate EventName = "GUILD_INTEGRATIONS_UPDATE"
	EventNameGuildMemberAdd          EventName = "GUILD_MEMBER_ADD"
	EventNameGuildMemberRemove       EventName = "GUILD_MEMBER_REMOVE"
	EventNameGuildMemberUpdate       EventName = "GUILD_MEMBER_UPDATE"
	EventNameGuildMembersChunk       EventName = "GUILD_MEMBERS_CHUNK"
	EventNameGuildRoleCreate         EventName = "GUILD_ROLE_CREATE"
	EventNameGuildRoleUpdate         EventName = "GUILD_ROLE_UPDATE"
	EventNameGuildRoleDelete         EventName = "GUILD_ROLE_DELETE"
	EventNameMessageCreate           EventName = "MESSAGE_CREATE"
	EventNameMessageUpdate           EventName = "MESSAGE_UPDATE"
	EventNameMessageDelete           EventName = "MESSAGE_DELETE"
	EventNameMessageDeleteBulk       EventName = "MESSAGE_DELETE_BULK"
	EventNamePresenceUpdate          EventName = "PRESENCE_UPDATE"
	EventNameTypingStart             EventName = "TYPING_START"
	EventNameUserSettingsUpdate      EventName = "USER_SETTINGS_UPDATE"
	EventNameUserUpdate              EventName = "USER_UPDATE"
	EventNameVoiceStateUpdate        EventName = "VOICE_STATE_UPDATE"
	EventNameVoiceServerUpdate       EventName = "VOICE_SERVER_UPDATE"

	EventNameUnknown EventName = "UNKNOWN"
)

type Event interface {
	EventName() EventName
}

type EventReady struct {
	Version         int                `json:"v"`
	User            User               `json:"user"`
	PrivateChannels []DMChannel        `json:"private_channels"`
	Guilds          []UnavailableGuild `json:"guilds"`
	SessionID       string             `json:"session_id"`
}

func (e *EventReady) EventName() EventName {
	return EventNameReady
}

type EventResumed struct {
	Trace []string `json:"_trace"`
}

func (e *EventResumed) EventName() EventName {
	return EventNameResumed
}

type EventChannelCreate struct{} //TODO

func (e *EventChannelCreate) EventName() EventName {
	return EventNameChannelCreate
}

type EventChannelUpdate Guild

func (e *EventChannelUpdate) EventName() EventName {
	return EventNameChannelUpdate
}

type EventChannelDelete struct{} //TODO

func (e *EventChannelDelete) EventName() EventName {
	return EventNameChannelDelete
}

type EventGuildCreate Guild

func (e *EventGuildCreate) EventName() EventName {
	return EventNameGuildCreate
}

type EventGuildUpdate Guild

func (e *EventGuildUpdate) EventName() EventName {
	return EventNameGuildUpdate
}

type EventGuildDelete struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

func (e *EventGuildDelete) EventName() EventName {
	return EventNameGuildDelete
}

type EventGuildBanAdd User

func (e *EventGuildBanAdd) EventName() EventName {
	return EventNameGuildBanAdd
}

type EventGuildbanRemove User

func (e *EventGuildbanRemove) EventName() EventName {
	return EventNameGuildbanRemove
}

type EventGuildEmojiUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []Emoji   `json:"emojis"`
}

func (e *EventGuildEmojiUpdate) EventName() EventName {
	return EventNameGuildEmojiUpdate
}

type EventGuildIntegrationsUpdate struct {
	GuildID Snowflake `json:"guild_id"`
}

func (e *EventGuildIntegrationsUpdate) EventName() EventName {
	return EventNameGuildIntegrationsUpdate
}

type EventGuildMemberAdd struct {
	GuildID Snowflake `json:"guild_id"`
}

func (e *EventGuildMemberAdd) EventName() EventName {
	return EventNameGuildMemberAdd
}

type EventGuildMemberRemove struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

func (e *EventGuildMemberRemove) EventName() EventName {
	return EventNameGuildMemberRemove
}

type EventGuildMemberUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Roles   []Role    `json:"roles"`
	User    User      `json:"user"`
}

func (e *EventGuildMemberUpdate) EventName() EventName {
	return EventNameGuildMemberUpdate
}

type EventGuildMembersChunk struct {
	GuildID Snowflake     `json:"guild_id"`
	Members []interface{} `json:"members"` //TODO
}

func (e *EventGuildMembersChunk) EventName() EventName {
	return EventNameGuildMembersChunk
}

type EventGuildRoleCreate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

func (e *EventGuildRoleCreate) EventName() EventName {
	return EventNameGuildRoleCreate
}

type EventGuildRoleUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

func (e *EventGuildRoleUpdate) EventName() EventName {
	return EventNameGuildRoleUpdate
}

type EventGuildRoleDelete struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
}

func (e *EventGuildRoleDelete) EventName() EventName {
	return EventNameGuildRoleDelete
}

type EventMessageCreate Message

func (e *EventMessageCreate) EventName() EventName {
	return EventNameMessageCreate
}

type EventMessageUpdate Message

func (e *EventMessageUpdate) EventName() EventName {
	return EventNameMessageUpdate
}

type EventMessageDelete struct {
	ID        Snowflake `json:"id"`
	ChannelID Snowflake `json:"channel_id"`
}

func (e *EventMessageDeleteBulk) EventName() EventName {
	return EventNameMessageDeleteBulk
}

type EventMessageDeleteBulk struct {
	IDs       []Snowflake `json:"ids"`
	ChannelID Snowflake   `json:"channel_id"`
}

func (e *EventMessageDelete) EventName() EventName {
	return EventNameMessageDelete
}

type EventPresenceUpdate struct {
	User    User        `json:"user"`
	Roles   []Snowflake `json:"roles"`
	Game    Game        `json:"game"` //TODO
	Nick    string      `json:"nick"`
	GuildID Snowflake   `json:"guild_id"`
	Status  Status      `json:"status"`
}

func (e *EventPresenceUpdate) EventName() EventName {
	return EventNamePresenceUpdate
}

type EventTypingStart struct {
	ChannelID Snowflake `json:"channel_id"`
	UserID    Snowflake `json:"user_id"`
	Timestamp Unixtime  `json:"timestamp"`
}

func (e *EventTypingStart) EventName() EventName {
	return EventNameTypingStart
}

type EventUserSettingsUpdate struct{} //TODO

func (e *EventUserSettingsUpdate) EventName() EventName {
	return EventNameUserSettingsUpdate
}

type EventUserUpdate User

func (e *EventUserUpdate) EventName() EventName {
	return EventNameUserUpdate
}

type EventVoiceStateUpdate VoiceState

func (e *EventVoiceStateUpdate) EventName() EventName {
	return EventNameVoiceStateUpdate
}

type EventVoiceServerUpdate struct {
	Token    string    `json:"token"`
	GuildID  Snowflake `json:"guild_id"`
	Endpoint string    `json:"endpoint"`
}

func (e *EventVoiceServerUpdate) EventName() EventName {
	return EventNameVoiceServerUpdate
}

type EventUnknown map[string]interface{}

func (e *EventUnknown) EventName() EventName {
	return EventNameUnknown
}

func (p *payloadDispatch) decode(name EventName) error {
	var event Event
	switch name {
	case EventNameReady:
		event = new(EventReady)
	case EventNameResumed:
		event = new(EventResumed)
	case EventNameChannelCreate:
		event = new(EventChannelCreate)
	case EventNameChannelUpdate:
		event = new(EventChannelUpdate)
	case EventNameChannelDelete:
		event = new(EventChannelDelete)
	case EventNameGuildBanAdd:
		event = new(EventGuildBanAdd)
	case EventNameGuildbanRemove:
		event = new(EventGuildbanRemove)
	case EventNameGuildCreate:
		event = new(EventGuildCreate)
	case EventNameGuildUpdate:
		event = new(EventGuildUpdate)
	case EventNameGuildEmojiUpdate:
		event = new(EventGuildEmojiUpdate)
	case EventNameGuildDelete:
		event = new(EventGuildDelete)
	case EventNameGuildIntegrationsUpdate:
		event = new(EventGuildIntegrationsUpdate)
	case EventNameGuildMemberAdd:
		event = new(EventGuildMemberAdd)
	case EventNameGuildMemberRemove:
		event = new(EventGuildMemberRemove)
	case EventNameGuildMemberUpdate:
		event = new(EventGuildMemberUpdate)
	case EventNameGuildMembersChunk:
		event = new(EventGuildMembersChunk)
	case EventNameGuildRoleCreate:
		event = new(EventGuildRoleCreate)
	case EventNameGuildRoleUpdate:
		event = new(EventGuildRoleUpdate)
	case EventNameGuildRoleDelete:
		event = new(EventGuildRoleDelete)
	case EventNameMessageCreate:
		event = new(EventMessageCreate)
	case EventNameMessageUpdate:
		event = new(EventMessageUpdate)
	case EventNameMessageDelete:
		event = new(EventMessageDelete)
	case EventNameMessageDeleteBulk:
		event = new(EventMessageDeleteBulk)
	case EventNamePresenceUpdate:
		event = new(EventPresenceUpdate)
	case EventNameTypingStart:
		event = new(EventTypingStart)
	case EventNameUserSettingsUpdate:
		event = new(EventUserSettingsUpdate)
	case EventNameUserUpdate:
		event = new(EventUserUpdate)
	case EventNameVoiceStateUpdate:
		event = new(EventVoiceStateUpdate)
	case EventNameVoiceServerUpdate:
		event = new(EventVoiceServerUpdate)
	default:
		event = new(EventUnknown)
	}
	err := json.Unmarshal([]byte(*p.Raw), event)
	p.Event = event
	return err
}
