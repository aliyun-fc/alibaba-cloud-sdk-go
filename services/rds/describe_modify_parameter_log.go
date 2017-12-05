package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (response *DescribeModifyParameterLogResponse, err error) {
	response = CreateDescribeModifyParameterLogResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeModifyParameterLogWithChan(request *DescribeModifyParameterLogRequest) (<-chan *DescribeModifyParameterLogResponse, <-chan error) {
	responseChan := make(chan *DescribeModifyParameterLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeModifyParameterLog(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeModifyParameterLogWithCallback(request *DescribeModifyParameterLogRequest, callback func(response *DescribeModifyParameterLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeModifyParameterLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeModifyParameterLog(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeModifyParameterLogRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeModifyParameterLogResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Engine           string `json:"Engine" xml:"Engine"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	EngineVersion    string `json:"EngineVersion" xml:"EngineVersion"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            []struct {
		ModifyTime        string `json:"ModifyTime" xml:"ModifyTime"`
		OldParameterValue string `json:"OldParameterValue" xml:"OldParameterValue"`
		NewParameterValue string `json:"NewParameterValue" xml:"NewParameterValue"`
		ParameterName     string `json:"ParameterName" xml:"ParameterName"`
		Status            string `json:"Status" xml:"Status"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeModifyParameterLogRequest() (request *DescribeModifyParameterLogRequest) {
	request = &DescribeModifyParameterLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeModifyParameterLog", "", "")
	return
}

func CreateDescribeModifyParameterLogResponse() (response *DescribeModifyParameterLogResponse) {
	response = &DescribeModifyParameterLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
