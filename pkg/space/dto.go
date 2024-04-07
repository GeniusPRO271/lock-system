package space

type SpaceDTO struct {
	ID            uint
	Name          string
	ParentSpaceID *uint
	SubSpaces     []SpaceDTO
}
