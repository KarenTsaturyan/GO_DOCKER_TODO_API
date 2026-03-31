package core_http_types

import (
	"encoding/json"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	n.Set = true

	if string(data) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}

/*
{}
    -value *nil
    -set false

{"value": null}
    -value *nil
    -set true

{"value": "some value"}
    -value *"some value"
    -set true
*/
