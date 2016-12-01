package discord

import (
	"bytes"
	"io"
	"mime/multipart"
)

type ChannelsService struct {
	client    *Client
	ChannelID Snowflake
}

func (s *ChannelsService) baseURL() string {
	return "channels/" + s.ChannelID.String()
}

type ChannelType string

const (
	ChannelTypeText  = ChannelType("text")
	ChannelTypeVoice = ChannelType("voice")
)

//Guild channels represent an isolated set of users and messages within a Guild.
type GuildChannel struct {
	ID                   Snowflake   `json:"id,string"`
	GuildID              Snowflake   `json:"guild_id,string"`
	Name                 string      `json:"name"`
	Type                 ChannelType `json:"type"`
	Position             int         `json:"position"`
	IsPrivate            bool        `json:"is_private"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites"`
	Topic                *string     `json:"topic,omitempty"`
	LastmessageID        *Snowflake  `json:"last_message_id,string,omitempty"`
	Bitrate              *int        `json:"bitrate,omitempty"`
	UserLimit            *int        `json:"user_limit,omitempty"`
}

//DM Channels represent a one-to-one conversation between two users, outside of the scope of guilds.
type DMChannel struct {
	ID            Snowflake `json:"id,string"`
	IsPrivate     bool      `json:"is_private"`
	Recipient     User      `json:"recipient"`
	LastMessageID Snowflake `json:"last_message_id,string"`
}

//Represents a message sent in a channel within Discord.
type Message struct {
	ID              Snowflake    `json:"id,string"`
	ChannelID       Snowflake    `json:"channel_id,string"`
	Author          *User        `json:"author,omitempty"`
	Content         string       `json:"content"`
	Timestamp       Timestamp    `json:"timestamp"`
	EditedTimestamp *Timestamp   `json:"edited_timestamp"`
	TTS             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	Mentions        []User       `json:"mentions"`
	MentionRoles    []Snowflake  `json:"mention_roles,string"`
	Attachments     []Attachment `json:"attachments"`
	Embeds          []Embed      `json:"embeds"`
	Reactions       []Reaction   `json:"reactions"`
	Nonce           *Snowflake   `json:"nonce"`
	Pinned          bool         `json:"pinned"`
	WebhookID       string       `json:"webhook_id"`
}

type Reaction struct {
	Count int  `json:"count"`
	Me    bool `json:"me"`
	Emoji struct {
		ID   *Snowflake `json:"id,string"`
		Name string     `json:"name"`
	} `json:"emoji"`
}

type OverwriteType string

const (
	OverwriteTypeRole   = OverwriteType("role")
	OverwriteTypeMember = OverwriteType("member")
)

type Overwrite struct {
	ID    Snowflake     `json:"id,string"`
	Type  OverwriteType `json:"type"`
	Allow Permissions   `json:"allow"`
	Deny  Permissions   `json:"deny"`
}

type Embed struct {
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Timestamp   Timestamp      `json:"timestamp"`
	Color       int            `json:"color"`
	Footer      EmbedFooter    `json:"footer"`
	Image       EmbedImage     `json:"image"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Video       EmbedVideo     `json:"video"`
	Provider    EmbedProvider  `json:"provider"`
	Author      EmbedAuthor    `json:"author"`
	Fields      []EmbedField   `json:"fields"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedVideo struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Attachment struct {
	ID       Snowflake `json:"id,string"`
	Filename string    `json:"filename"`
	Size     int       `json:"size"`
	URL      string    `json:"url"`
	ProxyURL string    `json:"proxy_url"`
	Height   *int      `json:"height"`
	Width    *int      `json:"width"`
}

func (s *ChannelsService) CreateMessage(content string) (*Message, error) {
	var data struct {
		Content string `json:"content"`
		Nonce   string `json:"nonce,omitempty"`
		TTS     bool   `json:"tts,omitempty"`
	}
	data.Content = content
	url := s.baseURL() + "/messages"
	req, err := s.client.NewRequest("POST", url, &data)
	if err != nil {
		return nil, err
	}
	var msg Message
	_, err = s.client.Do(req, &msg)
	return &msg, err
}

func (s *ChannelsService) UploadFile(filename string, body io.Reader) error {
	buf := new(bytes.Buffer)
	mp := multipart.NewWriter(buf)
	writer, err := mp.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, body); err != nil {
		return err
	}
	if err := mp.Close(); err != nil {
		return err
	}
	url := s.baseURL() + "/messages"
	req, err := s.client.newRequest("POST", url, buf, "multipart/form-data")
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	return err
}
