# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .application import *
from .custom_layer import *
from .ganglia_layer import *
from .haproxy_layer import *
from .instance import *
from .java_app_layer import *
from .memcached_layer import *
from .mysql_layer import *
from .nodejs_app_layer import *
from .permission import *
from .php_app_layer import *
from .rails_app_layer import *
from .rds_db_instance import *
from .stack import *
from .static_web_layer import *
from .user_profile import *
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
            if typ == "aws:opsworks/application:Application":
                return Application(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/customLayer:CustomLayer":
                return CustomLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/gangliaLayer:GangliaLayer":
                return GangliaLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/haproxyLayer:HaproxyLayer":
                return HaproxyLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/instance:Instance":
                return Instance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/javaAppLayer:JavaAppLayer":
                return JavaAppLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/memcachedLayer:MemcachedLayer":
                return MemcachedLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/mysqlLayer:MysqlLayer":
                return MysqlLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/nodejsAppLayer:NodejsAppLayer":
                return NodejsAppLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/permission:Permission":
                return Permission(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/phpAppLayer:PhpAppLayer":
                return PhpAppLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/railsAppLayer:RailsAppLayer":
                return RailsAppLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/rdsDbInstance:RdsDbInstance":
                return RdsDbInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/stack:Stack":
                return Stack(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/staticWebLayer:StaticWebLayer":
                return StaticWebLayer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:opsworks/userProfile:UserProfile":
                return UserProfile(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "opsworks/application", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/customLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/gangliaLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/haproxyLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/instance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/javaAppLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/memcachedLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/mysqlLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/nodejsAppLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/permission", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/phpAppLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/railsAppLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/rdsDbInstance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/stack", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/staticWebLayer", _module_instance)
    pulumi.runtime.register_resource_module("aws", "opsworks/userProfile", _module_instance)

_register_module()
