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

func (client *Client) CopyCasterSceneConfig(request *CopyCasterSceneConfigRequest) (response *CopyCasterSceneConfigResponse, err error) {
	response = CreateCopyCasterSceneConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CopyCasterSceneConfigWithChan(request *CopyCasterSceneConfigRequest) (<-chan *CopyCasterSceneConfigResponse, <-chan error) {
	responseChan := make(chan *CopyCasterSceneConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyCasterSceneConfig(request)
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

func (client *Client) CopyCasterSceneConfigWithCallback(request *CopyCasterSceneConfigRequest, callback func(response *CopyCasterSceneConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyCasterSceneConfigResponse
		var err error
		defer close(result)
		response, err = client.CopyCasterSceneConfig(request)
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

type CopyCasterSceneConfigRequest struct {
	*requests.RpcRequest
	ToSceneId     string `position:"Query" name:"ToSceneId"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	CasterId      string `position:"Query" name:"CasterId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	FromSceneId   string `position:"Query" name:"FromSceneId"`
	Version       string `position:"Query" name:"Version"`
}

type CopyCasterSceneConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCopyCasterSceneConfigRequest() (request *CopyCasterSceneConfigRequest) {
	request = &CopyCasterSceneConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CopyCasterSceneConfig", "", "")
	return
}

func CreateCopyCasterSceneConfigResponse() (response *CopyCasterSceneConfigResponse) {
	response = &CopyCasterSceneConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
