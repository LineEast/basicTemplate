package server

import (
	"kmfRedirect/internal/models"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func (s *Server) CreateUser() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		request := models.User{}

		err := json.Unmarshal(ctx.Request.Body(), &request)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = s.database.Create(ctx, &request)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}

func (s *Server) GetAllUserList() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		users, err := s.database.ReadAllUserList(ctx)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(users)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetBody(body)
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}

func (s *Server) GetUser() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		user := models.User{}

		err := json.Unmarshal(ctx.Request.Body(), &user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = s.database.ReadUser(ctx, &user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetBody(body)

		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}
