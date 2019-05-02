# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkAcl(pulumi.CustomResource):
    egress: pulumi.Output[list]
    """
    Specifies an egress rule. Parameters defined below.
    This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
    """
    ingress: pulumi.Output[list]
    """
    Specifies an ingress rule. Parameters defined below.
    This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
    """
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the network ACL.
    """
    subnet_ids: pulumi.Output[list]
    """
    A list of Subnet IDs to apply the ACL to
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the associated VPC.
    """
    def __init__(__self__, resource_name, opts=None, egress=None, ingress=None, subnet_ids=None, tags=None, vpc_id=None, __name__=None, __opts__=None):
        """
        Provides an network ACL resource. You might set up network ACLs with rules similar
        to your security groups in order to add an additional layer of security to your VPC.
        
        > **NOTE on Network ACLs and Network ACL Rules:** Terraform currently
        provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
        defined in-line. At this time you cannot use a Network ACL with in-line rules
        in conjunction with any Network ACL Rule resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] egress: Specifies an egress rule. Parameters defined below.
               This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        :param pulumi.Input[list] ingress: Specifies an ingress rule. Parameters defined below.
               This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        :param pulumi.Input[list] subnet_ids: A list of Subnet IDs to apply the ACL to
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the associated VPC.
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

        __props__['egress'] = egress

        __props__['ingress'] = ingress

        __props__['subnet_ids'] = subnet_ids

        __props__['tags'] = tags

        if vpc_id is None:
            raise TypeError("Missing required property 'vpc_id'")
        __props__['vpc_id'] = vpc_id

        __props__['owner_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NetworkAcl, __self__).__init__(
            'aws:ec2/networkAcl:NetworkAcl',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

