package core

import (
	"fmt"
	"homesensor/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendReturnErrorUnhandledQuery(t *testing.T) {
	bus := NewQueryBus()
	queryType := "no_handler"
	result, err := bus.Send(queryType, "")

	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf("unhandled query: %s", queryType))
}

func TestSendReturnErrorFromHandler(t *testing.T) {
	bus := NewQueryBus()
	queryType := "test_query"
	bus.SetHandler(queryType, shared.QueryHandlerFunc(func(payload shared.Query) (shared.Result, error) {
		return nil, fmt.Errorf("handler returned error")
	}))

	result, err := bus.Send(queryType, "")

	assert.Nil(t, result)
	assert.EqualError(t, err, "handler returned error")
}

func TestSendReturnHandlerResult(t *testing.T) {
	bus := NewQueryBus()
	queryType := "test_query"

	bus.SetHandler(queryType, shared.QueryHandlerFunc(func(payload shared.Query) (shared.Result, error) {
		return "", nil
	}))

	result, err := bus.Send(queryType, "")

	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

func TestUnsetHandler(t *testing.T) {
	bus := NewQueryBus()
	queryType := "test_query"

	unset := bus.SetHandler(queryType, shared.QueryHandlerFunc(func(payload shared.Query) (shared.Result, error) {
		return "", nil
	}))

	result, err := bus.Send(queryType, "")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	unset()

	result, err = bus.Send(queryType, "")

	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf("unhandled query: %s", queryType))
}
