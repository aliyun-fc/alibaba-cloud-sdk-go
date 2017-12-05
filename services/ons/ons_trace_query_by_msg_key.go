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

func (client *Client) OnsTraceQueryByMsgKey(request *OnsTraceQueryByMsgKeyRequest) (response *OnsTraceQueryByMsgKeyResponse, err error) {
	response = CreateOnsTraceQueryByMsgKeyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsTraceQueryByMsgKeyWithChan(request *OnsTraceQueryByMsgKeyRequest) (<-chan *OnsTraceQueryByMsgKeyResponse, <-chan error) {
	responseChan := make(chan *OnsTraceQueryByMsgKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTraceQueryByMsgKey(request)
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

func (client *Client) OnsTraceQueryByMsgKeyWithCallback(request *OnsTraceQueryByMsgKeyRequest, callback func(response *OnsTraceQueryByMsgKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTraceQueryByMsgKeyResponse
		var err error
		defer close(result)
		response, err = client.OnsTraceQueryByMsgKey(request)
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

type OnsTraceQueryByMsgKeyRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	Topic        string `position:"Query" name:"Topic"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	BeginTime    string `position:"Query" name:"BeginTime"`
	PreventCache string `position:"Query" name:"PreventCache"`
	MsgKey       string `position:"Query" name:"MsgKey"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsTraceQueryByMsgKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	QueryId   string `json:"QueryId" xml:"QueryId"`
}

func CreateOnsTraceQueryByMsgKeyRequest() (request *OnsTraceQueryByMsgKeyRequest) {
	request = &OnsTraceQueryByMsgKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceQueryByMsgKey", "", "")
	return
}

func CreateOnsTraceQueryByMsgKeyResponse() (response *OnsTraceQueryByMsgKeyResponse) {
	response = &OnsTraceQueryByMsgKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
