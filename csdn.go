package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	csdn自动签到
*/
func main() {
	client := &http.Client{}
	/*
		https://me.csdn.net/api/LuckyDraw_v2/signIn
		签到成功的时候

		重复签到的结果
		{
			"code": 200,
			"message": "成功",
			"data": {
			"msg": "用户已签到",
			"isSigned": true
			}
		}
	*/
	var cookie = getCsdnLoginInfo()
	var url = "https://me.csdn.net/api/LuckyDraw_v2/signIn"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	header := request.Header
	header.Set("Cookie", cookie)
	header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var csdnRespBody csdnRespBody
	fmt.Println(string(body))
	err = json.Unmarshal(body, &csdnRespBody)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	if csdnRespBody.Code != 200 {
		fmt.Println(now, "csnd自动签到失败,失败原因:", csdnRespBody.Message)
	} else {
		fmt.Println(now, "csdn签到成功", csdnRespBody.Message)
	}

}

/*
	从配置文件中取出csdn登录信息
*/
func getCsdnLoginInfo() string {
	file, err := ioutil.ReadFile("./login_info.yaml")
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	login := data["login"]
	cookie := login.(map[string]interface{})["cookie"]
	csdn := cookie.(map[string]interface{})["csdn"]

	return csdn.(string)
}

type csdnRespBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
