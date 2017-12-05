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

func (client *Client) ListResourcePool(request *ListResourcePoolRequest) (response *ListResourcePoolResponse, err error) {
	response = CreateListResourcePoolResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListResourcePoolWithChan(request *ListResourcePoolRequest) (<-chan *ListResourcePoolResponse, <-chan error) {
	responseChan := make(chan *ListResourcePoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourcePool(request)
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

func (client *Client) ListResourcePoolWithCallback(request *ListResourcePoolRequest, callback func(response *ListResourcePoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourcePoolResponse
		var err error
		defer close(result)
		response, err = client.ListResourcePool(request)
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

type ListResourcePoolRequest struct {
	*requests.RpcRequest
	PoolType        string `position:"Query" name:"PoolType"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListResourcePoolResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	Total        int    `json:"Total" xml:"Total"`
	PoolInfoList []struct {
		EcmResourcePool struct {
			Id             int64  `json:"Id" xml:"Id"`
			Name           string `json:"Name" xml:"Name"`
			ClusterId      int64  `json:"ClusterId" xml:"ClusterId"`
			PoolType       string `json:"PoolType" xml:"PoolType"`
			Active         bool   `json:"Active" xml:"Active"`
			Note           string `json:"Note" xml:"Note"`
			UserId         string `json:"UserId" xml:"UserId"`
			YarnSiteConfig string `json:"YarnSiteConfig" xml:"YarnSiteConfig"`
		} `json:"EcmResourcePool" xml:"EcmResourcePool"`
		QueueList []struct {
			EcmResourceQueue struct {
				Id             int64  `json:"Id" xml:"Id"`
				Name           string `json:"Name" xml:"Name"`
				QualifiedName  string `json:"QualifiedName" xml:"QualifiedName"`
				ClusterId      int64  `json:"ClusterId" xml:"ClusterId"`
				QueueType      string `json:"QueueType" xml:"QueueType"`
				ParentQueueId  int64  `json:"ParentQueueId" xml:"ParentQueueId"`
				Leaf           bool   `json:"Leaf" xml:"Leaf"`
				Status         string `json:"Status" xml:"Status"`
				UserId         string `json:"UserId" xml:"UserId"`
				ResourcePoolId int64  `json:"ResourcePoolId" xml:"ResourcePoolId"`
			} `json:"EcmResourceQueue" xml:"EcmResourceQueue"`
			EcmResourcePoolConfigList []struct {
				Id          int64  `json:"Id" xml:"Id"`
				ConfigKey   string `json:"ConfigKey" xml:"ConfigKey"`
				ConfigValue string `json:"ConfigValue" xml:"ConfigValue"`
				ClusterId   int64  `json:"ClusterId" xml:"ClusterId"`
				ConfigType  string `json:"ConfigType" xml:"ConfigType"`
				TargetId    int64  `json:"TargetId" xml:"TargetId"`
				Category    string `json:"Category" xml:"Category"`
				Status      string `json:"Status" xml:"Status"`
				Note        string `json:"Note" xml:"Note"`
			} `json:"EcmResourcePoolConfigList" xml:"EcmResourcePoolConfigList"`
		} `json:"QueueList" xml:"QueueList"`
		DefaultSettingList []struct {
			Id          int64  `json:"Id" xml:"Id"`
			ConfigKey   string `json:"ConfigKey" xml:"ConfigKey"`
			ConfigValue string `json:"ConfigValue" xml:"ConfigValue"`
			ClusterId   int64  `json:"ClusterId" xml:"ClusterId"`
			ConfigType  string `json:"ConfigType" xml:"ConfigType"`
			TargetId    int64  `json:"TargetId" xml:"TargetId"`
			Category    string `json:"Category" xml:"Category"`
			Status      string `json:"Status" xml:"Status"`
			Note        string `json:"Note" xml:"Note"`
		} `json:"DefaultSettingList" xml:"DefaultSettingList"`
		AccessControlSettingList []struct {
			Id          int64  `json:"Id" xml:"Id"`
			ConfigKey   string `json:"ConfigKey" xml:"ConfigKey"`
			ConfigValue string `json:"ConfigValue" xml:"ConfigValue"`
			ClusterId   int64  `json:"ClusterId" xml:"ClusterId"`
			ConfigType  string `json:"ConfigType" xml:"ConfigType"`
			TargetId    int64  `json:"TargetId" xml:"TargetId"`
			Category    string `json:"Category" xml:"Category"`
			Status      string `json:"Status" xml:"Status"`
			Note        string `json:"Note" xml:"Note"`
		} `json:"AccessControlSettingList" xml:"AccessControlSettingList"`
	} `json:"PoolInfoList" xml:"PoolInfoList"`
}

func CreateListResourcePoolRequest() (request *ListResourcePoolRequest) {
	request = &ListResourcePoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListResourcePool", "", "")
	return
}

func CreateListResourcePoolResponse() (response *ListResourcePoolResponse) {
	response = &ListResourcePoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
