// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accessanalyzer

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).
//
// ## Example Usage
// ### Account Analyzer
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/accessanalyzer"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accessanalyzer.NewAnalyzer(ctx, "example", &accessanalyzer.AnalyzerArgs{
// 			AnalyzerName: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Organization Analyzer
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/accessanalyzer"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
// 			AwsServiceAccessPrincipals: pulumi.StringArray{
// 				pulumi.String("access-analyzer.amazonaws.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accessanalyzer.NewAnalyzer(ctx, "exampleAnalyzer", &accessanalyzer.AnalyzerArgs{
// 			AnalyzerName: pulumi.String("example"),
// 			Type:         pulumi.String("ORGANIZATION"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleOrganization,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Analyzer struct {
	pulumi.CustomResourceState

	// Name of the Analyzer.
	AnalyzerName pulumi.StringOutput `pulumi:"analyzerName"`
	Arn          pulumi.StringOutput `pulumi:"arn"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAnalyzer registers a new resource with the given unique name, arguments, and options.
func NewAnalyzer(ctx *pulumi.Context,
	name string, args *AnalyzerArgs, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AnalyzerName == nil {
		return nil, errors.New("invalid value for required argument 'AnalyzerName'")
	}
	var resource Analyzer
	err := ctx.RegisterResource("aws:accessanalyzer/analyzer:Analyzer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyzer gets an existing Analyzer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyzer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyzerState, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	var resource Analyzer
	err := ctx.ReadResource("aws:accessanalyzer/analyzer:Analyzer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analyzer resources.
type analyzerState struct {
	// Name of the Analyzer.
	AnalyzerName *string `pulumi:"analyzerName"`
	Arn          *string `pulumi:"arn"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
	Type *string `pulumi:"type"`
}

type AnalyzerState struct {
	// Name of the Analyzer.
	AnalyzerName pulumi.StringPtrInput
	Arn          pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrInput
}

func (AnalyzerState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerState)(nil)).Elem()
}

type analyzerArgs struct {
	// Name of the Analyzer.
	AnalyzerName string `pulumi:"analyzerName"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Analyzer resource.
type AnalyzerArgs struct {
	// Name of the Analyzer.
	AnalyzerName pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrInput
}

func (AnalyzerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerArgs)(nil)).Elem()
}
