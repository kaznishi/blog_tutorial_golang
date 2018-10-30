package data_model

type User struct {
	ID int
	Name string
	Password string
	Salt string
	IsActive int
}
