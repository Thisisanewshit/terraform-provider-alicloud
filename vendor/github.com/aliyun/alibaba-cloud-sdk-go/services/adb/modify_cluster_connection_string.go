package adb

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

// ModifyClusterConnectionString invokes the adb.ModifyClusterConnectionString API synchronously
// api document: https://help.aliyun.com/api/adb/modifyclusterconnectionstring.html
func (client *Client) ModifyClusterConnectionString(request *ModifyClusterConnectionStringRequest) (response *ModifyClusterConnectionStringResponse, err error) {
	response = CreateModifyClusterConnectionStringResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClusterConnectionStringWithChan invokes the adb.ModifyClusterConnectionString API asynchronously
// api document: https://help.aliyun.com/api/adb/modifyclusterconnectionstring.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClusterConnectionStringWithChan(request *ModifyClusterConnectionStringRequest) (<-chan *ModifyClusterConnectionStringResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterConnectionStringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterConnectionString(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyClusterConnectionStringWithCallback invokes the adb.ModifyClusterConnectionString API asynchronously
// api document: https://help.aliyun.com/api/adb/modifyclusterconnectionstring.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClusterConnectionStringWithCallback(request *ModifyClusterConnectionStringRequest, callback func(response *ModifyClusterConnectionStringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterConnectionStringResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterConnectionString(request)
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

// ModifyClusterConnectionStringRequest is the request struct for api ModifyClusterConnectionString
type ModifyClusterConnectionStringRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string           `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string           `position:"Query" name:"DBClusterId"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	CurrentConnectionString string           `position:"Query" name:"CurrentConnectionString"`
	Port                    requests.Integer `position:"Query" name:"Port"`
}

// ModifyClusterConnectionStringResponse is the response struct for api ModifyClusterConnectionString
type ModifyClusterConnectionStringResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyClusterConnectionStringRequest creates a request to invoke ModifyClusterConnectionString API
func CreateModifyClusterConnectionStringRequest() (request *ModifyClusterConnectionStringRequest) {
	request = &ModifyClusterConnectionStringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyClusterConnectionString", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyClusterConnectionStringResponse creates a response to parse from ModifyClusterConnectionString response
func CreateModifyClusterConnectionStringResponse() (response *ModifyClusterConnectionStringResponse) {
	response = &ModifyClusterConnectionStringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}