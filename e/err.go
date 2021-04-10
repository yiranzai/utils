package e

import (
	"fmt"
)

// Error封装
type MyError struct {
	code    int    // 状态码
	message string // 状态内容
}

var codes = map[int]string{}

// 定制新的错误
func NewErr(code int, msg string) *MyError {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("code: %d has been exist", code))
	}
	codes[code] = msg
	return &MyError{code: code, message: msg}
}

// Error 输出错误
func (e *MyError) Error() string {
	return fmt.Sprintf("Error - code: %d, msg: %s", e.code, e.message)
}

// Code 输出错误code
func (e *MyError) Code() int {
	return e.code
}

// Message 输出错误信息
func (e *MyError) Message() string {
	return e.message
}

// 解析错误
func DecodeErr(err error) *MyError {
	if err == nil {
		return Success
	}
	switch errType := err.(type) {
	case *MyError:
		return &MyError{
			code:    errType.Code(),
			message: errType.Message(),
		}
	default:
		return ServerError
	}
}
