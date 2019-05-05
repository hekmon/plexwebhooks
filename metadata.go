package plexwebhook

import (
	"fmt"
	"net/url"
)

/*
	Metadata
*/

// Metadata represents all the metadata associated with a media sent by the webhook
type Metadata struct {
	LibrarySectionType    string             `json:"librarySectionType"`    // movie + show
	RatingKey             string             `json:"ratingKey"`             // movie + show
	Key                   string             `json:"key"`                   // movie + show
	ParentRatingKey       string             `json:"parentRatingKey"`       // show
	GrandparentRatingKey  string             `json:"grandparentRatingKey"`  // show
	GUID                  string             `json:"guid"`                  // movie + show
	LibrarySectionTitle   string             `json:"librarySectionTitle"`   // movie + show
	LibrarySectionID      int                `json:"librarySectionID"`      // movie + show
	LibrarySectionKey     string             `json:"librarySectionKey"`     // movie + show
	Studio                string             `json:"studio"`                // movie
	Type                  string             `json:"type"`                  // movie + show
	Title                 string             `json:"title"`                 // movie + show
	TitleSort             string             `json:"titleSort"`             // show (should be movie too)
	ContentRating         string             `json:"contentRating"`         // movie + show
	Summary               string             `json:"summary"`               // movie + show
	Rating                float64            `json:"rating"`                // movie + show
	ViewCount             int                `json:"viewCount"`             // movie + show
	LastViewedAt          int                `json:"lastViewedAt"`          // movie + show
	Year                  int                `json:"year"`                  // movie + show
	Tagline               string             `json:"tagline"`               // movie
	Thumb                 string             `json:"thumb"`                 // movie + show
	Art                   string             `json:"art"`                   // movie + show
	ParentThumb           string             `json:"parentThumb"`           // show
	ParentArt             string             `json:"parentArt"`             // show
	GrandparentThumb      string             `json:"grandparentThumb"`      // show
	GrandparentArt        string             `json:"grandparentArt"`        // show
	GrandparentTheme      string             `json:"grandparentTheme"`      // show
	Duration              int                `json:"duration"`              // movie
	OriginallyAvailableAt string             `json:"originallyAvailableAt"` // movie
	AddedAt               int                `json:"addedAt"`               // movie + show
	UpdatedAt             int                `json:"updatedAt"`             // movie + show
	PrimaryExtraKey       string             `json:"primaryExtraKey"`       // movie
	RatingImage           string             `json:"ratingImage"`           // movie
	ChapterSource         string             `json:"chapterSource"`         // show
	Genre                 []MetadataItem     `json:"Genre"`                 // movie
	Director              []MetadataItem     `json:"Director"`              // movie + show
	Writer                []MetadataItem     `json:"Writer"`                // movie + show
	Producer              []MetadataItem     `json:"Producer"`              // movie
	Country               []MetadataItem     `json:"Country"`               // movie
	Collection            []MetadataItem     `json:"Collection"`            // movie
	Role                  []MetadataItemRole `json:"Role"`                  // movie
	Similar               []MetadataItem     `json:"Similar"`               // movie
}

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
	if mir.Thumb, err = url.Parse(tmp.Thumb); err != nil {
		err = fmt.Errorf("can't parse MetadataItemRole thumb as URL: %v", err)
	}
	return
}
