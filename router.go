package easy_http

type Router struct {
	Handler string
	Method  string
	Path    string
}

type Handlers map[string]HandlerFun

type HandlerFun func()
type RouterArray []Router