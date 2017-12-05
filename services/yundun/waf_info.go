package yundun

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

func (client *Client) WafInfo(request *WafInfoRequest) (response *WafInfoResponse, err error) {
	response = CreateWafInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) WafInfoWithChan(request *WafInfoRequest) (<-chan *WafInfoResponse, <-chan error) {
	responseChan := make(chan *WafInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WafInfo(request)
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

func (client *Client) WafInfoWithCallback(request *WafInfoRequest, callback func(response *WafInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WafInfoResponse
		var err error
		defer close(result)
		response, err = client.WafInfo(request)
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

type WafInfoRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

type WafInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	WafDomainNum int    `json:"WafDomainNum" xml:"WafDomainNum"`
	WafInfos     []struct {
		Id     int    `json:"Id" xml:"Id"`
		Domain string `json:"Domain" xml:"Domain"`
		Cname  string `json:"Cname" xml:"Cname"`
		Status int    `json:"Status" xml:"Status"`
	} `json:"WafInfos" xml:"WafInfos"`
}

func CreateWafInfoRequest() (request *WafInfoRequest) {
	request = &WafInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "WafInfo", "", "")
	return
}

func CreateWafInfoResponse() (response *WafInfoResponse) {
	response = &WafInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
