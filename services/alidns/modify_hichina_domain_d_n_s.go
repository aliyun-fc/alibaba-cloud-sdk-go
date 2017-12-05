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

func (client *Client) ModifyHichinaDomainDNS(request *ModifyHichinaDomainDNSRequest) (response *ModifyHichinaDomainDNSResponse, err error) {
	response = CreateModifyHichinaDomainDNSResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyHichinaDomainDNSWithChan(request *ModifyHichinaDomainDNSRequest) (<-chan *ModifyHichinaDomainDNSResponse, <-chan error) {
	responseChan := make(chan *ModifyHichinaDomainDNSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHichinaDomainDNS(request)
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

func (client *Client) ModifyHichinaDomainDNSWithCallback(request *ModifyHichinaDomainDNSRequest, callback func(response *ModifyHichinaDomainDNSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHichinaDomainDNSResponse
		var err error
		defer close(result)
		response, err = client.ModifyHichinaDomainDNS(request)
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

type ModifyHichinaDomainDNSRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type ModifyHichinaDomainDNSResponse struct {
	*responses.BaseResponse
	RequestId          string   `json:"RequestId" xml:"RequestId"`
	OriginalDnsServers []string `json:"OriginalDnsServers" xml:"OriginalDnsServers"`
	NewDnsServers      []string `json:"NewDnsServers" xml:"NewDnsServers"`
}

func CreateModifyHichinaDomainDNSRequest() (request *ModifyHichinaDomainDNSRequest) {
	request = &ModifyHichinaDomainDNSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "ModifyHichinaDomainDNS", "", "")
	return
}

func CreateModifyHichinaDomainDNSResponse() (response *ModifyHichinaDomainDNSResponse) {
	response = &ModifyHichinaDomainDNSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
