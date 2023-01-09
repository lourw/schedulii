package data_model

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
}
