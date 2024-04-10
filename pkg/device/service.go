package device

import (
	model "github.com/GeniusPRO271/lock-system/pkg/database"
	// qrcode "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type DeviceService interface {
	CreateDevice(device model.Device) error
	GetDevices() ([]*DevicesGetResponse, error)
	GetDeviceById(id string) (*DevicesGetResponse, error)
}

type DeviceServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *DeviceServiceImpl) CreateDevice(device model.Device) error {

	//Generate QR Code to the instructions

	// "https://example.org" == "https://localhost:8080/device/{id}/instructions"
	// png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256) // qr code dimensions 256x256

	// if err != nil {
	// 	return err
	// }

	//Add intructions

	//Add device Info

	//Add whitelist

	if err := s.Db.Create(&device).Error; err != nil {
		return err
	}

	return nil
}

func (s *DeviceServiceImpl) GetDevices() ([]*DevicesGetResponse, error) {
	var devices []*model.Device

	// Fetch all users from the database
	if err := s.Db.Find(&devices).Error; err != nil {
		return nil, err
	}

	// Prepare the response
	response := []*DevicesGetResponse{}

	// Populate the response slice
	for _, device := range devices {
		response = append(response, &DevicesGetResponse{
			Id:    device.ID,
			Name:  device.Name,
			Type:  device.Type,
			Model: device.DeviceModel,
		})
	}

	return response, nil
}

func (s *DeviceServiceImpl) GetDeviceById(id string) (*DevicesGetResponse, error) {
	var deivce model.Device

	if err := s.Db.First(&deivce, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &DevicesGetResponse{
		Id:    deivce.ID,
		Name:  deivce.Name,
		Type:  deivce.Type,
		Model: deivce.DeviceModel,
	}, nil
}
