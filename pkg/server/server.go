package server

import (
	"net/http"

	log "github.com/rs/zerolog/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/orangesys/orangeapi/pkg/common"
	"github.com/orangesys/orangeapi/pkg/controller"
	"github.com/orangesys/orangeapi/pkg/storage"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func storageusage(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	consumerID := c.QueryParam("consumerID")
	if err := controller.CheckConsumer(uuid); err != nil {
		log.Error().Msgf("Not Found UUID in Firebase: %v", err)
		return c.String(http.StatusNotFound, "Not Found UUID in Firebase")
	}
	i := storage.InfluxDBClient(consumerID)
	s, err := storage.GetStorageUsed(i)
	if err != nil {
		log.Error().Msgf("Not Found host in orangesys-k8s: %v", err)
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
			log.Error().Msgf("Not Found UUID in Firebase: %v", err)
			return c.String(http.StatusNotFound, "Not Found UUID in Firebase")
		}

		name := common.ReleaseName()
		if err := controller.CreateConsumer(name, retention, pvcsize, uuid); err != nil {
			log.Error().Msgf("Some Wrong with UUI or RP: %v", err)
			return c.String(http.StatusNotFound, "Some Wrong with UUI or RP")
		}
		return c.String(http.StatusOK, "Processing")
	}
	return c.String(http.StatusNotFound, "Not Found Retention Plan with RP")
}

// Run is start echo server
func Run() {
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
