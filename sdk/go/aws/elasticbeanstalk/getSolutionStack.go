// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the name of a elastic beanstalk solution stack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticbeanstalk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := elasticbeanstalk.GetSolutionStack(ctx, &elasticbeanstalk.GetSolutionStackArgs{
// 			MostRecent: &opt0,
// 			NameRegex:  fmt.Sprintf("%v%v", "^64bit Amazon Linux (.*) Multi-container Docker (.*)", "$"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSolutionStack(ctx *pulumi.Context, args *GetSolutionStackArgs, opts ...pulumi.InvokeOption) (*GetSolutionStackResult, error) {
	var rv GetSolutionStackResult
	err := ctx.Invoke("aws:elasticbeanstalk/getSolutionStack:getSolutionStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSolutionStack.
type GetSolutionStackArgs struct {
	// If more than one result is returned, use the most
	// recent solution stack.
	MostRecent *bool `pulumi:"mostRecent"`
	// A regex string to apply to the solution stack list returned
	// by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from
	// AWS documentation for reference solution stack names.
	NameRegex string `pulumi:"nameRegex"`
}

// A collection of values returned by getSolutionStack.
type GetSolutionStackResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// The name of the solution stack.
	Name      string `pulumi:"name"`
	NameRegex string `pulumi:"nameRegex"`
}

func GetSolutionStackApply(ctx *pulumi.Context, args GetSolutionStackApplyInput, opts ...pulumi.InvokeOption) GetSolutionStackResultOutput {
	return args.ToGetSolutionStackApplyOutput().ApplyT(func(v GetSolutionStackArgs) (GetSolutionStackResult, error) {
		r, err := GetSolutionStack(ctx, &v, opts...)
		return *r, err

	}).(GetSolutionStackResultOutput)
}

// GetSolutionStackApplyInput is an input type that accepts GetSolutionStackApplyArgs and GetSolutionStackApplyOutput values.
// You can construct a concrete instance of `GetSolutionStackApplyInput` via:
//
//          GetSolutionStackApplyArgs{...}
type GetSolutionStackApplyInput interface {
	pulumi.Input

	ToGetSolutionStackApplyOutput() GetSolutionStackApplyOutput
	ToGetSolutionStackApplyOutputWithContext(context.Context) GetSolutionStackApplyOutput
}

// A collection of arguments for invoking getSolutionStack.
type GetSolutionStackApplyArgs struct {
	// If more than one result is returned, use the most
	// recent solution stack.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// A regex string to apply to the solution stack list returned
	// by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from
	// AWS documentation for reference solution stack names.
	NameRegex pulumi.StringInput `pulumi:"nameRegex"`
}

func (GetSolutionStackApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSolutionStackArgs)(nil)).Elem()
}

func (i GetSolutionStackApplyArgs) ToGetSolutionStackApplyOutput() GetSolutionStackApplyOutput {
	return i.ToGetSolutionStackApplyOutputWithContext(context.Background())
}

func (i GetSolutionStackApplyArgs) ToGetSolutionStackApplyOutputWithContext(ctx context.Context) GetSolutionStackApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSolutionStackApplyOutput)
}

// A collection of arguments for invoking getSolutionStack.
type GetSolutionStackApplyOutput struct{ *pulumi.OutputState }

func (GetSolutionStackApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSolutionStackArgs)(nil)).Elem()
}

func (o GetSolutionStackApplyOutput) ToGetSolutionStackApplyOutput() GetSolutionStackApplyOutput {
	return o
}

func (o GetSolutionStackApplyOutput) ToGetSolutionStackApplyOutputWithContext(ctx context.Context) GetSolutionStackApplyOutput {
	return o
}

// If more than one result is returned, use the most
// recent solution stack.
func (o GetSolutionStackApplyOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSolutionStackArgs) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// A regex string to apply to the solution stack list returned
// by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from
// AWS documentation for reference solution stack names.
func (o GetSolutionStackApplyOutput) NameRegex() pulumi.StringOutput {
	return o.ApplyT(func(v GetSolutionStackArgs) string { return v.NameRegex }).(pulumi.StringOutput)
}

// A collection of values returned by getSolutionStack.
type GetSolutionStackResultOutput struct{ *pulumi.OutputState }

func (GetSolutionStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSolutionStackResult)(nil)).Elem()
}

func (o GetSolutionStackResultOutput) ToGetSolutionStackResultOutput() GetSolutionStackResultOutput {
	return o
}

func (o GetSolutionStackResultOutput) ToGetSolutionStackResultOutputWithContext(ctx context.Context) GetSolutionStackResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSolutionStackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSolutionStackResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSolutionStackResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSolutionStackResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// The name of the solution stack.
func (o GetSolutionStackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSolutionStackResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSolutionStackResultOutput) NameRegex() pulumi.StringOutput {
	return o.ApplyT(func(v GetSolutionStackResult) string { return v.NameRegex }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSolutionStackApplyOutput{})
	pulumi.RegisterOutputType(GetSolutionStackResultOutput{})
}
