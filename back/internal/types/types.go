package types

type ClientID int64

type Client struct {
	ID   ClientID `json:"id"`
	Name string   `json:"name"`
}

type ServiceID int64
type Currency int64

type Service struct {
	ID   ServiceID `json:"id"`
	Name string    `json:"name"`
	Cost Currency  `json:"cost"`
}

type OrderID int64

type Order struct {
	ID           OrderID   `json:"id"`
	ServiceID    ServiceID `json:"service_id"`
	ClientID     ClientID  `json:"client_id"`
	CreationTime string    `json:"creation_time"`
	OrderTime    string    `json:"order_time"`
}
