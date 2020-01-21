// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package CrawlerCatalogTarget

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CrawlerCatalogTarget struct {
	// The name of the Glue database to be synchronized.
	DatabaseName string `pulumi:"databaseName"`
	// A list of catalog tables to be synchronized.
	Tables []string `pulumi:"tables"`
}

type CrawlerCatalogTargetInput interface {
	pulumi.Input

	ToCrawlerCatalogTargetOutput() CrawlerCatalogTargetOutput
	ToCrawlerCatalogTargetOutputWithContext(context.Context) CrawlerCatalogTargetOutput
}

type CrawlerCatalogTargetArgs struct {
	// The name of the Glue database to be synchronized.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// A list of catalog tables to be synchronized.
	Tables pulumi.StringArrayInput `pulumi:"tables"`
}

func (CrawlerCatalogTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrawlerCatalogTarget)(nil)).Elem()
}

func (i CrawlerCatalogTargetArgs) ToCrawlerCatalogTargetOutput() CrawlerCatalogTargetOutput {
	return i.ToCrawlerCatalogTargetOutputWithContext(context.Background())
}

func (i CrawlerCatalogTargetArgs) ToCrawlerCatalogTargetOutputWithContext(ctx context.Context) CrawlerCatalogTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrawlerCatalogTargetOutput)
}

type CrawlerCatalogTargetArrayInput interface {
	pulumi.Input

	ToCrawlerCatalogTargetArrayOutput() CrawlerCatalogTargetArrayOutput
	ToCrawlerCatalogTargetArrayOutputWithContext(context.Context) CrawlerCatalogTargetArrayOutput
}

type CrawlerCatalogTargetArray []CrawlerCatalogTargetInput

func (CrawlerCatalogTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CrawlerCatalogTarget)(nil)).Elem()
}

func (i CrawlerCatalogTargetArray) ToCrawlerCatalogTargetArrayOutput() CrawlerCatalogTargetArrayOutput {
	return i.ToCrawlerCatalogTargetArrayOutputWithContext(context.Background())
}

func (i CrawlerCatalogTargetArray) ToCrawlerCatalogTargetArrayOutputWithContext(ctx context.Context) CrawlerCatalogTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrawlerCatalogTargetArrayOutput)
}

type CrawlerCatalogTargetOutput struct { *pulumi.OutputState }

func (CrawlerCatalogTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrawlerCatalogTarget)(nil)).Elem()
}

func (o CrawlerCatalogTargetOutput) ToCrawlerCatalogTargetOutput() CrawlerCatalogTargetOutput {
	return o
}

func (o CrawlerCatalogTargetOutput) ToCrawlerCatalogTargetOutputWithContext(ctx context.Context) CrawlerCatalogTargetOutput {
	return o
}

// The name of the Glue database to be synchronized.
func (o CrawlerCatalogTargetOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func (v CrawlerCatalogTarget) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// A list of catalog tables to be synchronized.
func (o CrawlerCatalogTargetOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func (v CrawlerCatalogTarget) []string { return v.Tables }).(pulumi.StringArrayOutput)
}

type CrawlerCatalogTargetArrayOutput struct { *pulumi.OutputState}

func (CrawlerCatalogTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CrawlerCatalogTarget)(nil)).Elem()
}

func (o CrawlerCatalogTargetArrayOutput) ToCrawlerCatalogTargetArrayOutput() CrawlerCatalogTargetArrayOutput {
	return o
}

func (o CrawlerCatalogTargetArrayOutput) ToCrawlerCatalogTargetArrayOutputWithContext(ctx context.Context) CrawlerCatalogTargetArrayOutput {
	return o
}

func (o CrawlerCatalogTargetArrayOutput) Index(i pulumi.IntInput) CrawlerCatalogTargetOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) CrawlerCatalogTarget {
		return vs[0].([]CrawlerCatalogTarget)[vs[1].(int)]
	}).(CrawlerCatalogTargetOutput)
}

func init() {
	pulumi.RegisterOutputType(CrawlerCatalogTargetOutput{})
	pulumi.RegisterOutputType(CrawlerCatalogTargetArrayOutput{})
}
