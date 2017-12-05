package jaq

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

func (client *Client) GetAntiCheatExcelResult(request *GetAntiCheatExcelResultRequest) (response *GetAntiCheatExcelResultResponse, err error) {
	response = CreateGetAntiCheatExcelResultResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetAntiCheatExcelResultWithChan(request *GetAntiCheatExcelResultRequest) (<-chan *GetAntiCheatExcelResultResponse, <-chan error) {
	responseChan := make(chan *GetAntiCheatExcelResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAntiCheatExcelResult(request)
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

func (client *Client) GetAntiCheatExcelResultWithCallback(request *GetAntiCheatExcelResultRequest, callback func(response *GetAntiCheatExcelResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAntiCheatExcelResultResponse
		var err error
		defer close(result)
		response, err = client.GetAntiCheatExcelResult(request)
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

type GetAntiCheatExcelResultRequest struct {
	*requests.RpcRequest
	Packages   *[]string `position:"Query" name:"Packages"  type:"Repeated"`
	CallerName string    `position:"Query" name:"CallerName"`
	StartDay   string    `position:"Query" name:"StartDay"`
	EndDay     string    `position:"Query" name:"EndDay"`
}

type GetAntiCheatExcelResultResponse struct {
	*responses.BaseResponse
	Code int    `json:"Code" xml:"Code"`
	Msg  string `json:"Msg" xml:"Msg"`
	Data []struct {
		Code        int    `json:"Code" xml:"Code"`
		Msg         string `json:"Msg" xml:"Msg"`
		PackageName string `json:"PackageName" xml:"PackageName"`
		DateAndUrl  []struct {
			Date string `json:"Date" xml:"Date"`
			Url  string `json:"Url" xml:"Url"`
		} `json:"DateAndUrl" xml:"DateAndUrl"`
	} `json:"Data" xml:"Data"`
}

func CreateGetAntiCheatExcelResultRequest() (request *GetAntiCheatExcelResultRequest) {
	request = &GetAntiCheatExcelResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jaq", "2017-08-25", "GetAntiCheatExcelResult", "", "")
	return
}

func CreateGetAntiCheatExcelResultResponse() (response *GetAntiCheatExcelResultResponse) {
	response = &GetAntiCheatExcelResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
