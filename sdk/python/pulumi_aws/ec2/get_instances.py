# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, filters=None, id=None, ids=None, instance_state_names=None, instance_tags=None, private_ips=None, public_ips=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_state_names and not isinstance(instance_state_names, list):
            raise TypeError("Expected argument 'instance_state_names' to be a list")
        pulumi.set(__self__, "instance_state_names", instance_state_names)
        if instance_tags and not isinstance(instance_tags, dict):
            raise TypeError("Expected argument 'instance_tags' to be a dict")
        pulumi.set(__self__, "instance_tags", instance_tags)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if public_ips and not isinstance(public_ips, list):
            raise TypeError("Expected argument 'public_ips' to be a list")
        pulumi.set(__self__, "public_ips", public_ips)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetInstancesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        IDs of instances found through the filter
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceStateNames")
    def instance_state_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_state_names")

    @property
    @pulumi.getter(name="instanceTags")
    def instance_tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "instance_tags")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Sequence[str]:
        """
        Private IP addresses of instances found through the filter
        """
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="publicIps")
    def public_ips(self) -> Sequence[str]:
        """
        Public IP addresses of instances found through the filter
        """
        return pulumi.get(self, "public_ips")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            instance_state_names=self.instance_state_names,
            instance_tags=self.instance_tags,
            private_ips=self.private_ips,
            public_ips=self.public_ips)


def get_instances(filters: Optional[Sequence[pulumi.InputType['GetInstancesFilterArgs']]] = None,
                  instance_state_names: Optional[Sequence[str]] = None,
                  instance_tags: Optional[Mapping[str, str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
    e.g. to allow easier migration from another management solution
    or to make it easier for an operator to connect through bastion host(s).

    > **Note:** It's strongly discouraged to use this data source for querying ephemeral
    instances (e.g. managed via autoscaling group), as the output may change at any time
    and you'd need to re-run `apply` every time an instance comes up or dies.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    test_instances = aws.ec2.get_instances(instance_tags={
            "Role": "HardWorker",
        },
        filters=[aws.ec2.GetInstancesFilterArgs(
            name="instance.group-id",
            values=["sg-12345678"],
        )],
        instance_state_names=[
            "running",
            "stopped",
        ])
    test_eip = []
    for range in [{"value": i} for i in range(0, len(test_instances.ids))]:
        test_eip.append(aws.ec2.Eip(f"testEip-{range['value']}", instance=test_instances.ids[range["value"]]))
    ```


    :param Sequence[pulumi.InputType['GetInstancesFilterArgs']] filters: One or more name/value pairs to use as filters. There are
           several valid keys, for a full reference, check out
           [describe-instances in the AWS CLI reference][1].
    :param Sequence[str] instance_state_names: A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
    :param Mapping[str, str] instance_tags: A map of tags, each pair of which must
           exactly match a pair on desired instances.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['instanceStateNames'] = instance_state_names
    __args__['instanceTags'] = instance_tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_state_names=__ret__.instance_state_names,
        instance_tags=__ret__.instance_tags,
        private_ips=__ret__.private_ips,
        public_ips=__ret__.public_ips)
