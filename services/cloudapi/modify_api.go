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

func (client *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
	response = CreateModifyApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyApiWithChan(request *ModifyApiRequest) (<-chan *ModifyApiResponse, <-chan error) {
	responseChan := make(chan *ModifyApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApi(request)
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

func (client *Client) ModifyApiWithCallback(request *ModifyApiRequest, callback func(response *ModifyApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiResponse
		var err error
		defer close(result)
		response, err = client.ModifyApi(request)
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

type ModifyApiRequest struct {
	*requests.RpcRequest
	AuthType             string `position:"Query" name:"AuthType"`
	ServiceParameters    string `position:"Query" name:"ServiceParameters"`
	Visibility           string `position:"Query" name:"Visibility"`
	ApiName              string `position:"Query" name:"ApiName"`
	ErrorCodeSamples     string `position:"Query" name:"ErrorCodeSamples"`
	SystemParameters     string `position:"Query" name:"SystemParameters"`
	RequestParameters    string `position:"Query" name:"RequestParameters"`
	ConstantParameters   string `position:"Query" name:"ConstantParameters"`
	AllowSignatureMethod string `position:"Query" name:"AllowSignatureMethod"`
	FailResultSample     string `position:"Query" name:"FailResultSample"`
	ResultType           string `position:"Query" name:"ResultType"`
	RequestConfig        string `position:"Query" name:"RequestConfig"`
	ResultSample         string `position:"Query" name:"ResultSample"`
	Description          string `position:"Query" name:"Description"`
	ServiceConfig        string `position:"Query" name:"ServiceConfig"`
	OpenIdConnectConfig  string `position:"Query" name:"OpenIdConnectConfig"`
	GroupId              string `position:"Query" name:"GroupId"`
	ResultDescriptions   string `position:"Query" name:"ResultDescriptions"`
	ServiceParametersMap string `position:"Query" name:"ServiceParametersMap"`
	ApiId                string `position:"Query" name:"ApiId"`
}

type ModifyApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyApiRequest() (request *ModifyApiRequest) {
	request = &ModifyApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApi", "", "")
	return
}

func CreateModifyApiResponse() (response *ModifyApiResponse) {
	response = &ModifyApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
