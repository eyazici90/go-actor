package infrastructure

import (
	"sync"
)

type Persister interface {
	Create(Stock) error
	Update(Stock) error
	Get(string) Stock
}

var stocks = make(map[string]Stock)

var lockMutex = new(sync.RWMutex)

type inMemoryStore struct{}

func NewPersister() Persister {
	return &inMemoryStore{}
}

func (i *inMemoryStore) Get(id string) Stock {
	return stocks[id]
}

func (i *inMemoryStore) Update(s Stock) error {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	stocks[s.Id] = s

	return nil
}

func (i *inMemoryStore) Create(s Stock) error {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	stocks[s.Id] = s
	return nil
}
