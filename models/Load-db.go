package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func LoadDb() {
	user := Partner{Relation: "friend"}
	user_added, err := Post(user)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	logs.Info("user added")
	fmt.Println("user added successfully, id : ", &user_added)
	fmt.Println("user added successfully, id : ", user_added)
}
