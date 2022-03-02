package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	k8sErr "k8s.io/apimachinery/pkg/api/errors"
)

// ApplicationError struct
type ApplicationError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	_error  error
}

// NewHttpError function description
func NewHttpError(code int, err error) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Type:    http.StatusText(code),
		Message: err.Error(),
		_error:  err,
	}
}

func InternalServerError(err error) *ApplicationError {
	return NewHttpError(http.StatusInternalServerError, err)
}

func NotFoundError(err error) *ApplicationError {
	return NewHttpError(http.StatusNotFound, err)
}

func BadRequestError(err error) *ApplicationError {
	return NewHttpError(http.StatusBadRequest, err)
}

func KubeError(err error) *ApplicationError {
	if status := k8sErr.APIStatus(nil); errors.As(err, &status) {
		code := int(status.Status().Code)
		return NewHttpError(code, errors.Wrap(err, string(status.Status().Reason)))
	}
	return InternalServerError(errors.New("unknown error"))
}

func GinerateError(context *gin.Context, err *ApplicationError) *gin.Context {
	context.JSON(err.Code, err)
	context.AbortWithError(err.Code, err._error)
	return context
}
