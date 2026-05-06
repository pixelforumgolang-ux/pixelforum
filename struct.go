package main

import (
	"time"
)

type User struct {
	// user information
	ID 				uint    	`gorm:"primaryKey;column:id"`
	UserName      	string  	`gorm:"column:UserName"`
	FirstName      	string  	`gorm:"column:firstName"`
	LastName      	string  	`gorm:"column:lastName"`
	UserMail      	string  	`gorm:"column:UserMail"`
	UserMdp      	string  	`gorm:"column:UserMdp"`
	UserStatus      string  	`gorm:"column:UserStatus"`
}

type Post struct {
	PostId 			uint    	`gorm:"primaryKey;column:PostId"`

	//user information
	IdUser 			uint  		`gorm:"column:idUser"`
	User 			User 		`gorm:"foreignKey:idUser"`

	//post information
	PostMessage 	string  	`gorm:"column:PostMessage"`
	PostImg     	string  	`gorm:"column:PostImg"`
	PostDate     	time.Time  	`gorm:"column:PostDate"`
}


type Commentary struct {
	Commentary_id 	uint    	`gorm:"primaryKey;column:id"`

	//user information
	UserComId      	uint  		`gorm:"column:UserComId"`
	User 			User 		`gorm:"foreignKey:UserComId"`

	//Post information
	PostComId		uint		`gorm:"column:PostComId"`
	Post 			Post		`gorm:"foreignKey:PostComId"`

	//Commentary information
	Message     	string  	`gorm:"column:ComMessage"`
	Image     		string  	`gorm:"column:image"`
	ComDate     	time.Time  	`gorm:"column:ComDate"`
}
