// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Configuration struct {
	pulumi.CustomResourceState

	Arn            pulumi.StringOutput    `pulumi:"arn"`
	Data           pulumi.StringOutput    `pulumi:"data"`
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	EngineType     pulumi.StringOutput    `pulumi:"engineType"`
	EngineVersion  pulumi.StringOutput    `pulumi:"engineVersion"`
	LatestRevision pulumi.IntOutput       `pulumi:"latestRevision"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	if args == nil || args.EngineType == nil {
		return nil, errors.New("missing required argument 'EngineType'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil {
		args = &ConfigurationArgs{}
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:mq/configuration:Configuration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:mq/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	Arn            *string           `pulumi:"arn"`
	Data           *string           `pulumi:"data"`
	Description    *string           `pulumi:"description"`
	EngineType     *string           `pulumi:"engineType"`
	EngineVersion  *string           `pulumi:"engineVersion"`
	LatestRevision *int              `pulumi:"latestRevision"`
	Name           *string           `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
}

type ConfigurationState struct {
	Arn            pulumi.StringPtrInput
	Data           pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	EngineType     pulumi.StringPtrInput
	EngineVersion  pulumi.StringPtrInput
	LatestRevision pulumi.IntPtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.StringMapInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	Data          string            `pulumi:"data"`
	Description   *string           `pulumi:"description"`
	EngineType    string            `pulumi:"engineType"`
	EngineVersion string            `pulumi:"engineVersion"`
	Name          *string           `pulumi:"name"`
	Tags          map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	Data          pulumi.StringInput
	Description   pulumi.StringPtrInput
	EngineType    pulumi.StringInput
	EngineVersion pulumi.StringInput
	Name          pulumi.StringPtrInput
	Tags          pulumi.StringMapInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}
