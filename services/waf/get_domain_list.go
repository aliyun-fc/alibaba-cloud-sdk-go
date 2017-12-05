package waf

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

func (client *Client) GetDomainList(request *GetDomainListRequest) (response *GetDomainListResponse, err error) {
	response = CreateGetDomainListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetDomainListWithChan(request *GetDomainListRequest) (<-chan *GetDomainListResponse, <-chan error) {
	responseChan := make(chan *GetDomainListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDomainList(request)
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

func (client *Client) GetDomainListWithCallback(request *GetDomainListRequest, callback func(response *GetDomainListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDomainListResponse
		var err error
		defer close(result)
		response, err = client.GetDomainList(request)
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

type GetDomainListRequest struct {
	*requests.RpcRequest
	Region      string `position:"Query" name:"Region"`
	PageSize    string `position:"Query" name:"PageSize"`
	Domain      string `position:"Query" name:"Domain"`
	CurrentPage string `position:"Query" name:"CurrentPage"`
}

type GetDomainListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      struct {
		PageInfo struct {
			CurrentPage int `json:"CurrentPage" xml:"CurrentPage"`
			PageSize    int `json:"PageSize" xml:"PageSize"`
			Total       int `json:"Total" xml:"Total"`
		} `json:"PageInfo" xml:"PageInfo"`
		List []struct {
			IsNonStandardPort string   `json:"IsNonStandardPort" xml:"IsNonStandardPort"`
			IsCcActive        int      `json:"IsCcActive" xml:"IsCcActive"`
			IsWafActive       int      `json:"IsWafActive" xml:"IsWafActive"`
			IsAclActive       int      `json:"IsAclActive" xml:"IsAclActive"`
			IsAccessProduct   string   `json:"IsAccessProduct" xml:"IsAccessProduct"`
			Domain            string   `json:"Domain" xml:"Domain"`
			Cname             string   `json:"Cname" xml:"Cname"`
			Protocols         []string `json:"Protocols" xml:"Protocols"`
			SourceIps         []string `json:"SourceIps" xml:"SourceIps"`
		} `json:"List" xml:"List"`
	} `json:"Data" xml:"Data"`
}

func CreateGetDomainListRequest() (request *GetDomainListRequest) {
	request = &GetDomainListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf", "2017-09-28", "GetDomainList", "", "")
	return
}

func CreateGetDomainListResponse() (response *GetDomainListResponse) {
	response = &GetDomainListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
