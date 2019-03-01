package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"sync"
	"time"
)

//func sum(min,max int, c chan int)  {
//	result := 0
//	for i:=min; i<=max; i++  {
//		result = result+i
//	}
//	//result := (1+max)*max/2
//	//fmt.Println(result)
//	c <- result
//}

func test() (ret int) {
	defer func() { //defer 必须写在panic前面 , panic defer为后进先出原则
		ret++
		fmt.Println(ret)
	}()
	panic("exp") //异常抛出就不会往下执行
}

func main() {
	var wg sync.WaitGroup
	var locker sync.Mutex

	file, _ := os.OpenFile("./test/test", os.O_RDONLY, 666)
	defer file.Close()
	fw := bufio.NewReader(file)

	for i := 1; i <= 2; i++ {
		go func(index int) {
			defer wg.Done()
			for {
				locker.Lock()
				//用换行符作为分割来读
				str, err := fw.ReadString('\n')
				if err != nil {
					// io.EOF 代表读到结尾
					if err == io.EOF {
						locker.Unlock()
						break
					}
					fmt.Println(err)
				}
				time.Sleep(time.Second * 1)
				fmt.Printf("[协程%d]:%s", index, str)
				locker.Unlock()
			}
		}(i)
	}
	wg.Add(2)
	wg.Wait()
	fmt.Println("读取完成")

	//刷新缓冲区，无论是否达到限制都写入
	//fw.Flush()
	//c:=time.After(time.Second*3)
	//fmt.Println("不堵塞")
	//fmt.Println(<-c)
	//var wg sync.WaitGroup
	//var mutex sync.Mutex
	//list := make([]int, 0)
	//go func() {
	//	defer wg.Done()
	//	defer mutex.Unlock()
	//	mutex.Lock()
	//	for i:=1; i<=1000000; i++ {
	//		list = append(list, 1)
	//	}
	//	fmt.Println(len(list))
	//}()
	//go func() {
	//	defer wg.Done()
	//	defer mutex.Unlock()
	//	mutex.Lock()
	//	for i:=1; i<=1000000; i++ {
	//		list = append(list, 1)
	//	}
	//	fmt.Println(len(list))
	//}()
	//wg.Add(2)
	//wg.Wait()
	//fmt.Println(len(list))

	//url := "https://news.cnblogs.com/n/page/%d/"
	//
	//var c = make(chan map[int][]byte)
	//for i:=1; i<=10;i++  {
	//	go func(index int) {
	//		url := fmt.Sprintf(url, index)
	//		res, _ := http.Get(url)
	//		cnt, _ := ioutil.ReadAll(res.Body)
	//
	//		c <- map[int][]byte{index:cnt}
	//		//ioutil.WriteFile(fmt.Sprintf("./files/%d", i), cnt, 666)
	//		//if index==3 {
	//		//	close(c)
	//		//}
	//	}(i)
	//}
	////for getcnt := range c {
	////	for k,v := range getcnt {
	////	ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
	////	}
	////}
	//result := map[int][]byte{}
	//myloop:for {
	//	select {
	//		case result = <-c:
	//			for k,v := range result {
	//				ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
	//			}
	//		case <-time.After(time.Second*3):
	//			break myloop
	//
	//	}
	//}
	//return
	//-------------------------------------------------------------------------

	//res, _ := http.Get(url)
	//cnt, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(cnt))
	//	return
	//
	//
	//
	//
	//	return
	//	users := strings.Split("azimao, azi, tiancai, tiancaiA, tcA", ",")
	//	ages := strings.Split("12, 15, 22, 33, 44", ",")
	//
	//	//allList := make([]string, 0)
	//	//allList = append(allList, users...)
	//	//allList = append(allList, ages...)
	///*	for i, v := range users {
	//		allList = append(allList, v)
	//		allList = append(allList, ages[i])
	//	}*/
	//	//c1 := make(chan string, len(users))
	//	//c2 := make(chan string, len(ages))
	//	c1, c2 := make(chan bool), make(chan bool)
	//
	//	res := make([]string, 0)
	//	//匿名函数
	//	go func() {
	//		for _,v := range users {
	//			<-c1
	//			res = append(res, v)
	//			c2<-true
	//		}
	//	}()
	//	go func() {
	//		for _,v := range ages {
	//			<-c2
	//			res = append(res, v)
	//			c1<-true
	//		}
	//	}()
	//	c1<-true
	//	//for v := range res{
	//	//	fmt.Println(v)
	//	//}
	//	//
	//	//
	//	fmt.Println(res)
	//	return
	//	num := 1000000000;
	//	start := time.Now()
	//	c := make(chan int, 2)
	//	go sum(1, num/2, c)
	//	go sum(1+num/2, num, c)
	//	//sum(1, num, c)
	//
	//	ret1, ret2 := <- c, <-c
	//	end := time.Now()
	//	fmt.Println(end.Sub(start))
	//	fmt.Println(ret1+ret2)
	//	return
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

	//var sn submodels.SportsNews
	//sn.NewsId = 123
	//sn.NewsTitle = "abc"
	//sn.Tags = []string{"az", "tc", "A"};
	//fmt.Println(sn.ToJson())

	//var arr []string = []string{"a", "b", "c"}
	//var arr_1 [3]string = [3]string{0:"a",2:"c"}
	//fmt.Println(arr)
	//fmt.Println(arr_1)
	//fmt.Println(len(arr))

	//var service services.IService = services.NewsService{}
	//var service services.IService = new(services.UserService)
	//var service services.IService = services.NewServiceFactory().Create("user")
	//fmt.Println(service.Get(123))
	//s := []int{1,2,3}
	//fmt.Println(len(s), cap(s))

	//arr := [5]int{1,2,3,4,5}
	//// <=0   >5
	//s := arr[:2]
	//s = append(s, 4,5,6,7)
	//fmt.Println(s)
	//var i interface{} = "123"
	//var list = make([]interface{}, 2)
	//list[0] = 1
	//list[1] = "abc"
	//fmt.Println(list)

	//db, err := sql.Open("mysql", "root:320199@tcp(localhost:3306)/ceshi?charset=utf8mb4")
	//if err != nil{
	//	fmt.Println("连接错误:"+err.Error())
	//	return
	//}
	//
	//rows,err := db.Query("select * from admins limit 10")
	//if err != nil{
	//	fmt.Println("查询错误:"+err.Error())
	//	return
	//}
	//columns,_ :=rows.Columns()
	////userModels := []models.UserModel{}
	//allRows := make([]interface{}, 0)
	//fieldMap := make(map[string]interface{}, 0)
	//for rows.Next(){
	//	//tmp := models.UserModel{}
	//	//var uid int
	//	//var uname string
	//	oneRow := make([]interface{}, len(columns))
	//	scanRow := make([]interface{}, len(columns))
	//	for i,_ := range oneRow{
	//		scanRow[i] = &oneRow[i]
	//	}
	//	//rows.Scan(&oneRow[0], &oneRow[1])
	//	rows.Scan(scanRow...)
	//	for i,val:=range oneRow{
	//		v,ok := val.([]byte)
	//		if (ok) {
	//			fieldMap[columns[i]] = string(v)
	//		}
	//	}
	//	allRows = append(allRows, fieldMap)
	//	//oneRow[0] = uid
	//	//oneRow[1] = uname
	//	//allRows = append(allRows, []interface{}{uid, uname})
	//	//userModels = append(userModels, tmp)
	//	//fmt.Println(userModels)
	//}
	////m:=map[string]string{"username":"azimao"}
	//fmt.Println(allRows)
	//return
	//
	//m := make(map[string]string)
	//m["username"] = "azimao"
	//m["age"] = "19"
	//delete(m, "age")
	//fmt.Println(m)

}
