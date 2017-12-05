package cr

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

func (client *Client) GetRepoBuildLogs(request *GetRepoBuildLogsRequest) (response *GetRepoBuildLogsResponse, err error) {
	response = CreateGetRepoBuildLogsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetRepoBuildLogsWithChan(request *GetRepoBuildLogsRequest) (<-chan *GetRepoBuildLogsResponse, <-chan error) {
	responseChan := make(chan *GetRepoBuildLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoBuildLogs(request)
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

func (client *Client) GetRepoBuildLogsWithCallback(request *GetRepoBuildLogsRequest, callback func(response *GetRepoBuildLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoBuildLogsResponse
		var err error
		defer close(result)
		response, err = client.GetRepoBuildLogs(request)
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

type GetRepoBuildLogsRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	BuildId       string `position:"Path" name:"BuildId"`
	RepoName      string `position:"Path" name:"RepoName"`
}

type GetRepoBuildLogsResponse struct {
	*responses.BaseResponse
}

func CreateGetRepoBuildLogsRequest() (request *GetRepoBuildLogsRequest) {
	request = &GetRepoBuildLogsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildLogs", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/logs", "", "")
	return
}

func CreateGetRepoBuildLogsResponse() (response *GetRepoBuildLogsResponse) {
	response = &GetRepoBuildLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
