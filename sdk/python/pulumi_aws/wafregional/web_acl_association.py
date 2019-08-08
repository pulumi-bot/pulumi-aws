# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class WebAclAssociation(pulumi.CustomResource):
    resource_arn: pulumi.Output[str]
    """
    ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
    """
    web_acl_id: pulumi.Output[str]
    """
    The ID of the WAF Regional WebACL to create an association.
    """
    def __init__(__self__, resource_name, opts=None, resource_arn=None, web_acl_id=None, __name__=None, __opts__=None):
        """
        Manages an association with WAF Regional Web ACL.
        
        > **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        :param pulumi.Input[str] web_acl_id: The ID of the WAF Regional WebACL to create an association.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_web_acl_association.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if resource_arn is None:
            raise TypeError("Missing required property 'resource_arn'")
        __props__['resource_arn'] = resource_arn
        if web_acl_id is None:
            raise TypeError("Missing required property 'web_acl_id'")
        __props__['web_acl_id'] = web_acl_id
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(WebAclAssociation, __self__).__init__(
            'aws:wafregional/webAclAssociation:WebAclAssociation',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

