package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-k8s/greet/internal/logic"
	"go-k8s/greet/internal/svc"
	"go-k8s/greet/internal/types"
)

func funcHandler(userLogic *logic.UserLogic, req *types.UserRequest, funcName string) (resp *types.UserResponse, err error) {
	switch funcName {
	case "get":
		resp, err = userLogic.UserInfo(req)
	case "list":
		resp, err = userLogic.UserList()
	}

	return
}

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UserFunctionHandler(serverCtx *svc.ServiceContext, funcName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), serverCtx)
		resp, err := funcHandler(l, &req, funcName)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
