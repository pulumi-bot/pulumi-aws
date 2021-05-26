// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
// in a given region for the purpose of permitting in S3 bucket policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elb"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := elb.GetServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		elbLogs, err := s3.NewBucket(ctx, "elbLogs", &s3.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Id\": \"Policy\",\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancer(ctx, "bar", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-west-2a"),
// 			},
// 			AccessLogs: &elb.LoadBalancerAccessLogsArgs{
// 				Bucket:   elbLogs.Bucket,
// 				Interval: pulumi.Int(5),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(8000),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(80),
// 					LbProtocol:       pulumi.String("http"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetServiceAccount(ctx *pulumi.Context, args *GetServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetServiceAccountResult, error) {
	var rv GetServiceAccountResult
	err := ctx.Invoke("aws:elb/getServiceAccount:getServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountArgs struct {
	// Name of the region whose AWS ELB account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResult struct {
	// The ARN of the AWS ELB service account in the selected region.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}

func GetServiceAccountApply(ctx *pulumi.Context, args GetServiceAccountApplyInput, opts ...pulumi.InvokeOption) GetServiceAccountResultOutput {
	return args.ToGetServiceAccountApplyOutput().ApplyT(func(v GetServiceAccountArgs) (GetServiceAccountResult, error) {
		r, err := GetServiceAccount(ctx, &v, opts...)
		return *r, err

	}).(GetServiceAccountResultOutput)
}

// GetServiceAccountApplyInput is an input type that accepts GetServiceAccountApplyArgs and GetServiceAccountApplyOutput values.
// You can construct a concrete instance of `GetServiceAccountApplyInput` via:
//
//          GetServiceAccountApplyArgs{...}
type GetServiceAccountApplyInput interface {
	pulumi.Input

	ToGetServiceAccountApplyOutput() GetServiceAccountApplyOutput
	ToGetServiceAccountApplyOutputWithContext(context.Context) GetServiceAccountApplyOutput
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountApplyArgs struct {
	// Name of the region whose AWS ELB account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetServiceAccountApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountArgs)(nil)).Elem()
}

func (i GetServiceAccountApplyArgs) ToGetServiceAccountApplyOutput() GetServiceAccountApplyOutput {
	return i.ToGetServiceAccountApplyOutputWithContext(context.Background())
}

func (i GetServiceAccountApplyArgs) ToGetServiceAccountApplyOutputWithContext(ctx context.Context) GetServiceAccountApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceAccountApplyOutput)
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountApplyOutput struct{ *pulumi.OutputState }

func (GetServiceAccountApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountArgs)(nil)).Elem()
}

func (o GetServiceAccountApplyOutput) ToGetServiceAccountApplyOutput() GetServiceAccountApplyOutput {
	return o
}

func (o GetServiceAccountApplyOutput) ToGetServiceAccountApplyOutputWithContext(ctx context.Context) GetServiceAccountApplyOutput {
	return o
}

// Name of the region whose AWS ELB account ID is desired.
// Defaults to the region from the AWS provider configuration.
func (o GetServiceAccountApplyOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceAccountArgs) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountResult)(nil)).Elem()
}

func (o GetServiceAccountResultOutput) ToGetServiceAccountResultOutput() GetServiceAccountResultOutput {
	return o
}

func (o GetServiceAccountResultOutput) ToGetServiceAccountResultOutputWithContext(ctx context.Context) GetServiceAccountResultOutput {
	return o
}

// The ARN of the AWS ELB service account in the selected region.
func (o GetServiceAccountResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceAccountResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceAccountResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceAccountApplyOutput{})
	pulumi.RegisterOutputType(GetServiceAccountResultOutput{})
}
