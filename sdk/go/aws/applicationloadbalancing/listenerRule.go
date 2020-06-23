// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

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
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "static", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					TargetGroupArn: pulumi.String(aws_lb_target_group.Static.Arn),
// 					Type:           pulumi.String("forward"),
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
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(100),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedRouting", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Forward: &lb.ListenerRuleActionForwardArgs{
// 						Stickiness: &lb.ListenerRuleActionForwardStickinessArgs{
// 							Duration: pulumi.Int(600),
// 							Enabled:  pulumi.Bool(true),
// 						},
// 						TargetGroup: pulumi.MapArray{
// 							pulumi.Map{
// 								"arn":    pulumi.String(aws_lb_target_group.Main.Arn),
// 								"weight": pulumi.Float64(80),
// 							},
// 							pulumi.Map{
// 								"arn":    pulumi.String(aws_lb_target_group.Canary.Arn),
// 								"weight": pulumi.Float64(20),
// 							},
// 						},
// 					},
// 					Type: pulumi.String("forward"),
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
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedWeightedRouting", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					TargetGroupArn: pulumi.String(aws_lb_target_group.Static.Arn),
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("my-service.*.mydomain.io"),
// 						},
// 					},
// 				},
// 			},
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "redirectHttpToHttps", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Redirect: &lb.ListenerRuleActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
// 					Type: pulumi.String("redirect"),
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
// 			ListenerArn: frontEndListener.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "healthCheck", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					FixedResponse: &lb.ListenerRuleActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("HEALTHY"),
// 						StatusCode:  pulumi.String("200"),
// 					},
// 					Type: pulumi.String("fixed-response"),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					QueryString: pulumi.Array{
// 						pulumi.Map{
// 							"key":   pulumi.String("health"),
// 							"value": pulumi.String("check"),
// 						},
// 						pulumi.Map{
// 							"value": pulumi.String("bar"),
// 						},
// 					},
// 				},
// 			},
// 			ListenerArn: frontEndListener.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewUserPool(ctx, "pool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewUserPoolClient(ctx, "client", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewUserPoolDomain(ctx, "domain", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "admin", &lb.ListenerRuleArgs{
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					AuthenticateOidc: &lb.ListenerRuleActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("client_id"),
// 						ClientSecret:          pulumi.String("client_secret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 					Type: pulumi.String("authenticate-oidc"),
// 				},
// 				&lb.ListenerRuleActionArgs{
// 					TargetGroupArn: pulumi.String(aws_lb_target_group.Static.Arn),
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			ListenerArn: frontEndListener.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule
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
	err := ctx.RegisterResource("aws:applicationloadbalancing/listenerRule:ListenerRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:applicationloadbalancing/listenerRule:ListenerRule", name, id, state, &resource, opts...)
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
