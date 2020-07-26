package main

import "light-gin/router"

func main()  {
	app := router.InitRouter()

	app.Run()
}
