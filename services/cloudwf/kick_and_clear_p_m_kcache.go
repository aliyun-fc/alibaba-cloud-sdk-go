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

func (client *Client) KickAndClearPMKcache(request *KickAndClearPMKcacheRequest) (response *KickAndClearPMKcacheResponse, err error) {
	response = CreateKickAndClearPMKcacheResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) KickAndClearPMKcacheWithChan(request *KickAndClearPMKcacheRequest) (<-chan *KickAndClearPMKcacheResponse, <-chan error) {
	responseChan := make(chan *KickAndClearPMKcacheResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.KickAndClearPMKcache(request)
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

func (client *Client) KickAndClearPMKcacheWithCallback(request *KickAndClearPMKcacheRequest, callback func(response *KickAndClearPMKcacheResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *KickAndClearPMKcacheResponse
		var err error
		defer close(result)
		response, err = client.KickAndClearPMKcache(request)
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

type KickAndClearPMKcacheRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

type KickAndClearPMKcacheResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateKickAndClearPMKcacheRequest() (request *KickAndClearPMKcacheRequest) {
	request = &KickAndClearPMKcacheRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "KickAndClearPMKcache", "", "")
	return
}

func CreateKickAndClearPMKcacheResponse() (response *KickAndClearPMKcacheResponse) {
	response = &KickAndClearPMKcacheResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
