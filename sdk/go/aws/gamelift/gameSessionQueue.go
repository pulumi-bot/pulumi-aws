// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GameSessionQueue struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput                            `pulumi:"arn"`
	Destinations          pulumi.StringArrayOutput                       `pulumi:"destinations"`
	Name                  pulumi.StringOutput                            `pulumi:"name"`
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayOutput `pulumi:"playerLatencyPolicies"`
	Tags                  pulumi.StringMapOutput                         `pulumi:"tags"`
	TimeoutInSeconds      pulumi.IntPtrOutput                            `pulumi:"timeoutInSeconds"`
}

// NewGameSessionQueue registers a new resource with the given unique name, arguments, and options.
func NewGameSessionQueue(ctx *pulumi.Context,
	name string, args *GameSessionQueueArgs, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	if args == nil {
		args = &GameSessionQueueArgs{}
	}
	var resource GameSessionQueue
	err := ctx.RegisterResource("aws:gamelift/gameSessionQueue:GameSessionQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameSessionQueue gets an existing GameSessionQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameSessionQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameSessionQueueState, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	var resource GameSessionQueue
	err := ctx.ReadResource("aws:gamelift/gameSessionQueue:GameSessionQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameSessionQueue resources.
type gameSessionQueueState struct {
	Arn                   *string                               `pulumi:"arn"`
	Destinations          []string                              `pulumi:"destinations"`
	Name                  *string                               `pulumi:"name"`
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	Tags                  map[string]string                     `pulumi:"tags"`
	TimeoutInSeconds      *int                                  `pulumi:"timeoutInSeconds"`
}

type GameSessionQueueState struct {
	Arn                   pulumi.StringPtrInput
	Destinations          pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayInput
	Tags                  pulumi.StringMapInput
	TimeoutInSeconds      pulumi.IntPtrInput
}

func (GameSessionQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueState)(nil)).Elem()
}

type gameSessionQueueArgs struct {
	Destinations          []string                              `pulumi:"destinations"`
	Name                  *string                               `pulumi:"name"`
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	Tags                  map[string]string                     `pulumi:"tags"`
	TimeoutInSeconds      *int                                  `pulumi:"timeoutInSeconds"`
}

// The set of arguments for constructing a GameSessionQueue resource.
type GameSessionQueueArgs struct {
	Destinations          pulumi.StringArrayInput
	Name                  pulumi.StringPtrInput
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayInput
	Tags                  pulumi.StringMapInput
	TimeoutInSeconds      pulumi.IntPtrInput
}

func (GameSessionQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueArgs)(nil)).Elem()
}
