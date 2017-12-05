package cdn

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

func (client *Client) DescribeLiveStreamDomainAppInfo(request *DescribeLiveStreamDomainAppInfoRequest) (response *DescribeLiveStreamDomainAppInfoResponse, err error) {
	response = CreateDescribeLiveStreamDomainAppInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamDomainAppInfoWithChan(request *DescribeLiveStreamDomainAppInfoRequest) (<-chan *DescribeLiveStreamDomainAppInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamDomainAppInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamDomainAppInfo(request)
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

func (client *Client) DescribeLiveStreamDomainAppInfoWithCallback(request *DescribeLiveStreamDomainAppInfoRequest, callback func(response *DescribeLiveStreamDomainAppInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamDomainAppInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamDomainAppInfo(request)
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

type DescribeLiveStreamDomainAppInfoRequest struct {
	*requests.RpcRequest
	AppDomain     string `position:"Query" name:"AppDomain"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamDomainAppInfoResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	DomainAppList []struct {
		AppDomain    string `json:"AppDomain" xml:"AppDomain"`
		AppId        string `json:"AppId" xml:"AppId"`
		AppKey       string `json:"AppKey" xml:"AppKey"`
		AppOssBucket string `json:"AppOssBucket" xml:"AppOssBucket"`
		AppOssHost   string `json:"AppOssHost" xml:"AppOssHost"`
		AppOwnerId   string `json:"AppOwnerId" xml:"AppOwnerId"`
		AppSecret    string `json:"AppSecret" xml:"AppSecret"`
		UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
	} `json:"DomainAppList" xml:"DomainAppList"`
}

func CreateDescribeLiveStreamDomainAppInfoRequest() (request *DescribeLiveStreamDomainAppInfoRequest) {
	request = &DescribeLiveStreamDomainAppInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamDomainAppInfo", "", "")
	return
}

func CreateDescribeLiveStreamDomainAppInfoResponse() (response *DescribeLiveStreamDomainAppInfoResponse) {
	response = &DescribeLiveStreamDomainAppInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
