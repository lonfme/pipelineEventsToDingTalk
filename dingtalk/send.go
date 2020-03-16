package dingtalk

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendMdMessage(mdText, accessToken string) error {
	message := MarkdownMessage{ // 构建 post 消息体
		MsgType: MsgTypeMarkdown,
		Markdown: MarkdownParams{
			Title: "Gitlab-CI 通知",
			Text:  mdText,
		}}

	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	dingReq, err := http.NewRequest("POST",
		"https://oapi.dingtalk.com/robot/send", body)
	if err != nil {
		return err
	}
	dingReq.Header.Set("Content-Type", "application/json")

	params := dingReq.URL.Query()
	params.Add("access_token", accessToken)
	dingReq.URL.RawQuery = params.Encode()

	dingResp, err := http.DefaultClient.Do(dingReq) // 发送请求到钉钉
	if err != nil {
		return err
	}
	if dingResp != nil {
		defer dingResp.Body.Close()
	}

	return nil
}
