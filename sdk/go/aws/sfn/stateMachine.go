// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Step Function State Machine resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
// 			RoleArn:    pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
// 			Definition: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Comment\": \"A Hello World example of the Amazon States Language using an AWS Lambda Function\",\n", "  \"StartAt\": \"HelloWorld\",\n", "  \"States\": {\n", "    \"HelloWorld\": {\n", "      \"Type\": \"Task\",\n", "      \"Resource\": \"", aws_lambda_function.Lambda.Arn, "\",\n", "      \"End\": true\n", "    }\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type StateMachine struct {
	pulumi.CustomResourceState

	// The ARN of the state machine.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date the state machine was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringOutput `pulumi:"definition"`
	// The name of the state machine.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewStateMachine registers a new resource with the given unique name, arguments, and options.
func NewStateMachine(ctx *pulumi.Context,
	name string, args *StateMachineArgs, opts ...pulumi.ResourceOption) (*StateMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource StateMachine
	err := ctx.RegisterResource("aws:sfn/stateMachine:StateMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStateMachine gets an existing StateMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStateMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StateMachineState, opts ...pulumi.ResourceOption) (*StateMachine, error) {
	var resource StateMachine
	err := ctx.ReadResource("aws:sfn/stateMachine:StateMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StateMachine resources.
type stateMachineState struct {
	// The ARN of the state machine.
	Arn *string `pulumi:"arn"`
	// The date the state machine was created.
	CreationDate *string `pulumi:"creationDate"`
	// The Amazon States Language definition of the state machine.
	Definition *string `pulumi:"definition"`
	// The name of the state machine.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn *string `pulumi:"roleArn"`
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status *string `pulumi:"status"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type StateMachineState struct {
	// The ARN of the state machine.
	Arn pulumi.StringPtrInput
	// The date the state machine was created.
	CreationDate pulumi.StringPtrInput
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringPtrInput
	// The name of the state machine.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringPtrInput
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (StateMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*stateMachineState)(nil)).Elem()
}

type stateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition string `pulumi:"definition"`
	// The name of the state machine.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn string `pulumi:"roleArn"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StateMachine resource.
type StateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringInput
	// The name of the state machine.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (StateMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stateMachineArgs)(nil)).Elem()
}

type StateMachineInput interface {
	pulumi.Input

	ToStateMachineOutput() StateMachineOutput
	ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput
}

func (StateMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachine)(nil)).Elem()
}

func (i StateMachine) ToStateMachineOutput() StateMachineOutput {
	return i.ToStateMachineOutputWithContext(context.Background())
}

func (i StateMachine) ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineOutput)
}

type StateMachineOutput struct {
	*pulumi.OutputState
}

func (StateMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineOutput)(nil)).Elem()
}

func (o StateMachineOutput) ToStateMachineOutput() StateMachineOutput {
	return o
}

func (o StateMachineOutput) ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StateMachineOutput{})
}
