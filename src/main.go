package main

import (
	"fmt"
	"golearn/src/notifier"
	"golearn/src/order"
	"golearn/src/user"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type HelloReq struct {
	Name string            `json: "name"`
	Tags []string          `json: "tags"`
	Meta map[string]string `json: "meta"`
}

func startWebServer() {
	// 创建默认路由引擎，带 Logger 和 Recovery 中间件
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(ctx *gin.Context) {
		// 从查询参数里取 name，比如 /hello?name=Tom
		name := ctx.DefaultQuery("name", "world") // 如果没传 name，就用 "world"

		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello," + name,
		})
	})

	// 2) GET /hello/Tom         —— 路径参数
	//    :name 是一个占位符，对应 ctx.Param("name")
	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name == "" {
			name = "world"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "success",
			"message": "hello, " + name,
			"from":    "path",
		})

	})

	// 3) GET /hello/json        —— json 对象
	//    请求体示例：
	//    {
	//      "name": "Tom",
	//      "tags": ["student", "vip"],
	//      "meta": {"city": "Tokyo", "level": "gold"}
	//    }
	r.POST("/hello/json", func(ctx *gin.Context) {
		var req HelloReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "success",
			"message": "hello, " + req.Name,
			"tags":    req.Tags,
			"meta":    req.Meta,
			"from":    "json",
		})
	})

	// 4) GET /hello/list        —— 纯切片 JSON
	//   请求体示例：["Tom", "Jerry", "Bob"]
	r.POST("/hello/list", func(ctx *gin.Context) {
		var names []string
		if err := ctx.ShouldBindJSON(&names); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "success",
			"count": len(names),
			"names": names,
			"from":  "json",
		})
	})

	// 5) GET /hello/map        —— 纯 map JSON
	//    请求体示例：{"Tom": 18, "Jerry": 20}
	r.POST("/hello/map", func(ctx *gin.Context) {
		var m map[string]string
		if err := ctx.ShouldBindJSON(&m); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": m,
			"from": "json map",
		})
	})

	r.Run(":8090")
}

func main() {

	var name string = "John"
	var age = 30
	const conutry = "China"

	fmt.Println("Hello World")
	fmt.Println("你好: ", name, "来自：", conutry, "年龄：", age)

	sum := add(1, 2)
	fmt.Println("Sum: ", sum)
	printSum(1, 2)

	score := 80
	if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println("偶数: ", i)
		} else {
			fmt.Println("奇数: ", i)
		}
	}

	fmt.Println("Sum: ", resultSum(11111))

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", j, i, j*i)
		}
		fmt.Println()
	}

	arr()
	arrayList()
	mapList()
	worldfindSum()
	getUser()

	// u := User{"alan", 40, "519798002@qq.com"}
	// u.SayHello()
	// u.ChangeEmail("alanliurong@qq.com")
	// fmt.Println(u.Email)

	// order := Order{"1001", 3000.00, false}
	// order.Pay()
	// fmt.Println(order.Paid)

	userTest := user.Users{"刘笑笑", "alanlsakjsk@Email.com", 50}
	userTest.PrintUser()

	userTest.ChangeAmount(100)
	orderTest := order.NewOrder("A1001", 80, &userTest)
	orderTest.PrintOrder()

	orderTest.Pay()
	orderTest.PrintOrder()

	userTest.PrintUser()

	email := notifier.EmailNotifier{}
	sms := notifier.SmsNotifier{}
	ai := notifier.AINotifier{}

	NotifyAll(email, "欢迎使用系统！")
	NotifyAll(sms, "验证码：123456")
	ReportAll(email)
	ReportAll(sms)

	notifiers := []notifier.Notifier{email, sms, ai}

	NotifyAllType(notifiers)

	list, err := notifier.LoadNotifierFromConfig("src/notifier/config.json")

	if err != nil {
		fmt.Println("加载配置失败", err)
	}

	if len(list) == 0 {
		fmt.Println("没有可用的通用类型，请检查配置文件")
		return
	}

	fmt.Println("从配置文件加载成功，开始发送通知...")

	for _, n := range list {
		fmt.Printf("通知类型： %s\n", n.GetType())
		n.Send("系统通知，配置驱动加载成功")
		fmt.Println("---------------------------")
	}

	startWebServer()
}

func add(a int, b int) int {
	return a + b
}

func printSum(a int, b int) {
	fmt.Println("Sum: ", add(a, b))
}

func resultSum(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += i
	}
	return total
}

func arr() {
	var sums [3]int = [3]int{10, 20, 30}
	fmt.Println(sums)
	fmt.Println(sums[0])

	// 遍历
	for i, v := range sums {
		fmt.Println("index=%d value=%d\n", i, v)
	}
}

func arrayList() {

	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println(slice)

	slice = append(slice, 50)
	fmt.Println(slice)

	slice = append(slice, 60)
	fmt.Println(slice)

	slice = append(slice, 70)
	fmt.Println(slice)

	fmt.Println(slice[1:4])

	for i, v := range slice {
		fmt.Printf("index=%d value=%d\n", i, v)
	}
}

func mapList() {

	scores := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	scores["haha"] = 100
	fmt.Println(scores)
	fmt.Println(scores["apple"])

	v, ok := scores["apple"]
	if ok {
		fmt.Println("apple score is ", v)
	} else {
		fmt.Println("apple score is not found")
	}

	map2 := make(map[string]int)

	map1 := make(map[string]string)
	map1["alan"] = "a"
	map1["bbbbb1"] = "b"
	map1["cccccc"] = "c"

	fmt.Println(map1)

	map2["apple"] = 1
	map2["banana"] = 2
	map2["cherry"] = 3
	fmt.Println(map2)

	for k, v := range map2 {
		fmt.Printf("key=%s value=%d\n", k, v)
	}
}

// 词频统计

func worldfindSum() {
	text := "I love you , liu hai xianm I love go and I love code"

	// 分割字符串成单词
	world := strings.Fields(text)

	// 统计单词的数量
	counter := make(map[string]int)
	for _, w := range world {
		counter[w]++
	}

	// 把 map 转成切片方便排序
	type kv struct {
		Key   string
		Value int
	}

	var freq []kv
	for k, v := range counter {
		freq = append(freq, kv{k, v})
	}

	//按出现次数从高到低排序
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].Value > freq[j].Value
	})

	// 从高到低词频排序
	fmt.Println("从高到低词频排序")
	for _, kv := range freq {
		fmt.Printf("%-10s : %d\n", kv.Key, kv.Value)
	}

	fmt.Println("词频统计结果：")
	for k, v := range counter {
		fmt.Printf("%s: %d\n", k, v)
	}
}

type User struct {
	Name  string
	Age   int
	Email string
}

func getUser() {
	u1 := User{"刘荣", 34, "519798002@qq.com"}
	u2 := User{Name: "刘荣", Age: 34, Email: "519798002@qq.com"}

	fmt.Println(u1)
	fmt.Println(u2.Name)
}

func (u User) SayHello() {
	fmt.Printf("你好。我是 %s,今年 %d岁\n", u.Name, u.Age)
}

func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

type Order struct {
	ID     string
	Amount float64
	Paid   bool
}

// 打印订单
func (o Order) printOrder() {
	fmt.Printf("订单号：%s, 金额：%.2f, 支付状态： %v\n", o.ID, o.Amount, o.Paid)
}

// 改变订单状态

func (o *Order) Pay() {
	o.Paid = true
	fmt.Println("支付成功")
}

func NotifyAll(n notifier.Notifier, msg string) {
	n.Send(msg)
}

func ReportAll(n notifier.Notifier) {
	typeName := n.GetType()
	fmt.Printf("通知类型是： %s\n", typeName)
}

func NotifyAllType(list []notifier.Notifier) {
	for _, n := range list {
		fmt.Printf("通知类型是：%s\n", n.GetType())
		n.Send("系统提醒接口多态已经生效!")
		fmt.Println("-----------------------------------")
	}
}
