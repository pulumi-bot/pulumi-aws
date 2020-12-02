# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .get_listener import *
from .get_load_balancer import *
from .get_target_group import *
from .listener import *
from .listener_certificate import *
from .listener_rule import *
from .load_balancer import *
from .target_group import *
from .target_group_attachment import *
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
            if typ == "aws:alb/listener:Listener":
                return Listener(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:alb/listenerCertificate:ListenerCertificate":
                return ListenerCertificate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:alb/listenerRule:ListenerRule":
                return ListenerRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:alb/loadBalancer:LoadBalancer":
                return LoadBalancer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:alb/targetGroup:TargetGroup":
                return TargetGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:alb/targetGroupAttachment:TargetGroupAttachment":
                return TargetGroupAttachment(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "alb/listener", _module_instance)
    pulumi.runtime.register_resource_module("aws", "alb/listenerCertificate", _module_instance)
    pulumi.runtime.register_resource_module("aws", "alb/listenerRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "alb/loadBalancer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "alb/targetGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "alb/targetGroupAttachment", _module_instance)

_register_module()
