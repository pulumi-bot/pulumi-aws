# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .api import *
from .api_mapping import *
from .authorizer import *
from .deployment import *
from .domain_name import *
from .integration import *
from .integration_response import *
from .model import *
from .route import *
from .route_response import *
from .stage import *
from .vpc_link import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:apigatewayv2/api:Api":
                return Api(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/apiMapping:ApiMapping":
                return ApiMapping(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/authorizer:Authorizer":
                return Authorizer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/deployment:Deployment":
                return Deployment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/domainName:DomainName":
                return DomainName(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/integration:Integration":
                return Integration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/integrationResponse:IntegrationResponse":
                return IntegrationResponse(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/model:Model":
                return Model(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/route:Route":
                return Route(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/routeResponse:RouteResponse":
                return RouteResponse(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/stage:Stage":
                return Stage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:apigatewayv2/vpcLink:VpcLink":
                return VpcLink(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/api", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/apiMapping", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/authorizer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/deployment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/domainName", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/integration", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/integrationResponse", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/model", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/route", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/routeResponse", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/stage", _module_instance)
    pulumi.runtime.register_resource_module("aws", "apigatewayv2/vpcLink", _module_instance)

_register_module()
