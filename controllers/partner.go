package controllers

import (
	"encoding/json"
	"strconv"
	"training_project/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type PartnerController struct {
	beego.Controller
}

func (u *PartnerController) Get() {
	uid := u.GetString(":uid")
	uid_i, _ := strconv.Atoi(uid)
	user, err := models.Get(uid_i)
	if err != nil {
		u.Data["json"] = map[string]interface{}{"status": 400, "error": err.Error()}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"status": 200, "data": user}
	u.ServeJSON()
	return
}

func (u *PartnerController) Post() {
	var data map[string]interface{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	var partner models.Partner
	partner.Relation = data["relation"].(string)
	user, err := models.Post(partner)
	if err != nil {
		u.Data["json"] = map[string]interface{}{"status": 400, "error": err.Error()}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"status": 200, "added data": user}
	u.ServeJSON()
	return
}
