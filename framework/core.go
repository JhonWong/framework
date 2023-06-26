package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("Add ruote err:", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("Add ruote err:", err)
	}
}
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("Add ruote err:", err)
	}
}
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("Add ruote err:", err)
	}
}

func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if tree, ok := c.router[upperMethod]; ok {
		return tree.FindHandler(uri)
	}
	return nil
}

func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if tree, ok := c.router[upperMethod]; ok {
		return tree.root.matchNode(uri)
	}
	return nil
}

// 框架核心结构实现handle接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//创建context
	ctx := NewContext(request, response)

	node := c.FindRouteNodeByRequest(request)
	if node == nil {
		ctx.SetStatus(404).Json("not found")
		return
	}

	//set route param
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.setParams(params)

	ctx.SetHandlers(node.handlers)

	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
