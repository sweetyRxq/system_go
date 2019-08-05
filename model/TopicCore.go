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
	func AddTopic(stub shim.ChaincodeStubInterface, TopicJson string, fileArr []string) (error, string){
		var TopicObj *Topic
	
		err := json.Unmarshal([]byte(TopicJson), &TopicObj)
		if err != nil{
			log.Error(systematic.ERRNMCC0001)
			return err, systematic.ERRNMCC0001
		}
		
		var TopicKey = TopicObj.Uuid//获取主键id
	
		log.Info("Function Name:AddTopic ---- Message:PutState execute...")
		err = stub.PutState(TopicKey, []byte(TopicJson))
		if err != nil{
			log.Error(systematic.ERRNMCC0005)
			return err, systematic.ERRNMCC0005
		}else {
			log.Info("Function Name:AddTopic ---- Message:PutState SUCCESS!!!")
		}
		return nil, ""
	}
	func DelTopic(stub shim.ChaincodeStubInterface, TopicId string) (error, string) {
		log.Info("Function Name:DelTopic ---- Message:DelState execute...")
		//先取到数据，判断要删除的数据信息在数据库中存不存在
		tmpTopicByteArr, err := stub.GetState(TopicId)
		if err != nil{
			log.Error(systematic.ERRNMCC0004)
			return err, systematic.ERRNMCC0004
		}else if tmpTopicByteArr == nil{
			//说明没有对应的数据信息
			var customError = errors.New(systematic.ERRNMCC0012)
			log.Error(systematic.ERRNMCC0012)
			return customError, systematic.ERRNMCC0012
		}
	
		err = stub.DelState(TopicId)
		if err != nil{
			log.Error(systematic.ERRNMCC0004)
			return err, systematic.ERRNMCC0004
		}else {
			log.Info("Function Name:DelTopic ---- Message:DelState SUCCESS!!!")
		}
	
		return nil, ""
	}
	func UpdateTopic(stub shim.ChaincodeStubInterface, TopicJson string) (error, string) {
		var TopicObj *Topic
	
		err := json.Unmarshal([]byte(TopicJson), &TopicObj)
		if err != nil{
			log.Error(systematic.ERRNMCC0001)
			return err, systematic.ERRNMCC0001
		}
		
		var TopicKey = TopicObj.Uuid//获取主键id
		tmpTopicByteArr, err := stub.GetState(TopicKey)
		if err != nil{
			log.Error(systematic.ERRNMCC0006)
			return err, systematic.ERRNMCC0006
		}else if tmpTopicByteArr == nil{
			var customError = errors.New(systematic.ERRNMCC0015)
			log.Error(systematic.ERRNMCC0015)
			return customError, systematic.ERRNMCC0015
		}
	
		log.Info("Function Name:UpdateTopic ---- Message:PutState execute...")
		err = stub.PutState(TopicKey, []byte(TopicJson))
		if err != nil{
			log.Error(systematic.ERRNMCC0011)
			return err, systematic.ERRNMCC0011
		}else {
			log.Info("Function Name:UpdateTopic ---- Message:PutState SUCCESS!!!")
		}
	
		return nil, ""
	}
	func SelectTopic(stub shim.ChaincodeStubInterface, TopicId string) (error, string, string) {
		TopicByteArr, err := stub.GetState(TopicId)
		if err != nil{
			log.Error(systematic.ERRNMCC0006)
			return err, systematic.ERRNMCC0006, ""
		}else if TopicByteArr == nil{
			var customError = errors.New(systematic.ERRNMCC0012)
			log.Error(systematic.ERRNMCC0012)
			return customError, systematic.ERRNMCC0012, ""
		}
	
		return nil, "", string(TopicByteArr)
	}
	func SelectAllTopic(stub shim.ChaincodeStubInterface) (error, string, []string) {
		var queryString = "{\"selector\": {\"dataType\": \"Topic\"}}"
	
		resArr, err, errCode := systematic.ConditionQuery(stub, queryString)
		if err != nil{
			return  err, errCode, nil
		}
	
		return nil, "", resArr
	}
	
	func QueryTopic(stub shim.ChaincodeStubInterface, queryString string) (error, string, []string) {
		var pQueryString = fmt.Sprintf("{\"selector\": %s}", queryString)
	
		resArr, err, errCode := systematic.ConditionQuery(stub, pQueryString)
		if err != nil{
			return  err, errCode, nil
		}
	
		return nil, "", resArr
	}
//GENERATE_END