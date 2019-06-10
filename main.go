package main

import (
	"fmt"

	"github.com/identixone/identixone-go/api/notification"
	"github.com/identixone/identixone-go/api/utility"
)

func main() {
	c := notification.CreateRequest{}

	err := c.Validate()
	fmt.Println(err)

	req := utility.AsmRequest{}
	err = req.FromFile("img/v2878.png")
	fmt.Println(err)
}
