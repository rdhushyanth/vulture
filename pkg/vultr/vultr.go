package vultr

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func getAPIKey() string {
	return os.Getenv("VULTR_API_KEY")
}

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type Account struct {
	Balance           float32   `json:"balance,omitempty"`
	PendingCharges    float32   `json:"pending_charges,omitempty"`
	LastPaymentDate   time.Time `json:"last_payment_date,omitempty"`
	LastPaymentAmount float32   `json:"last_payment_amount,omitempty"`
}

func (c *Client) AccountInfo() (*Account, error) {
	rel := &url.URL{Path: "v1/account/info"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("API-Key", getAPIKey())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var account Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		fmt.Println(resp.Body)
		fmt.Println(req)
		fmt.Println(err)
	}
	return &account, err
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient}
	return c
}
