package main

import "time"

type User struct {
	id 				uint    	`gorm:"primaryKey;column:id"`
	UserName      	string  	`gorm:"column:UserName"`
	UserMail      	string  	`gorm:"column:UserMail"`
	UserMdp      	string  	`gorm:"column:UserMdp"`
	UserStatus      string  	`gorm:"column:UserStatus"`
}

type Post struct {
	PostId 			uint    	`gorm:"primaryKey;column:PostId"`
	idUser 			uint  		`gorm:"column:idUser"`
	User 			User 		`gorm:"foreignKey:idUser"`
	PostMessage 	string  	`gorm:"column:PostMessage"`
	PostImg     	string  	`gorm:"column:PostImg"`
	PostDate     	time.Time  	`gorm:"column:PostDate"`
	PostCommentary  []uint  	`gorm:"column:PostCommentary"`
}

type Commentary struct {
	Commentary_id 	uint    	`gorm:"primaryKey;column:id"`
	UserComId      	uint  		`gorm:"column:UserComId"`
	User 			User 		`gorm:"foreignKey:UserComId"`
	Message     	string  	`gorm:"column:ComMessage"`
	Image     		string  	`gorm:"column:image"`
}