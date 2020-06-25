// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Config Configuration Aggregator
//
// ## Example Usage
// ### Account Based Aggregation
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cfg.NewConfigurationAggregator(ctx, "account", &cfg.ConfigurationAggregatorArgs{
// 			AccountAggregationSource: &cfg.ConfigurationAggregatorAccountAggregationSourceArgs{
// 				AccountIds: pulumi.StringArray{
// 					pulumi.String("123456789012"),
// 				},
// 				Regions: pulumi.StringArray{
// 					pulumi.String("us-west-2"),
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
// ### Organization Based Aggregation
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		organizationRole, err := iam.NewRole(ctx, "organizationRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewConfigurationAggregator(ctx, "organizationConfigurationAggregator", &cfg.ConfigurationAggregatorArgs{
// 			OrganizationAggregationSource: &cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs{
// 				AllRegions: pulumi.Bool(true),
// 				RoleArn:    organizationRole.Arn,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "organizationRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations"),
// 			Role:      organizationRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ConfigurationAggregator struct {
	pulumi.CustomResourceState

	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrOutput `pulumi:"accountAggregationSource"`
	// The ARN of the aggregator
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the configuration aggregator.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrOutput `pulumi:"organizationAggregationSource"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewConfigurationAggregator registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAggregator(ctx *pulumi.Context,
	name string, args *ConfigurationAggregatorArgs, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	if args == nil {
		args = &ConfigurationAggregatorArgs{}
	}
	var resource ConfigurationAggregator
	err := ctx.RegisterResource("aws:cfg/configurationAggregator:ConfigurationAggregator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAggregator gets an existing ConfigurationAggregator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAggregator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAggregatorState, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	var resource ConfigurationAggregator
	err := ctx.ReadResource("aws:cfg/configurationAggregator:ConfigurationAggregator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAggregator resources.
type configurationAggregatorState struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource *ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSource"`
	// The ARN of the aggregator
	Arn *string `pulumi:"arn"`
	// The name of the configuration aggregator.
	Name *string `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ConfigurationAggregatorState struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrInput
	// The ARN of the aggregator
	Arn pulumi.StringPtrInput
	// The name of the configuration aggregator.
	Name pulumi.StringPtrInput
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ConfigurationAggregatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorState)(nil)).Elem()
}

type configurationAggregatorArgs struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource *ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSource"`
	// The name of the configuration aggregator.
	Name *string `pulumi:"name"`
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationAggregator resource.
type ConfigurationAggregatorArgs struct {
	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource ConfigurationAggregatorAccountAggregationSourcePtrInput
	// The name of the configuration aggregator.
	Name pulumi.StringPtrInput
	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ConfigurationAggregatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorArgs)(nil)).Elem()
}
