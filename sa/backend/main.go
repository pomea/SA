package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pomea/app/controllers"
	_ "github.com/pomea/app/docs"
	"github.com/pomea/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Users is ...
type Users struct {
	User []User
}

// User is ...
type User struct {
	Name  string
	Email string
}

// Menutypes is ...
type Menutypes struct {
	Menutype []Menutype
}

// Menutype is ...
type Menutype struct {
	Type string
}

// Menugroups is ...
type Menugroups struct {
	Menugroup []Menugroup
}

// Menugroup is ...
type Menugroup struct {
	Group string
}

// @title SUT SA Example API Menu
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

	// v1 for controller....
	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewMenutypeController(v1, client)
	controllers.NewMenugroupController(v1, client)
	controllers.NewMenuController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"Thanapoom Pongpracha", "poom@gmail.com"},
			User{"Name Surname", "me@example.com"},
			User{"Tanjiro Omakase", "Tanjiro@gmail.com"},
			User{"The O  german", "TheO@hotmail.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.Email).
			SetName(u.Name).
			Save(context.Background())
	}

	menutypes := Menutypes{
		Menutype: []Menutype{
			Menutype{"meat dish"},
			Menutype{"dessert"},
			Menutype{"snack"},
			Menutype{"drink"},
			
		},
	}

	for _, t := range menutypes.Menutype {
		client.Menutype.
			Create().
			SetType(t.Type).
			Save(context.Background())
	}

	menugroups := Menugroups{
		Menugroup: []Menugroup{
			Menugroup{"instant food"},
			Menugroup{"soup"},
			Menugroup{"noodle"},
			Menugroup{"fried"},
			Menugroup{"seafood"},
			Menugroup{"steamed"},
			Menugroup{"curry"},
			Menugroup{"boiled"},
			Menugroup{"pasta"},
			Menugroup{"fast food"},
		},
	}
	for _, g := range menugroups.Menugroup {
		client.Menugroup.
			Create().
			SetGroup(g.Group).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
