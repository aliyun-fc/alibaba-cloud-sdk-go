package ros

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

func (client *Client) DescribeStacks(request *DescribeStacksRequest) (response *DescribeStacksResponse, err error) {
	response = CreateDescribeStacksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeStacksWithChan(request *DescribeStacksRequest) (<-chan *DescribeStacksResponse, <-chan error) {
	responseChan := make(chan *DescribeStacksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStacks(request)
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

func (client *Client) DescribeStacksWithCallback(request *DescribeStacksRequest, callback func(response *DescribeStacksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStacksResponse
		var err error
		defer close(result)
		response, err = client.DescribeStacks(request)
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

type DescribeStacksRequest struct {
	*requests.RoaRequest
	PageSize   string `position:"Query" name:"PageSize"`
	PageNumber string `position:"Query" name:"PageNumber"`
	Status     string `position:"Query" name:"Status"`
	Name       string `position:"Query" name:"Name"`
	StackId    string `position:"Query" name:"StackId"`
}

type DescribeStacksResponse struct {
	*responses.BaseResponse
}

func CreateDescribeStacksRequest() (request *DescribeStacksRequest) {
	request = &DescribeStacksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "DescribeStacks", "/stacks", "", "")
	return
}

func CreateDescribeStacksResponse() (response *DescribeStacksResponse) {
	response = &DescribeStacksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
