# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .cluster import *
from .fargate_profile import *
from .get_cluster import *
from .get_cluster_auth import *
from .node_group import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:eks/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:eks/fargateProfile:FargateProfile":
                return FargateProfile(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:eks/nodeGroup:NodeGroup":
                return NodeGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "eks/cluster", _module_instance)
    pulumi.runtime.register_resource_module("aws", "eks/fargateProfile", _module_instance)
    pulumi.runtime.register_resource_module("aws", "eks/nodeGroup", _module_instance)

_register_module()
