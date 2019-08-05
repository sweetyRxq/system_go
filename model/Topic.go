package model

type Topic struct {
//GENERATE_START 
	Name string`json:"name"` 
	Uuid string`json:"uuid"`	 
	CommentLogs []CommentLog`json:"commentLogs"`	 
	TopicName string`json:"topicName"`	
	DataType string `json:"dataType"`
//GENERATE_END
}