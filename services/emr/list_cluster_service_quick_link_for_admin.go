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

func (client *Client) ListClusterServiceQuickLinkForAdmin(request *ListClusterServiceQuickLinkForAdminRequest) (response *ListClusterServiceQuickLinkForAdminResponse, err error) {
	response = CreateListClusterServiceQuickLinkForAdminResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListClusterServiceQuickLinkForAdminWithChan(request *ListClusterServiceQuickLinkForAdminRequest) (<-chan *ListClusterServiceQuickLinkForAdminResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceQuickLinkForAdminResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterServiceQuickLinkForAdmin(request)
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

func (client *Client) ListClusterServiceQuickLinkForAdminWithCallback(request *ListClusterServiceQuickLinkForAdminRequest, callback func(response *ListClusterServiceQuickLinkForAdminResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceQuickLinkForAdminResponse
		var err error
		defer close(result)
		response, err = client.ListClusterServiceQuickLinkForAdmin(request)
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

type ListClusterServiceQuickLinkForAdminRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListClusterServiceQuickLinkForAdminResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	QuickLinkList []struct {
		ServiceName        string `json:"ServiceName" xml:"ServiceName"`
		ServiceDisplayName string `json:"ServiceDisplayName" xml:"ServiceDisplayName"`
		QuickLinkAddress   string `json:"QuickLinkAddress" xml:"QuickLinkAddress"`
		Protocol           string `json:"Protocol" xml:"Protocol"`
	} `json:"QuickLinkList" xml:"QuickLinkList"`
}

func CreateListClusterServiceQuickLinkForAdminRequest() (request *ListClusterServiceQuickLinkForAdminRequest) {
	request = &ListClusterServiceQuickLinkForAdminRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceQuickLinkForAdmin", "", "")
	return
}

func CreateListClusterServiceQuickLinkForAdminResponse() (response *ListClusterServiceQuickLinkForAdminResponse) {
	response = &ListClusterServiceQuickLinkForAdminResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
