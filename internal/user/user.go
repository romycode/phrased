package user

import (
	"fmt"
	"sync"
)

type User struct {
	ID string
}

func NewUser(ID string) *User {
	return &User{ID: ID}
}

type Repository interface {
	FindById(ID string) (*User, error)
	Save(u *User) error
}

type InMemoryRepository struct {
	mu sync.Mutex
	db map[string]*User
}

func NewInMemoryRepository(db map[string]*User) *InMemoryRepository {
	return &InMemoryRepository{db: db}
}

func (ur *InMemoryRepository) Save(u *User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	ur.db[u.ID] = u
	return nil
}

func (ur *InMemoryRepository) FindById(ID string) (*User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	if _, ok := ur.db[ID]; !ok {
		return nil, fmt.Errorf("userRepository: not found user with ID: %s", ID)
	}

	return ur.db[ID], nil
}
