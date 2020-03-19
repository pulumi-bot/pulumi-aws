# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VpcEndpointServiceAllowedPrinciple(pulumi.CustomResource):
    principal_arn: pulumi.Output[str]
    """
    The ARN of the principal to allow permissions.
    """
    vpc_endpoint_service_id: pulumi.Output[str]
    """
    The ID of the VPC endpoint service to allow permission.
    """
    def __init__(__self__, resource_name, opts=None, principal_arn=None, vpc_endpoint_service_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to allow a principal to discover a VPC endpoint service.

        > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
        both a standalone VPC Endpoint Service Allowed Principal resource
        and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
        a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
        and will overwrite the association.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_service_allowed_principal.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] principal_arn: The ARN of the principal to allow permissions.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC endpoint service to allow permission.
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

            if principal_arn is None:
                raise TypeError("Missing required property 'principal_arn'")
            __props__['principal_arn'] = principal_arn
            if vpc_endpoint_service_id is None:
                raise TypeError("Missing required property 'vpc_endpoint_service_id'")
            __props__['vpc_endpoint_service_id'] = vpc_endpoint_service_id
        super(VpcEndpointServiceAllowedPrinciple, __self__).__init__(
            'aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, principal_arn=None, vpc_endpoint_service_id=None):
        """
        Get an existing VpcEndpointServiceAllowedPrinciple resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] principal_arn: The ARN of the principal to allow permissions.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC endpoint service to allow permission.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["principal_arn"] = principal_arn
        __props__["vpc_endpoint_service_id"] = vpc_endpoint_service_id
        return VpcEndpointServiceAllowedPrinciple(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

