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

func (client *Client) ListClusterOperationHost(request *ListClusterOperationHostRequest) (response *ListClusterOperationHostResponse, err error) {
	response = CreateListClusterOperationHostResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListClusterOperationHostWithChan(request *ListClusterOperationHostRequest) (<-chan *ListClusterOperationHostResponse, <-chan error) {
	responseChan := make(chan *ListClusterOperationHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterOperationHost(request)
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

func (client *Client) ListClusterOperationHostWithCallback(request *ListClusterOperationHostRequest, callback func(response *ListClusterOperationHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterOperationHostResponse
		var err error
		defer close(result)
		response, err = client.ListClusterOperationHost(request)
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

type ListClusterOperationHostRequest struct {
	*requests.RpcRequest
	PageSize        string `position:"Query" name:"PageSize"`
	PageNumber      string `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListClusterOperationHostResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	TotalCount               int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber               int    `json:"PageNumber" xml:"PageNumber"`
	PageSize                 int    `json:"PageSize" xml:"PageSize"`
	ClusterOperationHostList []struct {
		HostId     string `json:"HostId" xml:"HostId"`
		HostName   string `json:"HostName" xml:"HostName"`
		Status     string `json:"Status" xml:"Status"`
		Percentage string `json:"Percentage" xml:"Percentage"`
	} `json:"ClusterOperationHostList" xml:"ClusterOperationHostList"`
}

func CreateListClusterOperationHostRequest() (request *ListClusterOperationHostRequest) {
	request = &ListClusterOperationHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationHost", "", "")
	return
}

func CreateListClusterOperationHostResponse() (response *ListClusterOperationHostResponse) {
	response = &ListClusterOperationHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
