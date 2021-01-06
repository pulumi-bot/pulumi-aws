// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		lb, err := elb.NewLoadBalancer(ctx, "lb", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(8000),
// 					InstanceProtocol: pulumi.String("https"),
// 					LbPort:           pulumi.Int(443),
// 					LbProtocol:       pulumi.String("https"),
// 					SslCertificateId: pulumi.String("arn:aws:iam::123456789012:server-certificate/certName"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewSslNegotiationPolicy(ctx, "foo", &elb.SslNegotiationPolicyArgs{
// 			LoadBalancer: lb.ID(),
// 			LbPort:       pulumi.Int(443),
// 			Attributes: elb.SslNegotiationPolicyAttributeArray{
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("Protocol-TLSv1"),
// 					Value: pulumi.String("false"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("Protocol-TLSv1.1"),
// 					Value: pulumi.String("false"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("Protocol-TLSv1.2"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("Server-Defined-Cipher-Order"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("ECDHE-RSA-AES128-GCM-SHA256"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("AES128-GCM-SHA256"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.SslNegotiationPolicyAttributeArgs{
// 					Name:  pulumi.String("EDH-RSA-DES-CBC3-SHA"),
// 					Value: pulumi.String("false"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SslNegotiationPolicy struct {
	pulumi.CustomResourceState

	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayOutput `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`
	// The name of the attribute
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSslNegotiationPolicy registers a new resource with the given unique name, arguments, and options.
func NewSslNegotiationPolicy(ctx *pulumi.Context,
	name string, args *SslNegotiationPolicyArgs, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbPort == nil {
		return nil, errors.New("invalid value for required argument 'LbPort'")
	}
	if args.LoadBalancer == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancer'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource SslNegotiationPolicy
	err := ctx.RegisterResource("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslNegotiationPolicy gets an existing SslNegotiationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslNegotiationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslNegotiationPolicyState, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	var resource SslNegotiationPolicy
	err := ctx.ReadResource("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslNegotiationPolicy resources.
type sslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes []SslNegotiationPolicyAttribute `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort *int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer *string `pulumi:"loadBalancer"`
	// The name of the attribute
	Name *string `pulumi:"name"`
}

type SslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntPtrInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringPtrInput
	// The name of the attribute
	Name pulumi.StringPtrInput
}

func (SslNegotiationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslNegotiationPolicyState)(nil)).Elem()
}

type sslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes []SslNegotiationPolicyAttribute `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer string `pulumi:"loadBalancer"`
	// The name of the attribute
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SslNegotiationPolicy resource.
type SslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput
	// The name of the attribute
	Name pulumi.StringPtrInput
}

func (SslNegotiationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslNegotiationPolicyArgs)(nil)).Elem()
}

type SslNegotiationPolicyInput interface {
	pulumi.Input

	ToSslNegotiationPolicyOutput() SslNegotiationPolicyOutput
	ToSslNegotiationPolicyOutputWithContext(ctx context.Context) SslNegotiationPolicyOutput
}

func (*SslNegotiationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SslNegotiationPolicy)(nil))
}

func (i *SslNegotiationPolicy) ToSslNegotiationPolicyOutput() SslNegotiationPolicyOutput {
	return i.ToSslNegotiationPolicyOutputWithContext(context.Background())
}

func (i *SslNegotiationPolicy) ToSslNegotiationPolicyOutputWithContext(ctx context.Context) SslNegotiationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslNegotiationPolicyOutput)
}

func (i *SslNegotiationPolicy) ToSslNegotiationPolicyPtrOutput() SslNegotiationPolicyPtrOutput {
	return i.ToSslNegotiationPolicyPtrOutputWithContext(context.Background())
}

func (i *SslNegotiationPolicy) ToSslNegotiationPolicyPtrOutputWithContext(ctx context.Context) SslNegotiationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslNegotiationPolicyPtrOutput)
}

type SslNegotiationPolicyPtrInput interface {
	pulumi.Input

	ToSslNegotiationPolicyPtrOutput() SslNegotiationPolicyPtrOutput
	ToSslNegotiationPolicyPtrOutputWithContext(ctx context.Context) SslNegotiationPolicyPtrOutput
}

type sslNegotiationPolicyPtrType SslNegotiationPolicyArgs

func (*sslNegotiationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslNegotiationPolicy)(nil))
}

func (i *sslNegotiationPolicyPtrType) ToSslNegotiationPolicyPtrOutput() SslNegotiationPolicyPtrOutput {
	return i.ToSslNegotiationPolicyPtrOutputWithContext(context.Background())
}

func (i *sslNegotiationPolicyPtrType) ToSslNegotiationPolicyPtrOutputWithContext(ctx context.Context) SslNegotiationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslNegotiationPolicyOutput).ToSslNegotiationPolicyPtrOutput()
}

// SslNegotiationPolicyArrayInput is an input type that accepts SslNegotiationPolicyArray and SslNegotiationPolicyArrayOutput values.
// You can construct a concrete instance of `SslNegotiationPolicyArrayInput` via:
//
//          SslNegotiationPolicyArray{ SslNegotiationPolicyArgs{...} }
type SslNegotiationPolicyArrayInput interface {
	pulumi.Input

	ToSslNegotiationPolicyArrayOutput() SslNegotiationPolicyArrayOutput
	ToSslNegotiationPolicyArrayOutputWithContext(context.Context) SslNegotiationPolicyArrayOutput
}

type SslNegotiationPolicyArray []SslNegotiationPolicyInput

func (SslNegotiationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SslNegotiationPolicy)(nil))
}

func (i SslNegotiationPolicyArray) ToSslNegotiationPolicyArrayOutput() SslNegotiationPolicyArrayOutput {
	return i.ToSslNegotiationPolicyArrayOutputWithContext(context.Background())
}

func (i SslNegotiationPolicyArray) ToSslNegotiationPolicyArrayOutputWithContext(ctx context.Context) SslNegotiationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslNegotiationPolicyArrayOutput)
}

// SslNegotiationPolicyMapInput is an input type that accepts SslNegotiationPolicyMap and SslNegotiationPolicyMapOutput values.
// You can construct a concrete instance of `SslNegotiationPolicyMapInput` via:
//
//          SslNegotiationPolicyMap{ "key": SslNegotiationPolicyArgs{...} }
type SslNegotiationPolicyMapInput interface {
	pulumi.Input

	ToSslNegotiationPolicyMapOutput() SslNegotiationPolicyMapOutput
	ToSslNegotiationPolicyMapOutputWithContext(context.Context) SslNegotiationPolicyMapOutput
}

type SslNegotiationPolicyMap map[string]SslNegotiationPolicyInput

func (SslNegotiationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SslNegotiationPolicy)(nil))
}

func (i SslNegotiationPolicyMap) ToSslNegotiationPolicyMapOutput() SslNegotiationPolicyMapOutput {
	return i.ToSslNegotiationPolicyMapOutputWithContext(context.Background())
}

func (i SslNegotiationPolicyMap) ToSslNegotiationPolicyMapOutputWithContext(ctx context.Context) SslNegotiationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslNegotiationPolicyMapOutput)
}

type SslNegotiationPolicyOutput struct {
	*pulumi.OutputState
}

func (SslNegotiationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslNegotiationPolicy)(nil))
}

func (o SslNegotiationPolicyOutput) ToSslNegotiationPolicyOutput() SslNegotiationPolicyOutput {
	return o
}

func (o SslNegotiationPolicyOutput) ToSslNegotiationPolicyOutputWithContext(ctx context.Context) SslNegotiationPolicyOutput {
	return o
}

func (o SslNegotiationPolicyOutput) ToSslNegotiationPolicyPtrOutput() SslNegotiationPolicyPtrOutput {
	return o.ToSslNegotiationPolicyPtrOutputWithContext(context.Background())
}

func (o SslNegotiationPolicyOutput) ToSslNegotiationPolicyPtrOutputWithContext(ctx context.Context) SslNegotiationPolicyPtrOutput {
	return o.ApplyT(func(v SslNegotiationPolicy) *SslNegotiationPolicy {
		return &v
	}).(SslNegotiationPolicyPtrOutput)
}

type SslNegotiationPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (SslNegotiationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslNegotiationPolicy)(nil))
}

func (o SslNegotiationPolicyPtrOutput) ToSslNegotiationPolicyPtrOutput() SslNegotiationPolicyPtrOutput {
	return o
}

func (o SslNegotiationPolicyPtrOutput) ToSslNegotiationPolicyPtrOutputWithContext(ctx context.Context) SslNegotiationPolicyPtrOutput {
	return o
}

type SslNegotiationPolicyArrayOutput struct{ *pulumi.OutputState }

func (SslNegotiationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SslNegotiationPolicy)(nil))
}

func (o SslNegotiationPolicyArrayOutput) ToSslNegotiationPolicyArrayOutput() SslNegotiationPolicyArrayOutput {
	return o
}

func (o SslNegotiationPolicyArrayOutput) ToSslNegotiationPolicyArrayOutputWithContext(ctx context.Context) SslNegotiationPolicyArrayOutput {
	return o
}

func (o SslNegotiationPolicyArrayOutput) Index(i pulumi.IntInput) SslNegotiationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SslNegotiationPolicy {
		return vs[0].([]SslNegotiationPolicy)[vs[1].(int)]
	}).(SslNegotiationPolicyOutput)
}

type SslNegotiationPolicyMapOutput struct{ *pulumi.OutputState }

func (SslNegotiationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SslNegotiationPolicy)(nil))
}

func (o SslNegotiationPolicyMapOutput) ToSslNegotiationPolicyMapOutput() SslNegotiationPolicyMapOutput {
	return o
}

func (o SslNegotiationPolicyMapOutput) ToSslNegotiationPolicyMapOutputWithContext(ctx context.Context) SslNegotiationPolicyMapOutput {
	return o
}

func (o SslNegotiationPolicyMapOutput) MapIndex(k pulumi.StringInput) SslNegotiationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SslNegotiationPolicy {
		return vs[0].(map[string]SslNegotiationPolicy)[vs[1].(string)]
	}).(SslNegotiationPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(SslNegotiationPolicyOutput{})
	pulumi.RegisterOutputType(SslNegotiationPolicyPtrOutput{})
	pulumi.RegisterOutputType(SslNegotiationPolicyArrayOutput{})
	pulumi.RegisterOutputType(SslNegotiationPolicyMapOutput{})
}
