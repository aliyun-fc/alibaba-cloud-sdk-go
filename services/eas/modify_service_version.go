package eas

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

func (client *Client) ModifyServiceVersion(request *ModifyServiceVersionRequest) (response *ModifyServiceVersionResponse, err error) {
	response = CreateModifyServiceVersionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyServiceVersionWithChan(request *ModifyServiceVersionRequest) (<-chan *ModifyServiceVersionResponse, <-chan error) {
	responseChan := make(chan *ModifyServiceVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyServiceVersion(request)
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

func (client *Client) ModifyServiceVersionWithCallback(request *ModifyServiceVersionRequest, callback func(response *ModifyServiceVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyServiceVersionResponse
		var err error
		defer close(result)
		response, err = client.ModifyServiceVersion(request)
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

type ModifyServiceVersionRequest struct {
	*requests.RoaRequest
	Region      string `position:"Path" name:"Region"`
	ServiceName string `position:"Path" name:"ServiceName"`
}

type ModifyServiceVersionResponse struct {
	*responses.BaseResponse
}

func CreateModifyServiceVersionRequest() (request *ModifyServiceVersionRequest) {
	request = &ModifyServiceVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2017-10-23", "ModifyServiceVersion", "/api/services/[Region]/[ServiceName]/version", "", "")
	return
}

func CreateModifyServiceVersionResponse() (response *ModifyServiceVersionResponse) {
	response = &ModifyServiceVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
