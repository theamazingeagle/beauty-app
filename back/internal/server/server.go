package server

import (
	"bbs-back/internal/types"
	"log"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

type ClientController interface {
	Get(ID types.ClientID) (types.Client, error)
	GetAll() (map[types.ClientID]types.Client, error)
	Create(newClient types.Client) error
	Update(Client types.Client) error
	Delete(ID types.ClientID) error
}

type Conf struct {
	Addr string
}

type Server struct {
	conf             Conf
	ClientController ClientController
	HTTPServer       http.Server
}

func New(conf Conf, ClientController ClientController) *Server {
	server := &Server{
		conf:             conf,
		ClientController: ClientController,
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/client/get/{id:[0-9]+}", server.getClient).Methods("GET")
	mux.HandleFunc("/client/get", server.getAllClients).Methods("GET")
	mux.HandleFunc("/client/create", server.createClient).Methods("POST")
	mux.HandleFunc("/client/update/{id:[0-9]+}", server.updateClient).Methods("PATCH")
	mux.HandleFunc("/client/delete/{id:[0-9]+}", server.deleteClient).Methods("DELETE")

	server.HTTPServer = http.Server{Addr: conf.Addr, Handler: mux}
	return server
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler))
}

func (s *Server) getClient(w http.ResponseWriter, r *http.Request) {
	clientIDStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("bad pam Client id")
	}
	clientID := types.ClientID(clientIDStr)
	if err != nil {
		log.Println("invalid Client id")
		return
	}

	Client, err := s.ClientController.Get(clientID)
	if err != nil {
		log.Println("Failed to get Client")
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getAllClients(w http.ResponseWriter, r *http.Request) {
	tempClients, err := s.ClientController.GetAll()
	if err != nil {
		log.Println("Failed to get Client")
	}
	// we need slice for data marshalling
	Clients := []*types.Client{}
	for _, copy := range tempClients {
		newCopy := copy // ?
		Clients = append(Clients, &newCopy)
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, Clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}

	if err := jsonapi.UnmarshalPayload(r.Body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := s.ClientController.Create(Client)
	if err != nil {
		log.Println("Failed to create Client")
	}
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	// if err := jsonapi.MarshalPayload(w, struct{}{}); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func (s *Server) updateClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}

	if err := jsonapi.UnmarshalPayload(r.Body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := s.ClientController.Update(Client)
	if err != nil {
		log.Println("Failed to create Client")
	}
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) deleteClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}

	if err := jsonapi.UnmarshalPayload(r.Body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := s.ClientController.Delete(Client.ID)
	if err != nil {
		log.Println("Failed to create Client")
	}
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)
}
