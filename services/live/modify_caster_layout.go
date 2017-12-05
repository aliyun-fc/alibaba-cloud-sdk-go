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

func (client *Client) ModifyCasterLayout(request *ModifyCasterLayoutRequest) (response *ModifyCasterLayoutResponse, err error) {
	response = CreateModifyCasterLayoutResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyCasterLayoutWithChan(request *ModifyCasterLayoutRequest) (<-chan *ModifyCasterLayoutResponse, <-chan error) {
	responseChan := make(chan *ModifyCasterLayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCasterLayout(request)
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

func (client *Client) ModifyCasterLayoutWithCallback(request *ModifyCasterLayoutRequest, callback func(response *ModifyCasterLayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCasterLayoutResponse
		var err error
		defer close(result)
		response, err = client.ModifyCasterLayout(request)
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

type ModifyCasterLayoutRequest struct {
	*requests.RpcRequest
	LayoutId      string                          `position:"Query" name:"LayoutId"`
	MixList       *[]string                       `position:"Query" name:"MixList"  type:"Repeated"`
	BlendList     *[]string                       `position:"Query" name:"BlendList"  type:"Repeated"`
	AudioLayer    *[]ModifyCasterLayoutAudioLayer `position:"Query" name:"AudioLayer"  type:"Repeated"`
	VideoLayer    *[]ModifyCasterLayoutVideoLayer `position:"Query" name:"VideoLayer"  type:"Repeated"`
	OwnerId       string                          `position:"Query" name:"OwnerId"`
	CasterId      string                          `position:"Query" name:"CasterId"`
	SecurityToken string                          `position:"Query" name:"SecurityToken"`
	Version       string                          `position:"Query" name:"Version"`
}

type ModifyCasterLayoutAudioLayer struct {
	VolumeRate   string `name:"VolumeRate"`
	ValidChannel string `name:"ValidChannel"`
}
type ModifyCasterLayoutVideoLayer struct {
	HeightNormalized   string    `name:"HeightNormalized"`
	WidthNormalized    string    `name:"WidthNormalized"`
	PositionRefer      string    `name:"PositionRefer"`
	PositionNormalized *[]string `name:"PositionNormalized" type:"Repeated"`
}

type ModifyCasterLayoutResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	LayoutId  string `json:"LayoutId" xml:"LayoutId"`
}

func CreateModifyCasterLayoutRequest() (request *ModifyCasterLayoutRequest) {
	request = &ModifyCasterLayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyCasterLayout", "", "")
	return
}

func CreateModifyCasterLayoutResponse() (response *ModifyCasterLayoutResponse) {
	response = &ModifyCasterLayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
