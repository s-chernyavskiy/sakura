package hashmap

import "sync"

type HashMap struct {
	m  map[string]string
	mx sync.Mutex
}

func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[string]string),
	}
}

func (m *HashMap) Set(key, value string) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.m[key] = value
}

func (m *HashMap) Get(key, value string) string {
	m.mx.Lock()
	defer m.mx.Unlock()

	if v, ok := m.m[key]; ok {
		return v
	}

	return ""
}

func (m *HashMap) Delete(keys ...string) bool {
	m.mx.Lock()
	defer m.mx.Unlock()

	for _, key := range keys {
		if _, ok := m.m[key]; !ok {
			return false
		}

		delete(m.m, key)
	}

	return true
}

func (m *HashMap) Exists(key string) bool {
	m.mx.Lock()
	defer m.mx.Unlock()
	_, ok := m.m[key]

	return ok
}
