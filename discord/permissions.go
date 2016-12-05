package discord

type Permissions uint64

const (
	PermissionCreateInstantInvite = Permissions(1 << iota)
	PermissionKickMembers
	PermissionBanMembers
	PermissionAdministrator
	PermissionManageChannels
	PermissionManageGuild
)

const (
	PermissionReadMessages = Permissions(1 << (iota + 10))
	PermissionSendMessages
	PermissionSendTTSMessages
	PermissionManageMessages
	PermissionEmbedLinks
	PermissionAttachFiles
	PermissionReadMessageHistory
	PermissionMentionEveryone
)

const (
	PermissionConnect = Permissions(1 << (iota + 20))
	PermissionSpeak
	PermissionMuteMembers
	PermissionDeafenMembers
	PermissionMoveMembers
	PermissionUserVAD
	PermissionChangeNickname
	PermissionManageNicknames
	PermissionManageRoles
)

type Role struct {
	ID          Snowflake      `json:"id"`
	Name        string      `json:"name"`
	Color       int         `json:"color"`
	Hoist       bool        `json:"hoist"`
	Position    int         `json:"position"`
	Permissions Permissions `json:"permissions"`
	Managed     bool        `json:"managed"`
	Mentionable bool        `json:"mentionable"`
}
