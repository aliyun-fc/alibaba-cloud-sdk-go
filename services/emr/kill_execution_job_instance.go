package emr

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

func (client *Client) KillExecutionJobInstance(request *KillExecutionJobInstanceRequest) (response *KillExecutionJobInstanceResponse, err error) {
	response = CreateKillExecutionJobInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) KillExecutionJobInstanceWithChan(request *KillExecutionJobInstanceRequest) (<-chan *KillExecutionJobInstanceResponse, <-chan error) {
	responseChan := make(chan *KillExecutionJobInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.KillExecutionJobInstance(request)
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

func (client *Client) KillExecutionJobInstanceWithCallback(request *KillExecutionJobInstanceRequest, callback func(response *KillExecutionJobInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *KillExecutionJobInstanceResponse
		var err error
		defer close(result)
		response, err = client.KillExecutionJobInstance(request)
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

type KillExecutionJobInstanceRequest struct {
	*requests.RpcRequest
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type KillExecutionJobInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateKillExecutionJobInstanceRequest() (request *KillExecutionJobInstanceRequest) {
	request = &KillExecutionJobInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "KillExecutionJobInstance", "", "")
	return
}

func CreateKillExecutionJobInstanceResponse() (response *KillExecutionJobInstanceResponse) {
	response = &KillExecutionJobInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
