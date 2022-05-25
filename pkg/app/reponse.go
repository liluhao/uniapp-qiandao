package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//使之实现error接口，有一个方法
type Errno struct {
	Code    int
	Message string
}

//Err表示一个错误，使之实现error接口，有一个方法
type Err struct {
	Code    int
	Message string
	Err     error
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (err Errno) Error() string {
	return err.Message
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// DecodeErr判断错误的类型，如果为空，表示正常返回；如果不为空就判断是哪种类型的错误返回，如果都不是就统一返回服务器异常
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

// SendResponse 同意封装返回结果集
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
