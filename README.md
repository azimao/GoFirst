### 常用类型
#### 切片(Slice)
>可以形容为动态数组
```go
# 创建切片
slice1 := make([]type, len)

#初始化切片
s :=[] int {1,2,3 }
```

### 常用实例
#### 异常
> 实例
```go
func test() (ret int) {
    //获取异常
	defer func() {
		err := recover()
		if err != nil{

		}
	}()
	defer func() {   //defer 必须写在panic前面 , panic defer为后进先出原则
		ret++
		fmt.Println(ret)
	}()
	panic("exp")  //异常抛出就不会往下执行
}
```

#### select
> 类似switch case 只不过用于channel通信（要么send 要么receive）
```go
select {
    case
}
```
#### sync包
> `sync.WaitGroup`
```go
	var wg sync.WaitGroup
	for i:=1; i<=5; i++ {
		go func(index int) {
			defer func() {
				wg.Done()
			}()
			fmt.Println(index, "执行完成")
		}(i)
		wg.Add(1)3
	}
	fmt.Println(runtime.NumGoroutine());
	wg.Wait()
```
> `sync.Mutex`
```go
#互斥锁
	var wg sync.WaitGroup
	var mutex sync.Mutex
	list := make([]int, 0)
	go func() {
		defer wg.Done()
		defer mutex.Unlock()
		mutex.Lock()
		for i:=1; i<=1000000; i++ {
			list = append(list, 1)
		}
		fmt.Println(len(list))
	}()
	go func() {
		defer wg.Done()
		defer mutex.Unlock()
		mutex.Lock()
		for i:=1; i<=1000000; i++ {
			list = append(list, 1)
		}
		fmt.Println(len(list))
	}()
	wg.Add(2)
	wg.Wait()
	fmt.Println(len(list))
```
#### bufio 包
> 自带缓冲区，到达一定量才会触发下一步操作
```go
	/***写**/
	//os.O_APPEND|os.O_TRUNC   代表打开清空 写入追加
	file, _ := os.OpenFile("./test/test", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 666)
	defer file.Close()
	//设置超过3字节就写入
	fw := bufio.NewWriterSize(file, 3)
	for i:=1; i<=10; i++ {
		fw.Write([]byte("azimao\n"))
		time.Sleep(time.Second*2)
	}
	//刷新缓冲区，无论是否达到限制都写入
	//fw.Flush()

	/***读**/
    file, _ := os.OpenFile("./test/test", os.O_RDONLY, 666)
    defer file.Close()
    fw := bufio.NewReader(file)

    for {
        //用换行符作为分割来读
        str, err := fw.ReadString('\n')
        if err != nil {
            // io.EOF 代表读到结尾
            if err == io.EOF {
                break
            }
        }
        fmt.Println(str)
    }

	/***多协程读**/
    file, _ := os.OpenFile("./test/test", os.O_RDONLY, 666)
    defer file.Close()
    fw := bufio.NewReader(file)

    for i:=1; i<=2; i++ {
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
                time.Sleep(time.Second*1)
                fmt.Printf("[协程%d]:%s" ,index,str)
                locker.Unlock()
            }
        }(i)
    }
    wg.Add(2)
    wg.Wait()
    fmt.Println("读取完成")
```
