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

func (client *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (response *DescribeTagKeysResponse, err error) {
	response = CreateDescribeTagKeysResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTagKeysWithChan(request *DescribeTagKeysRequest) (<-chan *DescribeTagKeysResponse, <-chan error) {
	responseChan := make(chan *DescribeTagKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTagKeys(request)
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

func (client *Client) DescribeTagKeysWithCallback(request *DescribeTagKeysRequest, callback func(response *DescribeTagKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTagKeysResponse
		var err error
		defer close(result)
		response, err = client.DescribeTagKeys(request)
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

type DescribeTagKeysRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeTagKeysResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	TagKeys    []string `json:"TagKeys" xml:"TagKeys"`
}

func CreateDescribeTagKeysRequest() (request *DescribeTagKeysRequest) {
	request = &DescribeTagKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTagKeys", "", "")
	return
}

func CreateDescribeTagKeysResponse() (response *DescribeTagKeysResponse) {
	response = &DescribeTagKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
