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

func (client *Client) OnsSubscriptionDelete(request *OnsSubscriptionDeleteRequest) (response *OnsSubscriptionDeleteResponse, err error) {
	response = CreateOnsSubscriptionDeleteResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsSubscriptionDeleteWithChan(request *OnsSubscriptionDeleteRequest) (<-chan *OnsSubscriptionDeleteResponse, <-chan error) {
	responseChan := make(chan *OnsSubscriptionDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsSubscriptionDelete(request)
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

func (client *Client) OnsSubscriptionDeleteWithCallback(request *OnsSubscriptionDeleteRequest, callback func(response *OnsSubscriptionDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsSubscriptionDeleteResponse
		var err error
		defer close(result)
		response, err = client.OnsSubscriptionDelete(request)
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

type OnsSubscriptionDeleteRequest struct {
	*requests.RpcRequest
	Topic        string `position:"Query" name:"Topic"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	PreventCache string `position:"Query" name:"PreventCache"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsSubscriptionDeleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

func CreateOnsSubscriptionDeleteRequest() (request *OnsSubscriptionDeleteRequest) {
	request = &OnsSubscriptionDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionDelete", "", "")
	return
}

func CreateOnsSubscriptionDeleteResponse() (response *OnsSubscriptionDeleteResponse) {
	response = &OnsSubscriptionDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
