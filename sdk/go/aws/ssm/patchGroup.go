// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SSM Patch Group resource
//
// ## Example Usage
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
// 		production, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
// 			ApprovedPatches: pulumi.StringArray{
// 				pulumi.String("KB123456"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewPatchGroup(ctx, "patchgroup", &ssm.PatchGroupArgs{
// 			BaselineId: production.ID(),
// 			PatchGroup: pulumi.String("patch-group-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PatchGroup struct {
	pulumi.CustomResourceState

	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringOutput `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringOutput `pulumi:"patchGroup"`
}

// NewPatchGroup registers a new resource with the given unique name, arguments, and options.
func NewPatchGroup(ctx *pulumi.Context,
	name string, args *PatchGroupArgs, opts ...pulumi.ResourceOption) (*PatchGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.BaselineId == nil {
		return nil, errors.New("invalid value for required argument 'BaselineId'")
	}
	if args.PatchGroup == nil {
		return nil, errors.New("invalid value for required argument 'PatchGroup'")
	}
	var resource PatchGroup
	err := ctx.RegisterResource("aws:ssm/patchGroup:PatchGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchGroup gets an existing PatchGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchGroupState, opts ...pulumi.ResourceOption) (*PatchGroup, error) {
	var resource PatchGroup
	err := ctx.ReadResource("aws:ssm/patchGroup:PatchGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchGroup resources.
type patchGroupState struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId *string `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup *string `pulumi:"patchGroup"`
}

type PatchGroupState struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringPtrInput
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringPtrInput
}

func (PatchGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchGroupState)(nil)).Elem()
}

type patchGroupArgs struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId string `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup string `pulumi:"patchGroup"`
}

// The set of arguments for constructing a PatchGroup resource.
type PatchGroupArgs struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringInput
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringInput
}

func (PatchGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchGroupArgs)(nil)).Elem()
}
