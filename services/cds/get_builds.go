package cds

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

func (client *Client) GetBuilds(request *GetBuildsRequest) (response *GetBuildsResponse, err error) {
	response = CreateGetBuildsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetBuildsWithChan(request *GetBuildsRequest) (<-chan *GetBuildsResponse, <-chan error) {
	responseChan := make(chan *GetBuildsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBuilds(request)
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

func (client *Client) GetBuildsWithCallback(request *GetBuildsRequest, callback func(response *GetBuildsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBuildsResponse
		var err error
		defer close(result)
		response, err = client.GetBuilds(request)
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

type GetBuildsRequest struct {
	*requests.RoaRequest
	NumberPerPage string `position:"Query" name:"NumberPerPage"`
	JobName       string `position:"Path" name:"JobName"`
	Start         string `position:"Query" name:"Start"`
}

type GetBuildsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Builds    []struct {
		BuildNumber int    `json:"BuildNumber" xml:"BuildNumber"`
		Duration    int    `json:"Duration" xml:"Duration"`
		StartTime   int64  `json:"StartTime" xml:"StartTime"`
		Log         string `json:"Log" xml:"Log"`
		BuildEnv    string `json:"BuildEnv" xml:"BuildEnv"`
	} `json:"Builds" xml:"Builds"`
}

func CreateGetBuildsRequest() (request *GetBuildsRequest) {
	request = &GetBuildsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cds", "2017-09-25", "GetBuilds", "/v1/job/[JobName]/builds", "", "")
	return
}

func CreateGetBuildsResponse() (response *GetBuildsResponse) {
	response = &GetBuildsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
