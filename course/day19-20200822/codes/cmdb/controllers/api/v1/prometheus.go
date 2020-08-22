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

func (c *PrometheusController) Alert() {
	var form forms.AlertGroupForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		// 处理告警组
		services.AlertService.Notice(&form)
		// 处理告警
		for _, alert := range form.Alerts {
			services.AlertService.Alert(alert)
		}
	}

	c.Data["json"] = response.NewJsonResponse(200, "ok", nil)
	// fmt.Println(string(c.Ctx.Input.RequestBody))
	// gjson.GetBytes(c.Ctx.Input.RequestBody, "alerts").ForEach(func(key, alert gjson.Result) bool {
	// 	var form forms.AlertForm
	// 	if err := json.Unmarshal([]byte(alert.Raw), &form); err == nil {
	// 		services.AlertService.Alert(&form)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// 	return true
	// })
}
