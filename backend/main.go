package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jirayuSai/app/controllers"
	"github.com/jirayuSai/app/ent"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Customers struct {
	Customer []Customer
}
type Customer struct {
	Customer string
}
type Genders struct {
	Gender []Gender
}
type Gender struct {
	GenderName string
}
type Personals struct {
	Personal []Personal
}
type Personal struct {
	Name     string
	Email    string
	Password string
}
type Titles struct {
	Title []Title
}
type Title struct {
	TitleName string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//Set gender Data
	genders := Genders{
		Gender: []Gender{
			Gender{"none"},
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}
	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGenderName(g.GenderName).
			Save(context.Background())

	}

	//Set title Data
	titles := Titles{
		Title: []Title{
			Title{"none"},
			Title{"นาย"},
			Title{"นางสาว"},
		},
	}
	for _, t := range titles.Title {
		client.Title.
			Create().
			SetTitleName(t.TitleName).
			Save(context.Background())

	}

	//Set personal Data
	personals := Personals{
		Personal: []Personal{
			Personal{"sai", "sai@ex.com", "12345678"},
			Personal{"som", "som@ex.com", "87654321"},
			Personal{"ben", "ben@ex.com", "65478961"},
		},
	}
	for _, p := range personals.Personal {
		client.Personal.
			Create().
			SetName(p.Name).
			SetEmail(p.Email).
			SetPassword(p.Password).
			Save(context.Background())

	}

	v1 := router.Group("/api/v1")
	controllers.NewCustomerController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewPersonalController(v1, client)
	controllers.NewTitleController(v1, client)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
