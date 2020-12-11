// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Global Accelerator listener.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/globalaccelerator"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccelerator, err := globalaccelerator.NewAccelerator(ctx, "exampleAccelerator", &globalaccelerator.AcceleratorArgs{
// 			IpAddressType: pulumi.String("IPV4"),
// 			Enabled:       pulumi.Bool(true),
// 			Attributes: &globalaccelerator.AcceleratorAttributesArgs{
// 				FlowLogsEnabled:  pulumi.Bool(true),
// 				FlowLogsS3Bucket: pulumi.String("example-bucket"),
// 				FlowLogsS3Prefix: pulumi.String("flow-logs/"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = globalaccelerator.NewListener(ctx, "exampleListener", &globalaccelerator.ListenerArgs{
// 			AcceleratorArn: exampleAccelerator.ID(),
// 			ClientAffinity: pulumi.String("SOURCE_IP"),
// 			Protocol:       pulumi.String("TCP"),
// 			PortRanges: globalaccelerator.ListenerPortRangeArray{
// 				&globalaccelerator.ListenerPortRangeArgs{
// 					FromPort: pulumi.Int(80),
// 					ToPort:   pulumi.Int(80),
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
// Global Accelerator listeners can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:globalaccelerator/listener:Listener example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
// ```
type Listener struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringOutput `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrOutput `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayOutput `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorArn == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorArn'")
	}
	if args.PortRanges == nil {
		return nil, errors.New("invalid value for required argument 'PortRanges'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	var resource Listener
	err := ctx.RegisterResource("aws:globalaccelerator/listener:Listener", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:globalaccelerator/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn *string `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol *string `pulumi:"protocol"`
}

type ListenerState struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringPtrInput
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrInput
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayInput
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn string `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol string `pulumi:"protocol"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringInput
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrInput
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayInput
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringInput
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
