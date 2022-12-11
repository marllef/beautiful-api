package server


type Router interface {
	GetRoutes() Routes
	AddRoute(key string, route Route)
	RemoveRoute(key string)
}

type router struct {
	routes Routes
	controlers []interface{}
	server Server
}

func NewRouter(httpServer Server ,controlers ...interface{}) *router {
	return &router{
		routes: make(Routes),
		controlers: controlers,
	}
}

func (r *router) GetRoutes() Routes {
	return r.routes
}

func (r *router) AddRoute(key string, route Route) {
	r.routes[key] = route
}

func (r *router) RemoveRoute(key string) {
	delete(r.routes, key)
}
