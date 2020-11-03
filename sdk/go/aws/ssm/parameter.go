// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SSM Parameter resource.
//
// ## Example Usage
//
// To store a basic string parameter:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewParameter(ctx, "foo", &ssm.ParameterArgs{
// 			Type:  pulumi.String("String"),
// 			Value: pulumi.String("bar"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To store an encrypted string using the default SSM KMS key:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewInstance(ctx, "_default", &rds.InstanceArgs{
// 			AllocatedStorage:   pulumi.Int(10),
// 			StorageType:        pulumi.String("gp2"),
// 			Engine:             pulumi.String("mysql"),
// 			EngineVersion:      pulumi.String("5.7.16"),
// 			InstanceClass:      pulumi.String("db.t2.micro"),
// 			Name:               pulumi.String("mydb"),
// 			Username:           pulumi.String("foo"),
// 			Password:           pulumi.Any(_var.Database_master_password),
// 			DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
// 			ParameterGroupName: pulumi.String("default.mysql5.7"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewParameter(ctx, "secret", &ssm.ParameterArgs{
// 			Description: pulumi.String("The parameter description"),
// 			Type:        pulumi.String("SecureString"),
// 			Value:       pulumi.Any(_var.Database_master_password),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Note:** The unencrypted value of a SecureString will be stored in the raw state as plain-text.
type Parameter struct {
	pulumi.CustomResourceState

	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrOutput `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// The description of the parameter.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringOutput `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrOutput `pulumi:"overwrite"`
	// A map of tags to assign to the object.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the parameter.
	Value pulumi.StringOutput `pulumi:"value"`
	// The version of the parameter.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewParameter registers a new resource with the given unique name, arguments, and options.
func NewParameter(ctx *pulumi.Context,
	name string, args *ParameterArgs, opts ...pulumi.ResourceOption) (*Parameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Parameter
	err := ctx.RegisterResource("aws:ssm/parameter:Parameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameter gets an existing Parameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterState, opts ...pulumi.ResourceOption) (*Parameter, error) {
	var resource Parameter
	err := ctx.ReadResource("aws:ssm/parameter:Parameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameter resources.
type parameterState struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType *string `pulumi:"dataType"`
	// The description of the parameter.
	Description *string `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite *bool `pulumi:"overwrite"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type *string `pulumi:"type"`
	// The value of the parameter.
	Value *string `pulumi:"value"`
	// The version of the parameter.
	Version *int `pulumi:"version"`
}

type ParameterState struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// The ARN of the parameter.
	Arn pulumi.StringPtrInput
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringPtrInput
	// The description of the parameter.
	Description pulumi.StringPtrInput
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringPtrInput
	// The value of the parameter.
	Value pulumi.StringPtrInput
	// The version of the parameter.
	Version pulumi.IntPtrInput
}

func (ParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterState)(nil)).Elem()
}

type parameterArgs struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType *string `pulumi:"dataType"`
	// The description of the parameter.
	Description *string `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite *bool `pulumi:"overwrite"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type string `pulumi:"type"`
	// The value of the parameter.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Parameter resource.
type ParameterArgs struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// The ARN of the parameter.
	Arn pulumi.StringPtrInput
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringPtrInput
	// The description of the parameter.
	Description pulumi.StringPtrInput
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringInput
	// The value of the parameter.
	Value pulumi.StringInput
}

func (ParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterArgs)(nil)).Elem()
}
