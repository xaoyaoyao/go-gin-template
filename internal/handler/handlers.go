/**
 * Package handler
 * @file      : handlers.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:30
 **/

package handler

import "github.com/coverai/api/internal/domain/user"

type HandlerImpl struct {
	credential *user.Credential
}

func NewHandlerImpl(
	credential *user.Credential,
) HandlerImpl {
	return HandlerImpl{
		credential: credential,
	}
}
