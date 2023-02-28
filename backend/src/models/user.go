package models

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
}
