package models

// User - модель пользователя
type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Type     string `json:"type,omitempty"` // {manager, user}
	TypeId   int    `json:"type_id"`
}

// UserType - модель роли пользователя
type UserType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Good - модель товара
type Good struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

// CartsForManager - модель для показа списка корзин всех пользоватлей
type CartsForManager struct {
	UserId int `json:"user_id"`
	Cart []Good `json:"cart"`
}

// // Good - модель корзины
type ReservedGood struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	GoodId int `json:"good_id"`
	Count int `json:"count"`
}