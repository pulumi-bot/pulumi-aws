# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .account import *
from .get_organization import *
from .get_organizational_units import *
from .organization import *
from .organizational_unit import *
from .policy import *
from .policy_attachment import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:organizations/account:Account":
                return Account(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:organizations/organization:Organization":
                return Organization(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:organizations/organizationalUnit:OrganizationalUnit":
                return OrganizationalUnit(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:organizations/policy:Policy":
                return Policy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:organizations/policyAttachment:PolicyAttachment":
                return PolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "organizations/account", _module_instance)
    pulumi.runtime.register_resource_module("aws", "organizations/organization", _module_instance)
    pulumi.runtime.register_resource_module("aws", "organizations/organizationalUnit", _module_instance)
    pulumi.runtime.register_resource_module("aws", "organizations/policy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "organizations/policyAttachment", _module_instance)

_register_module()
