package services

import "com.azimao/models"

func GetUser() string {
	user:=new(models.UserModel)
	user.Uname = "aziamo"
	return user.ToString()
}