package client

import (
	"beauty/internal/service/memorydb"
	"beauty/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientModel(t *testing.T) {
	storage := memorydb.New()
	ClientController := New(&storage)
	// get
	testValue := types.Client{2, "Ksenia"}
	receivedValue, err := ClientController.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, testValue, receivedValue)
	// create
	testValue = types.Client{6, "Rita"}
	err = ClientController.Create(testValue)
	assert.Nil(t, err)
	receivedValue, err = ClientController.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, testValue, receivedValue)
	// update
	testValue = types.Client{6, "Dana"}
	err = ClientController.Update(types.Client{6, "Dana"})
	assert.Nil(t, err)
	receivedValue, err = ClientController.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, testValue, receivedValue)
	// delete TODO
	// err = ClientController.Delete(types.ClientID(6))
	// assert.Nil(t, err)
	// Client, err := ClientController.Get(6)
	// fmt.Println(Client)
	// assert.Equal(t, ErrClientNotFound, err)
}
