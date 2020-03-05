package controller

import (
	"github.com/chenyingdi/web-toolkits/data_gf/v1"
	"github.com/chenyingdi/web-toolkits/req_gf"
	"github.com/chenyingdi/web-toolkits/resp_gf"
	"github.com/gogf/gf/net/ghttp"
)

type REST struct {
	Model          data.Model
	AllowedMethods []string
	RequestFilter  []string
}

func NewRest(model data.Model, methods []string) *REST {
	return &REST{
		Model:          model,
		AllowedMethods: methods,
	}
}

func (rest *REST) Get(r *ghttp.Request) {
	res := resp.Response{}
	dao := data.DAO{}

	args, err := req.GetArgs(r)
	if err != nil {
		res.ServerError(r, err.Error())
		return
	}

	if !rest.checkMethod("GET") {
		res.UnAuthorized(r)
		return
	}

	result, err := dao.Get(args.(data.Args), rest.Model)
	if err != nil {
		res.ServerError(r, err.Error())
		return
	}

	res.Success(r, result)
}

func (rest *REST) Post(r *ghttp.Request) {
	res := resp.Response{}
	dao := data.DAO{}
	args := r.GetMap()

	if !rest.checkMethod("POST") {
		res.UnAuthorized(r)
		return
	}

	_, err := dao.Create(args, rest.Model)
	if err != nil {
		res.ServerError(r, err.Error())
		return
	}

	res.Success(r, nil)
}

func (rest *REST) Put(r *ghttp.Request) {
	res := resp.Response{}
	dao := data.DAO{}
	args := r.GetMap()

	if !rest.checkMethod("PUT") {
		res.UnAuthorized(r)
		return
	}

	err := dao.Update(args, rest.Model)
	if err != nil {
		res.ServerError(r, err.Error())
		return
	}

	res.Success(r, nil)
}

func (rest *REST) Delete(r *ghttp.Request) {
	res := resp.Response{}
	dao := data.DAO{}
	id := r.GetInt("id")

	if !rest.checkMethod("DELETE") {
		res.UnAuthorized(r)
		return
	}

	err := dao.DeleteByID(id, rest.Model)
	if err != nil {
		res.ServerError(r, err.Error())
		return
	}

	res.Success(r, nil)
}

func (rest *REST) checkMethod(method string) bool {
	var switcher = false

	for i := range rest.AllowedMethods {
		if rest.AllowedMethods[i] == method {
			switcher = true
		}
	}

	return switcher
}
