package perror

import "fmt"

var codeErrMap = map[int]string{
	0:     "SUCCESS",
	80403: "AUTH",
	80501: "ARGS",
	80502: "REMOTE",
	80999: "UNKNOWN",
}

type Error struct {
	/*
		-1000 ~ -100 为 http code 的保留code，其值为 -httpCode
		-1 位保留code，用于默认的错误code
		0  为请求成功
		>0 为业务保留错误
	*/
	Code int
	Msg  string
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func CreateErr(code int, msg string) error {
	if code == 0 {
		return nil
	}

	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func GenCodeErr(code int) error {
	if code == 0 {
		return nil
	}

	msg, ok := codeErrMap[code]
	if ok {
		return CreateErr(code, msg)
	}

	return CreateErr(code, "UNKNOWN")
}

func GetCode(err error) int {
	if err == nil {
		return 0
	}

	if x, ok := err.(*Error); ok {
		return x.Code
	}
	return -1
}
