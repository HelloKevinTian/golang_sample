package routes

import (
	"errors"
)

type Handler func(ctx interface{}, params interface{}) (string, interface{})

var routes = map[string]Handler{}

func Add(protocol string, handler Handler) {
	routes[protocol] = handler
}

func Route(protocol string) (Handler, error) {
	handler, ok := routes[protocol]
	if ok {
		return handler, nil
	} else {
		return handler, errors.New("Router not found!")
	}
}
