package main

import (
	"fmt"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	appid := "1259654469"
	bucket := "test-1259654469"
	c := sts.NewClient(
		os.Getenv("COS_SECRETID"),
		os.Getenv("COS_SECRETKEY"),
		nil,
	)
	// 设置域名, 默认域名sts.tencentcloudapi.com
	// c.SetHost("")
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-guangzhou",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:ap-guangzhou:uid/" + appid + ":" + bucket + "/exampleobject",
					},
				},
			},
		},
	}
	// case 1 请求临时密钥
	res, err := c.GetCredential(opt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Credentials)

	// case 2 发起临时密钥请求，需自行解析密钥，自行判断临时密钥是否请求成功
	resp, err := c.RequestCredential(opt)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body:%v\n", string(bs))
}
