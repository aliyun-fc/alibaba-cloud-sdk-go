package dds

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

func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (response *ModifySecurityIpsResponse, err error) {
	response = CreateModifySecurityIpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifySecurityIpsWithChan(request *ModifySecurityIpsRequest) (<-chan *ModifySecurityIpsResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityIps(request)
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

func (client *Client) ModifySecurityIpsWithCallback(request *ModifySecurityIpsRequest, callback func(response *ModifySecurityIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityIpsResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityIps(request)
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

type ModifySecurityIpsRequest struct {
	*requests.RpcRequest
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	ModifyMode               string `position:"Query" name:"ModifyMode"`
	SecurityIps              string `position:"Query" name:"SecurityIps"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
	SecurityIpGroupAttribute string `position:"Query" name:"SecurityIpGroupAttribute"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId          string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	SecurityIpGroupName      string `position:"Query" name:"SecurityIpGroupName"`
}

type ModifySecurityIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifySecurityIpsRequest() (request *ModifySecurityIpsRequest) {
	request = &ModifySecurityIpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifySecurityIps", "", "")
	return
}

func CreateModifySecurityIpsResponse() (response *ModifySecurityIpsResponse) {
	response = &ModifySecurityIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
