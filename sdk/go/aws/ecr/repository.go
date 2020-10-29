// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic Container Registry Repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecr.NewRepository(ctx, "foo", &ecr.RepositoryArgs{
// 			ImageScanningConfiguration: &ecr.RepositoryImageScanningConfigurationArgs{
// 				ScanOnPush: pulumi.Bool(true),
// 			},
// 			ImageTagMutability: pulumi.String("MUTABLE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Full ARN of the repository.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations RepositoryEncryptionConfigurationArrayOutput `pulumi:"encryptionConfigurations"`
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrOutput `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrOutput `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}
	var resource Repository
	err := ctx.RegisterResource("aws:ecr/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws:ecr/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Full ARN of the repository.
	Arn *string `pulumi:"arn"`
	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations []RepositoryEncryptionConfiguration `pulumi:"encryptionConfigurations"`
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration *RepositoryImageScanningConfiguration `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability *string `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type RepositoryState struct {
	// Full ARN of the repository.
	Arn pulumi.StringPtrInput
	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations RepositoryEncryptionConfigurationArrayInput
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrInput
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations []RepositoryEncryptionConfiguration `pulumi:"encryptionConfigurations"`
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration *RepositoryImageScanningConfiguration `pulumi:"imageScanningConfiguration"`
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability *string `pulumi:"imageTagMutability"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations RepositoryEncryptionConfigurationArrayInput
	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration RepositoryImageScanningConfigurationPtrInput
	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil)).Elem()
}

func (i Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct {
	*pulumi.OutputState
}

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryOutput)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryOutput{})
}
