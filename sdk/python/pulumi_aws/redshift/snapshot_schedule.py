# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SnapshotScheduleArgs', 'SnapshotSchedule']

@pulumi.input_type
class SnapshotScheduleArgs:
    def __init__(__self__, *,
                 definitions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SnapshotSchedule resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        :param pulumi.Input[str] description: The description of the snapshot schedule.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        :param pulumi.Input[str] identifier: The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        pulumi.set(__self__, "definitions", definitions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if identifier is not None:
            pulumi.set(__self__, "identifier", identifier)
        if identifier_prefix is not None:
            pulumi.set(__self__, "identifier_prefix", identifier_prefix)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        """
        return pulumi.get(self, "definitions")

    @definitions.setter
    def definitions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "definitions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the snapshot schedule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique
        identifier beginning with the specified prefix. Conflicts with `identifier`.
        """
        return pulumi.get(self, "identifier_prefix")

    @identifier_prefix.setter
    def identifier_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier_prefix", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SnapshotSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.redshift.SnapshotSchedule("default",
            definitions=["rate(12 hours)"],
            identifier="tf-redshift-snapshot-schedule")
        ```

        ## Import

        Redshift Snapshot Schedule can be imported using the `identifier`, e.g.

        ```sh
         $ pulumi import aws:redshift/snapshotSchedule:SnapshotSchedule default tf-redshift-snapshot-schedule
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        :param pulumi.Input[str] description: The description of the snapshot schedule.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        :param pulumi.Input[str] identifier: The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.redshift.SnapshotSchedule("default",
            definitions=["rate(12 hours)"],
            identifier="tf-redshift-snapshot-schedule")
        ```

        ## Import

        Redshift Snapshot Schedule can be imported using the `identifier`, e.g.

        ```sh
         $ pulumi import aws:redshift/snapshotSchedule:SnapshotSchedule default tf-redshift-snapshot-schedule
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if definitions is None and not opts.urn:
                raise TypeError("Missing required property 'definitions'")
            __props__['definitions'] = definitions
            __props__['description'] = description
            __props__['force_destroy'] = force_destroy
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            __props__['tags'] = tags
            __props__['arn'] = None
        super(SnapshotSchedule, __self__).__init__(
            'aws:redshift/snapshotSchedule:SnapshotSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            identifier_prefix: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SnapshotSchedule':
        """
        Get an existing SnapshotSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        :param pulumi.Input[str] description: The description of the snapshot schedule.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        :param pulumi.Input[str] identifier: The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["definitions"] = definitions
        __props__["description"] = description
        __props__["force_destroy"] = force_destroy
        __props__["identifier"] = identifier
        __props__["identifier_prefix"] = identifier_prefix
        __props__["tags"] = tags
        return SnapshotSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Output[Sequence[str]]:
        """
        The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        """
        return pulumi.get(self, "definitions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the snapshot schedule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique
        identifier beginning with the specified prefix. Conflicts with `identifier`.
        """
        return pulumi.get(self, "identifier_prefix")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

