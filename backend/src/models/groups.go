package models

type Groups struct {
	GroupID int `json:"groupID" form:"groupID" binding:"required"`
    GroupName string `json:"groupName" form:"groupName" binding:"required"`
    GroupURL string `json:"groupURL" form:"groupURL" binding:"required"`
    AvailableStartHour int `json:"availableStartHour" form:"availableStartHour" binding:"required"`
    AvailableEndHour int `json:"avaiilableEndHour" form:"avaiilableEndHour" binding:"required"`
}
