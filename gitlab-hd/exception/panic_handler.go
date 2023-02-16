package exception

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"gitlab-hd/config/token"
	"gitlab-hd/dto"
	"net/http"
	"runtime/debug"
)

// ErrHandle 统一异常处理
func ErrHandle(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			apiErr, isApiErr := r.(*ApiError)
			tokenErr, isTokenErr := r.(*token.TokenError)
			if isApiErr {
				// 打印错误堆栈信息
				logger.WithField("ErrMsg", apiErr.Error()).Error("PanicHandler handled apiError: ")
				// 封装通用json返回
				c.JSON(http.StatusInternalServerError, apiErr)
			} else if isTokenErr {
				// 打印错误堆栈信息
				logger.WithField("ErrMsg", tokenErr.Error()).Error("PanicHandler handled tokenError: ")
				// 封装通用json返回
				c.JSON(http.StatusUnauthorized, tokenErr)
			} else {
				// 打印错误堆栈信息
				err := r.(error)
				logger.WithField("ErrMsg", err.Error()).Error("PanicHandler handled ordinaryError: ")
				debug.PrintStack()
				// 封装通用json返回
				c.JSON(http.StatusInternalServerError, NewApiError(dto.InternalServerError, dto.GetResultMsg(dto.InternalServerError)))
			}
			c.Abort()
		}
	}()
	c.Next()
}
