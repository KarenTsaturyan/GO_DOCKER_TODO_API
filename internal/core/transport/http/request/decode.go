package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	var err error

	if v, ok := dest.(validatable); ok {
		// Custom validation
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("request validation: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
