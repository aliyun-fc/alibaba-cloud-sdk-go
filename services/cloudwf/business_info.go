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

func (client *Client) BusinessInfo(request *BusinessInfoRequest) (response *BusinessInfoResponse, err error) {
	response = CreateBusinessInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) BusinessInfoWithChan(request *BusinessInfoRequest) (<-chan *BusinessInfoResponse, <-chan error) {
	responseChan := make(chan *BusinessInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BusinessInfo(request)
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

func (client *Client) BusinessInfoWithCallback(request *BusinessInfoRequest, callback func(response *BusinessInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BusinessInfoResponse
		var err error
		defer close(result)
		response, err = client.BusinessInfo(request)
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

type BusinessInfoRequest struct {
	*requests.RpcRequest
	Bid string `position:"Query" name:"Bid"`
}

type BusinessInfoResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateBusinessInfoRequest() (request *BusinessInfoRequest) {
	request = &BusinessInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessInfo", "", "")
	return
}

func CreateBusinessInfoResponse() (response *BusinessInfoResponse) {
	response = &BusinessInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
