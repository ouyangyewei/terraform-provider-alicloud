package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateAccount invokes the gpdb.CreateAccount API synchronously
// api document: https://help.aliyun.com/api/gpdb/createaccount.html
func (client *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
	response = CreateCreateAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAccountWithChan invokes the gpdb.CreateAccount API asynchronously
// api document: https://help.aliyun.com/api/gpdb/createaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccountWithChan(request *CreateAccountRequest) (<-chan *CreateAccountResponse, <-chan error) {
	responseChan := make(chan *CreateAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAccount(request)
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

// CreateAccountWithCallback invokes the gpdb.CreateAccount API asynchronously
// api document: https://help.aliyun.com/api/gpdb/createaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccountWithCallback(request *CreateAccountRequest, callback func(response *CreateAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateAccount(request)
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

// CreateAccountRequest is the request struct for api CreateAccount
type CreateAccountRequest struct {
	*requests.RpcRequest
	AccountPassword    string           `position:"Query" name:"AccountPassword"`
	AccountName        string           `position:"Query" name:"AccountName"`
	DatabaseName       string           `position:"Query" name:"DatabaseName"`
	DBInstanceId       string           `position:"Query" name:"DBInstanceId"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	AccountDescription string           `position:"Query" name:"AccountDescription"`
}

// CreateAccountResponse is the response struct for api CreateAccount
type CreateAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateAccountRequest creates a request to invoke CreateAccount API
func CreateCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "CreateAccount", "gpdb", "openAPI")
	return
}

// CreateCreateAccountResponse creates a response to parse from CreateAccount response
func CreateCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
