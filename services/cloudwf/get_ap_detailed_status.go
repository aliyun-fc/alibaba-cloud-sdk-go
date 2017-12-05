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

func (client *Client) GetApDetailedStatus(request *GetApDetailedStatusRequest) (response *GetApDetailedStatusResponse, err error) {
	response = CreateGetApDetailedStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetApDetailedStatusWithChan(request *GetApDetailedStatusRequest) (<-chan *GetApDetailedStatusResponse, <-chan error) {
	responseChan := make(chan *GetApDetailedStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApDetailedStatus(request)
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

func (client *Client) GetApDetailedStatusWithCallback(request *GetApDetailedStatusRequest, callback func(response *GetApDetailedStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApDetailedStatusResponse
		var err error
		defer close(result)
		response, err = client.GetApDetailedStatus(request)
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

type GetApDetailedStatusRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

type GetApDetailedStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateGetApDetailedStatusRequest() (request *GetApDetailedStatusRequest) {
	request = &GetApDetailedStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetApDetailedStatus", "", "")
	return
}

func CreateGetApDetailedStatusResponse() (response *GetApDetailedStatusResponse) {
	response = &GetApDetailedStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
