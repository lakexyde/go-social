package controllers

import (
    "github.com/astaxie/beego"
    _"golang.org/x/crypto/bcrypt"
    _"github.com/lakexyde/gosocial/models"
)


type MainController struct{
    beego.Controller
}

func (this *MainController) Homepage(){
    this.TplName = "home/index.tpl"
}

func (this *MainController) Login(){
    v := this.GetSession("lakex")
    if this.Ctx.Request.Method == "POST"{
        if v == nil{
            this.SetSession("lakex", int(1))
            this.Redirect("/lakex/timeline", 302)
        } else {
            this.SetSession("lakex", v.(int)+1)
            this.Redirect("/lakex/timeline", 302)
        }
    }
        

    this.TplName = "home/index.tpl"
}

func (this *MainController) Register(){
//var user models.User
    if this.Ctx.Request.Method == "POST" {
        // if err := this.ParseForm(&user); err == nil{
        //     pass := user.Password
        //     hp, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
        //     if err != nil{
        //         beego.Error(err, "Server error")
        //         return
        //     }
        //     user.Password = string(hp)
        //     err = models.RegisterUser(&user)
        //     if err == nil{
        //         this.Redirect("/", 302)
        //     }

        // }
        

    }
    this.Data["json"] = "Error"
    this.ServeJSON()
}

func (this *MainController) Logout(){
    this.DelSession("lakex")
    this.Redirect("/", 302)
}