package dm

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

func (client *Client) CheckReplyToMailAddress(request *CheckReplyToMailAddressRequest) (response *CheckReplyToMailAddressResponse, err error) {
	response = CreateCheckReplyToMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CheckReplyToMailAddressWithChan(request *CheckReplyToMailAddressRequest) (<-chan *CheckReplyToMailAddressResponse, <-chan error) {
	responseChan := make(chan *CheckReplyToMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckReplyToMailAddress(request)
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

func (client *Client) CheckReplyToMailAddressWithCallback(request *CheckReplyToMailAddressRequest, callback func(response *CheckReplyToMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckReplyToMailAddressResponse
		var err error
		defer close(result)
		response, err = client.CheckReplyToMailAddress(request)
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

type CheckReplyToMailAddressRequest struct {
	*requests.RpcRequest
	Region               string `position:"Query" name:"Region"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MailAddressId        string `position:"Query" name:"MailAddressId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Lang                 string `position:"Query" name:"Lang"`
}

type CheckReplyToMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCheckReplyToMailAddressRequest() (request *CheckReplyToMailAddressRequest) {
	request = &CheckReplyToMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2017-06-22", "CheckReplyToMailAddress", "", "")
	return
}

func CreateCheckReplyToMailAddressResponse() (response *CheckReplyToMailAddressResponse) {
	response = &CheckReplyToMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
