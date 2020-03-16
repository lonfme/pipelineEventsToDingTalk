## pipelineEventsToDingTalk

GitLab Pipeline events webhook 事件信息发送到 [钉钉](https://www.dingtalk.com/) 群，使用的是钉钉群通用机器人，目前支持群机器人安全策略是 ip 白名单。


### 使用

1. 运行此项目程序 `go run main.go`，默认监听端口为 `8080`，如服务器 ip 为 `172.17.1.13`

2. 在钉钉群里创建自定义机器人，并拿到机器人的 access_token

3. 在需要通知的 GitLab 项目打开 `Settings` ➔ `Integrations`，然后添加 Pipeline events webhook，URL 为 `http://172.17.1.13:8080/robot/send`， `Secret Token` 为第二步申请到的机器人 access_token 值，只勾选 `Pipeline events`，取消勾选 `Enable SSL verification`

4. 可以发送测试信息，如没问题，可尝试运行 pipelines，验收实际接收到的信息
