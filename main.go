package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pipelineEventsToDingTalk/dingtalk"
	"pipelineEventsToDingTalk/gitlab"
	"strconv"
)

const ReqMD = `### Gitlab-CI 通知

> project name: %s  
> commit user: %s  
> ref: %s  
> status: %s  
> duration: %v  
> commit message: %s  
> pipeline detail: %s  

`

func RobotSendPipeline(rw http.ResponseWriter, req *http.Request) {
	var err error

	accessToken := req.Header.Get("X-Gitlab-Token")

	// 解析 pipeline 信息
	var pipelinePayload gitlab.PipelineEventPayload
	err = json.NewDecoder(req.Body).Decode(&pipelinePayload)
	defer req.Body.Close()
	if err != nil {
		rw.WriteHeader(500)
		_, _ = rw.Write([]byte(err.Error()))
		fmt.Printf("Get GitLab pipelinePayload error: %s", err.Error())
		return
	}

	// 构建 markdown 文本
	mdText := fmt.Sprintf(ReqMD, pipelinePayload.Project.Name, pipelinePayload.User.UserName, // 格式化 markdown
		pipelinePayload.ObjectAttributes.Ref, pipelinePayload.ObjectAttributes.Status,
		pipelinePayload.ObjectAttributes.Duration, pipelinePayload.Commit.Message,
		pipelinePayload.Project.WebURL + "/pipelines/" + strconv.FormatInt(pipelinePayload.ObjectAttributes.ID, 10))

	err = dingtalk.SendMdMessage(mdText, accessToken)
	if err != nil {
		rw.WriteHeader(500)
		_, _ = rw.Write([]byte(err.Error()))
		fmt.Printf("SendMdMessage to DingTalk error: %s", err.Error())
		return
	}

	rw.WriteHeader(200)
}


func main() {
	http.HandleFunc("/robot/send", RobotSendPipeline)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
