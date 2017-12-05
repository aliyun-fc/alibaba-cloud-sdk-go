package qualitycheck

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

func (client *Client) SaveReviewResult(request *SaveReviewResultRequest) (response *SaveReviewResultResponse, err error) {
	response = CreateSaveReviewResultResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveReviewResultWithChan(request *SaveReviewResultRequest) (<-chan *SaveReviewResultResponse, <-chan error) {
	responseChan := make(chan *SaveReviewResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveReviewResult(request)
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

func (client *Client) SaveReviewResultWithCallback(request *SaveReviewResultRequest, callback func(response *SaveReviewResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveReviewResultResponse
		var err error
		defer close(result)
		response, err = client.SaveReviewResult(request)
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

type SaveReviewResultRequest struct {
	*requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

type SaveReviewResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Success   bool   `json:"success" xml:"success"`
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	Data      string `json:"data" xml:"data"`
}

func CreateSaveReviewResultRequest() (request *SaveReviewResultRequest) {
	request = &SaveReviewResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2016-08-01", "SaveReviewResult", "", "")
	return
}

func CreateSaveReviewResultResponse() (response *SaveReviewResultResponse) {
	response = &SaveReviewResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
