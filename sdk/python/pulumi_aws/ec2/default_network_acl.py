# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DefaultNetworkAcl(pulumi.CustomResource):
    default_network_acl_id: pulumi.Output[str]
    """
    The Network ACL ID to manage. This
    attribute is exported from `aws_vpc`, or manually found via the AWS Console.
    """
    egress: pulumi.Output[list]
    """
    Specifies an egress rule. Parameters defined below.
    """
    ingress: pulumi.Output[list]
    """
    Specifies an ingress rule. Parameters defined below.
    """
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the Default Network ACL
    """
    subnet_ids: pulumi.Output[list]
    """
    A list of Subnet IDs to apply the ACL to. See the
    notes below on managing Subnets in the Default Network ACL
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the associated VPC
    """
    def __init__(__self__, resource_name, opts=None, default_network_acl_id=None, egress=None, ingress=None, subnet_ids=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage the default AWS Network ACL. VPC Only.
        
        Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
        destroyed. **This is an advanced resource**, and has special caveats to be aware
        of when using it. Please read this document in its entirety before using this
        resource.
        
        The `aws_default_network_acl` behaves differently from normal resources, in that
        Terraform does not _create_ this resource, but instead attempts to "adopt" it
        into management. We can do this because each VPC created has a Default Network
        ACL that cannot be destroyed, and is created with a known set of default rules.
        
        When Terraform first adopts the Default Network ACL, it **immediately removes all
        rules in the ACL**. It then proceeds to create any rules specified in the
        configuration. This step is required so that only the rules specified in the
        configuration are created.
        
        This resource treats its inline rules as absolute; only the rules defined
        inline are created, and any additions/removals external to this resource will
        result in diffs being shown. For these reasons, this resource is incompatible with the
        `aws_network_acl_rule` resource.
        
        For more information about Network ACLs, see the AWS Documentation on
        [Network ACLs][aws-network-acls].
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_network_acl_id: The Network ACL ID to manage. This
               attribute is exported from `aws_vpc`, or manually found via the AWS Console.
        :param pulumi.Input[list] egress: Specifies an egress rule. Parameters defined below.
        :param pulumi.Input[list] ingress: Specifies an ingress rule. Parameters defined below.
        :param pulumi.Input[list] subnet_ids: A list of Subnet IDs to apply the ACL to. See the
               notes below on managing Subnets in the Default Network ACL
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

        if default_network_acl_id is None:
            raise TypeError("Missing required property 'default_network_acl_id'")
        __props__['default_network_acl_id'] = default_network_acl_id

        __props__['egress'] = egress

        __props__['ingress'] = ingress

        __props__['subnet_ids'] = subnet_ids

        __props__['tags'] = tags

        __props__['owner_id'] = None
        __props__['vpc_id'] = None

        super(DefaultNetworkAcl, __self__).__init__(
            'aws:ec2/defaultNetworkAcl:DefaultNetworkAcl',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

