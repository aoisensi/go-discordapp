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
	return "channels/" + string(s.ChannelID)
}

type ChannelType string

const (
	ChannelTypeText  = ChannelType("text")
	ChannelTypeVoice = ChannelType("voice")
)

//Guild channels represent an isolated set of users and messages within a Guild.
type GuildChannel struct {
	ID                   Snowflake   `json:"id"`
	GuildID              Snowflake   `json:"guild_id"`
	Name                 string      `json:"name"`
	Type                 ChannelType `json:"type"`
	Position             int         `json:"position"`
	IsPrivate            bool        `json:"is_private"`
	PermissionOverwrites interface{} `json:"permission_overwrites"`
	Topic                string      `json:"topic"`
	LastMessageID        Snowflake   `json:"last_message_id"`
	Bitrate              int         `json:"bitrate"`
}

//DM Channels represent a one-to-one conversation between two users, outside of the scope of guilds.
type DMChannel struct {
	ID            Snowflake `json:"id"`
	IsPrivate     bool      `json:"is_private"`
	Recipient     User      `json:"recipient"`
	LastMessageID Snowflake `json:"last_message_id"`
}

//Represents a message sent in a channel within Discord.
type Message struct {
	ID              Snowflake    `json:"id"`
	ChannelID       Snowflake    `json:"channel_id"`
	Author          User         `json:"author"`
	Content         string       `json:"content"`
	Timestamp       Timestamp    `json:"timestamp"`
	EditedTimestamp *Timestamp   `json:"edited_timestamp"`
	TTS             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	Mentions        []User       `json:"mentions"`
	Attachments     []Attachment `json:"attachments"`
	Embeds          []Embed      `json:"embeds"`
	Nonce           string       `json:"nonce"`
}

type OverwriteType string

const (
	OverwriteTypeRole   = OverwriteType("role")
	OverwriteTypeMember = OverwriteType("member")
)

type Overwrite struct {
	ID    Snowflake     `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow Permissions   `json:"allow"`
	Deny  Permissions   `json:"deny"`
}

type Embed struct {
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Provider    EmbedProvider  `json:"provider"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Attachment struct {
	ID       Snowflake `json:"id"`
	Filename string    `json:"filename"`
	Size     int       `json:"size"`
	URL      string    `json:"url"`
	ProxyURL string    `json:"proxy_url"`
	Height   int       `json:"height"`
	Width    int       `json:"width"`
}

func (s *ChannelsService) CreateMessage(content string) error {
	var data struct {
		Content string `json:"content"`
		Nonce   string `json:"nonce,omitempty"`
		TTS     bool   `json:"tts,omitempty"`
	}
	data.Content = content
	url := s.baseURL() + "/messages"
	req, err := s.client.NewRequest("POST", url, &data)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
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
