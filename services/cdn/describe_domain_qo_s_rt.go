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

func (client *Client) DescribeDomainQoSRt(request *DescribeDomainQoSRtRequest) (response *DescribeDomainQoSRtResponse, err error) {
	response = CreateDescribeDomainQoSRtResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainQoSRtWithChan(request *DescribeDomainQoSRtRequest) (<-chan *DescribeDomainQoSRtResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainQoSRtResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainQoSRt(request)
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

func (client *Client) DescribeDomainQoSRtWithCallback(request *DescribeDomainQoSRtRequest, callback func(response *DescribeDomainQoSRtResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainQoSRtResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainQoSRt(request)
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

type DescribeDomainQoSRtRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	Ip            string `position:"Query" name:"Ip"`
	Version       string `position:"Query" name:"Version"`
	DomainName    string `position:"Query" name:"DomainName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeDomainQoSRtResponse struct {
	*responses.BaseResponse
	DomainName string `json:"DomainName" xml:"DomainName"`
	StartTime  string `json:"StartTime" xml:"StartTime"`
	EndTime    string `json:"EndTime" xml:"EndTime"`
	Ip         string `json:"Ip" xml:"Ip"`
	Content    []struct {
		More5s string `json:"More5s" xml:"More5s"`
		Time   string `json:"Time" xml:"Time"`
		More3s string `json:"More3s" xml:"More3s"`
	} `json:"Content" xml:"Content"`
}

func CreateDescribeDomainQoSRtRequest() (request *DescribeDomainQoSRtRequest) {
	request = &DescribeDomainQoSRtRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainQoSRt", "", "")
	return
}

func CreateDescribeDomainQoSRtResponse() (response *DescribeDomainQoSRtResponse) {
	response = &DescribeDomainQoSRtResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
