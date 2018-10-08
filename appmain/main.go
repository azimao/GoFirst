package main

import (
	"com.azimao/submodels"
	"fmt"
)

func main()  {
	//var me string = "azimao"
	//var u *string
	//u = &me
	//test(u)
	//fmt.Print(*u,me)

	//res:=services.GetNews()+services.GetUser()
	//res:=services.GetUser()

	//user:=new(models.UserModel)
	//user.SetValue(123, "azimao")
	//fmt.Println(user.ToString())


	//var news NewsModle = NewsModle{123, "cs", "azimao"}
	//fmt.Println(news)
	//var news_1 NewsModle = NewsModle{NewsTitle:"tcA", NewsContent:"tiancaiA"}
	//fmt.Println(news_1)
	//
	//var news_2 *NewsModle
	//news_2 = new(NewsModle)
	//
	//*news_2 = NewsModle{123, "cs", "azimao"}
	//news_2 = &NewsModle{123, "cs", "azimao"}
	//fmt.Println(news_2)
	//fmt.Print(res)

	//news := NewsModle{123, "cs", "azimao"}
	//data := NewsModle{3010, "这是测试标题", "tiancaiazimao"}
	//fmt.Println(data.ToJson())

	var sn submodels.SportsNews
	sn.NewsId = 123
	sn.NewsTitle = "abc"
	sn.Tags = []string{"az", "tc", "A"};
	fmt.Println(sn.ToJson())


	//var arr []string = []string{"a", "b", "c"}
	//var arr_1 [3]string = [3]string{0:"a",2:"c"}
	//fmt.Println(arr)
	//fmt.Println(arr_1)
	//fmt.Println(len(arr))


}

func test(p *string)  {

	*p = "tcA"
}