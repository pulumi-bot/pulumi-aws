# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Connection(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the connection.
    """
    aws_device: pulumi.Output[str]
    """
    The Direct Connect endpoint on which the physical connection terminates.
    """
    bandwidth: pulumi.Output[str]
    """
    The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
    """
    has_logical_redundancy: pulumi.Output[str]
    """
    Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
    """
    jumbo_frame_capable: pulumi.Output[bool]
    """
    Boolean value representing if jumbo frames have been enabled for this connection.
    """
    location: pulumi.Output[str]
    """
    The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
    """
    name: pulumi.Output[str]
    """
    The name of the connection.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, bandwidth=None, location=None, name=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a Connection of Direct Connect.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_connection.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if bandwidth is None:
            raise TypeError("Missing required property 'bandwidth'")
        __props__['bandwidth'] = bandwidth

        if location is None:
            raise TypeError("Missing required property 'location'")
        __props__['location'] = location

        __props__['name'] = name

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['aws_device'] = None
        __props__['has_logical_redundancy'] = None
        __props__['jumbo_frame_capable'] = None

        super(Connection, __self__).__init__(
            'aws:directconnect/connection:Connection',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

