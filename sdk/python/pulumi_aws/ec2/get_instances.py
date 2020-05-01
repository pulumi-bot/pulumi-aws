# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, filters=None, id=None, ids=None, instance_state_names=None, instance_tags=None, private_ips=None, public_ips=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        IDs of instances found through the filter
        """
        if instance_state_names and not isinstance(instance_state_names, list):
            raise TypeError("Expected argument 'instance_state_names' to be a list")
        __self__.instance_state_names = instance_state_names
        if instance_tags and not isinstance(instance_tags, dict):
            raise TypeError("Expected argument 'instance_tags' to be a dict")
        __self__.instance_tags = instance_tags
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        __self__.private_ips = private_ips
        """
        Private IP addresses of instances found through the filter
        """
        if public_ips and not isinstance(public_ips, list):
            raise TypeError("Expected argument 'public_ips' to be a list")
        __self__.public_ips = public_ips
        """
        Public IP addresses of instances found through the filter
        """
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

def get_instances(filters=None,instance_state_names=None,instance_tags=None,opts=None):
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

    test_instances = aws.ec2.get_instances(filters=[{
            "name": "instance.group-id",
            "values": ["sg-12345678"],
        }],
        instance_state_names=[
            "running",
            "stopped",
        ],
        instance_tags={
            "Role": "HardWorker",
        })
    test_eip = []
    for range in [{"value": i} for i in range(0, len(test_instances.ids))]:
        test_eip.append(aws.ec2.Eip(f"testEip-{range['value']}", instance=test_instances.ids[range["value"]]))
    ```



    :param list filters: One or more name/value pairs to use as filters. There are
           several valid keys, for a full reference, check out
           [describe-instances in the AWS CLI reference][1].
    :param list instance_state_names: A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
    :param dict instance_tags: A mapping of tags, each pair of which must
           exactly match a pair on desired instances.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['instanceStateNames'] = instance_state_names
    __args__['instanceTags'] = instance_tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getInstances:getInstances', __args__, opts=opts).value

    return AwaitableGetInstancesResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_state_names=__ret__.get('instanceStateNames'),
        instance_tags=__ret__.get('instanceTags'),
        private_ips=__ret__.get('privateIps'),
        public_ips=__ret__.get('publicIps'))
