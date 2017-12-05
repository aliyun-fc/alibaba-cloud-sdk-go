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

func (client *Client) DescribeNewProjectEipMonitorData(request *DescribeNewProjectEipMonitorDataRequest) (response *DescribeNewProjectEipMonitorDataResponse, err error) {
	response = CreateDescribeNewProjectEipMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeNewProjectEipMonitorDataWithChan(request *DescribeNewProjectEipMonitorDataRequest) (<-chan *DescribeNewProjectEipMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeNewProjectEipMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNewProjectEipMonitorData(request)
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

func (client *Client) DescribeNewProjectEipMonitorDataWithCallback(request *DescribeNewProjectEipMonitorDataRequest, callback func(response *DescribeNewProjectEipMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNewProjectEipMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeNewProjectEipMonitorData(request)
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

type DescribeNewProjectEipMonitorDataRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Period               string `position:"Query" name:"Period"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AllocationId         string `position:"Query" name:"AllocationId"`
}

type DescribeNewProjectEipMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	EipMonitorDatas []struct {
		EipRX        int    `json:"EipRX" xml:"EipRX"`
		EipTX        int    `json:"EipTX" xml:"EipTX"`
		EipFlow      int    `json:"EipFlow" xml:"EipFlow"`
		EipBandwidth int    `json:"EipBandwidth" xml:"EipBandwidth"`
		EipPackets   int    `json:"EipPackets" xml:"EipPackets"`
		TimeStamp    string `json:"TimeStamp" xml:"TimeStamp"`
	} `json:"EipMonitorDatas" xml:"EipMonitorDatas"`
}

func CreateDescribeNewProjectEipMonitorDataRequest() (request *DescribeNewProjectEipMonitorDataRequest) {
	request = &DescribeNewProjectEipMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNewProjectEipMonitorData", "", "")
	return
}

func CreateDescribeNewProjectEipMonitorDataResponse() (response *DescribeNewProjectEipMonitorDataResponse) {
	response = &DescribeNewProjectEipMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
