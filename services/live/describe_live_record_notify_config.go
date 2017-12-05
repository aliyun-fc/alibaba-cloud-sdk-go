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

func (client *Client) DescribeLiveRecordNotifyConfig(request *DescribeLiveRecordNotifyConfigRequest) (response *DescribeLiveRecordNotifyConfigResponse, err error) {
	response = CreateDescribeLiveRecordNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveRecordNotifyConfigWithChan(request *DescribeLiveRecordNotifyConfigRequest) (<-chan *DescribeLiveRecordNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveRecordNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveRecordNotifyConfig(request)
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

func (client *Client) DescribeLiveRecordNotifyConfigWithCallback(request *DescribeLiveRecordNotifyConfigRequest, callback func(response *DescribeLiveRecordNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveRecordNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveRecordNotifyConfig(request)
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

type DescribeLiveRecordNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveRecordNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	LiveRecordNotifyConfig struct {
		DomainName       string `json:"DomainName" xml:"DomainName"`
		NotifyUrl        string `json:"NotifyUrl" xml:"NotifyUrl"`
		NeedStatusNotify bool   `json:"NeedStatusNotify" xml:"NeedStatusNotify"`
	} `json:"LiveRecordNotifyConfig" xml:"LiveRecordNotifyConfig"`
}

func CreateDescribeLiveRecordNotifyConfigRequest() (request *DescribeLiveRecordNotifyConfigRequest) {
	request = &DescribeLiveRecordNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordNotifyConfig", "", "")
	return
}

func CreateDescribeLiveRecordNotifyConfigResponse() (response *DescribeLiveRecordNotifyConfigResponse) {
	response = &DescribeLiveRecordNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
