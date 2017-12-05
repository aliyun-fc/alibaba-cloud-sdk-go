package rds

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

func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (response *DescribeParameterTemplatesResponse, err error) {
	response = CreateDescribeParameterTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeParameterTemplatesWithChan(request *DescribeParameterTemplatesRequest) (<-chan *DescribeParameterTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeParameterTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParameterTemplates(request)
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

func (client *Client) DescribeParameterTemplatesWithCallback(request *DescribeParameterTemplatesRequest, callback func(response *DescribeParameterTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParameterTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeParameterTemplates(request)
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

type DescribeParameterTemplatesRequest struct {
	*requests.RpcRequest
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
}

type DescribeParameterTemplatesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Engine         string `json:"Engine" xml:"Engine"`
	EngineVersion  string `json:"EngineVersion" xml:"EngineVersion"`
	ParameterCount string `json:"ParameterCount" xml:"ParameterCount"`
	Parameters     []struct {
		ParameterName        string `json:"ParameterName" xml:"ParameterName"`
		ParameterValue       string `json:"ParameterValue" xml:"ParameterValue"`
		ForceModify          string `json:"ForceModify" xml:"ForceModify"`
		ForceRestart         string `json:"ForceRestart" xml:"ForceRestart"`
		CheckingCode         string `json:"CheckingCode" xml:"CheckingCode"`
		ParameterDescription string `json:"ParameterDescription" xml:"ParameterDescription"`
	} `json:"Parameters" xml:"Parameters"`
}

func CreateDescribeParameterTemplatesRequest() (request *DescribeParameterTemplatesRequest) {
	request = &DescribeParameterTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeParameterTemplates", "", "")
	return
}

func CreateDescribeParameterTemplatesResponse() (response *DescribeParameterTemplatesResponse) {
	response = &DescribeParameterTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
