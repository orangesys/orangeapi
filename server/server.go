package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/orangesys/orangeapi/common"
	"github.com/orangesys/orangeapi/controller"
	"github.com/orangesys/orangeapi/storage"
)


func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func storageusage(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	consumerId := c.QueryParam("consumerId")
	if err := controller.CheckConsumer(uuid); err != nil {
		log.Println(err)
		return c.String(http.StatusNotFound, "Not Found UUID in Firebase")
	}
	i := storage.InfluxDBClient(consumerId)
	s, err := storage.GetStorageUsed(i)
	if err != nil {
	    log.Println(err)
            return c.String(http.StatusNotFound, "Not Found host in orangesys-k8s")
	}
        var content struct {
                StorageUsage int64 `json:"storageUsage"`
        }
        content.StorageUsage = s

	return c.JSON(http.StatusOK, &content)
}

func create(c echo.Context) error {
	retention := c.QueryParam("rp")
	uuid := c.QueryParam("uuid")
	pvc := map[string]string{
		"10d":  "10Gi",
		"40d":  "50Gi",
		"400d": "100Gi",
	}

	//	uuid := "iGzNX6QzfudVlwKtR8CQCj0itIU2"
	if pvcsize, ok := pvc[retention]; ok {
		if err := controller.CheckConsumer(uuid); err != nil {
			log.Println(err)
			return c.String(http.StatusNotFound, "Not Found UUID in Firebase")
		}

		name := common.ReleaseName()
		if err := controller.CreateConsumer(name, retention, pvcsize, uuid); err != nil {
			log.Println(err)
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

	// Get Storage Used
	e.GET("/storageusage", storageusage)

	// Unauthenticated route
	e.GET("/", accessible)

	e.Logger.Fatal(e.Start(":1323"))
}
