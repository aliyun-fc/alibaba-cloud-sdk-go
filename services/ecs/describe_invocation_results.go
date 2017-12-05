package ecs

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

func (client *Client) DescribeInvocationResults(request *DescribeInvocationResultsRequest) (response *DescribeInvocationResultsResponse, err error) {
	response = CreateDescribeInvocationResultsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInvocationResultsWithChan(request *DescribeInvocationResultsRequest) (<-chan *DescribeInvocationResultsResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocationResults(request)
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

func (client *Client) DescribeInvocationResultsWithCallback(request *DescribeInvocationResultsRequest, callback func(response *DescribeInvocationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationResultsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocationResults(request)
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

type DescribeInvocationResultsRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

type DescribeInvocationResultsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Invocation struct {
		InvokeId    string `json:"InvokeId" xml:"InvokeId"`
		PageSize    int64  `json:"PageSize" xml:"PageSize"`
		PageNumber  int64  `json:"PageNumber" xml:"PageNumber"`
		TotalCount  int64  `json:"TotalCount" xml:"TotalCount"`
		Status      string `json:"Status" xml:"Status"`
		ResultLists []struct {
			InstanceId   string `json:"InstanceId" xml:"InstanceId"`
			FinishedTime string `json:"FinishedTime" xml:"FinishedTime"`
			Output       string `json:"Output" xml:"Output"`
			ExitCode     int64  `json:"ExitCode" xml:"ExitCode"`
		} `json:"ResultLists" xml:"ResultLists"`
	} `json:"Invocation" xml:"Invocation"`
}

func CreateDescribeInvocationResultsRequest() (request *DescribeInvocationResultsRequest) {
	request = &DescribeInvocationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocationResults", "", "")
	return
}

func CreateDescribeInvocationResultsResponse() (response *DescribeInvocationResultsResponse) {
	response = &DescribeInvocationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
