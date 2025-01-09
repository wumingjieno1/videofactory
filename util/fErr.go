package util

import (
	"fmt"
)

type WErr struct {
	method string
	err    error
	msg    []any
}

func (e WErr) ToError() error {
	if e.err == nil && len(e.msg) == 0 {
		return nil
	}
	if e.err == nil {
		return fmt.Errorf("[%s] %s", e.method, genMsg(e.msg))
	}
	return fmt.Errorf("[%s] %s: %w", e.method, genMsg(e.msg), e.err)
}

func (e *WErr) Err(err error) *WErr {
	e.err = err
	return e
}

func (e *WErr) Msg(msg ...any) *WErr {
	e.msg = msg
	return e
}

func (e *WErr) Method(method string) *WErr {
	e.method = method
	return e
}


func EmptyWErr(method string) *WErr {
	return &WErr{method: method}
}

func Err(method string, err error, msg ...any) error {
	if err == nil && len(msg) == 0 {
		return nil
	}
	if err == nil {
		return fmt.Errorf("[%s] %s", method, genMsg(msg))
	}
	return fmt.Errorf("[%s] %s: %w", method, genMsg(msg), err)
}

func genMsg(msg []any) string {
	if len(msg) == 0 {
		return ""
	}
	msg0, _ := msg[0].(string)
	if len(msg) > 1 && msg0 != "" {
		msg0 = fmt.Sprintf(msg0, msg[1:]...)
	}
	return msg0
}
