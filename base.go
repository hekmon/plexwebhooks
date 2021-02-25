package plexwebhooks

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"
)

// Payload represents the base structure for all plex webhooks
type Payload struct {
	Rating   int       // only present for Event == EventTypeRate
	Event    EventType `json:"event"`
	User     bool      `json:"user"`
	Owner    bool      `json:"owner"`
	Account  Account   `json:"Account"`
	Server   Server    `json:"Server"`
	Player   Player    `json:"Player"`
	Metadata Metadata  `json:"Metadata"`
}

// UnmarshalJSON allows to convert json values to Go types
func (p *Payload) UnmarshalJSON(data []byte) (err error) {
	type Shadow Payload
	tmp := struct {
		Rating string `json:"rating"`
		*Shadow
	}{
		Shadow: (*Shadow)(p),
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	if tmp.Event == EventTypeRate {
		if p.Rating, err = strconv.Atoi(tmp.Rating); err != nil {
			err = fmt.Errorf("can't convert rating as int: %w", err)
		}
	}
	return
}

/*
	Event type
*/

// EventType represent the webhook event type
type EventType string

const (
	// EventTypePause represents the play event
	EventTypePause EventType = "media.pause"
	// EventTypePlay represents the play event
	EventTypePlay EventType = "media.play"
	// EventTypeRate represents the rate event
	EventTypeRate EventType = "media.rate"
	// EventTypeResume represents the resume event
	EventTypeResume EventType = "media.resume"
	// EventTypeScrobble represents the scrobble event
	EventTypeScrobble EventType = "media.scrobble"
	// EventTypeStop represents the stop event
	EventTypeStop EventType = "media.stop"
)

/*
	Account
*/

// Account represent the account which has generated the wehhook launch on the server
type Account struct {
	ID    int `json:"id"` // server relative id (owner is 1)
	Thumb *url.URL
	Title string `json:"title"` // username
}

// UnmarshalJSON allows to convert json values to Go types
func (a *Account) UnmarshalJSON(data []byte) (err error) {
	type Shadow Account
	tmp := struct {
		Thumb string `json:"thumb"`
		*Shadow
	}{
		Shadow: (*Shadow)(a),
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	if a.Thumb, err = url.Parse(tmp.Thumb); err != nil {
		err = fmt.Errorf("can't parse account thumb as URL: %w", err)
	}
	return
}

/*
	Server
*/

// Server holds informations about the server which generated the webhook
type Server struct {
	Title string `json:"title"`
	UUID  string `json:"uuid"`
}

/*
	Player
*/

// Player holds informations about the user's player
type Player struct {
	Local         bool `json:"local"`
	PublicAddress net.IP
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
}

// UnmarshalJSON allows to convert json values to Go types
func (p *Player) UnmarshalJSON(data []byte) (err error) {
	type Shadow Player
	tmp := struct {
		PublicAddress string `json:"publicAddress"`
		*Shadow
	}{
		Shadow: (*Shadow)(p),
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	p.PublicAddress = net.ParseIP(tmp.PublicAddress)
	return
}
