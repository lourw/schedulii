package models

type Groups struct {
    GroupID int `json:"groupID" form:"groupID" binding:"required"`
    GroupName string `json:"groupName" form:"groupName"`
    GroupURL string `json:"groupURL" form:"groupURL"`
    AvailableStartHour int `json:"availableStartHour" form:"availableStartHour"`
    AvailableEndHour int `json:"availableEndHour" form:"availableEndHour"`
}
