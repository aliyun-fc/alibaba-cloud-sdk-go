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

func (client *Client) DescribeReadOnlyAccount(request *DescribeReadOnlyAccountRequest) (response *DescribeReadOnlyAccountResponse, err error) {
	response = CreateDescribeReadOnlyAccountResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeReadOnlyAccountWithChan(request *DescribeReadOnlyAccountRequest) (<-chan *DescribeReadOnlyAccountResponse, <-chan error) {
	responseChan := make(chan *DescribeReadOnlyAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReadOnlyAccount(request)
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

func (client *Client) DescribeReadOnlyAccountWithCallback(request *DescribeReadOnlyAccountRequest, callback func(response *DescribeReadOnlyAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReadOnlyAccountResponse
		var err error
		defer close(result)
		response, err = client.DescribeReadOnlyAccount(request)
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

type DescribeReadOnlyAccountRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

type DescribeReadOnlyAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      struct {
		DbName         string `json:"DbName" xml:"DbName"`
		DrdsInstanceId string `json:"DrdsInstanceId" xml:"DrdsInstanceId"`
		AccountName    string `json:"AccountName" xml:"AccountName"`
	} `json:"Data" xml:"Data"`
}

func CreateDescribeReadOnlyAccountRequest() (request *DescribeReadOnlyAccountRequest) {
	request = &DescribeReadOnlyAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DescribeReadOnlyAccount", "", "")
	return
}

func CreateDescribeReadOnlyAccountResponse() (response *DescribeReadOnlyAccountResponse) {
	response = &DescribeReadOnlyAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
