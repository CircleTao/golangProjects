package main

import (
	"golangproject4/router"
)

func main() {
	r := router.Router()

	r.Run(":9999")
}
