package vultr

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	info = log.New(os.Stderr,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	debug = log.New(os.Stderr,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	baseURL, _ = url.Parse("https://api.vultr.com/")
)

// SetLogger specifies a custom logger for the library
func SetLogger(info, debug *log.Logger) {
	info = info
	debug = debug
}

func getAPIKey() string {
	return os.Getenv("VULTR_API_KEY")
}

// Client for use in library
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// AccountInfo returns the account details
func (c *Client) AccountInfo() (*Account, error) {
	rel := &url.URL{Path: "v1/account/info"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {

		info.Println(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("API-Key", getAPIKey())

	resp, err := c.httpClient.Do(req)

	if err != nil {

		info.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var account Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		info.Println(err)
	}
	return &account, err
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient, UserAgent: "rdhushyanth-vulture", BaseURL: baseURL}
	return c
}
