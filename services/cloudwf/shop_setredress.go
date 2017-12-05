package cloudwf

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

func (client *Client) ShopSetredress(request *ShopSetredressRequest) (response *ShopSetredressResponse, err error) {
	response = CreateShopSetredressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ShopSetredressWithChan(request *ShopSetredressRequest) (<-chan *ShopSetredressResponse, <-chan error) {
	responseChan := make(chan *ShopSetredressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopSetredress(request)
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

func (client *Client) ShopSetredressWithCallback(request *ShopSetredressRequest, callback func(response *ShopSetredressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopSetredressResponse
		var err error
		defer close(result)
		response, err = client.ShopSetredress(request)
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

type ShopSetredressRequest struct {
	*requests.RpcRequest
	Sid         string `position:"Query" name:"Sid"`
	Hnum        string `position:"Query" name:"Hnum"`
	Workday     string `position:"Query" name:"Workday"`
	State       string `position:"Query" name:"State"`
	Filterstate string `position:"Query" name:"Filterstate"`
	Holiday     string `position:"Query" name:"Holiday"`
	Clerk       string `position:"Query" name:"Clerk"`
	Crowdfixed  string `position:"Query" name:"Crowdfixed"`
	Filterclose string `position:"Query" name:"Filterclose"`
	Maxstoptime string `position:"Query" name:"Maxstoptime"`
	Wnum        string `position:"Query" name:"Wnum"`
	Minstoptime string `position:"Query" name:"Minstoptime"`
}

type ShopSetredressResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateShopSetredressRequest() (request *ShopSetredressRequest) {
	request = &ShopSetredressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopSetredress", "", "")
	return
}

func CreateShopSetredressResponse() (response *ShopSetredressResponse) {
	response = &ShopSetredressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
