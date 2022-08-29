package region

type RegionInput struct {
	Name string `json:"name" binding:"required"`
}

type GetDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
