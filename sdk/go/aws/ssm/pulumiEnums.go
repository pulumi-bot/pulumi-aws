// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ParameterType pulumi.String

const (
	String       = ParameterType("String")
	StringList   = ParameterType("StringList")
	SecureString = ParameterType("SecureString")
)

func (ParameterType) possibleValues() map[ParameterType]bool {
	return map[ParameterType]bool{
		String:       true,
		StringList:   true,
		SecureString: true,
	}
}

func (e ParameterType) Validate() error {
	if !e.possibleValues()[e] {
		return fmt.Errorf("unexpected value: %v", e)
	}
	return nil
}

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
