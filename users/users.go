package users

import (
	"encoding/json"
	"os"
	"time"
)

type User struct {
	Name        string    `json:"name"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	Rol         string    `json:"rol"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
}

func deserializeCollection[T any](data []byte) ([]T, error) {
	var collection []T
	err := json.Unmarshal(data, &collection)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func SaveUsersToFile() error {
	usersCollection := []User{
		{Name: "John", Lastname: "Milligton", Email: "Milligton_software@gmail.com", Rol: "engineer", Permissions: []string{"Read", "Write", "Update", "Delete"}, CreatedAt: time.Now()},
		{Name: "Michael", Lastname: "Smith", Email: "smith_software@gmail.com", Rol: "engineer", Permissions: []string{"Read", "Write", "Update", "Delete"}, CreatedAt: time.Now()},
		{Name: "Emma", Lastname: "Jarrington", Email: "emma_hr@gmail.com", Rol: "engineer", Permissions: []string{"Read", "Write", "Update"}, CreatedAt: time.Now()},
	}
	usersJson, err := json.Marshal(usersCollection)
	if err != nil {
		return err
	}
	os.WriteFile("users.json", []byte(usersJson), 0644)
	return nil
}

func GetUsers() ([]User, error) {

	data, err := os.ReadFile("users.json") // Corregido el nombre del archivo
	if err != nil {
		return nil, err
	}
	deserializedUsers, err := deserializeCollection[User](data)
	var users = make([]User, cap(deserializedUsers))
	users = append(users, deserializedUsers...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
