// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getApplicationAppversionLifecycle

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetApplicationAppversionLifecycle struct {
	// Specifies whether delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 bool `pulumi:"deleteSourceFromS3"`
	// The number of days to retain an application version.
	MaxAgeInDays int `pulumi:"maxAgeInDays"`
	// The maximum number of application versions to retain.
	MaxCount int `pulumi:"maxCount"`
	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole string `pulumi:"serviceRole"`
}

type GetApplicationAppversionLifecycleInput interface {
	pulumi.Input

	ToGetApplicationAppversionLifecycleOutput() GetApplicationAppversionLifecycleOutput
	ToGetApplicationAppversionLifecycleOutputWithContext(context.Context) GetApplicationAppversionLifecycleOutput
}

type GetApplicationAppversionLifecycleArgs struct {
	// Specifies whether delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 pulumi.BoolInput `pulumi:"deleteSourceFromS3"`
	// The number of days to retain an application version.
	MaxAgeInDays pulumi.IntInput `pulumi:"maxAgeInDays"`
	// The maximum number of application versions to retain.
	MaxCount pulumi.IntInput `pulumi:"maxCount"`
	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole pulumi.StringInput `pulumi:"serviceRole"`
}

func (GetApplicationAppversionLifecycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationAppversionLifecycle)(nil)).Elem()
}

func (i GetApplicationAppversionLifecycleArgs) ToGetApplicationAppversionLifecycleOutput() GetApplicationAppversionLifecycleOutput {
	return i.ToGetApplicationAppversionLifecycleOutputWithContext(context.Background())
}

func (i GetApplicationAppversionLifecycleArgs) ToGetApplicationAppversionLifecycleOutputWithContext(ctx context.Context) GetApplicationAppversionLifecycleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetApplicationAppversionLifecycleOutput)
}

type GetApplicationAppversionLifecycleOutput struct { *pulumi.OutputState }

func (GetApplicationAppversionLifecycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationAppversionLifecycle)(nil)).Elem()
}

func (o GetApplicationAppversionLifecycleOutput) ToGetApplicationAppversionLifecycleOutput() GetApplicationAppversionLifecycleOutput {
	return o
}

func (o GetApplicationAppversionLifecycleOutput) ToGetApplicationAppversionLifecycleOutputWithContext(ctx context.Context) GetApplicationAppversionLifecycleOutput {
	return o
}

// Specifies whether delete a version's source bundle from S3 when the application version is deleted.
func (o GetApplicationAppversionLifecycleOutput) DeleteSourceFromS3() pulumi.BoolOutput {
	return o.ApplyT(func (v GetApplicationAppversionLifecycle) bool { return v.DeleteSourceFromS3 }).(pulumi.BoolOutput)
}

// The number of days to retain an application version.
func (o GetApplicationAppversionLifecycleOutput) MaxAgeInDays() pulumi.IntOutput {
	return o.ApplyT(func (v GetApplicationAppversionLifecycle) int { return v.MaxAgeInDays }).(pulumi.IntOutput)
}

// The maximum number of application versions to retain.
func (o GetApplicationAppversionLifecycleOutput) MaxCount() pulumi.IntOutput {
	return o.ApplyT(func (v GetApplicationAppversionLifecycle) int { return v.MaxCount }).(pulumi.IntOutput)
}

// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
func (o GetApplicationAppversionLifecycleOutput) ServiceRole() pulumi.StringOutput {
	return o.ApplyT(func (v GetApplicationAppversionLifecycle) string { return v.ServiceRole }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationAppversionLifecycleOutput{})
}
