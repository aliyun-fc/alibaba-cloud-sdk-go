package drds

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

func (client *Client) DescribeCreateDrdsInstanceStatus(request *DescribeCreateDrdsInstanceStatusRequest) (response *DescribeCreateDrdsInstanceStatusResponse, err error) {
	response = CreateDescribeCreateDrdsInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCreateDrdsInstanceStatusWithChan(request *DescribeCreateDrdsInstanceStatusRequest) (<-chan *DescribeCreateDrdsInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeCreateDrdsInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCreateDrdsInstanceStatus(request)
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

func (client *Client) DescribeCreateDrdsInstanceStatusWithCallback(request *DescribeCreateDrdsInstanceStatusRequest, callback func(response *DescribeCreateDrdsInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCreateDrdsInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeCreateDrdsInstanceStatus(request)
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

type DescribeCreateDrdsInstanceStatusRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      struct {
		Status string `json:"Status" xml:"Status"`
	} `json:"Data" xml:"Data"`
}

func CreateDescribeCreateDrdsInstanceStatusRequest() (request *DescribeCreateDrdsInstanceStatusRequest) {
	request = &DescribeCreateDrdsInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DescribeCreateDrdsInstanceStatus", "", "")
	return
}

func CreateDescribeCreateDrdsInstanceStatusResponse() (response *DescribeCreateDrdsInstanceStatusResponse) {
	response = &DescribeCreateDrdsInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
