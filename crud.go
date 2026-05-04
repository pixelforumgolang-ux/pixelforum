package main

import (
	"time"
	"fmt"
)

//---------------
//CRUD FUNCTION |
//---------------


//CREATE RUD

func createUser(userName string, userMail string, userMdp string, lastName string, firstName string){
	
	if (userName == "" || userMail == "" || userMdp == "" || lastName == "" || firstName  == "" ){
		return nil
	} else {
		user := User{UserName: userName, 
			FirstName : firstName, 
			LastName : LastName, 
			UserMail = userMail,   	
			UserMdp = userMdp, 	
			UserStatus = "client"}

		result := db.Create(&user)

		if (result.Error != nil) {
			fmt.println("erreur")
		}
	}
}

func createPost(theIdUser uint, postMessage string, postImg string, postDate time.time, list []uint){

	if (PostMessage == ""){
		fmt.println("pas de message")
	} else {
		post := Post{		
				idUser : theIdUser,		
				PostMessage : postMessage,
				PostImg : postImg,
				PostDate : postDate,	
				PostCommentary : list}

		result := db.Create(&post)

		if (result.Error != nil) {
			fmt.println("erreur")
		}
	}

	
}

func createCommentary(theIdUser uint,  ComMessage string, ComImg string, ComDate time.time){
	if (ComMessage == ""){
		fmt.println("pas de message")
	} else {
		com := Commentary{		
			UserComId : theIdUser,		
			Message : comMessage,
			Image  : postImg,
			comDate : comDateDate}

		result := db.Create(&com)

		if (result.Error != nil) {
			fmt.println("erreur")
		} else {
			fmt.println(result.RowsAffected)
		}
	}
}

//C READ UD

func readUser(id uint, theUser User){
	result := db.First(&theUser)
}

func readPost(id uint, thePost Post){
	
	if (id == -1){
		result := db.Find(&thePost)
	} else {
		result := db.First(&thePost)
	}
}

func readCommentary(idList []uint, CommentaryList []Commentary){
	if (idList > 0){
		result := db.Find(&CommentaryList, idList)
	} 
}


//CR UPDATE D

func updateUser(id uint, userName string, userMail string, userMdp string, userStatus string, lastName string, firstName string){
	dataUser := User{UserName: userName, 
			FirstName : firstName, 
			LastName : LastName, 
			UserMail : userMail,	
			UserMdp : userMdp,
			UserStatus : userStatus}
	
	result := db.Model(User{}).Where("id = ?", id).Updates(dataUser)
}

func updatePost(id uint, postCommentary []uint){
	result := db.Model(Post{}).Where("id = ?", id).Updates(Post{PostCommentary : postCommentary})
}


//CRU DELETE

func deleteUser(id uint){
	result := db.Delete(&User{}, id)
}

func deletePost(id uint){
	result := db.Delete(&Post{}, id)
}

func deleteCom(id uint){
	result := db.Delete(&Commentary{}, id)
}