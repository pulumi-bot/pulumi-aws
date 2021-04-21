# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .cluster import *
from .cluster_instance import *
from .cluster_parameter_group import *
from .cluster_snapshot import *
from .event_subscription import *
from .get_engine_version import *
from .get_orderable_db_instance import *
from .parameter_group import *
from .subnet_group import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:neptune/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/clusterInstance:ClusterInstance":
                return ClusterInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/clusterParameterGroup:ClusterParameterGroup":
                return ClusterParameterGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/clusterSnapshot:ClusterSnapshot":
                return ClusterSnapshot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/eventSubscription:EventSubscription":
                return EventSubscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/parameterGroup:ParameterGroup":
                return ParameterGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:neptune/subnetGroup:SubnetGroup":
                return SubnetGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "neptune/cluster", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/clusterInstance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/clusterParameterGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/clusterSnapshot", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/eventSubscription", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/parameterGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "neptune/subnetGroup", _module_instance)

_register_module()
