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

func (client *Client) DescribeCdnDomainBaseDetail(request *DescribeCdnDomainBaseDetailRequest) (response *DescribeCdnDomainBaseDetailResponse, err error) {
	response = CreateDescribeCdnDomainBaseDetailResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCdnDomainBaseDetailWithChan(request *DescribeCdnDomainBaseDetailRequest) (<-chan *DescribeCdnDomainBaseDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainBaseDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainBaseDetail(request)
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

func (client *Client) DescribeCdnDomainBaseDetailWithCallback(request *DescribeCdnDomainBaseDetailRequest, callback func(response *DescribeCdnDomainBaseDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainBaseDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainBaseDetail(request)
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

type DescribeCdnDomainBaseDetailRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeCdnDomainBaseDetailResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	DomainBaseDetailModel struct {
		Cname        string   `json:"Cname" xml:"Cname"`
		CdnType      string   `json:"CdnType" xml:"CdnType"`
		DomainStatus string   `json:"DomainStatus" xml:"DomainStatus"`
		SourceType   string   `json:"SourceType" xml:"SourceType"`
		Region       string   `json:"Region" xml:"Region"`
		DomainName   string   `json:"DomainName" xml:"DomainName"`
		Remark       string   `json:"Remark" xml:"Remark"`
		GmtModified  string   `json:"GmtModified" xml:"GmtModified"`
		GmtCreated   string   `json:"GmtCreated" xml:"GmtCreated"`
		Sources      []string `json:"Sources" xml:"Sources"`
	} `json:"DomainBaseDetailModel" xml:"DomainBaseDetailModel"`
}

func CreateDescribeCdnDomainBaseDetailRequest() (request *DescribeCdnDomainBaseDetailRequest) {
	request = &DescribeCdnDomainBaseDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainBaseDetail", "", "")
	return
}

func CreateDescribeCdnDomainBaseDetailResponse() (response *DescribeCdnDomainBaseDetailResponse) {
	response = &DescribeCdnDomainBaseDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
