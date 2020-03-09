// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancingv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener resource.
// 
// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener.html.markdown.
type Listener struct {
	pulumi.CustomResourceState

	// The ARN of the listener (matches `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayOutput `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringOutput `pulumi:"loadBalancerArn"`
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntOutput `pulumi:"port"`
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringOutput `pulumi:"sslPolicy"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil || args.DefaultActions == nil {
		return nil, errors.New("missing required argument 'DefaultActions'")
	}
	if args == nil || args.LoadBalancerArn == nil {
		return nil, errors.New("missing required argument 'LoadBalancerArn'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil {
		args = &ListenerArgs{}
	}
	var resource Listener
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("aws:elasticloadbalancingv2/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The ARN of the listener (matches `id`)
	Arn *string `pulumi:"arn"`
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn *string `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port *int `pulumi:"port"`
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol *string `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
}

type ListenerState struct {
	// The ARN of the listener (matches `id`)
	Arn pulumi.StringPtrInput
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn pulumi.StringPtrInput
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayInput
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringPtrInput
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntPtrInput
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrInput
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn *string `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port int `pulumi:"port"`
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol *string `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn pulumi.StringPtrInput
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayInput
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringInput
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntInput
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrInput
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

