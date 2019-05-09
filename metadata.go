package plexwebhooks

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

/*
	Metadata
*/

// Metadata represents all the metadata associated with a media sent by the webhook
type Metadata struct {
	LibrarySectionType    LibrarySection     `json:"librarySectionType"`    // movie + show + music
	RatingKey             string             `json:"ratingKey"`             // movie + show + music
	Key                   string             `json:"key"`                   // movie + show + music
	ParentRatingKey       string             `json:"parentRatingKey"`       // show + music
	GrandparentRatingKey  string             `json:"grandparentRatingKey"`  // show + music
	GUID                  *url.URL           `json:"guid"`                  // movie + show + music
	LibrarySectionTitle   string             `json:"librarySectionTitle"`   // movie + show + music
	LibrarySectionID      int                `json:"librarySectionID"`      // movie + show + music
	LibrarySectionKey     string             `json:"librarySectionKey"`     // movie + show + music
	Studio                string             `json:"studio"`                // movie
	Type                  MediaType          `json:"type"`                  // movie + show + music
	Title                 string             `json:"title"`                 // movie + show + music
	TitleSort             string             `json:"titleSort"`             // show (should be movie too)
	GrandparentKey        string             `json:"grandparentKey"`        // music
	ParentKey             string             `json:"parentKey"`             // music
	GrandparentTitle      string             `json:"grandparentTitle"`      // music
	ParentTitle           string             `json:"parentTitle"`           // music
	OriginalTitle         string             `json:"originalTitle"`         // music
	ContentRating         string             `json:"contentRating"`         // movie + show
	Summary               string             `json:"summary"`               // movie + show + music
	Index                 int                `json:"index"`                 // show + music
	ParentIndex           int                `json:"parentIndex"`           // show + music
	Rating                float64            `json:"rating"`                // movie + show
	RatingCount           int                `json:"ratingCount"`           // music
	ViewCount             int                `json:"viewCount"`             // movie + show + music
	LastViewedAt          time.Time          `json:"lastViewedAt"`          // movie + show + music
	Year                  int                `json:"year"`                  // movie + show
	Tagline               string             `json:"tagline"`               // movie
	Thumb                 string             `json:"thumb"`                 // movie + show + music
	Art                   string             `json:"art"`                   // movie + show
	ParentThumb           string             `json:"parentThumb"`           // show + music
	ParentArt             string             `json:"parentArt"`             // show
	GrandparentThumb      string             `json:"grandparentThumb"`      // show + music
	GrandparentArt        string             `json:"grandparentArt"`        // show
	GrandparentTheme      string             `json:"grandparentTheme"`      // show
	Duration              time.Duration      `json:"duration"`              // movie
	OriginallyAvailableAt time.Time          `json:"originallyAvailableAt"` // movie
	AddedAt               time.Time          `json:"addedAt"`               // movie + show + music
	UpdatedAt             time.Time          `json:"updatedAt"`             // movie + show + music
	PrimaryExtraKey       string             `json:"primaryExtraKey"`       // movie
	RatingImage           string             `json:"ratingImage"`           // movie
	ChapterSource         string             `json:"chapterSource"`         // show (movie too ?)
	Genre                 []MetadataItem     `json:"Genre"`                 // movie
	Director              []MetadataItem     `json:"Director"`              // movie + show
	Writer                []MetadataItem     `json:"Writer"`                // movie + show
	Producer              []MetadataItem     `json:"Producer"`              // movie
	Country               []MetadataItem     `json:"Country"`               // movie
	Collection            []MetadataItem     `json:"Collection"`            // movie
	Role                  []MetadataItemRole `json:"Role"`                  // movie
	Similar               []MetadataItem     `json:"Similar"`               // movie
	Mood                  []MetadataItem     `json:"Mood"`                  // music
}

// UnmarshalJSON allows to convert json values to Go types
func (m *Metadata) UnmarshalJSON(data []byte) (err error) {
	// Prepare the catcher
	type Shadow Metadata
	tmp := struct {
		GUID                  string `json:"guid"`
		LastViewedAt          int64  `json:"lastViewedAt"`
		Duration              int64  `json:"duration"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		AddedAt               int64  `json:"addedAt"`
		UpdatedAt             int64  `json:"updatedAt"`
		*Shadow
	}{
		Shadow: (*Shadow)(m),
	}
	// Unmarshal within the catcher
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	// Use catcher values to build the golang one
	if m.GUID, err = url.Parse(tmp.GUID); err != nil {
		return fmt.Errorf("can't convert GUID string as URL: %v", err)
	}
	m.LastViewedAt = time.Unix(tmp.LastViewedAt, 0)
	m.Duration = time.Duration(tmp.Duration) * time.Millisecond
	if tmp.OriginallyAvailableAt != "" {
		if m.OriginallyAvailableAt, err = time.Parse("2006-01-02", tmp.OriginallyAvailableAt); err != nil {
			return fmt.Errorf("can't parse 'OriginallyAvailableAt' as time.Time: %v", err)
		}
	}
	m.AddedAt = time.Unix(tmp.AddedAt, 0)
	m.UpdatedAt = time.Unix(tmp.UpdatedAt, 0)
	return
}

/*
	Metadata custom
*/

// LibrarySection represents the of the library section
type LibrarySection string

const (
	// LibrarySectionShow represents the shows library type
	LibrarySectionShow LibrarySection = "show"
	// LibrarySectionMusic represents the music library type
	LibrarySectionMusic LibrarySection = "artist"
	// LibrarySectionMovie represents the movies library type
	LibrarySectionMovie LibrarySection = "movie"
)

// MediaType represente the type of media related to the webhook event
type MediaType string

const (
	// MediaTypeMovie represents the media type for a movie
	MediaTypeMovie MediaType = "movie"
	// MediaTypeEpisode represents the media type for a show episode
	MediaTypeEpisode MediaType = "episode"
	// MediaTypeTrack represents the media type for an audio track
	MediaTypeTrack MediaType = "track"
)

/*
	Metadata items
*/

// MetadataItem represents a ref to an external entity
type MetadataItem struct {
	ID     int    `json:"id"`
	Filter string `json:"filter"`
	Tag    string `json:"tag"`
	Count  int    `json:"count"`
}

// MetadataItemRole is a specialisation for roles of MetadataItem
type MetadataItemRole struct {
	MetadataItem
	Role  string `json:"role"`
	Thumb *url.URL
}

// UnmarshalJSON allows to convert json values to Go types
func (mir *MetadataItemRole) UnmarshalJSON(data []byte) (err error) {
	type Shadow MetadataItemRole
	tmp := struct {
		Thumb string `json:"thumb"`
		*Shadow
	}{
		Shadow: (*Shadow)(mir),
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	if tmp.Thumb != "" {
		if mir.Thumb, err = url.Parse(tmp.Thumb); err != nil {
			err = fmt.Errorf("can't parse MetadataItemRole thumb as URL: %v", err)
		}
	}
	return
}
