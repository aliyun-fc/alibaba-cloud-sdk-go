package cloudapi

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

func (client *Client) DescribeApiHistories(request *DescribeApiHistoriesRequest) (response *DescribeApiHistoriesResponse, err error) {
	response = CreateDescribeApiHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeApiHistoriesWithChan(request *DescribeApiHistoriesRequest) (<-chan *DescribeApiHistoriesResponse, <-chan error) {
	responseChan := make(chan *DescribeApiHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiHistories(request)
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

func (client *Client) DescribeApiHistoriesWithCallback(request *DescribeApiHistoriesRequest, callback func(response *DescribeApiHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiHistoriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiHistories(request)
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

type DescribeApiHistoriesRequest struct {
	*requests.RpcRequest
	PageSize   string `position:"Query" name:"PageSize"`
	PageNumber string `position:"Query" name:"PageNumber"`
	StageName  string `position:"Query" name:"StageName"`
	ApiName    string `position:"Query" name:"ApiName"`
	GroupId    string `position:"Query" name:"GroupId"`
	ApiId      string `position:"Query" name:"ApiId"`
}

type DescribeApiHistoriesResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	PageNumber  int    `json:"PageNumber" xml:"PageNumber"`
	ApiHisItems []struct {
		RegionId       string `json:"RegionId" xml:"RegionId"`
		ApiId          string `json:"ApiId" xml:"ApiId"`
		ApiName        string `json:"ApiName" xml:"ApiName"`
		GroupId        string `json:"GroupId" xml:"GroupId"`
		GroupName      string `json:"GroupName" xml:"GroupName"`
		StageName      string `json:"StageName" xml:"StageName"`
		HistoryVersion string `json:"HistoryVersion" xml:"HistoryVersion"`
		Status         string `json:"Status" xml:"Status"`
		Description    string `json:"Description" xml:"Description"`
		DeployedTime   string `json:"DeployedTime" xml:"DeployedTime"`
	} `json:"ApiHisItems" xml:"ApiHisItems"`
}

func CreateDescribeApiHistoriesRequest() (request *DescribeApiHistoriesRequest) {
	request = &DescribeApiHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiHistories", "", "")
	return
}

func CreateDescribeApiHistoriesResponse() (response *DescribeApiHistoriesResponse) {
	response = &DescribeApiHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
