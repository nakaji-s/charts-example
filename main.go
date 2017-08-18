package main

import (
	"net/http"

	"math/rand"
	"time"

	"github.com/labstack/echo"
)

type DataSet struct {
	Data []int
}

func (d *DataSet) Add(num int) {
	d.Data = append(d.Data[1:], num)
}

func main() {
	e := echo.New()

	e.File("/", "charts-example/index.html")

	dataSet := DataSet{Data: []int{}}
	for i := 0; i < 100; i++ {
		dataSet.Data = append(dataSet.Data, rand.Intn(100))
	}
	go func() {
		for {
			dataSet.Add(rand.Intn(100))
			time.Sleep(1 * time.Second)
		}
	}()

	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dataSet)
	})

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
