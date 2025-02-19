package private

import (
	prjerror "ifigurin12/todo/interactor/error"

	"github.com/jackc/pgx/v5"
)

func TransformError(err error) error {
	if err == pgx.ErrNoRows {
		return &prjerror.NotFoundError{}
	} else {
		return err
	}
}
