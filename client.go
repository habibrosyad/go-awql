package awql

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// HTTP client to post and get AWQL query.
type Client struct {
	auth   *Auth
	client *http.Client
}

// Create a new client based on the provided auth configuration.
func NewClient(auth *Auth) *Client {
	oauth2Token := &oauth2.Token{
		AccessToken:  auth.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: auth.RefreshToken,
		Expiry:       time.Now(),
	}

	oauth2Config := &oauth2.Config{
		ClientID:     auth.ClientId,
		ClientSecret: auth.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/adwords",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	return &Client{auth, oauth2Config.Client(context.TODO(), oauth2Token)}
}

// Execute an AWQL query for a report.
func (c *Client) Query(q string) (Rows, error) {
	return newReportQuery(c, q)
}

// Execute an AWQL query for a service.
func (c *Client) QueryService(q, s string) (Rows, error) {
	return newServiceQuery(q, s)
}
