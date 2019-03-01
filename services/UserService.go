package services

//import "com.azimao/models"
//
//func GetUser() string {
//	user:=new(models.UserModel)
//	user.Uname = "aziamo"
//	return user.ToString()
//}
type UserService struct {
}

func (us UserService) Get(id int) string {
	return "单个用户信息"
}
