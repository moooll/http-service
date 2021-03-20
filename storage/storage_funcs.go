package storage

import (
	"errors"
	"sync"

	"http-service/dto"
)

type Storage struct {
	storage map[string]string
	m       sync.Mutex
}

var storage Storage

func (db *Storage) Upsert(req []dto.Request) {
	if db.storage == nil || len(db.storage) == 0 {
		db.storage = make(map[string]string)
	}
	for _, v := range req {
		db.m.Lock()
		db.storage[v.Key] = v.Value
		db.m.Unlock()
	}
}

func (db *Storage) Delete(keys []string) error {
	for _, v := range keys {
		_, found := db.storage[v]
		if found == false {
			return errors.New("the record with this key wasn't found in the storage")
		}
		db.m.Lock()
		delete(db.storage, v)
		db.m.Unlock()
	}
	return nil
}

func (db *Storage) Get(keys []string) (map[string]string, error) {
	values := make(map[string]string)
	for _, v := range keys {
		db.m.Lock()
		value, found := db.storage[v]
		db.m.Unlock()
		if found {
			values[v] = value
		} else {
			return nil, errors.New("nothing found")
		}
	}
	return values, nil
}

func (db *Storage) List() (map[string]string, error) {
	if len(db.storage) == 0 {
		return nil, errors.New("no elements in a storage")
	}
	return db.storage, nil
}
