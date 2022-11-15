package handler

import (
	"NewsAppApi/common/response"
	"NewsAppApi/model"
	"NewsAppApi/service"
	"NewsAppApi/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserHandler interface {
	AllPosts() http.HandlerFunc
	AddPost() http.HandlerFunc
	SendVerificationMail() http.HandlerFunc
	VerifyAccount() http.HandlerFunc
	AddComment() http.HandlerFunc
	MarkVote() http.HandlerFunc
	ReportPost() http.HandlerFunc
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (c *userHandler) ReportPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var report model.Report
		json.NewDecoder(r.Body).Decode(&report)
		report.UserID, _ = strconv.Atoi(r.Header.Get("user_id"))
		err := c.userService.ReportPost(report)
		if err != nil {
			response := response.ErrorResponse("Failed to mark vote", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", report)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

func (c *userHandler) MarkVote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newVote model.Vote
		json.NewDecoder(r.Body).Decode(&newVote)
		newVote.UserID, _ = strconv.Atoi(r.Header.Get("user_id"))
		err, msg := c.userService.MarkVote(newVote)
		if err != nil {
			response := response.ErrorResponse("Failed to mark vote", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, msg, newVote)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}

}

func (c *userHandler) AddComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newComment model.Comment
		json.NewDecoder(r.Body).Decode(&newComment)
		newComment.UserID, _ = strconv.Atoi(r.Header.Get("user_id"))
		err := c.userService.AddComment(newComment)
		if err != nil {
			response := response.ErrorResponse("Failed to add new comment", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", newComment)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

func (c *userHandler) AddPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newPost model.Post
		json.NewDecoder(r.Body).Decode(&newPost)
		newPost.UserID, _ = strconv.Atoi(r.Header.Get("user_id"))
		_, err := c.userService.AddPost(newPost)
		if err != nil {
			response := response.ErrorResponse("Failed to add new post", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", newPost)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

func (c *userHandler) AllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))

		sortby := r.URL.Query().Get("sortby")

		log.Println(page, "   ", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		//log.Println("pagenation from handler", pagenation)

		posts, metadata, err := c.userService.AllPosts(pagenation, sortby)

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

func (c *userHandler) VerifyAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")
		code, _ := strconv.Atoi(r.URL.Query().Get("Code"))

		err := c.userService.VerifyAccount(email, code)

		if err != nil {
			response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Account verified successfully", email)
		utils.ResponseJSON(w, response)
	}
}

func (c *userHandler) SendVerificationMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")

		_, err := c.userService.FindUser(email)

		if err == nil {
			err = c.userService.SendVerificationEmail(email)
		}

		if err != nil {
			response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Verification mail sent successfully", email)
		utils.ResponseJSON(w, response)
	}
}
