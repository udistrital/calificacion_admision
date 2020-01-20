package controllers

import (

	// "strconv"

	"github.com/astaxie/beego"
)

// operations for TrArchivoIcfesController
type TrArchivoIcfesController struct {
	beego.Controller
}

func (c *TrArchivoIcfesController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrArchivoIcfesController
// @Description create the TrArchivoIcfesController
// @Param	body		body 	models.TrArchivoIcfes	true	"body for TrArchivoIcfesController content"
// @Success 201 {int} models.TrArchivoIcfes
// @Failure 400 the request contains incorrect syntax
// @router / [post]
// func (c *TrArchivoIcfesController) Post() {
// 	var v models.TrArchivoIcfes
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
// 		if err := models.AddTransaccionArchivoIcfes(&v); err == nil {
// 			c.Ctx.Output.SetStatus(201)
// 			c.Data["json"] = v
// 		} else {
// 			logs.Error(err)
// 			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
// 			c.Data["system"] = err
// 			c.Abort("400")
// 		}
// 	} else {
// 		logs.Error(err)
// 		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
// 		c.Data["system"] = err
// 		c.Abort("400")
// 	}
// 	c.ServeJSON()
// }
