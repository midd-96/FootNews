package model

import "time"

type PostResponse struct {
	ID           int       `json:"id,omitempty"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       int       `json:"user_id"`
	Votes        int       `json:"votes,omitempty"`
	UserName     string    `json:"user_name,omitempty"`
	CommentCount int       `json:"comment_count,omitempty"`
	Approved     bool      `json:"approved"`
	TotalRecords int       `json:"total_records,omitempty"`
}

type AdminResponse struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Created_At time.Time `json:"created_at"`
	Email      string    `json:"email"`
	//Name       string    `json:"name"`
	Password string `json:"hash_password"`
	Token    string `json:"token,omitempty"`
}

type UserResponse struct {
	ID         int       `json:"id"`
	Created_At time.Time `json:"created_at"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
}

type DoctorResponse struct {
	ID         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Approvel   bool   `json:"approvel"`
	Token      string `json:"token,omitempty"`
}

type Appointments struct {
	Day_consult    string `json:"consulting_day"`
	Time_consult   string `json:"consulting_time"`
	Payment_mode   string `json:"payment_mode"`
	Payment_status bool   `json:"payment_status"`
	Email          string `json:"email"`
}

type AppointmentByDoctor struct {
	Time_consult   string `json:"consulting_time"`
	Payment_mode   string `json:"payment_mode"`
	Payment_status bool   `json:"payment_status"`
	Email          string `json:"email"`
}

type Filter struct {
	Day      []string `json:"day"`
	DoctorId []string `json:"doc_id"`
	Sort     []string `json:"sort"`
}

type Day struct {
	Day string `json:"day"`
}

type DoctorId struct {
	DoctorId string `json:"doc_id"`
}

type Sort struct {
	Sort string `json:"sort"`
}
