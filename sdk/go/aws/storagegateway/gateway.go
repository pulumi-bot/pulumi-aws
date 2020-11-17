// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.
//
// > NOTE: The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).
//
// ## Example Usage
// ### File Gateway
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("FILE_S3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Tape Gateway
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress:  pulumi.String("1.2.3.4"),
// 			GatewayName:       pulumi.String("example"),
// 			GatewayTimezone:   pulumi.String("GMT"),
// 			GatewayType:       pulumi.String("VTL"),
// 			MediumChangerType: pulumi.String("AWS-Gateway-VTL"),
// 			TapeDriveType:     pulumi.String("IBM-ULT3580-TD5"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Volume Gateway (Cached)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("CACHED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Volume Gateway (Stored)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("STORED"),
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
// `aws_storagegateway_gateway` can be imported by using the gateway Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:storagegateway/gateway:Gateway example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
// ```
//
//  Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the Terraform configuration or use [`ignore_changes`](/docs/configuration/resources.html#ignore_changes) to hide the difference, e.g. hcl resource "aws_storagegateway_gateway" "example" {
//
// # ... other configuration ...
//
//  gateway_ip_address = aws_instance.sgw.private_ip
//
// # There is no Storage Gateway API for reading gateway_ip_address
//
//  lifecycle {
//
//  ignore_changes = ["gateway_ip_address"]
//
//  } }
type Gateway struct {
	pulumi.CustomResourceState

	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringOutput `pulumi:"activationKey"`
	// Amazon Resource Name (ARN) of the gateway.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrOutput `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrOutput `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrOutput `pulumi:"cloudwatchLogGroupArn"`
	// Identifier of the gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringOutput `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName pulumi.StringOutput `pulumi:"gatewayName"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringOutput `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrOutput `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrOutput `pulumi:"gatewayVpcEndpoint"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`.
	MediumChangerType pulumi.StringPtrOutput `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrOutput `pulumi:"smbActiveDirectorySettings"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrOutput `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringOutput `pulumi:"smbSecurityStrategy"`
	// Key-value mapping of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrOutput `pulumi:"tapeDriveType"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil || args.GatewayName == nil {
		return nil, errors.New("missing required argument 'GatewayName'")
	}
	if args == nil || args.GatewayTimezone == nil {
		return nil, errors.New("missing required argument 'GatewayTimezone'")
	}
	if args == nil {
		args = &GatewayArgs{}
	}
	var resource Gateway
	err := ctx.RegisterResource("aws:storagegateway/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("aws:storagegateway/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey *string `pulumi:"activationKey"`
	// Amazon Resource Name (ARN) of the gateway.
	Arn *string `pulumi:"arn"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec *int `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec *int `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// Identifier of the gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName *string `pulumi:"gatewayName"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone *string `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType *string `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint *string `pulumi:"gatewayVpcEndpoint"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`.
	MediumChangerType *string `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings *GatewaySmbActiveDirectorySettings `pulumi:"smbActiveDirectorySettings"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword *string `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy *string `pulumi:"smbSecurityStrategy"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType *string `pulumi:"tapeDriveType"`
}

type GatewayState struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the gateway.
	Arn pulumi.StringPtrInput
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// Identifier of the gateway.
	GatewayId pulumi.StringPtrInput
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringPtrInput
	// Name of the gateway.
	GatewayName pulumi.StringPtrInput
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringPtrInput
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrInput
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrInput
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`.
	MediumChangerType pulumi.StringPtrInput
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrInput
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrInput
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey *string `pulumi:"activationKey"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec *int `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec *int `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone string `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType *string `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint *string `pulumi:"gatewayVpcEndpoint"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`.
	MediumChangerType *string `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings *GatewaySmbActiveDirectorySettings `pulumi:"smbActiveDirectorySettings"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword *string `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy *string `pulumi:"smbSecurityStrategy"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType *string `pulumi:"tapeDriveType"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringPtrInput
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringPtrInput
	// Name of the gateway.
	GatewayName pulumi.StringInput
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringInput
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrInput
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrInput
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`.
	MediumChangerType pulumi.StringPtrInput
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrInput
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrInput
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil)).Elem()
}

func (i Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

type GatewayOutput struct {
	*pulumi.OutputState
}

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayOutput)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}
