package main

import routers "gin-framework/router"

func main() {

	var PORT = ":7070"

	routers.StartServer().Run(PORT)

}
