package main

import "fmt"

func verifLoginPage(theMail string, theMdp string) bool{
	var UserTest User

	result := db.Where("UserMail = ?", theMail).First(&UserTest)
	
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if (UserTest.UserMdp != theMdp){
		return false
	}

	return true
}

func disconnect(theUser **User,verif *bool){
	if (*verif == true){
		*theUser = nil
		fmt.Println("l'utilisateur c'est déconnecté")
		*verif = false
	}
}

func deleteActualUser(oneUser *User, verif *bool){
	if (*verif == true){
		var lastId uint = oneUser.id
		disconnect(&oneUser,verif)
		deleteUser(lastId)
		fmt.Println("l'utilisateur a été supprimer")
	} else {
		fmt.Println("Impossible de supprimer un utilisateur en étant supprimer")
	}
}

func UpdateActualUser(id uint, userName string, firstName string, lastName string, userMail string, userMdp string, userStatus string,verif bool){
	if (verif == true){
		updateUser(id,userName,firstName,lastName,userMail,userMdp,userStatus)
	} else {
		fmt.Println("Impossible de modifié un utilisateur en étant déconnecté")
	}
}

func login(mail string, mdp string, oneUser *User){
	if (verifLoginPage(mail,mfdp)== true){
		readUser(&oneUser)
	} else {
		fmt.Println("Aucun profile n'a été trouver - reverifier l'adresse et le mdp")
	}
}

func ExistOrNot(theMail string) bool{
	var UserTest User
	result := db.Where("UserMail = ?", theMail).First(&UserTest)
	
	if (result.Error == nil){
		fmt.Println("Deja pris")
		return true
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	
	return false
}

func createActalUser(userName string, firstName string, lastName string, userMail string, userMdp string){
	if (ExistOrNot(userMail)==false){
		createUser(userName, firstName, lastName, userMail, userMdp)
	} else {
		fmt.Println("existe deja dans la db")
	}
}


