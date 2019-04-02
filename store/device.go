package store

import (
	"errors"
	"galileo/types"
	"sync"
)

type deviceStore struct {
	mx      sync.RWMutex
	byID    map[uint64]*types.Device
	byToken map[string]*types.Device
	c       uint64
}

func NewDeviceStore() types.DeviceStore {
	return &deviceStore{
		byID:    make(map[uint64]*types.Device),
		byToken: make(map[string]*types.Device),
		c:       0,
	}
}

func (store *deviceStore) Incr() uint64 {
	store.c += 1

	return store.c
}

func (store *deviceStore) Add(v *types.Device) {
	store.mx.Lock()
	defer store.mx.Unlock()
	v.ID = store.Incr()
	store.byID[v.ID] = v
	store.byToken[v.Token] = v
}

func (store *deviceStore) FindByID(k uint64) (*types.Device, error) {
	store.mx.RLock()
	defer store.mx.RUnlock()

	v, ok := store.byID[k]

	if ok {
		return v, nil
	}

	return nil, errors.New("record not found")
}

func (store *deviceStore) FindByToken(k string) (*types.Device, error) {
	store.mx.RLock()
	defer store.mx.RUnlock()

	v, ok := store.byToken[k]

	if ok {
		return v, nil
	}

	return nil, errors.New("record not found")
}

func (store *deviceStore) All() []*types.Device {
	var devices []*types.Device
	for _, d := range store.byID {
		devices = append(devices, d)
	}

	return devices
}
