package controllers

import (
	"log"

	raspao "github.com/wawandco/raspao/base"
)

type PrincipalController struct {
	raspao.Controller
}

func init() {
	raspao.RegisterController("PrincipalController", &PrincipalController{})
}

func (p *PrincipalController) Index(context *raspao.Context) {
	log.Println("| Visited PrincipalController.Index")
	context.Response.Write([]byte("Welcome!"))
}
