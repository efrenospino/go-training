package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

//User type for users
type User struct {
	ID   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "db.sqlite3", 30)

	// Error.
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func insertInto(o orm.Ormer, name string) bool {
	user := User{Name: name}
	_, err := o.Insert(&user)
	checkError(err)
	return true
}

func readOneUser(o orm.Ormer, id int) User {
	user := User{ID: id}
	checkError(o.Read(&user))
	return user
}

func updateUser(o orm.Ormer, id int, newName string) bool {
	user := readOneUser(o, id)
	user.Name = newName
	_, err := o.Update(&user)
	checkError(err)
	return true
}

func deleteUser(o orm.Ormer, id int) bool {
	user := User{ID: id}
	_, err := o.Delete(&user)
	checkError(err)
	return true
}

func main() {

	o := orm.NewOrm()

	fmt.Println(insertInto(o, "Efren Ospino"))

}
