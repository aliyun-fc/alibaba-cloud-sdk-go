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

func (client *Client) SoftwareDataCollection(request *SoftwareDataCollectionRequest) (response *SoftwareDataCollectionResponse, err error) {
	response = CreateSoftwareDataCollectionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SoftwareDataCollectionWithChan(request *SoftwareDataCollectionRequest) (<-chan *SoftwareDataCollectionResponse, <-chan error) {
	responseChan := make(chan *SoftwareDataCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SoftwareDataCollection(request)
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

func (client *Client) SoftwareDataCollectionWithCallback(request *SoftwareDataCollectionRequest, callback func(response *SoftwareDataCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SoftwareDataCollectionResponse
		var err error
		defer close(result)
		response, err = client.SoftwareDataCollection(request)
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

type SoftwareDataCollectionRequest struct {
	*requests.RpcRequest
	Uid              string `position:"Query" name:"Uid"`
	SoftwareInfojson string `position:"Query" name:"SoftwareInfojson"`
	Sid              string `position:"Query" name:"Sid"`
	Language         string `position:"Query" name:"Language"`
	Cid              string `position:"Query" name:"Cid"`
	Umid             string `position:"Query" name:"Umid"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
}

type SoftwareDataCollectionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
	} `json:"ResultData" xml:"ResultData"`
}

func CreateSoftwareDataCollectionRequest() (request *SoftwareDataCollectionRequest) {
	request = &SoftwareDataCollectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "SoftwareDataCollection", "", "")
	return
}

func CreateSoftwareDataCollectionResponse() (response *SoftwareDataCollectionResponse) {
	response = &SoftwareDataCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
