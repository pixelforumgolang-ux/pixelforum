package main

import (
	"time"
	"fmt"
)

//---------------
//CRUD FUNCTION |
//---------------


//CREATE RUD

func createUser(userName string, firstName string, lastName string, userMail string, userMdp string){
	
	if (userName == "" || userMail == "" || userMdp == "" || lastName == "" || firstName  == "" ){

		fmt.Println("pas de message")
	
	} else {

		user := User{
			UserName: userName, 
			FirstName : firstName, 
			LastName : lastName, 
			UserMail : userMail,   	
			UserMdp : userMdp, 	
			UserStatus : "client"}

		result := db.Create(&user)

		if (result.Error != nil) {
			fmt.Println("erreur lors de la création du profile")
		}

		fmt.Println("Vous pouvez maintenant vous conncter")
	}
}

func createPost(userID uint, subject string, message string, image string, time time.Time){

	if (comMessage == "" || subject == "" ){

		fmt.Println("pas de message")

	} else {

		time = Date.now()
		DataCom := Post{		
			UserID		: userID
			Subject 	: subject,
			Message 	: message,	
			Image 		: image,	
			Date		: time
		}

		result := db.Create(&DataCom)

		if (result.Error != nil) {
			fmt.Println("erreur")
		}

	}
}

func createCommentary(userID uint, postID uint, message string, image string){

	if (message == ""){

		fmt.Println("pas de message")

	} else {

		time = Date.now()
		DataCom := Commentary{		
			UserID 		: userID,		
			PostID 		: postID,
			Message 	: message,
			Image  		: image,
			Date 		: time
		}

		result := db.Create(&DataCom)

		if (result.Error != nil) {
			fmt.Println("erreur")
		}

	}
}

func createSession(userID uint){
	
	session := Session{
		UserID : userID
		token : true
	}

	result := db.Create(&session)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}

}

func createlikes(comID uint){
	
	like := Likes{
		ComID : comID
		NbLike : 0
	}

	result := db.Create(&like)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}

}

//C READ UD

func readUser(mdp string, theUser *User){
	result := db.Preload("Post").Preload("Commentary").Where("UserMdp = ?",mdp).First(theUser)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readUserAll(users *[]User){
	result := db.Preload("Post").Preload("Commentary").Find(users)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readPost(id uint, thePost *Post){
	result := db.Preload("Commentary").First(thePost, id)
	
	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readPostAll(thePost *[]Post){
	result := db.Preload("Commentary").Find(thePost)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readCommentary(idList []uint, CommentaryList *[]Commentary){
	result := db.Preload("Post").Preload("User").Find(CommentaryList, "PostID IN ?", idList)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readCommentaryAll(CommentaryList *[]Commentary){
	result := db.Preload("Post").Preload("User").Find(CommentaryList)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readLikes(likes *Likes, id uint){
	result := db.First(likes, id)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

readAllLikes(likes *[]Likes){
	result := db.Find(likes)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func readSession(session *Session, id uint){
	result := db.Where("ComID = ?", id).First(Session)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

readAllSession(session *[]Sessions){
	result := db.Find(session)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}



//CR UPDATE D

func updateUser(id uint, userName string, firstName string, lastName string, userMail string, userMdp string, userStatus string){
	dataUser := User{
			UserName: userName, 
			FirstName : firstName, 
			LastName : lastName, 
			UserMail : userMail,	
			UserMdp : userMdp,
			UserStatus : userStatus
		}
	
	result := db.Model(User{}).Where("id = ?", id).Updates(dataUser)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func updateLike(comID uint, like int){
	
	Likes := Likes{
		NbLike : like
	}

	result := db.Model(Likes{}).Where("UserID = ?", comID uint).Updates(&session)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}

}

func updateSessionON(userID uint){
	
	session := Session{
		token : true
	}

	result := db.Model(Session{}).Where("UserID = ?", userID uint).Updates(&session)

	if ( result.Error != nil){
		fmt.Println("erreur")
	}

}

func updateSessionOFF(userID uint){
	
	session := Session{
		token : false
	}

	result := db.Model(Session{}).Where("UserID = ?", userID uint).Updates(&session)

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

func deleteLikes(id uint){
	result := db.Where("ComID = ?", id).Delete(&Likes{})

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}

func deleteSession(id uint){
	result := db.Where("UserID = ?", id).Delete(&Session{})

	if ( result.Error != nil){
		fmt.Println("erreur")
	}
}