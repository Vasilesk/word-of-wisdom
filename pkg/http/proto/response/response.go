package response

import (
	"fmt"
	"net/http"
)

type Response interface {
	Body() interface{}
	HTTPCode() int
}

func Ok(data interface{}) Response {
	return &resp{body: data, code: http.StatusOK}
}

func OkStatus() Response {
	return Ok(okResp{Status: "ok"})
}

func ErrResponse(httpCode int, msg string) Response {
	return &resp{
		body: errResp{Error: errRespBody{
			Msg: msg,
		}},
		code: httpCode,
	}
}

func ErrServer() Response {
	const serverErrorMsg = "server error occurred"

	return ErrResponse(http.StatusInternalServerError, serverErrorMsg)
}

func ErrClient(format string, a ...interface{}) Response {
	return ErrResponse(http.StatusBadRequest, fmt.Sprintf(format, a...))
}

type resp struct {
	body interface{}
	code int
}

func (r *resp) Body() interface{} {
	return r.body
}

func (r *resp) HTTPCode() int {
	return r.code
}

type okResp struct {
	Status string `json:"status"`
}

type errResp struct {
	Error errRespBody `json:"error"`
}

type errRespBody struct {
	Msg string `json:"msg"`
}
