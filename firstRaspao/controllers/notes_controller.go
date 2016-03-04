package controllers

import (
	"encoding/json"

	"github.com/efrenospino/firstRaspao/config"
	"github.com/efrenospino/firstRaspao/models"
	raspao "github.com/wawandco/raspao/base"
)

func init() {
	raspao.RegisterController("NotesController", &NotesController{})
}

type NotesController struct {
	raspao.Controller
}

func (n *NotesController) Initialize(context *raspao.Context) {
	db := config.SharedApplication.Database
	db.DropTable(&models.Note{})
	db.CreateTable(&models.Note{})
}

func (n *NotesController) Index(context *raspao.Context) {
	db := config.SharedApplication.Database
	var notes []models.Note
	var count int
	db.Model(&models.Note{}).Count(&count)
	db.Find(&notes)

	rs := struct {
		Meta map[string]int `json:"meta"`
		Data []models.Note  `json:"data"`
	}{
		Meta: map[string]int{
			"total": count,
		},
		Data: notes,
	}

	b, _ := json.Marshal(rs)
	context.Response.Write([]byte(b))
}

func (n *NotesController) Create(context *raspao.Context) {
	db := config.SharedApplication.Database
	note := models.Note{}
	db.Create(&note)
}

func (n *NotesController) Delete(context *raspao.Context) {
	db := config.SharedApplication.Database
	noteID := context.Params["id"][0]
	db.Where("ID = ?", noteID).Delete(models.Note{})
}

func (n *NotesController) Find(context *raspao.Context) {
	db := config.SharedApplication.Database
	noteID := context.Params["id"][0]
	note := models.Note{}
	db.Where("ID = ?", noteID).First(&note)

	rs := struct {
		Data models.Note `json:"data"`
	}{
		Data: note,
	}

	b, _ := json.Marshal(rs)
	context.Response.Write([]byte(b))
}
