// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic Container Registry Repository Policy.
//
// Note that currently only one policy may be applied to a repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := ecr.NewRepository(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecr.NewRepositoryPolicy(ctx, "foopolicy", &ecr.RepositoryPolicyArgs{
// 			Repository: foo.Name,
// 			Policy:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2008-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"new policy\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": \"*\",\n", "            \"Action\": [\n", "                \"ecr:GetDownloadUrlForLayer\",\n", "                \"ecr:BatchGetImage\",\n", "                \"ecr:BatchCheckLayerAvailability\",\n", "                \"ecr:PutImage\",\n", "                \"ecr:InitiateLayerUpload\",\n", "                \"ecr:UploadLayerPart\",\n", "                \"ecr:CompleteLayerUpload\",\n", "                \"ecr:DescribeRepositories\",\n", "                \"ecr:GetRepositoryPolicy\",\n", "                \"ecr:ListImages\",\n", "                \"ecr:DeleteRepository\",\n", "                \"ecr:BatchDeleteImage\",\n", "                \"ecr:SetRepositoryPolicy\",\n", "                \"ecr:DeleteRepositoryPolicy\"\n", "            ]\n", "        }\n", "    ]\n", "}\n")),
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
// ECR Repository Policy can be imported using the repository name, e.g.
//
// ```sh
//  $ pulumi import aws:ecr/repositoryPolicy:RepositoryPolicy example example
// ```
type RepositoryPolicy struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPolicy(ctx *pulumi.Context,
	name string, args *RepositoryPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryPolicy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil {
		args = &RepositoryPolicyArgs{}
	}
	var resource RepositoryPolicy
	err := ctx.RegisterResource("aws:ecr/repositoryPolicy:RepositoryPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPolicy gets an existing RepositoryPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPolicyState, opts ...pulumi.ResourceOption) (*RepositoryPolicy, error) {
	var resource RepositoryPolicy
	err := ctx.ReadResource("aws:ecr/repositoryPolicy:RepositoryPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPolicy resources.
type repositoryPolicyState struct {
	// The policy document. This is a JSON formatted string.
	Policy *string `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository *string `pulumi:"repository"`
}

type RepositoryPolicyState struct {
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringPtrInput
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// Name of the repository to apply the policy.
	Repository pulumi.StringPtrInput
}

func (RepositoryPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyState)(nil)).Elem()
}

type repositoryPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
	// Name of the repository to apply the policy.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryPolicy resource.
type RepositoryPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
	// Name of the repository to apply the policy.
	Repository pulumi.StringInput
}

func (RepositoryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPolicyArgs)(nil)).Elem()
}
