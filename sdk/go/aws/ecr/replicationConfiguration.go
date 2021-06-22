// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic Container Registry Replication Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRegions, err := aws.GetRegions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecr.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &ecr.ReplicationConfigurationArgs{
// 			ReplicationConfiguration: &ecr.ReplicationConfigurationReplicationConfigurationArgs{
// 				Rule: &ecr.ReplicationConfigurationReplicationConfigurationRuleArgs{
// 					Destinations: []ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
// 						&ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
// 							Region:     pulumi.String(exampleRegions.Names[0]),
// 							RegistryId: pulumi.String(current.AccountId),
// 						},
// 					},
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
//
// ## Import
//
// ECR Replication Configuration can be imported using the `registry_id`, e.g.
//
// ```sh
//  $ pulumi import aws:ecr/replicationConfiguration:ReplicationConfiguration service 012345678912
// ```
type ReplicationConfiguration struct {
	pulumi.CustomResourceState

	// The account ID of the destination registry to replicate to.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration ReplicationConfigurationReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
}

// NewReplicationConfiguration registers a new resource with the given unique name, arguments, and options.
func NewReplicationConfiguration(ctx *pulumi.Context,
	name string, args *ReplicationConfigurationArgs, opts ...pulumi.ResourceOption) (*ReplicationConfiguration, error) {
	if args == nil {
		args = &ReplicationConfigurationArgs{}
	}

	var resource ReplicationConfiguration
	err := ctx.RegisterResource("aws:ecr/replicationConfiguration:ReplicationConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationConfiguration gets an existing ReplicationConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationConfigurationState, opts ...pulumi.ResourceOption) (*ReplicationConfiguration, error) {
	var resource ReplicationConfiguration
	err := ctx.ReadResource("aws:ecr/replicationConfiguration:ReplicationConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationConfiguration resources.
type replicationConfigurationState struct {
	// The account ID of the destination registry to replicate to.
	RegistryId *string `pulumi:"registryId"`
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration *ReplicationConfigurationReplicationConfiguration `pulumi:"replicationConfiguration"`
}

type ReplicationConfigurationState struct {
	// The account ID of the destination registry to replicate to.
	RegistryId pulumi.StringPtrInput
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration ReplicationConfigurationReplicationConfigurationPtrInput
}

func (ReplicationConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigurationState)(nil)).Elem()
}

type replicationConfigurationArgs struct {
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration *ReplicationConfigurationReplicationConfiguration `pulumi:"replicationConfiguration"`
}

// The set of arguments for constructing a ReplicationConfiguration resource.
type ReplicationConfigurationArgs struct {
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration ReplicationConfigurationReplicationConfigurationPtrInput
}

func (ReplicationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigurationArgs)(nil)).Elem()
}

type ReplicationConfigurationInput interface {
	pulumi.Input

	ToReplicationConfigurationOutput() ReplicationConfigurationOutput
	ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput
}

func (*ReplicationConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationConfiguration)(nil))
}

func (i *ReplicationConfiguration) ToReplicationConfigurationOutput() ReplicationConfigurationOutput {
	return i.ToReplicationConfigurationOutputWithContext(context.Background())
}

func (i *ReplicationConfiguration) ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationOutput)
}

func (i *ReplicationConfiguration) ToReplicationConfigurationPtrOutput() ReplicationConfigurationPtrOutput {
	return i.ToReplicationConfigurationPtrOutputWithContext(context.Background())
}

func (i *ReplicationConfiguration) ToReplicationConfigurationPtrOutputWithContext(ctx context.Context) ReplicationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationPtrOutput)
}

type ReplicationConfigurationPtrInput interface {
	pulumi.Input

	ToReplicationConfigurationPtrOutput() ReplicationConfigurationPtrOutput
	ToReplicationConfigurationPtrOutputWithContext(ctx context.Context) ReplicationConfigurationPtrOutput
}

type replicationConfigurationPtrType ReplicationConfigurationArgs

func (*replicationConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfiguration)(nil))
}

func (i *replicationConfigurationPtrType) ToReplicationConfigurationPtrOutput() ReplicationConfigurationPtrOutput {
	return i.ToReplicationConfigurationPtrOutputWithContext(context.Background())
}

func (i *replicationConfigurationPtrType) ToReplicationConfigurationPtrOutputWithContext(ctx context.Context) ReplicationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationPtrOutput)
}

// ReplicationConfigurationArrayInput is an input type that accepts ReplicationConfigurationArray and ReplicationConfigurationArrayOutput values.
// You can construct a concrete instance of `ReplicationConfigurationArrayInput` via:
//
//          ReplicationConfigurationArray{ ReplicationConfigurationArgs{...} }
type ReplicationConfigurationArrayInput interface {
	pulumi.Input

	ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput
	ToReplicationConfigurationArrayOutputWithContext(context.Context) ReplicationConfigurationArrayOutput
}

type ReplicationConfigurationArray []ReplicationConfigurationInput

func (ReplicationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReplicationConfiguration)(nil))
}

func (i ReplicationConfigurationArray) ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput {
	return i.ToReplicationConfigurationArrayOutputWithContext(context.Background())
}

func (i ReplicationConfigurationArray) ToReplicationConfigurationArrayOutputWithContext(ctx context.Context) ReplicationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationArrayOutput)
}

// ReplicationConfigurationMapInput is an input type that accepts ReplicationConfigurationMap and ReplicationConfigurationMapOutput values.
// You can construct a concrete instance of `ReplicationConfigurationMapInput` via:
//
//          ReplicationConfigurationMap{ "key": ReplicationConfigurationArgs{...} }
type ReplicationConfigurationMapInput interface {
	pulumi.Input

	ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput
	ToReplicationConfigurationMapOutputWithContext(context.Context) ReplicationConfigurationMapOutput
}

type ReplicationConfigurationMap map[string]ReplicationConfigurationInput

func (ReplicationConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReplicationConfiguration)(nil))
}

func (i ReplicationConfigurationMap) ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput {
	return i.ToReplicationConfigurationMapOutputWithContext(context.Background())
}

func (i ReplicationConfigurationMap) ToReplicationConfigurationMapOutputWithContext(ctx context.Context) ReplicationConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationMapOutput)
}

type ReplicationConfigurationOutput struct {
	*pulumi.OutputState
}

func (ReplicationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationConfiguration)(nil))
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationOutput() ReplicationConfigurationOutput {
	return o
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput {
	return o
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationPtrOutput() ReplicationConfigurationPtrOutput {
	return o.ToReplicationConfigurationPtrOutputWithContext(context.Background())
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationPtrOutputWithContext(ctx context.Context) ReplicationConfigurationPtrOutput {
	return o.ApplyT(func(v ReplicationConfiguration) *ReplicationConfiguration {
		return &v
	}).(ReplicationConfigurationPtrOutput)
}

type ReplicationConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (ReplicationConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfiguration)(nil))
}

func (o ReplicationConfigurationPtrOutput) ToReplicationConfigurationPtrOutput() ReplicationConfigurationPtrOutput {
	return o
}

func (o ReplicationConfigurationPtrOutput) ToReplicationConfigurationPtrOutputWithContext(ctx context.Context) ReplicationConfigurationPtrOutput {
	return o
}

type ReplicationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ReplicationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicationConfiguration)(nil))
}

func (o ReplicationConfigurationArrayOutput) ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput {
	return o
}

func (o ReplicationConfigurationArrayOutput) ToReplicationConfigurationArrayOutputWithContext(ctx context.Context) ReplicationConfigurationArrayOutput {
	return o
}

func (o ReplicationConfigurationArrayOutput) Index(i pulumi.IntInput) ReplicationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicationConfiguration {
		return vs[0].([]ReplicationConfiguration)[vs[1].(int)]
	}).(ReplicationConfigurationOutput)
}

type ReplicationConfigurationMapOutput struct{ *pulumi.OutputState }

func (ReplicationConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicationConfiguration)(nil))
}

func (o ReplicationConfigurationMapOutput) ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput {
	return o
}

func (o ReplicationConfigurationMapOutput) ToReplicationConfigurationMapOutputWithContext(ctx context.Context) ReplicationConfigurationMapOutput {
	return o
}

func (o ReplicationConfigurationMapOutput) MapIndex(k pulumi.StringInput) ReplicationConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicationConfiguration {
		return vs[0].(map[string]ReplicationConfiguration)[vs[1].(string)]
	}).(ReplicationConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationConfigurationOutput{})
	pulumi.RegisterOutputType(ReplicationConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ReplicationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ReplicationConfigurationMapOutput{})
}
