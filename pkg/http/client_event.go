package http

import (
	"fmt"

	"github.com/saniyar-dev/xk6-new-http/pkg/events"
	"github.com/saniyar-dev/xk6-new-http/pkg/helpers"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
)

const (
	OPEN     = "open"
	CLOSE    = "close"
	ERROR    = "error"
	RESPONSE = "responseReceived"
	REQUEST  = "requestToBeSent"
)

type eventListeners struct {
	open     *events.EventListener
	close    *events.EventListener
	error    *events.EventListener
	response *events.EventListener
	request  *events.EventListener
}

var _ interfaces.EventListeners = &eventListeners{}

func (es *eventListeners) GetListener(t string) (*events.EventListener, error) {
	switch t {
	case OPEN:
		return es.open, nil
	case CLOSE:
		return es.close, nil
	case ERROR:
		return es.error, nil
	case RESPONSE:
		return es.response, nil
	case REQUEST:
		return es.request, nil
	default:
		return nil, fmt.Errorf("unsupported event type for client %s", t)
	}
}

func (es *eventListeners) New() interfaces.EventListeners {
	return &eventListeners{
		open:     helpers.NewEventListener(OPEN),
		close:    helpers.NewEventListener(CLOSE),
		error:    helpers.NewEventListener(ERROR),
		response: helpers.NewEventListener(RESPONSE),
		request:  helpers.NewEventListener(REQUEST),
	}
}
