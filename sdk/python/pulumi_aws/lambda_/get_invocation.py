# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetInvocationResult',
    'AwaitableGetInvocationResult',
    'get_invocation',
]

@pulumi.output_type
class GetInvocationResult:
    """
    A collection of values returned by getInvocation.
    """
    def __init__(__self__, function_name=None, id=None, input=None, qualifier=None, result=None):
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        pulumi.set(__self__, "function_name", function_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if input and not isinstance(input, str):
            raise TypeError("Expected argument 'input' to be a str")
        pulumi.set(__self__, "input", input)
        if qualifier and not isinstance(qualifier, str):
            raise TypeError("Expected argument 'qualifier' to be a str")
        pulumi.set(__self__, "qualifier", qualifier)
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> str:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def input(self) -> str:
        return pulumi.get(self, "input")

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[str]:
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableGetInvocationResult(GetInvocationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInvocationResult(
            function_name=self.function_name,
            id=self.id,
            input=self.input,
            qualifier=self.qualifier,
            result=self.result)


def get_invocation(function_name: Optional[str] = None,
                   input: Optional[str] = None,
                   qualifier: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInvocationResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['input'] = input
    __args__['qualifier'] = qualifier
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lambda/getInvocation:getInvocation', __args__, opts=opts, typ=GetInvocationResult).value

    return AwaitableGetInvocationResult(
        function_name=__ret__.function_name,
        id=__ret__.id,
        input=__ret__.input,
        qualifier=__ret__.qualifier,
        result=__ret__.result)
