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

func (client *Client) DeleteExecutionPlan(request *DeleteExecutionPlanRequest) (response *DeleteExecutionPlanResponse, err error) {
	response = CreateDeleteExecutionPlanResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteExecutionPlanWithChan(request *DeleteExecutionPlanRequest) (<-chan *DeleteExecutionPlanResponse, <-chan error) {
	responseChan := make(chan *DeleteExecutionPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteExecutionPlan(request)
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

func (client *Client) DeleteExecutionPlanWithCallback(request *DeleteExecutionPlanRequest, callback func(response *DeleteExecutionPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteExecutionPlanResponse
		var err error
		defer close(result)
		response, err = client.DeleteExecutionPlan(request)
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

type DeleteExecutionPlanRequest struct {
	*requests.RpcRequest
	Id              string `position:"Query" name:"Id"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type DeleteExecutionPlanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteExecutionPlanRequest() (request *DeleteExecutionPlanRequest) {
	request = &DeleteExecutionPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteExecutionPlan", "", "")
	return
}

func CreateDeleteExecutionPlanResponse() (response *DeleteExecutionPlanResponse) {
	response = &DeleteExecutionPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
