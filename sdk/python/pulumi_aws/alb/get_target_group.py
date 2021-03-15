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
    'GetTargetGroupResult',
    'AwaitableGetTargetGroupResult',
    'get_target_group',
]

@pulumi.output_type
class GetTargetGroupResult:
    """
    A collection of values returned by getTargetGroup.
    """
    def __init__(__self__, arn=None, arn_suffix=None, deregistration_delay=None, health_check=None, id=None, lambda_multi_value_headers_enabled=None, load_balancing_algorithm_type=None, name=None, port=None, protocol=None, protocol_version=None, proxy_protocol_v2=None, slow_start=None, stickiness=None, tags=None, target_type=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if arn_suffix and not isinstance(arn_suffix, str):
            raise TypeError("Expected argument 'arn_suffix' to be a str")
        pulumi.set(__self__, "arn_suffix", arn_suffix)
        if deregistration_delay and not isinstance(deregistration_delay, int):
            raise TypeError("Expected argument 'deregistration_delay' to be a int")
        pulumi.set(__self__, "deregistration_delay", deregistration_delay)
        if health_check and not isinstance(health_check, dict):
            raise TypeError("Expected argument 'health_check' to be a dict")
        pulumi.set(__self__, "health_check", health_check)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lambda_multi_value_headers_enabled and not isinstance(lambda_multi_value_headers_enabled, bool):
            raise TypeError("Expected argument 'lambda_multi_value_headers_enabled' to be a bool")
        pulumi.set(__self__, "lambda_multi_value_headers_enabled", lambda_multi_value_headers_enabled)
        if load_balancing_algorithm_type and not isinstance(load_balancing_algorithm_type, str):
            raise TypeError("Expected argument 'load_balancing_algorithm_type' to be a str")
        pulumi.set(__self__, "load_balancing_algorithm_type", load_balancing_algorithm_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if protocol_version and not isinstance(protocol_version, str):
            raise TypeError("Expected argument 'protocol_version' to be a str")
        pulumi.set(__self__, "protocol_version", protocol_version)
        if proxy_protocol_v2 and not isinstance(proxy_protocol_v2, bool):
            raise TypeError("Expected argument 'proxy_protocol_v2' to be a bool")
        pulumi.set(__self__, "proxy_protocol_v2", proxy_protocol_v2)
        if slow_start and not isinstance(slow_start, int):
            raise TypeError("Expected argument 'slow_start' to be a int")
        pulumi.set(__self__, "slow_start", slow_start)
        if stickiness and not isinstance(stickiness, dict):
            raise TypeError("Expected argument 'stickiness' to be a dict")
        pulumi.set(__self__, "stickiness", stickiness)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_type and not isinstance(target_type, str):
            raise TypeError("Expected argument 'target_type' to be a str")
        pulumi.set(__self__, "target_type", target_type)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="arnSuffix")
    def arn_suffix(self) -> str:
        return pulumi.get(self, "arn_suffix")

    @property
    @pulumi.getter(name="deregistrationDelay")
    def deregistration_delay(self) -> int:
        return pulumi.get(self, "deregistration_delay")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> 'outputs.GetTargetGroupHealthCheckResult':
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lambdaMultiValueHeadersEnabled")
    def lambda_multi_value_headers_enabled(self) -> bool:
        return pulumi.get(self, "lambda_multi_value_headers_enabled")

    @property
    @pulumi.getter(name="loadBalancingAlgorithmType")
    def load_balancing_algorithm_type(self) -> str:
        return pulumi.get(self, "load_balancing_algorithm_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolVersion")
    def protocol_version(self) -> str:
        return pulumi.get(self, "protocol_version")

    @property
    @pulumi.getter(name="proxyProtocolV2")
    def proxy_protocol_v2(self) -> bool:
        return pulumi.get(self, "proxy_protocol_v2")

    @property
    @pulumi.getter(name="slowStart")
    def slow_start(self) -> int:
        return pulumi.get(self, "slow_start")

    @property
    @pulumi.getter
    def stickiness(self) -> 'outputs.GetTargetGroupStickinessResult':
        return pulumi.get(self, "stickiness")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> str:
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetTargetGroupResult(GetTargetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetGroupResult(
            arn=self.arn,
            arn_suffix=self.arn_suffix,
            deregistration_delay=self.deregistration_delay,
            health_check=self.health_check,
            id=self.id,
            lambda_multi_value_headers_enabled=self.lambda_multi_value_headers_enabled,
            load_balancing_algorithm_type=self.load_balancing_algorithm_type,
            name=self.name,
            port=self.port,
            protocol=self.protocol,
            protocol_version=self.protocol_version,
            proxy_protocol_v2=self.proxy_protocol_v2,
            slow_start=self.slow_start,
            stickiness=self.stickiness,
            tags=self.tags,
            target_type=self.target_type,
            vpc_id=self.vpc_id)


def get_target_group(arn: Optional[str] = None,
                     name: Optional[str] = None,
                     tags: Optional[Mapping[str, str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetGroupResult:
    """
    > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.

    Provides information about a Load Balancer Target Group.

    This data source can prove useful when a module accepts an LB Target Group as an
    input variable and needs to know its attributes. It can also be used to get the ARN of
    an LB Target Group for use in other resources, given LB Target Group name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    lb_tg_arn = config.get("lbTgArn")
    if lb_tg_arn is None:
        lb_tg_arn = ""
    lb_tg_name = config.get("lbTgName")
    if lb_tg_name is None:
        lb_tg_name = ""
    test = aws.lb.get_target_group(arn=lb_tg_arn,
        name=lb_tg_name)
    ```


    :param str arn: The full ARN of the target group.
    :param str name: The unique name of the target group.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:alb/getTargetGroup:getTargetGroup', __args__, opts=opts, typ=GetTargetGroupResult).value

    return AwaitableGetTargetGroupResult(
        arn=__ret__.arn,
        arn_suffix=__ret__.arn_suffix,
        deregistration_delay=__ret__.deregistration_delay,
        health_check=__ret__.health_check,
        id=__ret__.id,
        lambda_multi_value_headers_enabled=__ret__.lambda_multi_value_headers_enabled,
        load_balancing_algorithm_type=__ret__.load_balancing_algorithm_type,
        name=__ret__.name,
        port=__ret__.port,
        protocol=__ret__.protocol,
        protocol_version=__ret__.protocol_version,
        proxy_protocol_v2=__ret__.proxy_protocol_v2,
        slow_start=__ret__.slow_start,
        stickiness=__ret__.stickiness,
        tags=__ret__.tags,
        target_type=__ret__.target_type,
        vpc_id=__ret__.vpc_id)
