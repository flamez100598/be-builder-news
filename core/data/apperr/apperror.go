package apperror
import (
	"errors"
)

var (
	ErrorNotFound = errors.New("error not found")
	ErrNoRows = errors.New("error no rows")
)