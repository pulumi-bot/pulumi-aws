# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MaintenanceWindowTarget']


class MaintenanceWindowTarget(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SSM Maintenance Window Target resource

        ## Instance Target Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        window = aws.ssm.MaintenanceWindow("window",
            schedule="cron(0 16 ? * TUE *)",
            duration=3,
            cutoff=1)
        target1 = aws.ssm.MaintenanceWindowTarget("target1",
            window_id=window.id,
            description="This is a maintenance window target",
            resource_type="INSTANCE",
            targets=[{
                "key": "tag:Name",
                "values": ["acceptance_test"],
            }])
        ```

        ## Resource Group Target Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        window = aws.ssm.MaintenanceWindow("window",
            schedule="cron(0 16 ? * TUE *)",
            duration=3,
            cutoff=1)
        target1 = aws.ssm.MaintenanceWindowTarget("target1",
            window_id=window.id,
            description="This is a maintenance window target",
            resource_type="RESOURCE_GROUP",
            targets=[{
                "key": "resource-groups:ResourceTypeFilters",
                "values": [
                    "AWS::EC2::INSTANCE",
                    "AWS::EC2::VPC",
                ],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
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
            opts.version = _utilities.get_version()
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_information: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
            window_id: Optional[pulumi.Input[str]] = None) -> 'MaintenanceWindowTarget':
        """
        Get an existing MaintenanceWindowTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["owner_information"] = owner_information
        __props__["resource_type"] = resource_type
        __props__["targets"] = targets
        __props__["window_id"] = window_id
        return MaintenanceWindowTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the maintenance window target.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the maintenance window target.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> Optional[str]:
        """
        User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        """
        return pulumi.get(self, "owner_information")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def targets(self) -> List['outputs.MaintenanceWindowTargetTarget']:
        """
        The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
        (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> str:
        """
        The Id of the maintenance window to register the target with.
        """
        return pulumi.get(self, "window_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

