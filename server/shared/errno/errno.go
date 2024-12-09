package errno

import (
	"fmt"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/errno"
)

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_code:%d,err_msg:%d", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) Errno {
	return Errno{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e Errno) WithMessage(msg string) Errno {
	e.ErrMsg = msg
	return e
}

var (
	Success        = NewErrNo(int64(errno.Err_Success), "success")
	NoRoute        = NewErrNo(int64(errno.Err_NoRoute), "noRoute")
	NoMethod       = NewErrNo(int64(errno.Err_NoMethod), "NoMethod")
	BadRequest     = NewErrNo(int64(errno.Err_BadRequest), "BadRequest")
	ParamsErr      = NewErrNo(int64(errno.Err_ParamsErr), "paramsErr")
	AuthorizeFail  = NewErrNo(int64(errno.Err_AuthorizeFail), "AuthorizeFail")
	TooManyRequest = NewErrNo(int64(errno.Err_TooManyRequest), "Too many requests")
)
