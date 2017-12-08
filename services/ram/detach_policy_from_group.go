
package ram

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

func (client *Client) DetachPolicyFromGroup(request *DetachPolicyFromGroupRequest) (response *DetachPolicyFromGroupResponse, err error) {
response = CreateDetachPolicyFromGroupResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DetachPolicyFromGroupWithChan(request *DetachPolicyFromGroupRequest) (<-chan *DetachPolicyFromGroupResponse, <-chan error) {
responseChan := make(chan *DetachPolicyFromGroupResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DetachPolicyFromGroup(request)
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

func (client *Client) DetachPolicyFromGroupWithCallback(request *DetachPolicyFromGroupRequest, callback func(response *DetachPolicyFromGroupResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DetachPolicyFromGroupResponse
var err error
defer close(result)
response, err = client.DetachPolicyFromGroup(request)
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

type DetachPolicyFromGroupRequest struct {
*requests.RpcRequest
                PolicyType  string `position:"Query" name:"PolicyType"`
                GroupName  string `position:"Query" name:"GroupName"`
                PolicyName  string `position:"Query" name:"PolicyName"`
}


type DetachPolicyFromGroupResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateDetachPolicyFromGroupRequest() (request *DetachPolicyFromGroupRequest) {
request = &DetachPolicyFromGroupRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromGroup", "", "")
return
}

func CreateDetachPolicyFromGroupResponse() (response *DetachPolicyFromGroupResponse) {
response = &DetachPolicyFromGroupResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
