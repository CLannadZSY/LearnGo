package main

import (
	"GinDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 使用 New() 替代 Default()
	router := gin.New()
	//router := gin.Default()

	// 1. 最简单示例
	//r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()

	// 2. 不同的请求方式
	//router := gin.Default()
	//
	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePus", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	//
	//router.Run()

	// 3. Patameters in Path
	//router := gin.Default()
	//
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//
	//router.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//
	////For each matched request Context will hold the route definition
	//router.POST("/user/:name/*action", func(c *gin.Context) {
	//	c.FullPath() == "/user/:name/*action"
	//})
	//
	//router.Run(":8080")

	//4. Querystring parameters
	//router.GET("/welcome", func(c *gin.Context) {
	//		firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname")
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})

	// 5. Multipart/Urlencode Form
	//router.POST("/form_post", func(c *gin.Context) {
	//	message := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick", "anonymous")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": "posted",
	//		"message": message,
	//		"nick": nick,
	//	})
	//})

	// 6. 另一个示例：查询+发布表单
	//router.POST("/post", func(c *gin.Context) {
	//	id := c.Query("id")
	//	page := c.DefaultQuery("page", "0")
	//	name := c.PostForm("name")
	//	message := c.PostForm("message")
	//	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	//})

	// 7. 映射为querystring或postform参数
	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	// Content-Type: application/x-www-form-urlencoded
	// names[first]=thinkerou&names[second]=tianou
	// ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
	//router.POST("/post", func(c *gin.Context) {
	//	ids := c.QueryMap("ids")
	//	names := c.PostFormMap("names")
	//	fmt.Printf("ids: %v; names: %v", ids, names)
	//})

	// 8. 上传文件
	//router.MaxMultipartMemory = 8 << 20
	//router.POST("/upload", func(c *gin.Context) {
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//
	//	c.SaveUploadedFile(file, dst)
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//
	//})

	// 9. 上传多个文件, 详细版本代码
	//router.MaxMultipartMemory = 8 << 20
	//router.Static("/", "./public_upload_file")
	//router.POST("/upload", func(c *gin.Context) {
	//	name := c.PostForm("name")
	//	email := c.PostForm("email")
	//
	//	form, err := c.MultipartForm()
	//	if err != nil {
	//		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	//		return
	//	}
	//
	//	files := form.File["fiels"]
	//	for _, file := range files {
	//		filename := filepath.Base(file.Filename)
	//		if err := c.SaveUploadedFile(file, filename); err != nil {
	//			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	//			return
	//		}
	//	}
	//	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	//
	//})

	// 10. Grouping routes
	//v1 := router.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}

	// 11. Using middleware
	//全局中间件
	//记录程序中间件会将日志写入gin.DefaultWriter，即使您使用GIN_MODE = release进行设置。
	//默认情况下gin.DefaultWriter = os.Stdout
	//router.Use(gin.Logger())

	//恢复中间件可从任何紧急情况中恢复，如果有，则写入500。
	//router.Use(gin.Recovery())

	// 对于每个路由中间件，您可以根据需要添加任意数量。
	//router.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	//authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// 12. Custom Recovery behavior, 自定义恢复行为
	//router.Use(gin.Logger())
	//
	//router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
	//	if err, ok := recovered.(string); ok {
	//		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	//	}
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//}))
	//
	//router.GET("/panic", func(c *gin.Context) {
	//	panic("panic error")
	//})
	//
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello world")
	//})

	// 13. How to write log file 将日志写入文件
	// Disable Console Color, you don't need console color when writing the logs to file.
	// 将日志写入文件, 不需要控制台的颜色
	//gin.DisableConsoleColor()
	//
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//
	//// 既写入控制台, 又写入日志, 使用一下代码
	//// Use the following code if you need to write the logs to file and console at the same time.
	//// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})

	//14. Custom Log Format 自定义日志格式
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// your custom format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//router.Use(gin.Recovery())
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})

	// 15. Model binding and validation 模型绑定和验证
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json model.Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml model.Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form model.Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}
