/**
 * Package rsp
 * @file      : response.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:22
 **/

package rsp

import (
	"net/http"
)

type ResponseEntity struct {
	Code    int32       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type RetResult struct {
	Status int
	Err    error
	Data   interface{}
}

func MakeResponseEntity(code int32, message string) *ResponseEntity {
	return &ResponseEntity{
		Code:    code,
		Message: message,
	}
}

func MakeResponseData(msg string, data interface{}) *ResponseEntity {
	if data == nil {
		return MakeResponseEntityOK(msg)
	}
	return &ResponseEntity{
		Code:    http.StatusOK,
		Message: msg,
		Data:    &data,
	}
}

func MakeResponseEntityOK(msg string) *ResponseEntity {
	return &ResponseEntity{
		Code:    http.StatusOK,
		Message: msg,
	}
}

func MakeResponseEntityNotFound(e error) *ResponseEntity {
	return MakeResponseEntityError(http.StatusNotFound, e)
}

func MakeResponseEntityError(code int32, e error) *ResponseEntity {
	return &ResponseEntity{
		Code:    code,
		Message: e.Error(),
	}
}

func NewRetResult(status int, err error, data interface{}) RetResult {
	if data == nil {
		return RetResult{
			Status: status,
			Err:    err,
		}
	}
	return RetResult{
		Status: status,
		Err:    err,
		Data:   &data,
	}
}

func IsOK(entity RetResult) bool {
	return entity.Status == http.StatusOK
}

func NewInternalServer(err error) RetResult {
	return NewRetResult(http.StatusInternalServerError, err, nil)
}

func NewBadRequest(err error) RetResult {
	return NewRetResult(http.StatusBadRequest, err, nil)
}
