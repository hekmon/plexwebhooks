# Plex Webhooks

[![Go Reference](https://pkg.go.dev/badge/github.com/hekmon/plexwebhooks.svg)](https://pkg.go.dev/github.com/hekmon/plexwebhooks) [![Go Report Card](https://goreportcard.com/badge/github.com/hekmon/plexwebhooks)](https://goreportcard.com/report/github.com/hekmon/plexwebhooks)

Golang binding for [Plex Webhooks](https://support.plex.tv/articles/115002267687-webhooks/).

This library provides:

* Golang binding for webhook (JSON) payloads
* Auto convert values to Golang types when possible (time, duration, IP, URL, etc...)
* Multipart reader extractor which returns the hook payload and the thumbnail if present

## Example

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/hekmon/plexwebhooks"
)

func main() {
    http.HandleFunc("/", processHandler)
    http.ListenAndServe(":7095", http.DefaultServeMux)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    // Create the multi part reader
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
        // Log the error
        fmt.Println("can't create a multipart reader from request:", err)
        return
    }
    // Use the multipart reader to parse the request body
    payload, thumb, err := plexwebhooks.Extract(multiPartReader)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        // Try to write the error as http body
        _, wErr := w.Write([]byte(err.Error()))
        if wErr != nil {
            err = fmt.Errorf("request error: %v | write error: %v", err, wErr)
        }
        // Log the error
        fmt.Println("can't create a multipart reader from request:", err)
        return
    }
    // Do something
    fmt.Println()
    fmt.Println(time.Now())
    fmt.Printf("%+v\n", *payload)
    if thumb != nil {
        fmt.Printf("Name: %s | Size: %d\n", thumb.Filename, len(thumb.Data))
    }
    fmt.Println()
}

```
