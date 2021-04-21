# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .cluster import *
from .cluster_instance import *
from .cluster_parameter_group import *
from .cluster_snapshot import *
from .get_engine_version import *
from .get_orderable_db_instance import *
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
            if typ == "aws:docdb/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:docdb/clusterInstance:ClusterInstance":
                return ClusterInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:docdb/clusterParameterGroup:ClusterParameterGroup":
                return ClusterParameterGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:docdb/clusterSnapshot:ClusterSnapshot":
                return ClusterSnapshot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:docdb/subnetGroup:SubnetGroup":
                return SubnetGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "docdb/cluster", _module_instance)
    pulumi.runtime.register_resource_module("aws", "docdb/clusterInstance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "docdb/clusterParameterGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "docdb/clusterSnapshot", _module_instance)
    pulumi.runtime.register_resource_module("aws", "docdb/subnetGroup", _module_instance)

_register_module()
