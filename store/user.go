package store

import (
	"galileo/types"
	"sync"
)

type userStore struct {
	mx      sync.RWMutex
	byEmail map[string]*types.User
	byToken map[string]*types.User
}

func NewUser() *types.User {
	return &types.User{
		Token:   types.NewToken(),
		Devices: NewDeviceStore(),
	}
}

func NewUserStore() types.UserStore {
	return &userStore{
		byEmail: make(map[string]*types.User),
		byToken: make(map[string]*types.User),
	}
}

func (store *userStore) Add(v *types.User) {
	store.mx.Lock()
	defer store.mx.Unlock()

	store.byEmail[v.Email] = v
	store.byToken[v.Token] = v
}

func (store *userStore) FindByEmail(k string) (*types.User, error) {
	store.mx.RLock()
	defer store.mx.RUnlock()

	v, ok := store.byEmail[k]

	if ok {
		return v, nil
	}

	return nil, ErrNotFound
}

func (store *userStore) FindByToken(k string) (*types.User, error) {
	store.mx.RLock()
	defer store.mx.RUnlock()

	v, ok := store.byToken[k]

	if ok {
		return v, nil
	}

	return nil, ErrNotFound
}

func (store *userStore) AppendDevice(userToken string, device *types.Device) error {
	user, err := store.FindByToken(userToken)

	if err != nil {
		return nil
	}

	user.AppendDevice(device)

	return nil
}
