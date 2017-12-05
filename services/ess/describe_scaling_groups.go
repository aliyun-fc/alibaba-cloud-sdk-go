package ess

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

func (client *Client) DescribeScalingGroups(request *DescribeScalingGroupsRequest) (response *DescribeScalingGroupsResponse, err error) {
	response = CreateDescribeScalingGroupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeScalingGroupsWithChan(request *DescribeScalingGroupsRequest) (<-chan *DescribeScalingGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingGroups(request)
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

func (client *Client) DescribeScalingGroupsWithCallback(request *DescribeScalingGroupsRequest, callback func(response *DescribeScalingGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingGroups(request)
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

type DescribeScalingGroupsRequest struct {
	*requests.RpcRequest
	ScalingGroupName5    string `position:"Query" name:"ScalingGroupName.5"`
	ScalingGroupName6    string `position:"Query" name:"ScalingGroupName.6"`
	PageSize             string `position:"Query" name:"PageSize"`
	ScalingGroupName3    string `position:"Query" name:"ScalingGroupName.3"`
	ScalingGroupName4    string `position:"Query" name:"ScalingGroupName.4"`
	ScalingGroupId10     string `position:"Query" name:"ScalingGroupId.10"`
	ScalingGroupName9    string `position:"Query" name:"ScalingGroupName.9"`
	ScalingGroupId13     string `position:"Query" name:"ScalingGroupId.13"`
	ScalingGroupId12     string `position:"Query" name:"ScalingGroupId.12"`
	ScalingGroupName7    string `position:"Query" name:"ScalingGroupName.7"`
	ScalingGroupId15     string `position:"Query" name:"ScalingGroupId.15"`
	ScalingGroupName8    string `position:"Query" name:"ScalingGroupName.8"`
	ScalingGroupId14     string `position:"Query" name:"ScalingGroupId.14"`
	ScalingGroupName1    string `position:"Query" name:"ScalingGroupName.1"`
	ScalingGroupName2    string `position:"Query" name:"ScalingGroupName.2"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName20   string `position:"Query" name:"ScalingGroupName.20"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ScalingGroupId6      string `position:"Query" name:"ScalingGroupId.6"`
	ScalingGroupId5      string `position:"Query" name:"ScalingGroupId.5"`
	ScalingGroupId4      string `position:"Query" name:"ScalingGroupId.4"`
	ScalingGroupId3      string `position:"Query" name:"ScalingGroupId.3"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ScalingGroupId9      string `position:"Query" name:"ScalingGroupId.9"`
	ScalingGroupId8      string `position:"Query" name:"ScalingGroupId.8"`
	ScalingGroupId7      string `position:"Query" name:"ScalingGroupId.7"`
	ScalingGroupName16   string `position:"Query" name:"ScalingGroupName.16"`
	ScalingGroupId2      string `position:"Query" name:"ScalingGroupId.2"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ScalingGroupName17   string `position:"Query" name:"ScalingGroupName.17"`
	ScalingGroupId1      string `position:"Query" name:"ScalingGroupId.1"`
	ScalingGroupName18   string `position:"Query" name:"ScalingGroupName.18"`
	ScalingGroupName19   string `position:"Query" name:"ScalingGroupName.19"`
	ScalingGroupId20     string `position:"Query" name:"ScalingGroupId.20"`
	ScalingGroupName13   string `position:"Query" name:"ScalingGroupName.13"`
	ScalingGroupName12   string `position:"Query" name:"ScalingGroupName.12"`
	ScalingGroupName15   string `position:"Query" name:"ScalingGroupName.15"`
	ScalingGroupName14   string `position:"Query" name:"ScalingGroupName.14"`
	ScalingGroupId19     string `position:"Query" name:"ScalingGroupId.19"`
	ScalingGroupId18     string `position:"Query" name:"ScalingGroupId.18"`
	ScalingGroupName11   string `position:"Query" name:"ScalingGroupName.11"`
	ScalingGroupId17     string `position:"Query" name:"ScalingGroupId.17"`
	ScalingGroupName10   string `position:"Query" name:"ScalingGroupName.10"`
	ScalingGroupId16     string `position:"Query" name:"ScalingGroupId.16"`
}

type DescribeScalingGroupsResponse struct {
	*responses.BaseResponse
	TotalCount    int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int    `json:"PageNumber" xml:"PageNumber"`
	PageSize      int    `json:"PageSize" xml:"PageSize"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ScalingGroups []struct {
		DefaultCooldown              int      `json:"DefaultCooldown" xml:"DefaultCooldown"`
		MaxSize                      int      `json:"MaxSize" xml:"MaxSize"`
		PendingCapacity              int      `json:"PendingCapacity" xml:"PendingCapacity"`
		RemovingCapacity             int      `json:"RemovingCapacity" xml:"RemovingCapacity"`
		ScalingGroupName             string   `json:"ScalingGroupName" xml:"ScalingGroupName"`
		ActiveCapacity               int      `json:"ActiveCapacity" xml:"ActiveCapacity"`
		ActiveScalingConfigurationId string   `json:"ActiveScalingConfigurationId" xml:"ActiveScalingConfigurationId"`
		ScalingGroupId               string   `json:"ScalingGroupId" xml:"ScalingGroupId"`
		RegionId                     string   `json:"RegionId" xml:"RegionId"`
		TotalCapacity                int      `json:"TotalCapacity" xml:"TotalCapacity"`
		MinSize                      int      `json:"MinSize" xml:"MinSize"`
		LifecycleState               string   `json:"LifecycleState" xml:"LifecycleState"`
		CreationTime                 string   `json:"CreationTime" xml:"CreationTime"`
		VSwitchId                    string   `json:"VSwitchId" xml:"VSwitchId"`
		RemovalPolicies              []string `json:"RemovalPolicies" xml:"RemovalPolicies"`
		DBInstanceIds                []string `json:"DBInstanceIds" xml:"DBInstanceIds"`
		LoadBalancerIds              []string `json:"LoadBalancerIds" xml:"LoadBalancerIds"`
	} `json:"ScalingGroups" xml:"ScalingGroups"`
}

func CreateDescribeScalingGroupsRequest() (request *DescribeScalingGroupsRequest) {
	request = &DescribeScalingGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2016-07-22", "DescribeScalingGroups", "", "")
	return
}

func CreateDescribeScalingGroupsResponse() (response *DescribeScalingGroupsResponse) {
	response = &DescribeScalingGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
