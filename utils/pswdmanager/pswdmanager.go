package pswdmanager

import (
	"OSlash/constants"
	"OSlash/utils/log"

	"golang.org/x/crypto/bcrypt"
)

func GetPswdHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Error.Println("Get Password Hash Error: ", err)
	}
	return string(hash)
}

func VerifyPassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Error.Println(constants.ERR_MSG_PSWD_DECODE, err)
		return false
	}
	return true
}
