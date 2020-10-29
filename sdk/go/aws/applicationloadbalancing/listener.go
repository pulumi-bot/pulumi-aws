// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Load Balancer Listener resource.
//
// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
//
// ## Example Usage
//
// Deprecated: aws.applicationloadbalancing.Listener has been deprecated in favor of aws.alb.Listener
type Listener struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayOutput `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringOutput `pulumi:"loadBalancerArn"`
	// The port on which the load balancer is listening.
	Port pulumi.IntOutput `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
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
	err := ctx.RegisterResource("aws:applicationloadbalancing/listener:Listener", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:applicationloadbalancing/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The Amazon Resource Name (ARN) of the target group.
	Arn *string `pulumi:"arn"`
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn *string `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// The port on which the load balancer is listening.
	Port *int `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
}

type ListenerState struct {
	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringPtrInput
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrInput
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayInput
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringPtrInput
	// The port on which the load balancer is listening.
	Port pulumi.IntPtrInput
	// The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
	Protocol pulumi.StringPtrInput
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn *string `pulumi:"certificateArn"`
	// An Action block. Action blocks are documented below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// The ARN of the load balancer.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	// The port on which the load balancer is listening.
	Port int `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrInput
	// An Action block. Action blocks are documented below.
	DefaultActions ListenerDefaultActionArrayInput
	// The ARN of the load balancer.
	LoadBalancerArn pulumi.StringInput
	// The port on which the load balancer is listening.
	Port pulumi.IntInput
	// The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
	Protocol pulumi.StringPtrInput
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}
