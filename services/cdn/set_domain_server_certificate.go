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

func (client *Client) SetDomainServerCertificate(request *SetDomainServerCertificateRequest) (response *SetDomainServerCertificateResponse, err error) {
	response = CreateSetDomainServerCertificateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetDomainServerCertificateWithChan(request *SetDomainServerCertificateRequest) (<-chan *SetDomainServerCertificateResponse, <-chan error) {
	responseChan := make(chan *SetDomainServerCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainServerCertificate(request)
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

func (client *Client) SetDomainServerCertificateWithCallback(request *SetDomainServerCertificateRequest, callback func(response *SetDomainServerCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainServerCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetDomainServerCertificate(request)
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

type SetDomainServerCertificateRequest struct {
	*requests.RpcRequest
	Region                  string `position:"Query" name:"Region"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	CertName                string `position:"Query" name:"CertName"`
	DomainName              string `position:"Query" name:"DomainName"`
	OwnerId                 string `position:"Query" name:"OwnerId"`
	SecurityToken           string `position:"Query" name:"SecurityToken"`
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ServerCertificateStatus string `position:"Query" name:"ServerCertificateStatus"`
}

type SetDomainServerCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetDomainServerCertificateRequest() (request *SetDomainServerCertificateRequest) {
	request = &SetDomainServerCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetDomainServerCertificate", "", "")
	return
}

func CreateSetDomainServerCertificateResponse() (response *SetDomainServerCertificateResponse) {
	response = &SetDomainServerCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
