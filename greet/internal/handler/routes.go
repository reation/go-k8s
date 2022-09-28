// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-k8s/greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GreetHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/from/user/get/:id",
				Handler: UserFunctionHandler(serverCtx, "get"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/from/user/list/:id",
				Handler: UserFunctionHandler(serverCtx, "list"),
			},
		},
	)
}