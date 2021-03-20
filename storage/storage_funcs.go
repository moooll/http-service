package storage

import (
	"errors"
	"fmt"
	"sync"

	"http-service/dto"
)

type Storage struct {
	storage map[string]string
	m       sync.Mutex
}

var storage Storage

func (db *Storage) Upsert(req []dto.Request) {
	if db.storage == nil {
		storage.storage = make(map[string]string)
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

// todo: mb change map to dto.struct in return value
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

//todo: mb check db.storage length??
// todo: mb change map to dto.struct in return value
func (db *Storage) List() (map[string]string, error) {
	// if len(db.storage) == 0 {
	// 	return nil, errors.New("no elements in a storage")
	// }
	return db.storage, nil
}

func (db *Storage) FillStorage() error {
	db.storage = make(map[string]string)
	// rand:
	// k := make([]byte, 20)
	// v := make([]byte, 45)
	// for i := 0; i < 150; i++ {
	// 	_, err := rand.Read(k)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	_, err = rand.Read(v)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	db.storage[string(k)] = string(v)
	// }
	//
	db.storage["lol"] = "kek"
	fmt.Println(db.storage)
	return nil
}
