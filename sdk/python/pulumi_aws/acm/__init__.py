# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .certificate import *
from .certificate_validation import *
from .get_certificate import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:acm/certificate:Certificate":
                return Certificate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:acm/certificateValidation:CertificateValidation":
                return CertificateValidation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "acm/certificate", _module_instance)
    pulumi.runtime.register_resource_module("aws", "acm/certificateValidation", _module_instance)

_register_module()
