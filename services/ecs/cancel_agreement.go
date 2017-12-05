package ecs

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

func (client *Client) CancelAgreement(request *CancelAgreementRequest) (response *CancelAgreementResponse, err error) {
	response = CreateCancelAgreementResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CancelAgreementWithChan(request *CancelAgreementRequest) (<-chan *CancelAgreementResponse, <-chan error) {
	responseChan := make(chan *CancelAgreementResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelAgreement(request)
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

func (client *Client) CancelAgreementWithCallback(request *CancelAgreementRequest, callback func(response *CancelAgreementResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelAgreementResponse
		var err error
		defer close(result)
		response, err = client.CancelAgreement(request)
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

type CancelAgreementRequest struct {
	*requests.RpcRequest
	AgreementType        string `position:"Query" name:"AgreementType"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

type CancelAgreementResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCancelAgreementRequest() (request *CancelAgreementRequest) {
	request = &CancelAgreementRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CancelAgreement", "", "")
	return
}

func CreateCancelAgreementResponse() (response *CancelAgreementResponse) {
	response = &CancelAgreementResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
