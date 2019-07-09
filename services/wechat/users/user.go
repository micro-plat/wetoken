
package users

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type UserHandler struct {
	container component.IContainer
}

func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{container: container}
}


func (u *UserHandler) GetHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *UserHandler) PutHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *UserHandler) DeleteHandle(ctx *context.Context) (r interface{}) {
	return "success"
}





