# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class DefaultVpcDhcpOptions(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the DHCP Options Set.
    """
    domain_name: pulumi.Output[str]
    domain_name_servers: pulumi.Output[str]
    netbios_name_servers: pulumi.Output[list]
    """
    List of NETBIOS name servers.
    """
    netbios_node_type: pulumi.Output[str]
    """
    The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
    """
    ntp_servers: pulumi.Output[str]
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the DHCP options set.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, netbios_name_servers=None, netbios_node_type=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage the [default AWS DHCP Options Set](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html#AmazonDNS)
        in the current region.

        Each AWS region comes with a default set of DHCP options.
        **This is an advanced resource**, and has special caveats to be aware of when
        using it. Please read this document in its entirety before using this resource.

        The `ec2.DefaultVpcDhcpOptions` behaves differently from normal resources, in that
        this provider does not _create_ this resource, but instead "adopts" it
        into management.

        ## Example Usage

        Basic usage with tags:

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.ec2.DefaultVpcDhcpOptions("default", tags={
            "Name": "Default DHCP Option Set",
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] netbios_name_servers: List of NETBIOS name servers.
        :param pulumi.Input[str] netbios_node_type: The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['netbiosNameServers'] = netbios_name_servers
            __props__['netbiosNodeType'] = netbios_node_type
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['domain_name'] = None
            __props__['domain_name_servers'] = None
            __props__['ntp_servers'] = None
            __props__['owner_id'] = None
        super(DefaultVpcDhcpOptions, __self__).__init__(
            'aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, domain_name=None, domain_name_servers=None, netbios_name_servers=None, netbios_node_type=None, ntp_servers=None, owner_id=None, tags=None):
        """
        Get an existing DefaultVpcDhcpOptions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the DHCP Options Set.
        :param pulumi.Input[list] netbios_name_servers: List of NETBIOS name servers.
        :param pulumi.Input[str] netbios_node_type: The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the DHCP options set.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["domain_name"] = domain_name
        __props__["domain_name_servers"] = domain_name_servers
        __props__["netbios_name_servers"] = netbios_name_servers
        __props__["netbios_node_type"] = netbios_node_type
        __props__["ntp_servers"] = ntp_servers
        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        return DefaultVpcDhcpOptions(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
