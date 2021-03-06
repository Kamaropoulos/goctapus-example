package main

import (
	"os"

	"github.com/Kamaropoulos/goctapus"
	"github.com/Kamaropoulos/goctapus-example/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args, "debug")

	goctapus.ConnectDB("goapp")
	goctapus.Migrate(goctapus.Databases["goapp"], "./models/tasks.sql")

	goctapus.AddStatic(goctapus.Route{
		Path: "/",
		File: "public/index.html",
		Rate: 10,
	})

	goctapus.AddEndpoint(goctapus.Route{
		Method:  "GET",
		Path:    "/tasks",
		Handler: handlers.GetTasks(goctapus.Databases["goapp"]),
		Rate:    10,
	})

	goctapus.AddEndpoint(goctapus.Route{
		Method:  "PUT",
		Path:    "/tasks",
		Handler: handlers.PutTask(goctapus.Databases["goapp"]),
		Rate:    10,
	})

	goctapus.AddEndpoint(goctapus.Route{
		Method:  "DELETE",
		Path:    "/tasks/:id",
		Handler: handlers.DeleteTask(goctapus.Databases["goapp"]),
		Rate:    10,
	})

	goctapus.Start()
}
