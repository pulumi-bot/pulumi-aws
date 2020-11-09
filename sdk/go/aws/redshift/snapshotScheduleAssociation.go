// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultCluster, err := redshift.NewCluster(ctx, "defaultCluster", &redshift.ClusterArgs{
// 			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
// 			DatabaseName:      pulumi.String("mydb"),
// 			MasterUsername:    pulumi.String("foo"),
// 			MasterPassword:    pulumi.String("Mustbe8characters"),
// 			NodeType:          pulumi.String("dc1.large"),
// 			ClusterType:       pulumi.String("single-node"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSnapshotSchedule, err := redshift.NewSnapshotSchedule(ctx, "defaultSnapshotSchedule", &redshift.SnapshotScheduleArgs{
// 			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
// 			Definitions: pulumi.StringArray{
// 				pulumi.String("rate(12 hours)"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redshift.NewSnapshotScheduleAssociation(ctx, "defaultSnapshotScheduleAssociation", &redshift.SnapshotScheduleAssociationArgs{
// 			ClusterIdentifier:  defaultCluster.ID(),
// 			ScheduleIdentifier: defaultSnapshotSchedule.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SnapshotScheduleAssociation struct {
	pulumi.CustomResourceState

	// The cluster identifier.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The snapshot schedule identifier.
	ScheduleIdentifier pulumi.StringOutput `pulumi:"scheduleIdentifier"`
}

// NewSnapshotScheduleAssociation registers a new resource with the given unique name, arguments, and options.
func NewSnapshotScheduleAssociation(ctx *pulumi.Context,
	name string, args *SnapshotScheduleAssociationArgs, opts ...pulumi.ResourceOption) (*SnapshotScheduleAssociation, error) {
	if args == nil || args.ClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterIdentifier'")
	}
	if args == nil || args.ScheduleIdentifier == nil {
		return nil, errors.New("missing required argument 'ScheduleIdentifier'")
	}
	if args == nil {
		args = &SnapshotScheduleAssociationArgs{}
	}
	var resource SnapshotScheduleAssociation
	err := ctx.RegisterResource("aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotScheduleAssociation gets an existing SnapshotScheduleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotScheduleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotScheduleAssociationState, opts ...pulumi.ResourceOption) (*SnapshotScheduleAssociation, error) {
	var resource SnapshotScheduleAssociation
	err := ctx.ReadResource("aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotScheduleAssociation resources.
type snapshotScheduleAssociationState struct {
	// The cluster identifier.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The snapshot schedule identifier.
	ScheduleIdentifier *string `pulumi:"scheduleIdentifier"`
}

type SnapshotScheduleAssociationState struct {
	// The cluster identifier.
	ClusterIdentifier pulumi.StringPtrInput
	// The snapshot schedule identifier.
	ScheduleIdentifier pulumi.StringPtrInput
}

func (SnapshotScheduleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotScheduleAssociationState)(nil)).Elem()
}

type snapshotScheduleAssociationArgs struct {
	// The cluster identifier.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// The snapshot schedule identifier.
	ScheduleIdentifier string `pulumi:"scheduleIdentifier"`
}

// The set of arguments for constructing a SnapshotScheduleAssociation resource.
type SnapshotScheduleAssociationArgs struct {
	// The cluster identifier.
	ClusterIdentifier pulumi.StringInput
	// The snapshot schedule identifier.
	ScheduleIdentifier pulumi.StringInput
}

func (SnapshotScheduleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotScheduleAssociationArgs)(nil)).Elem()
}

type SnapshotScheduleAssociationInput interface {
	pulumi.Input

	ToSnapshotScheduleAssociationOutput() SnapshotScheduleAssociationOutput
	ToSnapshotScheduleAssociationOutputWithContext(ctx context.Context) SnapshotScheduleAssociationOutput
}

func (SnapshotScheduleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotScheduleAssociation)(nil)).Elem()
}

func (i SnapshotScheduleAssociation) ToSnapshotScheduleAssociationOutput() SnapshotScheduleAssociationOutput {
	return i.ToSnapshotScheduleAssociationOutputWithContext(context.Background())
}

func (i SnapshotScheduleAssociation) ToSnapshotScheduleAssociationOutputWithContext(ctx context.Context) SnapshotScheduleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotScheduleAssociationOutput)
}

type SnapshotScheduleAssociationOutput struct {
	*pulumi.OutputState
}

func (SnapshotScheduleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotScheduleAssociationOutput)(nil)).Elem()
}

func (o SnapshotScheduleAssociationOutput) ToSnapshotScheduleAssociationOutput() SnapshotScheduleAssociationOutput {
	return o
}

func (o SnapshotScheduleAssociationOutput) ToSnapshotScheduleAssociationOutputWithContext(ctx context.Context) SnapshotScheduleAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnapshotScheduleAssociationOutput{})
}
