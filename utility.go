package main

func mdpCheck(){

}

func mdpExist(){
	
}

func Search(){

}

func MakeStructPostPage(id uint, subject PostPage){
	readPost(id,subject.subjectPost)
	readCommentary(subject.subjectPost.PostCommentary, subject.subjectCommentary)
}


func MakeStructAdminPage(everything AdminPage){

	readUserAll(everything.everyUser)
	readPostAll(everything.everyPost)
	readCommentaryAll(everything.everyCommentary)
}

func verifLoginPage(theMail string, theMdp string) bool{
	var mail User
	result := db.Where("UserMail = ?", theMail)First(&mail)

	if (result.Error != nil){
		fmt.Println("information érroné ou manquante")
	}  else {
		var mdp User
		result1 := db.Where("UserMdp = ?", theMdp).First(&mdp)
		if (result1.Error != nil){
			fmt.Println("information érroné ou manquante")
		} else {
			return true
		}
	}
}

func disconnect(){
	theUser = nil
	fmt.Println("l'utilisateur c'est déconnecté")
}