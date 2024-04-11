package db

import (
	"errors"
	"sync"
)

var ErrRecordNotFound = errors.New("record not found")

type InMemoryUserRepository struct {
	mutex sync.RWMutex
	users map[int]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]User)}
}

func (r *InMemoryUserRepository) CreateUser(u User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	newID := len(r.users) + 1
	u.ID = newID
	r.users[newID] = u
	return nil
}

func (r *InMemoryUserRepository) GetUserByID(userID int) (User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	u, ok := r.users[userID]
	if !ok {
		return User{}, ErrRecordNotFound
	}
	return u, nil
}
