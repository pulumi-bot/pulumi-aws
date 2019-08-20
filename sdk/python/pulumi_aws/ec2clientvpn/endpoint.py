# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Endpoint(pulumi.CustomResource):
    authentication_options: pulumi.Output[dict]
    """
    Information about the authentication method to be used to authenticate clients.
    """
    client_cidr_block: pulumi.Output[str]
    """
    The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
    """
    connection_log_options: pulumi.Output[dict]
    """
    Information about the client connection logging options.
    """
    description: pulumi.Output[str]
    """
    Name of the repository.
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name to be used by clients when establishing their VPN session.
    """
    dns_servers: pulumi.Output[list]
    """
    Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
    """
    server_certificate_arn: pulumi.Output[str]
    """
    The ARN of the ACM server certificate.
    """
    split_tunnel: pulumi.Output[bool]
    """
    Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
    """
    status: pulumi.Output[str]
    """
    The current state of the Client VPN endpoint.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    transport_protocol: pulumi.Output[str]
    """
    The transport protocol to be used by the VPN session. Default value is `udp`.
    """
    def __init__(__self__, resource_name, opts=None, authentication_options=None, client_cidr_block=None, connection_log_options=None, description=None, dns_servers=None, server_certificate_arn=None, split_tunnel=None, tags=None, transport_protocol=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
        [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input[dict] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[list] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_client_vpn_endpoint.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if authentication_options is None:
                raise TypeError("Missing required property 'authentication_options'")
            __props__['authentication_options'] = authentication_options
            if client_cidr_block is None:
                raise TypeError("Missing required property 'client_cidr_block'")
            __props__['client_cidr_block'] = client_cidr_block
            if connection_log_options is None:
                raise TypeError("Missing required property 'connection_log_options'")
            __props__['connection_log_options'] = connection_log_options
            __props__['description'] = description
            __props__['dns_servers'] = dns_servers
            if server_certificate_arn is None:
                raise TypeError("Missing required property 'server_certificate_arn'")
            __props__['server_certificate_arn'] = server_certificate_arn
            __props__['split_tunnel'] = split_tunnel
            __props__['tags'] = tags
            __props__['transport_protocol'] = transport_protocol
            __props__['dns_name'] = None
            __props__['status'] = None
        super(Endpoint, __self__).__init__(
            'aws:ec2clientvpn/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, authentication_options=None, client_cidr_block=None, connection_log_options=None, description=None, dns_name=None, dns_servers=None, server_certificate_arn=None, split_tunnel=None, status=None, tags=None, transport_protocol=None):
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] authentication_options: Information about the authentication method to be used to authenticate clients.
        :param pulumi.Input[str] client_cidr_block: The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        :param pulumi.Input[dict] connection_log_options: Information about the client connection logging options.
        :param pulumi.Input[str] description: Name of the repository.
        :param pulumi.Input[str] dns_name: The DNS name to be used by clients when establishing their VPN session.
        :param pulumi.Input[list] dns_servers: Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        :param pulumi.Input[str] server_certificate_arn: The ARN of the ACM server certificate.
        :param pulumi.Input[bool] split_tunnel: Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        :param pulumi.Input[str] status: The current state of the Client VPN endpoint.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transport_protocol: The transport protocol to be used by the VPN session. Default value is `udp`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_client_vpn_endpoint.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["authentication_options"] = authentication_options
        __props__["client_cidr_block"] = client_cidr_block
        __props__["connection_log_options"] = connection_log_options
        __props__["description"] = description
        __props__["dns_name"] = dns_name
        __props__["dns_servers"] = dns_servers
        __props__["server_certificate_arn"] = server_certificate_arn
        __props__["split_tunnel"] = split_tunnel
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["transport_protocol"] = transport_protocol
        return Endpoint(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

