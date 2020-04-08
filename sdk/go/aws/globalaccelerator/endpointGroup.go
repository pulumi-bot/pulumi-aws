// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Global Accelerator endpoint group.
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayOutput `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrOutput `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringPtrOutput `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntPtrOutput `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrOutput `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrOutput `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrOutput `pulumi:"trafficDialPercentage"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	if args == nil {
		args = &EndpointGroupArgs{}
	}
	var resource EndpointGroup
	err := ctx.RegisterResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `pulumi:"listenerArn"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

type EndpointGroupState struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringPtrInput
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringPtrInput
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringPtrInput
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrInput
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrInput
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []EndpointGroupEndpointConfigurationArgs `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn string `pulumi:"listenerArn"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArgsArrayInput
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringPtrInput
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringPtrInput
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringInput
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrInput
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}
