package itaas

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

func (client *Client) SyncOTP(request *SyncOTPRequest) (response *SyncOTPResponse, err error) {
	response = CreateSyncOTPResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SyncOTPWithChan(request *SyncOTPRequest) (<-chan *SyncOTPResponse, <-chan error) {
	responseChan := make(chan *SyncOTPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncOTP(request)
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

func (client *Client) SyncOTPWithCallback(request *SyncOTPRequest, callback func(response *SyncOTPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncOTPResponse
		var err error
		defer close(result)
		response, err = client.SyncOTP(request)
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

type SyncOTPRequest struct {
	*requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	Sid           string `position:"Query" name:"Sid"`
	AuthType      string `position:"Query" name:"AuthType"`
	OsType        string `position:"Query" name:"OsType"`
	AppVersion    string `position:"Query" name:"AppVersion"`
	Cid           string `position:"Query" name:"Cid"`
	EncryptData   string `position:"Query" name:"EncryptData"`
	AppName       string `position:"Query" name:"AppName"`
	Language      string `position:"Query" name:"Language"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Umid          string `position:"Query" name:"Umid"`
	TokenKey      string `position:"Query" name:"TokenKey"`
}

type SyncOTPResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
		StatusCode int `json:"StatusCode" xml:"StatusCode"`
		ServerTime int `json:"ServerTime" xml:"ServerTime"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateSyncOTPRequest() (request *SyncOTPRequest) {
	request = &SyncOTPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "SyncOTP", "", "")
	return
}

func CreateSyncOTPResponse() (response *SyncOTPResponse) {
	response = &SyncOTPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
