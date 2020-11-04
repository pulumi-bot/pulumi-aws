// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Traffic mirror filter rule.\
// Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
//
// ## Example Usage
//
// To create a basic traffic mirror session
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		filter, err := ec2.NewTrafficMirrorFilter(ctx, "filter", &ec2.TrafficMirrorFilterArgs{
// 			Description: pulumi.String("traffic mirror filter - example"),
// 			NetworkServices: pulumi.StringArray{
// 				pulumi.String("amazon-dns"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "ruleout", &ec2.TrafficMirrorFilterRuleArgs{
// 			Description:           pulumi.String("test rule"),
// 			TrafficMirrorFilterId: filter.ID(),
// 			DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
// 			SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
// 			RuleNumber:            pulumi.Int(1),
// 			RuleAction:            pulumi.String("accept"),
// 			TrafficDirection:      pulumi.String("egress"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "rulein", &ec2.TrafficMirrorFilterRuleArgs{
// 			Description:           pulumi.String("test rule"),
// 			TrafficMirrorFilterId: filter.ID(),
// 			DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
// 			SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
// 			RuleNumber:            pulumi.Int(1),
// 			RuleAction:            pulumi.String("accept"),
// 			TrafficDirection:      pulumi.String("ingress"),
// 			Protocol:              pulumi.Int(6),
// 			DestinationPortRange: &ec2.TrafficMirrorFilterRuleDestinationPortRangeArgs{
// 				FromPort: pulumi.Int(22),
// 				ToPort:   pulumi.Int(53),
// 			},
// 			SourcePortRange: &ec2.TrafficMirrorFilterRuleSourcePortRangeArgs{
// 				FromPort: pulumi.Int(0),
// 				ToPort:   pulumi.Int(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TrafficMirrorFilterRule struct {
	pulumi.CustomResourceState

	// A description of the traffic mirror filter rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrOutput `pulumi:"destinationPortRange"`
	// The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrOutput `pulumi:"protocol"`
	// The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntOutput `pulumi:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringOutput `pulumi:"sourceCidrBlock"`
	// The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrOutput `pulumi:"sourcePortRange"`
	// The direction of traffic to be captured. Valid values are `ingress` and `egress`
	TrafficDirection pulumi.StringOutput `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
}

// NewTrafficMirrorFilterRule registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterRuleArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.RuleNumber == nil {
		return nil, errors.New("invalid value for required argument 'RuleNumber'")
	}
	if args.SourceCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidrBlock'")
	}
	if args.TrafficDirection == nil {
		return nil, errors.New("invalid value for required argument 'TrafficDirection'")
	}
	if args.TrafficMirrorFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorFilterId'")
	}
	var resource TrafficMirrorFilterRule
	err := ctx.RegisterResource("aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilterRule gets an existing TrafficMirrorFilterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterRuleState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	var resource TrafficMirrorFilterRule
	err := ctx.ReadResource("aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilterRule resources.
type trafficMirrorFilterRuleState struct {
	// A description of the traffic mirror filter rule.
	Description *string `pulumi:"description"`
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange *TrafficMirrorFilterRuleDestinationPortRange `pulumi:"destinationPortRange"`
	// The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int `pulumi:"protocol"`
	// The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction *string `pulumi:"ruleAction"`
	// The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber *int `pulumi:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `pulumi:"sourceCidrBlock"`
	// The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange *TrafficMirrorFilterRuleSourcePortRange `pulumi:"sourcePortRange"`
	// The direction of traffic to be captured. Valid values are `ingress` and `egress`
	TrafficDirection *string `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
}

type TrafficMirrorFilterRuleState struct {
	// A description of the traffic mirror filter rule.
	Description pulumi.StringPtrInput
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrInput
	// The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrInput
	// The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringPtrInput
	// The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntPtrInput
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringPtrInput
	// The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrInput
	// The direction of traffic to be captured. Valid values are `ingress` and `egress`
	TrafficDirection pulumi.StringPtrInput
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringPtrInput
}

func (TrafficMirrorFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleState)(nil)).Elem()
}

type trafficMirrorFilterRuleArgs struct {
	// A description of the traffic mirror filter rule.
	Description *string `pulumi:"description"`
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange *TrafficMirrorFilterRuleDestinationPortRange `pulumi:"destinationPortRange"`
	// The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int `pulumi:"protocol"`
	// The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction string `pulumi:"ruleAction"`
	// The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber int `pulumi:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock string `pulumi:"sourceCidrBlock"`
	// The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange *TrafficMirrorFilterRuleSourcePortRange `pulumi:"sourcePortRange"`
	// The direction of traffic to be captured. Valid values are `ingress` and `egress`
	TrafficDirection string `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
}

// The set of arguments for constructing a TrafficMirrorFilterRule resource.
type TrafficMirrorFilterRuleArgs struct {
	// A description of the traffic mirror filter rule.
	Description pulumi.StringPtrInput
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringInput
	// The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrInput
	// The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrInput
	// The action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringInput
	// The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntInput
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringInput
	// The source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrInput
	// The direction of traffic to be captured. Valid values are `ingress` and `egress`
	TrafficDirection pulumi.StringInput
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringInput
}

func (TrafficMirrorFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleArgs)(nil)).Elem()
}
