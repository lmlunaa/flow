package flow

import (
	"net/http"
)

var barrelPayload http.Handler

type Payload struct {
	Handler func(http.ResponseWriter, *http.Request)
}

type Stream func(http.Handler) http.Handler

// Barrel specifies handler to corresponding endpoint.
func Barrel(handler func(http.ResponseWriter, *http.Request)) *Payload {
	return &Payload{
		Handler: handler,
	}
}

// Thru will process Barrel through all Stream given by order.
func (p *Payload) Thru(streams ...Stream) http.Handler {
	barrelPayload = http.HandlerFunc(p.Handler)
	for i := len(streams) - 1; i >= 0; i-- {
		newBarrelPayload := UseStream(barrelPayload, streams[i])
		barrelPayload = newBarrelPayload
	}
	return barrelPayload
}

// useStream will pass Barrel to Stream given.
func UseStream(handler http.Handler, stream Stream) http.Handler {
	return stream(handler)
}
