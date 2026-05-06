package main

import (
	"time"
)

type User struct {
	// user information
	ID 				uint    	`gorm:"primaryKey;column:ID"`

	UserName      	string  	`gorm:"column:UserName"`
	FirstName      	string  	`gorm:"column:firstName"`
	LastName      	string  	`gorm:"column:lastName"`
	UserMail      	string  	`gorm:"column:UserMail"`
	UserMdp      	string  	`gorm:"column:UserMdp"`
	UserStatus      string  	`gorm:"column:UserStatus"`
}

type Post struct {
	PostID 			uint    	`gorm:"primaryKey;column:PostID"`

	//user information
	IdUser 			uint  		`gorm:"column:idUser"`
	User 			User 		`gorm:"foreignKey:idUser"`

	//post information
	PostMessage 	string  	`gorm:"column:PostMessage"`
	PostImg     	string  	`gorm:"column:PostImg"`
	PostDate     	time.Time  	`gorm:"column:PostDate"`

	ComID 			*uint		`gorm:"column:ComID"`
  	CommentaryList  []Commentary `gorm:"foreignkey:ComID"`
}


type Commentary struct {
	ComID 	uint    	`gorm:"primaryKey;column:ComID"`

	//user information
	UserComId      	uint  		`gorm:"column:UserComId"`
	User 			User 		`gorm:"foreignKey:UserComId"`

	//Commentary information
	Message     	string  	`gorm:"column:ComMessage"`
	Image     		string  	`gorm:"column:image"`
	ComDate     	time.Time  	`gorm:"column:ComDate"`
}
