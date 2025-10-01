package models

type User struct{
	ID int64 
	Name string
	Email string 
	PassWord string
	PasswordHash string
	Role string
	CreatedAt int64
	UpdatedAt int64
}