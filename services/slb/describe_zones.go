package slb

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

func (client *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = CreateDescribeZonesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeZonesWithChan(request *DescribeZonesRequest) (<-chan *DescribeZonesResponse, <-chan error) {
	responseChan := make(chan *DescribeZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeZones(request)
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

func (client *Client) DescribeZonesWithCallback(request *DescribeZonesRequest, callback func(response *DescribeZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeZonesResponse
		var err error
		defer close(result)
		response, err = client.DescribeZones(request)
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

type DescribeZonesRequest struct {
	*requests.RpcRequest
	Tags                 string `position:"Query" name:"Tags"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeZonesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Zones     []struct {
		ZoneId     string `json:"ZoneId" xml:"ZoneId"`
		LocalName  string `json:"LocalName" xml:"LocalName"`
		SlaveZones []struct {
			ZoneId    string `json:"ZoneId" xml:"ZoneId"`
			LocalName string `json:"LocalName" xml:"LocalName"`
		} `json:"SlaveZones" xml:"SlaveZones"`
	} `json:"Zones" xml:"Zones"`
}

func CreateDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeZones", "", "")
	return
}

func CreateDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
