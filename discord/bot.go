package discord

import "net/http"

func NewBotClient(token string) *http.Client {
	client := http.DefaultClient
	t := new(tripper)
	t.token = token
	client.Transport = t
	return client
}

type tripper struct {
	token string
}

func (t *tripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bot "+t.token)
	return http.DefaultTransport.RoundTrip(req)
}

type ApplicationInfo struct {
	Description string        `json:"description"`
	Icon        *string       `json:"icon"`
	ID          Snowflake     `json:"id"`
	Name        string        `json:"name"`
	RPCOrigins  []interface{} `json:"rpc_origins"`
	Flags       uint64        `json:"flags"`
	Owner       struct {
		Username      string    `json:"username"`
		Discriminator string    `json:"discriminator"`
		ID            Snowflake `json:"id"`
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
