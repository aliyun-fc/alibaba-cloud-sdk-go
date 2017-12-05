package cloudwf

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

func (client *Client) InnerCheckOrder(request *InnerCheckOrderRequest) (response *InnerCheckOrderResponse, err error) {
	response = CreateInnerCheckOrderResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) InnerCheckOrderWithChan(request *InnerCheckOrderRequest) (<-chan *InnerCheckOrderResponse, <-chan error) {
	responseChan := make(chan *InnerCheckOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InnerCheckOrder(request)
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

func (client *Client) InnerCheckOrderWithCallback(request *InnerCheckOrderRequest, callback func(response *InnerCheckOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InnerCheckOrderResponse
		var err error
		defer close(result)
		response, err = client.InnerCheckOrder(request)
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

type InnerCheckOrderRequest struct {
	*requests.RpcRequest
	Data string `position:"Query" name:"data"`
}

type InnerCheckOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"success" xml:"success"`
	Message   string `json:"message" xml:"message"`
	Code      string `json:"code" xml:"code"`
	Data      string `json:"data" xml:"data"`
}

func CreateInnerCheckOrderRequest() (request *InnerCheckOrderRequest) {
	request = &InnerCheckOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "InnerCheckOrder", "", "")
	return
}

func CreateInnerCheckOrderResponse() (response *InnerCheckOrderResponse) {
	response = &InnerCheckOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
