// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package eipiface provides an interface to enable mocking the eip service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package eipiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/ksc/ksc-sdk-go/service/eip"
)

// EipAPI provides an interface to enable mocking the
// eip.Eip service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // eip.
//    func myFunc(svc eipiface.EipAPI) bool {
//        // Make svc.AllocateAddress request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := eip.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEipClient struct {
//        eipiface.EipAPI
//    }
//    func (m *mockEipClient) AllocateAddress(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEipClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EipAPI interface {
	AllocateAddress(*map[string]interface{}) (*map[string]interface{}, error)
	AllocateAddressWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AllocateAddressRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateAddress(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateAddressWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateAddressRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAddresses(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAddressesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAddressesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateAddress(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateAddressWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateAddressRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetLines(*map[string]interface{}) (*map[string]interface{}, error)
	GetLinesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetLinesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAddress(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAddressWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAddressRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RegionList(*map[string]interface{}) (*map[string]interface{}, error)
	RegionListWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RegionListRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReleaseAddress(*map[string]interface{}) (*map[string]interface{}, error)
	ReleaseAddressWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReleaseAddressRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ EipAPI = (*eip.Eip)(nil)
