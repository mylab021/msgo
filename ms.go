package msgo

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	handleFuncMap map[string]HandleFunc
}

type Engine struct {
	router
}

func (r *router) Add(path string, handler HandleFunc) {
	r.handleFuncMap[path] = handler
}

func (r *router) Get(path string, f HandleFunc) {

}

// 通过调用New方法将返回一个Engine的实例对象
func New() *Engine {
	return &Engine{
		router: router{
			make(map[string]HandleFunc),
		},
	}
}

// 定义Engine结构体的一个方法Run，通过调用该方法将产生一个Http的套接字
func (engine *Engine) Run() {
	// 从engine中取出router，然后遍历router，从中取出path->handleFunc，
	// NOTE: 实例化后的结构体，可以直接调用该结构体成员（该成员也是一个结构体）下的属性。
	for path, handle := range engine.handleFuncMap {
		// 从自定义router中取出的path->handleFunc，添加到Go Http 原生框架中。
		http.HandleFunc(path, handle)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
