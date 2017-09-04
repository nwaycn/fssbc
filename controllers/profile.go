package controllers

import (
	"fmt"

	"nwaycn/fssbc/models"

	"os"
	"runtime"
	"time"

	"github.com/astaxie/beego"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) List() {
	if this.userId > 0 {
		this.redirect("/")
	}

}
