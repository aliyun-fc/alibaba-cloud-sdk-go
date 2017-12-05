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

func (client *Client) DescribeRecycleBin(request *DescribeRecycleBinRequest) (response *DescribeRecycleBinResponse, err error) {
	response = CreateDescribeRecycleBinResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRecycleBinWithChan(request *DescribeRecycleBinRequest) (<-chan *DescribeRecycleBinResponse, <-chan error) {
	responseChan := make(chan *DescribeRecycleBinResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecycleBin(request)
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

func (client *Client) DescribeRecycleBinWithCallback(request *DescribeRecycleBinRequest, callback func(response *DescribeRecycleBinResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecycleBinResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecycleBin(request)
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

type DescribeRecycleBinRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeRecycleBinResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalCount       int    `json:"TotalCount" xml:"TotalCount"`
	RecycleBinModels []struct {
		ResourceId        string `json:"ResourceId" xml:"ResourceId"`
		RegionId          string `json:"RegionId" xml:"RegionId"`
		ResourceType      string `json:"ResourceType" xml:"ResourceType"`
		CreationTime      string `json:"CreationTime" xml:"CreationTime"`
		Status            string `json:"Status" xml:"Status"`
		RelationResources []struct {
			RelationResourceId   string `json:"RelationResourceId" xml:"RelationResourceId"`
			RelationResourceType string `json:"RelationResourceType" xml:"RelationResourceType"`
		} `json:"RelationResources" xml:"RelationResources"`
	} `json:"RecycleBinModels" xml:"RecycleBinModels"`
}

func CreateDescribeRecycleBinRequest() (request *DescribeRecycleBinRequest) {
	request = &DescribeRecycleBinRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecycleBin", "", "")
	return
}

func CreateDescribeRecycleBinResponse() (response *DescribeRecycleBinResponse) {
	response = &DescribeRecycleBinResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
