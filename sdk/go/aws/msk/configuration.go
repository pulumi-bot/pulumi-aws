// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Configuration struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput      `pulumi:"arn"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	KafkaVersions    pulumi.StringArrayOutput `pulumi:"kafkaVersions"`
	LatestRevision   pulumi.IntOutput         `pulumi:"latestRevision"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	ServerProperties pulumi.StringOutput      `pulumi:"serverProperties"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil || args.KafkaVersions == nil {
		return nil, errors.New("missing required argument 'KafkaVersions'")
	}
	if args == nil || args.ServerProperties == nil {
		return nil, errors.New("missing required argument 'ServerProperties'")
	}
	if args == nil {
		args = &ConfigurationArgs{}
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:msk/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws:msk/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	Arn              *string  `pulumi:"arn"`
	Description      *string  `pulumi:"description"`
	KafkaVersions    []string `pulumi:"kafkaVersions"`
	LatestRevision   *int     `pulumi:"latestRevision"`
	Name             *string  `pulumi:"name"`
	ServerProperties *string  `pulumi:"serverProperties"`
}

type ConfigurationState struct {
	Arn              pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	KafkaVersions    pulumi.StringArrayInput
	LatestRevision   pulumi.IntPtrInput
	Name             pulumi.StringPtrInput
	ServerProperties pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	Description      *string  `pulumi:"description"`
	KafkaVersions    []string `pulumi:"kafkaVersions"`
	Name             *string  `pulumi:"name"`
	ServerProperties string   `pulumi:"serverProperties"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	Description      pulumi.StringPtrInput
	KafkaVersions    pulumi.StringArrayInput
	Name             pulumi.StringPtrInput
	ServerProperties pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}
