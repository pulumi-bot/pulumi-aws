// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BrokerUser

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BrokerUser struct {
	// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
	ConsoleAccess *bool `pulumi:"consoleAccess"`
	// The list of groups (20 maximum) to which the ActiveMQ user belongs.
	Groups []string `pulumi:"groups"`
	// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
	Password string `pulumi:"password"`
	// The username of the user.
	Username string `pulumi:"username"`
}

type BrokerUserInput interface {
	pulumi.Input

	ToBrokerUserOutput() BrokerUserOutput
	ToBrokerUserOutputWithContext(context.Context) BrokerUserOutput
}

type BrokerUserArgs struct {
	// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
	ConsoleAccess pulumi.BoolPtrInput `pulumi:"consoleAccess"`
	// The list of groups (20 maximum) to which the ActiveMQ user belongs.
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
	Password pulumi.StringInput `pulumi:"password"`
	// The username of the user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (BrokerUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BrokerUser)(nil)).Elem()
}

func (i BrokerUserArgs) ToBrokerUserOutput() BrokerUserOutput {
	return i.ToBrokerUserOutputWithContext(context.Background())
}

func (i BrokerUserArgs) ToBrokerUserOutputWithContext(ctx context.Context) BrokerUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrokerUserOutput)
}

type BrokerUserArrayInput interface {
	pulumi.Input

	ToBrokerUserArrayOutput() BrokerUserArrayOutput
	ToBrokerUserArrayOutputWithContext(context.Context) BrokerUserArrayOutput
}

type BrokerUserArray []BrokerUserInput

func (BrokerUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BrokerUser)(nil)).Elem()
}

func (i BrokerUserArray) ToBrokerUserArrayOutput() BrokerUserArrayOutput {
	return i.ToBrokerUserArrayOutputWithContext(context.Background())
}

func (i BrokerUserArray) ToBrokerUserArrayOutputWithContext(ctx context.Context) BrokerUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrokerUserArrayOutput)
}

type BrokerUserOutput struct { *pulumi.OutputState }

func (BrokerUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BrokerUser)(nil)).Elem()
}

func (o BrokerUserOutput) ToBrokerUserOutput() BrokerUserOutput {
	return o
}

func (o BrokerUserOutput) ToBrokerUserOutputWithContext(ctx context.Context) BrokerUserOutput {
	return o
}

// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
func (o BrokerUserOutput) ConsoleAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v BrokerUser) *bool { return v.ConsoleAccess }).(pulumi.BoolPtrOutput)
}

// The list of groups (20 maximum) to which the ActiveMQ user belongs.
func (o BrokerUserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v BrokerUser) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
func (o BrokerUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func (v BrokerUser) string { return v.Password }).(pulumi.StringOutput)
}

// The username of the user.
func (o BrokerUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func (v BrokerUser) string { return v.Username }).(pulumi.StringOutput)
}

type BrokerUserArrayOutput struct { *pulumi.OutputState}

func (BrokerUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BrokerUser)(nil)).Elem()
}

func (o BrokerUserArrayOutput) ToBrokerUserArrayOutput() BrokerUserArrayOutput {
	return o
}

func (o BrokerUserArrayOutput) ToBrokerUserArrayOutputWithContext(ctx context.Context) BrokerUserArrayOutput {
	return o
}

func (o BrokerUserArrayOutput) Index(i pulumi.IntInput) BrokerUserOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) BrokerUser {
		return vs[0].([]BrokerUser)[vs[1].(int)]
	}).(BrokerUserOutput)
}

func init() {
	pulumi.RegisterOutputType(BrokerUserOutput{})
	pulumi.RegisterOutputType(BrokerUserArrayOutput{})
}
