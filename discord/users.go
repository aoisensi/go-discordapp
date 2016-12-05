package discord

import (
	"fmt"
	"net/http"
)

type UsersService struct {
	client *Client
}

type User struct {
	ID            Snowflake `json:"id,string"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        *string   `json:"avatar,omitempty"`
	Bot           bool      `json:"bot"`
	MFAEnabled    bool      `json:"mfa_enabled"`
	Verified      bool      `json:"verified"`
	Email         *string   `json:"email,omitempty"`
}

//A brief version of a guild object
type UserGuild struct {
	ID          Snowflake `json:"id,string"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Owner       bool      `json:"owner"`
	Permissions int       `json:"permissions"`
}

//The connection object that the user has attached.
type Connection struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	Revoked      bool          `json:"revoked"`
	Integrations []Integration `json:"integrations"`
}

//Returns the user object of the requestors account.
//For OAuth2, this requires the identify scope, which will return the object without an email, and optionally the email scope, which returns the object with an email.
func (s *UsersService) GetCurrentUser() (*User, *http.Response, error) {
	return s.getUser("@me")
}

//Returns a user for a given user ID.
func (s *UsersService) GetUser(userID Snowflake) (*User, *http.Response, error) {
	return s.getUser(userID.String())
}

func (s *UsersService) getUser(userID string) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%v", userID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}

//Modify the requestors user account settings.
//Returns a user object on success.
func (s *UsersService) ModifyCurrentUser(username string, avatar *AvatarData) (*User, *http.Response, error) {
	data := struct {
		Username string `json:"username,omitempty"`
		Avatar   string `json:"username,omitempty"`
	}{}
	if username != "" {
		data.Username = username
	}
	if avatar != nil {
		data.Avatar = avatar.toString()
	}

	req, err := s.client.NewRequest("PATCH", "users/@me", &data)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}

//Returns a list of user guild objects the current user is a member of.
//Requires the guilds OAuth2 scope.
func (s *UsersService) GetCurrentUserGuild() ([]*UserGuild, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "users/@me", nil)
	if err != nil {
		return nil, nil, err
	}

	var guilds []*UserGuild
	resp, err := s.client.Do(req, guilds)
	if err != nil {
		return nil, resp, err
	}

	return guilds, resp, err
}

//Leave a guild.
//Returns a 204 empty response on success.
func (s *UsersService) LeaveGuild(Snowflake string) (*http.Response, error) {
	u := fmt.Sprintf("users/@me/guilds/%v", Snowflake)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//Returns a list of DM channel objects.
func (s *UsersService) GetUserDMs() ([]*DMChannel, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "users/@me/channels", nil)
	if err != nil {
		return nil, nil, err
	}

	var channels []*DMChannel
	resp, err := s.client.Do(req, channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, err
}

//Create a new DM channel with a user.
//Returns a DM channel object.
func (s *UsersService) CreateDM(recipientID Snowflake) (*DMChannel, *http.Response, error) {
	data := struct {
		RecipientID Snowflake `json:"recipient_id,string"`
	}{
		RecipientID: recipientID,
	}
	req, err := s.client.NewRequest("POST", "users/@me/channels", &data)
	if err != nil {
		return nil, nil, err
	}

	channel := new(DMChannel)
	resp, err := s.client.Do(req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, err
}

//Returns a list of connection objects. Requires the connections OAuth2 scope.
func (s *UsersService) GetUsersConnections() ([]*Connection, *http.Response, error) {

	req, err := s.client.NewRequest("GET", "users/@me/connections", nil)
	if err != nil {
		return nil, nil, err
	}

	var conns []*Connection
	resp, err := s.client.Do(req, conns)
	if err != nil {
		return nil, resp, err
	}

	return conns, resp, err
}
