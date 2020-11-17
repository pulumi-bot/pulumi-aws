// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeCommit Repository Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codecommit"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codecommit.NewRepository(ctx, "test", &codecommit.RepositoryArgs{
// 			Description:    pulumi.String("This is the Sample App Repository"),
// 			RepositoryName: pulumi.String("MyTestRepository"),
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
// Codecommit repository can be imported using repository name, e.g.
//
// ```sh
//  $ pulumi import aws:codecommit/repository:Repository imported ExistingRepo
// ```
type Repository struct {
	pulumi.CustomResourceState

	// The ARN of the repository
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp pulumi.StringOutput `pulumi:"cloneUrlHttp"`
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh pulumi.StringOutput `pulumi:"cloneUrlSsh"`
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrOutput `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the repository
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil || args.RepositoryName == nil {
		return nil, errors.New("missing required argument 'RepositoryName'")
	}
	if args == nil {
		args = &RepositoryArgs{}
	}
	var resource Repository
	err := ctx.RegisterResource("aws:codecommit/repository:Repository", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:codecommit/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// The ARN of the repository
	Arn *string `pulumi:"arn"`
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp *string `pulumi:"cloneUrlHttp"`
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh *string `pulumi:"cloneUrlSsh"`
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description *string `pulumi:"description"`
	// The ID of the repository
	RepositoryId *string `pulumi:"repositoryId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName *string `pulumi:"repositoryName"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type RepositoryState struct {
	// The ARN of the repository
	Arn pulumi.StringPtrInput
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp pulumi.StringPtrInput
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh pulumi.StringPtrInput
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrInput
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrInput
	// The ID of the repository
	RepositoryId pulumi.StringPtrInput
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description *string `pulumi:"description"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName string `pulumi:"repositoryName"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrInput
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrInput
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringInput
	// Key-value map of resource tags
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
