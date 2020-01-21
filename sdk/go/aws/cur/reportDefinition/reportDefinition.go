// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package reportDefinition

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages Cost and Usage Report Definitions.
// 
// > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
// 
// > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cur_report_definition.html.markdown.
type ReportDefinition struct {
	pulumi.CustomResourceState

	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT.
	AdditionalArtifacts pulumi.StringArrayOutput `pulumi:"additionalArtifacts"`
	// A list of schema elements. Valid values are: RESOURCES.
	AdditionalSchemaElements pulumi.StringArrayOutput `pulumi:"additionalSchemaElements"`
	// Compression format for report. Valid values are: GZIP, ZIP.
	Compression pulumi.StringOutput `pulumi:"compression"`
	// Format for report. Valid values are: textORcsv.
	Format pulumi.StringOutput `pulumi:"format"`
	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName pulumi.StringOutput `pulumi:"reportName"`
	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket pulumi.StringOutput `pulumi:"s3Bucket"`
	// Report path prefix. Limited to 256 characters.
	S3Prefix pulumi.StringPtrOutput `pulumi:"s3Prefix"`
	// Region of the existing S3 bucket to hold generated reports.
	S3Region pulumi.StringOutput `pulumi:"s3Region"`
	// The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
	TimeUnit pulumi.StringOutput `pulumi:"timeUnit"`
}

// NewReportDefinition registers a new resource with the given unique name, arguments, and options.
func NewReportDefinition(ctx *pulumi.Context,
	name string, args *ReportDefinitionArgs, opts ...pulumi.ResourceOption) (*ReportDefinition, error) {
	if args == nil || args.AdditionalSchemaElements == nil {
		return nil, errors.New("missing required argument 'AdditionalSchemaElements'")
	}
	if args == nil || args.Compression == nil {
		return nil, errors.New("missing required argument 'Compression'")
	}
	if args == nil || args.Format == nil {
		return nil, errors.New("missing required argument 'Format'")
	}
	if args == nil || args.ReportName == nil {
		return nil, errors.New("missing required argument 'ReportName'")
	}
	if args == nil || args.S3Bucket == nil {
		return nil, errors.New("missing required argument 'S3Bucket'")
	}
	if args == nil || args.S3Region == nil {
		return nil, errors.New("missing required argument 'S3Region'")
	}
	if args == nil || args.TimeUnit == nil {
		return nil, errors.New("missing required argument 'TimeUnit'")
	}
	if args == nil {
		args = &ReportDefinitionArgs{}
	}
	var resource ReportDefinition
	err := ctx.RegisterResource("aws:cur/reportDefinition:ReportDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportDefinition gets an existing ReportDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportDefinitionState, opts ...pulumi.ResourceOption) (*ReportDefinition, error) {
	var resource ReportDefinition
	err := ctx.ReadResource("aws:cur/reportDefinition:ReportDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportDefinition resources.
type reportDefinitionState struct {
	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT.
	AdditionalArtifacts []string `pulumi:"additionalArtifacts"`
	// A list of schema elements. Valid values are: RESOURCES.
	AdditionalSchemaElements []string `pulumi:"additionalSchemaElements"`
	// Compression format for report. Valid values are: GZIP, ZIP.
	Compression *string `pulumi:"compression"`
	// Format for report. Valid values are: textORcsv.
	Format *string `pulumi:"format"`
	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName *string `pulumi:"reportName"`
	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket *string `pulumi:"s3Bucket"`
	// Report path prefix. Limited to 256 characters.
	S3Prefix *string `pulumi:"s3Prefix"`
	// Region of the existing S3 bucket to hold generated reports.
	S3Region *string `pulumi:"s3Region"`
	// The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
	TimeUnit *string `pulumi:"timeUnit"`
}

type ReportDefinitionState struct {
	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT.
	AdditionalArtifacts pulumi.StringArrayInput
	// A list of schema elements. Valid values are: RESOURCES.
	AdditionalSchemaElements pulumi.StringArrayInput
	// Compression format for report. Valid values are: GZIP, ZIP.
	Compression pulumi.StringPtrInput
	// Format for report. Valid values are: textORcsv.
	Format pulumi.StringPtrInput
	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName pulumi.StringPtrInput
	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket pulumi.StringPtrInput
	// Report path prefix. Limited to 256 characters.
	S3Prefix pulumi.StringPtrInput
	// Region of the existing S3 bucket to hold generated reports.
	S3Region pulumi.StringPtrInput
	// The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
	TimeUnit pulumi.StringPtrInput
}

func (ReportDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportDefinitionState)(nil)).Elem()
}

type reportDefinitionArgs struct {
	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT.
	AdditionalArtifacts []string `pulumi:"additionalArtifacts"`
	// A list of schema elements. Valid values are: RESOURCES.
	AdditionalSchemaElements []string `pulumi:"additionalSchemaElements"`
	// Compression format for report. Valid values are: GZIP, ZIP.
	Compression string `pulumi:"compression"`
	// Format for report. Valid values are: textORcsv.
	Format string `pulumi:"format"`
	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName string `pulumi:"reportName"`
	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket string `pulumi:"s3Bucket"`
	// Report path prefix. Limited to 256 characters.
	S3Prefix *string `pulumi:"s3Prefix"`
	// Region of the existing S3 bucket to hold generated reports.
	S3Region string `pulumi:"s3Region"`
	// The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
	TimeUnit string `pulumi:"timeUnit"`
}

// The set of arguments for constructing a ReportDefinition resource.
type ReportDefinitionArgs struct {
	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT.
	AdditionalArtifacts pulumi.StringArrayInput
	// A list of schema elements. Valid values are: RESOURCES.
	AdditionalSchemaElements pulumi.StringArrayInput
	// Compression format for report. Valid values are: GZIP, ZIP.
	Compression pulumi.StringInput
	// Format for report. Valid values are: textORcsv.
	Format pulumi.StringInput
	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName pulumi.StringInput
	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket pulumi.StringInput
	// Report path prefix. Limited to 256 characters.
	S3Prefix pulumi.StringPtrInput
	// Region of the existing S3 bucket to hold generated reports.
	S3Region pulumi.StringInput
	// The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
	TimeUnit pulumi.StringInput
}

func (ReportDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportDefinitionArgs)(nil)).Elem()
}

