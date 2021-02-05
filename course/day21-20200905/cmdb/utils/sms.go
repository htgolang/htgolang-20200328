package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func SendSms(phones []string, templateId string, templateParams []string) error {
	credential := common.NewCredential(
		beego.AppConfig.DefaultString("tencent_sms::secretId", "AKID44qsOsP1g5GB9qxu1ndW8CzuZIYYfr3y"),
		beego.AppConfig.DefaultString("tencent_sms::secretKey", "5qYz4uSnrUrJUe5GDvLUheZEcsQYQKyZ"),
	)
	sign := beego.AppConfig.DefaultString("tencent_sms::sign", "iamuk网")
	appId := beego.AppConfig.DefaultString("tencent_sms::appId", "1400287583")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf) // http client

	request := sms.NewSendSmsRequest() // 参数

	phoneSet := make([]*string, len(phones))
	templateParamSet := make([]*string, len(templateParams))
	for i, phone := range phones {
		phoneSet[i] = &phone
	}
	for i, param := range templateParams {
		templateParamSet[i] = &param
	}

	request.PhoneNumberSet = phoneSet
	request.TemplateID = &templateId
	request.TemplateParamSet = templateParamSet
	request.SmsSdkAppid = &appId
	request.Sign = &sign
	response, err := client.SendSms(request)
	// fmt.Println(phones, templateParams, err)
	fmt.Printf("%s\n", response.ToJsonString())
	return err
}
