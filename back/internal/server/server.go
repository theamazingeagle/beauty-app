package server

import (
	"bbs-back/internal/types"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClientController interface {
	Get(ID types.ClientID) (types.Client, error)
	GetAll() ([]types.Client, error)
	Create(newClient types.Client) error
	Update(Client types.Client) error
	Delete(ID types.ClientID) error
}

type ServiceController interface {
	Get(ID types.ServiceID) (types.Service, error)
	GetAll() ([]types.Service, error)
	Create(service types.Service) error
	Update(service types.Service) error
	Delete(ID types.ServiceID) error
}

type OrderController interface {
	Get(ID types.OrderID) (types.Order, error)
	GetAll() ([]types.Order, error)
	Create(newOrder types.Order) error
	Update(order types.Order) error
	Delete(ID types.OrderID) error
}

type Conf struct {
	Addr string `json:"addr"`
}

type Server struct {
	conf              Conf
	ClientController  ClientController
	ServiceController ServiceController
	OrderController   OrderController
	HTTPServer        http.Server
}

func New(conf Conf, clientController ClientController, orderController OrderController, serviceController ServiceController) *Server {
	server := &Server{
		conf:              conf,
		ClientController:  clientController,
		ServiceController: serviceController,
		OrderController:   orderController,
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/api/client/get/{id:[0-9]+}", server.getClient).Methods("GET")
	mux.HandleFunc("/api/client/get", server.getAllClients).Methods("GET")
	mux.HandleFunc("/api/client/create", server.createClient).Methods("POST")
	mux.HandleFunc("/api/client/update", server.updateClient).Methods("PATCH")
	mux.HandleFunc("/api/client/delete/{id:[0-9]+}", server.deleteClient).Methods("DELETE")

	mux.HandleFunc("/api/service/get/{id:[0-9]+}", server.getClient).Methods("GET")
	mux.HandleFunc("/api/service/get", server.getAllServices).Methods("GET")
	mux.HandleFunc("/api/service/create", server.createService).Methods("POST")
	mux.HandleFunc("/api/service/update", server.updateService).Methods("PATCH")
	mux.HandleFunc("/api/service/delete/{id:[0-9]+}", server.deleteService).Methods("DELETE")

	mux.HandleFunc("/api/order/get/{id:[0-9]+}", server.getOrder).Methods("GET")
	mux.HandleFunc("/api/order/get", server.getAllOrders).Methods("GET")
	mux.HandleFunc("/api/order/create", server.createOrder).Methods("POST")
	mux.HandleFunc("/api/order/update", server.updateOrder).Methods("PATCH")
	mux.HandleFunc("/api/order/delete/{id:[0-9]+}", server.deleteOrder).Methods("DELETE")

	server.HTTPServer = http.Server{Addr: conf.Addr, Handler: mux}
	return server
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler))
}

//
// Clients
//

func (s *Server) getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam client id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	clientID := types.ClientID(clientIDStr)
	if err != nil {
		log.Println("invalid client id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Client, err := s.ClientController.Get(clientID)
	if err != nil {
		log.Println("Failed to get client")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	answer, err := json.Marshal(&Client)
	if err != nil {
		log.Println("Failed to marshall client data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) getAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clients, err := s.ClientController.GetAll()
	if err != nil {
		log.Println("Failed to get clients")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	answer, err := json.Marshal(&clients)
	if err != nil {
		log.Println("Failed to marshall clients data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &Client); err != nil {
		log.Println("Failed to unmarshall request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.ClientController.Create(Client)
	if err != nil {
		log.Println("Failed to create Client")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) updateClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &Client); err != nil {
		log.Println("Failed to unmarshall request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.ClientController.Update(Client)
	if err != nil {
		log.Println("Failed to update client info")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam client id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	clientID := types.ClientID(clientIDStr)
	if err != nil {
		log.Println("invalid client id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.ClientController.Delete(clientID)
	if err != nil {
		log.Println("Failed to delete Client")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//
// Services
//

func (s *Server) getService(w http.ResponseWriter, r *http.Request) {
	serviceIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam service id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	serviceID := types.ServiceID(serviceIDStr)
	if err != nil {
		log.Println("invalid service id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service, err := s.ServiceController.Get(serviceID)
	if err != nil {
		log.Println("Failed to get service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	answer, err := json.Marshal(&service)
	if err != nil {
		log.Println("Failed to marshall service data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) getAllServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services, err := s.ServiceController.GetAll()
	if err != nil {
		log.Println("Failed to get services list")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	answer, err := json.Marshal(&services)
	if err != nil {
		log.Println("Failed to marshall service data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) createService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := types.Service{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println("Failed to unmarshall service data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.ServiceController.Create(service)
	if err != nil {
		log.Println("Failed to create service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) updateService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := types.Service{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println("Failed to unmarshall service data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.ServiceController.Update(service)
	if err != nil {
		log.Println("Failed to update service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) deleteService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serviceIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam service id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	serviceID := types.ServiceID(serviceIDStr)
	if err != nil {
		log.Println("bad pam service id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.ServiceController.Delete(serviceID)
	if err != nil {
		log.Println("Failed to delete service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//
// Orders
//

func (s *Server) getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam order id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orderID := types.OrderID(orderIDStr)
	if err != nil {
		log.Println("invalid order id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := s.OrderController.Get(orderID)
	if err != nil {
		log.Println("Failed to get order")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	answer, err := json.Marshal(&order)
	if err != nil {
		log.Println("Failed to marshall answer")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) getAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := s.OrderController.GetAll()
	if err != nil {
		log.Println("Failed to get orders")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	answer, err := json.Marshal(orders)
	if err != nil {
		log.Println("Failed to marshall answer")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	order := types.Order{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &order); err != nil {
		log.Println("Failed to unmarshall request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.OrderController.Create(order)
	if err != nil {
		log.Println("Failed to create order")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	order := types.Order{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &order); err != nil {
		log.Println("Failed to unmarshall request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.OrderController.Update(order)
	if err != nil {
		log.Println("Failed to update order")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam order id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orderID := types.OrderID(orderIDStr)
	if err != nil {
		log.Println("invalid order id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.OrderController.Delete(orderID)
	if err != nil {
		log.Println("Failed to delete order")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
