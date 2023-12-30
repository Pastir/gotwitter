package gotwitter

type Events interface {
	EventMethod() string
	EventName() string
	EventEndpoint() string
	EventHeader() map[string]string
}
