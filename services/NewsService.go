package services

//
//func GetNews() string {
//	user:=GetUser()
//	return "用户信息"+user
//}
type NewsService struct {
}

func (ns NewsService) Get(id int) string {
	return "新闻内容"
}
