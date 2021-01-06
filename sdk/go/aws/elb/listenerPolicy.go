// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a load balancer policy to an ELB Listener.
//
// ## Example Usage
// ### Custom Policy
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
// 		_, err := elb.NewLoadBalancer(ctx, "wu_tang", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(443),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(443),
// 					LbProtocol:       pulumi.String("https"),
// 					SslCertificateId: pulumi.String("arn:aws:iam::000000000000:server-certificate/wu-tang.net"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("wu-tang"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ssl", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-ssl"),
// 			PolicyTypeName:   pulumi.String("SSLNegotiationPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("ECDHE-ECDSA-AES128-GCM-SHA256"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("Protocol-TLSv1.2"),
// 					Value: pulumi.String("true"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewListenerPolicy(ctx, "wu_tang_listener_policies_443", &elb.ListenerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			LoadBalancerPort: pulumi.Int(443),
// 			PolicyNames: pulumi.StringArray{
// 				wu_tang_ssl.PolicyName,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This example shows how to customize the TLS settings of an HTTPS listener.
// ### AWS Predefined Security Policy
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
// 		_, err := elb.NewLoadBalancer(ctx, "wu_tang", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(443),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(443),
// 					LbProtocol:       pulumi.String("https"),
// 					SslCertificateId: pulumi.String("arn:aws:iam::000000000000:server-certificate/wu-tang.net"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("wu-tang"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ssl_tls_1_1", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-ssl"),
// 			PolicyTypeName:   pulumi.String("SSLNegotiationPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("Reference-Security-Policy"),
// 					Value: pulumi.String("ELBSecurityPolicy-TLS-1-1-2017-01"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewListenerPolicy(ctx, "wu_tang_listener_policies_443", &elb.ListenerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			LoadBalancerPort: pulumi.Int(443),
// 			PolicyNames: pulumi.StringArray{
// 				wu_tang_ssl_tls_1_1.PolicyName,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This example shows how to add a [Predefined Security Policy for ELBs](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-security-policy-table.html)
type ListenerPolicy struct {
	pulumi.CustomResourceState

	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntOutput `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayOutput `pulumi:"policyNames"`
}

// NewListenerPolicy registers a new resource with the given unique name, arguments, and options.
func NewListenerPolicy(ctx *pulumi.Context,
	name string, args *ListenerPolicyArgs, opts ...pulumi.ResourceOption) (*ListenerPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.LoadBalancerPort == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerPort'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/listenerPolicy:ListenerPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ListenerPolicy
	err := ctx.RegisterResource("aws:elb/listenerPolicy:ListenerPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerPolicy gets an existing ListenerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerPolicyState, opts ...pulumi.ResourceOption) (*ListenerPolicy, error) {
	var resource ListenerPolicy
	err := ctx.ReadResource("aws:elb/listenerPolicy:ListenerPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerPolicy resources.
type listenerPolicyState struct {
	// The load balancer to attach the policy to.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort *int `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

type ListenerPolicyState struct {
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringPtrInput
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntPtrInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (ListenerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerPolicyState)(nil)).Elem()
}

type listenerPolicyArgs struct {
	// The load balancer to attach the policy to.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort int `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

// The set of arguments for constructing a ListenerPolicy resource.
type ListenerPolicyArgs struct {
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringInput
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (ListenerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerPolicyArgs)(nil)).Elem()
}

type ListenerPolicyInput interface {
	pulumi.Input

	ToListenerPolicyOutput() ListenerPolicyOutput
	ToListenerPolicyOutputWithContext(ctx context.Context) ListenerPolicyOutput
}

func (*ListenerPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerPolicy)(nil))
}

func (i *ListenerPolicy) ToListenerPolicyOutput() ListenerPolicyOutput {
	return i.ToListenerPolicyOutputWithContext(context.Background())
}

func (i *ListenerPolicy) ToListenerPolicyOutputWithContext(ctx context.Context) ListenerPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPolicyOutput)
}

func (i *ListenerPolicy) ToListenerPolicyPtrOutput() ListenerPolicyPtrOutput {
	return i.ToListenerPolicyPtrOutputWithContext(context.Background())
}

func (i *ListenerPolicy) ToListenerPolicyPtrOutputWithContext(ctx context.Context) ListenerPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPolicyPtrOutput)
}

type ListenerPolicyPtrInput interface {
	pulumi.Input

	ToListenerPolicyPtrOutput() ListenerPolicyPtrOutput
	ToListenerPolicyPtrOutputWithContext(ctx context.Context) ListenerPolicyPtrOutput
}

type listenerPolicyPtrType ListenerPolicyArgs

func (*listenerPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerPolicy)(nil))
}

func (i *listenerPolicyPtrType) ToListenerPolicyPtrOutput() ListenerPolicyPtrOutput {
	return i.ToListenerPolicyPtrOutputWithContext(context.Background())
}

func (i *listenerPolicyPtrType) ToListenerPolicyPtrOutputWithContext(ctx context.Context) ListenerPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPolicyOutput).ToListenerPolicyPtrOutput()
}

// ListenerPolicyArrayInput is an input type that accepts ListenerPolicyArray and ListenerPolicyArrayOutput values.
// You can construct a concrete instance of `ListenerPolicyArrayInput` via:
//
//          ListenerPolicyArray{ ListenerPolicyArgs{...} }
type ListenerPolicyArrayInput interface {
	pulumi.Input

	ToListenerPolicyArrayOutput() ListenerPolicyArrayOutput
	ToListenerPolicyArrayOutputWithContext(context.Context) ListenerPolicyArrayOutput
}

type ListenerPolicyArray []ListenerPolicyInput

func (ListenerPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerPolicy)(nil))
}

func (i ListenerPolicyArray) ToListenerPolicyArrayOutput() ListenerPolicyArrayOutput {
	return i.ToListenerPolicyArrayOutputWithContext(context.Background())
}

func (i ListenerPolicyArray) ToListenerPolicyArrayOutputWithContext(ctx context.Context) ListenerPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPolicyArrayOutput)
}

// ListenerPolicyMapInput is an input type that accepts ListenerPolicyMap and ListenerPolicyMapOutput values.
// You can construct a concrete instance of `ListenerPolicyMapInput` via:
//
//          ListenerPolicyMap{ "key": ListenerPolicyArgs{...} }
type ListenerPolicyMapInput interface {
	pulumi.Input

	ToListenerPolicyMapOutput() ListenerPolicyMapOutput
	ToListenerPolicyMapOutputWithContext(context.Context) ListenerPolicyMapOutput
}

type ListenerPolicyMap map[string]ListenerPolicyInput

func (ListenerPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ListenerPolicy)(nil))
}

func (i ListenerPolicyMap) ToListenerPolicyMapOutput() ListenerPolicyMapOutput {
	return i.ToListenerPolicyMapOutputWithContext(context.Background())
}

func (i ListenerPolicyMap) ToListenerPolicyMapOutputWithContext(ctx context.Context) ListenerPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPolicyMapOutput)
}

type ListenerPolicyOutput struct {
	*pulumi.OutputState
}

func (ListenerPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerPolicy)(nil))
}

func (o ListenerPolicyOutput) ToListenerPolicyOutput() ListenerPolicyOutput {
	return o
}

func (o ListenerPolicyOutput) ToListenerPolicyOutputWithContext(ctx context.Context) ListenerPolicyOutput {
	return o
}

func (o ListenerPolicyOutput) ToListenerPolicyPtrOutput() ListenerPolicyPtrOutput {
	return o.ToListenerPolicyPtrOutputWithContext(context.Background())
}

func (o ListenerPolicyOutput) ToListenerPolicyPtrOutputWithContext(ctx context.Context) ListenerPolicyPtrOutput {
	return o.ApplyT(func(v ListenerPolicy) *ListenerPolicy {
		return &v
	}).(ListenerPolicyPtrOutput)
}

type ListenerPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ListenerPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerPolicy)(nil))
}

func (o ListenerPolicyPtrOutput) ToListenerPolicyPtrOutput() ListenerPolicyPtrOutput {
	return o
}

func (o ListenerPolicyPtrOutput) ToListenerPolicyPtrOutputWithContext(ctx context.Context) ListenerPolicyPtrOutput {
	return o
}

type ListenerPolicyArrayOutput struct{ *pulumi.OutputState }

func (ListenerPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerPolicy)(nil))
}

func (o ListenerPolicyArrayOutput) ToListenerPolicyArrayOutput() ListenerPolicyArrayOutput {
	return o
}

func (o ListenerPolicyArrayOutput) ToListenerPolicyArrayOutputWithContext(ctx context.Context) ListenerPolicyArrayOutput {
	return o
}

func (o ListenerPolicyArrayOutput) Index(i pulumi.IntInput) ListenerPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ListenerPolicy {
		return vs[0].([]ListenerPolicy)[vs[1].(int)]
	}).(ListenerPolicyOutput)
}

type ListenerPolicyMapOutput struct{ *pulumi.OutputState }

func (ListenerPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ListenerPolicy)(nil))
}

func (o ListenerPolicyMapOutput) ToListenerPolicyMapOutput() ListenerPolicyMapOutput {
	return o
}

func (o ListenerPolicyMapOutput) ToListenerPolicyMapOutputWithContext(ctx context.Context) ListenerPolicyMapOutput {
	return o
}

func (o ListenerPolicyMapOutput) MapIndex(k pulumi.StringInput) ListenerPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ListenerPolicy {
		return vs[0].(map[string]ListenerPolicy)[vs[1].(string)]
	}).(ListenerPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerPolicyOutput{})
	pulumi.RegisterOutputType(ListenerPolicyPtrOutput{})
	pulumi.RegisterOutputType(ListenerPolicyArrayOutput{})
	pulumi.RegisterOutputType(ListenerPolicyMapOutput{})
}
