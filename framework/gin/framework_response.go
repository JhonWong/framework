package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// IResponse代表返回方法
type IResponse interface {
	// Json输出
	IJson(obj interface{}) IResponse

	// Jsonp输出
	IJsonp(obj interface{}) IResponse

	//xml输出
	IXml(obj interface{}) IResponse

	// html输出
	IHtml(file string, obj interface{}) IResponse

	// string
	IText(format string, values ...interface{}) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置200状态
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

// Jsonp输出
func (ctx *Context) IJsonp(obj interface{}) IResponse {
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	//进行字符串过滤，防止xss攻击
	callback := template.JSEscapeString(callbackFunc)

	//输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	//输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	//数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	//输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

// xml输出
func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

// html输出
func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	//读取模板文件，创建template实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	//执行execute方法将obj和模板结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

// string
func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

// 重定向
func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

// header
func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

// Cookie
func (ctx *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// 设置状态码
func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) ISetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}
