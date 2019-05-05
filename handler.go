package plexwebhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// Event contains all the information extract from the webhook.
// Start by checking the error.
type Event struct {
	Payload *Payload
	Thumb   *EventFile
}

// EventFile contains all the relevant data about the thumb file (if sended).
type EventFile struct {
	Filename string
	Data     []byte
}

// HTTPHandler yield a valid HTTP handler to receive HTTP multi part form from Plex webhooks.
// It will send extracted information as an Event on the eventChan.
func HTTPHandler(process func(event *Event, err error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare to stream the multi part form body
		defer r.Body.Close()
		multiPartReader, err := r.MultipartReader()
		if err != nil {
			// Detect error type for the http answer
			if err == http.ErrNotMultipart || err == http.ErrMissingBoundary {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			// Try to write the error as http body
			_, wErr := w.Write([]byte(err.Error()))
			if wErr != nil {
				err = fmt.Errorf("request error: %v | write error: %v", err, wErr)
			}
			// send the eventwith the error(s)
			process(nil, err)
			return
		}
		// Read parts
		event := new(Event)
		var formPart *multipart.Part
		for formPart, err = multiPartReader.NextPart(); err == nil; formPart, err = multiPartReader.NextPart() {
			if err = handleFormPart(event, formPart); err != nil {
				break
			}
		}
		// Handle errors
		if err == io.EOF {
			err = nil
		}
		if err == nil && event.Payload == nil {
			err = errors.New("payload not found within request")
		}
		if err != nil {
			event = nil
		}
		// Send event
		process(event, err)
		// Prepare clean http close
		w.WriteHeader(http.StatusOK)
	}
}

func handleFormPart(event *Event, formPart *multipart.Part) (err error) {
	switch formPart.FormName() {
	case "payload":
		// Only one payload
		if event.Payload != nil {
			return errors.New("payload part is present more than once")
		}
		// Extract payload
		event.Payload = new(Payload)
		decoder := json.NewDecoder(formPart)
		// decoder.DisallowUnknownFields() // dev
		if err = decoder.Decode(event.Payload); err != nil {
			return fmt.Errorf("payload JSON decode failed: %v", err)
		}
	case "thumb":
		// Only one thumb can be present
		if event.Thumb != nil {
			return errors.New("thumb part is present more than once")
		}
		// Prepare thumb event payload & set filename
		event.Thumb = &EventFile{
			Filename: formPart.FileName(),
		}
		// Extract thumb data
		if event.Thumb.Data, err = ioutil.ReadAll(formPart); err != nil {
			return fmt.Errorf("error while reading thumb form part data: %v", err)
		}
	default:
		return fmt.Errorf("unexpected form part encountered: %s", formPart.FormName())
	}
	return
}
