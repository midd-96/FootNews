package routes

import (
	h "NewsAppApi/handler"
	m "NewsAppApi/middleware"

	"github.com/go-chi/chi"
)

type UserRoute interface {
	UserRouter(router chi.Router,
		authHandler h.AuthHandler,
		middleware m.Middleware,
		userHandler h.UserHandler,
	)
}

type userRoute struct{}

func NewUserRoute() UserRoute {
	return &userRoute{}
}

func (r *userRoute) UserRouter(routes chi.Router,
	authHandler h.AuthHandler,
	middleware m.Middleware,
	userHandler h.UserHandler) {

	routes.Post("/user/signup", authHandler.UserSignup())
	routes.Post("/user/login", authHandler.UserLogin())
	routes.Post("/user/send/verification", userHandler.SendVerificationMail())
	routes.Patch("/user/verify/account", userHandler.VerifyAccount())
	routes.Group(func(r chi.Router) {
		r.Use(middleware.AuthorizeJwt)
		r.Post("/newpost", userHandler.AddPost())
		r.Get("/user/token/refresh", authHandler.UserRefreshToken())
		r.Get("/allpost", userHandler.AllPosts())
		r.Post("/user/addcomment", userHandler.AddComment())
		r.Patch("/user/vote", userHandler.MarkVote())
		r.Post("/user/report", userHandler.ReportPost())
	})

}
