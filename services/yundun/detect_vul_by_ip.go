package yundun

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

func (client *Client) DetectVulByIp(request *DetectVulByIpRequest) (response *DetectVulByIpResponse, err error) {
	response = CreateDetectVulByIpResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DetectVulByIpWithChan(request *DetectVulByIpRequest) (<-chan *DetectVulByIpResponse, <-chan error) {
	responseChan := make(chan *DetectVulByIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVulByIp(request)
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

func (client *Client) DetectVulByIpWithCallback(request *DetectVulByIpRequest, callback func(response *DetectVulByIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVulByIpResponse
		var err error
		defer close(result)
		response, err = client.DetectVulByIp(request)
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

type DetectVulByIpRequest struct {
	*requests.RpcRequest
	VulIp      string `position:"Query" name:"VulIp"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

type DetectVulByIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDetectVulByIpRequest() (request *DetectVulByIpRequest) {
	request = &DetectVulByIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "DetectVulByIp", "", "")
	return
}

func CreateDetectVulByIpResponse() (response *DetectVulByIpResponse) {
	response = &DetectVulByIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
