# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from .composite_alarm import *
from .dashboard import *
from .event_archive import *
from .event_bus import *
from .event_permission import *
from .event_rule import *
from .event_target import *
from .get_log_group import *
from .log_destination import *
from .log_destination_policy import *
from .log_group import *
from .log_metric_filter import *
from .log_resource_policy import *
from .log_stream import *
from .log_subscription_filter import *
from .metric_alarm import *
from .query_definition import *
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
            if typ == "aws:cloudwatch/compositeAlarm:CompositeAlarm":
                return CompositeAlarm(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/dashboard:Dashboard":
                return Dashboard(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/eventArchive:EventArchive":
                return EventArchive(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/eventBus:EventBus":
                return EventBus(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/eventPermission:EventPermission":
                return EventPermission(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/eventRule:EventRule":
                return EventRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/eventTarget:EventTarget":
                return EventTarget(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logDestination:LogDestination":
                return LogDestination(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy":
                return LogDestinationPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logGroup:LogGroup":
                return LogGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logMetricFilter:LogMetricFilter":
                return LogMetricFilter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logResourcePolicy:LogResourcePolicy":
                return LogResourcePolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logStream:LogStream":
                return LogStream(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter":
                return LogSubscriptionFilter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/metricAlarm:MetricAlarm":
                return MetricAlarm(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:cloudwatch/queryDefinition:QueryDefinition":
                return QueryDefinition(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "cloudwatch/compositeAlarm", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/dashboard", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/eventArchive", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/eventBus", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/eventPermission", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/eventRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/eventTarget", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logDestination", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logDestinationPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logMetricFilter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logResourcePolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logStream", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/logSubscriptionFilter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/metricAlarm", _module_instance)
    pulumi.runtime.register_resource_module("aws", "cloudwatch/queryDefinition", _module_instance)

_register_module()
