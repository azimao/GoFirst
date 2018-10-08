package models

import "github.com/pquerna/ffjson/ffjson"

type NewsModle struct {
	NewsId int
	NewsTitle,NewsContent string
}

func (news NewsModle) ToJson() string {
	res,err := ffjson.Marshal(news)
	if err != nil {
		return err.Error()
	} else {
		return string(res)
	}
}