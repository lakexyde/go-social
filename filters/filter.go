package filters

import (
    "github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context){
    _, ok := ctx.Input.Session("lakex").(int)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}