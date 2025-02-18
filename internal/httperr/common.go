/**
 * Package httperr
 * @file      : common.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 15:23
 **/

package httperr

import (
	"database/sql"
	"github.com/coverai/api/internal/logs"
	"github.com/coverai/api/internal/xerrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Respond(ctx *gin.Context, err error) {
	if err == sql.ErrNoRows {
		err = xerrors.NewNotFoundError("not found")
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		validationError(ctx, err)
		return
	}
	xerror, ok := err.(xerrors.Error)
	if !ok {
		httpRespondWithError(
			ctx,
			"INTERNAL_ERROR",
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	errorType := xerror.ErrorType()
	httpRespondWithError(
		ctx,
		errorType.String(),
		err.Error(),
		errorType.Status(),
	)
}

func validationError(ctx *gin.Context, err error) {
	httpRespondWithError(
		ctx,
		"INCORRECT_INPUT",
		err.Error(),
		http.StatusBadRequest,
	)
}

func httpRespondWithError(ctx *gin.Context, errType, message string, status int) {
	logs.FromContext(ctx).WithFields(logrus.Fields{
		"type": errType,
		"push": message,
		"code": status,
	}).Warn("http error")
	ctx.JSON(status, ErrorResponse{status, errType, message})
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	ErrorType string `json:"type"`
	Message   string `json:"push"`
}
