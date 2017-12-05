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

func (client *Client) GetRegisterBoxList(request *GetRegisterBoxListRequest) (response *GetRegisterBoxListResponse, err error) {
	response = CreateGetRegisterBoxListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetRegisterBoxListWithChan(request *GetRegisterBoxListRequest) (<-chan *GetRegisterBoxListResponse, <-chan error) {
	responseChan := make(chan *GetRegisterBoxListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRegisterBoxList(request)
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

func (client *Client) GetRegisterBoxListWithCallback(request *GetRegisterBoxListRequest, callback func(response *GetRegisterBoxListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRegisterBoxListResponse
		var err error
		defer close(result)
		response, err = client.GetRegisterBoxList(request)
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

type GetRegisterBoxListRequest struct {
	*requests.RpcRequest
	AuthType string `position:"Query" name:"AuthType"`
	SysFrom  string `position:"Query" name:"SysFrom"`
	Operator string `position:"Query" name:"Operator"`
}

type GetRegisterBoxListResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
		ActivedNumber int `json:"ActivedNumber" xml:"ActivedNumber"`
		BuyNumber     int `json:"BuyNumber" xml:"BuyNumber"`
		BoxesList     []struct {
			CurVersion         string `json:"CurVersion" xml:"CurVersion"`
			DrName             string `json:"DrName" xml:"DrName"`
			DrSessionId        string `json:"DrSessionId" xml:"DrSessionId"`
			DrStatus           string `json:"DrStatus" xml:"DrStatus"`
			DrStatusTxt        string `json:"DrStatusTxt" xml:"DrStatusTxt"`
			IpAddress          string `json:"IpAddress" xml:"IpAddress"`
			LastReportTimeL    int64  `json:"LastReportTimeL" xml:"LastReportTimeL"`
			OnlineTimeL        int64  `json:"OnlineTimeL" xml:"OnlineTimeL"`
			ScreenCode         string `json:"ScreenCode" xml:"ScreenCode"`
			SysVersion         string `json:"SysVersion" xml:"SysVersion"`
			AuthUserCode       string `json:"AuthUserCode" xml:"AuthUserCode"`
			AuthRefreshTokenId string `json:"AuthRefreshTokenId" xml:"AuthRefreshTokenId"`
		} `json:"BoxesList" xml:"BoxesList"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateGetRegisterBoxListRequest() (request *GetRegisterBoxListRequest) {
	request = &GetRegisterBoxListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "GetRegisterBoxList", "", "")
	return
}

func CreateGetRegisterBoxListResponse() (response *GetRegisterBoxListResponse) {
	response = &GetRegisterBoxListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
