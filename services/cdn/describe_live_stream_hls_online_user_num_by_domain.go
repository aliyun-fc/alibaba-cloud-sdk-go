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

func (client *Client) DescribeLiveStreamHlsOnlineUserNumByDomain(request *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) (response *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, err error) {
	response = CreateDescribeLiveStreamHlsOnlineUserNumByDomainResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamHlsOnlineUserNumByDomainWithChan(request *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) (<-chan *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamHlsOnlineUserNumByDomain(request)
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

func (client *Client) DescribeLiveStreamHlsOnlineUserNumByDomainWithCallback(request *DescribeLiveStreamHlsOnlineUserNumByDomainRequest, callback func(response *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamHlsOnlineUserNumByDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamHlsOnlineUserNumByDomain(request)
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

type DescribeLiveStreamHlsOnlineUserNumByDomainRequest struct {
	*requests.RpcRequest
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	PageSize      string `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageNumber    string `position:"Query" name:"PageNumber"`
	AppName       string `position:"Query" name:"AppName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamHlsOnlineUserNumByDomainResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TotalUserNumber int64  `json:"TotalUserNumber" xml:"TotalUserNumber"`
	Count           int64  `json:"Count" xml:"Count"`
	PageNumber      int64  `json:"pageNumber" xml:"pageNumber"`
	PageSize        int64  `json:"pageSize" xml:"pageSize"`
	OnlineUserInfo  []struct {
		StreamUrl  string `json:"StreamUrl" xml:"StreamUrl"`
		UserNumber int64  `json:"UserNumber" xml:"UserNumber"`
		Time       string `json:"Time" xml:"Time"`
	} `json:"OnlineUserInfo" xml:"OnlineUserInfo"`
}

func CreateDescribeLiveStreamHlsOnlineUserNumByDomainRequest() (request *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) {
	request = &DescribeLiveStreamHlsOnlineUserNumByDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamHlsOnlineUserNumByDomain", "", "")
	return
}

func CreateDescribeLiveStreamHlsOnlineUserNumByDomainResponse() (response *DescribeLiveStreamHlsOnlineUserNumByDomainResponse) {
	response = &DescribeLiveStreamHlsOnlineUserNumByDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
