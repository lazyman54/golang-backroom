// @Author @bytedance
// @Time 2024/8/12
package ierror

import "fmt"

type IErrorNew interface {
	error
	GetCode() int64
	GetMsg() string
}
type IError struct {
}

func (i *IError) GetCode() int64 {
	return 0
}
func (i *IError) GetMsg() string {
	return ""
}

func (i *IError) Error() string {
	return ""
}

type MyError struct {
	Code int64
	Msg  string
}

func (e *MyError) GetCode() int64 {
	return e.Code
}
func (e *MyError) GetMsg() string {
	return e.Msg
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

type MyError2 struct {
	Code    int32
	Message string
}

func (e *MyError2) GetCode() int64 {
	return int64(e.Code)
}
func (e *MyError2) GetMsg() string {
	return e.Message
}

func (e *MyError2) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Message)
}

func Solution() {
	var e IErrorNew = &MyError{Code: 1, Msg: "hello"}

	var e2 IErrorNew = &MyError2{Code: 2, Message: "world"}
	commonErrHandler(e)
	commonErrHandler(e2)
}

func commonErrHandler(err error) {
	if err.(IErrorNew) != nil {
		println(err.(IErrorNew).GetCode())
		println(err.(IErrorNew).GetMsg())
	} else {
		fmt.Println(err)
	}
}
