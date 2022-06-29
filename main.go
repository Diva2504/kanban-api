package main

import (
	"log"

	"github.com/takadev15/kanban-api/config"
	"github.com/takadev15/kanban-api/routers"
)

func main(){
  config.InitDB()
  r := routers.RoutesList()
  log.Fatal(r.Run(":3030"))
}
