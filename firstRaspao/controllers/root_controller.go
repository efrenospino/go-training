package controllers
import (
  raspao "github.com/wawandco/raspao/base"
  "log"
)

func init() {
	raspao.RegisterController("RootController", &RootController{})
}

type RootController struct {
  raspao.Controller
}

/* Index is an auto-generated route for the landing page of your Raspao app */
func (c *RootController) Index(context *raspao.Context) {
	log.Println("| Visited RootController.Index")
	context.Response.Write([]byte("Welcome To Raspao!"))
}
