package handler

import (
	"NewsAppApi/common/response"
	"NewsAppApi/model"
	"NewsAppApi/service"
	"NewsAppApi/utils"
	"log"
	"net/http"
	"strconv"
)

type AdminHandler interface {
	ApprovePost() http.HandlerFunc
	ListAllPosts() http.HandlerFunc
}

type adminHandler struct {
	adminService service.AdminService
	userService  service.UserService
}

func NewAdminHandler(
	adminService service.AdminService,
	userService service.UserService,
) AdminHandler {
	return &adminHandler{
		adminService: adminService,
		userService:  userService,
	}
}

func (c *adminHandler) ListAllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))

		sortby := r.URL.Query().Get("sortby")

		log.Println(page, "   ", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		posts, metadata, err := c.adminService.ListAllPosts(pagenation, sortby)

		result := struct {
			Posts *[]model.PostResponse
			Meta  *utils.Metadata
		}{
			Posts: posts,
			Meta:  metadata,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All posts", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}

func (c *adminHandler) ApprovePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post_id, _ := strconv.Atoi(r.URL.Query().Get("post_id"))
		//post_id, _ := strconv.Atoi(chi.URLParam(r, "Post_id"))
		log.Println("post id :", post_id)
		err := c.adminService.ApprovePost(post_id)

		if err != nil {
			response := response.ErrorResponse("Approvel failed", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Post approved successfully", post_id)
		utils.ResponseJSON(w, response)
	}
}
