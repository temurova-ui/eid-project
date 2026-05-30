package storage

import (
	"encoding/json"
	"errors"
	"os"
	"project/models"
	"sync"
)

type UserStorage struct {
	Mu       sync.Mutex
	FileName string
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	file, err := os.ReadFile(s.FileName)
	if err != nil {
		return nil, err
	}
	var users []models.User

	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

	func (s *UserStorage) GetByID(id int) (*models.User, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	file, err := os.ReadFile(s.FileName)
	if err != nil {
		return nil, err
	}
	var users []models.User
	var user models.User

	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}
	for _, val := range users {
		if val.ID == id {
			user = val
			return &user, nil
		}
	}
	return  nil, errors.New("user not found")
}

func (s *UserStorage) Create(user models.User) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	file, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	var users []models.User

	err = json.Unmarshal(file, &users)
	if err != nil {
		return err
	}
	users = append(users, user)

	file1, err := json.Marshal(users)
	if err != nil{
		return err
	}
	err = os.WriteFile(s.FileName, file1, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (s *UserStorage) Update(id int, user models.User) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	file, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	var users []models.User

	err = json.Unmarshal(file, &users)
	if err != nil {
		return err
	}
	for i, val := range users {
		if val.ID == id {
			users[i] = user
		}
	}
	fileData, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.FileName, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}
