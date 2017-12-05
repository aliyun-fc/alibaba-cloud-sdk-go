package alidns

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

func (client *Client) DescribeDomainRecords(request *DescribeDomainRecordsRequest) (response *DescribeDomainRecordsResponse, err error) {
	response = CreateDescribeDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainRecordsWithChan(request *DescribeDomainRecordsRequest) (<-chan *DescribeDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRecords(request)
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

func (client *Client) DescribeDomainRecordsWithCallback(request *DescribeDomainRecordsRequest, callback func(response *DescribeDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRecords(request)
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

type DescribeDomainRecordsRequest struct {
	*requests.RpcRequest
	PageSize     string `position:"Query" name:"PageSize"`
	DomainName   string `position:"Query" name:"DomainName"`
	ValueKeyWord string `position:"Query" name:"ValueKeyWord"`
	PageNumber   string `position:"Query" name:"PageNumber"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	TypeKeyWord  string `position:"Query" name:"TypeKeyWord"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RRKeyWord    string `position:"Query" name:"RRKeyWord"`
}

type DescribeDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TotalCount    int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize      int64  `json:"PageSize" xml:"PageSize"`
	DomainRecords []struct {
		DomainName string `json:"DomainName" xml:"DomainName"`
		RecordId   string `json:"RecordId" xml:"RecordId"`
		RR         string `json:"RR" xml:"RR"`
		Type       string `json:"Type" xml:"Type"`
		Value      string `json:"Value" xml:"Value"`
		TTL        int64  `json:"TTL" xml:"TTL"`
		Priority   int64  `json:"Priority" xml:"Priority"`
		Line       string `json:"Line" xml:"Line"`
		Status     string `json:"Status" xml:"Status"`
		Locked     bool   `json:"Locked" xml:"Locked"`
		Weight     int    `json:"Weight" xml:"Weight"`
		Remark     string `json:"Remark" xml:"Remark"`
	} `json:"DomainRecords" xml:"DomainRecords"`
}

func CreateDescribeDomainRecordsRequest() (request *DescribeDomainRecordsRequest) {
	request = &DescribeDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainRecords", "", "")
	return
}

func CreateDescribeDomainRecordsResponse() (response *DescribeDomainRecordsResponse) {
	response = &DescribeDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
