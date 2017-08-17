package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type DataSet struct {
	Data []int
}

func main() {
	e := echo.New()

	e.File("/", "charts-example/index.html")

	dataSet := DataSet{Data: []int{40, 39, 10, 40, 39, 80}}

	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dataSet)
	})

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
