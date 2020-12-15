// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Cognito User Pool Domain resource.
//
// ## Example Usage
// ### Amazon Cognito domain
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cognito.NewUserPool(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewUserPoolDomain(ctx, "main", &cognito.UserPoolDomainArgs{
// 			Domain:     pulumi.String("example-domain"),
// 			UserPoolId: example.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Custom Cognito domain
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		main, err := cognito.NewUserPoolDomain(ctx, "main", &cognito.UserPoolDomainArgs{
// 			Domain:         pulumi.String("example-domain.example.com"),
// 			CertificateArn: pulumi.Any(aws_acm_certificate.Cert.Arn),
// 			UserPoolId:     exampleUserPool.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "example.com"
// 		exampleZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
// 			Name: _opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "auth_cognito_A", &route53.RecordArgs{
// 			Name:   main.Domain,
// 			Type:   pulumi.String("A"),
// 			ZoneId: pulumi.String(exampleZone.ZoneId),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					EvaluateTargetHealth: pulumi.Bool(false),
// 					Name:                 main.CloudfrontDistributionArn,
// 					ZoneId:               pulumi.String("Z2FDTNDATAQYW2"),
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
// Cognito User Pool Domains can be imported using the `domain`, e.g.
//
// ```sh
//  $ pulumi import aws:cognito/userPoolDomain:UserPoolDomain main <domain>
// ```
type UserPoolDomain struct {
	pulumi.CustomResourceState

	// The AWS account ID for the user pool owner.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// The URL of the CloudFront distribution. This is required to generate the ALIAS `route53.Record`
	CloudfrontDistributionArn pulumi.StringOutput `pulumi:"cloudfrontDistributionArn"`
	// The domain string.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The S3 bucket where the static files for this domain are stored.
	S3Bucket pulumi.StringOutput `pulumi:"s3Bucket"`
	// The user pool ID.
	UserPoolId pulumi.StringOutput `pulumi:"userPoolId"`
	// The app version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewUserPoolDomain registers a new resource with the given unique name, arguments, and options.
func NewUserPoolDomain(ctx *pulumi.Context,
	name string, args *UserPoolDomainArgs, opts ...pulumi.ResourceOption) (*UserPoolDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	var resource UserPoolDomain
	err := ctx.RegisterResource("aws:cognito/userPoolDomain:UserPoolDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPoolDomain gets an existing UserPoolDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPoolDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolDomainState, opts ...pulumi.ResourceOption) (*UserPoolDomain, error) {
	var resource UserPoolDomain
	err := ctx.ReadResource("aws:cognito/userPoolDomain:UserPoolDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPoolDomain resources.
type userPoolDomainState struct {
	// The AWS account ID for the user pool owner.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn *string `pulumi:"certificateArn"`
	// The URL of the CloudFront distribution. This is required to generate the ALIAS `route53.Record`
	CloudfrontDistributionArn *string `pulumi:"cloudfrontDistributionArn"`
	// The domain string.
	Domain *string `pulumi:"domain"`
	// The S3 bucket where the static files for this domain are stored.
	S3Bucket *string `pulumi:"s3Bucket"`
	// The user pool ID.
	UserPoolId *string `pulumi:"userPoolId"`
	// The app version.
	Version *string `pulumi:"version"`
}

type UserPoolDomainState struct {
	// The AWS account ID for the user pool owner.
	AwsAccountId pulumi.StringPtrInput
	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn pulumi.StringPtrInput
	// The URL of the CloudFront distribution. This is required to generate the ALIAS `route53.Record`
	CloudfrontDistributionArn pulumi.StringPtrInput
	// The domain string.
	Domain pulumi.StringPtrInput
	// The S3 bucket where the static files for this domain are stored.
	S3Bucket pulumi.StringPtrInput
	// The user pool ID.
	UserPoolId pulumi.StringPtrInput
	// The app version.
	Version pulumi.StringPtrInput
}

func (UserPoolDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolDomainState)(nil)).Elem()
}

type userPoolDomainArgs struct {
	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn *string `pulumi:"certificateArn"`
	// The domain string.
	Domain string `pulumi:"domain"`
	// The user pool ID.
	UserPoolId string `pulumi:"userPoolId"`
}

// The set of arguments for constructing a UserPoolDomain resource.
type UserPoolDomainArgs struct {
	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn pulumi.StringPtrInput
	// The domain string.
	Domain pulumi.StringInput
	// The user pool ID.
	UserPoolId pulumi.StringInput
}

func (UserPoolDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolDomainArgs)(nil)).Elem()
}

type UserPoolDomainInput interface {
	pulumi.Input

	ToUserPoolDomainOutput() UserPoolDomainOutput
	ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput
}

func (*UserPoolDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolDomain)(nil))
}

func (i *UserPoolDomain) ToUserPoolDomainOutput() UserPoolDomainOutput {
	return i.ToUserPoolDomainOutputWithContext(context.Background())
}

func (i *UserPoolDomain) ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolDomainOutput)
}

type UserPoolDomainOutput struct {
	*pulumi.OutputState
}

func (UserPoolDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolDomain)(nil))
}

func (o UserPoolDomainOutput) ToUserPoolDomainOutput() UserPoolDomainOutput {
	return o
}

func (o UserPoolDomainOutput) ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserPoolDomainOutput{})
}
