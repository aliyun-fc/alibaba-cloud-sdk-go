package itaas

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

func (client *Client) AddIPSegment(request *AddIPSegmentRequest) (response *AddIPSegmentResponse, err error) {
	response = CreateAddIPSegmentResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddIPSegmentWithChan(request *AddIPSegmentRequest) (<-chan *AddIPSegmentResponse, <-chan error) {
	responseChan := make(chan *AddIPSegmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddIPSegment(request)
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

func (client *Client) AddIPSegmentWithCallback(request *AddIPSegmentRequest, callback func(response *AddIPSegmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddIPSegmentResponse
		var err error
		defer close(result)
		response, err = client.AddIPSegment(request)
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

type AddIPSegmentRequest struct {
	*requests.RpcRequest
	IpSegment string `position:"Query" name:"IpSegment"`
	AuthType  string `position:"Query" name:"AuthType"`
	Memo      string `position:"Query" name:"Memo"`
	SysFrom   string `position:"Query" name:"SysFrom"`
	Operator  string `position:"Query" name:"Operator"`
}

type AddIPSegmentResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   string `json:"ResultData" xml:"ResultData"`
}

func CreateAddIPSegmentRequest() (request *AddIPSegmentRequest) {
	request = &AddIPSegmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "AddIPSegment", "", "")
	return
}

func CreateAddIPSegmentResponse() (response *AddIPSegmentResponse) {
	response = &AddIPSegmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
