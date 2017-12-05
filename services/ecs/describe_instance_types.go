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

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
	response = CreateDescribeInstanceTypesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceTypesWithChan(request *DescribeInstanceTypesRequest) (<-chan *DescribeInstanceTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTypes(request)
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

func (client *Client) DescribeInstanceTypesWithCallback(request *DescribeInstanceTypesRequest, callback func(response *DescribeInstanceTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTypes(request)
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

type DescribeInstanceTypesRequest struct {
	*requests.RpcRequest
	InstanceTypeFamily   string `position:"Query" name:"InstanceTypeFamily"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeInstanceTypesResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	InstanceTypes []struct {
		InstanceTypeId       string  `json:"InstanceTypeId" xml:"InstanceTypeId"`
		CpuCoreCount         int     `json:"CpuCoreCount" xml:"CpuCoreCount"`
		MemorySize           float64 `json:"MemorySize" xml:"MemorySize"`
		InstanceTypeFamily   string  `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
		LocalStorageCapacity int64   `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
		LocalStorageAmount   int     `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
		LocalStorageCategory string  `json:"LocalStorageCategory" xml:"LocalStorageCategory"`
		GPUAmount            int     `json:"GPUAmount" xml:"GPUAmount"`
		GPUSpec              string  `json:"GPUSpec" xml:"GPUSpec"`
		InitialCredit        int     `json:"InitialCredit" xml:"InitialCredit"`
		BaselineCredit       int     `json:"BaselineCredit" xml:"BaselineCredit"`
	} `json:"InstanceTypes" xml:"InstanceTypes"`
}

func CreateDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
	request = &DescribeInstanceTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceTypes", "", "")
	return
}

func CreateDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
	response = &DescribeInstanceTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
