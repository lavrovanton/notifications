package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error"`
}

func errorResponse(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, response{err.Error()})
}

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("resource not found")
	ErrBadParamInput       = errors.New("param is not valid")
)
