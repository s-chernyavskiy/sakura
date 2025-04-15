package db

import (
	"fmt"
	"sync"
)

type DB struct {
	file map[string]*DataNode
	m    sync.Mutex
}

func NewDB() *DB {
	return &DB{file: make(map[string]*DataNode)}
}

type KeyNotFoundError struct {
	key string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.key)
}

func (db *DB) Get(key string) (*DataNode, error) {
	db.m.Lock()
	defer db.m.Unlock()

	if v, ok := db.file[key]; ok {
		return v, nil
	}

	return nil, &KeyNotFoundError{key: key}
}

func (db *DB) Set(key string, val *DataNode) {
	db.m.Lock()
	defer db.m.Unlock()

	db.file[key] = val
}

func (db *DB) Delete(keys []string) bool {
	db.m.Lock()
	defer db.m.Unlock()

	for _, k := range keys {
		if _, ok := db.file[k]; !ok {
			return false
		}

		delete(db.file, k)
	}

	return true
}

func (db *DB) Exists(key string) bool {
	db.m.Lock()
	defer db.m.Unlock()

	if _, ok := db.file[key]; ok {
		return true
	}

	return false
}
