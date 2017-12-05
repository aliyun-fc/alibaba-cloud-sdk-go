package aas

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

func (client *Client) UpdatePasswordForAccount(request *UpdatePasswordForAccountRequest) (response *UpdatePasswordForAccountResponse, err error) {
	response = CreateUpdatePasswordForAccountResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdatePasswordForAccountWithChan(request *UpdatePasswordForAccountRequest) (<-chan *UpdatePasswordForAccountResponse, <-chan error) {
	responseChan := make(chan *UpdatePasswordForAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePasswordForAccount(request)
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

func (client *Client) UpdatePasswordForAccountWithCallback(request *UpdatePasswordForAccountRequest, callback func(response *UpdatePasswordForAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePasswordForAccountResponse
		var err error
		defer close(result)
		response, err = client.UpdatePasswordForAccount(request)
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

type UpdatePasswordForAccountRequest struct {
	*requests.RpcRequest
	NewPassword string `position:"Query" name:"NewPassword"`
	PK          string `position:"Query" name:"PK"`
}

type UpdatePasswordForAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PK        string `json:"PK" xml:"PK"`
	Result    string `json:"Result" xml:"Result"`
}

func CreateUpdatePasswordForAccountRequest() (request *UpdatePasswordForAccountRequest) {
	request = &UpdatePasswordForAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Aas", "2015-07-01", "UpdatePasswordForAccount", "", "")
	return
}

func CreateUpdatePasswordForAccountResponse() (response *UpdatePasswordForAccountResponse) {
	response = &UpdatePasswordForAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
