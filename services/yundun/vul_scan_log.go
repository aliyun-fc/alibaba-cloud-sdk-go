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

func (client *Client) VulScanLog(request *VulScanLogRequest) (response *VulScanLogResponse, err error) {
	response = CreateVulScanLogResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) VulScanLogWithChan(request *VulScanLogRequest) (<-chan *VulScanLogResponse, <-chan error) {
	responseChan := make(chan *VulScanLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VulScanLog(request)
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

func (client *Client) VulScanLogWithCallback(request *VulScanLogRequest, callback func(response *VulScanLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VulScanLogResponse
		var err error
		defer close(result)
		response, err = client.VulScanLog(request)
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

type VulScanLogRequest struct {
	*requests.RpcRequest
	PageSize   string `position:"Query" name:"PageSize"`
	PageNumber string `position:"Query" name:"PageNumber"`
	VulStatus  string `position:"Query" name:"VulStatus"`
	InstanceId string `position:"Query" name:"InstanceId"`
	JstOwnerId string `position:"Query" name:"JstOwnerId"`
}

type VulScanLogResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	StartTime  string `json:"StartTime" xml:"StartTime"`
	EndTime    string `json:"EndTime" xml:"EndTime"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	LogList    []struct {
		Id           int    `json:"Id" xml:"Id"`
		Type         string `json:"Type" xml:"Type"`
		Url          string `json:"Url" xml:"Url"`
		HelpAddress  string `json:"HelpAddress" xml:"HelpAddress"`
		VulParameter string `json:"VulParameter" xml:"VulParameter"`
		Status       int    `json:"Status" xml:"Status"`
	} `json:"LogList" xml:"LogList"`
}

func CreateVulScanLogRequest() (request *VulScanLogRequest) {
	request = &VulScanLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "VulScanLog", "", "")
	return
}

func CreateVulScanLogResponse() (response *VulScanLogResponse) {
	response = &VulScanLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
