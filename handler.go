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
	Error   error
}

// EventFile contains all the relevent data about the thumb file (if sended).
type EventFile struct {
	Filename string
	Data     []byte
}

// HTTPHandler yield a valid HTTP handler to receive HTTP multi part form from Plex webhooks.
// It will send extracted information as an Event on the eventChan.
func HTTPHandler(eventChan chan<- Event) http.HandlerFunc {
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
			eventChan <- Event{Error: err}
			return
		}
		// Read parts
		var (
			formPart *multipart.Part
			event    Event
		)
		for formPart, err = multiPartReader.NextPart(); err == nil; formPart, err = multiPartReader.NextPart() {
			switch formPart.FormName() {
			case "payload":
				// Only one payload
				if event.Payload != nil {
					//TODO
				}
				// Extract payload
				event.Payload = new(Payload)
				if err = json.NewDecoder(formPart).Decode(event.Payload); err != nil {
					eventError(&event, fmt.Errorf("payload JSON decode failed: %v", err))
					break
				}
			case "thumb":
				// Only one thumb can be present
				if event.Thumb != nil {
					eventError(&event, errors.New("thumb part is present more than once"))
					break
				}
				// Prepare thumb event payload & set filename
				event.Thumb = &EventFile{
					Filename: formPart.FileName(),
				}
				// Extract thumb data
				if event.Thumb.Data, err = ioutil.ReadAll(formPart); err != nil {
					eventError(&event, fmt.Errorf("error while reading thumb form part data: %v", err))
					break
				}
			default:
				eventError(&event, fmt.Errorf("unexpected form part encountered: %s", formPart.FormName()))
				break
			}
		}
		// Handle multi part errors
		if event.Error == nil && err != io.EOF {
			eventError(&event, fmt.Errorf("moving to next multi form part failed: %s", formPart.FormName()))
		}
		// Send event
		eventChan <- event
		// Prepare clean http close
		w.WriteHeader(http.StatusOK)
	}
}

func eventError(event *Event, err error) {
	event.Payload = nil
	event.Thumb = nil
	event.Error = err
}
