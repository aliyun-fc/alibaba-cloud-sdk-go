package itaas

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

func (client *Client) CheckBuyVideoConference(request *CheckBuyVideoConferenceRequest) (response *CheckBuyVideoConferenceResponse, err error) {
	response = CreateCheckBuyVideoConferenceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CheckBuyVideoConferenceWithChan(request *CheckBuyVideoConferenceRequest) (<-chan *CheckBuyVideoConferenceResponse, <-chan error) {
	responseChan := make(chan *CheckBuyVideoConferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckBuyVideoConference(request)
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

func (client *Client) CheckBuyVideoConferenceWithCallback(request *CheckBuyVideoConferenceRequest, callback func(response *CheckBuyVideoConferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckBuyVideoConferenceResponse
		var err error
		defer close(result)
		response, err = client.CheckBuyVideoConference(request)
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

type CheckBuyVideoConferenceRequest struct {
	*requests.RpcRequest
	AuthType string `position:"Query" name:"AuthType"`
	SysFrom  string `position:"Query" name:"SysFrom"`
	Operator string `position:"Query" name:"Operator"`
}

type CheckBuyVideoConferenceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
		Buy bool `json:"Buy" xml:"Buy"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateCheckBuyVideoConferenceRequest() (request *CheckBuyVideoConferenceRequest) {
	request = &CheckBuyVideoConferenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "CheckBuyVideoConference", "", "")
	return
}

func CreateCheckBuyVideoConferenceResponse() (response *CheckBuyVideoConferenceResponse) {
	response = &CheckBuyVideoConferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
