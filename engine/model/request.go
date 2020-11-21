package model

// 运行容器
type ContainerRunRequest struct {
	TemplateName  string                 `json:"templateName" binding:"required"`  // 模版名
	FunctionParam map[string]interface{} `json:"functionParam" binding:"required"` // 函数参数
	Synchronized  bool                   `json:"synchronized"`                     // 同步等待结果
}

// 函数执行完成
type FunctionEndRequest struct {
	Id             string      `json:"id"  binding:"required"` // 函数实例 id
	Error          interface{} `json:"error"`                  // 错误信息
	FunctionResult interface{} `json:"functionResult" `        // 函数执行结果

	ContainerProcessRunAt int64 `json:"containerProcessRunAt" binding:"required"` // 容器进程开始运行的时间 bootstrap记录的
	FunctionRunTimestamp  int64 `json:"functionRunTimestamp" binding:"required"`  // 函数开始运行的时间
	FunctionEndTimestamp  int64 `json:"functionEndTimestamp" binding:"required"`  // 函数结束运行的时间
}
