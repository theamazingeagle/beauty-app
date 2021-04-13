package client

import (
	"bbs-back/internal/types"
	"errors"
	"log"
)

type Storage interface {
	GetClientByID(ID int64) (types.Client, bool, error)
	GetAll() (map[int64]types.Client, error)
	Create(newClient types.Client) error
	Update(Client types.Client) error
	DeleteByID(ID int64) error
}

var (
	ErrClientNotFound   = errors.New("Client not found")
	ErrFailedToFetch    = errors.New("Failed to fetch Client info")
	ErrFailedToFetchAll = errors.New("Failed to fetch all Clients info")
	ErrFailedToCreate   = errors.New("Failed to create new Client record")
	ErrFailedToUpdate   = errors.New("Failed to update Client info")
	ErrFailedToDelete   = errors.New("Failed to delete Client info")
)

type ClientModel struct {
	storage Storage
}

func New(s Storage) ClientModel {
	return ClientModel{storage: s}
}

func (cm *ClientModel) Get(ID int64) (types.Client, error) {
	Client, exist, err := cm.storage.GetClientByID(ID)
	if err != nil {
		log.Println("Failed to get Client info")
		return types.Client{}, ErrFailedToFetch
	}
	if !exist {
		log.Println("Client not found")
		return types.Client{}, ErrClientNotFound
	}
	return Client, nil
}

func (cm *ClientModel) GetAll() (map[int64]types.Client, error) {
	allClients, err := cm.storage.GetAll()
	if err != nil {
		log.Println("Failed to get all Clients info")
		return nil, ErrFailedToFetchAll
	}
	return allClients, nil
}

func (cm *ClientModel) Create(Client types.Client) error {
	err := cm.storage.Create(Client)
	if err != nil {
		log.Println("Failed to create Client record")
		return ErrFailedToCreate
	}
	return nil
}

func (cm *ClientModel) Update(Client types.Client) error {
	err := cm.storage.Update(Client)
	if err != nil {
		log.Println("Failed to update Client record")
		return ErrFailedToUpdate
	}
	return nil
}

func (cm *ClientModel) Delete(ID int64) error {
	err := cm.storage.DeleteByID(ID)
	if err != nil {
		log.Println("Failed to delete Client record")
		return ErrFailedToDelete
	}
	return nil
}