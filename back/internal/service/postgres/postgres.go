package postgres

import (
	"beauty/internal/types"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	ErrNotFound  = errors.New("Not found")
	ErrQueryExec = errors.New("Failed to exec query")
)

type PostgresConf struct {
	HostName   string `json:"hostname"`
	DriverName string `json:"drivername"`
	DBName     string `json:"dbname"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Port       int    `json:"port"`
}

type Postgres struct {
	conn *sql.DB
}

func New(conf PostgresConf) (*Postgres, error) {
	postgres := &Postgres{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.HostName, conf.Port, conf.UserName, conf.Password, conf.DBName)
	var err error
	postgres.conn, err = sql.Open(conf.DriverName, dsn)
	if err != nil {
		log.Println("NewPostgres err: ", err)
		return &Postgres{}, err
	}
	return postgres, nil
}

func (p *Postgres) CreateClient(client types.Client) error {
	_, err := p.conn.Exec(`INSERT INTO clients(name)
						  VALUES($1);`,
		client.Name)
	if err != nil {
		log.Println("Failed to create client's record", err)
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) GetClientByID(clientID types.ClientID) (types.Client, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM clients WHERE id=$1 ;", clientID)
	client := types.Client{}
	err := row.Scan(&client.ID, &client.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Client not found")
			return types.Client{}, false, ErrNotFound
		}
		log.Println("Failed to get client info")
		return types.Client{}, false, ErrQueryExec
	}
	return client, true, nil
}

func (p *Postgres) GetClients() ([]types.Client, error) {
	rows, err := p.conn.Query("SELECT * FROM clients ;")
	clients := []types.Client{}
	client := types.Client{}
	for rows.Next() {
		err = rows.Scan(&client.ID, &client.Name)
		if err != nil {
			log.Println("Failed to get clients info")
			return []types.Client{}, ErrQueryExec
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (p *Postgres) UpdateClient(client types.Client) error {
	_, err := p.conn.Exec("UPDATE messages SET name = $1 WHERE id = $3", client.Name, client.ID)
	if err != nil {
		log.Println("Failed to update client info")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) DeleteClientByID(clientID types.ClientID) error {
	_, err := p.conn.Exec("DELETE FROM messages WHERE id=$1", clientID)
	if err != nil {
		log.Println("Failed to delete client info")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) CreateService(service types.Service) error {
	_, err := p.conn.Exec(`INSERT INTO services(name, cost)
						VALUES($1, $2);`,
		service.Name, service.Cost)
	if err != nil {
		log.Println("Failed to create service record", err)
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) GetServiceByID(serviceID types.ServiceID) (types.Service, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM services WHERE id=$1 ;", serviceID)
	service := types.Service{}
	err := row.Scan(&service.ID, &service.Name, &service.Cost)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Service not found")
			return types.Service{}, false, ErrNotFound
		}
		log.Println("Failed to get service info")
		return types.Service{}, false, ErrQueryExec
	}
	return service, true, nil
}

func (p *Postgres) GetServices() ([]types.Service, error) {
	rows, err := p.conn.Query("SELECT * FROM services ;")
	services := []types.Service{}
	service := types.Service{}
	for rows.Next() {
		err = rows.Scan(&service.ID, &service.Name, &service.Cost)
		if err != nil {
			log.Println("Failed to get clients info")
			return []types.Service{}, ErrQueryExec
		}
		services = append(services, service)
	}
	return services, nil
}

func (p *Postgres) UpdateService(service types.Service) error {
	_, err := p.conn.Exec("UPDATE services SET name = $1, cost = $2 WHERE id = $3",
		service.Name, service.Cost, service.ID)
	if err != nil {
		log.Println("Failed to update service info")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) DeleteServiceByID(serviceID types.ServiceID) error {
	_, err := p.conn.Exec("DELETE FROM services WHERE id=$1", serviceID)
	if err != nil {
		log.Println("Failed to delete service info")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) CreateOrder(order types.Order) error {
	_, err := p.conn.Exec(`INSERT INTO orders(service_id, client_id, creation_time, order_time)
						  VALUES($1, $2, $3, $4);`,
		order.ServiceID, order.ClientID, order.CreationTime, order.OrderTime)
	if err != nil {
		log.Println("Failed to create order record", err)
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) GetOrderByID(orderID types.OrderID) (types.Order, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM orders WHERE id=$1 ;", orderID)
	order := types.Order{}
	err := row.Scan(&order.ID, &order.ServiceID, &order.ClientID, &order.CreationTime, &order.OrderTime)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Order not found")
			return types.Order{}, false, ErrNotFound
		}
		log.Println("Failed to get order info")
		return types.Order{}, false, ErrQueryExec
	}
	return order, true, nil
}

func (p *Postgres) GetOrders() ([]types.Order, error) {
	rows, err := p.conn.Query("SELECT * FROM orders ;")
	orders := []types.Order{}
	order := types.Order{}
	for rows.Next() {
		err = rows.Scan(&order.ID, &order.ServiceID, &order.ClientID, &order.CreationTime, &order.OrderTime)
		if err != nil {
			log.Println("Failed to get clients info")
			return []types.Order{}, ErrQueryExec
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (p *Postgres) UpdateOrder(order types.Order) error {
	_, err := p.conn.Exec("UPDATE orders SET order_time = $1 WHERE id = $2", order.OrderTime, order.ID)
	if err != nil {
		log.Println("Failed to update order info")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) DeleteOrderByID(orderID types.OrderID) error {
	_, err := p.conn.Exec("DELETE FROM orders WHERE id=$1", orderID)
	if err != nil {
		log.Println("Failed to delete order info")
		return ErrQueryExec
	}
	return nil
}
