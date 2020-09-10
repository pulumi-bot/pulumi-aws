// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Repository struct {
	pulumi.CustomResourceState

	Arn            pulumi.StringOutput    `pulumi:"arn"`
	CloneUrlHttp   pulumi.StringOutput    `pulumi:"cloneUrlHttp"`
	CloneUrlSsh    pulumi.StringOutput    `pulumi:"cloneUrlSsh"`
	DefaultBranch  pulumi.StringPtrOutput `pulumi:"defaultBranch"`
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	RepositoryId   pulumi.StringOutput    `pulumi:"repositoryId"`
	RepositoryName pulumi.StringOutput    `pulumi:"repositoryName"`
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
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
	Arn            *string           `pulumi:"arn"`
	CloneUrlHttp   *string           `pulumi:"cloneUrlHttp"`
	CloneUrlSsh    *string           `pulumi:"cloneUrlSsh"`
	DefaultBranch  *string           `pulumi:"defaultBranch"`
	Description    *string           `pulumi:"description"`
	RepositoryId   *string           `pulumi:"repositoryId"`
	RepositoryName *string           `pulumi:"repositoryName"`
	Tags           map[string]string `pulumi:"tags"`
}

type RepositoryState struct {
	Arn            pulumi.StringPtrInput
	CloneUrlHttp   pulumi.StringPtrInput
	CloneUrlSsh    pulumi.StringPtrInput
	DefaultBranch  pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	RepositoryId   pulumi.StringPtrInput
	RepositoryName pulumi.StringPtrInput
	Tags           pulumi.StringMapInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	DefaultBranch  *string           `pulumi:"defaultBranch"`
	Description    *string           `pulumi:"description"`
	RepositoryName string            `pulumi:"repositoryName"`
	Tags           map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	DefaultBranch  pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	RepositoryName pulumi.StringInput
	Tags           pulumi.StringMapInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}
