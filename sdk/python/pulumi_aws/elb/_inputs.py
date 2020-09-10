# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'LoadBalancerAccessLogsArgs',
    'LoadBalancerHealthCheckArgs',
    'LoadBalancerListenerArgs',
    'LoadBalancerPolicyPolicyAttributeArgs',
    'SslNegotiationPolicyAttributeArgs',
]

@pulumi.input_type
class LoadBalancerAccessLogsArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 bucket_prefix: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 interval: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "bucket", bucket)
        if bucket_prefix is not None:
            pulumi.set(__self__, "bucket_prefix", bucket_prefix)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket_prefix")

    @bucket_prefix.setter
    def bucket_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_prefix", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "interval", value)


@pulumi.input_type
class LoadBalancerHealthCheckArgs:
    def __init__(__self__, *,
                 healthy_threshold: pulumi.Input[float],
                 interval: pulumi.Input[float],
                 target: pulumi.Input[str],
                 timeout: pulumi.Input[float],
                 unhealthy_threshold: pulumi.Input[float]):
        pulumi.set(__self__, "healthy_threshold", healthy_threshold)
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "timeout", timeout)
        pulumi.set(__self__, "unhealthy_threshold", unhealthy_threshold)

    @property
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> pulumi.Input[float]:
        return pulumi.get(self, "healthy_threshold")

    @healthy_threshold.setter
    def healthy_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "healthy_threshold", value)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[float]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[float]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Input[float]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: pulumi.Input[float]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> pulumi.Input[float]:
        return pulumi.get(self, "unhealthy_threshold")

    @unhealthy_threshold.setter
    def unhealthy_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "unhealthy_threshold", value)


@pulumi.input_type
class LoadBalancerListenerArgs:
    def __init__(__self__, *,
                 instance_port: pulumi.Input[float],
                 instance_protocol: pulumi.Input[str],
                 lb_port: pulumi.Input[float],
                 lb_protocol: pulumi.Input[str],
                 ssl_certificate_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "instance_port", instance_port)
        pulumi.set(__self__, "instance_protocol", instance_protocol)
        pulumi.set(__self__, "lb_port", lb_port)
        pulumi.set(__self__, "lb_protocol", lb_protocol)
        if ssl_certificate_id is not None:
            pulumi.set(__self__, "ssl_certificate_id", ssl_certificate_id)

    @property
    @pulumi.getter(name="instancePort")
    def instance_port(self) -> pulumi.Input[float]:
        return pulumi.get(self, "instance_port")

    @instance_port.setter
    def instance_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "instance_port", value)

    @property
    @pulumi.getter(name="instanceProtocol")
    def instance_protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_protocol")

    @instance_protocol.setter
    def instance_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_protocol", value)

    @property
    @pulumi.getter(name="lbPort")
    def lb_port(self) -> pulumi.Input[float]:
        return pulumi.get(self, "lb_port")

    @lb_port.setter
    def lb_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "lb_port", value)

    @property
    @pulumi.getter(name="lbProtocol")
    def lb_protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "lb_protocol")

    @lb_protocol.setter
    def lb_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_protocol", value)

    @property
    @pulumi.getter(name="sslCertificateId")
    def ssl_certificate_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_certificate_id")

    @ssl_certificate_id.setter
    def ssl_certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_certificate_id", value)


@pulumi.input_type
class LoadBalancerPolicyPolicyAttributeArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SslNegotiationPolicyAttributeArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


