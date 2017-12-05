package alidns

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

func (client *Client) UpdateDomainGroup(request *UpdateDomainGroupRequest) (response *UpdateDomainGroupResponse, err error) {
	response = CreateUpdateDomainGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateDomainGroupWithChan(request *UpdateDomainGroupRequest) (<-chan *UpdateDomainGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateDomainGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDomainGroup(request)
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

func (client *Client) UpdateDomainGroupWithCallback(request *UpdateDomainGroupRequest, callback func(response *UpdateDomainGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDomainGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateDomainGroup(request)
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

type UpdateDomainGroupRequest struct {
	*requests.RpcRequest
	GroupName    string `position:"Query" name:"GroupName"`
	GroupId      string `position:"Query" name:"GroupId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type UpdateDomainGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
	GroupName string `json:"GroupName" xml:"GroupName"`
}

func CreateUpdateDomainGroupRequest() (request *UpdateDomainGroupRequest) {
	request = &UpdateDomainGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainGroup", "", "")
	return
}

func CreateUpdateDomainGroupResponse() (response *UpdateDomainGroupResponse) {
	response = &UpdateDomainGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
