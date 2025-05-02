package config

type Config struct {
	Name   string  `json:"name,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value any    `json:"value,omitempty"`
}

type Group struct {
	Name string     `json:"name,omitempty"`
	Vars []KeyValue `json:"vars,omitempty"`
}
