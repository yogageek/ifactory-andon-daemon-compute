package errors

import (
	"fmt"
)

// swagger:response ErrorResponse
type EnsaasError struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (e *EnsaasError) Error() string {
	return fmt.Sprintf("%v: %v %v", e.Code, e.Title, e.Detail)
}

func (e *EnsaasError) ErrorCode() int {
	return e.Code
}

func (e *EnsaasError) HttpStatusCode() int {
	return GetHttpCode(e.Code)
}

func NewEnsaasError(errorCode int, msgs ...string) error {
	e := new(EnsaasError)
	e.Code = errorCode
	e.Title = Title[errorCode]
	e.Detail = Msg[errorCode]

	if len(msgs) > 0 {
		if msgs[0] != "" {
			e.Title = msgs[0]
		}
	}

	if len(msgs) > 1 {
		e.Detail = msgs[1]
	}

	return e
}
