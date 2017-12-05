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

func (client *Client) QueryClusterOrders(request *QueryClusterOrdersRequest) (response *QueryClusterOrdersResponse, err error) {
	response = CreateQueryClusterOrdersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryClusterOrdersWithChan(request *QueryClusterOrdersRequest) (<-chan *QueryClusterOrdersResponse, <-chan error) {
	responseChan := make(chan *QueryClusterOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryClusterOrders(request)
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

func (client *Client) QueryClusterOrdersWithCallback(request *QueryClusterOrdersRequest, callback func(response *QueryClusterOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryClusterOrdersResponse
		var err error
		defer close(result)
		response, err = client.QueryClusterOrders(request)
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

type QueryClusterOrdersRequest struct {
	*requests.RpcRequest
	Id              string `position:"Query" name:"Id"`
	Status          string `position:"Query" name:"Status"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type QueryClusterOrdersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderList []struct {
		OrderId         string `json:"OrderId" xml:"OrderId"`
		CreateTimestamp int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	} `json:"OrderList" xml:"OrderList"`
}

func CreateQueryClusterOrdersRequest() (request *QueryClusterOrdersRequest) {
	request = &QueryClusterOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "QueryClusterOrders", "", "")
	return
}

func CreateQueryClusterOrdersResponse() (response *QueryClusterOrdersResponse) {
	response = &QueryClusterOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
