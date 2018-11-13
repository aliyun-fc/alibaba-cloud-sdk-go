package lubancloud

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

// SubmitCutoutTask invokes the lubancloud.SubmitCutoutTask API synchronously
// api document: https://help.aliyun.com/api/lubancloud/submitcutouttask.html
func (client *Client) SubmitCutoutTask(request *SubmitCutoutTaskRequest) (response *SubmitCutoutTaskResponse, err error) {
	response = CreateSubmitCutoutTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCutoutTaskWithChan invokes the lubancloud.SubmitCutoutTask API asynchronously
// api document: https://help.aliyun.com/api/lubancloud/submitcutouttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCutoutTaskWithChan(request *SubmitCutoutTaskRequest) (<-chan *SubmitCutoutTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitCutoutTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCutoutTask(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SubmitCutoutTaskWithCallback invokes the lubancloud.SubmitCutoutTask API asynchronously
// api document: https://help.aliyun.com/api/lubancloud/submitcutouttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCutoutTaskWithCallback(request *SubmitCutoutTaskRequest, callback func(response *SubmitCutoutTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCutoutTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitCutoutTask(request)
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

// SubmitCutoutTaskRequest is the request struct for api SubmitCutoutTask
type SubmitCutoutTaskRequest struct {
	*requests.RpcRequest
	PictureUrl *[]string `position:"Query" name:"PictureUrl"  type:"Repeated"`
}

// SubmitCutoutTaskResponse is the response struct for api SubmitCutoutTask
type SubmitCutoutTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int    `json:"TaskId" xml:"TaskId"`
}

// CreateSubmitCutoutTaskRequest creates a request to invoke SubmitCutoutTask API
func CreateSubmitCutoutTaskRequest() (request *SubmitCutoutTaskRequest) {
	request = &SubmitCutoutTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lubancloud", "2018-05-09", "SubmitCutoutTask", "luban", "openAPI")
	return
}

// CreateSubmitCutoutTaskResponse creates a response to parse from SubmitCutoutTask response
func CreateSubmitCutoutTaskResponse() (response *SubmitCutoutTaskResponse) {
	response = &SubmitCutoutTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}