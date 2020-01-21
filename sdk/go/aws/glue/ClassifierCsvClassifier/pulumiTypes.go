// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClassifierCsvClassifier

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClassifierCsvClassifier struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn *bool `pulumi:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader *string `pulumi:"containsHeader"`
	// The delimiter used in the Csv to separate columns.
	Delimiter *string `pulumi:"delimiter"`
	// Specifies whether to trim column values. 
	DisableValueTrimming *bool `pulumi:"disableValueTrimming"`
	// A list of strings representing column names.
	Headers []string `pulumi:"headers"`
	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol *string `pulumi:"quoteSymbol"`
}

type ClassifierCsvClassifierInput interface {
	pulumi.Input

	ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput
	ToClassifierCsvClassifierOutputWithContext(context.Context) ClassifierCsvClassifierOutput
}

type ClassifierCsvClassifierArgs struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn pulumi.BoolPtrInput `pulumi:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader pulumi.StringPtrInput `pulumi:"containsHeader"`
	// The delimiter used in the Csv to separate columns.
	Delimiter pulumi.StringPtrInput `pulumi:"delimiter"`
	// Specifies whether to trim column values. 
	DisableValueTrimming pulumi.BoolPtrInput `pulumi:"disableValueTrimming"`
	// A list of strings representing column names.
	Headers pulumi.StringArrayInput `pulumi:"headers"`
	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol pulumi.StringPtrInput `pulumi:"quoteSymbol"`
}

func (ClassifierCsvClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassifierCsvClassifier)(nil)).Elem()
}

func (i ClassifierCsvClassifierArgs) ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput {
	return i.ToClassifierCsvClassifierOutputWithContext(context.Background())
}

func (i ClassifierCsvClassifierArgs) ToClassifierCsvClassifierOutputWithContext(ctx context.Context) ClassifierCsvClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassifierCsvClassifierOutput)
}

func (i ClassifierCsvClassifierArgs) ToClassifierCsvClassifierPtrOutput() ClassifierCsvClassifierPtrOutput {
	return i.ToClassifierCsvClassifierPtrOutputWithContext(context.Background())
}

func (i ClassifierCsvClassifierArgs) ToClassifierCsvClassifierPtrOutputWithContext(ctx context.Context) ClassifierCsvClassifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassifierCsvClassifierOutput).ToClassifierCsvClassifierPtrOutputWithContext(ctx)
}

type ClassifierCsvClassifierPtrInput interface {
	pulumi.Input

	ToClassifierCsvClassifierPtrOutput() ClassifierCsvClassifierPtrOutput
	ToClassifierCsvClassifierPtrOutputWithContext(context.Context) ClassifierCsvClassifierPtrOutput
}

type classifierCsvClassifierPtrType ClassifierCsvClassifierArgs

func ClassifierCsvClassifierPtr(v *ClassifierCsvClassifierArgs) ClassifierCsvClassifierPtrInput {	return (*classifierCsvClassifierPtrType)(v)
}

func (*classifierCsvClassifierPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassifierCsvClassifier)(nil)).Elem()
}

func (i *classifierCsvClassifierPtrType) ToClassifierCsvClassifierPtrOutput() ClassifierCsvClassifierPtrOutput {
	return i.ToClassifierCsvClassifierPtrOutputWithContext(context.Background())
}

func (i *classifierCsvClassifierPtrType) ToClassifierCsvClassifierPtrOutputWithContext(ctx context.Context) ClassifierCsvClassifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassifierCsvClassifierPtrOutput)
}

type ClassifierCsvClassifierOutput struct { *pulumi.OutputState }

func (ClassifierCsvClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassifierCsvClassifier)(nil)).Elem()
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput {
	return o
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierOutputWithContext(ctx context.Context) ClassifierCsvClassifierOutput {
	return o
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierPtrOutput() ClassifierCsvClassifierPtrOutput {
	return o.ToClassifierCsvClassifierPtrOutputWithContext(context.Background())
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierPtrOutputWithContext(ctx context.Context) ClassifierCsvClassifierPtrOutput {
	return o.ApplyT(func(v ClassifierCsvClassifier) *ClassifierCsvClassifier {
		return &v
	}).(ClassifierCsvClassifierPtrOutput)
}
// Enables the processing of files that contain only one column.
func (o ClassifierCsvClassifierOutput) AllowSingleColumn() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *bool { return v.AllowSingleColumn }).(pulumi.BoolPtrOutput)
}

// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
func (o ClassifierCsvClassifierOutput) ContainsHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.ContainsHeader }).(pulumi.StringPtrOutput)
}

// The delimiter used in the Csv to separate columns.
func (o ClassifierCsvClassifierOutput) Delimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.Delimiter }).(pulumi.StringPtrOutput)
}

// Specifies whether to trim column values. 
func (o ClassifierCsvClassifierOutput) DisableValueTrimming() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *bool { return v.DisableValueTrimming }).(pulumi.BoolPtrOutput)
}

// A list of strings representing column names.
func (o ClassifierCsvClassifierOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
func (o ClassifierCsvClassifierOutput) QuoteSymbol() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.QuoteSymbol }).(pulumi.StringPtrOutput)
}

type ClassifierCsvClassifierPtrOutput struct { *pulumi.OutputState}

func (ClassifierCsvClassifierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassifierCsvClassifier)(nil)).Elem()
}

func (o ClassifierCsvClassifierPtrOutput) ToClassifierCsvClassifierPtrOutput() ClassifierCsvClassifierPtrOutput {
	return o
}

func (o ClassifierCsvClassifierPtrOutput) ToClassifierCsvClassifierPtrOutputWithContext(ctx context.Context) ClassifierCsvClassifierPtrOutput {
	return o
}

func (o ClassifierCsvClassifierPtrOutput) Elem() ClassifierCsvClassifierOutput {
	return o.ApplyT(func (v *ClassifierCsvClassifier) ClassifierCsvClassifier { return *v }).(ClassifierCsvClassifierOutput)
}

// Enables the processing of files that contain only one column.
func (o ClassifierCsvClassifierPtrOutput) AllowSingleColumn() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *bool { return v.AllowSingleColumn }).(pulumi.BoolPtrOutput)
}

// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
func (o ClassifierCsvClassifierPtrOutput) ContainsHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.ContainsHeader }).(pulumi.StringPtrOutput)
}

// The delimiter used in the Csv to separate columns.
func (o ClassifierCsvClassifierPtrOutput) Delimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.Delimiter }).(pulumi.StringPtrOutput)
}

// Specifies whether to trim column values. 
func (o ClassifierCsvClassifierPtrOutput) DisableValueTrimming() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *bool { return v.DisableValueTrimming }).(pulumi.BoolPtrOutput)
}

// A list of strings representing column names.
func (o ClassifierCsvClassifierPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
func (o ClassifierCsvClassifierPtrOutput) QuoteSymbol() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClassifierCsvClassifier) *string { return v.QuoteSymbol }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClassifierCsvClassifierOutput{})
	pulumi.RegisterOutputType(ClassifierCsvClassifierPtrOutput{})
}
