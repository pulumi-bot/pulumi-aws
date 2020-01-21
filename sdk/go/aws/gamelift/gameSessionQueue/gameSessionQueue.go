// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gameSessionQueue

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/gamelift/GameSessionQueuePlayerLatencyPolicy"
)

// Provides an Gamelift Game Session Queue resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_game_session_queue.html.markdown.
type GameSessionQueue struct {
	pulumi.CustomResourceState

	// Game Session Queue ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayOutput `pulumi:"destinations"`
	// Name of the session queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies gameliftGameSessionQueuePlayerLatencyPolicy.GameSessionQueuePlayerLatencyPolicyArrayOutput `pulumi:"playerLatencyPolicies"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrOutput `pulumi:"timeoutInSeconds"`
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
	// Game Session Queue ARN.
	Arn *string `pulumi:"arn"`
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []string `pulumi:"destinations"`
	// Name of the session queue.
	Name *string `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies []gameliftGameSessionQueuePlayerLatencyPolicy.GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
}

type GameSessionQueueState struct {
	// Game Session Queue ARN.
	Arn pulumi.StringPtrInput
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayInput
	// Name of the session queue.
	Name pulumi.StringPtrInput
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies gameliftGameSessionQueuePlayerLatencyPolicy.GameSessionQueuePlayerLatencyPolicyArrayInput
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrInput
}

func (GameSessionQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueState)(nil)).Elem()
}

type gameSessionQueueArgs struct {
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []string `pulumi:"destinations"`
	// Name of the session queue.
	Name *string `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies []gameliftGameSessionQueuePlayerLatencyPolicy.GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
}

// The set of arguments for constructing a GameSessionQueue resource.
type GameSessionQueueArgs struct {
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayInput
	// Name of the session queue.
	Name pulumi.StringPtrInput
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies gameliftGameSessionQueuePlayerLatencyPolicy.GameSessionQueuePlayerLatencyPolicyArrayInput
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrInput
}

func (GameSessionQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueArgs)(nil)).Elem()
}

