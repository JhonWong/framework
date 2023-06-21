package framework

// IResponse代表返回方法
type IResponse interface {
	// Json输出
	Json(obj interface{}) IResponse

	// Jsonp输出
	Jsonp(obj interface{}) IResponse

	//xml输出
	Xml(obj interface{}) IResponse

	// html输出
	Html(file string, obj interface{}) IResponse

	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}
