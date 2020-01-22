// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package accessanalyzer

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/accessanalyzer_analyzer.html.markdown.
type Analyzer struct {
	pulumi.CustomResourceState

	// Name of the Analyzer.
	AnalyzerName pulumi.StringOutput `pulumi:"analyzerName"`
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Key-value mapping of resource tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAnalyzer registers a new resource with the given unique name, arguments, and options.
func NewAnalyzer(ctx *pulumi.Context,
	name string, args *AnalyzerArgs, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	if args == nil || args.AnalyzerName == nil {
		return nil, errors.New("missing required argument 'AnalyzerName'")
	}
	if args == nil {
		args = &AnalyzerArgs{}
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
	Arn *string `pulumi:"arn"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
	Type *string `pulumi:"type"`
}

type AnalyzerState struct {
	// Name of the Analyzer.
	AnalyzerName pulumi.StringPtrInput
	Arn pulumi.StringPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
	// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrInput
}

func (AnalyzerState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerState)(nil)).Elem()
}

type analyzerArgs struct {
	// Name of the Analyzer.
	AnalyzerName string `pulumi:"analyzerName"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Analyzer resource.
type AnalyzerArgs struct {
	// Name of the Analyzer.
	AnalyzerName pulumi.StringInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
	// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
	Type pulumi.StringPtrInput
}

func (AnalyzerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerArgs)(nil)).Elem()
}

