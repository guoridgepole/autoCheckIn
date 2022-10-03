package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
什么值得买自动签到
*/
func main() {
	client := &http.Client{}
	/**
	https://zhiyou.smzdm.com/user/checkin/jsonp_checkin?callback=jQuery1124017798994082517394_1664604435190&_=1664604435204

	失败时候
	{
		"error_code": 9999,
		"error_msg": "",
		"data": []
	}

	成功时候
	{
		"error_code": 0,
		"error_msg": "",
		"data": {
			"add_point": 0,
			"checkin_num": "464",
			"point": 432,
			"exp": 2429,
			"gold": 63,
			"prestige": 0,
			"rank": 9,
			"slogan": "<div class=\"signIn_data\">\u4eca\u65e5\u5df2\u9886<span class=\"red\">0<\/span>\u79ef\u5206<\/div>",
			"cards": "3",
			"can_contract": 0,
			"continue_checkin_days": 3,
			"continue_checkin_reward_show": false
		}
	}

	*/
	url := "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	header := request.Header
	header.Set("Cookie", "这里写入自己的cookie")
	header.Set("Host", "zhiyou.smzdm.com")
	header.Set("Referer", "https://www.smzdm.com/")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var smzdmBody shenmezhidemaiRespBody
	err = json.Unmarshal(body, &smzdmBody)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	if smzdmBody.ErrorCode != 0 {
		fmt.Println(now, "什么值得买签到失败，失败原因:", smzdmBody.ErrorMsg)
	} else {
		fmt.Println(now, "什么值得买签到成功:", smzdmBody.Data)
	}

}

type shenmezhidemaiRespBody struct {
	ErrorCode int                    `json:"error_code"`
	ErrorMsg  string                 `json:"error_msg"`
	Data      map[string]interface{} `json:"data"`
}
