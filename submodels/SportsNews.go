package submodels

import (
	"com.azimao/models"
	"github.com/pquerna/ffjson/ffjson"
)

type SportsNews struct {
	Tags []string
	models.NewsModle
}

func (sn SportsNews) ToJson() string {
	res,err := ffjson.Marshal(sn)
	if err != nil {
		return err.Error()
	} else {
		return string(res)
	}
}
