// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package swf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SWF Domain resource.
//
// ## Example Usage
//
// To register a basic SWF domain:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/swf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := swf.NewDomain(ctx, "foo", &swf.DomainArgs{
// 			Description:                            pulumi.String("SWF Domain"),
// 			WorkflowExecutionRetentionPeriodInDays: pulumi.String("30"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the domain. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
	WorkflowExecutionRetentionPeriodInDays pulumi.StringOutput `pulumi:"workflowExecutionRetentionPeriodInDays"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil || args.WorkflowExecutionRetentionPeriodInDays == nil {
		return nil, errors.New("missing required argument 'WorkflowExecutionRetentionPeriodInDays'")
	}
	if args == nil {
		args = &DomainArgs{}
	}
	var resource Domain
	err := ctx.RegisterResource("aws:swf/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:swf/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The domain description.
	Description *string `pulumi:"description"`
	// The name of the domain. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
	WorkflowExecutionRetentionPeriodInDays *string `pulumi:"workflowExecutionRetentionPeriodInDays"`
}

type DomainState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The domain description.
	Description pulumi.StringPtrInput
	// The name of the domain. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
	WorkflowExecutionRetentionPeriodInDays pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The domain description.
	Description *string `pulumi:"description"`
	// The name of the domain. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
	WorkflowExecutionRetentionPeriodInDays string `pulumi:"workflowExecutionRetentionPeriodInDays"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The domain description.
	Description pulumi.StringPtrInput
	// The name of the domain. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
	WorkflowExecutionRetentionPeriodInDays pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (i Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOutput)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
