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

func (client *Client) DescribeDomainSrcBpsData(request *DescribeDomainSrcBpsDataRequest) (response *DescribeDomainSrcBpsDataResponse, err error) {
	response = CreateDescribeDomainSrcBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainSrcBpsDataWithChan(request *DescribeDomainSrcBpsDataRequest) (<-chan *DescribeDomainSrcBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSrcBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSrcBpsData(request)
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

func (client *Client) DescribeDomainSrcBpsDataWithCallback(request *DescribeDomainSrcBpsDataRequest, callback func(response *DescribeDomainSrcBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSrcBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSrcBpsData(request)
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

type DescribeDomainSrcBpsDataRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	DomainName    string `position:"Query" name:"DomainName"`
	Interval      string `position:"Query" name:"Interval"`
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
}

type DescribeDomainSrcBpsDataResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	DomainName            string `json:"DomainName" xml:"DomainName"`
	DataInterval          string `json:"DataInterval" xml:"DataInterval"`
	StartTime             string `json:"StartTime" xml:"StartTime"`
	EndTime               string `json:"EndTime" xml:"EndTime"`
	SrcBpsDataPerInterval []struct {
		TimeStamp string `json:"TimeStamp" xml:"TimeStamp"`
		Value     string `json:"Value" xml:"Value"`
	} `json:"SrcBpsDataPerInterval" xml:"SrcBpsDataPerInterval"`
}

func CreateDescribeDomainSrcBpsDataRequest() (request *DescribeDomainSrcBpsDataRequest) {
	request = &DescribeDomainSrcBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcBpsData", "", "")
	return
}

func CreateDescribeDomainSrcBpsDataResponse() (response *DescribeDomainSrcBpsDataResponse) {
	response = &DescribeDomainSrcBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
