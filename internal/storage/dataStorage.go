package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pulse/internal/store"
)

type Storage struct {
	Path    string                       `json:"path"`
	URLs    map[string]string            `json:"url"`
	Headers map[string]map[string]string `json:"headers"`
	Bodys   map[string]map[string]string `json:"body"`
}

func NewStorage() (*Storage, error) {
	path, err := store.EnsureFile("data.json")
	if err != nil {
		return nil, err
	}
	st := &Storage{
		Path:    path,
		URLs:    map[string]string{},
		Headers: map[string]map[string]string{},
		Bodys:   map[string]map[string]string{},
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
		st.Headers = map[string]map[string]string{}
	}
	if st.Bodys == nil {
		st.Bodys = map[string]map[string]string{}
	}

	return st, nil
}
func (st *Storage) save() error {
	res, _ := json.MarshalIndent(st, "", "  ")
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
func (st *Storage) AddHeader(name, key, value string) error {
	if st.Headers[name] == nil {
		st.Headers[name] = map[string]string{}
	}
	st.Headers[name][key] = value
	return st.save()

}
func (st *Storage) RemoveHeader(name, key string) error {
	if st.Headers[name] == nil {
		return nil
	}
	delete(st.Headers[name], key)
	return st.save()
}
func (st *Storage) ListHeader() error {
	if len(st.Headers) == 0 {
		fmt.Println("Kayıtlı url yok")
		return nil
	}
	fmt.Println("kayıtlı urller:")
	for k, v := range st.Headers {
		for vk, vv := range v {
			fmt.Println(k, ":  ", vk, ":", vv)
		}
	}
	return nil
}
func (st *Storage) AddBody(name, key, value string) error {
	if st.Bodys[name] == nil {
		st.Bodys[name] = map[string]string{}
	}
	st.Bodys[name][key] = value
	return st.save()

}
func (st *Storage) RemoveBody(name, key string) error {
	if st.Bodys[name] == nil {
		return nil
	}
	if key == "" {
		delete(st.Bodys, name)
	} else {
		delete(st.Bodys[name], key)
	}
	return st.save()
}
func (st *Storage) ListBody() error {
	if len(st.Bodys) == 0 {
		fmt.Println("Kayıtlı body yok")
		return nil
	}
	fmt.Println("kayıtlı Body:")
	for k, v := range st.Bodys {
		for vk, vv := range v {
			fmt.Println(k, ":  ", vk, ":", vv)
		}
	}
	return nil
}

func SendRequest(url, method string, headers map[string]string, body interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		if string(jsonBytes) != "null" {
			bodyReader = bytes.NewReader(jsonBytes)
		}
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	fmt.Println("------------------- REQUEST -------------------")
	fmt.Printf("%s %s HTTP/1.1\n", method, req.URL.Path)

	if body != nil {
		jsonBytes, _ := json.MarshalIndent(body, "", "  ")
		fmt.Printf("Giden Body:\n%s\n", string(jsonBytes))
	}
	fmt.Println("-----------------------------------------------")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("------------- RESPONSE -------------")
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	return nil
}
