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
		fmt.Println("pas de message")
	} else {
		user := User{UserName: userName, 
			FirstName : firstName, 
			LastName : lastName, 
			UserMail : userMail,   	
			UserMdp : userMdp, 	
			UserStatus : "client"}

		result := db.Create(&user)

		if (result.Error != nil) {
			fmt.Println("erreur")
		}
	}
}

func createPost(theIdUser uint, postMessage string, postImg string, postDate time.Time, list []uint){

	if (postMessage == ""){
		fmt.Println("pas de message")
	} else {
		post := Post{		
				idUser : theIdUser,		
				PostMessage : postMessage,
				PostImg : postImg,
				PostDate : postDate,	
				PostCommentary : list}

		result := db.Create(&post)

		if (result.Error != nil) {
			fmt.Println("erreur")
		}
	}

	
}

func createCommentary(theIdUser uint, comMessage string, comImg string, comDate time.Time){
	if (comMessage == ""){
		fmt.Println("pas de message")
	} else {
		com := Commentary{		
			UserComId : theIdUser,		
			Message : comMessage,
			Image  : comImg,
			ComDate : comDate}

		result := db.Create(&com)

		if (result.Error != nil) {
			fmt.Println("erreur")
		} else {
			fmt.Println(result.RowsAffected)
		}
	}
}

//C READ UD

func readUser(id uint, theUser User){
	result := db.First(&theUser, id)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readUserAll(theUser []User){
	result := db.Find(&theUser)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readPost(id uint, thePost Post){
	
		result := db.First(&thePost)
		if ( result.Error != nil){
			fmt.Println("erreur")
		}
}

func readPostAll(thePost Post){
	result := db.Find(&thePost)

		if ( result.Error != nil){
			fmt.Println("erreur")
		}
}

func readCommentary(idList []uint, CommentaryList []Commentary){
	if (len(idList) > 0){
		result := db.First(&CommentaryList, idList)

		if ( result.Error != nil){
			fmt.Println("erreur")
		}
	} 
}

func readCommentaryAll(CommentaryList []Commentary){
	
	result := db.Find(&CommentaryList)

	if ( result.Error != nil){
		fmt.Println("erreur")
	} 
}


//CR UPDATE D

func updateUser(id uint, userName string, userMail string, userMdp string, userStatus string, lastName string, firstName string){
	dataUser := User{UserName: userName, 
			FirstName : firstName, 
			LastName : lastName, 
			UserMail : userMail,	
			UserMdp : userMdp,
			UserStatus : userStatus}
	
	result := db.Model(User{}).Where("id = ?", id).Updates(dataUser)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func updatePost(id uint, postCommentary []uint){
	result := db.Model(Post{}).Where("id = ?", id).Updates(Post{PostCommentary : postCommentary})

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}


//CRU DELETE

func deleteUser(id uint){
	result := db.Delete(&User{}, id)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func deletePost(id uint){
	result := db.Delete(&Post{}, id)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func deleteCom(id uint){
	result := db.Delete(&Commentary{}, id)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}