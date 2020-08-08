package v1

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/response"
	"cmdb/forms"
	"cmdb/services"
	"encoding/json"
)

type PrometheusController struct {
	auth.APIController
}

func (c *PrometheusController) Register() {
	form := &forms.NodeRegisterForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, form); err == nil {
		// 验证
		services.NodeService.Register(form)
		c.Data["json"] = response.Ok
	} else {
		c.Data["json"] = response.BadRequest
	}
}

func (c *PrometheusController) Config() {
	uuid := c.GetString("uuid")
	// job target
	/*
		[
			{
				"key" : "",
				"targets": [
					{"addr" : ""}, {"addr" : ""}
				]
			}
		]
	*/
	rt := services.JobService.QueryByUUID(uuid)

	c.Data["json"] = response.NewJsonResponse(200, "ok", rt)
}
