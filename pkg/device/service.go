package device

import (
	"context"
	"fmt"

	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/tuya/tuya-connector-go/connector"

	// qrcode "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type DeviceService interface {
	SyncDeviceList() ([]model.Device, error)
	GetDevices() ([]*DevicesGetResponse, error)
	GetDeviceById(id string) (*DevicesGetResponse, error)
	UpdateDevicesSpace(data *UpdateDeviceSpaceDto) error
}

type DeviceServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *DeviceServiceImpl) SyncDeviceList() ([]model.Device, error) {
	resp := &GetDevicesResponse{}
	err := connector.MakeGetRequest(
		context.Background(),
		connector.WithAPIUri("/v1.0/users/eu1710916580378iZqjE/devices?page_no=1&page_size=20"),
		connector.WithResp(resp))

	if err != nil {
		return nil, err
	}

	var devicesToAdd []model.Device

	// Check if devices are registered in the database and add new ones to devicesToAdd
	for _, item := range resp.Result {
		// Check if the device exists in the database
		var existingDevice model.Device
		if err := s.Db.Where("provider_device_id = ?", item.ID).First(&existingDevice).Error; err == gorm.ErrRecordNotFound {
			device := model.Device{
				Name:             item.Name,
				ProductName:      item.ProductName,
				ProviderDeviceID: item.ID,
				Online:           item.Online,
			}
			devicesToAdd = append(devicesToAdd, device)
		}
	}

	// Save new devices to the database
	for _, newDevice := range devicesToAdd {
		// Create a new device record in the database
		fmt.Println("DEVICE ADDED: ", newDevice.Name)
		if err := s.Db.Create(&newDevice).Error; err != nil {
			// Handle error
			return nil, err
		}
	}

	return devicesToAdd, nil
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
			Id:               device.ID,
			ProviderDeviceID: device.ProviderDeviceID,
			Name:             device.Name,
			Online:           device.Online,
			ProductName:      device.ProductName,
			SpaceID:          device.SpaceID,
		})
	}

	return response, nil
}

func (s *DeviceServiceImpl) GetDeviceById(id string) (*DevicesGetResponse, error) {
	var device model.Device

	if err := s.Db.First(&device, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &DevicesGetResponse{
		Id:               device.ID,
		ProviderDeviceID: device.ProviderDeviceID,
		Name:             device.Name,
		Online:           device.Online,
		ProductName:      device.ProductName,
		SpaceID:          device.SpaceID,
	}, nil
}

func (s *DeviceServiceImpl) UpdateDevicesSpace(data *UpdateDeviceSpaceDto) error {
	var device model.Device

	if err := s.Db.First(&device, "id = ?", data.DeviceId).Error; err != nil {
		return err
	}

	device.SpaceID = &data.SpaceId

	if err := s.Db.Save(&device).Error; err != nil {
		return err
	}

	return nil
}
