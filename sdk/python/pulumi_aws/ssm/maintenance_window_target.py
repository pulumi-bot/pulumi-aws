# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MaintenanceWindowTarget(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the maintenance window target.
    """
    name: pulumi.Output[str]
    """
    The name of the maintenance window target.
    """
    owner_information: pulumi.Output[str]
    """
    User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
    """
    resource_type: pulumi.Output[str]
    """
    The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
    """
    targets: pulumi.Output[list]
    """
    The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
    """
    window_id: pulumi.Output[str]
    """
    The Id of the maintenance window to register the target with.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, owner_information=None, resource_type=None, targets=None, window_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an SSM Maintenance Window Target resource
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
        :param pulumi.Input[list] targets: The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_maintenance_window_target.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['name'] = name
            __props__['owner_information'] = owner_information
            if resource_type is None:
                raise TypeError("Missing required property 'resource_type'")
            __props__['resource_type'] = resource_type
            if targets is None:
                raise TypeError("Missing required property 'targets'")
            __props__['targets'] = targets
            if window_id is None:
                raise TypeError("Missing required property 'window_id'")
            __props__['window_id'] = window_id
        super(MaintenanceWindowTarget, __self__).__init__(
            'aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, name=None, owner_information=None, resource_type=None, targets=None, window_id=None):
        """
        Get an existing MaintenanceWindowTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
        :param pulumi.Input[list] targets: The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_maintenance_window_target.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["name"] = name
        __props__["owner_information"] = owner_information
        __props__["resource_type"] = resource_type
        __props__["targets"] = targets
        __props__["window_id"] = window_id
        return MaintenanceWindowTarget(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

