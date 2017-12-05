package cms

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

func (client *Client) ListAlarmHistory(request *ListAlarmHistoryRequest) (response *ListAlarmHistoryResponse, err error) {
	response = CreateListAlarmHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListAlarmHistoryWithChan(request *ListAlarmHistoryRequest) (<-chan *ListAlarmHistoryResponse, <-chan error) {
	responseChan := make(chan *ListAlarmHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlarmHistory(request)
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

func (client *Client) ListAlarmHistoryWithCallback(request *ListAlarmHistoryRequest, callback func(response *ListAlarmHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlarmHistoryResponse
		var err error
		defer close(result)
		response, err = client.ListAlarmHistory(request)
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

type ListAlarmHistoryRequest struct {
	*requests.RpcRequest
	EndTime        string `position:"Query" name:"EndTime"`
	Id             string `position:"Query" name:"Id"`
	Cursor         string `position:"Query" name:"Cursor"`
	StartTime      string `position:"Query" name:"StartTime"`
	Size           string `position:"Query" name:"Size"`
	CallbyCmsOwner string `position:"Query" name:"callby_cms_owner"`
}

type ListAlarmHistoryResponse struct {
	*responses.BaseResponse
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
	Cursor           string `json:"Cursor" xml:"Cursor"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	AlarmHistoryList []struct {
		Id              string `json:"Id" xml:"Id"`
		Name            string `json:"Name" xml:"Name"`
		Namespace       string `json:"Namespace" xml:"Namespace"`
		MetricName      string `json:"MetricName" xml:"MetricName"`
		Dimension       string `json:"Dimension" xml:"Dimension"`
		EvaluationCount int    `json:"EvaluationCount" xml:"EvaluationCount"`
		Value           string `json:"Value" xml:"Value"`
		AlarmTime       int64  `json:"AlarmTime" xml:"AlarmTime"`
		LastTime        int64  `json:"LastTime" xml:"LastTime"`
		State           string `json:"State" xml:"State"`
		Status          int    `json:"Status" xml:"Status"`
		ContactGroups   string `json:"ContactGroups" xml:"ContactGroups"`
	} `json:"AlarmHistoryList" xml:"AlarmHistoryList"`
}

func CreateListAlarmHistoryRequest() (request *ListAlarmHistoryRequest) {
	request = &ListAlarmHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListAlarmHistory", "", "")
	return
}

func CreateListAlarmHistoryResponse() (response *ListAlarmHistoryResponse) {
	response = &ListAlarmHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
