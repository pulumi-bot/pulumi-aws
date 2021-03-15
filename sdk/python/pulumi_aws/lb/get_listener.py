# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetListenerResult',
    'AwaitableGetListenerResult',
    'get_listener',
]

@pulumi.output_type
class GetListenerResult:
    """
    A collection of values returned by getListener.
    """
    def __init__(__self__, arn=None, certificate_arn=None, default_actions=None, id=None, load_balancer_arn=None, port=None, protocol=None, ssl_policy=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if certificate_arn and not isinstance(certificate_arn, str):
            raise TypeError("Expected argument 'certificate_arn' to be a str")
        pulumi.set(__self__, "certificate_arn", certificate_arn)
        if default_actions and not isinstance(default_actions, list):
            raise TypeError("Expected argument 'default_actions' to be a list")
        pulumi.set(__self__, "default_actions", default_actions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balancer_arn and not isinstance(load_balancer_arn, str):
            raise TypeError("Expected argument 'load_balancer_arn' to be a str")
        pulumi.set(__self__, "load_balancer_arn", load_balancer_arn)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if ssl_policy and not isinstance(ssl_policy, str):
            raise TypeError("Expected argument 'ssl_policy' to be a str")
        pulumi.set(__self__, "ssl_policy", ssl_policy)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> str:
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="defaultActions")
    def default_actions(self) -> Sequence['outputs.GetListenerDefaultActionResult']:
        return pulumi.get(self, "default_actions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalancerArn")
    def load_balancer_arn(self) -> str:
        return pulumi.get(self, "load_balancer_arn")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> str:
        return pulumi.get(self, "ssl_policy")


class AwaitableGetListenerResult(GetListenerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetListenerResult(
            arn=self.arn,
            certificate_arn=self.certificate_arn,
            default_actions=self.default_actions,
            id=self.id,
            load_balancer_arn=self.load_balancer_arn,
            port=self.port,
            protocol=self.protocol,
            ssl_policy=self.ssl_policy)


def get_listener(arn: Optional[str] = None,
                 load_balancer_arn: Optional[str] = None,
                 port: Optional[int] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetListenerResult:
    """
    > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.

    Provides information about a Load Balancer Listener.

    This data source can prove useful when a module accepts an LB Listener as an
    input variable and needs to know the LB it is attached to, or other
    information specific to the listener in question.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    listener_arn = config.require("listenerArn")
    listener = aws.lb.get_listener(arn=listener_arn)
    selected = aws.lb.get_load_balancer(name="default-public")
    selected443 = aws.lb.get_listener(load_balancer_arn=selected.arn,
        port=443)
    ```


    :param str arn: The arn of the listener. Required if `load_balancer_arn` and `port` is not set.
    :param str load_balancer_arn: The arn of the load balancer. Required if `arn` is not set.
    :param int port: The port of the listener. Required if `arn` is not set.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['loadBalancerArn'] = load_balancer_arn
    __args__['port'] = port
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lb/getListener:getListener', __args__, opts=opts, typ=GetListenerResult).value

    return AwaitableGetListenerResult(
        arn=__ret__.arn,
        certificate_arn=__ret__.certificate_arn,
        default_actions=__ret__.default_actions,
        id=__ret__.id,
        load_balancer_arn=__ret__.load_balancer_arn,
        port=__ret__.port,
        protocol=__ret__.protocol,
        ssl_policy=__ret__.ssl_policy)
