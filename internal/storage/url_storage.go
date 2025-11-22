package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"pulse/internal/store"
)

type URLStore struct {
	Path string
	Data map[string]string
}

func NewURL() (*URLStore, error) {
	path, err := store.EnsureFile("url.json")
	if err != nil {
		return nil, err
	}
	store := &URLStore{
		Path: path,
		Data: map[string]string{},
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &store)
	return store, nil
}

func (s *URLStore) Save() error {
	b, _ := json.MarshalIndent(s.Data, "", " ")
	return os.WriteFile(s.Path, b, 0644)
}

func (s *URLStore) Add(name, url string) error {
	s.Data[name] = url
	return s.Save()
}

func (s *URLStore) Remove(name string) error {
	delete(s.Data, name)
	return s.Save()
}
func (s *URLStore) Get(name string) (string, bool) {
	v, ok := s.Data[name]
	return v, ok
}
func (s *URLStore) List() error {
	if len(s.Data) == 0 {
		fmt.Println("Kayıtlı url yok!")
		return nil
	}

	fmt.Println("\n Kayıtlı Url'ler:")
	for k, v := range s.Data {
		fmt.Println("-", k, ":", v)
	}
	return nil
}
