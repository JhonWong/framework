package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func NewContext(request *http.Request, responseWriter http.ResponseWriter) *Context {
	return &Context{
		request:        request,
		responseWriter: responseWriter,
		ctx:            request.Context(),
		handlers:       []ControllerHandler{},
		index:          -1,
		writerMux:      &sync.Mutex{},
	}
}

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler
	index          int

	writerMux  *sync.Mutex
	hasTimeout bool

	params map[string]string
}

func (c *Context) Next() error {
	c.index++
	if c.index < len(c.handlers) {
		if err := c.handlers[c.index](c); err != nil {
			return err
		}
	}
	return nil
}

// #region base function

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) setParams(params map[string]string) {
	ctx.params = params
}

//  #endregion

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// #region implement context.Context
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// #endregion

// #region form post
func (ctx *Context) FormAll() map[string][]string {

	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// #endregion
