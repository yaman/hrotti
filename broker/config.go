package broker

import (
	"log"
	"net/url"
	"os"
)

//loggers
var (
	INFO     *log.Logger
	PROTOCOL *log.Logger
	ERROR    *log.Logger
	DEBUG    *log.Logger
)

func init() {
	INFO = log.New(os.Stderr, "", 0)
	PROTOCOL = log.New(os.Stderr, "", 0)
	ERROR = log.New(os.Stderr, "", 0)
	DEBUG = log.New(os.Stderr, "", 0)
}

type ListenerConfig struct {
	URL  *url.URL `json:"url"`
	stop chan struct{}
}

func NewListenerConfig(rawURL string) *ListenerConfig {
	listenerURL, err := url.Parse(rawURL)
	if err != nil {
		return nil
	}
	l := &ListenerConfig{URL: listenerURL}
	return l
}
