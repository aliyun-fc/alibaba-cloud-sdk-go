
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

func (client *Client) SetLoadBalancerStatus(request *SetLoadBalancerStatusRequest) (response *SetLoadBalancerStatusResponse, err error) {
response = CreateSetLoadBalancerStatusResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) SetLoadBalancerStatusWithChan(request *SetLoadBalancerStatusRequest) (<-chan *SetLoadBalancerStatusResponse, <-chan error) {
responseChan := make(chan *SetLoadBalancerStatusResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SetLoadBalancerStatus(request)
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

func (client *Client) SetLoadBalancerStatusWithCallback(request *SetLoadBalancerStatusRequest, callback func(response *SetLoadBalancerStatusResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SetLoadBalancerStatusResponse
var err error
defer close(result)
response, err = client.SetLoadBalancerStatus(request)
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

type SetLoadBalancerStatusRequest struct {
*requests.RpcRequest
                Tags  string `position:"Query" name:"Tags"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                LoadBalancerStatus  string `position:"Query" name:"LoadBalancerStatus"`
                AccessKeyId  string `position:"Query" name:"access_key_id"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                LoadBalancerId  string `position:"Query" name:"LoadBalancerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type SetLoadBalancerStatusResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerStatusRequest() (request *SetLoadBalancerStatusRequest) {
request = &SetLoadBalancerStatusRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerStatus", "", "")
return
}

func CreateSetLoadBalancerStatusResponse() (response *SetLoadBalancerStatusResponse) {
response = &SetLoadBalancerStatusResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
