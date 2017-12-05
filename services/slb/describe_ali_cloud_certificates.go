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

func (client *Client) DescribeAliCloudCertificates(request *DescribeAliCloudCertificatesRequest) (response *DescribeAliCloudCertificatesResponse, err error) {
	response = CreateDescribeAliCloudCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAliCloudCertificatesWithChan(request *DescribeAliCloudCertificatesRequest) (<-chan *DescribeAliCloudCertificatesResponse, <-chan error) {
	responseChan := make(chan *DescribeAliCloudCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAliCloudCertificates(request)
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

func (client *Client) DescribeAliCloudCertificatesWithCallback(request *DescribeAliCloudCertificatesRequest, callback func(response *DescribeAliCloudCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAliCloudCertificatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAliCloudCertificates(request)
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

type DescribeAliCloudCertificatesRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeAliCloudCertificatesResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	AliCloudCertificates []struct {
		AliCloudCertificateId   string `json:"AliCloudCertificateId" xml:"AliCloudCertificateId"`
		AliCloudCertificateName string `json:"AliCloudCertificateName" xml:"AliCloudCertificateName"`
		Fingerprint             string `json:"Fingerprint" xml:"Fingerprint"`
		DomainName              string `json:"DomainName" xml:"DomainName"`
		Issuer                  string `json:"Issuer" xml:"Issuer"`
	} `json:"AliCloudCertificates" xml:"AliCloudCertificates"`
}

func CreateDescribeAliCloudCertificatesRequest() (request *DescribeAliCloudCertificatesRequest) {
	request = &DescribeAliCloudCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeAliCloudCertificates", "", "")
	return
}

func CreateDescribeAliCloudCertificatesResponse() (response *DescribeAliCloudCertificatesResponse) {
	response = &DescribeAliCloudCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
