package plexwebhooks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
)

// Thumbnail contains all the relevant data about the thumbnail file (if sended).
type Thumbnail struct {
	Filename string
	Data     []byte
}

// Extract extracts the payload and the thumbnail (if present) from a multipart reader
func Extract(mpr *multipart.Reader) (payload *Payload, thumbnail *Thumbnail, err error) {
	if mpr == nil {
		err = errors.New("multipart reader can't be nil")
		return
	}
	// Read
	var formPart *multipart.Part
	for formPart, err = mpr.NextPart(); err == nil; formPart, err = mpr.NextPart() {
		switch formPart.FormName() {
		case "payload":
			// Only one payload
			if payload != nil {
				err = errors.New("payload part is present more than once")
				return
			}
			// Extract payload
			payload = new(Payload)
			decoder := json.NewDecoder(formPart)
			// decoder.DisallowUnknownFields() // dev
			if err = decoder.Decode(payload); err != nil {
				err = fmt.Errorf("payload JSON decode failed: %w", err)
				return
			}
		case "thumb":
			// Only one thumb can be present
			if thumbnail != nil {
				err = errors.New("thumb part is present more than once")
				return
			}
			// Prepare thumb event payload & set filename
			thumbnail = &Thumbnail{
				Filename: formPart.FileName(),
			}
			// Extract thumb data
			if thumbnail.Data, err = ioutil.ReadAll(formPart); err != nil {
				err = fmt.Errorf("error while reading thumb form part data: %w", err)
			}
		default:
			err = fmt.Errorf("unexpected form part encountered: %s", formPart.FormName())
			return
		}
	}
	// Handle errors
	if err == io.EOF {
		err = nil
	}
	if err == nil && payload == nil {
		err = errors.New("payload not found within request")
	}
	// Done
	return
}
