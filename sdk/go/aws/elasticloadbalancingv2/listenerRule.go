// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Load Balancer Listener Rule resource.
//
// > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.
//
// ## Example Usage
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
// 		_, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "static", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(100),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					PathPattern: &lb.ListenerRuleConditionPathPatternArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("/static/*"),
// 						},
// 					},
// 				},
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("example.com"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedWeightedRouting", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("my-service.*.mycompany.io"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedRouting", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("forward"),
// 					Forward: &lb.ListenerRuleActionForwardArgs{
// 						TargetGroups: lb.ListenerRuleActionForwardTargetGroupArray{
// 							&lb.ListenerRuleActionForwardTargetGroupArgs{
// 								Arn:    pulumi.Any(aws_lb_target_group.Main.Arn),
// 								Weight: pulumi.Int(80),
// 							},
// 							&lb.ListenerRuleActionForwardTargetGroupArgs{
// 								Arn:    pulumi.Any(aws_lb_target_group.Canary.Arn),
// 								Weight: pulumi.Int(20),
// 							},
// 						},
// 						Stickiness: &lb.ListenerRuleActionForwardStickinessArgs{
// 							Enabled:  pulumi.Bool(true),
// 							Duration: pulumi.Int(600),
// 						},
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("my-service.*.mycompany.io"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "redirectHttpToHttps", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("redirect"),
// 					Redirect: &lb.ListenerRuleActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HttpHeader: &lb.ListenerRuleConditionHttpHeaderArgs{
// 						HttpHeaderName: pulumi.String("X-Forwarded-For"),
// 						Values: pulumi.StringArray{
// 							pulumi.String("192.168.1.*"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "healthCheck", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("fixed-response"),
// 					FixedResponse: &lb.ListenerRuleActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("HEALTHY"),
// 						StatusCode:  pulumi.String("200"),
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					QueryStrings: lb.ListenerRuleConditionQueryStringArray{
// 						&lb.ListenerRuleConditionQueryStringArgs{
// 							Key:   pulumi.String("health"),
// 							Value: pulumi.String("check"),
// 						},
// 						&lb.ListenerRuleConditionQueryStringArgs{
// 							Value: pulumi.String("bar"),
// 						},
// 					},
// 				},
// 			},
// 		})
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
// 		_, err = lb.NewListenerRule(ctx, "admin", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("authenticate-cognito"),
// 					AuthenticateCognito: &lb.ListenerRuleActionAuthenticateCognitoArgs{
// 						UserPoolArn:      pool.Arn,
// 						UserPoolClientId: client.ID(),
// 						UserPoolDomain:   domain.Domain,
// 					},
// 				},
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "oidc", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("authenticate-oidc"),
// 					AuthenticateOidc: &lb.ListenerRuleActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("client_id"),
// 						ClientSecret:          pulumi.String("client_secret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 				},
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
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
// Rules can be imported using their ARN, e.g.
//
// ```sh
//  $ pulumi import aws:elasticloadbalancingv2/listenerRule:ListenerRule front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener-rule/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b
// ```
//
// Deprecated: aws.elasticloadbalancingv2.ListenerRule has been deprecated in favor of aws.lb.ListenerRule
type ListenerRule struct {
	pulumi.CustomResourceState

	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayOutput `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayOutput `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
}

// NewListenerRule registers a new resource with the given unique name, arguments, and options.
func NewListenerRule(ctx *pulumi.Context,
	name string, args *ListenerRuleArgs, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	if args == nil || args.Actions == nil {
		return nil, errors.New("missing required argument 'Actions'")
	}
	if args == nil || args.Conditions == nil {
		return nil, errors.New("missing required argument 'Conditions'")
	}
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	if args == nil {
		args = &ListenerRuleArgs{}
	}
	var resource ListenerRule
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerRule gets an existing ListenerRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerRuleState, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	var resource ListenerRule
	err := ctx.ReadResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerRule resources.
type listenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the target group.
	Arn *string `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn *string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
}

type ListenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringPtrInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringPtrInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
}

func (ListenerRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleState)(nil)).Elem()
}

type listenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
}

// The set of arguments for constructing a ListenerRule resource.
type ListenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
}

func (ListenerRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleArgs)(nil)).Elem()
}
