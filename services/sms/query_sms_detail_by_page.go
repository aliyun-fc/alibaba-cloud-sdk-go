package sms

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

func (client *Client) QuerySmsDetailByPage(request *QuerySmsDetailByPageRequest) (response *QuerySmsDetailByPageResponse, err error) {
	response = CreateQuerySmsDetailByPageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QuerySmsDetailByPageWithChan(request *QuerySmsDetailByPageRequest) (<-chan *QuerySmsDetailByPageResponse, <-chan error) {
	responseChan := make(chan *QuerySmsDetailByPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySmsDetailByPage(request)
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

func (client *Client) QuerySmsDetailByPageWithCallback(request *QuerySmsDetailByPageRequest, callback func(response *QuerySmsDetailByPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySmsDetailByPageResponse
		var err error
		defer close(result)
		response, err = client.QuerySmsDetailByPage(request)
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

type QuerySmsDetailByPageRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecNum               string `position:"Query" name:"RecNum"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	PageNo               string `position:"Query" name:"PageNo"`
	QueryTime            string `position:"Query" name:"QueryTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Version              string `position:"Query" name:"Version"`
}

type QuerySmsDetailByPageResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Data       []struct {
		ReceiverNum string `json:"ReceiverNum" xml:"ReceiverNum"`
		SmsCode     string `json:"SmsCode" xml:"SmsCode"`
		SmsContent  string `json:"SmsContent" xml:"SmsContent"`
		SmsStatus   int    `json:"SmsStatus" xml:"SmsStatus"`
		ResultCode  string `json:"ResultCode" xml:"ResultCode"`
	} `json:"data" xml:"data"`
}

func CreateQuerySmsDetailByPageRequest() (request *QuerySmsDetailByPageRequest) {
	request = &QuerySmsDetailByPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sms", "2017-02-28", "QuerySmsDetailByPage", "", "")
	return
}

func CreateQuerySmsDetailByPageResponse() (response *QuerySmsDetailByPageResponse) {
	response = &QuerySmsDetailByPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
