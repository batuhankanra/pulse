package models

type Config struct {
	URLs    map[string]string            `json:"urls"`
	Headers map[string]map[string]string `json:"headers"`
	Body    map[string]map[string]string `json:"body"`
}
