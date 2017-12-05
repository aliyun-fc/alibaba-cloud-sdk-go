package emr

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

func (client *Client) QueryLogKey(request *QueryLogKeyRequest) (response *QueryLogKeyResponse, err error) {
	response = CreateQueryLogKeyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryLogKeyWithChan(request *QueryLogKeyRequest) (<-chan *QueryLogKeyResponse, <-chan error) {
	responseChan := make(chan *QueryLogKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLogKey(request)
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

func (client *Client) QueryLogKeyWithCallback(request *QueryLogKeyRequest, callback func(response *QueryLogKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLogKeyResponse
		var err error
		defer close(result)
		response, err = client.QueryLogKey(request)
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

type QueryLogKeyRequest struct {
	*requests.RpcRequest
	KeyBase         string `position:"Query" name:"KeyBase"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ContainerId     string `position:"Query" name:"ContainerId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	LogName         string `position:"Query" name:"LogName"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	JobId           string `position:"Query" name:"JobId"`
}

type QueryLogKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	LogKey    string `json:"LogKey" xml:"LogKey"`
}

func CreateQueryLogKeyRequest() (request *QueryLogKeyRequest) {
	request = &QueryLogKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "QueryLogKey", "", "")
	return
}

func CreateQueryLogKeyResponse() (response *QueryLogKeyResponse) {
	response = &QueryLogKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
