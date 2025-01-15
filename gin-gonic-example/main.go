package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	_ "gin-gonic-example/docs"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	/*
		// Server swagger 2.0 (doc.json)
		r.StaticFile("/swagger/doc.json", "./docs/swagger.json")
		// Swagger UI pointing to Swagger 2.0
		r.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))
	*/

	// Serve OpenAPI 3.0 (openapi3.yaml)
	r.StaticFile("/swagger/openapi3.yaml", "./docs/openapi3.yaml")

	// Swagger UI pointing to OpenAPI 3.0
	r.GET("/swagger/v3/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/openapi3.yaml")))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", Admin)

	//Ejemplos de gin gonic web framework

	//https://gin-gonic.com/docs/examples/ascii-json/

	r.GET("/someJSON", SomeJson)

	//Bind form-data request with custom struct
	r.GET("/getb", GetDataB)

	//Bind html checkboxes
	r.POST("/formHandler", FormHandler)

	///Bind query string or post data
	r.POST("/testing", startPage)

	//Bind Uri
	r.GET("/:name/:id", func(c *gin.Context) {
		var person PersonBindUri
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person})
	})

	// Load templates
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", ReturnTmpl)

	return r
}

// Admin godoc
// @Summary      Ingresar datos especificos
// @Description  Updates a value associated with a user after Basic Auth validation
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        body  body      object  true  "Input Json con algun valor"
// @Success      200   {object}  SuccessResponse  "status: ok"
// @Failure      400   {object}  ErrorResponse    "error: invalid request"
// @Failure      401   {object}  ErrorResponse    "error: unauthorized"
// @Router       /admin [post]
func Admin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, SuccessResponse{Status: "ok"})
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
	}
}

// SuccessResponse represents a successful response
type SuccessResponse struct {
	Status string `json:"status" example:"ok"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

// SomeJson godoc
// @Summary      Show a JSON response
// @Description  Returns a simple JSON object
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Successful Response"
// @Router       /someJSON [get]
func SomeJson(c *gin.Context) {

	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	c.AsciiJSON(http.StatusOK, data)
}

func ReturnTmpl(c *gin.Context) {
	c.HTML(http.StatusOK, "/html/index.tmpl", nil)
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func FormHandler(c *gin.Context) {
	var fakeForm myForm

	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

//Bind query string or post data

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// startPage godoc
// @Summary      Process form data and display result
// @Description  Accepts form data and returns the formatted response as a string
// @Tags         default
// @Accept       application/x-www-form-urlencoded
// @Produce      plain
// @Param        name      formData  string  true  "Name of the person"
// @Param        address   formData  string  true  "Address of the person"
// @Param        birthday  formData  string  true  "Birthday of the person (YYYY-MM-DD)"
// @Success      200  {string}  string  "Formatted response with name, address, and birthday"
// @Failure      400  {object}  map[string]interface{}  "Invalid form data"
// @Router       /testing [post]
func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	// Format the person struct data into the response string
	response := fmt.Sprintf("Success: Name=%s, Address=%s, Birthday=%s",
		person.Name, person.Address, person.Birthday)

	// Send the formatted response
	c.String(200, response)
}

type PersonBindUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
