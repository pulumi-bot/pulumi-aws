// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudhsmv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an HSM module in Amazon CloudHSM v2 cluster.
//
// ## Example Usage
//
// The following example below creates an HSM module in CloudHSM cluster.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudhsmv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cluster, err := cloudhsmv2.LookupCluster(ctx, &cloudhsmv2.LookupClusterArgs{
// 			ClusterId: _var.Cloudhsm_cluster_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudhsmv2.NewHsm(ctx, "cloudhsmV2Hsm", &cloudhsmv2.HsmArgs{
// 			SubnetId:  pulumi.String(cluster.SubnetIds[0]),
// 			ClusterId: pulumi.String(cluster.ClusterId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// HSM modules can be imported using their HSM ID, e.g.
//
// ```sh
//  $ pulumi import aws:cloudhsmv2/hsm:Hsm bar hsm-quo8dahtaca
// ```
type Hsm struct {
	pulumi.CustomResourceState

	// The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The id of the ENI interface allocated for HSM module.
	HsmEniId pulumi.StringOutput `pulumi:"hsmEniId"`
	// The id of the HSM module.
	HsmId pulumi.StringOutput `pulumi:"hsmId"`
	// The state of the HSM module.
	HsmState pulumi.StringOutput `pulumi:"hsmState"`
	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The ID of subnet in which HSM module will be located.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewHsm registers a new resource with the given unique name, arguments, and options.
func NewHsm(ctx *pulumi.Context,
	name string, args *HsmArgs, opts ...pulumi.ResourceOption) (*Hsm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	var resource Hsm
	err := ctx.RegisterResource("aws:cloudhsmv2/hsm:Hsm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHsm gets an existing Hsm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHsm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HsmState, opts ...pulumi.ResourceOption) (*Hsm, error) {
	var resource Hsm
	err := ctx.ReadResource("aws:cloudhsmv2/hsm:Hsm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hsm resources.
type hsmState struct {
	// The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId *string `pulumi:"clusterId"`
	// The id of the ENI interface allocated for HSM module.
	HsmEniId *string `pulumi:"hsmEniId"`
	// The id of the HSM module.
	HsmId *string `pulumi:"hsmId"`
	// The state of the HSM module.
	HsmState *string `pulumi:"hsmState"`
	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of subnet in which HSM module will be located.
	SubnetId *string `pulumi:"subnetId"`
}

type HsmState struct {
	// The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
	AvailabilityZone pulumi.StringPtrInput
	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId pulumi.StringPtrInput
	// The id of the ENI interface allocated for HSM module.
	HsmEniId pulumi.StringPtrInput
	// The id of the HSM module.
	HsmId pulumi.StringPtrInput
	// The state of the HSM module.
	HsmState pulumi.StringPtrInput
	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress pulumi.StringPtrInput
	// The ID of subnet in which HSM module will be located.
	SubnetId pulumi.StringPtrInput
}

func (HsmState) ElementType() reflect.Type {
	return reflect.TypeOf((*hsmState)(nil)).Elem()
}

type hsmArgs struct {
	// The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId string `pulumi:"clusterId"`
	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of subnet in which HSM module will be located.
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a Hsm resource.
type HsmArgs struct {
	// The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
	AvailabilityZone pulumi.StringPtrInput
	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId pulumi.StringInput
	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress pulumi.StringPtrInput
	// The ID of subnet in which HSM module will be located.
	SubnetId pulumi.StringPtrInput
}

func (HsmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hsmArgs)(nil)).Elem()
}

type HsmInput interface {
	pulumi.Input

	ToHsmOutput() HsmOutput
	ToHsmOutputWithContext(ctx context.Context) HsmOutput
}

func (Hsm) ElementType() reflect.Type {
	return reflect.TypeOf((*Hsm)(nil)).Elem()
}

func (i Hsm) ToHsmOutput() HsmOutput {
	return i.ToHsmOutputWithContext(context.Background())
}

func (i Hsm) ToHsmOutputWithContext(ctx context.Context) HsmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HsmOutput)
}

type HsmOutput struct {
	*pulumi.OutputState
}

func (HsmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HsmOutput)(nil)).Elem()
}

func (o HsmOutput) ToHsmOutput() HsmOutput {
	return o
}

func (o HsmOutput) ToHsmOutputWithContext(ctx context.Context) HsmOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HsmOutput{})
}
