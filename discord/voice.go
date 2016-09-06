package discord

type VoiceState struct {
	UserID    Snowflake `json:"user_id"`
	SessionID string    `json:"session_id"`
}
