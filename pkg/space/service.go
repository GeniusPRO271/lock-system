package space

import (
	"log"
	"strconv"

	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"gorm.io/gorm"
)

type SpaceService interface {
	CreateSpace(space model.Space) error
	GetSpaceByID(spaceID string) (*SpaceDTO, error)
	GetAllSpaces() ([]SpaceDTO, error)
	LoadSubSpaces(space *model.Space) error
	SpaceToDTO(spaceData model.Space) SpaceDTO
}

type SpaceServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *SpaceServiceImpl) CreateSpace(space model.Space) error {
	// Check if the space has a parentID
	if space.ParentSpaceID != nil {
		// Retrieve the parent space
		var parentSpace model.Space
		if err := s.Db.First(&parentSpace, *space.ParentSpaceID).Error; err != nil {
			return err
		}
		// Set the level of the new space to the level of the parent space plus 1
		space.Level = parentSpace.Level + 1
		log.Println("Adding lvl +1 ")
	} else {
		// If parentID is null, set the level to 1
		log.Println("Adding lvl +1 ")
		space.Level = 1
	}

	// Create the space with the associated whitelist
	log.Println("The new space lvl = %d", space.Level)
	if err := s.Db.Create(&space).Error; err != nil {
		return err
	}

	return nil
}

func (s *SpaceServiceImpl) GetSpaceByID(spaceID string) (*SpaceDTO, error) {
	log.Printf("GetSpaceByID")
	id, err := strconv.ParseUint(spaceID, 10, 64)
	if err != nil {
		return nil, err
	}

	var space model.Space
	if err := s.Db.First(&space, uint(id)).Error; err != nil {
		return nil, err
	}

	// Load subspaces recursively
	if err := s.LoadSubSpaces(&space); err != nil {
		return nil, err
	}

	spaceDTO := s.SpaceToDTO(space)

	return &spaceDTO, nil
}

func (s *SpaceServiceImpl) GetAllSpaces() ([]SpaceDTO, error) {
	var spaces []*model.Space
	if err := s.Db.Find(&spaces).Error; err != nil {
		return nil, err
	}

	var spaceDTOs []SpaceDTO
	for _, space := range spaces {
		if err := s.LoadSubSpaces(space); err != nil {
			return nil, err
		}
		spaceDTO := s.SpaceToDTO(*space)
		spaceDTOs = append(spaceDTOs, spaceDTO)
	}

	return spaceDTOs, nil
}

func (s *SpaceServiceImpl) LoadSubSpaces(space *model.Space) error {
	if space == nil {
		return nil
	}

	var subSpaces []*model.Space
	if err := s.Db.Preload("Whitelist.Users").Preload("Devices").Where("parent_space_id = ?", space.ID).Find(&subSpaces).Error; err != nil {
		return err
	}

	for _, subSpace := range subSpaces {
		if err := s.LoadSubSpaces(subSpace); err != nil {
			return err
		}
	}

	space.SubSpaces = subSpaces
	return nil
}

func (s *SpaceServiceImpl) SpaceToDTO(spaceData model.Space) SpaceDTO {
	var subSpacesDTO []SpaceDTO
	for _, subSpace := range spaceData.SubSpaces {
		subSpaceDTO := s.SpaceToDTO(*subSpace)
		subSpacesDTO = append(subSpacesDTO, subSpaceDTO)
	}

	spaceDTO := SpaceDTO{
		ID:            spaceData.ID,
		Name:          spaceData.Name,
		ParentSpaceID: spaceData.ParentSpaceID,
		SubSpaces:     subSpacesDTO,
	}

	return spaceDTO
}
