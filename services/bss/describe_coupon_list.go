package bss

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

func (client *Client) DescribeCouponList(request *DescribeCouponListRequest) (response *DescribeCouponListResponse, err error) {
	response = CreateDescribeCouponListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCouponListWithChan(request *DescribeCouponListRequest) (<-chan *DescribeCouponListResponse, <-chan error) {
	responseChan := make(chan *DescribeCouponListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCouponList(request)
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

func (client *Client) DescribeCouponListWithCallback(request *DescribeCouponListRequest, callback func(response *DescribeCouponListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCouponListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCouponList(request)
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

type DescribeCouponListRequest struct {
	*requests.RpcRequest
	PageSize          string `position:"Query" name:"PageSize"`
	Status            string `position:"Query" name:"Status"`
	PageNum           string `position:"Query" name:"PageNum"`
	StartDeliveryTime string `position:"Query" name:"StartDeliveryTime"`
	EndDeliveryTime   string `position:"Query" name:"EndDeliveryTime"`
}

type DescribeCouponListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Coupons   []struct {
		CouponTemplateId int64    `json:"CouponTemplateId" xml:"CouponTemplateId"`
		TotalAmount      string   `json:"TotalAmount" xml:"TotalAmount"`
		BalanceAmount    string   `json:"BalanceAmount" xml:"BalanceAmount"`
		FrozenAmount     string   `json:"FrozenAmount" xml:"FrozenAmount"`
		ExpiredAmount    string   `json:"ExpiredAmount" xml:"ExpiredAmount"`
		DeliveryTime     string   `json:"DeliveryTime" xml:"DeliveryTime"`
		ExpiredTime      string   `json:"ExpiredTime" xml:"ExpiredTime"`
		CouponNumber     string   `json:"CouponNumber" xml:"CouponNumber"`
		Status           string   `json:"Status" xml:"Status"`
		Description      string   `json:"Description" xml:"Description"`
		CreationTime     string   `json:"CreationTime" xml:"CreationTime"`
		ModificationTime string   `json:"ModificationTime" xml:"ModificationTime"`
		PriceLimit       string   `json:"PriceLimit" xml:"PriceLimit"`
		Application      string   `json:"Application" xml:"Application"`
		ProductCodes     []string `json:"ProductCodes" xml:"ProductCodes"`
		TradeTypes       []string `json:"TradeTypes" xml:"TradeTypes"`
	} `json:"Coupons" xml:"Coupons"`
}

func CreateDescribeCouponListRequest() (request *DescribeCouponListRequest) {
	request = &DescribeCouponListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Bss", "2014-07-14", "DescribeCouponList", "", "")
	return
}

func CreateDescribeCouponListResponse() (response *DescribeCouponListResponse) {
	response = &DescribeCouponListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
