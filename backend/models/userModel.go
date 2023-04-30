package models

type User struct {
	Name string `json:"name"`;
	Username string `json:"username"`;
	Email string `json:"email"`;
	Password string `json:"password"`;
}
type Shirt struct {
	ID string  `json:"id"`;
	Title string `json:"title"`;
	Cost int `json:"cost"`;
	Quantity int `json:"quantity"`;
}
type Pants struct {
	ID string  `json:"id"`;
	Title string `json:"title"`;
	Cost int `json:"cost"`;
	Quantity int `json:"quantity"`;
}
type Item struct { 
	Title string `json:"title"`; 
	Sold int `json:"sold"`; 
}