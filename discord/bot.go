package discord

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func NewBotClient(ctx context.Context, token string) *Client {
	t := oauth2.Token{
		TokenType:   "Bot",
		AccessToken: token,
	}

	cli := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&t))
	r := newClient(cli)
	r.token = token
	r.bot = true
	return r
}

type ApplicationInfo struct {
	Description string        `json:"description"`
	Icon        *string       `json:"icon"`
	ID          Snowflake     `json:"id,string"`
	Name        string        `json:"name"`
	RPCOrigins  []interface{} `json:"rpc_origins"`
	Flags       uint64        `json:"flags"`
	Owner       struct {
		Username      string    `json:"username"`
		Discriminator string    `json:"discriminator"`
		ID            Snowflake `json:"id,string"`
		Avatar        *string   `json:"avatar"`
	} `json:"owner"`
}

func (c *Client) GetCurrentApplicationInfomation() (*ApplicationInfo, error) {
	req, err := c.NewRequest("GET", "oauth2/applications/@me", nil)
	if err != nil {
		return nil, err
	}
	info := new(ApplicationInfo)
	_, err = c.Do(req, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
