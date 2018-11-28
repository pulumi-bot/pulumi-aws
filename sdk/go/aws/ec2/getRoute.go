// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `aws_route` provides details about a specific Route.
// 
// This resource can prove useful when finding the resource
// associated with a CIDR. For example, finding the peering
// connection associated with a CIDR value.
func LookupRoute(ctx *pulumi.Context, args *GetRouteArgs) (*GetRouteResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["destinationCidrBlock"] = args.DestinationCidrBlock
		inputs["destinationIpv6CidrBlock"] = args.DestinationIpv6CidrBlock
		inputs["egressOnlyGatewayId"] = args.EgressOnlyGatewayId
		inputs["gatewayId"] = args.GatewayId
		inputs["instanceId"] = args.InstanceId
		inputs["natGatewayId"] = args.NatGatewayId
		inputs["networkInterfaceId"] = args.NetworkInterfaceId
		inputs["routeTableId"] = args.RouteTableId
		inputs["transitGatewayId"] = args.TransitGatewayId
		inputs["vpcPeeringConnectionId"] = args.VpcPeeringConnectionId
	}
	outputs, err := ctx.Invoke("aws:ec2/getRoute:getRoute", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRouteResult{
		DestinationCidrBlock: outputs["destinationCidrBlock"],
		DestinationIpv6CidrBlock: outputs["destinationIpv6CidrBlock"],
		EgressOnlyGatewayId: outputs["egressOnlyGatewayId"],
		GatewayId: outputs["gatewayId"],
		InstanceId: outputs["instanceId"],
		NatGatewayId: outputs["natGatewayId"],
		NetworkInterfaceId: outputs["networkInterfaceId"],
		TransitGatewayId: outputs["transitGatewayId"],
		VpcPeeringConnectionId: outputs["vpcPeeringConnectionId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRoute.
type GetRouteArgs struct {
	// The CIDR block of the Route belonging to the Route Table.
	DestinationCidrBlock interface{}
	// The IPv6 CIDR block of the Route belonging to the Route Table.
	DestinationIpv6CidrBlock interface{}
	// The Egress Only Gateway ID of the Route belonging to the Route Table.
	EgressOnlyGatewayId interface{}
	// The Gateway ID of the Route belonging to the Route Table.
	GatewayId interface{}
	// The Instance ID of the Route belonging to the Route Table.
	InstanceId interface{}
	// The NAT Gateway ID of the Route belonging to the Route Table.
	NatGatewayId interface{}
	// The Network Interface ID of the Route belonging to the Route Table.
	NetworkInterfaceId interface{}
	// The id of the specific Route Table containing the Route entry.
	RouteTableId interface{}
	// The EC2 Transit Gateway ID of the Route belonging to the Route Table.
	TransitGatewayId interface{}
	// The VPC Peering Connection ID of the Route belonging to the Route Table.
	VpcPeeringConnectionId interface{}
}

// A collection of values returned by getRoute.
type GetRouteResult struct {
	DestinationCidrBlock interface{}
	DestinationIpv6CidrBlock interface{}
	EgressOnlyGatewayId interface{}
	GatewayId interface{}
	InstanceId interface{}
	NatGatewayId interface{}
	NetworkInterfaceId interface{}
	TransitGatewayId interface{}
	VpcPeeringConnectionId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
