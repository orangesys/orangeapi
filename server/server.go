package server

import (
    "log"
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "github.com/orangesys/orangeapi/controller"
    "github.com/orangesys/orangeapi/common"
)

func accessible(c echo.Context) error {
    return c.String(http.StatusOK, "Accessible")
}

func create(c echo.Context) error {
	rp   := c.QueryParam("rp")
	uuid := c.QueryParam("uuid")

        pvc := map[string]string{
             "10d": "10Gi",
             "40d": "50Gi",
             "400d": "100Gi",
        }
	wp := "mypassword"
//	uuid := "iGzNX6QzfudVlwKtR8CQCj0itIU2"

	if _, ok := pvc[rp]; ok {
		if err := controller.CheckConsumer(uuid); err != nil {
			return c.String(http.StatusNotFound, "Not Found UUID in Firebase")
		}

		name := common.ReleaseName()
		if err := controller.CreateConsumer(name, wp, uuid); err != nil {
			return c.String(http.StatusNotFound, "Some Wrong with UUI or RP")
		}
		return c.String(http.StatusOK, "Processing")
	}

	return c.String(http.StatusNotFound, "Not Found Retention Plan with RP")
}


func Run() {
    log.Println("Starting orangeapi...")
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		echo.GET,
		echo.POST,
	},
    }))

    // Login route
    e.POST("/create", create)

    // Unauthenticated route
    e.GET("/", accessible)

    e.Logger.Fatal(e.Start(":1323"))
}
