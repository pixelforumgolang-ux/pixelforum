package main

import (
	"time"
)

//---------------
//CRUD FUNCTION |
//---------------


//CREATE RUD

func createUser(&User, userName string, userMail string, userMdp string, lastName string, firstName string){
	
	if (userName == "" || userMail == "" || userMdp == "" || lastName == "" || firstName  == "" ){
		return "no complete information"
	} else {
		return db.Create(&User, userName, userMail, userMdp, "client", lastName, firstName)
	}
	
}

func createPost(&Post, idUser uint, User User, PostMessage string, PostImg string, PostDate time.time){

	if (PostMessage == ""){
		return "no complete information"
	} else {
		if ( PostImg == nil) {
			PostImg = ""
		}
		return db.Create(&Post, idUser, User, PostMessage, PostImg, PostDate, [])
	}
}

func createCommentary(&Commentary, idUser uint, User User,  ComMessage string, ComImg string, ComDate time.time){
	if (ComMessage == ""){
		return "no complete information"
	} else {
		if ( ComImg == nil) {
			ComImg = ""
		}
		return db.Create(&Commentary, idUser, User, ComMessage, ComImg, ComDate)
	}
}

//C READ UD

func readUser(id uint, &User){
	return db.Find(&User, "id = ?", id)
}

func readPost(id uint, &Post){
	return db.Find(&Post, "id = ?", id)
}

func readCommentary(idList []uint, &Commentary){
	if (idList > 0){
		return db.Find(&Commentary, "Commentary.Commentary_id IN ?", idList)
	} else {
		return "no commentary"
	}
	
}

//CR UPDATE D

func updateUser(id uint, userName string, userMail string, userMdp string, userStatus string, lastName string, firstName string){
	return db.Model(&User).Update(id, userName, userMail, userMdp, userStatus, lastName, firstName)
}

func updatePost(id uint, idUser uint, User User, PostMessage string, PostImg string, PostDate time.time, PostCommentary []uint){
	return db.Model(&Post).Update(id, idUser, User, PostMessage, PostImg, PostDate, PostCommentary)
}


//CRU DELETE

func deleteUser(id uint){
	return db.Where("id = ?", id).Delete(&User)
}

func deletePost(id uint){
	return db.Where("idPost = ?", id).Delete(&Post)
}

func deleteCom(id uint){
	return db.Where("idCom = ?", id).Delete(&Com)
}