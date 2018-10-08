package models

type UserModel struct {
	Uid int
	Uname string
}


func (u *UserModel) SetValue(id int, name string) string {
	u.Uid = id
	u.Uname = name
	return "";
}

func (u UserModel) ToString() string {
	return "用户名是："+u.Uname
}

