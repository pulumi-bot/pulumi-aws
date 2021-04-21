# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .policy import *
from .scheduled_action import *
from .target import *
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
            if typ == "aws:appautoscaling/policy:Policy":
                return Policy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:appautoscaling/scheduledAction:ScheduledAction":
                return ScheduledAction(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:appautoscaling/target:Target":
                return Target(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "appautoscaling/policy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "appautoscaling/scheduledAction", _module_instance)
    pulumi.runtime.register_resource_module("aws", "appautoscaling/target", _module_instance)

_register_module()
