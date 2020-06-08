// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

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
// ### Forward Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			CertificateArn: pulumi.String("arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					TargetGroupArn: frontEndTargetGroup.Arn,
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(443),
// 			Protocol:        pulumi.String("HTTPS"),
// 			SslPolicy:       pulumi.String("ELBSecurityPolicy-2016-08"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Redirect Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Redirect: &lb.ListenerDefaultActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
// 					Type: pulumi.String("redirect"),
// 				},
// 			},
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Fixed-response Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("Fixed response content"),
// 						StatusCode:  pulumi.String("200"),
// 					},
// 					Type: pulumi.String("fixed-response"),
// 				},
// 			},
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Authenticate-cognito Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := cognito.NewUserPool(ctx, "pool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		client, err := cognito.NewUserPoolClient(ctx, "client", nil)
// 		if err != nil {
// 			return err
// 		}
// 		domain, err := cognito.NewUserPoolDomain(ctx, "domain", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					AuthenticateCognito: &lb.ListenerDefaultActionAuthenticateCognitoArgs{
// 						UserPoolArn:      pool.Arn,
// 						UserPoolClientId: client.ID(),
// 						UserPoolDomain:   domain.Domain,
// 					},
// 					Type: pulumi.String("authenticate-cognito"),
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					TargetGroupArn: frontEndTargetGroup.Arn,
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Authenticate-oidc Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					AuthenticateOidc: &lb.ListenerDefaultActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("clientId"),
// 						ClientSecret:          pulumi.String("clientSecret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 					Type: pulumi.String("authenticate-oidc"),
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					TargetGroupArn: frontEndTargetGroup.Arn,
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancingv2/listener:Listener"),
		},
	})
	opts = append(opts, aliases)
	var resource Listener
	err := ctx.RegisterResource("aws:lb/listener:Listener", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:lb/listener:Listener", name, id, state, &resource, opts...)
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
