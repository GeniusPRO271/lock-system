package space

type SpaceDTO struct {
	ID            uint
	Name          string
	ParentSpaceID *uint
	SubSpaces     []SpaceDTO
}

type UpdateSpaceDTO struct {
	Name          string `json:"Name" binding:"required"`
	ParentSpaceID *uint  `json:"ParentSpaceID"`
}
