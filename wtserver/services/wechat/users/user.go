package users

import (
	"github.com/micro-plat/hydra"
)

type UserHandler struct {
}

func NewUserHandler() (u *UserHandler) {
	return &UserHandler{}
}

func (u *UserHandler) GetHandle(ctx hydra.IContext) (r interface{}) {
	return "success"
}
func (u *UserHandler) PostHandle(ctx hydra.IContext) (r interface{}) {
	return "success"
}
func (u *UserHandler) PutHandle(ctx hydra.IContext) (r interface{}) {
	return "success"
}
func (u *UserHandler) DeleteHandle(ctx hydra.IContext) (r interface{}) {
	return "success"
}
