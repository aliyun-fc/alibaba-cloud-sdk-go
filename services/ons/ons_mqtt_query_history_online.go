package ons

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

func (client *Client) OnsMqttQueryHistoryOnline(request *OnsMqttQueryHistoryOnlineRequest) (response *OnsMqttQueryHistoryOnlineResponse, err error) {
	response = CreateOnsMqttQueryHistoryOnlineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsMqttQueryHistoryOnlineWithChan(request *OnsMqttQueryHistoryOnlineRequest) (<-chan *OnsMqttQueryHistoryOnlineResponse, <-chan error) {
	responseChan := make(chan *OnsMqttQueryHistoryOnlineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttQueryHistoryOnline(request)
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

func (client *Client) OnsMqttQueryHistoryOnlineWithCallback(request *OnsMqttQueryHistoryOnlineRequest, callback func(response *OnsMqttQueryHistoryOnlineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttQueryHistoryOnlineResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttQueryHistoryOnline(request)
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

type OnsMqttQueryHistoryOnlineRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	BeginTime    string `position:"Query" name:"BeginTime"`
	PreventCache string `position:"Query" name:"PreventCache"`
	GroupId      string `position:"Query" name:"GroupId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsMqttQueryHistoryOnlineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      struct {
		Title   string `json:"Title" xml:"Title"`
		XUnit   string `json:"XUnit" xml:"XUnit"`
		YUnit   string `json:"YUnit" xml:"YUnit"`
		Records []struct {
			X int64   `json:"X" xml:"X"`
			Y float64 `json:"Y" xml:"Y"`
		} `json:"Records" xml:"Records"`
	} `json:"Data" xml:"Data"`
}

func CreateOnsMqttQueryHistoryOnlineRequest() (request *OnsMqttQueryHistoryOnlineRequest) {
	request = &OnsMqttQueryHistoryOnlineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryHistoryOnline", "", "")
	return
}

func CreateOnsMqttQueryHistoryOnlineResponse() (response *OnsMqttQueryHistoryOnlineResponse) {
	response = &OnsMqttQueryHistoryOnlineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
