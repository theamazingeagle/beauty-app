package memorydb

import (
	"bbs-back/internal/types"
	"errors"
)

var (
	ErrNotFound = errors.New("Record not found")
)

type MemoryDB struct {
	ClientCollection map[types.ClientID]types.Client
	currentClient    int64
}

func New() MemoryDB {
	return MemoryDB{
		ClientCollection: map[types.ClientID]types.Client{
			1: {1, "Katya"},
			2: {2, "Ksenia"},
			3: {3, "Liza"},
			4: {4, "Tanya"},
			5: {5, "Nina"},
		},
		currentClient: 5,
	}
}

func (mdb *MemoryDB) GetClientByID(ID types.ClientID) (types.Client, bool, error) {
	value, exist := mdb.ClientCollection[ID]
	if exist {
		return value, true, nil
	}
	return types.Client{}, false, nil
}

func (mdb *MemoryDB) GetClients() (map[types.ClientID]types.Client, error) {
	return mdb.ClientCollection, nil
}

func (mdb *MemoryDB) CreateClient(newClient types.Client) error {
	mdb.currentClient++
	newClient.ID = types.ClientID(mdb.currentClient)
	mdb.ClientCollection[newClient.ID] = newClient
	return nil
}

func (mdb *MemoryDB) UpdateClient(Client types.Client) error {
	mdb.ClientCollection[Client.ID] = Client
	return nil
}

func (mdb *MemoryDB) DeleteClientByID(ID types.ClientID) error {
	delete(mdb.ClientCollection, ID)
	return nil
}
