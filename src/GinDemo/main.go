package main

import (
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
	//Example for binding JSON ({"user": "manu", "password": "123"})
	//router.POST("/loginJSON", func(c *gin.Context) {
	//	var json model.Login
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if json.User != "manu" || json.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	//router.POST("/loginXML", func(c *gin.Context) {
	//	var xml model.Login
	//	if err := c.ShouldBindXML(&xml); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if xml.User != "manu" || xml.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})
	//
	//// Example for binding a HTML form (user=manu&password=123)
	//router.POST("/loginForm", func(c *gin.Context) {
	//	var form model.Login
	//	// This will infer what binder to use depending on the content-type header.
	//	if err := c.ShouldBind(&form); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if form.User != "manu" || form.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})

	// 16. Custom Validators
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("bookabledate", model.BookableDate)
	//}
	//router.GET("/bookable", model.GetBookable)

	// 17. Only Bind Query String
	//router.Any("/testing", model.StartPage)

	// 18. Bind Query String or Post Data
	//router.GET("/testing", model.StartPagePost)

	//19 Bind Uri
	//router.GET("/:name/:id", func(c *gin.Context) {
	//	var person model.PersonUri
	//	if err := c.ShouldBindUri(&person); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	//})

	// 20. Bind Header
	//router.GET("/", func(c *gin.Context) {
	//	h := model.BindHeader{}
	//	if err := c.ShouldBindHeader(&h); err != nil {
	//		c.JSON(http.StatusOK, err)
	//	}
	//	fmt.Printf("%#v\n", h)
	//	c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	//})

	// 21. XML, JSON, YAML and ProtoBuf rendering
	// gin.H is a shortcut for map[string]interface{}
	//router.GET("/someJSON", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//router.GET("/moreJSON", func(c *gin.Context) {
	//	// You also can use a struct
	//	var msg struct {
	//		Name    string `json:"user"`
	//		Message string
	//		Number  int
	//	}
	//	msg.Name = "Lena"
	//	msg.Message = "hey"
	//	msg.Number = 123
	//	// Note that msg.Name becomes "user" in the JSON
	//	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	//	c.JSON(http.StatusOK, msg)
	//})
	//
	//router.GET("/someXML", func(c *gin.Context) {
	//	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//router.GET("/someYAML", func(c *gin.Context) {
	//	c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//router.GET("/someProtoBuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//	// The specific definition of protobuf is written in the testdata/protoexample file.
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	// Note that data becomes binary data in the response
	//	// Will output protoexample.Test protobuf serialized data
	//	c.ProtoBuf(http.StatusOK, data)
	//})

	// 22. Serving static files 提供静态文件
	//router.Static("/assets", "./assets")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	//23. Serving data from file
	//router.GET("/local/file", func(c *gin.Context) {
	//	c.File("local/file.go")
	//})
	//
	//var fs http.FileSystem = // ...
	//router.GET("/fs/file", func(c *gin.Context) {
	//	c.FileFromFS("fs/file.go", fs)
	//})

	// 24. Serving data from reader
	//router.GET("/someDataFromReader", func(c *gin.Context) {
	//	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	//	if err != nil || response.StatusCode != http.StatusOK {
	//		c.Status(http.StatusServiceUnavailable)
	//		return
	//	}
	//
	//	reader := response.Body
	//	contentLength := response.ContentLength
	//	contentType := response.Header.Get("Content-Type")
	//
	//	extraHeaders := map[string]string{
	//		"Content-Disposition": `attachment; filename="gopher.png"`,
	//	}
	//
	//	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//})

	// 25. HTML rendering
	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})

	//26. Using templates with same name in different directories
	// 在不同目录中使用具有相同名称的模板
	//router.LoadHTMLGlob("templates/**/*")
	//router.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	//		"title": "Posts",
	//	})
	//})
	//router.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	//		"title": "Users",
	//	})
	//})

	// 27. Redirects
	//router.GET("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	//})
	//
	//router.POST("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/foo")
	//})

	//router.GET("/test", func(c *gin.Context) {
	//	c.Request.URL.Path = "/test2"
	//	router.HandleContext(c)
	//})
	//router.GET("/test2", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"hello": "world"})
	//})

	// 28. Custom Middleware
	//router.Use(middleware.Logger())
	//
	//router.GET("/test", func(c *gin.Context) {
	//	example := c.MustGet("example").(string)
	//
	//	// it would print: "12345"
	//	log.Println(example)
	//})

	// 29. Using BasicAuth() middleware
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	router.Run(":8088")

}
