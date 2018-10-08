package services

func GetNews() string {
	user:=GetUser()
	return "用户信息"+user
}
