// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serverlessrepository

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an AWS Serverless Application Repository application. For example, this can be used to determine the required `capabilities` for an application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/serverlessrepository"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := serverlessrepository.GetApplication(ctx, &serverlessrepository.GetApplicationArgs{
// 			ApplicationId: "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serverlessrepository.NewCloudFormationStack(ctx, "exampleCloudFormationStack", &serverlessrepository.CloudFormationStackArgs{
// 			ApplicationId:   pulumi.String(exampleApplication.ApplicationId),
// 			SemanticVersion: pulumi.String(exampleApplication.SemanticVersion),
// 			Capabilities:    toPulumiStringArray(exampleApplication.RequiredCapabilities),
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
func GetApplication(ctx *pulumi.Context, args *GetApplicationArgs, opts ...pulumi.InvokeOption) (*GetApplicationResult, error) {
	var rv GetApplicationResult
	err := ctx.Invoke("aws:serverlessrepository/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type GetApplicationArgs struct {
	// The ARN of the application.
	ApplicationId string `pulumi:"applicationId"`
	// The requested version of the application. By default, retrieves the latest version.
	SemanticVersion *string `pulumi:"semanticVersion"`
}

// A collection of values returned by getApplication.
type GetApplicationResult struct {
	// The ARN of the application.
	ApplicationId string `pulumi:"applicationId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the application.
	Name string `pulumi:"name"`
	// A list of capabilities describing the permissions needed to deploy the application.
	RequiredCapabilities []string `pulumi:"requiredCapabilities"`
	SemanticVersion      string   `pulumi:"semanticVersion"`
	// A URL pointing to the source code of the application version.
	SourceCodeUrl string `pulumi:"sourceCodeUrl"`
	// A URL pointing to the Cloud Formation template for the application version.
	TemplateUrl string `pulumi:"templateUrl"`
}

func GetApplicationOutput(ctx *pulumi.Context, args GetApplicationOutputArgs, opts ...pulumi.InvokeOption) GetApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationResult, error) {
			args := v.(GetApplicationArgs)
			r, err := GetApplication(ctx, &args, opts...)
			return *r, err
		}).(GetApplicationResultOutput)
}

// A collection of arguments for invoking getApplication.
type GetApplicationOutputArgs struct {
	// The ARN of the application.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// The requested version of the application. By default, retrieves the latest version.
	SemanticVersion pulumi.StringPtrInput `pulumi:"semanticVersion"`
}

func (GetApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationArgs)(nil)).Elem()
}

// A collection of values returned by getApplication.
type GetApplicationResultOutput struct{ *pulumi.OutputState }

func (GetApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationResult)(nil)).Elem()
}

func (o GetApplicationResultOutput) ToGetApplicationResultOutput() GetApplicationResultOutput {
	return o
}

func (o GetApplicationResultOutput) ToGetApplicationResultOutputWithContext(ctx context.Context) GetApplicationResultOutput {
	return o
}

// The ARN of the application.
func (o GetApplicationResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the application.
func (o GetApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of capabilities describing the permissions needed to deploy the application.
func (o GetApplicationResultOutput) RequiredCapabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationResult) []string { return v.RequiredCapabilities }).(pulumi.StringArrayOutput)
}

func (o GetApplicationResultOutput) SemanticVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.SemanticVersion }).(pulumi.StringOutput)
}

// A URL pointing to the source code of the application version.
func (o GetApplicationResultOutput) SourceCodeUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.SourceCodeUrl }).(pulumi.StringOutput)
}

// A URL pointing to the Cloud Formation template for the application version.
func (o GetApplicationResultOutput) TemplateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationResult) string { return v.TemplateUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationResultOutput{})
}
