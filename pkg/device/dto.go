package device

type DevicesGetResponse struct {
	Id    uint   `binding:"required"`
	Name  string `binding:"required"`
	Model string `binding:"required"`
	Type  string `binding:"required"`
}
