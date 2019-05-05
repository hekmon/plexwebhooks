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
	LibrarySectionType    string             `json:"librarySectionType"`    // movie + show + music
	RatingKey             string             `json:"ratingKey"`             // movie + show + music
	Key                   string             `json:"key"`                   // movie + show + music
	ParentRatingKey       string             `json:"parentRatingKey"`       // show + music
	GrandparentRatingKey  string             `json:"grandparentRatingKey"`  // show + music
	GUID                  string             `json:"guid"`                  // movie + show + music
	LibrarySectionTitle   string             `json:"librarySectionTitle"`   // movie + show + music
	LibrarySectionID      int                `json:"librarySectionID"`      // movie + show + music
	LibrarySectionKey     string             `json:"librarySectionKey"`     // movie + show + music
	Studio                string             `json:"studio"`                // movie
	Type                  string             `json:"type"`                  // movie + show + music
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
	LastViewedAt          int                `json:"lastViewedAt"`          // movie + show + music
	Year                  int                `json:"year"`                  // movie + show
	Tagline               string             `json:"tagline"`               // movie
	Thumb                 string             `json:"thumb"`                 // movie + show + music
	Art                   string             `json:"art"`                   // movie + show
	ParentThumb           string             `json:"parentThumb"`           // show + music
	ParentArt             string             `json:"parentArt"`             // show
	GrandparentThumb      string             `json:"grandparentThumb"`      // show + music
	GrandparentArt        string             `json:"grandparentArt"`        // show
	GrandparentTheme      string             `json:"grandparentTheme"`      // show
	Duration              int                `json:"duration"`              // movie
	OriginallyAvailableAt string             `json:"originallyAvailableAt"` // movie
	AddedAt               int                `json:"addedAt"`               // movie + show + music
	UpdatedAt             int                `json:"updatedAt"`             // movie + show + music
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
