package main

import (
	"fmt"
	"log"
	"os"
	"time"

	gocommon "github.com/kurmidev/gocomman"
	"github.com/kurmidev/webox/handler"
	"github.com/kurmidev/webox/models"
	"github.com/kurmidev/webox/router"
)

func main() {
	start := time.Now()
	fmt.Println("Starting processing jobs.......")
	i := InitApplication()
	i.Common.ListenAndServe()
	elapsed := time.Since(start)
	fmt.Printf("Time taken for jobs %s", elapsed)

}

func InitApplication() *handler.Handlers {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	c := &gocommon.Common{}
	err = c.New(path)
	if err != nil {
		log.Fatal(err)
	}
	myHandlers := &handler.Handlers{
		Common: c,
		Models: models.New(c.DB),
	}
	c.AppName = "app"
	c.Routes = router.Routes(myHandlers)

	return myHandlers
}
