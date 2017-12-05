package sts

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

func (client *Client) GenerateSubUserMiniLoginToken(request *GenerateSubUserMiniLoginTokenRequest) (response *GenerateSubUserMiniLoginTokenResponse, err error) {
	response = CreateGenerateSubUserMiniLoginTokenResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GenerateSubUserMiniLoginTokenWithChan(request *GenerateSubUserMiniLoginTokenRequest) (<-chan *GenerateSubUserMiniLoginTokenResponse, <-chan error) {
	responseChan := make(chan *GenerateSubUserMiniLoginTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateSubUserMiniLoginToken(request)
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

func (client *Client) GenerateSubUserMiniLoginTokenWithCallback(request *GenerateSubUserMiniLoginTokenRequest, callback func(response *GenerateSubUserMiniLoginTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateSubUserMiniLoginTokenResponse
		var err error
		defer close(result)
		response, err = client.GenerateSubUserMiniLoginToken(request)
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

type GenerateSubUserMiniLoginTokenRequest struct {
	*requests.RpcRequest
	TargetSubUserAccountId string `position:"Query" name:"TargetSubUserAccountId"`
}

type GenerateSubUserMiniLoginTokenResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	TargetSubUserAccountId string `json:"TargetSubUserAccountId" xml:"TargetSubUserAccountId"`
	LoginToken             string `json:"LoginToken" xml:"LoginToken"`
}

func CreateGenerateSubUserMiniLoginTokenRequest() (request *GenerateSubUserMiniLoginTokenRequest) {
	request = &GenerateSubUserMiniLoginTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sts", "2017-10-01", "GenerateSubUserMiniLoginToken", "", "")
	return
}

func CreateGenerateSubUserMiniLoginTokenResponse() (response *GenerateSubUserMiniLoginTokenResponse) {
	response = &GenerateSubUserMiniLoginTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
