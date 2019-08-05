package model

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"fmt"
	log "test.go/lib/log4go"
	"test.go/systematic"
)
//GENERATE_START

	func AddCommentLog(stub shim.ChaincodeStubInterface, CommentLogJson string, fileArr []string) (error, string){
		var CommentLogObj *CommentLog
	
		err := json.Unmarshal([]byte(CommentLogJson), &CommentLogObj)
		if err != nil{
			log.Error(systematic.ERRNMCC0001)
			return err, systematic.ERRNMCC0001
		}
		
		var CommentLogKey = CommentLogObj.Uuid//获取主键id
	
		log.Info("Function Name:AddCommentLog ---- Message:PutState execute...")
		err = stub.PutState(CommentLogKey, []byte(CommentLogJson))
		if err != nil{
			log.Error(systematic.ERRNMCC0005)
			return err, systematic.ERRNMCC0005
		}else {
			log.Info("Function Name:AddCommentLog ---- Message:PutState SUCCESS!!!")
		}
		return nil, ""
	}
	func DelCommentLog(stub shim.ChaincodeStubInterface, CommentLogId string) (error, string) {
		log.Info("Function Name:DelCommentLog ---- Message:DelState execute...")
		//先取到数据，判断要删除的数据信息在数据库中存不存在
		tmpCommentLogByteArr, err := stub.GetState(CommentLogId)
		if err != nil{
			log.Error(systematic.ERRNMCC0004)
			return err, systematic.ERRNMCC0004
		}else if tmpCommentLogByteArr == nil{
			//说明没有对应的数据信息
			var customError = errors.New(systematic.ERRNMCC0012)
			log.Error(systematic.ERRNMCC0012)
			return customError, systematic.ERRNMCC0012
		}
	
		err = stub.DelState(CommentLogId)
		if err != nil{
			log.Error(systematic.ERRNMCC0004)
			return err, systematic.ERRNMCC0004
		}else {
			log.Info("Function Name:DelCommentLog ---- Message:DelState SUCCESS!!!")
		}
	
		return nil, ""
	}
	func UpdateCommentLog(stub shim.ChaincodeStubInterface, CommentLogJson string) (error, string) {
		var CommentLogObj *CommentLog
	
		err := json.Unmarshal([]byte(CommentLogJson), &CommentLogObj)
		if err != nil{
			log.Error(systematic.ERRNMCC0001)
			return err, systematic.ERRNMCC0001
		}
		
		var CommentLogKey = CommentLogObj.Uuid//获取主键id
		tmpCommentLogByteArr, err := stub.GetState(CommentLogKey)
		if err != nil{
			log.Error(systematic.ERRNMCC0006)
			return err, systematic.ERRNMCC0006
		}else if tmpCommentLogByteArr == nil{
			var customError = errors.New(systematic.ERRNMCC0015)
			log.Error(systematic.ERRNMCC0015)
			return customError, systematic.ERRNMCC0015
		}
	
		log.Info("Function Name:UpdateCommentLog ---- Message:PutState execute...")
		err = stub.PutState(CommentLogKey, []byte(CommentLogJson))
		if err != nil{
			log.Error(systematic.ERRNMCC0011)
			return err, systematic.ERRNMCC0011
		}else {
			log.Info("Function Name:UpdateCommentLog ---- Message:PutState SUCCESS!!!")
		}
	
		return nil, ""
	}
	func SelectCommentLog(stub shim.ChaincodeStubInterface, CommentLogId string) (error, string, string) {
		CommentLogByteArr, err := stub.GetState(CommentLogId)
		if err != nil{
			log.Error(systematic.ERRNMCC0006)
			return err, systematic.ERRNMCC0006, ""
		}else if CommentLogByteArr == nil{
			var customError = errors.New(systematic.ERRNMCC0012)
			log.Error(systematic.ERRNMCC0012)
			return customError, systematic.ERRNMCC0012, ""
		}
	
		return nil, "", string(CommentLogByteArr)
	}
	func SelectAllCommentLog(stub shim.ChaincodeStubInterface) (error, string, []string) {
		var queryString = "{\"selector\": {\"dataType\": \"CommentLog\"}}"
	
		resArr, err, errCode := systematic.ConditionQuery(stub, queryString)
		if err != nil{
			return  err, errCode, nil
		}
	
		return nil, "", resArr
	}
	
	func QueryCommentLog(stub shim.ChaincodeStubInterface, queryString string) (error, string, []string) {
		var pQueryString = fmt.Sprintf("{\"selector\": %s}", queryString)
	
		resArr, err, errCode := systematic.ConditionQuery(stub, pQueryString)
		if err != nil{
			return  err, errCode, nil
		}
	
		return nil, "", resArr
	}

//GENERATE_END
