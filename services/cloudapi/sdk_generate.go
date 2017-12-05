package cloudapi

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

func (client *Client) SdkGenerate(request *SdkGenerateRequest) (response *SdkGenerateResponse, err error) {
	response = CreateSdkGenerateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SdkGenerateWithChan(request *SdkGenerateRequest) (<-chan *SdkGenerateResponse, <-chan error) {
	responseChan := make(chan *SdkGenerateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SdkGenerate(request)
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

func (client *Client) SdkGenerateWithCallback(request *SdkGenerateRequest, callback func(response *SdkGenerateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SdkGenerateResponse
		var err error
		defer close(result)
		response, err = client.SdkGenerate(request)
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

type SdkGenerateRequest struct {
	*requests.RpcRequest
	AppId    string `position:"Query" name:"AppId"`
	Language string `position:"Query" name:"Language"`
	GroupId  string `position:"Query" name:"GroupId"`
}

type SdkGenerateResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DownloadLink string `json:"DownloadLink" xml:"DownloadLink"`
}

func CreateSdkGenerateRequest() (request *SdkGenerateRequest) {
	request = &SdkGenerateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerate", "", "")
	return
}

func CreateSdkGenerateResponse() (response *SdkGenerateResponse) {
	response = &SdkGenerateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
