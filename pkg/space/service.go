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
}

type SpaceServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *SpaceServiceImpl) CreateSpace(space model.Space) error {

	// Create the space with the associated whitelist
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
	if err := s.Db.Preload("SubSpaces").First(&space, uint(id)).Error; err != nil {
		return nil, err
	}

	spaceDTO := spaceToDTO(space)

	return &spaceDTO, nil
}

func (s *SpaceServiceImpl) GetAllSpaces() ([]SpaceDTO, error) {
	var spaces []*model.Space
	if err := s.Db.Find(&spaces).Error; err != nil {
		return nil, err
	}

	var spaceDTOs []SpaceDTO
	for _, space := range spaces {
		if err := s.loadSubSpaces(space); err != nil {
			return nil, err
		}
		spaceDTO := spaceToDTO(*space)
		spaceDTOs = append(spaceDTOs, spaceDTO)
	}

	return spaceDTOs, nil
}

func (s *SpaceServiceImpl) loadSubSpaces(space *model.Space) error {
	if space == nil {
		return nil
	}

	var subSpaces []*model.Space
	if err := s.Db.Where("parent_space_id = ?", space.ID).Find(&subSpaces).Error; err != nil {
		return err
	}

	for _, subSpace := range subSpaces {
		if err := s.loadSubSpaces(subSpace); err != nil {
			return err
		}
	}

	space.SubSpaces = subSpaces
	return nil
}

func spaceToDTO(spaceData model.Space) SpaceDTO {
	var subSpacesDTO []SpaceDTO
	for _, subSpace := range spaceData.SubSpaces {
		subSpaceDTO := spaceToDTO(*subSpace)
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
