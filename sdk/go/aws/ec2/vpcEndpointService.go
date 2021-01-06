// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC Endpoint Service resource.
// Service consumers can create an _Interface_ VPC Endpoint to connect to the service.
//
// > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
// both a standalone VPC Endpoint Service Allowed Principal resource
// and a VPC Endpoint Service resource with an `allowedPrincipals` attribute. Do not use the same principal ARN in both
// a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
// and will overwrite the association.
//
// ## Example Usage
//
// ## Import
//
// VPC Endpoint Services can be imported using the `VPC endpoint service id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/vpcEndpointService:VpcEndpointService foo vpce-svc-0f97a19d3fa8220bc
// ```
type VpcEndpointService struct {
	pulumi.CustomResourceState

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired pulumi.BoolOutput `pulumi:"acceptanceRequired"`
	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals pulumi.StringArrayOutput `pulumi:"allowedPrincipals"`
	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zones in which the service is available.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The DNS names for the service.
	BaseEndpointDnsNames pulumi.StringArrayOutput `pulumi:"baseEndpointDnsNames"`
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns pulumi.StringArrayOutput `pulumi:"gatewayLoadBalancerArns"`
	// Whether or not the service manages its VPC endpoints - `true` or `false`.
	ManagesVpcEndpoints pulumi.BoolOutput `pulumi:"managesVpcEndpoints"`
	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns pulumi.StringArrayOutput `pulumi:"networkLoadBalancerArns"`
	// The private DNS name for the service.
	PrivateDnsName pulumi.StringOutput `pulumi:"privateDnsName"`
	// The service name.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The service type, `Gateway` or `Interface`.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// The state of the VPC endpoint service.
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVpcEndpointService registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointService(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceArgs, opts ...pulumi.ResourceOption) (*VpcEndpointService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceptanceRequired == nil {
		return nil, errors.New("invalid value for required argument 'AcceptanceRequired'")
	}
	var resource VpcEndpointService
	err := ctx.RegisterResource("aws:ec2/vpcEndpointService:VpcEndpointService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointService gets an existing VpcEndpointService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceState, opts ...pulumi.ResourceOption) (*VpcEndpointService, error) {
	var resource VpcEndpointService
	err := ctx.ReadResource("aws:ec2/vpcEndpointService:VpcEndpointService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointService resources.
type vpcEndpointServiceState struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired *bool `pulumi:"acceptanceRequired"`
	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn *string `pulumi:"arn"`
	// The Availability Zones in which the service is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DNS names for the service.
	BaseEndpointDnsNames []string `pulumi:"baseEndpointDnsNames"`
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns []string `pulumi:"gatewayLoadBalancerArns"`
	// Whether or not the service manages its VPC endpoints - `true` or `false`.
	ManagesVpcEndpoints *bool `pulumi:"managesVpcEndpoints"`
	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns []string `pulumi:"networkLoadBalancerArns"`
	// The private DNS name for the service.
	PrivateDnsName *string `pulumi:"privateDnsName"`
	// The service name.
	ServiceName *string `pulumi:"serviceName"`
	// The service type, `Gateway` or `Interface`.
	ServiceType *string `pulumi:"serviceType"`
	// The state of the VPC endpoint service.
	State *string `pulumi:"state"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VpcEndpointServiceState struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired pulumi.BoolPtrInput
	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn pulumi.StringPtrInput
	// The Availability Zones in which the service is available.
	AvailabilityZones pulumi.StringArrayInput
	// The DNS names for the service.
	BaseEndpointDnsNames pulumi.StringArrayInput
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns pulumi.StringArrayInput
	// Whether or not the service manages its VPC endpoints - `true` or `false`.
	ManagesVpcEndpoints pulumi.BoolPtrInput
	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns pulumi.StringArrayInput
	// The private DNS name for the service.
	PrivateDnsName pulumi.StringPtrInput
	// The service name.
	ServiceName pulumi.StringPtrInput
	// The service type, `Gateway` or `Interface`.
	ServiceType pulumi.StringPtrInput
	// The state of the VPC endpoint service.
	State pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcEndpointServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceState)(nil)).Elem()
}

type vpcEndpointServiceArgs struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired bool `pulumi:"acceptanceRequired"`
	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns []string `pulumi:"gatewayLoadBalancerArns"`
	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns []string `pulumi:"networkLoadBalancerArns"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcEndpointService resource.
type VpcEndpointServiceArgs struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired pulumi.BoolInput
	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals pulumi.StringArrayInput
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns pulumi.StringArrayInput
	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcEndpointServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceArgs)(nil)).Elem()
}

type VpcEndpointServiceInput interface {
	pulumi.Input

	ToVpcEndpointServiceOutput() VpcEndpointServiceOutput
	ToVpcEndpointServiceOutputWithContext(ctx context.Context) VpcEndpointServiceOutput
}

func (*VpcEndpointService) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointService)(nil))
}

func (i *VpcEndpointService) ToVpcEndpointServiceOutput() VpcEndpointServiceOutput {
	return i.ToVpcEndpointServiceOutputWithContext(context.Background())
}

func (i *VpcEndpointService) ToVpcEndpointServiceOutputWithContext(ctx context.Context) VpcEndpointServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceOutput)
}

func (i *VpcEndpointService) ToVpcEndpointServicePtrOutput() VpcEndpointServicePtrOutput {
	return i.ToVpcEndpointServicePtrOutputWithContext(context.Background())
}

func (i *VpcEndpointService) ToVpcEndpointServicePtrOutputWithContext(ctx context.Context) VpcEndpointServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServicePtrOutput)
}

type VpcEndpointServicePtrInput interface {
	pulumi.Input

	ToVpcEndpointServicePtrOutput() VpcEndpointServicePtrOutput
	ToVpcEndpointServicePtrOutputWithContext(ctx context.Context) VpcEndpointServicePtrOutput
}

type vpcEndpointServicePtrType VpcEndpointServiceArgs

func (*vpcEndpointServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointService)(nil))
}

func (i *vpcEndpointServicePtrType) ToVpcEndpointServicePtrOutput() VpcEndpointServicePtrOutput {
	return i.ToVpcEndpointServicePtrOutputWithContext(context.Background())
}

func (i *vpcEndpointServicePtrType) ToVpcEndpointServicePtrOutputWithContext(ctx context.Context) VpcEndpointServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceOutput).ToVpcEndpointServicePtrOutput()
}

// VpcEndpointServiceArrayInput is an input type that accepts VpcEndpointServiceArray and VpcEndpointServiceArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceArrayInput` via:
//
//          VpcEndpointServiceArray{ VpcEndpointServiceArgs{...} }
type VpcEndpointServiceArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceArrayOutput() VpcEndpointServiceArrayOutput
	ToVpcEndpointServiceArrayOutputWithContext(context.Context) VpcEndpointServiceArrayOutput
}

type VpcEndpointServiceArray []VpcEndpointServiceInput

func (VpcEndpointServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointService)(nil))
}

func (i VpcEndpointServiceArray) ToVpcEndpointServiceArrayOutput() VpcEndpointServiceArrayOutput {
	return i.ToVpcEndpointServiceArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceArray) ToVpcEndpointServiceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceArrayOutput)
}

// VpcEndpointServiceMapInput is an input type that accepts VpcEndpointServiceMap and VpcEndpointServiceMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceMapInput` via:
//
//          VpcEndpointServiceMap{ "key": VpcEndpointServiceArgs{...} }
type VpcEndpointServiceMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceMapOutput() VpcEndpointServiceMapOutput
	ToVpcEndpointServiceMapOutputWithContext(context.Context) VpcEndpointServiceMapOutput
}

type VpcEndpointServiceMap map[string]VpcEndpointServiceInput

func (VpcEndpointServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcEndpointService)(nil))
}

func (i VpcEndpointServiceMap) ToVpcEndpointServiceMapOutput() VpcEndpointServiceMapOutput {
	return i.ToVpcEndpointServiceMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceMap) ToVpcEndpointServiceMapOutputWithContext(ctx context.Context) VpcEndpointServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceMapOutput)
}

type VpcEndpointServiceOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointService)(nil))
}

func (o VpcEndpointServiceOutput) ToVpcEndpointServiceOutput() VpcEndpointServiceOutput {
	return o
}

func (o VpcEndpointServiceOutput) ToVpcEndpointServiceOutputWithContext(ctx context.Context) VpcEndpointServiceOutput {
	return o
}

func (o VpcEndpointServiceOutput) ToVpcEndpointServicePtrOutput() VpcEndpointServicePtrOutput {
	return o.ToVpcEndpointServicePtrOutputWithContext(context.Background())
}

func (o VpcEndpointServiceOutput) ToVpcEndpointServicePtrOutputWithContext(ctx context.Context) VpcEndpointServicePtrOutput {
	return o.ApplyT(func(v VpcEndpointService) *VpcEndpointService {
		return &v
	}).(VpcEndpointServicePtrOutput)
}

type VpcEndpointServicePtrOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointService)(nil))
}

func (o VpcEndpointServicePtrOutput) ToVpcEndpointServicePtrOutput() VpcEndpointServicePtrOutput {
	return o
}

func (o VpcEndpointServicePtrOutput) ToVpcEndpointServicePtrOutputWithContext(ctx context.Context) VpcEndpointServicePtrOutput {
	return o
}

type VpcEndpointServiceArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointService)(nil))
}

func (o VpcEndpointServiceArrayOutput) ToVpcEndpointServiceArrayOutput() VpcEndpointServiceArrayOutput {
	return o
}

func (o VpcEndpointServiceArrayOutput) ToVpcEndpointServiceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceArrayOutput {
	return o
}

func (o VpcEndpointServiceArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpointService {
		return vs[0].([]VpcEndpointService)[vs[1].(int)]
	}).(VpcEndpointServiceOutput)
}

type VpcEndpointServiceMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcEndpointService)(nil))
}

func (o VpcEndpointServiceMapOutput) ToVpcEndpointServiceMapOutput() VpcEndpointServiceMapOutput {
	return o
}

func (o VpcEndpointServiceMapOutput) ToVpcEndpointServiceMapOutputWithContext(ctx context.Context) VpcEndpointServiceMapOutput {
	return o
}

func (o VpcEndpointServiceMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcEndpointService {
		return vs[0].(map[string]VpcEndpointService)[vs[1].(string)]
	}).(VpcEndpointServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointServiceOutput{})
	pulumi.RegisterOutputType(VpcEndpointServicePtrOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceMapOutput{})
}
