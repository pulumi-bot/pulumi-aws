# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 authentication_options: pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]],
                 client_cidr_block: pulumi.Input[str],
                 connection_log_options: pulumi.Input['EndpointConnectionLogOptionsArgs'],
                 server_certificate_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 split_tunnel: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transport_protocol: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        :param pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input['EndpointConnectionLogOptionsArgs'] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        pulumi.set(__self__, "authentication_options", authentication_options)
        pulumi.set(__self__, "client_cidr_block", client_cidr_block)
        pulumi.set(__self__, "connection_log_options", connection_log_options)
        pulumi.set(__self__, "server_certificate_arn", server_certificate_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns_servers is not None:
            pulumi.set(__self__, "dns_servers", dns_servers)
        if split_tunnel is not None:
            pulumi.set(__self__, "split_tunnel", split_tunnel)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if transport_protocol is not None:
            pulumi.set(__self__, "transport_protocol", transport_protocol)

    @property
    @pulumi.getter(name="authenticationOptions")
    def authentication_options(self) -> pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]]:
        """
        Information about the authentication method to be used to authenticate clients.
        """
        return pulumi.get(self, "authentication_options")

    @authentication_options.setter
    def authentication_options(self, value: pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]]):
        pulumi.set(self, "authentication_options", value)

    @property
    @pulumi.getter(name="clientCidrBlock")
    def client_cidr_block(self) -> pulumi.Input[str]:
        """
        The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        """
        return pulumi.get(self, "client_cidr_block")

    @client_cidr_block.setter
    def client_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_cidr_block", value)

    @property
    @pulumi.getter(name="connectionLogOptions")
    def connection_log_options(self) -> pulumi.Input['EndpointConnectionLogOptionsArgs']:
        """
        Information about the client connection logging options.
        """
        return pulumi.get(self, "connection_log_options")

    @connection_log_options.setter
    def connection_log_options(self, value: pulumi.Input['EndpointConnectionLogOptionsArgs']):
        pulumi.set(self, "connection_log_options", value)

    @property
    @pulumi.getter(name="serverCertificateArn")
    def server_certificate_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the ACM server certificate.
        """
        return pulumi.get(self, "server_certificate_arn")

    @server_certificate_arn.setter
    def server_certificate_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_certificate_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        """
        return pulumi.get(self, "dns_servers")

    @dns_servers.setter
    def dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_servers", value)

    @property
    @pulumi.getter(name="splitTunnel")
    def split_tunnel(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        """
        return pulumi.get(self, "split_tunnel")

    @split_tunnel.setter
    def split_tunnel(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "split_tunnel", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="transportProtocol")
    def transport_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        return pulumi.get(self, "transport_protocol")

    @transport_protocol.setter
    def transport_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport_protocol", value)


@pulumi.input_type
class _EndpointState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 authentication_options: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]]] = None,
                 client_cidr_block: Optional[pulumi.Input[str]] = None,
                 connection_log_options: Optional[pulumi.Input['EndpointConnectionLogOptionsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_certificate_arn: Optional[pulumi.Input[str]] = None,
                 split_tunnel: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transport_protocol: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Endpoint resources.
        :param pulumi.Input[str] arn: The ARN of the Client VPN endpoint.
        :param pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input['EndpointConnectionLogOptionsArgs'] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[str] dns_name: The DNS name to be used by clients when establishing their VPN session.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[str] status: The current state of the Client VPN endpoint.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if authentication_options is not None:
            pulumi.set(__self__, "authentication_options", authentication_options)
        if client_cidr_block is not None:
            pulumi.set(__self__, "client_cidr_block", client_cidr_block)
        if connection_log_options is not None:
            pulumi.set(__self__, "connection_log_options", connection_log_options)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns_name is not None:
            pulumi.set(__self__, "dns_name", dns_name)
        if dns_servers is not None:
            pulumi.set(__self__, "dns_servers", dns_servers)
        if server_certificate_arn is not None:
            pulumi.set(__self__, "server_certificate_arn", server_certificate_arn)
        if split_tunnel is not None:
            pulumi.set(__self__, "split_tunnel", split_tunnel)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if transport_protocol is not None:
            pulumi.set(__self__, "transport_protocol", transport_protocol)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Client VPN endpoint.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="authenticationOptions")
    def authentication_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]]]:
        """
        Information about the authentication method to be used to authenticate clients.
        """
        return pulumi.get(self, "authentication_options")

    @authentication_options.setter
    def authentication_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAuthenticationOptionArgs']]]]):
        pulumi.set(self, "authentication_options", value)

    @property
    @pulumi.getter(name="clientCidrBlock")
    def client_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        """
        return pulumi.get(self, "client_cidr_block")

    @client_cidr_block.setter
    def client_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cidr_block", value)

    @property
    @pulumi.getter(name="connectionLogOptions")
    def connection_log_options(self) -> Optional[pulumi.Input['EndpointConnectionLogOptionsArgs']]:
        """
        Information about the client connection logging options.
        """
        return pulumi.get(self, "connection_log_options")

    @connection_log_options.setter
    def connection_log_options(self, value: Optional[pulumi.Input['EndpointConnectionLogOptionsArgs']]):
        pulumi.set(self, "connection_log_options", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS name to be used by clients when establishing their VPN session.
        """
        return pulumi.get(self, "dns_name")

    @dns_name.setter
    def dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_name", value)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        """
        return pulumi.get(self, "dns_servers")

    @dns_servers.setter
    def dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_servers", value)

    @property
    @pulumi.getter(name="serverCertificateArn")
    def server_certificate_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the ACM server certificate.
        """
        return pulumi.get(self, "server_certificate_arn")

    @server_certificate_arn.setter
    def server_certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_certificate_arn", value)

    @property
    @pulumi.getter(name="splitTunnel")
    def split_tunnel(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        """
        return pulumi.get(self, "split_tunnel")

    @split_tunnel.setter
    def split_tunnel(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "split_tunnel", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the Client VPN endpoint.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="transportProtocol")
    def transport_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        return pulumi.get(self, "transport_protocol")

    @transport_protocol.setter
    def transport_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport_protocol", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAuthenticationOptionArgs']]]]] = None,
                 client_cidr_block: Optional[pulumi.Input[str]] = None,
                 connection_log_options: Optional[pulumi.Input[pulumi.InputType['EndpointConnectionLogOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_certificate_arn: Optional[pulumi.Input[str]] = None,
                 split_tunnel: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transport_protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
        [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2clientvpn.Endpoint("example",
            description="clientvpn-example",
            server_certificate_arn=aws_acm_certificate["cert"]["arn"],
            client_cidr_block="10.0.0.0/16",
            authentication_options=[aws.ec2clientvpn.EndpointAuthenticationOptionArgs(
                type="certificate-authentication",
                root_certificate_chain_arn=aws_acm_certificate["root_cert"]["arn"],
            )],
            connection_log_options=aws.ec2clientvpn.EndpointConnectionLogOptionsArgs(
                enabled=True,
                cloudwatch_log_group=aws_cloudwatch_log_group["lg"]["name"],
                cloudwatch_log_stream=aws_cloudwatch_log_stream["ls"]["name"],
            ))
        ```

        ## Import

        AWS Client VPN endpoints can be imported using the `id` value found via `aws ec2 describe-client-vpn-endpoints`, e.g.

        ```sh
         $ pulumi import aws:ec2clientvpn/endpoint:Endpoint example cvpn-endpoint-0ac3a1abbccddd666
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAuthenticationOptionArgs']]]] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input[pulumi.InputType['EndpointConnectionLogOptionsArgs']] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
        [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2clientvpn.Endpoint("example",
            description="clientvpn-example",
            server_certificate_arn=aws_acm_certificate["cert"]["arn"],
            client_cidr_block="10.0.0.0/16",
            authentication_options=[aws.ec2clientvpn.EndpointAuthenticationOptionArgs(
                type="certificate-authentication",
                root_certificate_chain_arn=aws_acm_certificate["root_cert"]["arn"],
            )],
            connection_log_options=aws.ec2clientvpn.EndpointConnectionLogOptionsArgs(
                enabled=True,
                cloudwatch_log_group=aws_cloudwatch_log_group["lg"]["name"],
                cloudwatch_log_stream=aws_cloudwatch_log_stream["ls"]["name"],
            ))
        ```

        ## Import

        AWS Client VPN endpoints can be imported using the `id` value found via `aws ec2 describe-client-vpn-endpoints`, e.g.

        ```sh
         $ pulumi import aws:ec2clientvpn/endpoint:Endpoint example cvpn-endpoint-0ac3a1abbccddd666
        ```

        :param str resource_name_: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAuthenticationOptionArgs']]]]] = None,
                 client_cidr_block: Optional[pulumi.Input[str]] = None,
                 connection_log_options: Optional[pulumi.Input[pulumi.InputType['EndpointConnectionLogOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_certificate_arn: Optional[pulumi.Input[str]] = None,
                 split_tunnel: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transport_protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointArgs.__new__(EndpointArgs)

            if authentication_options is None and not opts.urn:
                raise TypeError("Missing required property 'authentication_options'")
            __props__.__dict__["authentication_options"] = authentication_options
            if client_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'client_cidr_block'")
            __props__.__dict__["client_cidr_block"] = client_cidr_block
            if connection_log_options is None and not opts.urn:
                raise TypeError("Missing required property 'connection_log_options'")
            __props__.__dict__["connection_log_options"] = connection_log_options
            __props__.__dict__["description"] = description
            __props__.__dict__["dns_servers"] = dns_servers
            if server_certificate_arn is None and not opts.urn:
                raise TypeError("Missing required property 'server_certificate_arn'")
            __props__.__dict__["server_certificate_arn"] = server_certificate_arn
            __props__.__dict__["split_tunnel"] = split_tunnel
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["transport_protocol"] = transport_protocol
            __props__.__dict__["arn"] = None
            __props__.__dict__["dns_name"] = None
            __props__.__dict__["status"] = None
        super(Endpoint, __self__).__init__(
            'aws:ec2clientvpn/endpoint:Endpoint',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            authentication_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAuthenticationOptionArgs']]]]] = None,
            client_cidr_block: Optional[pulumi.Input[str]] = None,
            connection_log_options: Optional[pulumi.Input[pulumi.InputType['EndpointConnectionLogOptionsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            server_certificate_arn: Optional[pulumi.Input[str]] = None,
            split_tunnel: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transport_protocol: Optional[pulumi.Input[str]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Client VPN endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAuthenticationOptionArgs']]]] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input[pulumi.InputType['EndpointConnectionLogOptionsArgs']] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[str] dns_name: The DNS name to be used by clients when establishing their VPN session.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[str] status: The current state of the Client VPN endpoint.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointState.__new__(_EndpointState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["authentication_options"] = authentication_options
        __props__.__dict__["client_cidr_block"] = client_cidr_block
        __props__.__dict__["connection_log_options"] = connection_log_options
        __props__.__dict__["description"] = description
        __props__.__dict__["dns_name"] = dns_name
        __props__.__dict__["dns_servers"] = dns_servers
        __props__.__dict__["server_certificate_arn"] = server_certificate_arn
        __props__.__dict__["split_tunnel"] = split_tunnel
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["transport_protocol"] = transport_protocol
        return Endpoint(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Client VPN endpoint.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationOptions")
    def authentication_options(self) -> pulumi.Output[Sequence['outputs.EndpointAuthenticationOption']]:
        """
        Information about the authentication method to be used to authenticate clients.
        """
        return pulumi.get(self, "authentication_options")

    @property
    @pulumi.getter(name="clientCidrBlock")
    def client_cidr_block(self) -> pulumi.Output[str]:
        """
        The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        """
        return pulumi.get(self, "client_cidr_block")

    @property
    @pulumi.getter(name="connectionLogOptions")
    def connection_log_options(self) -> pulumi.Output['outputs.EndpointConnectionLogOptions']:
        """
        Information about the client connection logging options.
        """
        return pulumi.get(self, "connection_log_options")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        """
        The DNS name to be used by clients when establishing their VPN session.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter(name="serverCertificateArn")
    def server_certificate_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the ACM server certificate.
        """
        return pulumi.get(self, "server_certificate_arn")

    @property
    @pulumi.getter(name="splitTunnel")
    def split_tunnel(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        """
        return pulumi.get(self, "split_tunnel")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current state of the Client VPN endpoint.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="transportProtocol")
    def transport_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The transport protocol to be used by the VPN session. Default value is `udp`.
        """
        return pulumi.get(self, "transport_protocol")

