package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("golang echo!")
	e := echo.New()
	e.get("")
}
