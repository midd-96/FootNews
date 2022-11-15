package handler

import (
	"NewsAppApi/common/response"
	"NewsAppApi/model"
	"NewsAppApi/service"
	"NewsAppApi/utils"
	"encoding/json"
	"net/http"
	"strings"
)

type AuthHandler interface {
	AdminSignup() http.HandlerFunc
	AdminLogin() http.HandlerFunc
	UserLogin() http.HandlerFunc
	UserSignup() http.HandlerFunc
	AdminRefreshToken() http.HandlerFunc
	UserRefreshToken() http.HandlerFunc
}

type authHandler struct {
	jwtAdminService service.JWTService
	jwtUserService  service.JWTService
	authService     service.AuthService
	adminService    service.AdminService
	userService     service.UserService
}

func NewAuthHandler(
	jwtAdminService service.JWTService,
	jwtUserService service.JWTService,
	authService service.AuthService,
	adminService service.AdminService,
	userService service.UserService,

) AuthHandler {
	return &authHandler{
		jwtAdminService: jwtAdminService,
		jwtUserService:  jwtUserService,
		authService:     authService,
		adminService:    adminService,
		userService:     userService,
	}
}

func (c *authHandler) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var adminLogin model.Admin

		json.NewDecoder(r.Body).Decode(&adminLogin)

		//verifying  admin credentials
		err := c.authService.VerifyAdmin(adminLogin.Email, adminLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, response)
			return
		}

		//getting admin values
		admin, _ := c.adminService.FindAdmin(adminLogin.Email)
		token := c.jwtAdminService.GenerateToken(admin.ID, admin.Email, "admin")
		admin.Password = ""
		admin.Token = token
		response := response.SuccessResponse(true, "SUCCESS", admin.Token)
		utils.ResponseJSON(w, response)
	}

}

func (c *authHandler) UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userLogin model.User

		json.NewDecoder(r.Body).Decode(&userLogin)

		//verify User details
		err := c.authService.VerifyUser(userLogin.Email, userLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, response)
			return
		}

		//fetching user details
		user, _ := c.userService.FindUser(userLogin.Email)
		token := c.jwtUserService.GenerateToken(user.ID, user.Email, "user")
		user.Password = ""
		user.Token = token
		response := response.SuccessResponse(true, "SUCCESS", user.Token)
		utils.ResponseJSON(w, response)
	}
}

func (c *authHandler) UserSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newUser model.User

		//fetching data
		json.NewDecoder(r.Body).Decode(&newUser)

		if newUser.Password == newUser.ConfirmPassword {

			err := c.userService.CreateUser(newUser)

			if err != nil {
				response := response.ErrorResponse("Failed to create user", err.Error(), nil)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				utils.ResponseJSON(w, response)
				return
			}

			user, _ := c.userService.FindUser(newUser.Email)
			user.Password = ""
			response := response.SuccessResponse(true, "SUCCESS", user)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			utils.ResponseJSON(w, response)
		} else {
			response := response.ErrorResponse("Password Missmatch", "Enter same password two times", nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
	}
}

func (c *authHandler) AdminSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newAdmin model.Admin

		//fetching data
		json.NewDecoder(r.Body).Decode(&newAdmin)

		if newAdmin.Password == newAdmin.ConfirmPassword {

			err := c.adminService.CreateAdmin(newAdmin)

			if err != nil {
				response := response.ErrorResponse("Failed to signup", err.Error(), nil)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				utils.ResponseJSON(w, response)
				return
			}

			admin, _ := c.adminService.FindAdmin(newAdmin.Email)
			admin.Password = ""
			response := response.SuccessResponse(true, "SUCCESS", admin)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			utils.ResponseJSON(w, response)
		} else {
			response := response.ErrorResponse("Password Missmatch", "Enter same password two times", nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
	}
}
func (c *authHandler) AdminRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		autheader := r.Header.Get("Authorization")
		bearerToken := strings.Split(autheader, " ")
		token := bearerToken[1]

		refreshToken, err := c.jwtAdminService.GenerateRefreshToken(token)

		if err != nil {
			response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "SUCCESS", refreshToken)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}

func (c *authHandler) UserRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		autheader := r.Header.Get("Authorization")
		bearerToken := strings.Split(autheader, " ")
		token := bearerToken[1]

		refreshToken, err := c.jwtUserService.GenerateRefreshToken(token)

		if err != nil {
			response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "SUCCESS", refreshToken)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}
