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

func (client *Client) DescribeCdnDomainDetail(request *DescribeCdnDomainDetailRequest) (response *DescribeCdnDomainDetailResponse, err error) {
	response = CreateDescribeCdnDomainDetailResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCdnDomainDetailWithChan(request *DescribeCdnDomainDetailRequest) (<-chan *DescribeCdnDomainDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainDetail(request)
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

func (client *Client) DescribeCdnDomainDetailWithCallback(request *DescribeCdnDomainDetailRequest, callback func(response *DescribeCdnDomainDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainDetail(request)
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

type DescribeCdnDomainDetailRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeCdnDomainDetailResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	GetDomainDetailModel struct {
		GmtCreated              string   `json:"GmtCreated" xml:"GmtCreated"`
		GmtModified             string   `json:"GmtModified" xml:"GmtModified"`
		SourceType              string   `json:"SourceType" xml:"SourceType"`
		DomainStatus            string   `json:"DomainStatus" xml:"DomainStatus"`
		SourcePort              int      `json:"SourcePort" xml:"SourcePort"`
		CdnType                 string   `json:"CdnType" xml:"CdnType"`
		Cname                   string   `json:"Cname" xml:"Cname"`
		HttpsCname              string   `json:"HttpsCname" xml:"HttpsCname"`
		DomainName              string   `json:"DomainName" xml:"DomainName"`
		Description             string   `json:"Description" xml:"Description"`
		ServerCertificateStatus string   `json:"ServerCertificateStatus" xml:"ServerCertificateStatus"`
		ServerCertificate       string   `json:"ServerCertificate" xml:"ServerCertificate"`
		Region                  string   `json:"Region" xml:"Region"`
		Scope                   string   `json:"Scope" xml:"Scope"`
		CertificateName         string   `json:"CertificateName" xml:"CertificateName"`
		ResourceGroupId         string   `json:"ResourceGroupId" xml:"ResourceGroupId"`
		Sources                 []string `json:"Sources" xml:"Sources"`
		SourceModels            []struct {
			Content  string `json:"Content" xml:"Content"`
			Type     string `json:"Type" xml:"Type"`
			Port     int    `json:"Port" xml:"Port"`
			Enabled  string `json:"Enabled" xml:"Enabled"`
			Priority string `json:"Priority" xml:"Priority"`
		} `json:"SourceModels" xml:"SourceModels"`
	} `json:"GetDomainDetailModel" xml:"GetDomainDetailModel"`
}

func CreateDescribeCdnDomainDetailRequest() (request *DescribeCdnDomainDetailRequest) {
	request = &DescribeCdnDomainDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainDetail", "", "")
	return
}

func CreateDescribeCdnDomainDetailResponse() (response *DescribeCdnDomainDetailResponse) {
	response = &DescribeCdnDomainDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
