# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MaintenanceWindowTargetArgs', 'MaintenanceWindowTarget']

@pulumi.input_type
class MaintenanceWindowTargetArgs:
    def __init__(__self__, *,
                 resource_type: pulumi.Input[str],
                 targets: pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]],
                 window_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MaintenanceWindowTarget resource.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        """
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "targets", targets)
        pulumi.set(__self__, "window_id", window_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_information is not None:
            pulumi.set(__self__, "owner_information", owner_information)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]]:
        """
        The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
        (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> pulumi.Input[str]:
        """
        The Id of the maintenance window to register the target with.
        """
        return pulumi.get(self, "window_id")

    @window_id.setter
    def window_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "window_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the maintenance window target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the maintenance window target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        """
        return pulumi.get(self, "owner_information")

    @owner_information.setter
    def owner_information(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_information", value)


@pulumi.input_type
class _MaintenanceWindowTargetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MaintenanceWindowTarget resources.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_information is not None:
            pulumi.set(__self__, "owner_information", owner_information)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if window_id is not None:
            pulumi.set(__self__, "window_id", window_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the maintenance window target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the maintenance window target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        """
        return pulumi.get(self, "owner_information")

    @owner_information.setter
    def owner_information(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_information", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]]]:
        """
        The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
        (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of the maintenance window to register the target with.
        """
        return pulumi.get(self, "window_id")

    @window_id.setter
    def window_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "window_id", value)


class MaintenanceWindowTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an SSM Maintenance Window Target resource

        ## Example Usage
        ### Instance Target

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
            targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
                key="tag:Name",
                values=["acceptance_test"],
            )])
        ```
        ### Resource Group Target

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
            targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
                key="resource-groups:ResourceTypeFilters",
                values=["AWS::EC2::Instance"],
            )])
        ```

        ## Import

        SSM Maintenance Window targets can be imported using `WINDOW_ID/WINDOW_TARGET_ID`, e.g.

        ```sh
         $ pulumi import aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: MaintenanceWindowTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an SSM Maintenance Window Target resource

        ## Example Usage
        ### Instance Target

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
            targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
                key="tag:Name",
                values=["acceptance_test"],
            )])
        ```
        ### Resource Group Target

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
            targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
                key="resource-groups:ResourceTypeFilters",
                values=["AWS::EC2::Instance"],
            )])
        ```

        ## Import

        SSM Maintenance Window targets can be imported using `WINDOW_ID/WINDOW_TARGET_ID`, e.g.

        ```sh
         $ pulumi import aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
        ```

        :param str resource_name_: The name of the resource.
        :param MaintenanceWindowTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MaintenanceWindowTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MaintenanceWindowTargetArgs.__new__(MaintenanceWindowTargetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_information"] = owner_information
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            if targets is None and not opts.urn:
                raise TypeError("Missing required property 'targets'")
            __props__.__dict__["targets"] = targets
            if window_id is None and not opts.urn:
                raise TypeError("Missing required property 'window_id'")
            __props__.__dict__["window_id"] = window_id
        super(MaintenanceWindowTarget, __self__).__init__(
            'aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_information: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
            window_id: Optional[pulumi.Input[str]] = None) -> 'MaintenanceWindowTarget':
        """
        Get an existing MaintenanceWindowTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window target.
        :param pulumi.Input[str] name: The name of the maintenance window target.
        :param pulumi.Input[str] owner_information: User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        :param pulumi.Input[str] resource_type: The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]] targets: The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
               (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the target with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MaintenanceWindowTargetState.__new__(_MaintenanceWindowTargetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["owner_information"] = owner_information
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["targets"] = targets
        __props__.__dict__["window_id"] = window_id
        return MaintenanceWindowTarget(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the maintenance window target.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the maintenance window target.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> pulumi.Output[Optional[str]]:
        """
        User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        """
        return pulumi.get(self, "owner_information")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Sequence['outputs.MaintenanceWindowTargetTarget']]:
        """
        The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
        (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> pulumi.Output[str]:
        """
        The Id of the maintenance window to register the target with.
        """
        return pulumi.get(self, "window_id")

