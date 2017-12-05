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

func (client *Client) AuthUpgradeService(request *AuthUpgradeServiceRequest) (response *AuthUpgradeServiceResponse, err error) {
	response = CreateAuthUpgradeServiceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AuthUpgradeServiceWithChan(request *AuthUpgradeServiceRequest) (<-chan *AuthUpgradeServiceResponse, <-chan error) {
	responseChan := make(chan *AuthUpgradeServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthUpgradeService(request)
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

func (client *Client) AuthUpgradeServiceWithCallback(request *AuthUpgradeServiceRequest, callback func(response *AuthUpgradeServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthUpgradeServiceResponse
		var err error
		defer close(result)
		response, err = client.AuthUpgradeService(request)
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

type AuthUpgradeServiceRequest struct {
	*requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	AuthType      string `position:"Query" name:"AuthType"`
	Sid           string `position:"Query" name:"Sid"`
	OsType        string `position:"Query" name:"OsType"`
	AppName       string `position:"Query" name:"AppName"`
	CurVer        string `position:"Query" name:"CurVer"`
	Language      string `position:"Query" name:"Language"`
	Cid           string `position:"Query" name:"Cid"`
	Umid          string `position:"Query" name:"Umid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type AuthUpgradeServiceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
		PolicyDTO struct {
			Open string `json:"Open" xml:"Open"`
		} `json:"PolicyDTO" xml:"PolicyDTO"`
		InstallPackageDTO struct {
			Name         string `json:"Name" xml:"Name"`
			Version      string `json:"Version" xml:"Version"`
			SoftwareName string `json:"SoftwareName" xml:"SoftwareName"`
			Size         string `json:"Size" xml:"Size"`
			Md5          string `json:"Md5" xml:"Md5"`
			UpTime       string `json:"UpTime" xml:"UpTime"`
			DownloadURL  string `json:"DownloadURL" xml:"DownloadURL"`
			InstallType  string `json:"InstallType" xml:"InstallType"`
			Force        int    `json:"Force" xml:"Force"`
			Zip          int    `json:"Zip" xml:"Zip"`
			UpdateInfo   string `json:"UpdateInfo" xml:"UpdateInfo"`
			Strategy     int    `json:"Strategy" xml:"Strategy"`
		} `json:"InstallPackageDTO" xml:"InstallPackageDTO"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateAuthUpgradeServiceRequest() (request *AuthUpgradeServiceRequest) {
	request = &AuthUpgradeServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "AuthUpgradeService", "", "")
	return
}

func CreateAuthUpgradeServiceResponse() (response *AuthUpgradeServiceResponse) {
	response = &AuthUpgradeServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
