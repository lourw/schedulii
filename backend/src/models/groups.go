package models

type Groups struct {
	GroupID int `json:"groupID" form:"groupID" binding:"required"`
    GroupName string `json:"groupName" form:"groupName" binding:"required"`
    GroupURL string `json:"groupURL" form:"groupURL" binding:"required"`
    AvailabilityStart string `json:"availabilityStart" form:"availabilityStart" binding:"required"`
    AvailabilityEnd string `json:"availabilityEnd" form:"availabilityEnd" binding:"required"`
}
