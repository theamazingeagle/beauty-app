package service

import (
	"beauty/internal/types"
	"errors"
	"log"
)

type Storage interface {
	GetServiceByID(ID types.ServiceID) (types.Service, bool, error)
	GetServices() ([]types.Service, error)
	CreateService(newservice types.Service) error
	UpdateService(service types.Service) error
	DeleteServiceByID(ID types.ServiceID) error
}

var (
	ErrserviceNotFound  = errors.New("service not found")
	ErrFailedToFetch    = errors.New("Failed to fetch service info")
	ErrFailedToFetchAll = errors.New("Failed to fetch all services info")
	ErrFailedToCreate   = errors.New("Failed to create new service record")
	ErrFailedToUpdate   = errors.New("Failed to update service info")
	ErrFailedToDelete   = errors.New("Failed to delete service info")
)

type serviceModel struct {
	storage Storage
}

func New(s Storage) serviceModel {
	return serviceModel{storage: s}
}

func (sm *serviceModel) Get(ID types.ServiceID) (types.Service, error) {
	service, exist, err := sm.storage.GetServiceByID(ID)
	if err != nil {
		log.Println("Failed to get service info")
		return types.Service{}, ErrFailedToFetch
	}
	if !exist {
		log.Println("service not found")
		return types.Service{}, ErrserviceNotFound
	}
	return service, nil
}

func (sm *serviceModel) GetAll() ([]types.Service, error) {
	allservices, err := sm.storage.GetServices()
	if err != nil {
		log.Println("Failed to get all services info")
		return nil, ErrFailedToFetchAll
	}
	return allservices, nil
}

func (sm *serviceModel) Create(service types.Service) error {
	err := sm.storage.CreateService(service)
	if err != nil {
		log.Println("Failed to create service record")
		return ErrFailedToCreate
	}
	return nil
}

func (sm *serviceModel) Update(service types.Service) error {
	err := sm.storage.UpdateService(service)
	if err != nil {
		log.Println("Failed to update service record")
		return ErrFailedToUpdate
	}
	return nil
}

func (sm *serviceModel) Delete(ID types.ServiceID) error {
	err := sm.storage.DeleteServiceByID(ID)
	if err != nil {
		log.Println("Failed to delete service record")
		return ErrFailedToDelete
	}
	return nil
}
