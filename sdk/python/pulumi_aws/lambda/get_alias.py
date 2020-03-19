# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAliasResult:
    """
    A collection of values returned by getAlias.
    """
    def __init__(__self__, arn=None, description=None, function_name=None, function_version=None, id=None, invoke_arn=None, name=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) identifying the Lambda function alias.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of alias.
        """
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        __self__.function_name = function_name
        if function_version and not isinstance(function_version, str):
            raise TypeError("Expected argument 'function_version' to be a str")
        __self__.function_version = function_version
        """
        Lambda function version which the alias uses.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if invoke_arn and not isinstance(invoke_arn, str):
            raise TypeError("Expected argument 'invoke_arn' to be a str")
        __self__.invoke_arn = invoke_arn
        """
        The ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's `uri`.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetAliasResult(GetAliasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAliasResult(
            arn=self.arn,
            description=self.description,
            function_name=self.function_name,
            function_version=self.function_version,
            id=self.id,
            invoke_arn=self.invoke_arn,
            name=self.name)

def get_alias(function_name=None,name=None,opts=None):
    """
    Provides information about a Lambda Alias.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_alias.html.markdown.


    :param str function_name: Name of the aliased Lambda function.
    :param str name: Name of the Lambda alias.
    """
    __args__ = dict()


    __args__['functionName'] = function_name
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lambda/getAlias:getAlias', __args__, opts=opts).value

    return AwaitableGetAliasResult(
        arn=__ret__.get('arn'),
        description=__ret__.get('description'),
        function_name=__ret__.get('functionName'),
        function_version=__ret__.get('functionVersion'),
        id=__ret__.get('id'),
        invoke_arn=__ret__.get('invokeArn'),
        name=__ret__.get('name'))
