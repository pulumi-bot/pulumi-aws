// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS Inspector Rules Packages data source allows access to the list of AWS
// Inspector Rules Packages which can be used by AWS Inspector within the region
// configured in the provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/inspector"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rules, err := inspector.GetRulesPackages(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		group, err := inspector.NewResourceGroup(ctx, "group", &inspector.ResourceGroupArgs{
// 			Tags: pulumi.StringMap{
// 				"test": pulumi.String("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		assessmentAssessmentTarget, err := inspector.NewAssessmentTarget(ctx, "assessmentAssessmentTarget", &inspector.AssessmentTargetArgs{
// 			ResourceGroupArn: group.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = inspector.NewAssessmentTemplate(ctx, "assessmentAssessmentTemplate", &inspector.AssessmentTemplateArgs{
// 			TargetArn:        assessmentAssessmentTarget.Arn,
// 			Duration:         pulumi.Int(60),
// 			RulesPackageArns: toPulumiStringArray(rules.Arns),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
func GetRulesPackages(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetRulesPackagesResult, error) {
	var rv GetRulesPackagesResult
	err := ctx.Invoke("aws:inspector/getRulesPackages:getRulesPackages", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getRulesPackages.
type GetRulesPackagesResult struct {
	// A list of the AWS Inspector Rules Packages arns available in the AWS region.
	Arns []string `pulumi:"arns"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetRulesPackagesApply(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetRulesPackagesResultOutput {
	return pulumi.Any(opts).ApplyT(func(v interface{}) (GetRulesPackagesResult, error) {
		r, err := GetRulesPackages(ctx, opts...)
		return *r, err

	}).(GetRulesPackagesResultOutput)
}

// A collection of values returned by getRulesPackages.
type GetRulesPackagesResultOutput struct{ *pulumi.OutputState }

func (GetRulesPackagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesPackagesResult)(nil)).Elem()
}

func (o GetRulesPackagesResultOutput) ToGetRulesPackagesResultOutput() GetRulesPackagesResultOutput {
	return o
}

func (o GetRulesPackagesResultOutput) ToGetRulesPackagesResultOutputWithContext(ctx context.Context) GetRulesPackagesResultOutput {
	return o
}

// A list of the AWS Inspector Rules Packages arns available in the AWS region.
func (o GetRulesPackagesResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRulesPackagesResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRulesPackagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesPackagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRulesPackagesResultOutput{})
}
