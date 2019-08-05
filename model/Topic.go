package model

type Topic struct {
//GENERATE_START 
 
	Name string`json:"name"` 
	Uuid string`json:"uuid"`	 
	TopicName string`json:"topicName"`	 
	CommentLogs []CommentLog`json:"commentLogs"`	
	DataType string `json:"dataType"`

//GENERATE_END
}
