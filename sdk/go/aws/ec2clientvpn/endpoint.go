// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2clientvpn

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2clientvpn"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2clientvpn.NewEndpoint(ctx, "example", &ec2clientvpn.EndpointArgs{
// 			AuthenticationOptions: ec2clientvpn.EndpointAuthenticationOptionArray{
// 				&ec2clientvpn.EndpointAuthenticationOptionArgs{
// 					RootCertificateChainArn: pulumi.Any(aws_acm_certificate.Root_cert.Arn),
// 					Type:                    pulumi.String("certificate-authentication"),
// 				},
// 			},
// 			ClientCidrBlock: pulumi.String("10.0.0.0/16"),
// 			ConnectionLogOptions: &ec2clientvpn.EndpointConnectionLogOptionsArgs{
// 				CloudwatchLogGroup:  pulumi.Any(aws_cloudwatch_log_group.Lg.Name),
// 				CloudwatchLogStream: pulumi.Any(aws_cloudwatch_log_stream.Ls.Name),
// 				Enabled:             pulumi.Bool(true),
// 			},
// 			Description:          pulumi.String("clientvpn-example"),
// 			ServerCertificateArn: pulumi.Any(aws_acm_certificate.Cert.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// The ARN of the Client VPN endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionArrayOutput `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringOutput `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsOutput `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringOutput `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolPtrOutput `pulumi:"splitTunnel"`
	// The current state of the Client VPN endpoint.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringPtrOutput `pulumi:"transportProtocol"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.AuthenticationOptions == nil {
		return nil, errors.New("missing required argument 'AuthenticationOptions'")
	}
	if args == nil || args.ClientCidrBlock == nil {
		return nil, errors.New("missing required argument 'ClientCidrBlock'")
	}
	if args == nil || args.ConnectionLogOptions == nil {
		return nil, errors.New("missing required argument 'ConnectionLogOptions'")
	}
	if args == nil || args.ServerCertificateArn == nil {
		return nil, errors.New("missing required argument 'ServerCertificateArn'")
	}
	if args == nil {
		args = &EndpointArgs{}
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws:ec2clientvpn/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws:ec2clientvpn/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// The ARN of the Client VPN endpoint.
	Arn *string `pulumi:"arn"`
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions []EndpointAuthenticationOption `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock *string `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions *EndpointConnectionLogOptions `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description *string `pulumi:"description"`
	// The DNS name to be used by clients when establishing their VPN session.
	DnsName *string `pulumi:"dnsName"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers []string `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn *string `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel *bool `pulumi:"splitTunnel"`
	// The current state of the Client VPN endpoint.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol *string `pulumi:"transportProtocol"`
}

type EndpointState struct {
	// The ARN of the Client VPN endpoint.
	Arn pulumi.StringPtrInput
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionArrayInput
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringPtrInput
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsPtrInput
	// Name of the repository.
	Description pulumi.StringPtrInput
	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringPtrInput
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayInput
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringPtrInput
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolPtrInput
	// The current state of the Client VPN endpoint.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions []EndpointAuthenticationOption `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock string `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptions `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description *string `pulumi:"description"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers []string `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn string `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel *bool `pulumi:"splitTunnel"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol *string `pulumi:"transportProtocol"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionArrayInput
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringInput
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsInput
	// Name of the repository.
	Description pulumi.StringPtrInput
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayInput
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringInput
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}
