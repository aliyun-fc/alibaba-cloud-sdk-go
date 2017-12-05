package live

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

func (client *Client) StopMultipleStreamMixService(request *StopMultipleStreamMixServiceRequest) (response *StopMultipleStreamMixServiceResponse, err error) {
	response = CreateStopMultipleStreamMixServiceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) StopMultipleStreamMixServiceWithChan(request *StopMultipleStreamMixServiceRequest) (<-chan *StopMultipleStreamMixServiceResponse, <-chan error) {
	responseChan := make(chan *StopMultipleStreamMixServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopMultipleStreamMixService(request)
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

func (client *Client) StopMultipleStreamMixServiceWithCallback(request *StopMultipleStreamMixServiceRequest, callback func(response *StopMultipleStreamMixServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopMultipleStreamMixServiceResponse
		var err error
		defer close(result)
		response, err = client.StopMultipleStreamMixService(request)
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

type StopMultipleStreamMixServiceRequest struct {
	*requests.RpcRequest
	StreamName    string `position:"Query" name:"StreamName"`
	DomainName    string `position:"Query" name:"DomainName"`
	AppName       string `position:"Query" name:"AppName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type StopMultipleStreamMixServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateStopMultipleStreamMixServiceRequest() (request *StopMultipleStreamMixServiceRequest) {
	request = &StopMultipleStreamMixServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StopMultipleStreamMixService", "", "")
	return
}

func CreateStopMultipleStreamMixServiceResponse() (response *StopMultipleStreamMixServiceResponse) {
	response = &StopMultipleStreamMixServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
