package vpc

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

func (client *Client) ModifyRouterInterfaceSpec(request *ModifyRouterInterfaceSpecRequest) (response *ModifyRouterInterfaceSpecResponse, err error) {
	response = CreateModifyRouterInterfaceSpecResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyRouterInterfaceSpecWithChan(request *ModifyRouterInterfaceSpecRequest) (<-chan *ModifyRouterInterfaceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyRouterInterfaceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouterInterfaceSpec(request)
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

func (client *Client) ModifyRouterInterfaceSpecWithCallback(request *ModifyRouterInterfaceSpecRequest, callback func(response *ModifyRouterInterfaceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouterInterfaceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouterInterfaceSpec(request)
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

type ModifyRouterInterfaceSpecRequest struct {
	*requests.RpcRequest
	Spec                 string `position:"Query" name:"Spec"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type ModifyRouterInterfaceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Spec      string `json:"Spec" xml:"Spec"`
}

func CreateModifyRouterInterfaceSpecRequest() (request *ModifyRouterInterfaceSpecRequest) {
	request = &ModifyRouterInterfaceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouterInterfaceSpec", "", "")
	return
}

func CreateModifyRouterInterfaceSpecResponse() (response *ModifyRouterInterfaceSpecResponse) {
	response = &ModifyRouterInterfaceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
