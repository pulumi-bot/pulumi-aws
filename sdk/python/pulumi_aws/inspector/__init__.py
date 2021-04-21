# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .assessment_target import *
from .assessment_template import *
from .get_rules_packages import *
from .resource_group import *

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:inspector/assessmentTarget:AssessmentTarget":
                return AssessmentTarget(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:inspector/assessmentTemplate:AssessmentTemplate":
                return AssessmentTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:inspector/resourceGroup:ResourceGroup":
                return ResourceGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "inspector/assessmentTarget", _module_instance)
    pulumi.runtime.register_resource_module("aws", "inspector/assessmentTemplate", _module_instance)
    pulumi.runtime.register_resource_module("aws", "inspector/resourceGroup", _module_instance)

_register_module()
