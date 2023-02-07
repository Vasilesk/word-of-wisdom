package request

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request interface {
	ReadJSON(p interface{}) error
	GetHeader(key string) string
	FormValue(key string) string
	Ctx() context.Context
	Raw() *http.Request
}

type request struct {
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) Request {
	return &request{
		w: w,
		r: r,
	}
}

func (r *request) ReadJSON(p interface{}) error {
	defer r.r.Body.Close()

	body, err := io.ReadAll(r.r.Body)
	if err != nil {
		return fmt.Errorf("reading all body: %w", err)
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return fmt.Errorf("unmarshalling body: %w", err)
	}

	return nil
}

func (r *request) GetHeader(key string) string {
	return r.r.Header.Get(key)
}

func (r *request) FormValue(key string) string {
	return r.r.FormValue(key)
}

func (r *request) Ctx() context.Context {
	return r.r.Context()
}

func (r *request) Raw() *http.Request {
	return r.r
}
