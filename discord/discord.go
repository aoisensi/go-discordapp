package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://discordapp.com/api/"
	userAgent      = "DiscordBot (\"http://github.com/aoisensi/go-discordapp\", \"dev\")"
)

type Status string

const (
	StatusIdle    Status = "idle"
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(`"`+time.RFC3339Nano+`"`, string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(ts)
	return nil
}

type Unixtime time.Time

func (t *Unixtime) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	*t = Unixtime(time.Unix(int64(ts), 0))
	return nil
}

type Client struct {
	client    *http.Client
	bot       bool
	token     string
	BaseURL   *url.URL
	UserAgent string
}

func (c *Client) Channel(ChannelID Snowflake) *ChannelsService {
	return &ChannelsService{client: c, ChannelID: ChannelID}
}

func (c *Client) Guild(GuildID Snowflake) *GuildsService {
	return &GuildsService{client: c, GuildID: GuildID}
}

func (c *Client) User() *UsersService {
	return &UsersService{client: c}
}

func newClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if body == nil {
		return c.newRequest(method, urlStr, nil, "")
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}
	return c.newRequest(method, urlStr, buf, "application/json")
}

func (c *Client) newRequest(method, urlStr string, body io.Reader, mime string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if mime != "" {
		req.Header.Add("Content-Type", mime)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)

	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	return resp, json.Unmarshal(buf.Bytes(), v)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Code     int            `json:"code"`
	Message  string         `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %d %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Code, r.Message)
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
