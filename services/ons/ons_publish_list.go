package ons

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

func (client *Client) OnsPublishList(request *OnsPublishListRequest) (response *OnsPublishListResponse, err error) {
	response = CreateOnsPublishListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsPublishListWithChan(request *OnsPublishListRequest) (<-chan *OnsPublishListResponse, <-chan error) {
	responseChan := make(chan *OnsPublishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsPublishList(request)
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

func (client *Client) OnsPublishListWithCallback(request *OnsPublishListRequest, callback func(response *OnsPublishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsPublishListResponse
		var err error
		defer close(result)
		response, err = client.OnsPublishList(request)
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

type OnsPublishListRequest struct {
	*requests.RpcRequest
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	PreventCache string `position:"Query" name:"PreventCache"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsPublishListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      []struct {
		Id          int64  `json:"Id" xml:"Id"`
		ChannelId   int    `json:"ChannelId" xml:"ChannelId"`
		ChannelName string `json:"ChannelName" xml:"ChannelName"`
		OnsRegionId string `json:"OnsRegionId" xml:"OnsRegionId"`
		RegionName  string `json:"RegionName" xml:"RegionName"`
		Owner       string `json:"Owner" xml:"Owner"`
		ProducerId  string `json:"ProducerId" xml:"ProducerId"`
		Topic       string `json:"Topic" xml:"Topic"`
		Status      int    `json:"Status" xml:"Status"`
		StatusName  string `json:"StatusName" xml:"StatusName"`
		CreateTime  int64  `json:"CreateTime" xml:"CreateTime"`
		UpdateTime  int64  `json:"UpdateTime" xml:"UpdateTime"`
	} `json:"Data" xml:"Data"`
}

func CreateOnsPublishListRequest() (request *OnsPublishListRequest) {
	request = &OnsPublishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishList", "", "")
	return
}

func CreateOnsPublishListResponse() (response *OnsPublishListResponse) {
	response = &OnsPublishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
