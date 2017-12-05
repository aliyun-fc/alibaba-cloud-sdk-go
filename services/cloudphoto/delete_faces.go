package cloudphoto

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

func (client *Client) DeleteFaces(request *DeleteFacesRequest) (response *DeleteFacesResponse, err error) {
	response = CreateDeleteFacesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteFacesWithChan(request *DeleteFacesRequest) (<-chan *DeleteFacesResponse, <-chan error) {
	responseChan := make(chan *DeleteFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaces(request)
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

func (client *Client) DeleteFacesWithCallback(request *DeleteFacesRequest, callback func(response *DeleteFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFacesResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaces(request)
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

type DeleteFacesRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	StoreName string    `position:"Query" name:"StoreName"`
	FaceId    *[]string `position:"Query" name:"FaceId"  type:"Repeated"`
}

type DeleteFacesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Results   []struct {
		Id      int64  `json:"Id" xml:"Id"`
		Code    string `json:"Code" xml:"Code"`
		Message string `json:"Message" xml:"Message"`
	} `json:"Results" xml:"Results"`
}

func CreateDeleteFacesRequest() (request *DeleteFacesRequest) {
	request = &DeleteFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteFaces", "", "")
	return
}

func CreateDeleteFacesResponse() (response *DeleteFacesResponse) {
	response = &DeleteFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
