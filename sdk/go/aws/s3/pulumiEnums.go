// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// See https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
type CannedAcl string

const (
	CannedAclPrivate                = CannedAcl("private")
	CannedAclPublicRead             = CannedAcl("public-read")
	CannedAclPublicReadWrite        = CannedAcl("public-read-write")
	CannedAclAwsExecRead            = CannedAcl("aws-exec-read")
	CannedAclAuthenticatedRead      = CannedAcl("authenticated-read")
	CannedAclBucketOwnerRead        = CannedAcl("bucket-owner-read")
	CannedAclBucketOwnerFullControl = CannedAcl("bucket-owner-full-control")
	CannedAclLogDeliveryWrite       = CannedAcl("log-delivery-write")
)

func (CannedAcl) ElementType() reflect.Type {
	return reflect.TypeOf((*CannedAcl)(nil)).Elem()
}

func (e CannedAcl) ToCannedAclOutput() CannedAclOutput {
	return pulumi.ToOutput(CannedAcl(e)).(CannedAclOutput)
}

func (e CannedAcl) ToCannedAclOutputWithContext(ctx context.Context) CannedAclOutput {
	return pulumi.ToOutputWithContext(ctx, CannedAcl(e)).(CannedAclOutput)
}

func (e CannedAcl) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return CannedAcl(e).ToCannedAclPtrOutputWithContext(context.Background())
}

func (e CannedAcl) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return CannedAcl(e).ToCannedAclOutputWithContext(ctx).ToCannedAclPtrOutputWithContext(ctx)
}

func (e CannedAcl) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CannedAcl) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CannedAcl) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CannedAcl) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CannedAclOutput struct{ *pulumi.OutputState }

func (CannedAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CannedAcl)(nil)).Elem()
}

func (o CannedAclOutput) ToCannedAclOutput() CannedAclOutput {
	return o
}

func (o CannedAclOutput) ToCannedAclOutputWithContext(ctx context.Context) CannedAclOutput {
	return o
}

func (o CannedAclOutput) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return o.ToCannedAclPtrOutputWithContext(context.Background())
}

func (o CannedAclOutput) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CannedAcl) *CannedAcl {
		return &v
	}).(CannedAclPtrOutput)
}

type CannedAclPtrOutput struct{ *pulumi.OutputState }

func (CannedAclPtrOutput) ElementType() reflect.Type {
	return cannedAclPtrType
}

func (o CannedAclPtrOutput) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return o
}

func (o CannedAclPtrOutput) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return o
}

func (o CannedAclPtrOutput) Elem() CannedAclOutput {
	return o.ApplyT(func(v *CannedAcl) CannedAcl {
		var ret CannedAcl
		if v != nil {
			ret = *v
		}
		return ret
	}).(CannedAclOutput)
}

// CannedAclInput is an input type that accepts CannedAclArgs and CannedAclOutput values.
// You can construct a concrete instance of `CannedAclInput` via:
//
//          CannedAclArgs{...}
type CannedAclInput interface {
	pulumi.Input

	ToCannedAclOutput() CannedAclOutput
	ToCannedAclOutputWithContext(context.Context) CannedAclOutput
}

var cannedAclPtrType = reflect.TypeOf((**CannedAcl)(nil)).Elem()

type CannedAclPtrInput interface {
	pulumi.Input

	ToCannedAclPtrOutput() CannedAclPtrOutput
	ToCannedAclPtrOutputWithContext(context.Context) CannedAclPtrOutput
}

type cannedAclPtr string

func CannedAclPtr(v string) CannedAclPtrInput {
	return (*cannedAclPtr)(&v)
}

func (*cannedAclPtr) ElementType() reflect.Type {
	return cannedAclPtrType
}

func (in *cannedAclPtr) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return pulumi.ToOutput(in).(CannedAclPtrOutput)
}

func (in *cannedAclPtr) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CannedAclPtrOutput)
}
