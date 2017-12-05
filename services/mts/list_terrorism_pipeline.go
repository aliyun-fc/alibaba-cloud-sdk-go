package mts

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

func (client *Client) ListTerrorismPipeline(request *ListTerrorismPipelineRequest) (response *ListTerrorismPipelineResponse, err error) {
	response = CreateListTerrorismPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListTerrorismPipelineWithChan(request *ListTerrorismPipelineRequest) (<-chan *ListTerrorismPipelineResponse, <-chan error) {
	responseChan := make(chan *ListTerrorismPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTerrorismPipeline(request)
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

func (client *Client) ListTerrorismPipelineWithCallback(request *ListTerrorismPipelineRequest, callback func(response *ListTerrorismPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTerrorismPipelineResponse
		var err error
		defer close(result)
		response, err = client.ListTerrorismPipeline(request)
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

type ListTerrorismPipelineRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	State                string `position:"Query" name:"State"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type ListTerrorismPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TotalCount   int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize     int64  `json:"PageSize" xml:"PageSize"`
	PipelineList []struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		State        string `json:"State" xml:"State"`
		Priority     string `json:"Priority" xml:"Priority"`
		NotifyConfig struct {
			Topic string `json:"Topic" xml:"Topic"`
			Queue string `json:"Queue" xml:"Queue"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"PipelineList" xml:"PipelineList"`
}

func CreateListTerrorismPipelineRequest() (request *ListTerrorismPipelineRequest) {
	request = &ListTerrorismPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListTerrorismPipeline", "", "")
	return
}

func CreateListTerrorismPipelineResponse() (response *ListTerrorismPipelineResponse) {
	response = &ListTerrorismPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
