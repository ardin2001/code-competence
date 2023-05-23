package routes

import (
	"net/http"
	"os"

	"echo_golang/configs"
	"echo_golang/controllers"
	m "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/repositories"
	"echo_golang/services"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	db, _ = configs.InitDB()

	userR = repositories.NewUserRepositories(db)
	userS = services.NewUserServices(userR)
	userC = controllers.NewUserControllers(userS)

	productR = repositories.NewProductRepositories(db)
	productS = services.NewProductServices(productR)
	productC = controllers.NewProductControllers(productS)

	categoryR = repositories.NewCategoryRepositories(db)
	categoryS = services.NewCategoryServices(categoryR)
	categoryC = controllers.NewCategoryControllers(categoryS)
)

func Routers() *echo.Echo {
	e := echo.New()
	m.Logger(e)
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(dbHost),
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, login here !, localhost:8000/login")
	})

	// users
	e.POST("/login", userC.LoginUserController)
	e.POST("/register", userC.CreateUserController)
	e.GET("/users", userC.GetUserController, echojwt.WithConfig(config))
	e.PUT("/users", userC.UpdateUserController, echojwt.WithConfig(config))
	e.DELETE("/users", userC.DeleteUserController, echojwt.WithConfig(config))

	// products
	e.GET("/items", productC.GetProductsController, echojwt.WithConfig(config))
	e.GET("/items/:id", productC.GetProductController, echojwt.WithConfig(config))
	e.POST("/items", productC.CreateProductController, echojwt.WithConfig(config))
	e.PUT("/items/:id", productC.UpdateProductController, echojwt.WithConfig(config))
	e.DELETE("/items/:id", productC.DeleteProductController, echojwt.WithConfig(config))

	// categories
	e.GET("/categories", categoryC.GetCategoriesController, echojwt.WithConfig(config))
	e.GET("/items/category/:category_id", categoryC.GetCategoryController, echojwt.WithConfig(config))
	e.POST("/categories", categoryC.CreateCategoryController, echojwt.WithConfig(config))
	e.PUT("/categories/:id", categoryC.UpdateCategoryController, echojwt.WithConfig(config))
	e.DELETE("/categories/:id", categoryC.DeleteCategoryController, echojwt.WithConfig(config))

	return e
}
