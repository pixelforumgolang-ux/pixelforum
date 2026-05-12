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
	UserStatus     	string  	`gorm:"column:UserStatus"`

	Post			[]Post			`gorm:"foreignKey:UserID"` //permet de récupéré facilement via preload(Commentary) tous les Commentary lié a un user
	Commentary		[]Commentary	`gorm:"foreignKey:UserID"` //permet de récupéré facilement via preload(Post) tous les Post lié a un user
}

type Post struct { 
	//Post information 
	PostID 			uint    	`gorm:"primaryKey;column:PostID"`
	UserID 			uint		`gorm:"column:UserID"` 
	User   			User		`gorm:"foreignKey:UserID"`
	Subject 		string		`gorm:"column:Subject"`
	Message 		string		`gorm:"column:Message"`
	Image 			string		`gorm:"column:Image"`
	Date 			time.Time	`gorm:"column:Date"`

	Commentary 		[]Commentary	`gorm:"foreignKey:PostID"`  //permet de récupéré facilement via preload(Commentary) tous les Commentary lié a un post
}

type Commentary struct {
	//Commentary information
	ComID 			uint    	`gorm:"primaryKey;column:ComID"`
	MessageCom 		string		`gorm:"column:MessageCom"`
	ImageCom 		string		`gorm:"column:ImageCom"`
	DateCom 		time.Time	`gorm:"column:DateCom"`
	PostID 			uint		`gorm:"column:PostID"` //-----------------> association belong to pour récupéré facilement les post
	Post			Post		`gorm:"foreignKey:PostID"`
	UserID 			uint		`gorm:"column:UserID"` //-----------------> association belong to pour récupéré facilement les user
	User			User		`gorm:"foreignKey:UserID"`
}

type Likes struct {
	//like by Commentary / association "belong to"
	ComID 			uint		`gorm:"column:ComID"` //-----------------> association belong to pour récupéré facilement les Commentary
	Commentary		Commentary	`gorm:"foreignKey:ComID"`
	NbLike 			int			`gorm:"column:nbLike"`
}

type Session struct {
	//Session information / association "belong to"
	UserID 			uint		`gorm:"column:UserID"`
	User   			User		`gorm:"foreignKey:UserID"`
	Token 			int 		`gorm:"column:Token"`
}

