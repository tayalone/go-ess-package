package router

/*Context is Behavior of Route Context In Application*/

/*Bahavior is Method Which Router must have*/
type Bahavior interface {
	GET(path string, handlers ...func(Context))
	POST(path string, handlers ...func(Context))
	PATCH(path string, handlers ...func(Context))
	PUT(path string, handlers ...func(Context))
	DELETE(path string, handlers ...func(Context))
	Use(middleware func(Context))
}

/*RoterGrouper is Method Which RoterGrouper must have*/
type RoterGrouper interface {
	Group(path string, handlers ...func(Context)) RoterGrouper
	Bahavior
}

/*Route is Behavior of Route Method In Application*/
type Route interface {
	Start()
	Testing(method string, path string, body map[string]interface{}) (int, string)
	Group(path string, handlers ...func(Context)) RoterGrouper
	Bahavior
}
