package password

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\"<>,.?/~"
)

func GeneratePassword(length int, useUppercase, useDigits, useSpecialChars bool) string {
	var validChars string
	validChars += lowercaseLetters
	if useUppercase {
		validChars += uppercaseLetters
	}
	if useDigits {
		validChars += digits
	}
	if useSpecialChars {
		validChars += specialChars
	}
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(password)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erreur lors du hachage du mot de passe: %v", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// ValidatePassword valide un mot de passe selon les critères de sécurité
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Le mot de passe doit contenir au moins 8 caractères"
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, "Le mot de passe doit contenir au moins une lettre minuscule"
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, "Le mot de passe doit contenir au moins une lettre majuscule"
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, "Le mot de passe doit contenir au moins un chiffre"
	}
	return true, ""
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidateUsername(username string) (bool, string) {
	if len(username) < 3 {
		return false, "Le nom d'utilisateur doit contenir au moins 3 caractères"
	}
	if len(username) > 20 {
		return false, "Le nom d'utilisateur ne peut pas dépasser 20 caractères"
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(username) {
		return false, "Le nom d'utilisateur ne peut contenir que des lettres, chiffres, tirets et underscores"
	}
	return true, ""
}