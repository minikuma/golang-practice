package model

type Person struct {
	ID   string
	Name string `json:"name"`
	Age  int    `json:"age"`
}
