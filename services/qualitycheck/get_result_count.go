package qualitycheck

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

func (client *Client) GetResultCount(request *GetResultCountRequest) (response *GetResultCountResponse, err error) {
	response = CreateGetResultCountResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetResultCountWithChan(request *GetResultCountRequest) (<-chan *GetResultCountResponse, <-chan error) {
	responseChan := make(chan *GetResultCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResultCount(request)
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

func (client *Client) GetResultCountWithCallback(request *GetResultCountRequest, callback func(response *GetResultCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResultCountResponse
		var err error
		defer close(result)
		response, err = client.GetResultCount(request)
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

type GetResultCountRequest struct {
	*requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

type GetResultCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Success   bool   `json:"success" xml:"success"`
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	Data      int    `json:"data" xml:"data"`
}

func CreateGetResultCountRequest() (request *GetResultCountRequest) {
	request = &GetResultCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetResultCount", "", "")
	return
}

func CreateGetResultCountResponse() (response *GetResultCountResponse) {
	response = &GetResultCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
