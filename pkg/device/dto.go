package device

type DevicesGetResponse struct {
	Id               uint   `binding:"required"`
	Online           bool   `binding:"required"`
	Name             string `binding:"required"`
	ProductName      string `binding:"required"`
	ProviderDeviceID string `binding:"required"`
	SpaceID          *uint
}

type ResultItem struct {
	ID          string         `json:"id"`
	UUID        string         `json:"uuid"`
	UID         string         `json:"uid"`
	BizType     int            `json:"biz_type"`
	Name        string         `json:"name"`
	TimeZone    string         `json:"time_zone"`
	IP          string         `json:"ip"`
	LocalKey    string         `json:"local_key"`
	Sub         bool           `json:"sub"`
	CreateTime  int64          `json:"create_time"`
	UpdateTime  int64          `json:"update_time"`
	ActiveTime  int64          `json:"active_time"`
	Status      []StatusObject `json:"status"`
	OwnerID     string         `json:"owner_id"`
	ProductID   string         `json:"product_id"`
	ProductName string         `json:"product_name"`
	Category    string         `json:"category"`
	Icon        string         `json:"icon"`
	Online      bool           `json:"online"`
}

type StatusObject struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}

type GetDevicesResponse struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Success bool         `json:"success"`
	Result  []ResultItem `json:"result"`
	T       int64        `json:"t"`
}
