package cloudapi

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

func (client *Client) RemoveTrafficControlApis(request *RemoveTrafficControlApisRequest) (response *RemoveTrafficControlApisResponse, err error) {
	response = CreateRemoveTrafficControlApisResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RemoveTrafficControlApisWithChan(request *RemoveTrafficControlApisRequest) (<-chan *RemoveTrafficControlApisResponse, <-chan error) {
	responseChan := make(chan *RemoveTrafficControlApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTrafficControlApis(request)
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

func (client *Client) RemoveTrafficControlApisWithCallback(request *RemoveTrafficControlApisRequest, callback func(response *RemoveTrafficControlApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTrafficControlApisResponse
		var err error
		defer close(result)
		response, err = client.RemoveTrafficControlApis(request)
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

type RemoveTrafficControlApisRequest struct {
	*requests.RpcRequest
	ApiIds           string `position:"Query" name:"ApiIds"`
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	StageName        string `position:"Query" name:"StageName"`
	GroupId          string `position:"Query" name:"GroupId"`
}

type RemoveTrafficControlApisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateRemoveTrafficControlApisRequest() (request *RemoveTrafficControlApisRequest) {
	request = &RemoveTrafficControlApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveTrafficControlApis", "", "")
	return
}

func CreateRemoveTrafficControlApisResponse() (response *RemoveTrafficControlApisResponse) {
	response = &RemoveTrafficControlApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
