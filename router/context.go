package router

/*Context is Behavior of Route Context In Application*/
type Context interface {
	Next()
	JSON(int, interface{})
	Set(string, interface{})
	Get(string) (interface{}, bool)
}
