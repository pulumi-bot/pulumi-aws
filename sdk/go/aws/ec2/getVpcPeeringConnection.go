// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VPC Peering Connection data source provides details about
// a specific VPC peering connection.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := aws_vpc.Foo.Id
// 		opt1 := "10.0.1.0/22"
// 		pc, err := ec2.LookupVpcPeeringConnection(ctx, &ec2.LookupVpcPeeringConnectionArgs{
// 			VpcId:         &opt0,
// 			PeerCidrBlock: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		rt, err := ec2.NewRouteTable(ctx, "rt", &ec2.RouteTableArgs{
// 			VpcId: pulumi.Any(aws_vpc.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
// 			RouteTableId:           rt.ID(),
// 			DestinationCidrBlock:   pulumi.String(pc.PeerCidrBlock),
// 			VpcPeeringConnectionId: pulumi.String(pc.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVpcPeeringConnection(ctx *pulumi.Context, args *LookupVpcPeeringConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpcPeeringConnectionResult, error) {
	var rv LookupVpcPeeringConnectionResult
	err := ctx.Invoke("aws:ec2/getVpcPeeringConnection:getVpcPeeringConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPeeringConnection.
type LookupVpcPeeringConnectionArgs struct {
	// The primary CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Custom filter block as described below.
	Filters []GetVpcPeeringConnectionFilter `pulumi:"filters"`
	// The ID of the specific VPC Peering Connection to retrieve.
	Id *string `pulumi:"id"`
	// The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
	OwnerId *string `pulumi:"ownerId"`
	// The primary CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerCidrBlock *string `pulumi:"peerCidrBlock"`
	// The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerOwnerId *string `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerRegion *string `pulumi:"peerRegion"`
	// The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// The region of the requester VPC of the specific VPC Peering Connection to retrieve.
	Region *string `pulumi:"region"`
	// The status of the specific VPC Peering Connection to retrieve.
	Status *string `pulumi:"status"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired VPC Peering Connection.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpcPeeringConnection.
type LookupVpcPeeringConnectionResult struct {
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter map[string]bool `pulumi:"accepter"`
	// A CIDR block associated to the VPC of the specific VPC Peering Connection.
	CidrBlock string `pulumi:"cidrBlock"`
	// List of objects with CIDR blocks of the requester VPC.
	CidrBlockSets []GetVpcPeeringConnectionCidrBlockSet `pulumi:"cidrBlockSets"`
	Filters       []GetVpcPeeringConnectionFilter       `pulumi:"filters"`
	Id            string                                `pulumi:"id"`
	OwnerId       string                                `pulumi:"ownerId"`
	PeerCidrBlock string                                `pulumi:"peerCidrBlock"`
	// List of objects with CIDR blocks of the accepter VPC.
	PeerCidrBlockSets []GetVpcPeeringConnectionPeerCidrBlockSet `pulumi:"peerCidrBlockSets"`
	PeerOwnerId       string                                    `pulumi:"peerOwnerId"`
	PeerRegion        string                                    `pulumi:"peerRegion"`
	PeerVpcId         string                                    `pulumi:"peerVpcId"`
	Region            string                                    `pulumi:"region"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester map[string]bool   `pulumi:"requester"`
	Status    string            `pulumi:"status"`
	Tags      map[string]string `pulumi:"tags"`
	VpcId     string            `pulumi:"vpcId"`
}

func LookupVpcPeeringConnectionOutput(ctx *pulumi.Context, args LookupVpcPeeringConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPeeringConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPeeringConnectionResult, error) {
			args := v.(LookupVpcPeeringConnectionArgs)
			r, err := LookupVpcPeeringConnection(ctx, &args, opts...)
			return *r, err
		}).(LookupVpcPeeringConnectionResultOutput)
}

// A collection of arguments for invoking getVpcPeeringConnection.
type LookupVpcPeeringConnectionOutputArgs struct {
	// The primary CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
	// Custom filter block as described below.
	Filters GetVpcPeeringConnectionFilterArrayInput `pulumi:"filters"`
	// The ID of the specific VPC Peering Connection to retrieve.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
	OwnerId pulumi.StringPtrInput `pulumi:"ownerId"`
	// The primary CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerCidrBlock pulumi.StringPtrInput `pulumi:"peerCidrBlock"`
	// The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerOwnerId pulumi.StringPtrInput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerRegion pulumi.StringPtrInput `pulumi:"peerRegion"`
	// The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId pulumi.StringPtrInput `pulumi:"peerVpcId"`
	// The region of the requester VPC of the specific VPC Peering Connection to retrieve.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the specific VPC Peering Connection to retrieve.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired VPC Peering Connection.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupVpcPeeringConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPeeringConnectionArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPeeringConnection.
type LookupVpcPeeringConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPeeringConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPeeringConnectionResult)(nil)).Elem()
}

func (o LookupVpcPeeringConnectionResultOutput) ToLookupVpcPeeringConnectionResultOutput() LookupVpcPeeringConnectionResultOutput {
	return o
}

func (o LookupVpcPeeringConnectionResultOutput) ToLookupVpcPeeringConnectionResultOutputWithContext(ctx context.Context) LookupVpcPeeringConnectionResultOutput {
	return o
}

// A configuration block that describes [VPC Peering Connection]
// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
func (o LookupVpcPeeringConnectionResultOutput) Accepter() pulumi.BoolMapOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) map[string]bool { return v.Accepter }).(pulumi.BoolMapOutput)
}

// A CIDR block associated to the VPC of the specific VPC Peering Connection.
func (o LookupVpcPeeringConnectionResultOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// List of objects with CIDR blocks of the requester VPC.
func (o LookupVpcPeeringConnectionResultOutput) CidrBlockSets() GetVpcPeeringConnectionCidrBlockSetArrayOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) []GetVpcPeeringConnectionCidrBlockSet { return v.CidrBlockSets }).(GetVpcPeeringConnectionCidrBlockSetArrayOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Filters() GetVpcPeeringConnectionFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) []GetVpcPeeringConnectionFilter { return v.Filters }).(GetVpcPeeringConnectionFilterArrayOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) PeerCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.PeerCidrBlock }).(pulumi.StringOutput)
}

// List of objects with CIDR blocks of the accepter VPC.
func (o LookupVpcPeeringConnectionResultOutput) PeerCidrBlockSets() GetVpcPeeringConnectionPeerCidrBlockSetArrayOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) []GetVpcPeeringConnectionPeerCidrBlockSet {
		return v.PeerCidrBlockSets
	}).(GetVpcPeeringConnectionPeerCidrBlockSetArrayOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) PeerOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.PeerOwnerId }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) PeerRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.PeerRegion }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) PeerVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.PeerVpcId }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.Region }).(pulumi.StringOutput)
}

// A configuration block that describes [VPC Peering Connection]
// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
func (o LookupVpcPeeringConnectionResultOutput) Requester() pulumi.BoolMapOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) map[string]bool { return v.Requester }).(pulumi.BoolMapOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVpcPeeringConnectionResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPeeringConnectionResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPeeringConnectionResultOutput{})
}
