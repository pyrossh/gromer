package main

import (
	"github.com/pyros2097/gromer"
	_ "github.com/pyros2097/gromer/_example/components"
	_ "github.com/pyros2097/gromer/_example/routes"
)

func main() {
	gromer.Run("3000")
}
