package member

type MemberInput struct {
	RegionID int    `json:"region_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}

type GetDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
