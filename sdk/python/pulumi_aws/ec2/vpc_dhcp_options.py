# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpcDhcpOptions(pulumi.CustomResource):
    domain_name: pulumi.Output[str]
    """
    the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
    """
    domain_name_servers: pulumi.Output[list]
    """
    List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
    """
    netbios_name_servers: pulumi.Output[list]
    """
    List of NETBIOS name servers.
    """
    netbios_node_type: pulumi.Output[str]
    """
    The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
    """
    ntp_servers: pulumi.Output[list]
    """
    List of NTP servers to configure.
    """
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the DHCP options set.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, domain_name=None, domain_name_servers=None, netbios_name_servers=None, netbios_node_type=None, ntp_servers=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a VPC DHCP Options resource.
        
        ## Remarks
        
        * Notice that all arguments are optional but you have to specify at least one argument.
        * `domain_name_servers`, `netbios_name_servers`, `ntp_servers` are limited by AWS to maximum four servers only.
        * To actually use the DHCP Options Set you need to associate it to a VPC using [`aws_vpc_dhcp_options_association`](https://www.terraform.io/docs/providers/aws/r/vpc_dhcp_options_association.html).
        * If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
        * In most cases unless you're configuring your own DNS you'll want to set `domain_name_servers` to `AmazonProvidedDNS`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
        :param pulumi.Input[list] domain_name_servers: List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
        :param pulumi.Input[list] netbios_name_servers: List of NETBIOS name servers.
        :param pulumi.Input[str] netbios_node_type: The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
        :param pulumi.Input[list] ntp_servers: List of NTP servers to configure.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_dhcp_options.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['domain_name'] = domain_name
            __props__['domain_name_servers'] = domain_name_servers
            __props__['netbios_name_servers'] = netbios_name_servers
            __props__['netbios_node_type'] = netbios_node_type
            __props__['ntp_servers'] = ntp_servers
            __props__['tags'] = tags
            __props__['owner_id'] = None
        super(VpcDhcpOptions, __self__).__init__(
            'aws:ec2/vpcDhcpOptions:VpcDhcpOptions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, domain_name=None, domain_name_servers=None, netbios_name_servers=None, netbios_node_type=None, ntp_servers=None, owner_id=None, tags=None):
        """
        Get an existing VpcDhcpOptions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
        :param pulumi.Input[list] domain_name_servers: List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
        :param pulumi.Input[list] netbios_name_servers: List of NETBIOS name servers.
        :param pulumi.Input[str] netbios_node_type: The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
        :param pulumi.Input[list] ntp_servers: List of NTP servers to configure.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the DHCP options set.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_dhcp_options.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["domain_name"] = domain_name
        __props__["domain_name_servers"] = domain_name_servers
        __props__["netbios_name_servers"] = netbios_name_servers
        __props__["netbios_node_type"] = netbios_node_type
        __props__["ntp_servers"] = ntp_servers
        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        return VpcDhcpOptions(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

