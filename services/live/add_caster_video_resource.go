package live

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

func (client *Client) AddCasterVideoResource(request *AddCasterVideoResourceRequest) (response *AddCasterVideoResourceResponse, err error) {
	response = CreateAddCasterVideoResourceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddCasterVideoResourceWithChan(request *AddCasterVideoResourceRequest) (<-chan *AddCasterVideoResourceResponse, <-chan error) {
	responseChan := make(chan *AddCasterVideoResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterVideoResource(request)
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

func (client *Client) AddCasterVideoResourceWithCallback(request *AddCasterVideoResourceRequest, callback func(response *AddCasterVideoResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterVideoResourceResponse
		var err error
		defer close(result)
		response, err = client.AddCasterVideoResource(request)
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

type AddCasterVideoResourceRequest struct {
	*requests.RpcRequest
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	LocationId    string `position:"Query" name:"LocationId"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     string `position:"Query" name:"RepeatNum"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	CasterId      string `position:"Query" name:"CasterId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	MaterialId    string `position:"Query" name:"MaterialId"`
	Version       string `position:"Query" name:"Version"`
}

type AddCasterVideoResourceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ResourceId string `json:"ResourceId" xml:"ResourceId"`
}

func CreateAddCasterVideoResourceRequest() (request *AddCasterVideoResourceRequest) {
	request = &AddCasterVideoResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterVideoResource", "", "")
	return
}

func CreateAddCasterVideoResourceResponse() (response *AddCasterVideoResourceResponse) {
	response = &AddCasterVideoResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
