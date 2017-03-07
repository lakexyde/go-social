package models

import (
    "github.com/astaxie/beego/orm"
)

type User struct{
    Id int64 `orm:"unique"`
    Username string `orm:"unique;size(30)"`
    Firstname string `orm:""`
    Surname string `orm:""`
    Email string `orm:"unique"`
    Gender string `orm:""`
    Password string `orm:""`
}

func init(){
    orm.RegisterModel(new(User))
}

func (this *User) TableName() string{
    return "users"
}

func RegisterUser(u *User) (err error){
    o := orm.NewOrm()
    a := User{Email: u.Email}
    err = o.Read(&a); 
    if err == nil{
        o.Insert(a)
        return nil
    } else {
        return err
    }
    return
}


