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

func (client *Client) FrequencyAnalyse(request *FrequencyAnalyseRequest) (response *FrequencyAnalyseResponse, err error) {
	response = CreateFrequencyAnalyseResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) FrequencyAnalyseWithChan(request *FrequencyAnalyseRequest) (<-chan *FrequencyAnalyseResponse, <-chan error) {
	responseChan := make(chan *FrequencyAnalyseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FrequencyAnalyse(request)
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

func (client *Client) FrequencyAnalyseWithCallback(request *FrequencyAnalyseRequest, callback func(response *FrequencyAnalyseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FrequencyAnalyseResponse
		var err error
		defer close(result)
		response, err = client.FrequencyAnalyse(request)
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

type FrequencyAnalyseRequest struct {
	*requests.RpcRequest
	Gsid string `position:"Query" name:"Gsid"`
}

type FrequencyAnalyseResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateFrequencyAnalyseRequest() (request *FrequencyAnalyseRequest) {
	request = &FrequencyAnalyseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "FrequencyAnalyse", "", "")
	return
}

func CreateFrequencyAnalyseResponse() (response *FrequencyAnalyseResponse) {
	response = &FrequencyAnalyseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
