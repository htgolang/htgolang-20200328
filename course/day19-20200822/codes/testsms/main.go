package main

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func main() {

	credential := common.NewCredential(
		"AKID44qsOsP1g5GB9qxu1ndW8CzuZIYYfr3y",
		"5qYz4uSnrUrJUe5GDvLUheZEcsQYQKyZ",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf) // http client

	request := sms.NewSendSmsRequest() // 参数

	params := "{\"PhoneNumberSet\":[\"+8618696179145\",\"+8618618318200\"],\"TemplateID\":\"691782\",\"Sign\":\"iamuk网\",\"TemplateParamSet\":[\"各位\",\"账户余额\",\"CMDB2\"],\"SmsSdkAppid\":\"1400287583\"}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
