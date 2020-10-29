// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Traffic mirror session.\
// Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
type TrafficMirrorSession struct {
	pulumi.CustomResourceState

	// The ARN of the traffic mirror session.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the traffic mirror session.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength pulumi.IntPtrOutput `pulumi:"packetLength"`
	// - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber pulumi.IntOutput `pulumi:"sessionNumber"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId pulumi.StringOutput `pulumi:"trafficMirrorTargetId"`
	// - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId pulumi.IntOutput `pulumi:"virtualNetworkId"`
}

// NewTrafficMirrorSession registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorSession(ctx *pulumi.Context,
	name string, args *TrafficMirrorSessionArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorSession, error) {
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil || args.SessionNumber == nil {
		return nil, errors.New("missing required argument 'SessionNumber'")
	}
	if args == nil || args.TrafficMirrorFilterId == nil {
		return nil, errors.New("missing required argument 'TrafficMirrorFilterId'")
	}
	if args == nil || args.TrafficMirrorTargetId == nil {
		return nil, errors.New("missing required argument 'TrafficMirrorTargetId'")
	}
	if args == nil {
		args = &TrafficMirrorSessionArgs{}
	}
	var resource TrafficMirrorSession
	err := ctx.RegisterResource("aws:ec2/trafficMirrorSession:TrafficMirrorSession", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorSession gets an existing TrafficMirrorSession resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorSessionState, opts ...pulumi.ResourceOption) (*TrafficMirrorSession, error) {
	var resource TrafficMirrorSession
	err := ctx.ReadResource("aws:ec2/trafficMirrorSession:TrafficMirrorSession", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorSession resources.
type trafficMirrorSessionState struct {
	// The ARN of the traffic mirror session.
	Arn *string `pulumi:"arn"`
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength *int `pulumi:"packetLength"`
	// - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber *int `pulumi:"sessionNumber"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId *string `pulumi:"trafficMirrorTargetId"`
	// - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId *int `pulumi:"virtualNetworkId"`
}

type TrafficMirrorSessionState struct {
	// The ARN of the traffic mirror session.
	Arn pulumi.StringPtrInput
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId pulumi.StringPtrInput
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength pulumi.IntPtrInput
	// - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber pulumi.IntPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId pulumi.StringPtrInput
	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId pulumi.StringPtrInput
	// - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId pulumi.IntPtrInput
}

func (TrafficMirrorSessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorSessionState)(nil)).Elem()
}

type trafficMirrorSessionArgs struct {
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength *int `pulumi:"packetLength"`
	// - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber int `pulumi:"sessionNumber"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId string `pulumi:"trafficMirrorTargetId"`
	// - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId *int `pulumi:"virtualNetworkId"`
}

// The set of arguments for constructing a TrafficMirrorSession resource.
type TrafficMirrorSessionArgs struct {
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId pulumi.StringInput
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength pulumi.IntPtrInput
	// - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber pulumi.IntInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId pulumi.StringInput
	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId pulumi.StringInput
	// - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId pulumi.IntPtrInput
}

func (TrafficMirrorSessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorSessionArgs)(nil)).Elem()
}
