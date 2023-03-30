// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/handler/comment"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/comment",
				Handler: comment.GetCommentListInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/comment",
					Handler: comment.CreateCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)
}