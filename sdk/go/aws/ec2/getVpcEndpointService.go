// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VPC Endpoint Service data source details about a specific service that
// can be specified when creating a VPC endpoint within the region configured in the provider.
//
// ## Example Usage
// ### AWS Service
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
// 		opt0 := "s3"
// 		opt1 := "Gateway"
// 		s3, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
// 			Service:     &opt0,
// 			ServiceType: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewVpcEndpoint(ctx, "ep", &ec2.VpcEndpointArgs{
// 			VpcId:       foo.ID(),
// 			ServiceName: pulumi.String(s3.ServiceName),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Non-AWS Service
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
// 		opt0 := "com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8"
// 		_, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
// 			ServiceName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter
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
// 		_, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
// 			Filters: []ec2.GetVpcEndpointServiceFilter{
// 				ec2.GetVpcEndpointServiceFilter{
// 					Name: "service-name",
// 					Values: []string{
// 						"some-service",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVpcEndpointService(ctx *pulumi.Context, args *LookupVpcEndpointServiceArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointServiceResult, error) {
	var rv LookupVpcEndpointServiceResult
	err := ctx.Invoke("aws:ec2/getVpcEndpointService:getVpcEndpointService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEndpointService.
type LookupVpcEndpointServiceArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetVpcEndpointServiceFilter `pulumi:"filters"`
	// The common name of an AWS service (e.g. `s3`).
	Service *string `pulumi:"service"`
	// The service name that is specified when creating a VPC endpoint. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName *string `pulumi:"serviceName"`
	// The service type, `Gateway` or `Interface`.
	ServiceType *string `pulumi:"serviceType"`
	// A map of tags, each pair of which must exactly match a pair on the desired VPC Endpoint Service.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcEndpointService.
type LookupVpcEndpointServiceResult struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired bool `pulumi:"acceptanceRequired"`
	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn string `pulumi:"arn"`
	// The Availability Zones in which the service is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DNS names for the service.
	BaseEndpointDnsNames []string                      `pulumi:"baseEndpointDnsNames"`
	Filters              []GetVpcEndpointServiceFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether or not the service manages its VPC endpoints - `true` or `false`.
	ManagesVpcEndpoints bool `pulumi:"managesVpcEndpoints"`
	// The AWS account ID of the service owner or `amazon`.
	Owner string `pulumi:"owner"`
	// The private DNS name for the service.
	PrivateDnsName string  `pulumi:"privateDnsName"`
	Service        *string `pulumi:"service"`
	// The ID of the endpoint service.
	ServiceId   string `pulumi:"serviceId"`
	ServiceName string `pulumi:"serviceName"`
	ServiceType string `pulumi:"serviceType"`
	// A map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not the service supports endpoint policies - `true` or `false`.
	VpcEndpointPolicySupported bool `pulumi:"vpcEndpointPolicySupported"`
}

func LookupVpcEndpointServiceApply(ctx *pulumi.Context, args LookupVpcEndpointServiceApplyInput, opts ...pulumi.InvokeOption) LookupVpcEndpointServiceResultOutput {
	return args.ToLookupVpcEndpointServiceApplyOutput().ApplyT(func(v LookupVpcEndpointServiceArgs) (LookupVpcEndpointServiceResult, error) {
		r, err := LookupVpcEndpointService(ctx, &v, opts...)
		return *r, err

	}).(LookupVpcEndpointServiceResultOutput)
}

// LookupVpcEndpointServiceApplyInput is an input type that accepts LookupVpcEndpointServiceApplyArgs and LookupVpcEndpointServiceApplyOutput values.
// You can construct a concrete instance of `LookupVpcEndpointServiceApplyInput` via:
//
//          LookupVpcEndpointServiceApplyArgs{...}
type LookupVpcEndpointServiceApplyInput interface {
	pulumi.Input

	ToLookupVpcEndpointServiceApplyOutput() LookupVpcEndpointServiceApplyOutput
	ToLookupVpcEndpointServiceApplyOutputWithContext(context.Context) LookupVpcEndpointServiceApplyOutput
}

// A collection of arguments for invoking getVpcEndpointService.
type LookupVpcEndpointServiceApplyArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetVpcEndpointServiceFilterArrayInput `pulumi:"filters"`
	// The common name of an AWS service (e.g. `s3`).
	Service pulumi.StringPtrInput `pulumi:"service"`
	// The service name that is specified when creating a VPC endpoint. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// The service type, `Gateway` or `Interface`.
	ServiceType pulumi.StringPtrInput `pulumi:"serviceType"`
	// A map of tags, each pair of which must exactly match a pair on the desired VPC Endpoint Service.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVpcEndpointServiceApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointServiceArgs)(nil)).Elem()
}

func (i LookupVpcEndpointServiceApplyArgs) ToLookupVpcEndpointServiceApplyOutput() LookupVpcEndpointServiceApplyOutput {
	return i.ToLookupVpcEndpointServiceApplyOutputWithContext(context.Background())
}

func (i LookupVpcEndpointServiceApplyArgs) ToLookupVpcEndpointServiceApplyOutputWithContext(ctx context.Context) LookupVpcEndpointServiceApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupVpcEndpointServiceApplyOutput)
}

// A collection of arguments for invoking getVpcEndpointService.
type LookupVpcEndpointServiceApplyOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointServiceApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointServiceArgs)(nil)).Elem()
}

func (o LookupVpcEndpointServiceApplyOutput) ToLookupVpcEndpointServiceApplyOutput() LookupVpcEndpointServiceApplyOutput {
	return o
}

func (o LookupVpcEndpointServiceApplyOutput) ToLookupVpcEndpointServiceApplyOutputWithContext(ctx context.Context) LookupVpcEndpointServiceApplyOutput {
	return o
}

// Configuration block(s) for filtering. Detailed below.
func (o LookupVpcEndpointServiceApplyOutput) Filters() GetVpcEndpointServiceFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceArgs) []GetVpcEndpointServiceFilter { return v.Filters }).(GetVpcEndpointServiceFilterArrayOutput)
}

// The common name of an AWS service (e.g. `s3`).
func (o LookupVpcEndpointServiceApplyOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceArgs) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// The service name that is specified when creating a VPC endpoint. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
func (o LookupVpcEndpointServiceApplyOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceArgs) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// The service type, `Gateway` or `Interface`.
func (o LookupVpcEndpointServiceApplyOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceArgs) *string { return v.ServiceType }).(pulumi.StringPtrOutput)
}

// A map of tags, each pair of which must exactly match a pair on the desired VPC Endpoint Service.
func (o LookupVpcEndpointServiceApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getVpcEndpointService.
type LookupVpcEndpointServiceResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointServiceResult)(nil)).Elem()
}

func (o LookupVpcEndpointServiceResultOutput) ToLookupVpcEndpointServiceResultOutput() LookupVpcEndpointServiceResultOutput {
	return o
}

func (o LookupVpcEndpointServiceResultOutput) ToLookupVpcEndpointServiceResultOutputWithContext(ctx context.Context) LookupVpcEndpointServiceResultOutput {
	return o
}

// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) AcceptanceRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.AcceptanceRequired }).(pulumi.BoolOutput)
}

// The Amazon Resource Name (ARN) of the VPC endpoint service.
func (o LookupVpcEndpointServiceResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The Availability Zones in which the service is available.
func (o LookupVpcEndpointServiceResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The DNS names for the service.
func (o LookupVpcEndpointServiceResultOutput) BaseEndpointDnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []string { return v.BaseEndpointDnsNames }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointServiceResultOutput) Filters() GetVpcEndpointServiceFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []GetVpcEndpointServiceFilter { return v.Filters }).(GetVpcEndpointServiceFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcEndpointServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether or not the service manages its VPC endpoints - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) ManagesVpcEndpoints() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.ManagesVpcEndpoints }).(pulumi.BoolOutput)
}

// The AWS account ID of the service owner or `amazon`.
func (o LookupVpcEndpointServiceResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Owner }).(pulumi.StringOutput)
}

// The private DNS name for the service.
func (o LookupVpcEndpointServiceResultOutput) PrivateDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.PrivateDnsName }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// The ID of the endpoint service.
func (o LookupVpcEndpointServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceType }).(pulumi.StringOutput)
}

// A map of tags assigned to the resource.
func (o LookupVpcEndpointServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether or not the service supports endpoint policies - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) VpcEndpointPolicySupported() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.VpcEndpointPolicySupported }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointServiceApplyOutput{})
	pulumi.RegisterOutputType(LookupVpcEndpointServiceResultOutput{})
}
