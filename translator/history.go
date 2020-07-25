package translator

import (
	"sort"
	"sync"
)

var Storage *History

// History is a data structure for safely storing translations
type History struct {
	Data  map[string]string
	Mutex sync.RWMutex
}

func init() {
	s := &History{
		Data:  nil,
		Mutex: sync.RWMutex{},
	}

	Storage = s
}

// Store is a method for saving a translation
func (h *History) Store(ew string, gw string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	if h.Data == nil {
		h.Data = make(map[string]string, 0)
	}
	h.Data[ew] = gw
}

// Load is a method that returns a translation if stored already or false if it's not
func (h *History) Load(en string) string {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	str := h.Data[en]
	return str
}

// Load is a method that returns a translation if stored already or false if it's not
func (h *History) GetOrderedMap() map[string]string {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	// Put the keys in a slice and sort it.
	keys := make([]string, 0, len(h.Data))
	for key := range h.Data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	result := make(map[string]string, len(keys))
	for _, key := range keys {
		result[key] = h.Data[key]
	}

	return result
}

// Clear map so we can avoid caching between instances.
func (h *History) Clear() {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Data = nil
}

func ClearStorage() {
	Storage.Clear()
}
