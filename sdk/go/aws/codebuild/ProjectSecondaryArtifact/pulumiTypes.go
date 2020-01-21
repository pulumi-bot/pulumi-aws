// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ProjectSecondaryArtifact

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ProjectSecondaryArtifact struct {
	// The artifact identifier. Must be the same specified inside AWS CodeBuild buildspec.
	ArtifactIdentifier string `pulumi:"artifactIdentifier"`
	// If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
	EncryptionDisabled *bool `pulumi:"encryptionDisabled"`
	// The location of the source code from git or s3.
	Location *string `pulumi:"location"`
	// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
	Name *string `pulumi:"name"`
	// The namespace to use in storing build artifacts. If `type` is set to `S3`, then valid values for this parameter are: `BUILD_ID` or `NONE`.
	NamespaceType *string `pulumi:"namespaceType"`
	// If set to true, a name specified in the build spec file overrides the artifact name.
	OverrideArtifactName *bool `pulumi:"overrideArtifactName"`
	// The type of build output artifact to create. If `type` is set to `S3`, valid values for this parameter are: `NONE` or `ZIP`
	Packaging *string `pulumi:"packaging"`
	// If `type` is set to `S3`, this is the path to the output artifact
	Path *string `pulumi:"path"`
	// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
	Type string `pulumi:"type"`
}

type ProjectSecondaryArtifactInput interface {
	pulumi.Input

	ToProjectSecondaryArtifactOutput() ProjectSecondaryArtifactOutput
	ToProjectSecondaryArtifactOutputWithContext(context.Context) ProjectSecondaryArtifactOutput
}

type ProjectSecondaryArtifactArgs struct {
	// The artifact identifier. Must be the same specified inside AWS CodeBuild buildspec.
	ArtifactIdentifier pulumi.StringInput `pulumi:"artifactIdentifier"`
	// If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
	EncryptionDisabled pulumi.BoolPtrInput `pulumi:"encryptionDisabled"`
	// The location of the source code from git or s3.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The namespace to use in storing build artifacts. If `type` is set to `S3`, then valid values for this parameter are: `BUILD_ID` or `NONE`.
	NamespaceType pulumi.StringPtrInput `pulumi:"namespaceType"`
	// If set to true, a name specified in the build spec file overrides the artifact name.
	OverrideArtifactName pulumi.BoolPtrInput `pulumi:"overrideArtifactName"`
	// The type of build output artifact to create. If `type` is set to `S3`, valid values for this parameter are: `NONE` or `ZIP`
	Packaging pulumi.StringPtrInput `pulumi:"packaging"`
	// If `type` is set to `S3`, this is the path to the output artifact
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ProjectSecondaryArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSecondaryArtifact)(nil)).Elem()
}

func (i ProjectSecondaryArtifactArgs) ToProjectSecondaryArtifactOutput() ProjectSecondaryArtifactOutput {
	return i.ToProjectSecondaryArtifactOutputWithContext(context.Background())
}

func (i ProjectSecondaryArtifactArgs) ToProjectSecondaryArtifactOutputWithContext(ctx context.Context) ProjectSecondaryArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSecondaryArtifactOutput)
}

type ProjectSecondaryArtifactArrayInput interface {
	pulumi.Input

	ToProjectSecondaryArtifactArrayOutput() ProjectSecondaryArtifactArrayOutput
	ToProjectSecondaryArtifactArrayOutputWithContext(context.Context) ProjectSecondaryArtifactArrayOutput
}

type ProjectSecondaryArtifactArray []ProjectSecondaryArtifactInput

func (ProjectSecondaryArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectSecondaryArtifact)(nil)).Elem()
}

func (i ProjectSecondaryArtifactArray) ToProjectSecondaryArtifactArrayOutput() ProjectSecondaryArtifactArrayOutput {
	return i.ToProjectSecondaryArtifactArrayOutputWithContext(context.Background())
}

func (i ProjectSecondaryArtifactArray) ToProjectSecondaryArtifactArrayOutputWithContext(ctx context.Context) ProjectSecondaryArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSecondaryArtifactArrayOutput)
}

type ProjectSecondaryArtifactOutput struct { *pulumi.OutputState }

func (ProjectSecondaryArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSecondaryArtifact)(nil)).Elem()
}

func (o ProjectSecondaryArtifactOutput) ToProjectSecondaryArtifactOutput() ProjectSecondaryArtifactOutput {
	return o
}

func (o ProjectSecondaryArtifactOutput) ToProjectSecondaryArtifactOutputWithContext(ctx context.Context) ProjectSecondaryArtifactOutput {
	return o
}

// The artifact identifier. Must be the same specified inside AWS CodeBuild buildspec.
func (o ProjectSecondaryArtifactOutput) ArtifactIdentifier() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) string { return v.ArtifactIdentifier }).(pulumi.StringOutput)
}

// If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
func (o ProjectSecondaryArtifactOutput) EncryptionDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *bool { return v.EncryptionDisabled }).(pulumi.BoolPtrOutput)
}

// The location of the source code from git or s3.
func (o ProjectSecondaryArtifactOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
func (o ProjectSecondaryArtifactOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The namespace to use in storing build artifacts. If `type` is set to `S3`, then valid values for this parameter are: `BUILD_ID` or `NONE`.
func (o ProjectSecondaryArtifactOutput) NamespaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *string { return v.NamespaceType }).(pulumi.StringPtrOutput)
}

// If set to true, a name specified in the build spec file overrides the artifact name.
func (o ProjectSecondaryArtifactOutput) OverrideArtifactName() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *bool { return v.OverrideArtifactName }).(pulumi.BoolPtrOutput)
}

// The type of build output artifact to create. If `type` is set to `S3`, valid values for this parameter are: `NONE` or `ZIP`
func (o ProjectSecondaryArtifactOutput) Packaging() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *string { return v.Packaging }).(pulumi.StringPtrOutput)
}

// If `type` is set to `S3`, this is the path to the output artifact
func (o ProjectSecondaryArtifactOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
func (o ProjectSecondaryArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectSecondaryArtifact) string { return v.Type }).(pulumi.StringOutput)
}

type ProjectSecondaryArtifactArrayOutput struct { *pulumi.OutputState}

func (ProjectSecondaryArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectSecondaryArtifact)(nil)).Elem()
}

func (o ProjectSecondaryArtifactArrayOutput) ToProjectSecondaryArtifactArrayOutput() ProjectSecondaryArtifactArrayOutput {
	return o
}

func (o ProjectSecondaryArtifactArrayOutput) ToProjectSecondaryArtifactArrayOutputWithContext(ctx context.Context) ProjectSecondaryArtifactArrayOutput {
	return o
}

func (o ProjectSecondaryArtifactArrayOutput) Index(i pulumi.IntInput) ProjectSecondaryArtifactOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ProjectSecondaryArtifact {
		return vs[0].([]ProjectSecondaryArtifact)[vs[1].(int)]
	}).(ProjectSecondaryArtifactOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectSecondaryArtifactOutput{})
	pulumi.RegisterOutputType(ProjectSecondaryArtifactArrayOutput{})
}
