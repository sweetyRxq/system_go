package appRouter

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	log "test.go/lib/log4go"
	"test.go/model"
	"test.go/systematic"
)
//GENERATE_START
func ProcessBusiness_Invoke(stub shim.ChaincodeStubInterface, tranCode string, requestMsg *systematic.Message, returnMsg *systematic.Message ) {

	switch tranCode{
	case "CreateCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " CreateCommentLog perform start")

		//判断新增数据是否涉及到了附件的上传操作
		var hasFile string
		for _, argsItem := range requestMsg.Args{
			if argsItem.Name == "hasFile"{
				hasFile = argsItem.Value
			}
		}

		var commentLogJSON string
		var fileArr []string
		//如果涉及到了附件信息
		if hasFile == "true"{
			dataArr, err := systematic.GetDataFromMessage(requestMsg) //获取到的是个data数组
			if err != nil{
				systematic.CombinationErrorObj(err.Error(), "CreateCommentLog", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}

			for _, dataItem := range dataArr{
				if dataItem.DataType == "CommentLog"{
					commentLogJSON = dataItem.Content
				}else if dataItem.DataType == "FileInfo"{
					fileArr = append(fileArr, dataItem.Content)
				}
			}
		}else if hasFile == "false"{
			//如果是单纯的新增业务数据
			tmpCommentLogJSON, err := systematic.GetOneDataFromMessage(requestMsg)
			if err != nil{
				systematic.CombinationErrorObj(err.Error(), "CreateCommentLog", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}

			commentLogJSON = tmpCommentLogJSON
		}

		err, errCode := model.AddCommentLog(stub, commentLogJSON, fileArr)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "CreateCommentLog", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		log.Info("Return Message Id:" + returnMsg.Id + " CreateCommentLog perform finish")
		break
	case "DelCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " DelCommentLog perform start")
		var args = systematic.GetArgsOfDelOrQuery(requestMsg)
		var id string

		for _, argItem := range args{
			switch argItem.Name {
			case "id":
				id = argItem.Value
				break
			}
		}

		err, errCode := model.DelCommentLog(stub, id)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "DelCommentLog", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}
		log.Info("Return Message Id:" + returnMsg.Id + " DelCommentLog perform finish")
		break
	case "UpdateCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " UpdateCommentLog perform start")
		commentLogJSON, err := systematic.GetOneDataFromMessage(requestMsg)
		if err != nil{
			log.Error(systematic.ERRNMCC0012)
			systematic.CombinationErrorObj(err.Error(), "UpdateCommentLog", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		err, errCode := model.UpdateCommentLog(stub, commentLogJSON)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "UpdateCommentLog", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}
		log.Info("Return Message Id:" + returnMsg.Id + " UpdateCommentLog perform finish")
		break
	case "CreateTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " CreateTopic perform start")

		//判断新增数据是否涉及到了附件的上传操作
		var hasFile string
		for _, argsItem := range requestMsg.Args{
			if argsItem.Name == "hasFile"{
				hasFile = argsItem.Value
			}
		}

		var topicJSON string
		var fileArr []string
		//如果涉及到了附件信息
		if hasFile == "true"{
			dataArr, err := systematic.GetDataFromMessage(requestMsg) //获取到的是个data数组
			if err != nil{
				systematic.CombinationErrorObj(err.Error(), "CreateTopic", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}

			for _, dataItem := range dataArr{
				if dataItem.DataType == "Topic"{
					topicJSON = dataItem.Content
				}else if dataItem.DataType == "FileInfo"{
					fileArr = append(fileArr, dataItem.Content)
				}
			}
		}else if hasFile == "false"{
			//如果是单纯的新增业务数据
			tmpTopicJSON, err := systematic.GetOneDataFromMessage(requestMsg)
			if err != nil{
				systematic.CombinationErrorObj(err.Error(), "CreateTopic", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}

			topicJSON = tmpTopicJSON
		}

		err, errCode := model.AddTopic(stub, topicJSON, fileArr)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "CreateTopic", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		log.Info("Return Message Id:" + returnMsg.Id + " CreateTopic perform finish")
		break
	case "DelTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " DelTopic perform start")
		var args = systematic.GetArgsOfDelOrQuery(requestMsg)
		var id string

		for _, argItem := range args{
			switch argItem.Name {
			case "id":
				id = argItem.Value
				break
			}
		}

		err, errCode := model.DelTopic(stub, id)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "DelTopic", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}
		log.Info("Return Message Id:" + returnMsg.Id + " DelTopic perform finish")
		break
	case "UpdateTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " UpdateTopic perform start")
		topicJSON, err := systematic.GetOneDataFromMessage(requestMsg)
		if err != nil{
			log.Error(systematic.ERRNMCC0012)
			systematic.CombinationErrorObj(err.Error(), "UpdateTopic", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		err, errCode := model.UpdateTopic(stub, topicJSON)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "UpdateTopic", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}
		log.Info("Return Message Id:" + returnMsg.Id + " UpdateTopic perform finish")
		break
	case "FileInvoke":
		log.Info("Request Message Id:" + requestMsg.Id + " FileInvoke perform start")

		var fileArr []string

		dataArr, err := systematic.GetDataFromMessage(requestMsg) //获取到的是个data数组
		if err != nil {
			systematic.CombinationErrorObj(err.Error(), "FileInvoke", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		for _, dataItem := range dataArr {
			if dataItem.DataType == "FileInfo" {
				fileArr = append(fileArr, dataItem.Content)
			}
		}

		err, errCode := systematic.AddFile(stub, fileArr)
		if err != nil {
			systematic.CombinationErrorObj(err.Error(), "FileInvoke", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		log.Info("Return Message Id:" + returnMsg.Id + " FileInvoke perform finish")
		break
	default:
		var errDescription = fmt.Sprintf("Incorrect invoke transaction type routing function [ %s ]!", tranCode)
		log.Error(errDescription)
		systematic.CombinationErrorObj(errDescription, "ProcessBusiness_Invoke", "ERR_METHOD_NOT_FOUND", returnMsg)
		systematic.ErrorMessage(returnMsg)
	}
}

func ProcessBusiness_Query(stub shim.ChaincodeStubInterface, tranCode string, requestMsg *systematic.Message, returnMsg *systematic.Message ) {
	switch tranCode{
	case "GetVersion": //获取到版本号的接口
		log.Info("Request Message Id:" + requestMsg.Id + " GetVersion perform start")
		systematic.GetVersion(stub, requestMsg, returnMsg)
		log.Info("Return Message Id:" + returnMsg.Id + " GetVersion perform finish")
		break
	case "SelectAllCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " SelectAllCommentLog perform start")
		err, errCode, resArr := model.SelectAllCommentLog(stub)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectAllCommentLog", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			for _, item := range resArr{
				var data systematic.Data
				data.Content = item
				data.DataType = "CommentLog"

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}
		log.Info("Return Message Id:" + returnMsg.Id + " SelectAllCommentLog perform finish")
		break
	case "SelectCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " SelectCommentLog perform start")
		commentLogJSON, err := systematic.GetOneDataFromMessage(requestMsg)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectCommentLog", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		var commentLogObj *model.CommentLog
		err = json.Unmarshal([]byte(commentLogJSON), &commentLogObj)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectCommentLog", systematic.ERRNMCC0001, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			
			var commentLogKey = commentLogObj.Uuid//获取主键id
			if commentLogKey == ""{
				systematic.CombinationErrorObj(err.Error(), "SelectCommentLog", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}else {
				err, errCode, result := model.SelectCommentLog(stub, commentLogKey)
				if err != nil{
					systematic.CombinationErrorObj(err.Error(), "SelectCommentLog", errCode, returnMsg)
					systematic.ErrorMessage(returnMsg)
				}

				var data systematic.Data
				data.DataType = "CommentLog"
				data.Content = result

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}

		log.Info("Return Message Id:" + returnMsg.Id + " SelectCommentLog perform finish")
		break
	case "QueryCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " QueryCommentLog perform start")
		// 获取QueryString
		var args = systematic.GetArgsOfDelOrQuery(requestMsg)
		var queryString string
		for _, argsItem := range args {
			switch argsItem.Name {
			case "queryString":
				queryString = argsItem.Value
				break
			}
		}
		if queryString == "" {
			tempQueryString, err := systematic.GetOneDataFromMessage(requestMsg)
			if err != nil {
				queryString = tempQueryString
			}
		}
		// 如果queryString为空
		if queryString == "" {
			queryString = "{\"dataType\":\"CommentLog\"}"
		}
		err, errCode, resArr := model.QueryCommentLog(stub, queryString)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectAllCommentLog", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			for _, item := range resArr{
				var data systematic.Data
				data.Content = item
				data.DataType = "CommentLog"

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}
		log.Info("Return Message Id:" + returnMsg.Id + " QueryCommentLog perform finish")
		break
	case "PaginateCommentLog":
		log.Info("Request Message Id:" + requestMsg.Id + " PaginateCommentLog perform start")
		systematic.PaginateQuery(stub, requestMsg, returnMsg)
		log.Info("Return Message Id:" + returnMsg.Id + " PaginateCommentLog perform finish")
		break;
	case "SelectAllTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " SelectAllTopic perform start")
		err, errCode, resArr := model.SelectAllTopic(stub)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectAllTopic", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			for _, item := range resArr{
				var data systematic.Data
				data.Content = item
				data.DataType = "Topic"

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}
		log.Info("Return Message Id:" + returnMsg.Id + " SelectAllTopic perform finish")
		break
	case "SelectTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " SelectTopic perform start")
		topicJSON, err := systematic.GetOneDataFromMessage(requestMsg)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectTopic", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		var topicObj *model.Topic
		err = json.Unmarshal([]byte(topicJSON), &topicObj)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectTopic", systematic.ERRNMCC0001, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			
			var topicKey = topicObj.Uuid//获取主键id
			if topicKey == ""{
				systematic.CombinationErrorObj(err.Error(), "SelectTopic", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}else {
				err, errCode, result := model.SelectTopic(stub, topicKey)
				if err != nil{
					systematic.CombinationErrorObj(err.Error(), "SelectTopic", errCode, returnMsg)
					systematic.ErrorMessage(returnMsg)
				}

				var data systematic.Data
				data.DataType = "Topic"
				data.Content = result

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}

		log.Info("Return Message Id:" + returnMsg.Id + " SelectTopic perform finish")
		break
	case "QueryTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " QueryTopic perform start")
		// 获取QueryString
		var args = systematic.GetArgsOfDelOrQuery(requestMsg)
		var queryString string
		for _, argsItem := range args {
			switch argsItem.Name {
			case "queryString":
				queryString = argsItem.Value
				break
			}
		}
		if queryString == "" {
			tempQueryString, err := systematic.GetOneDataFromMessage(requestMsg)
			if err != nil {
				queryString = tempQueryString
			}
		}
		// 如果queryString为空
		if queryString == "" {
			queryString = "{\"dataType\":\"Topic\"}"
		}
		err, errCode, resArr := model.QueryTopic(stub, queryString)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "SelectAllTopic", errCode, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			for _, item := range resArr{
				var data systematic.Data
				data.Content = item
				data.DataType = "Topic"

				returnMsg.Data = append(returnMsg.Data, data)
			}
		}
		log.Info("Return Message Id:" + returnMsg.Id + " QueryTopic perform finish")
		break
	case "PaginateTopic":
		log.Info("Request Message Id:" + requestMsg.Id + " PaginateTopic perform start")
		systematic.PaginateQuery(stub, requestMsg, returnMsg)
		log.Info("Return Message Id:" + returnMsg.Id + " PaginateTopic perform finish")
		break;
	case "FileQuery":
		log.Info("Request Message Id:" + requestMsg.Id + " FileQuery perform start")
		fileJSON, err := systematic.GetOneDataFromMessage(requestMsg)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "FileQuery", systematic.ERRNMCC0012, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}

		var fileObj *systematic.FileInfo
		err = json.Unmarshal([]byte(fileJSON), &fileObj)
		if err != nil{
			systematic.CombinationErrorObj(err.Error(), "FileQuery", systematic.ERRNMCC0001, returnMsg)
			systematic.ErrorMessage(returnMsg)
		}else {
			var fileKey = fileObj.FileId
			if fileKey == ""{
				systematic.CombinationErrorObj(err.Error(), "FileQuery", systematic.ERRNMCC0012, returnMsg)
				systematic.ErrorMessage(returnMsg)
			}else {
				err, errCode, result := systematic.SelectFile(stub, fileKey)
				if err != nil{
					systematic.CombinationErrorObj(err.Error(), "FileQuery", errCode, returnMsg)
					systematic.ErrorMessage(returnMsg)
				}

				if result != ""{
					var data systematic.Data
					data.DataType = "FileInfo"
					data.Content = result

					returnMsg.Data = append(returnMsg.Data, data)
				}
			}
		}

		log.Info("Return Message Id:" + returnMsg.Id + " FileQuery perform finish")
		break
	default:
		var errDescription = fmt.Sprintf("Incorrect query transaction type routing function [ %s ]!", tranCode)
		log.Error(errDescription)
		systematic.CombinationErrorObj(errDescription, "ProcessBusiness_Query", "ERR_METHOD_NOT_FOUND", returnMsg)
		systematic.ErrorMessage(returnMsg)
	}
}
//GENERATE_END