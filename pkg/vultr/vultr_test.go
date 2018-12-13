package vultr

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestClient(t *testing.T) {
	u, err := url.Parse("https://api.vultr.com/")
	if err != nil {
		log.Fatal(err)
	}
	c := &Client{httpClient: http.DefaultClient, BaseURL: u, UserAgent: "My-client"}
	acc, err := c.AccountInfo()
	if err != nil {
		log.Fatal(err)
		t.Errorf("No Error Expected")
	}
	fmt.Println(acc)
}
