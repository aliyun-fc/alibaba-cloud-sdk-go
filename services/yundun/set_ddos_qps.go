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

func (client *Client) SetDdosQps(request *SetDdosQpsRequest) (response *SetDdosQpsResponse, err error) {
	response = CreateSetDdosQpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetDdosQpsWithChan(request *SetDdosQpsRequest) (<-chan *SetDdosQpsResponse, <-chan error) {
	responseChan := make(chan *SetDdosQpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDdosQps(request)
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

func (client *Client) SetDdosQpsWithCallback(request *SetDdosQpsRequest, callback func(response *SetDdosQpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDdosQpsResponse
		var err error
		defer close(result)
		response, err = client.SetDdosQps(request)
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

type SetDdosQpsRequest struct {
	*requests.RpcRequest
	QpsPosition  string `position:"Query" name:"QpsPosition"`
	Level        string `position:"Query" name:"Level"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

type SetDdosQpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetDdosQpsRequest() (request *SetDdosQpsRequest) {
	request = &SetDdosQpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "SetDdosQps", "", "")
	return
}

func CreateSetDdosQpsResponse() (response *SetDdosQpsResponse) {
	response = &SetDdosQpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
