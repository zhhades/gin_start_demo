package model

type Config struct {
	Port int
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
