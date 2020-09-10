// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceDnsConfig struct {
	DnsRecords    []ServiceDnsConfigDnsRecord `pulumi:"dnsRecords"`
	NamespaceId   string                      `pulumi:"namespaceId"`
	RoutingPolicy *string                     `pulumi:"routingPolicy"`
}

// ServiceDnsConfigInput is an input type that accepts ServiceDnsConfigArgs and ServiceDnsConfigOutput values.
// You can construct a concrete instance of `ServiceDnsConfigInput` via:
//
//          ServiceDnsConfigArgs{...}
type ServiceDnsConfigInput interface {
	pulumi.Input

	ToServiceDnsConfigOutput() ServiceDnsConfigOutput
	ToServiceDnsConfigOutputWithContext(context.Context) ServiceDnsConfigOutput
}

type ServiceDnsConfigArgs struct {
	DnsRecords    ServiceDnsConfigDnsRecordArrayInput `pulumi:"dnsRecords"`
	NamespaceId   pulumi.StringInput                  `pulumi:"namespaceId"`
	RoutingPolicy pulumi.StringPtrInput               `pulumi:"routingPolicy"`
}

func (ServiceDnsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDnsConfig)(nil)).Elem()
}

func (i ServiceDnsConfigArgs) ToServiceDnsConfigOutput() ServiceDnsConfigOutput {
	return i.ToServiceDnsConfigOutputWithContext(context.Background())
}

func (i ServiceDnsConfigArgs) ToServiceDnsConfigOutputWithContext(ctx context.Context) ServiceDnsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDnsConfigOutput)
}

func (i ServiceDnsConfigArgs) ToServiceDnsConfigPtrOutput() ServiceDnsConfigPtrOutput {
	return i.ToServiceDnsConfigPtrOutputWithContext(context.Background())
}

func (i ServiceDnsConfigArgs) ToServiceDnsConfigPtrOutputWithContext(ctx context.Context) ServiceDnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDnsConfigOutput).ToServiceDnsConfigPtrOutputWithContext(ctx)
}

// ServiceDnsConfigPtrInput is an input type that accepts ServiceDnsConfigArgs, ServiceDnsConfigPtr and ServiceDnsConfigPtrOutput values.
// You can construct a concrete instance of `ServiceDnsConfigPtrInput` via:
//
//          ServiceDnsConfigArgs{...}
//
//  or:
//
//          nil
type ServiceDnsConfigPtrInput interface {
	pulumi.Input

	ToServiceDnsConfigPtrOutput() ServiceDnsConfigPtrOutput
	ToServiceDnsConfigPtrOutputWithContext(context.Context) ServiceDnsConfigPtrOutput
}

type serviceDnsConfigPtrType ServiceDnsConfigArgs

func ServiceDnsConfigPtr(v *ServiceDnsConfigArgs) ServiceDnsConfigPtrInput {
	return (*serviceDnsConfigPtrType)(v)
}

func (*serviceDnsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDnsConfig)(nil)).Elem()
}

func (i *serviceDnsConfigPtrType) ToServiceDnsConfigPtrOutput() ServiceDnsConfigPtrOutput {
	return i.ToServiceDnsConfigPtrOutputWithContext(context.Background())
}

func (i *serviceDnsConfigPtrType) ToServiceDnsConfigPtrOutputWithContext(ctx context.Context) ServiceDnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDnsConfigPtrOutput)
}

type ServiceDnsConfigOutput struct{ *pulumi.OutputState }

func (ServiceDnsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDnsConfig)(nil)).Elem()
}

func (o ServiceDnsConfigOutput) ToServiceDnsConfigOutput() ServiceDnsConfigOutput {
	return o
}

func (o ServiceDnsConfigOutput) ToServiceDnsConfigOutputWithContext(ctx context.Context) ServiceDnsConfigOutput {
	return o
}

func (o ServiceDnsConfigOutput) ToServiceDnsConfigPtrOutput() ServiceDnsConfigPtrOutput {
	return o.ToServiceDnsConfigPtrOutputWithContext(context.Background())
}

func (o ServiceDnsConfigOutput) ToServiceDnsConfigPtrOutputWithContext(ctx context.Context) ServiceDnsConfigPtrOutput {
	return o.ApplyT(func(v ServiceDnsConfig) *ServiceDnsConfig {
		return &v
	}).(ServiceDnsConfigPtrOutput)
}
func (o ServiceDnsConfigOutput) DnsRecords() ServiceDnsConfigDnsRecordArrayOutput {
	return o.ApplyT(func(v ServiceDnsConfig) []ServiceDnsConfigDnsRecord { return v.DnsRecords }).(ServiceDnsConfigDnsRecordArrayOutput)
}

func (o ServiceDnsConfigOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDnsConfig) string { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o ServiceDnsConfigOutput) RoutingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceDnsConfig) *string { return v.RoutingPolicy }).(pulumi.StringPtrOutput)
}

type ServiceDnsConfigPtrOutput struct{ *pulumi.OutputState }

func (ServiceDnsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDnsConfig)(nil)).Elem()
}

func (o ServiceDnsConfigPtrOutput) ToServiceDnsConfigPtrOutput() ServiceDnsConfigPtrOutput {
	return o
}

func (o ServiceDnsConfigPtrOutput) ToServiceDnsConfigPtrOutputWithContext(ctx context.Context) ServiceDnsConfigPtrOutput {
	return o
}

func (o ServiceDnsConfigPtrOutput) Elem() ServiceDnsConfigOutput {
	return o.ApplyT(func(v *ServiceDnsConfig) ServiceDnsConfig { return *v }).(ServiceDnsConfigOutput)
}

func (o ServiceDnsConfigPtrOutput) DnsRecords() ServiceDnsConfigDnsRecordArrayOutput {
	return o.ApplyT(func(v *ServiceDnsConfig) []ServiceDnsConfigDnsRecord {
		if v == nil {
			return nil
		}
		return v.DnsRecords
	}).(ServiceDnsConfigDnsRecordArrayOutput)
}

func (o ServiceDnsConfigPtrOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceDnsConfig) *string {
		if v == nil {
			return nil
		}
		return &v.NamespaceId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceDnsConfigPtrOutput) RoutingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceDnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.RoutingPolicy
	}).(pulumi.StringPtrOutput)
}

type ServiceDnsConfigDnsRecord struct {
	Ttl  int    `pulumi:"ttl"`
	Type string `pulumi:"type"`
}

// ServiceDnsConfigDnsRecordInput is an input type that accepts ServiceDnsConfigDnsRecordArgs and ServiceDnsConfigDnsRecordOutput values.
// You can construct a concrete instance of `ServiceDnsConfigDnsRecordInput` via:
//
//          ServiceDnsConfigDnsRecordArgs{...}
type ServiceDnsConfigDnsRecordInput interface {
	pulumi.Input

	ToServiceDnsConfigDnsRecordOutput() ServiceDnsConfigDnsRecordOutput
	ToServiceDnsConfigDnsRecordOutputWithContext(context.Context) ServiceDnsConfigDnsRecordOutput
}

type ServiceDnsConfigDnsRecordArgs struct {
	Ttl  pulumi.IntInput    `pulumi:"ttl"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServiceDnsConfigDnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDnsConfigDnsRecord)(nil)).Elem()
}

func (i ServiceDnsConfigDnsRecordArgs) ToServiceDnsConfigDnsRecordOutput() ServiceDnsConfigDnsRecordOutput {
	return i.ToServiceDnsConfigDnsRecordOutputWithContext(context.Background())
}

func (i ServiceDnsConfigDnsRecordArgs) ToServiceDnsConfigDnsRecordOutputWithContext(ctx context.Context) ServiceDnsConfigDnsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDnsConfigDnsRecordOutput)
}

// ServiceDnsConfigDnsRecordArrayInput is an input type that accepts ServiceDnsConfigDnsRecordArray and ServiceDnsConfigDnsRecordArrayOutput values.
// You can construct a concrete instance of `ServiceDnsConfigDnsRecordArrayInput` via:
//
//          ServiceDnsConfigDnsRecordArray{ ServiceDnsConfigDnsRecordArgs{...} }
type ServiceDnsConfigDnsRecordArrayInput interface {
	pulumi.Input

	ToServiceDnsConfigDnsRecordArrayOutput() ServiceDnsConfigDnsRecordArrayOutput
	ToServiceDnsConfigDnsRecordArrayOutputWithContext(context.Context) ServiceDnsConfigDnsRecordArrayOutput
}

type ServiceDnsConfigDnsRecordArray []ServiceDnsConfigDnsRecordInput

func (ServiceDnsConfigDnsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceDnsConfigDnsRecord)(nil)).Elem()
}

func (i ServiceDnsConfigDnsRecordArray) ToServiceDnsConfigDnsRecordArrayOutput() ServiceDnsConfigDnsRecordArrayOutput {
	return i.ToServiceDnsConfigDnsRecordArrayOutputWithContext(context.Background())
}

func (i ServiceDnsConfigDnsRecordArray) ToServiceDnsConfigDnsRecordArrayOutputWithContext(ctx context.Context) ServiceDnsConfigDnsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDnsConfigDnsRecordArrayOutput)
}

type ServiceDnsConfigDnsRecordOutput struct{ *pulumi.OutputState }

func (ServiceDnsConfigDnsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDnsConfigDnsRecord)(nil)).Elem()
}

func (o ServiceDnsConfigDnsRecordOutput) ToServiceDnsConfigDnsRecordOutput() ServiceDnsConfigDnsRecordOutput {
	return o
}

func (o ServiceDnsConfigDnsRecordOutput) ToServiceDnsConfigDnsRecordOutputWithContext(ctx context.Context) ServiceDnsConfigDnsRecordOutput {
	return o
}

func (o ServiceDnsConfigDnsRecordOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceDnsConfigDnsRecord) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o ServiceDnsConfigDnsRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDnsConfigDnsRecord) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceDnsConfigDnsRecordArrayOutput struct{ *pulumi.OutputState }

func (ServiceDnsConfigDnsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceDnsConfigDnsRecord)(nil)).Elem()
}

func (o ServiceDnsConfigDnsRecordArrayOutput) ToServiceDnsConfigDnsRecordArrayOutput() ServiceDnsConfigDnsRecordArrayOutput {
	return o
}

func (o ServiceDnsConfigDnsRecordArrayOutput) ToServiceDnsConfigDnsRecordArrayOutputWithContext(ctx context.Context) ServiceDnsConfigDnsRecordArrayOutput {
	return o
}

func (o ServiceDnsConfigDnsRecordArrayOutput) Index(i pulumi.IntInput) ServiceDnsConfigDnsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceDnsConfigDnsRecord {
		return vs[0].([]ServiceDnsConfigDnsRecord)[vs[1].(int)]
	}).(ServiceDnsConfigDnsRecordOutput)
}

type ServiceHealthCheckConfig struct {
	FailureThreshold *int    `pulumi:"failureThreshold"`
	ResourcePath     *string `pulumi:"resourcePath"`
	Type             *string `pulumi:"type"`
}

// ServiceHealthCheckConfigInput is an input type that accepts ServiceHealthCheckConfigArgs and ServiceHealthCheckConfigOutput values.
// You can construct a concrete instance of `ServiceHealthCheckConfigInput` via:
//
//          ServiceHealthCheckConfigArgs{...}
type ServiceHealthCheckConfigInput interface {
	pulumi.Input

	ToServiceHealthCheckConfigOutput() ServiceHealthCheckConfigOutput
	ToServiceHealthCheckConfigOutputWithContext(context.Context) ServiceHealthCheckConfigOutput
}

type ServiceHealthCheckConfigArgs struct {
	FailureThreshold pulumi.IntPtrInput    `pulumi:"failureThreshold"`
	ResourcePath     pulumi.StringPtrInput `pulumi:"resourcePath"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
}

func (ServiceHealthCheckConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceHealthCheckConfig)(nil)).Elem()
}

func (i ServiceHealthCheckConfigArgs) ToServiceHealthCheckConfigOutput() ServiceHealthCheckConfigOutput {
	return i.ToServiceHealthCheckConfigOutputWithContext(context.Background())
}

func (i ServiceHealthCheckConfigArgs) ToServiceHealthCheckConfigOutputWithContext(ctx context.Context) ServiceHealthCheckConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckConfigOutput)
}

func (i ServiceHealthCheckConfigArgs) ToServiceHealthCheckConfigPtrOutput() ServiceHealthCheckConfigPtrOutput {
	return i.ToServiceHealthCheckConfigPtrOutputWithContext(context.Background())
}

func (i ServiceHealthCheckConfigArgs) ToServiceHealthCheckConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckConfigOutput).ToServiceHealthCheckConfigPtrOutputWithContext(ctx)
}

// ServiceHealthCheckConfigPtrInput is an input type that accepts ServiceHealthCheckConfigArgs, ServiceHealthCheckConfigPtr and ServiceHealthCheckConfigPtrOutput values.
// You can construct a concrete instance of `ServiceHealthCheckConfigPtrInput` via:
//
//          ServiceHealthCheckConfigArgs{...}
//
//  or:
//
//          nil
type ServiceHealthCheckConfigPtrInput interface {
	pulumi.Input

	ToServiceHealthCheckConfigPtrOutput() ServiceHealthCheckConfigPtrOutput
	ToServiceHealthCheckConfigPtrOutputWithContext(context.Context) ServiceHealthCheckConfigPtrOutput
}

type serviceHealthCheckConfigPtrType ServiceHealthCheckConfigArgs

func ServiceHealthCheckConfigPtr(v *ServiceHealthCheckConfigArgs) ServiceHealthCheckConfigPtrInput {
	return (*serviceHealthCheckConfigPtrType)(v)
}

func (*serviceHealthCheckConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHealthCheckConfig)(nil)).Elem()
}

func (i *serviceHealthCheckConfigPtrType) ToServiceHealthCheckConfigPtrOutput() ServiceHealthCheckConfigPtrOutput {
	return i.ToServiceHealthCheckConfigPtrOutputWithContext(context.Background())
}

func (i *serviceHealthCheckConfigPtrType) ToServiceHealthCheckConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckConfigPtrOutput)
}

type ServiceHealthCheckConfigOutput struct{ *pulumi.OutputState }

func (ServiceHealthCheckConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceHealthCheckConfig)(nil)).Elem()
}

func (o ServiceHealthCheckConfigOutput) ToServiceHealthCheckConfigOutput() ServiceHealthCheckConfigOutput {
	return o
}

func (o ServiceHealthCheckConfigOutput) ToServiceHealthCheckConfigOutputWithContext(ctx context.Context) ServiceHealthCheckConfigOutput {
	return o
}

func (o ServiceHealthCheckConfigOutput) ToServiceHealthCheckConfigPtrOutput() ServiceHealthCheckConfigPtrOutput {
	return o.ToServiceHealthCheckConfigPtrOutputWithContext(context.Background())
}

func (o ServiceHealthCheckConfigOutput) ToServiceHealthCheckConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckConfigPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckConfig) *ServiceHealthCheckConfig {
		return &v
	}).(ServiceHealthCheckConfigPtrOutput)
}
func (o ServiceHealthCheckConfigOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckConfig) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o ServiceHealthCheckConfigOutput) ResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckConfig) *string { return v.ResourcePath }).(pulumi.StringPtrOutput)
}

func (o ServiceHealthCheckConfigOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckConfig) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServiceHealthCheckConfigPtrOutput struct{ *pulumi.OutputState }

func (ServiceHealthCheckConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHealthCheckConfig)(nil)).Elem()
}

func (o ServiceHealthCheckConfigPtrOutput) ToServiceHealthCheckConfigPtrOutput() ServiceHealthCheckConfigPtrOutput {
	return o
}

func (o ServiceHealthCheckConfigPtrOutput) ToServiceHealthCheckConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckConfigPtrOutput {
	return o
}

func (o ServiceHealthCheckConfigPtrOutput) Elem() ServiceHealthCheckConfigOutput {
	return o.ApplyT(func(v *ServiceHealthCheckConfig) ServiceHealthCheckConfig { return *v }).(ServiceHealthCheckConfigOutput)
}

func (o ServiceHealthCheckConfigPtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceHealthCheckConfig) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func (o ServiceHealthCheckConfigPtrOutput) ResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceHealthCheckConfig) *string {
		if v == nil {
			return nil
		}
		return v.ResourcePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceHealthCheckConfigPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceHealthCheckConfig) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceHealthCheckCustomConfig struct {
	FailureThreshold *int `pulumi:"failureThreshold"`
}

// ServiceHealthCheckCustomConfigInput is an input type that accepts ServiceHealthCheckCustomConfigArgs and ServiceHealthCheckCustomConfigOutput values.
// You can construct a concrete instance of `ServiceHealthCheckCustomConfigInput` via:
//
//          ServiceHealthCheckCustomConfigArgs{...}
type ServiceHealthCheckCustomConfigInput interface {
	pulumi.Input

	ToServiceHealthCheckCustomConfigOutput() ServiceHealthCheckCustomConfigOutput
	ToServiceHealthCheckCustomConfigOutputWithContext(context.Context) ServiceHealthCheckCustomConfigOutput
}

type ServiceHealthCheckCustomConfigArgs struct {
	FailureThreshold pulumi.IntPtrInput `pulumi:"failureThreshold"`
}

func (ServiceHealthCheckCustomConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceHealthCheckCustomConfig)(nil)).Elem()
}

func (i ServiceHealthCheckCustomConfigArgs) ToServiceHealthCheckCustomConfigOutput() ServiceHealthCheckCustomConfigOutput {
	return i.ToServiceHealthCheckCustomConfigOutputWithContext(context.Background())
}

func (i ServiceHealthCheckCustomConfigArgs) ToServiceHealthCheckCustomConfigOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckCustomConfigOutput)
}

func (i ServiceHealthCheckCustomConfigArgs) ToServiceHealthCheckCustomConfigPtrOutput() ServiceHealthCheckCustomConfigPtrOutput {
	return i.ToServiceHealthCheckCustomConfigPtrOutputWithContext(context.Background())
}

func (i ServiceHealthCheckCustomConfigArgs) ToServiceHealthCheckCustomConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckCustomConfigOutput).ToServiceHealthCheckCustomConfigPtrOutputWithContext(ctx)
}

// ServiceHealthCheckCustomConfigPtrInput is an input type that accepts ServiceHealthCheckCustomConfigArgs, ServiceHealthCheckCustomConfigPtr and ServiceHealthCheckCustomConfigPtrOutput values.
// You can construct a concrete instance of `ServiceHealthCheckCustomConfigPtrInput` via:
//
//          ServiceHealthCheckCustomConfigArgs{...}
//
//  or:
//
//          nil
type ServiceHealthCheckCustomConfigPtrInput interface {
	pulumi.Input

	ToServiceHealthCheckCustomConfigPtrOutput() ServiceHealthCheckCustomConfigPtrOutput
	ToServiceHealthCheckCustomConfigPtrOutputWithContext(context.Context) ServiceHealthCheckCustomConfigPtrOutput
}

type serviceHealthCheckCustomConfigPtrType ServiceHealthCheckCustomConfigArgs

func ServiceHealthCheckCustomConfigPtr(v *ServiceHealthCheckCustomConfigArgs) ServiceHealthCheckCustomConfigPtrInput {
	return (*serviceHealthCheckCustomConfigPtrType)(v)
}

func (*serviceHealthCheckCustomConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHealthCheckCustomConfig)(nil)).Elem()
}

func (i *serviceHealthCheckCustomConfigPtrType) ToServiceHealthCheckCustomConfigPtrOutput() ServiceHealthCheckCustomConfigPtrOutput {
	return i.ToServiceHealthCheckCustomConfigPtrOutputWithContext(context.Background())
}

func (i *serviceHealthCheckCustomConfigPtrType) ToServiceHealthCheckCustomConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHealthCheckCustomConfigPtrOutput)
}

type ServiceHealthCheckCustomConfigOutput struct{ *pulumi.OutputState }

func (ServiceHealthCheckCustomConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceHealthCheckCustomConfig)(nil)).Elem()
}

func (o ServiceHealthCheckCustomConfigOutput) ToServiceHealthCheckCustomConfigOutput() ServiceHealthCheckCustomConfigOutput {
	return o
}

func (o ServiceHealthCheckCustomConfigOutput) ToServiceHealthCheckCustomConfigOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigOutput {
	return o
}

func (o ServiceHealthCheckCustomConfigOutput) ToServiceHealthCheckCustomConfigPtrOutput() ServiceHealthCheckCustomConfigPtrOutput {
	return o.ToServiceHealthCheckCustomConfigPtrOutputWithContext(context.Background())
}

func (o ServiceHealthCheckCustomConfigOutput) ToServiceHealthCheckCustomConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckCustomConfig) *ServiceHealthCheckCustomConfig {
		return &v
	}).(ServiceHealthCheckCustomConfigPtrOutput)
}
func (o ServiceHealthCheckCustomConfigOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceHealthCheckCustomConfig) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

type ServiceHealthCheckCustomConfigPtrOutput struct{ *pulumi.OutputState }

func (ServiceHealthCheckCustomConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHealthCheckCustomConfig)(nil)).Elem()
}

func (o ServiceHealthCheckCustomConfigPtrOutput) ToServiceHealthCheckCustomConfigPtrOutput() ServiceHealthCheckCustomConfigPtrOutput {
	return o
}

func (o ServiceHealthCheckCustomConfigPtrOutput) ToServiceHealthCheckCustomConfigPtrOutputWithContext(ctx context.Context) ServiceHealthCheckCustomConfigPtrOutput {
	return o
}

func (o ServiceHealthCheckCustomConfigPtrOutput) Elem() ServiceHealthCheckCustomConfigOutput {
	return o.ApplyT(func(v *ServiceHealthCheckCustomConfig) ServiceHealthCheckCustomConfig { return *v }).(ServiceHealthCheckCustomConfigOutput)
}

func (o ServiceHealthCheckCustomConfigPtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceHealthCheckCustomConfig) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceDnsConfigOutput{})
	pulumi.RegisterOutputType(ServiceDnsConfigPtrOutput{})
	pulumi.RegisterOutputType(ServiceDnsConfigDnsRecordOutput{})
	pulumi.RegisterOutputType(ServiceDnsConfigDnsRecordArrayOutput{})
	pulumi.RegisterOutputType(ServiceHealthCheckConfigOutput{})
	pulumi.RegisterOutputType(ServiceHealthCheckConfigPtrOutput{})
	pulumi.RegisterOutputType(ServiceHealthCheckCustomConfigOutput{})
	pulumi.RegisterOutputType(ServiceHealthCheckCustomConfigPtrOutput{})
}
