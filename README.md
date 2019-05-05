# Plex Webhooks

[![GoDoc](https://godoc.org/github.com/hekmon/plexwebhooks?status.svg)](https://godoc.org/github.com/hekmon/plexwebhooks) [![Go Report Card](https://goreportcard.com/badge/github.com/hekmon/plexwebhooks)](https://goreportcard.com/report/github.com/hekmon/plexwebhooks)

Golang binding for [Plex Webhooks](https://support.plex.tv/articles/115002267687-webhooks/).

This library provides:

* Golang binding for webhook (JSON) payloads
* Auto convert values to Golang types when possible (time, duration, IP, URL, etc...)
* HTTP handler with post multi form compatibility to extract both payload and attached thumbnail when sent (play & rate events)

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
    http.HandleFunc("/", plexwebhooks.HTTPHandler(process))
    http.ListenAndServe(":7095", http.DefaultServeMux)
}

func process(event *plexwebhooks.Event, err error) {
    fmt.Println()
    fmt.Println(time.Now())
    if err == nil {
        fmt.Printf("%+v\n", *event.Payload)
        if event.Thumb != nil {
            fmt.Printf("Name: %s | Size: %d\n", event.Thumb.Filename, len(event.Thumb.Data))
        }
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

```
