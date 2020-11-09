// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeArtifact Domain Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codeartifact"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
// 			Description: pulumi.String("domain key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codeartifact.NewDomain(ctx, "exampleDomain", &codeartifact.DomainArgs{
// 			Domain:        pulumi.String("example"),
// 			EncryptionKey: exampleKey.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The ARN of Domain.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The total size of all assets in the domain.
	AssetSizeBytes pulumi.IntOutput `pulumi:"assetSizeBytes"`
	// A timestamp that represents the date and time the domain was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The name of the domain to create. All domain names in an AWS Region that are in the same AWS account must be unique. The domain name is used as the prefix in DNS hostnames. Do not use sensitive information in a domain name because it is publicly discoverable.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The encryption key for the domain. This is used to encrypt content stored in a domain. The KMS Key Amazon Resource Name (ARN).
	EncryptionKey pulumi.StringOutput `pulumi:"encryptionKey"`
	// The AWS account ID that owns the domain.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The number of repositories in the domain.
	RepositoryCount pulumi.IntOutput `pulumi:"repositoryCount"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.EncryptionKey == nil {
		return nil, errors.New("missing required argument 'EncryptionKey'")
	}
	if args == nil {
		args = &DomainArgs{}
	}
	var resource Domain
	err := ctx.RegisterResource("aws:codeartifact/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:codeartifact/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The ARN of Domain.
	Arn *string `pulumi:"arn"`
	// The total size of all assets in the domain.
	AssetSizeBytes *int `pulumi:"assetSizeBytes"`
	// A timestamp that represents the date and time the domain was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedTime *string `pulumi:"createdTime"`
	// The name of the domain to create. All domain names in an AWS Region that are in the same AWS account must be unique. The domain name is used as the prefix in DNS hostnames. Do not use sensitive information in a domain name because it is publicly discoverable.
	Domain *string `pulumi:"domain"`
	// The encryption key for the domain. This is used to encrypt content stored in a domain. The KMS Key Amazon Resource Name (ARN).
	EncryptionKey *string `pulumi:"encryptionKey"`
	// The AWS account ID that owns the domain.
	Owner *string `pulumi:"owner"`
	// The number of repositories in the domain.
	RepositoryCount *int `pulumi:"repositoryCount"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type DomainState struct {
	// The ARN of Domain.
	Arn pulumi.StringPtrInput
	// The total size of all assets in the domain.
	AssetSizeBytes pulumi.IntPtrInput
	// A timestamp that represents the date and time the domain was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedTime pulumi.StringPtrInput
	// The name of the domain to create. All domain names in an AWS Region that are in the same AWS account must be unique. The domain name is used as the prefix in DNS hostnames. Do not use sensitive information in a domain name because it is publicly discoverable.
	Domain pulumi.StringPtrInput
	// The encryption key for the domain. This is used to encrypt content stored in a domain. The KMS Key Amazon Resource Name (ARN).
	EncryptionKey pulumi.StringPtrInput
	// The AWS account ID that owns the domain.
	Owner pulumi.StringPtrInput
	// The number of repositories in the domain.
	RepositoryCount pulumi.IntPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The name of the domain to create. All domain names in an AWS Region that are in the same AWS account must be unique. The domain name is used as the prefix in DNS hostnames. Do not use sensitive information in a domain name because it is publicly discoverable.
	Domain string `pulumi:"domain"`
	// The encryption key for the domain. This is used to encrypt content stored in a domain. The KMS Key Amazon Resource Name (ARN).
	EncryptionKey string `pulumi:"encryptionKey"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of the domain to create. All domain names in an AWS Region that are in the same AWS account must be unique. The domain name is used as the prefix in DNS hostnames. Do not use sensitive information in a domain name because it is publicly discoverable.
	Domain pulumi.StringInput
	// The encryption key for the domain. This is used to encrypt content stored in a domain. The KMS Key Amazon Resource Name (ARN).
	EncryptionKey pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (i Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOutput)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
