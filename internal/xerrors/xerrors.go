/**
 * Package xerrors
 * @file      : xerrors.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 15:24
 **/

package xerrors

import (
	"fmt"
	"net/http"
)

var (
	ErrorTypeUnknown        = ErrorType{"UNKNOWN_ERROR", http.StatusInternalServerError}
	ErrorTypeAuthorization  = ErrorType{"AUTHORIZATION_ERROR", http.StatusUnauthorized}
	ErrorTypeIncorrectInput = ErrorType{"INCORRECT_INPUT", http.StatusBadRequest}
	ErrorTypeInternal       = ErrorType{"INTERNAL_ERROR", http.StatusInternalServerError}
	ErrorTypeNotFound       = ErrorType{"NOT_FOUND", http.StatusNotFound}
)

type ErrorType struct {
	str    string
	status int
}

func (e ErrorType) String() string {
	return e.str
}

func (e ErrorType) Status() int {
	return e.status
}

type Error struct {
	err string
	typ ErrorType
}

func (e Error) Error() string {
	return e.err
}

func (e Error) ErrorType() ErrorType {
	return e.typ
}

func New(typ string, httpStatusCode int, msg ...string) Error {
	return Error{
		err: fmt.Sprintln(msg),
		typ: ErrorType{typ, http.StatusInternalServerError},
	}
}

func NewAuthorizationError(err string) Error {
	return Error{
		err: err,
		typ: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(err string) Error {
	return Error{
		err: err,
		typ: ErrorTypeIncorrectInput,
	}
}

func NewInternalError(err string) Error {
	return Error{
		err: err,
		typ: ErrorTypeInternal,
	}
}

func NewNotFoundError(err string) Error {
	return Error{
		err: err,
		typ: ErrorTypeNotFound,
	}
}
