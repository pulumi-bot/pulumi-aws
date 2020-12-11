// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Load Balancer Listener resource.
//
// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
//
// ## Example Usage
// ### Forward Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
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
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(443),
// 			Protocol:        pulumi.String("HTTPS"),
// 			SslPolicy:       pulumi.String("ELBSecurityPolicy-2016-08"),
// 			CertificateArn:  pulumi.String("arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// ### Redirect Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("redirect"),
// 					Redirect: &lb.ListenerDefaultActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
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
// ### Fixed-response Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("fixed-response"),
// 					FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("Fixed response content"),
// 						StatusCode:  pulumi.String("200"),
// 					},
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
// ### Authenticate-cognito Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
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
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("authenticate-cognito"),
// 					AuthenticateCognito: &lb.ListenerDefaultActionAuthenticateCognitoArgs{
// 						UserPoolArn:      pool.Arn,
// 						UserPoolClientId: client.ID(),
// 						UserPoolDomain:   domain.Domain,
// 					},
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// ### Authenticate-oidc Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
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
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("authenticate-oidc"),
// 					AuthenticateOidc: &lb.ListenerDefaultActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("client_id"),
// 						ClientSecret:          pulumi.String("client_secret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// ### Gateway Load Balancer Listener
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			LoadBalancerType: pulumi.String("gateway"),
// 			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
// 				&lb.LoadBalancerSubnetMappingArgs{
// 					SubnetId: pulumi.Any(aws_subnet.Example.Id),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTargetGroup, err := lb.NewTargetGroup(ctx, "exampleTargetGroup", &lb.TargetGroupArgs{
// 			Port:     pulumi.Int(6081),
// 			Protocol: pulumi.String("GENEVE"),
// 			VpcId:    pulumi.Any(aws_vpc.Example.Id),
// 			HealthCheck: &lb.TargetGroupHealthCheckArgs{
// 				Port:     pulumi.String("80"),
// 				Protocol: pulumi.String("HTTP"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "exampleListener", &lb.ListenerArgs{
// 			LoadBalancerArn: exampleLoadBalancer.ID(),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					TargetGroupArn: exampleTargetGroup.ID(),
// 					Type:           pulumi.String("forward"),
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
//
// ## Import
//
// Listeners can be imported using their ARN, e.g.
//
// ```sh
//  $ pulumi import aws:applicationloadbalancing/listener:Listener front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
// ```
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
	// The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringOutput `pulumi:"sslPolicy"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultActions == nil {
		return nil, errors.New("invalid value for required argument 'DefaultActions'")
	}
	if args.LoadBalancerArn == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerArn'")
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
	// The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port *int `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
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
	// The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port pulumi.IntPtrInput
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
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
	// The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port *int `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
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
	// The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port pulumi.IntPtrInput
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	Protocol pulumi.StringPtrInput
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

type ListenerPtrInput interface {
	pulumi.Input

	ToListenerPtrOutput() ListenerPtrOutput
	ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput
}

func (Listener) ElementType() reflect.Type {
	return reflect.TypeOf((*Listener)(nil)).Elem()
}

func (i Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

func (i Listener) ToListenerPtrOutput() ListenerPtrOutput {
	return i.ToListenerPtrOutputWithContext(context.Background())
}

func (i Listener) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPtrOutput)
}

type ListenerOutput struct {
	*pulumi.OutputState
}

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerOutput)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

type ListenerPtrOutput struct {
	*pulumi.OutputState
}

func (ListenerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerPtrOutput) ToListenerPtrOutput() ListenerPtrOutput {
	return o
}

func (o ListenerPtrOutput) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerPtrOutput{})
}
