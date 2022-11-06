package mock

import "github.com/tayalone/go-ess-package/router"

/*UseGlobal set value to ctx key global*/
func UseGlobal(ctx router.Context) {
	ctx.Set("global", 1)
	ctx.Next()
}

/*UseGroupV1 set value to ctx key global*/
func UseGroupV1(ctx router.Context) {
	ctx.Set("group-v1", "v1")
	ctx.Next()
}

/*UseSubGroupV1 set value to ctx key global*/
func UseSubGroupV1(ctx router.Context) {
	ctx.Set("sub-group-v1", "sub-v1")
	ctx.Next()
}
