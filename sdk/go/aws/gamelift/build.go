// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Gamelift Build resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewBuild(ctx, "test", &gamelift.BuildArgs{
// 			OperatingSystem: pulumi.String("WINDOWS_2012"),
// 			StorageLocation: &gamelift.BuildStorageLocationArgs{
// 				Bucket:  pulumi.Any(aws_s3_bucket.Test.Bucket),
// 				Key:     pulumi.Any(aws_s3_bucket_object.Test.Key),
// 				RoleArn: pulumi.Any(aws_iam_role.Test.Arn),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_iam_role_policy.Test,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Build struct {
	pulumi.CustomResourceState

	// Gamelift Build ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the build
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationOutput `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Version that is associated with this build.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewBuild registers a new resource with the given unique name, arguments, and options.
func NewBuild(ctx *pulumi.Context,
	name string, args *BuildArgs, opts ...pulumi.ResourceOption) (*Build, error) {
	if args == nil || args.OperatingSystem == nil {
		return nil, errors.New("missing required argument 'OperatingSystem'")
	}
	if args == nil || args.StorageLocation == nil {
		return nil, errors.New("missing required argument 'StorageLocation'")
	}
	if args == nil {
		args = &BuildArgs{}
	}
	var resource Build
	err := ctx.RegisterResource("aws:gamelift/build:Build", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuild gets an existing Build resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildState, opts ...pulumi.ResourceOption) (*Build, error) {
	var resource Build
	err := ctx.ReadResource("aws:gamelift/build:Build", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Build resources.
type buildState struct {
	// Gamelift Build ARN.
	Arn *string `pulumi:"arn"`
	// Name of the build
	Name *string `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation *BuildStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Version that is associated with this build.
	Version *string `pulumi:"version"`
}

type BuildState struct {
	// Gamelift Build ARN.
	Arn pulumi.StringPtrInput
	// Name of the build
	Name pulumi.StringPtrInput
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringPtrInput
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Version that is associated with this build.
	Version pulumi.StringPtrInput
}

func (BuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildState)(nil)).Elem()
}

type buildArgs struct {
	// Name of the build
	Name *string `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem string `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Version that is associated with this build.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Build resource.
type BuildArgs struct {
	// Name of the build
	Name pulumi.StringPtrInput
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringInput
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Version that is associated with this build.
	Version pulumi.StringPtrInput
}

func (BuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildArgs)(nil)).Elem()
}

type BuildInput interface {
	pulumi.Input

	ToBuildOutput() BuildOutput
	ToBuildOutputWithContext(ctx context.Context) BuildOutput
}

func (Build) ElementType() reflect.Type {
	return reflect.TypeOf((*Build)(nil)).Elem()
}

func (i Build) ToBuildOutput() BuildOutput {
	return i.ToBuildOutputWithContext(context.Background())
}

func (i Build) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildOutput)
}

type BuildOutput struct {
	*pulumi.OutputState
}

func (BuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildOutput)(nil)).Elem()
}

func (o BuildOutput) ToBuildOutput() BuildOutput {
	return o
}

func (o BuildOutput) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BuildOutput{})
}
