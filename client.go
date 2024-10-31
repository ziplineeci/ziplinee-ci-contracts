package contracts

import "time"

// Client represents a client application registered with Estafette
type Client struct {
	ID            string          `json:"id,omitempty"`
	Active        bool            `json:"active,omitempty"`
	Name          string          `json:"name,omitempty"`
	ClientID      string          `json:"clientID,omitempty"`
	ClientSecret  string          `json:"clientSecret,omitempty"`
	Roles         []*string       `json:"roles,omitempty"`
	Organizations []*Organization `json:"organizations,omitempty"`
	CreatedAt     *time.Time      `json:"createdAt,omitempty"`
}
