package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

func test01() {
	//定义一个默认的引擎
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.Run(":9000")
}
func test02() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/fqb", func(c *gin.Context) {
		var msg struct {
			Age     int
			Name    string
			Message string
		}
		msg.Name = "Gentle"
		msg.Age = 23
		msg.Message = "阳光"

		//map[string]interface{} 类型的数据
		// c.HTML(http.StatusOK, "fqb.tmpl", gin.H{
		// 	"Age":     18,
		// 	"Name":    "Gentle",
		// 	"Message": "帅气的",
		// })
		//传结构体
		c.HTML(200, "fqb.tmpl", msg)
	})
	r.Run(":9000")
}

func test03() {
	//定义一个默认模版
	router := gin.Default()
	//路由1
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	//路由2
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK, message)
	})
	//路由3
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(200, "%t", b)
	})

	//路由4
	router.GET("user/groups", func(c *gin.Context) {
		c.String(200, "The avaliable groups are [...]")
	})

	router.Run(":9000")
}

func test04() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(200, "Hello %s %s", firstname, lastname)
	})
	r.Run(":9000")
}

func test05() {
	//ps: http://127.0.0.1:9090/user/search?address=123456&username=gentle
	router := gin.Default()
	router.GET("/user/search", func(context *gin.Context) {
		username := context.DefaultQuery("username", "小王子") //取不到，就给默认值
		address := context.Query("address")                 //取不到给空字符
		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	router.Run(":9090")
}

func test06() {
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"username": username,
			"address":  address,
		})
	})

	r.POST("/json", func(c *gin.Context) {
		b, err := c.GetRawData() //从c.Request.Body读取请求数据 raw接收数据
		if err != nil {
			panic(err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(b, &m) //[]byte ->json格式
		c.JSON(200, m)
	})
	r.Run(":9090")
}

func test07() {
	r := gin.Default()
	r.POST("/json", func(c *gin.Context) {
		b, err := c.GetRawData() //从c.Request.Body读取请求数据
		if err != nil {
			panic(err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(b, &m)
		c.JSON(200, m)
	})
	r.Run(":9090")
}

//获取path参数
func test08() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")

		c.JSON(200, gin.H{
			"message":  "OK",
			"username": username,
			"address":  address,
		})

	})
	r.Run(":9090")
}

//参数绑定
type Login struct {
	User     string `form:"user"json:"user"binding:"required"`
	Password string `form:"password"json:"password"binding:"required"`
}

func test09() {
	r := gin.Default()

	//绑定JSON
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		//Body x-www-form-urlencoded
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user": login.User,
				"pwd":  login.Password,
			})
		} else {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		}
	})

	//绑定form表单示例
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info :%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

	//绑定QureyString
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(200, gin.H{
				"user": login.User,
				"pwd":  login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.Run(":8090")
}

//单个文件上传
func test10() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprint("./%s", file.Filename)
		c.SaveUploadedFile(file, dst) //上传文件到指定的目录
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%s'uploaded!", file.Filename),
		})
	})
	r.Run(":9090")
}

//上传多个文件
func test11() {
	r := gin.Default()
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"] //file is key

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	r.Run(":9090")
}

//HTTP重定向
func test12() {
	r := gin.Default()
	r.GET("/go", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	r.Run(":8090")
}

//路由重定向
func test13() {
	r := gin.Default()
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bbbb",
		})
	})
	r.Run(":8090")
}

//特殊的路由
func test14() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test-GET",
		})

	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test-POST",
		})
	})

	r.Any("/any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ANY",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	r.Run(":8090")
}

//Gin中间件Hook函数
//自定义统计耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		//通过set在上下文中设置值
		c.Set("name", "gentle")
		//调用该请求的剩余处理程序
		c.Next()
		//不调用该请求的剩余处理程序
		//c.Abort()
		cost := time.Since(start)
		log.Println(cost)
	}
}

func StatCost2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("age", 32)
	}
}
func test15() {
	//新建一个没有任何路由的中间件
	r := gin.New()
	//注册一个全局中间件
	r.Use(StatCost())
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) //从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "hello goland",
		})
	})
	r.GET("/test2", StatCost2(), func(c *gin.Context) {
		name := c.MustGet("name").(string)
		age := c.MustGet("age").(int)
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run()
}

//运行多个服务
var (
	g errgroup.Group
)

func router1() http.Handler {
	e := gin.New()
	//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":  200,
			"error": "welcome server 01",
		})
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "welcome server 02",
		})
	})
	return e
}

func test16() {
	server01 := &http.Server{
		Addr:         ":9080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr:              ":9090",
		Handler:           router02(),
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	//借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	test16()
}
