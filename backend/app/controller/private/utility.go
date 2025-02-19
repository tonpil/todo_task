package private

import (
	"errors"
	"fmt"
	"net/http"

	prjerror "ifigurin12/todo/interactor/error"
)

func TransformErrorToHttpError(err error) (int, string) {
	switch {
	case errors.Is(err, &prjerror.NotFoundError{}):
		return http.StatusNotFound, "Not found"
	case errors.Is(err, &prjerror.AlreadyExistsError{}):
		return http.StatusConflict, "Element already exists"
	default:
		fmt.Println(err)
		return http.StatusInternalServerError, "Internal server error"
	}
}
