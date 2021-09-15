// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vpciface provides an interface to enable mocking the vpc service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vpciface

import (
	"github.com/KscSDK/ksc-sdk-go/service/vpc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// VpcAPI provides an interface to enable mocking the
// vpc.Vpc service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // vpc.
//    func myFunc(svc vpciface.VpcAPI) bool {
//        // Make svc.AcceptVpcPeeringConnection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := vpc.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockVpcClient struct {
//        vpciface.VpcAPI
//    }
//    func (m *mockVpcClient) AcceptVpcPeeringConnection(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockVpcClient{}
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
type VpcAPI interface {
	AcceptVpcPeeringConnection(*map[string]interface{}) (*map[string]interface{}, error)
	AcceptVpcPeeringConnectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AcceptVpcPeeringConnectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AllocateSubnetIpv6CidrBlock(*map[string]interface{}) (*map[string]interface{}, error)
	AllocateSubnetIpv6CidrBlockWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AllocateSubnetIpv6CidrBlockRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateNat(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateNatWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateNatRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateNetworkAcl(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateNetworkAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateNetworkAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateVpcCidrBlock(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateVpcCidrBlockWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateVpcCidrBlockRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AuthorizeSecurityGroupEntry(*map[string]interface{}) (*map[string]interface{}, error)
	AuthorizeSecurityGroupEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AuthorizeSecurityGroupEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCustomerGateway(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCustomerGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCustomerGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateIpv6PublicIp(*map[string]interface{}) (*map[string]interface{}, error)
	CreateIpv6PublicIpWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateIpv6PublicIpRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNat(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNatWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNatRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNetworkAcl(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNetworkAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNetworkAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNetworkAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNetworkAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNetworkAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNetworkInterface(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNetworkInterfaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNetworkInterfaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRoute(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRouteWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRouteRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSubnet(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSubnetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSubnetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpc(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpcPeeringConnection(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpcPeeringConnectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpcPeeringConnectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnGateway(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnTunnel(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnTunnelWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnTunnelRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCustomerGateway(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCustomerGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCustomerGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNat(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNatWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNatRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNetworkAcl(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNetworkAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNetworkAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNetworkAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNetworkAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNetworkAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNetworkInterface(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNetworkInterfaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNetworkInterfaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRoute(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRouteWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRouteRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSubnet(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSubnetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSubnetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpc(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpcPeeringConnection(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpcPeeringConnectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpcPeeringConnectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnGateway(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnTunnel(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnTunnelWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnTunnelRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCustomerGateways(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCustomerGatewaysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCustomerGatewaysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInternetGateways(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInternetGatewaysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInternetGatewaysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeIpv6NetworkInterfaces(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeIpv6NetworkInterfacesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeIpv6NetworkInterfacesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeIpv6PublicIpAddresses(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeIpv6PublicIpAddressesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeIpv6PublicIpAddressesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNats(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNetworkAcls(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNetworkAclsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNetworkAclsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNetworkInterfaces(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNetworkInterfacesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNetworkInterfacesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRoutes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRoutesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRoutesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecurityGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecurityGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecurityGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubnetAllocatedIpAddresses(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubnetAllocatedIpAddressesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubnetAllocatedIpAddressesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubnetAvailableAddresses(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubnetAvailableAddressesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubnetAvailableAddressesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubnets(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubnetsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubnetsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcPeeringConnections(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcPeeringConnectionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcPeeringConnectionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcs(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGateways(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewaysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewaysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnTunnels(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnTunnelsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnTunnelsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateNat(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateNatWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateNatRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateNetworkAcl(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateNetworkAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateNetworkAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCustomerGateway(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCustomerGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCustomerGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyIpv6PublicIp(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyIpv6PublicIpWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyIpv6PublicIpRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNat(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNatWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNatRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkAcl(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkInterface(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkInterfaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkInterfaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroupEntry(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySubnet(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySubnetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySubnetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpc(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpcPeeringConnection(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpcPeeringConnectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpcPeeringConnectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpnGateway(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpnGatewayWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpnGatewayRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpnTunnel(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpnTunnelWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpnTunnelRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RejectVpcPeeringConnection(*map[string]interface{}) (*map[string]interface{}, error)
	RejectVpcPeeringConnectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RejectVpcPeeringConnectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReleaseIpv6PublicIp(*map[string]interface{}) (*map[string]interface{}, error)
	ReleaseIpv6PublicIpWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReleaseIpv6PublicIpRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeSecurityGroupEntry(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeSecurityGroupEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeSecurityGroupEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ VpcAPI = (*vpc.Vpc)(nil)
