// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/inspector"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rules, err := inspector.GetRulesPackages(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		group, err := inspector.NewResourceGroup(ctx, "group", &inspector.ResourceGroupArgs{
// 			Tags: map[string]interface{}{
// 				"test": "test",
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
// 			Duration:         pulumi.Int(60),
// 			RulesPackageArns: pulumi.StringArray(rules.Arns),
// 			TargetArn:        assessmentAssessmentTarget.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
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
