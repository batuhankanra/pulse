package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"pulse/internal/store"
)

type Storage struct {
	Path    string            `json:"path"`
	URLs    map[string]string `json:"url"`
	Headers map[string]string `json:"headers"`
}

func NewStorage() (*Storage, error) {
	path, err := store.EnsureFile("data.json")
	if err != nil {
		return nil, err
	}
	st := &Storage{
		Path:    path,
		URLs:    map[string]string{},
		Headers: map[string]string{},
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		b = []byte("{}")
	}
	_ = json.Unmarshal(b, st)
	if st.URLs == nil {
		st.URLs = map[string]string{}
	}
	if st.Headers == nil {
		st.Headers = map[string]string{}
	}

	return st, nil
}
func (st *Storage) save() error {
	res, _ := json.MarshalIndent(st, "", " ")
	return os.WriteFile(st.Path, res, 0644)
}
func (st *Storage) AddURL(name, newUrl string) error {
	st.URLs[name] = newUrl
	return st.save()
}
func (st *Storage) ListURL() error {
	if len(st.URLs) == 0 {
		fmt.Println("Kayıtlı url yok")
		return nil
	}
	fmt.Println("kayıtlı urller:")
	for k, v := range st.URLs {
		fmt.Println(k, ":", v)
	}
	return nil
}
func (st *Storage) DeleteURL(name string) error {
	delete(st.URLs, name)
	return st.save()
}
func (st *Storage) AddHeader(name, headers string) error {
	st.Headers[name] = headers
	return st.save()
}
func (st *Storage) RemoveHeader(name string) error {
	delete(st.Headers, name)
	return st.save()
}
func (st *Storage) ListHeader() error {
	if len(st.Headers) == 0 {
		fmt.Println("Kayıtlı url yok")
		return nil
	}
	fmt.Println("kayıtlı urller:")
	for k, v := range st.Headers {
		fmt.Println(k, ":", v)
	}
	return nil
}
