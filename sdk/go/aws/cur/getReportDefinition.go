// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cur

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information on an AWS Cost and Usage Report Definition.
//
// > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
//
// > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		reportDefinition, err := cur.LookupReportDefinition(ctx, &cur.LookupReportDefinitionArgs{
// 			ReportName: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupReportDefinition(ctx *pulumi.Context, args *LookupReportDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupReportDefinitionResult, error) {
	var rv LookupReportDefinitionResult
	err := ctx.Invoke("aws:cur/getReportDefinition:getReportDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReportDefinition.
type LookupReportDefinitionArgs struct {
	// The name of the report definition to match.
	ReportName string `pulumi:"reportName"`
}

// A collection of values returned by getReportDefinition.
type LookupReportDefinitionResult struct {
	// A list of additional artifacts.
	AdditionalArtifacts []string `pulumi:"additionalArtifacts"`
	// A list of schema elements.
	AdditionalSchemaElements []string `pulumi:"additionalSchemaElements"`
	// Preferred format for report.
	Compression string `pulumi:"compression"`
	// Preferred compression format for report.
	Format string `pulumi:"format"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	ReportName string `pulumi:"reportName"`
	// Name of customer S3 bucket.
	S3Bucket string `pulumi:"s3Bucket"`
	// Preferred report path prefix.
	S3Prefix string `pulumi:"s3Prefix"`
	// Region of customer S3 bucket.
	S3Region string `pulumi:"s3Region"`
	// The frequency on which report data are measured and displayed.
	TimeUnit string `pulumi:"timeUnit"`
}
