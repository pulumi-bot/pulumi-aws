// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type EndpointMongodbSettings struct {
	AuthMechanism     *string `pulumi:"authMechanism"`
	AuthSource        *string `pulumi:"authSource"`
	AuthType          *string `pulumi:"authType"`
	DocsToInvestigate *string `pulumi:"docsToInvestigate"`
	ExtractDocId      *string `pulumi:"extractDocId"`
	NestingLevel      *string `pulumi:"nestingLevel"`
}

type EndpointMongodbSettingsInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput
	ToEndpointMongodbSettingsOutputWithContext(context.Context) EndpointMongodbSettingsOutput
}

type EndpointMongodbSettingsArgs struct {
	AuthMechanism     pulumi.StringPtrInput `pulumi:"authMechanism"`
	AuthSource        pulumi.StringPtrInput `pulumi:"authSource"`
	AuthType          pulumi.StringPtrInput `pulumi:"authType"`
	DocsToInvestigate pulumi.StringPtrInput `pulumi:"docsToInvestigate"`
	ExtractDocId      pulumi.StringPtrInput `pulumi:"extractDocId"`
	NestingLevel      pulumi.StringPtrInput `pulumi:"nestingLevel"`
}

func (EndpointMongodbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettings)(nil)).Elem()
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput {
	return i.ToEndpointMongodbSettingsOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsOutputWithContext(ctx context.Context) EndpointMongodbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsOutput)
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return i.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgs) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsOutput).ToEndpointMongodbSettingsPtrOutputWithContext(ctx)
}

type EndpointMongodbSettingsPtrInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput
	ToEndpointMongodbSettingsPtrOutputWithContext(context.Context) EndpointMongodbSettingsPtrOutput
}

type endpointMongodbSettingsPtrType EndpointMongodbSettingsArgs

func EndpointMongodbSettingsPtr(v *EndpointMongodbSettingsArgs) EndpointMongodbSettingsPtrInput {
	return (*endpointMongodbSettingsPtrType)(v)
}

func (*endpointMongodbSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettings)(nil)).Elem()
}

func (i *endpointMongodbSettingsPtrType) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return i.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (i *endpointMongodbSettingsPtrType) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsPtrOutput)
}

type EndpointMongodbSettingsOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettings)(nil)).Elem()
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsOutput() EndpointMongodbSettingsOutput {
	return o
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsOutputWithContext(ctx context.Context) EndpointMongodbSettingsOutput {
	return o
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return o.ToEndpointMongodbSettingsPtrOutputWithContext(context.Background())
}

func (o EndpointMongodbSettingsOutput) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *EndpointMongodbSettings {
		return &v
	}).(EndpointMongodbSettingsPtrOutput)
}
func (o EndpointMongodbSettingsOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettings)(nil)).Elem()
}

func (o EndpointMongodbSettingsPtrOutput) ToEndpointMongodbSettingsPtrOutput() EndpointMongodbSettingsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsPtrOutput) ToEndpointMongodbSettingsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsPtrOutput) Elem() EndpointMongodbSettingsOutput {
	return o.ApplyT(func(v *EndpointMongodbSettings) EndpointMongodbSettings { return *v }).(EndpointMongodbSettingsOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsPtrOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettings) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettingsArgs struct {
	AuthMechanism     *string `pulumi:"authMechanism"`
	AuthSource        *string `pulumi:"authSource"`
	AuthType          *string `pulumi:"authType"`
	DocsToInvestigate *string `pulumi:"docsToInvestigate"`
	ExtractDocId      *string `pulumi:"extractDocId"`
	NestingLevel      *string `pulumi:"nestingLevel"`
}

type EndpointMongodbSettingsArgsInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsArgsOutput() EndpointMongodbSettingsArgsOutput
	ToEndpointMongodbSettingsArgsOutputWithContext(context.Context) EndpointMongodbSettingsArgsOutput
}

type EndpointMongodbSettingsArgsArgs struct {
	AuthMechanism     pulumi.StringPtrInput `pulumi:"authMechanism"`
	AuthSource        pulumi.StringPtrInput `pulumi:"authSource"`
	AuthType          pulumi.StringPtrInput `pulumi:"authType"`
	DocsToInvestigate pulumi.StringPtrInput `pulumi:"docsToInvestigate"`
	ExtractDocId      pulumi.StringPtrInput `pulumi:"extractDocId"`
	NestingLevel      pulumi.StringPtrInput `pulumi:"nestingLevel"`
}

func (EndpointMongodbSettingsArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettingsArgs)(nil)).Elem()
}

func (i EndpointMongodbSettingsArgsArgs) ToEndpointMongodbSettingsArgsOutput() EndpointMongodbSettingsArgsOutput {
	return i.ToEndpointMongodbSettingsArgsOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgsArgs) ToEndpointMongodbSettingsArgsOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsArgsOutput)
}

func (i EndpointMongodbSettingsArgsArgs) ToEndpointMongodbSettingsArgsPtrOutput() EndpointMongodbSettingsArgsPtrOutput {
	return i.ToEndpointMongodbSettingsArgsPtrOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsArgsArgs) ToEndpointMongodbSettingsArgsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsArgsOutput).ToEndpointMongodbSettingsArgsPtrOutputWithContext(ctx)
}

type EndpointMongodbSettingsArgsPtrInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsArgsPtrOutput() EndpointMongodbSettingsArgsPtrOutput
	ToEndpointMongodbSettingsArgsPtrOutputWithContext(context.Context) EndpointMongodbSettingsArgsPtrOutput
}

type endpointMongodbSettingsArgsPtrType EndpointMongodbSettingsArgsArgs

func EndpointMongodbSettingsArgsPtr(v *EndpointMongodbSettingsArgsArgs) EndpointMongodbSettingsArgsPtrInput {
	return (*endpointMongodbSettingsArgsPtrType)(v)
}

func (*endpointMongodbSettingsArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettingsArgs)(nil)).Elem()
}

func (i *endpointMongodbSettingsArgsPtrType) ToEndpointMongodbSettingsArgsPtrOutput() EndpointMongodbSettingsArgsPtrOutput {
	return i.ToEndpointMongodbSettingsArgsPtrOutputWithContext(context.Background())
}

func (i *endpointMongodbSettingsArgsPtrType) ToEndpointMongodbSettingsArgsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsArgsPtrOutput)
}

type EndpointMongodbSettingsArgsOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettingsArgs)(nil)).Elem()
}

func (o EndpointMongodbSettingsArgsOutput) ToEndpointMongodbSettingsArgsOutput() EndpointMongodbSettingsArgsOutput {
	return o
}

func (o EndpointMongodbSettingsArgsOutput) ToEndpointMongodbSettingsArgsOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsOutput {
	return o
}

func (o EndpointMongodbSettingsArgsOutput) ToEndpointMongodbSettingsArgsPtrOutput() EndpointMongodbSettingsArgsPtrOutput {
	return o.ToEndpointMongodbSettingsArgsPtrOutputWithContext(context.Background())
}

func (o EndpointMongodbSettingsArgsOutput) ToEndpointMongodbSettingsArgsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *EndpointMongodbSettingsArgs {
		return &v
	}).(EndpointMongodbSettingsArgsPtrOutput)
}
func (o EndpointMongodbSettingsArgsOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettingsArgsPtrOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMongodbSettingsArgs)(nil)).Elem()
}

func (o EndpointMongodbSettingsArgsPtrOutput) ToEndpointMongodbSettingsArgsPtrOutput() EndpointMongodbSettingsArgsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsArgsPtrOutput) ToEndpointMongodbSettingsArgsPtrOutputWithContext(ctx context.Context) EndpointMongodbSettingsArgsPtrOutput {
	return o
}

func (o EndpointMongodbSettingsArgsPtrOutput) Elem() EndpointMongodbSettingsArgsOutput {
	return o.ApplyT(func(v *EndpointMongodbSettingsArgs) EndpointMongodbSettingsArgs { return *v }).(EndpointMongodbSettingsArgsOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsArgsPtrOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsArgs) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointMongodbSettingsState struct {
	AuthMechanism     *string `pulumi:"authMechanism"`
	AuthSource        *string `pulumi:"authSource"`
	AuthType          *string `pulumi:"authType"`
	DocsToInvestigate *string `pulumi:"docsToInvestigate"`
	ExtractDocId      *string `pulumi:"extractDocId"`
	NestingLevel      *string `pulumi:"nestingLevel"`
}

type EndpointMongodbSettingsStateInput interface {
	pulumi.Input

	ToEndpointMongodbSettingsStateOutput() EndpointMongodbSettingsStateOutput
	ToEndpointMongodbSettingsStateOutputWithContext(context.Context) EndpointMongodbSettingsStateOutput
}

type EndpointMongodbSettingsStateArgs struct {
	AuthMechanism     pulumi.StringPtrInput `pulumi:"authMechanism"`
	AuthSource        pulumi.StringPtrInput `pulumi:"authSource"`
	AuthType          pulumi.StringPtrInput `pulumi:"authType"`
	DocsToInvestigate pulumi.StringPtrInput `pulumi:"docsToInvestigate"`
	ExtractDocId      pulumi.StringPtrInput `pulumi:"extractDocId"`
	NestingLevel      pulumi.StringPtrInput `pulumi:"nestingLevel"`
}

func (EndpointMongodbSettingsStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettingsState)(nil)).Elem()
}

func (i EndpointMongodbSettingsStateArgs) ToEndpointMongodbSettingsStateOutput() EndpointMongodbSettingsStateOutput {
	return i.ToEndpointMongodbSettingsStateOutputWithContext(context.Background())
}

func (i EndpointMongodbSettingsStateArgs) ToEndpointMongodbSettingsStateOutputWithContext(ctx context.Context) EndpointMongodbSettingsStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMongodbSettingsStateOutput)
}

type EndpointMongodbSettingsStateOutput struct{ *pulumi.OutputState }

func (EndpointMongodbSettingsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMongodbSettingsState)(nil)).Elem()
}

func (o EndpointMongodbSettingsStateOutput) ToEndpointMongodbSettingsStateOutput() EndpointMongodbSettingsStateOutput {
	return o
}

func (o EndpointMongodbSettingsStateOutput) ToEndpointMongodbSettingsStateOutputWithContext(ctx context.Context) EndpointMongodbSettingsStateOutput {
	return o
}

func (o EndpointMongodbSettingsStateOutput) AuthMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.AuthMechanism }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsStateOutput) AuthSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.AuthSource }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsStateOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsStateOutput) DocsToInvestigate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.DocsToInvestigate }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsStateOutput) ExtractDocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.ExtractDocId }).(pulumi.StringPtrOutput)
}

func (o EndpointMongodbSettingsStateOutput) NestingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointMongodbSettingsState) *string { return v.NestingLevel }).(pulumi.StringPtrOutput)
}

type EndpointS3Settings struct {
	BucketFolder            *string `pulumi:"bucketFolder"`
	BucketName              *string `pulumi:"bucketName"`
	CompressionType         *string `pulumi:"compressionType"`
	CsvDelimiter            *string `pulumi:"csvDelimiter"`
	CsvRowDelimiter         *string `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition *string `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    *string `pulumi:"serviceAccessRoleArn"`
}

type EndpointS3SettingsInput interface {
	pulumi.Input

	ToEndpointS3SettingsOutput() EndpointS3SettingsOutput
	ToEndpointS3SettingsOutputWithContext(context.Context) EndpointS3SettingsOutput
}

type EndpointS3SettingsArgs struct {
	BucketFolder            pulumi.StringPtrInput `pulumi:"bucketFolder"`
	BucketName              pulumi.StringPtrInput `pulumi:"bucketName"`
	CompressionType         pulumi.StringPtrInput `pulumi:"compressionType"`
	CsvDelimiter            pulumi.StringPtrInput `pulumi:"csvDelimiter"`
	CsvRowDelimiter         pulumi.StringPtrInput `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition pulumi.StringPtrInput `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    pulumi.StringPtrInput `pulumi:"serviceAccessRoleArn"`
}

func (EndpointS3SettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3Settings)(nil)).Elem()
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsOutput() EndpointS3SettingsOutput {
	return i.ToEndpointS3SettingsOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsOutputWithContext(ctx context.Context) EndpointS3SettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsOutput)
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return i.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgs) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsOutput).ToEndpointS3SettingsPtrOutputWithContext(ctx)
}

type EndpointS3SettingsPtrInput interface {
	pulumi.Input

	ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput
	ToEndpointS3SettingsPtrOutputWithContext(context.Context) EndpointS3SettingsPtrOutput
}

type endpointS3SettingsPtrType EndpointS3SettingsArgs

func EndpointS3SettingsPtr(v *EndpointS3SettingsArgs) EndpointS3SettingsPtrInput {
	return (*endpointS3SettingsPtrType)(v)
}

func (*endpointS3SettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3Settings)(nil)).Elem()
}

func (i *endpointS3SettingsPtrType) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return i.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (i *endpointS3SettingsPtrType) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsPtrOutput)
}

type EndpointS3SettingsOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3Settings)(nil)).Elem()
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsOutput() EndpointS3SettingsOutput {
	return o
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsOutputWithContext(ctx context.Context) EndpointS3SettingsOutput {
	return o
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return o.ToEndpointS3SettingsPtrOutputWithContext(context.Background())
}

func (o EndpointS3SettingsOutput) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *EndpointS3Settings {
		return &v
	}).(EndpointS3SettingsPtrOutput)
}
func (o EndpointS3SettingsOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

type EndpointS3SettingsPtrOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3Settings)(nil)).Elem()
}

func (o EndpointS3SettingsPtrOutput) ToEndpointS3SettingsPtrOutput() EndpointS3SettingsPtrOutput {
	return o
}

func (o EndpointS3SettingsPtrOutput) ToEndpointS3SettingsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsPtrOutput {
	return o
}

func (o EndpointS3SettingsPtrOutput) Elem() EndpointS3SettingsOutput {
	return o.ApplyT(func(v *EndpointS3Settings) EndpointS3Settings { return *v }).(EndpointS3SettingsOutput)
}

func (o EndpointS3SettingsPtrOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsPtrOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3Settings) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

type EndpointS3SettingsArgs struct {
	BucketFolder            *string `pulumi:"bucketFolder"`
	BucketName              *string `pulumi:"bucketName"`
	CompressionType         *string `pulumi:"compressionType"`
	CsvDelimiter            *string `pulumi:"csvDelimiter"`
	CsvRowDelimiter         *string `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition *string `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    *string `pulumi:"serviceAccessRoleArn"`
}

type EndpointS3SettingsArgsInput interface {
	pulumi.Input

	ToEndpointS3SettingsArgsOutput() EndpointS3SettingsArgsOutput
	ToEndpointS3SettingsArgsOutputWithContext(context.Context) EndpointS3SettingsArgsOutput
}

type EndpointS3SettingsArgsArgs struct {
	BucketFolder            pulumi.StringPtrInput `pulumi:"bucketFolder"`
	BucketName              pulumi.StringPtrInput `pulumi:"bucketName"`
	CompressionType         pulumi.StringPtrInput `pulumi:"compressionType"`
	CsvDelimiter            pulumi.StringPtrInput `pulumi:"csvDelimiter"`
	CsvRowDelimiter         pulumi.StringPtrInput `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition pulumi.StringPtrInput `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    pulumi.StringPtrInput `pulumi:"serviceAccessRoleArn"`
}

func (EndpointS3SettingsArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3SettingsArgs)(nil)).Elem()
}

func (i EndpointS3SettingsArgsArgs) ToEndpointS3SettingsArgsOutput() EndpointS3SettingsArgsOutput {
	return i.ToEndpointS3SettingsArgsOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgsArgs) ToEndpointS3SettingsArgsOutputWithContext(ctx context.Context) EndpointS3SettingsArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsArgsOutput)
}

func (i EndpointS3SettingsArgsArgs) ToEndpointS3SettingsArgsPtrOutput() EndpointS3SettingsArgsPtrOutput {
	return i.ToEndpointS3SettingsArgsPtrOutputWithContext(context.Background())
}

func (i EndpointS3SettingsArgsArgs) ToEndpointS3SettingsArgsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsArgsOutput).ToEndpointS3SettingsArgsPtrOutputWithContext(ctx)
}

type EndpointS3SettingsArgsPtrInput interface {
	pulumi.Input

	ToEndpointS3SettingsArgsPtrOutput() EndpointS3SettingsArgsPtrOutput
	ToEndpointS3SettingsArgsPtrOutputWithContext(context.Context) EndpointS3SettingsArgsPtrOutput
}

type endpointS3SettingsArgsPtrType EndpointS3SettingsArgsArgs

func EndpointS3SettingsArgsPtr(v *EndpointS3SettingsArgsArgs) EndpointS3SettingsArgsPtrInput {
	return (*endpointS3SettingsArgsPtrType)(v)
}

func (*endpointS3SettingsArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3SettingsArgs)(nil)).Elem()
}

func (i *endpointS3SettingsArgsPtrType) ToEndpointS3SettingsArgsPtrOutput() EndpointS3SettingsArgsPtrOutput {
	return i.ToEndpointS3SettingsArgsPtrOutputWithContext(context.Background())
}

func (i *endpointS3SettingsArgsPtrType) ToEndpointS3SettingsArgsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsArgsPtrOutput)
}

type EndpointS3SettingsArgsOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3SettingsArgs)(nil)).Elem()
}

func (o EndpointS3SettingsArgsOutput) ToEndpointS3SettingsArgsOutput() EndpointS3SettingsArgsOutput {
	return o
}

func (o EndpointS3SettingsArgsOutput) ToEndpointS3SettingsArgsOutputWithContext(ctx context.Context) EndpointS3SettingsArgsOutput {
	return o
}

func (o EndpointS3SettingsArgsOutput) ToEndpointS3SettingsArgsPtrOutput() EndpointS3SettingsArgsPtrOutput {
	return o.ToEndpointS3SettingsArgsPtrOutputWithContext(context.Background())
}

func (o EndpointS3SettingsArgsOutput) ToEndpointS3SettingsArgsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsArgsPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *EndpointS3SettingsArgs {
		return &v
	}).(EndpointS3SettingsArgsPtrOutput)
}
func (o EndpointS3SettingsArgsOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

type EndpointS3SettingsArgsPtrOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointS3SettingsArgs)(nil)).Elem()
}

func (o EndpointS3SettingsArgsPtrOutput) ToEndpointS3SettingsArgsPtrOutput() EndpointS3SettingsArgsPtrOutput {
	return o
}

func (o EndpointS3SettingsArgsPtrOutput) ToEndpointS3SettingsArgsPtrOutputWithContext(ctx context.Context) EndpointS3SettingsArgsPtrOutput {
	return o
}

func (o EndpointS3SettingsArgsPtrOutput) Elem() EndpointS3SettingsArgsOutput {
	return o.ApplyT(func(v *EndpointS3SettingsArgs) EndpointS3SettingsArgs { return *v }).(EndpointS3SettingsArgsOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsArgsPtrOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsArgs) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

type EndpointS3SettingsState struct {
	BucketFolder            *string `pulumi:"bucketFolder"`
	BucketName              *string `pulumi:"bucketName"`
	CompressionType         *string `pulumi:"compressionType"`
	CsvDelimiter            *string `pulumi:"csvDelimiter"`
	CsvRowDelimiter         *string `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition *string `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    *string `pulumi:"serviceAccessRoleArn"`
}

type EndpointS3SettingsStateInput interface {
	pulumi.Input

	ToEndpointS3SettingsStateOutput() EndpointS3SettingsStateOutput
	ToEndpointS3SettingsStateOutputWithContext(context.Context) EndpointS3SettingsStateOutput
}

type EndpointS3SettingsStateArgs struct {
	BucketFolder            pulumi.StringPtrInput `pulumi:"bucketFolder"`
	BucketName              pulumi.StringPtrInput `pulumi:"bucketName"`
	CompressionType         pulumi.StringPtrInput `pulumi:"compressionType"`
	CsvDelimiter            pulumi.StringPtrInput `pulumi:"csvDelimiter"`
	CsvRowDelimiter         pulumi.StringPtrInput `pulumi:"csvRowDelimiter"`
	ExternalTableDefinition pulumi.StringPtrInput `pulumi:"externalTableDefinition"`
	ServiceAccessRoleArn    pulumi.StringPtrInput `pulumi:"serviceAccessRoleArn"`
}

func (EndpointS3SettingsStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3SettingsState)(nil)).Elem()
}

func (i EndpointS3SettingsStateArgs) ToEndpointS3SettingsStateOutput() EndpointS3SettingsStateOutput {
	return i.ToEndpointS3SettingsStateOutputWithContext(context.Background())
}

func (i EndpointS3SettingsStateArgs) ToEndpointS3SettingsStateOutputWithContext(ctx context.Context) EndpointS3SettingsStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointS3SettingsStateOutput)
}

type EndpointS3SettingsStateOutput struct{ *pulumi.OutputState }

func (EndpointS3SettingsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointS3SettingsState)(nil)).Elem()
}

func (o EndpointS3SettingsStateOutput) ToEndpointS3SettingsStateOutput() EndpointS3SettingsStateOutput {
	return o
}

func (o EndpointS3SettingsStateOutput) ToEndpointS3SettingsStateOutputWithContext(ctx context.Context) EndpointS3SettingsStateOutput {
	return o
}

func (o EndpointS3SettingsStateOutput) BucketFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.BucketFolder }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) CompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.CompressionType }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) CsvDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.CsvDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) CsvRowDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.CsvRowDelimiter }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) ExternalTableDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.ExternalTableDefinition }).(pulumi.StringPtrOutput)
}

func (o EndpointS3SettingsStateOutput) ServiceAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointS3SettingsState) *string { return v.ServiceAccessRoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointMongodbSettingsOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsArgsOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsArgsPtrOutput{})
	pulumi.RegisterOutputType(EndpointMongodbSettingsStateOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsArgsOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsArgsPtrOutput{})
	pulumi.RegisterOutputType(EndpointS3SettingsStateOutput{})
}
