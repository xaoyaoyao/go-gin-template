/**
 * Package ret
 * @file      : code.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:24
 **/

package ret

import "fmt"

var (
	ErrUnknownConfigType = fmt.Errorf("unknown config Type")
)

var DefaultMinCode = 20000

var (
	Success             = StatusCode{"200", 200}
	BadRequest          = StatusCode{"400", 400}
	Unauthorized        = StatusCode{"401", 401}
	Forbidden           = StatusCode{"403", 403}
	NotFound            = StatusCode{"404", 404}
	ServerInternalError = StatusCode{"500", 500}

	availableRetCodes = []StatusCode{
		Success, BadRequest, Unauthorized, Forbidden, NotFound, ServerInternalError,
	}
)

type StatusCode struct {
	text string
	code int
}

func (t StatusCode) String() string {
	return t.text
}

func (t StatusCode) Code() int {
	return t.code
}

func ToStatusCode(value int) (StatusCode, error) {
	for _, v := range availableRetCodes {
		if v.Code() == value {
			return v, nil
		}
	}
	return StatusCode{}, fmt.Errorf("%w: %s", ErrUnknownConfigType, value)
}
