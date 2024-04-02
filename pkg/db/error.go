package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func HandleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrObjectNotFound{}
	} else if errors.Is(err, sql.ErrConnDone) {
		return ErrConnDone{}
	}
	return err
}

type ErrObjectNotFound struct{}

func (ErrObjectNotFound) Error() string {
	return "object not found"
}
func (ErrObjectNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}

type ErrConnDone struct{}

func (ErrConnDone) Error() string {
	return "connection done"
}
func (ErrConnDone) Unwrap() error {
	return fmt.Errorf("connection done")
}
