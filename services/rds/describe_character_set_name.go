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

func (client *Client) DescribeCharacterSetName(request *DescribeCharacterSetNameRequest) (response *DescribeCharacterSetNameResponse, err error) {
	response = CreateDescribeCharacterSetNameResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCharacterSetNameWithChan(request *DescribeCharacterSetNameRequest) (<-chan *DescribeCharacterSetNameResponse, <-chan error) {
	responseChan := make(chan *DescribeCharacterSetNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCharacterSetName(request)
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

func (client *Client) DescribeCharacterSetNameWithCallback(request *DescribeCharacterSetNameRequest, callback func(response *DescribeCharacterSetNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCharacterSetNameResponse
		var err error
		defer close(result)
		response, err = client.DescribeCharacterSetName(request)
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

type DescribeCharacterSetNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeCharacterSetNameResponse struct {
	*responses.BaseResponse
	RequestId             string   `json:"RequestId" xml:"RequestId"`
	Engine                string   `json:"Engine" xml:"Engine"`
	CharacterSetNameItems []string `json:"CharacterSetNameItems" xml:"CharacterSetNameItems"`
}

func CreateDescribeCharacterSetNameRequest() (request *DescribeCharacterSetNameRequest) {
	request = &DescribeCharacterSetNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeCharacterSetName", "", "")
	return
}

func CreateDescribeCharacterSetNameResponse() (response *DescribeCharacterSetNameResponse) {
	response = &DescribeCharacterSetNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
