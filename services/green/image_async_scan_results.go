package green

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

func (client *Client) ImageAsyncScanResults(request *ImageAsyncScanResultsRequest) (response *ImageAsyncScanResultsResponse, err error) {
	response = CreateImageAsyncScanResultsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImageAsyncScanResultsWithChan(request *ImageAsyncScanResultsRequest) (<-chan *ImageAsyncScanResultsResponse, <-chan error) {
	responseChan := make(chan *ImageAsyncScanResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImageAsyncScanResults(request)
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

func (client *Client) ImageAsyncScanResultsWithCallback(request *ImageAsyncScanResultsRequest, callback func(response *ImageAsyncScanResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImageAsyncScanResultsResponse
		var err error
		defer close(result)
		response, err = client.ImageAsyncScanResults(request)
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

type ImageAsyncScanResultsRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

type ImageAsyncScanResultsResponse struct {
	*responses.BaseResponse
}

func CreateImageAsyncScanResultsRequest() (request *ImageAsyncScanResultsRequest) {
	request = &ImageAsyncScanResultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-25", "ImageAsyncScanResults", "/green/image/results", "", "")
	return
}

func CreateImageAsyncScanResultsResponse() (response *ImageAsyncScanResultsResponse) {
	response = &ImageAsyncScanResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
