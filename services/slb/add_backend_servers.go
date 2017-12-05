package slb

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

func (client *Client) AddBackendServers(request *AddBackendServersRequest) (response *AddBackendServersResponse, err error) {
	response = CreateAddBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddBackendServersWithChan(request *AddBackendServersRequest) (<-chan *AddBackendServersResponse, <-chan error) {
	responseChan := make(chan *AddBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBackendServers(request)
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

func (client *Client) AddBackendServersWithCallback(request *AddBackendServersRequest, callback func(response *AddBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBackendServersResponse
		var err error
		defer close(result)
		response, err = client.AddBackendServers(request)
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

type AddBackendServersRequest struct {
	*requests.RpcRequest
	Tags                 string `position:"Query" name:"Tags"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
}

type AddBackendServersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	LoadBalancerId string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	BackendServers []struct {
		ServerId string `json:"ServerId" xml:"ServerId"`
		Weight   string `json:"Weight" xml:"Weight"`
	} `json:"BackendServers" xml:"BackendServers"`
}

func CreateAddBackendServersRequest() (request *AddBackendServersRequest) {
	request = &AddBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddBackendServers", "", "")
	return
}

func CreateAddBackendServersResponse() (response *AddBackendServersResponse) {
	response = &AddBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
