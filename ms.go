package msgo

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)

// 定义routerGroup结构体，该结构体是为了存储一个一级路由所对应的所有二级路由
// name 一级路由名字，不添加/
// handlerMap 用于存储每个二级路由->二级路由的处理方法
type routerGroup struct {
	name       string
	handlerMap map[string]HandleFunc
}

// 路由的入口，可以将多个routerGroup组合在一起
// groups 是所有一级路由的合集
type router struct {
	groups []*routerGroup
}

// 服务器的启用引擎
type Engine struct {
	*router
}

// 功能：新加一个一级路由
func (r *router) Group(name string) *routerGroup {
	group := &routerGroup{
		name:       name,
		handlerMap: make(map[string]HandleFunc),
	}
	r.groups = append(r.groups, group)
	return group
}

func (rg *routerGroup) Add(name string, handler HandleFunc) {
	rg.handlerMap[name] = handler
}

func New() *Engine {
	return &Engine{
		router: &router{},
	}
}

// 定义Engine结构体的一个方法Run，通过调用该方法将产生一个Http的套接字
func (engine *Engine) Run() {
	groups := engine.router.groups
	for _, group := range groups {
		for name, handler := range group.handlerMap {
			http.HandleFunc("/"+group.name+name, handler)
		}
	}
	//
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
