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
	GetAll() (map[types.ClientID]types.Client, error)
	Create(newClient types.Client) error
	Update(Client types.Client) error
	Delete(ID types.ClientID) error
}

type Conf struct {
	Addr string `json:"addr"`
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
	mux.HandleFunc("/api/client/get/{id:[0-9]+}", server.getClient).Methods("GET")
	mux.HandleFunc("/api/client/get", server.getAllClients).Methods("GET")
	mux.HandleFunc("/api/client/create", server.createClient).Methods("POST")
	mux.HandleFunc("/api/client/update/{id:[0-9]+}", server.updateClient).Methods("PATCH")
	mux.HandleFunc("/api/client/delete/{id:[0-9]+}", server.deleteClient).Methods("DELETE")

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

	w.Header().Set("Content-Type", "application/json")

	answer, err := json.Marshal(&Client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
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

	w.Header().Set("Content-Type", "application/json")
	answer, err := json.Marshal(&Clients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.ClientController.Create(Client)
	if err != nil {
		log.Println("Failed to create Client")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// if err := jsonapi.MarshalPayload(w, struct{}{}); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func (s *Server) updateClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.ClientController.Update(Client)
	if err != nil {
		log.Println("Failed to create Client")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) deleteClient(w http.ResponseWriter, r *http.Request) {
	Client := types.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &Client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.ClientController.Delete(Client.ID)
	if err != nil {
		log.Println("Failed to create Client")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
