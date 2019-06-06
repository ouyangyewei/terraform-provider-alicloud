package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBInstanceNetInfo invokes the gpdb.DescribeDBInstanceNetInfo API synchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstancenetinfo.html
func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (response *DescribeDBInstanceNetInfoResponse, err error) {
	response = CreateDescribeDBInstanceNetInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceNetInfoWithChan invokes the gpdb.DescribeDBInstanceNetInfo API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstancenetinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceNetInfoWithChan(request *DescribeDBInstanceNetInfoRequest) (<-chan *DescribeDBInstanceNetInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceNetInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceNetInfo(request)
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

// DescribeDBInstanceNetInfoWithCallback invokes the gpdb.DescribeDBInstanceNetInfo API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstancenetinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceNetInfoWithCallback(request *DescribeDBInstanceNetInfoRequest, callback func(response *DescribeDBInstanceNetInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceNetInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceNetInfo(request)
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

// DescribeDBInstanceNetInfoRequest is the request struct for api DescribeDBInstanceNetInfo
type DescribeDBInstanceNetInfoRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceNetInfoResponse is the response struct for api DescribeDBInstanceNetInfo
type DescribeDBInstanceNetInfoResponse struct {
	*responses.BaseResponse
	RequestId           string             `json:"RequestId" xml:"RequestId"`
	InstanceNetworkType string             `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	DBInstanceNetInfos  DBInstanceNetInfos `json:"DBInstanceNetInfos" xml:"DBInstanceNetInfos"`
}

// CreateDescribeDBInstanceNetInfoRequest creates a request to invoke DescribeDBInstanceNetInfo API
func CreateDescribeDBInstanceNetInfoRequest() (request *DescribeDBInstanceNetInfoRequest) {
	request = &DescribeDBInstanceNetInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBInstanceNetInfo", "gpdb", "openAPI")
	return
}

// CreateDescribeDBInstanceNetInfoResponse creates a response to parse from DescribeDBInstanceNetInfo response
func CreateDescribeDBInstanceNetInfoResponse() (response *DescribeDBInstanceNetInfoResponse) {
	response = &DescribeDBInstanceNetInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
