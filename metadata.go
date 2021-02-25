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
	AddedAt               time.Time          `json:"addedAt"`               // movie + show + music
	Art                   string             `json:"art"`                   // movie + show
	AttributionLogo       *url.URL           `json:"attributionLogo"`       // streaming movie
	AudienceRating        float64            `json:"audienceRating"`        // movie
	AudienceRatingImage   string             `json:"audienceRatingImage"`   // movie
	Banner                *url.URL           `json:"banner"`                // movie
	ChapterSource         string             `json:"chapterSource"`         // show (movie too ?)
	Collection            []MetadataItem     `json:"Collection"`            // movie
	ContentRating         string             `json:"contentRating"`         // movie + show
	Country               []MetadataItem     `json:"Country"`               // movie
	CreatedAtAccuracy     string             `json:"createdAtAccuracy"`     // photo
	CreatedAtTZOffset     string             `json:"createdAtTZOffset"`     // photo
	Director              []MetadataItem     `json:"Director"`              // movie + show
	Duration              time.Duration      `json:"duration"`              // movie
	Genre                 []MetadataItem     `json:"Genre"`                 // movie
	GenuineMediaAnalysis  string             `json:"genuineMediaAnalysis"`  // show
	GrandparentArt        string             `json:"grandparentArt"`        // show
	GrandparentKey        string             `json:"grandparentKey"`        // music
	GrandparentRatingKey  string             `json:"grandparentRatingKey"`  // show + music
	GrandparentTheme      string             `json:"grandparentTheme"`      // show
	GrandparentThumb      string             `json:"grandparentThumb"`      // show + music
	GrandparentTitle      string             `json:"grandparentTitle"`      // music
	GUID                  *url.URL           `json:"guid"`                  // movie + show + music
	Index                 int                `json:"index"`                 // show + music
	Indirect              bool               `json:"indirect"`              // movie
	Key                   string             `json:"key"`                   // movie + show + music
	LastRatedAt           time.Time          `json:"lastRatedAt"`           // movie + show + music
	LastViewedAt          time.Time          `json:"lastViewedAt"`          // movie + show + music
	LibrarySectionID      int                `json:"librarySectionID"`      // movie + show + music
	LibrarySectionKey     string             `json:"librarySectionKey"`     // movie + show + music
	LibrarySectionTitle   string             `json:"librarySectionTitle"`   // movie + show + music
	LibrarySectionType    LibrarySection     `json:"librarySectionType"`    // movie + show + music
	Live                  string             `json:"live"`                  // show
	Mood                  []MetadataItem     `json:"Mood"`                  // music
	OriginallyAvailableAt time.Time          `json:"originallyAvailableAt"` // movie
	OriginalTitle         string             `json:"originalTitle"`         // music
	ParentArt             string             `json:"parentArt"`             // show
	ParentIndex           int                `json:"parentIndex"`           // show + music
	ParentKey             string             `json:"parentKey"`             // music
	ParentRatingKey       string             `json:"parentRatingKey"`       // show + music
	ParentThumb           string             `json:"parentThumb"`           // show + music
	ParentTitle           string             `json:"parentTitle"`           // music
	PrimaryExtraKey       string             `json:"primaryExtraKey"`       // movie
	Producer              []MetadataItem     `json:"Producer"`              // movie
	Rating                float64            `json:"rating"`                // movie + show
	RatingCount           int                `json:"ratingCount"`           // music
	RatingImage           string             `json:"ratingImage"`           // movie
	RatingKey             string             `json:"ratingKey"`             // movie + show + music
	Role                  []MetadataItemRole `json:"Role"`                  // movie
	Similar               []MetadataItem     `json:"Similar"`               // movie
	Studio                string             `json:"studio"`                // movie
	SubType               MediaSubType       `json:"subtype"`               // clip
	Summary               string             `json:"summary"`               // movie + show + music
	Tagline               string             `json:"tagline"`               // movie
	Thumb                 string             `json:"thumb"`                 // movie + show + music
	Title                 string             `json:"title"`                 // movie + show + music
	TitleSort             string             `json:"titleSort"`             // show (should be movie too)
	Type                  MediaType          `json:"type"`                  // movie + show + music
	UpdatedAt             time.Time          `json:"updatedAt"`             // movie + show + music
	UserRating            float64            `json:"userRating"`            // movie
	ViewCount             int                `json:"viewCount"`             // movie + show + music
	ViewOffset            int                `json:"viewOffset"`            // movie
	Writer                []MetadataItem     `json:"Writer"`                // movie + show
	Year                  int                `json:"year"`                  // movie + show
}

// UnmarshalJSON allows to convert json values to Go types
func (m *Metadata) UnmarshalJSON(data []byte) (err error) {
	// Prepare the catcher
	type Shadow Metadata
	tmp := struct {
		AddedAt               int64  `json:"addedAt"`
		AttributionLogo       string `json:"attributionLogo"`
		Banner                string `json:"banner"`
		Duration              int64  `json:"duration"`
		GUID                  string `json:"guid"`
		LastRatedAt           int64  `json:"lastRatedAt"`
		LastViewedAt          int64  `json:"lastViewedAt"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
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
	m.AddedAt = time.Unix(tmp.AddedAt, 0)
	if tmp.AttributionLogo != "" {
		if m.AttributionLogo, err = url.Parse(tmp.AttributionLogo); err != nil {
			return fmt.Errorf("can not convert AttributionLogo string as URL: %w", err)
		}
	}
	if tmp.Banner != "" {
		if m.Banner, err = url.Parse(tmp.Banner); err != nil {
			return fmt.Errorf("can not convert Banner string as URL: %w", err)
		}
	}
	m.Duration = time.Duration(tmp.Duration) * time.Millisecond
	if m.GUID, err = url.Parse(tmp.GUID); err != nil {
		return fmt.Errorf("can not convert GUID string as URL: %w", err)
	}
	m.LastRatedAt = time.Unix(tmp.LastRatedAt, 0)
	m.LastViewedAt = time.Unix(tmp.LastViewedAt, 0)
	if tmp.OriginallyAvailableAt != "" {
		if m.OriginallyAvailableAt, err = time.Parse("2006-01-02", tmp.OriginallyAvailableAt); err != nil {
			return fmt.Errorf("can't parse 'OriginallyAvailableAt' as time.Time: %w", err)
		}
	}
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
	// MediaTypeClip represents the media type for a clip
	MediaTypeClip MediaType = "clip"
	// MediaTypePhoto represents the media type for a photo
	MediaTypePhoto MediaType = "photo"
)

// MediaSubType represente the subtype of media related to the webhook event
type MediaSubType string

const (
	// MediaSubTypeMovie represents the media sub type for a trailer
	MediaSubTypeMovie MediaSubType = "trailer"
	// MediaSubTypeBehindTheScenes represents the media sub type for a behind the scenes clip
	MediaSubTypeBehindTheScenes MediaSubType = "behindTheScenes"
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
			err = fmt.Errorf("can't parse MetadataItemRole thumb as URL: %w", err)
		}
	}
	return
}
