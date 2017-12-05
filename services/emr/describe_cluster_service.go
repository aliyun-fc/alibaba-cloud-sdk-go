package emr

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

func (client *Client) DescribeClusterService(request *DescribeClusterServiceRequest) (response *DescribeClusterServiceResponse, err error) {
	response = CreateDescribeClusterServiceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeClusterServiceWithChan(request *DescribeClusterServiceRequest) (<-chan *DescribeClusterServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterService(request)
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

func (client *Client) DescribeClusterServiceWithCallback(request *DescribeClusterServiceRequest, callback func(response *DescribeClusterServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterService(request)
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

type DescribeClusterServiceRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type DescribeClusterServiceResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ServiceInfo struct {
		ServiceName                  string   `json:"ServiceName" xml:"ServiceName"`
		ServiceVersion               string   `json:"ServiceVersion" xml:"ServiceVersion"`
		ServiceStatus                string   `json:"ServiceStatus" xml:"ServiceStatus"`
		NeedRestartInfo              string   `json:"NeedRestartInfo" xml:"NeedRestartInfo"`
		NeedRestartNum               int      `json:"NeedRestartNum" xml:"NeedRestartNum"`
		NeedRestartHostIdList        []string `json:"NeedRestartHostIdList" xml:"NeedRestartHostIdList"`
		NeedRestartComponentNameList []string `json:"NeedRestartComponentNameList" xml:"NeedRestartComponentNameList"`
		ServiceActionList            []struct {
			ServiceName   string `json:"ServiceName" xml:"ServiceName"`
			ComponentName string `json:"ComponentName" xml:"ComponentName"`
			ActionName    string `json:"ActionName" xml:"ActionName"`
			Command       string `json:"Command" xml:"Command"`
			DisplayName   string `json:"DisplayName" xml:"DisplayName"`
		} `json:"ServiceActionList" xml:"ServiceActionList"`
		ClusterServiceSummaryList []struct {
			Key         string `json:"Key" xml:"Key"`
			DisplayName string `json:"DisplayName" xml:"DisplayName"`
			Value       string `json:"Value" xml:"Value"`
			Status      string `json:"Status" xml:"Status"`
			Type        string `json:"Type" xml:"Type"`
			AlertInfo   string `json:"AlertInfo" xml:"AlertInfo"`
		} `json:"ClusterServiceSummaryList" xml:"ClusterServiceSummaryList"`
	} `json:"ServiceInfo" xml:"ServiceInfo"`
}

func CreateDescribeClusterServiceRequest() (request *DescribeClusterServiceRequest) {
	request = &DescribeClusterServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterService", "", "")
	return
}

func CreateDescribeClusterServiceResponse() (response *DescribeClusterServiceResponse) {
	response = &DescribeClusterServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
