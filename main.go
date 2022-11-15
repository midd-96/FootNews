package main

import (
	"NewsAppApi/config"
	h "NewsAppApi/handler"
	m "NewsAppApi/middleware"
	"NewsAppApi/repo"
	"NewsAppApi/routes"
	"NewsAppApi/service"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/subosito/gotenv"
)

//to call functions before main functions
func init() {
	gotenv.Load()
}

func main() {

	//Loading value from env file
	port := os.Getenv("PORT")

	//For making log file
	file, err := os.OpenFile("Logging Details", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Logging in File not done")
	}
	log.SetOutput(file)

	// creating an instance of chi r
	router := chi.NewRouter()

	// using logger to display each request
	router.Use(middleware.Logger)

	config.Init()

	var (
		db         *sql.DB              = config.ConnectDB()
		mailConfig config.MailConfig    = config.NewMailConfig()
		adminRepo  repo.AdminRepository = repo.NewAdminRepo(db)
		userRepo   repo.UserRepository  = repo.NewUserRepo(db)

		jwtAdminService service.JWTService   = service.NewJWTAdminService()
		jwtUserService  service.JWTService   = service.NewJWTUserService()
		authService     service.AuthService  = service.NewAuthService(adminRepo, userRepo)
		adminService    service.AdminService = service.NewAdminService(adminRepo, userRepo)
		userService     service.UserService  = service.NewUserService(userRepo, adminRepo, mailConfig)

		authHandler h.AuthHandler = h.NewAuthHandler(jwtAdminService,
			jwtUserService, authService,
			adminService,
			userService)
		//validate)
		adminMiddleware m.Middleware   = m.NewMiddlewareAdmin(jwtAdminService)
		userMiddleware  m.Middleware   = m.NewMiddlewareUser(jwtUserService)
		adminHandler    h.AdminHandler = h.NewAdminHandler(adminService, userService)
		userHandler     h.UserHandler  = h.NewUserHandler(userService)

		adminRoute routes.AdminRoute = routes.NewAdminRoute()
		userRoute  routes.UserRoute  = routes.NewUserRoute()
	)

	//routing
	adminRoute.AdminRouter(router,
		authHandler,
		adminHandler,
		adminMiddleware)

	userRoute.UserRouter(router,
		authHandler,
		userMiddleware,
		userHandler)

	log.Println("Api is listening on port:", port)
	// Starting server
	http.ListenAndServe(":"+port, router)

}
