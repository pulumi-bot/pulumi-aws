// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "  {\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\"Service\": \"ssm.amazonaws.com\"},\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testAttach, err := iam.NewRolePolicyAttachment(ctx, "testAttach", &iam.RolePolicyAttachmentArgs{
// 			Role:      testRole.Name,
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewActivation(ctx, "foo", &ssm.ActivationArgs{
// 			Description:       pulumi.String("Test"),
// 			IamRole:           testRole.ID(),
// 			RegistrationLimit: pulumi.Int(5),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			testAttach,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Activation struct {
	pulumi.CustomResourceState

	// The code the system generates when it processes the activation.
	ActivationCode pulumi.StringOutput `pulumi:"activationCode"`
	// The description of the resource that you want to register.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// If the current activation has expired.
	Expired pulumi.BoolOutput `pulumi:"expired"`
	// The IAM Role to attach to the managed instance.
	IamRole pulumi.StringOutput `pulumi:"iamRole"`
	// The default name of the registered managed instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of managed instances that are currently registered using this activation.
	RegistrationCount pulumi.IntOutput `pulumi:"registrationCount"`
	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit pulumi.IntPtrOutput `pulumi:"registrationLimit"`
	// A map of tags to assign to the object.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActivation registers a new resource with the given unique name, arguments, and options.
func NewActivation(ctx *pulumi.Context,
	name string, args *ActivationArgs, opts ...pulumi.ResourceOption) (*Activation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IamRole == nil {
		return nil, errors.New("invalid value for required argument 'IamRole'")
	}
	var resource Activation
	err := ctx.RegisterResource("aws:ssm/activation:Activation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivation gets an existing Activation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivationState, opts ...pulumi.ResourceOption) (*Activation, error) {
	var resource Activation
	err := ctx.ReadResource("aws:ssm/activation:Activation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Activation resources.
type activationState struct {
	// The code the system generates when it processes the activation.
	ActivationCode *string `pulumi:"activationCode"`
	// The description of the resource that you want to register.
	Description *string `pulumi:"description"`
	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate *string `pulumi:"expirationDate"`
	// If the current activation has expired.
	Expired *bool `pulumi:"expired"`
	// The IAM Role to attach to the managed instance.
	IamRole *string `pulumi:"iamRole"`
	// The default name of the registered managed instance.
	Name *string `pulumi:"name"`
	// The number of managed instances that are currently registered using this activation.
	RegistrationCount *int `pulumi:"registrationCount"`
	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit *int `pulumi:"registrationLimit"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
}

type ActivationState struct {
	// The code the system generates when it processes the activation.
	ActivationCode pulumi.StringPtrInput
	// The description of the resource that you want to register.
	Description pulumi.StringPtrInput
	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate pulumi.StringPtrInput
	// If the current activation has expired.
	Expired pulumi.BoolPtrInput
	// The IAM Role to attach to the managed instance.
	IamRole pulumi.StringPtrInput
	// The default name of the registered managed instance.
	Name pulumi.StringPtrInput
	// The number of managed instances that are currently registered using this activation.
	RegistrationCount pulumi.IntPtrInput
	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit pulumi.IntPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
}

func (ActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*activationState)(nil)).Elem()
}

type activationArgs struct {
	// The description of the resource that you want to register.
	Description *string `pulumi:"description"`
	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The IAM Role to attach to the managed instance.
	IamRole string `pulumi:"iamRole"`
	// The default name of the registered managed instance.
	Name *string `pulumi:"name"`
	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit *int `pulumi:"registrationLimit"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Activation resource.
type ActivationArgs struct {
	// The description of the resource that you want to register.
	Description pulumi.StringPtrInput
	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate pulumi.StringPtrInput
	// The IAM Role to attach to the managed instance.
	IamRole pulumi.StringInput
	// The default name of the registered managed instance.
	Name pulumi.StringPtrInput
	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit pulumi.IntPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
}

func (ActivationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activationArgs)(nil)).Elem()
}

type ActivationInput interface {
	pulumi.Input

	ToActivationOutput() ActivationOutput
	ToActivationOutputWithContext(ctx context.Context) ActivationOutput
}

func (Activation) ElementType() reflect.Type {
	return reflect.TypeOf((*Activation)(nil)).Elem()
}

func (i Activation) ToActivationOutput() ActivationOutput {
	return i.ToActivationOutputWithContext(context.Background())
}

func (i Activation) ToActivationOutputWithContext(ctx context.Context) ActivationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationOutput)
}

type ActivationOutput struct {
	*pulumi.OutputState
}

func (ActivationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivationOutput)(nil)).Elem()
}

func (o ActivationOutput) ToActivationOutput() ActivationOutput {
	return o
}

func (o ActivationOutput) ToActivationOutputWithContext(ctx context.Context) ActivationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActivationOutput{})
}
