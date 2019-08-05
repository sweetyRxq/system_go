package model

type CommentLog struct {
//GENERATE_START 
 
	Name string`json:"name"` 
	Uuid string`json:"uuid"`	 
	CommentDate string`json:"commentDate"`	 
	Content string`json:"content"`	 
	CommonetUser string`json:"commonetUser"`	
	DataType string `json:"dataType"`

//GENERATE_END
}
