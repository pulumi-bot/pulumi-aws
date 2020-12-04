# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .capacity_provider import *
from .cluster import *
from .get_cluster import *
from .get_container_definition import *
from .get_service import *
from .get_task_definition import *
from .service import *
from .task_definition import *
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
            if typ == "aws:ecs/capacityProvider:CapacityProvider":
                return CapacityProvider(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ecs/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ecs/service:Service":
                return Service(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ecs/taskDefinition:TaskDefinition":
                return TaskDefinition(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "ecs/capacityProvider", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ecs/cluster", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ecs/service", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ecs/taskDefinition", _module_instance)

_register_module()
