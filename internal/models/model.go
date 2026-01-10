package models

type Config struct {
	URLs    map[string]string `json:"urls"`
	Headers map[string]string `json:"headers"`
}
