# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Ami(pulumi.CustomResource):
    architecture: pulumi.Output[str]
    """
    Machine architecture for created instances. Defaults to "x86_64".
    """
    arn: pulumi.Output[str]
    """
    The ARN of the AMI.
    """
    description: pulumi.Output[str]
    """
    A longer, human-readable description for the AMI.
    """
    ebs_block_devices: pulumi.Output[list]
    """
    Nested block describing an EBS block device that should be
    attached to created instances. The structure of this block is described below.

      * `deleteOnTermination` (`bool`) - Boolean controlling whether the EBS volumes created to
        support each created instance will be deleted once that instance is terminated.
      * `device_name` (`str`) - The path at which the device is exposed to created instances.
      * `encrypted` (`bool`) - Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
      * `iops` (`float`) - Number of I/O operations per second the
        created volumes will support.
      * `snapshot_id` (`str`) - The id of an EBS snapshot that will be used to initialize the created
        EBS volumes. If set, the `volume_size` attribute must be at least as large as the referenced
        snapshot.
      * `volume_size` (`float`) - The size of created volumes in GiB.
        If `snapshot_id` is set and `volume_size` is omitted then the volume will have the same size
        as the selected snapshot.
      * `volumeType` (`str`) - The type of EBS volume to create. Can be one of "standard" (the
        default), "io1" or "gp2".
    """
    ena_support: pulumi.Output[bool]
    """
    Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
    """
    ephemeral_block_devices: pulumi.Output[list]
    """
    Nested block describing an ephemeral block device that
    should be attached to created instances. The structure of this block is described below.

      * `device_name` (`str`) - The path at which the device is exposed to created instances.
      * `virtualName` (`str`) - A name for the ephemeral device, of the form "ephemeralN" where
        *N* is a volume number starting from zero.
    """
    image_location: pulumi.Output[str]
    """
    Path to an S3 object containing an image manifest, e.g. created
    by the `ec2-upload-bundle` command in the EC2 command line tools.
    """
    kernel_id: pulumi.Output[str]
    """
    The id of the kernel image (AKI) that will be used as the paravirtual
    kernel in created instances.
    """
    manage_ebs_snapshots: pulumi.Output[bool]
    name: pulumi.Output[str]
    """
    A region-unique name for the AMI.
    """
    ramdisk_id: pulumi.Output[str]
    """
    The id of an initrd image (ARI) that will be used when booting the
    created instances.
    """
    root_device_name: pulumi.Output[str]
    """
    The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
    """
    root_snapshot_id: pulumi.Output[str]
    """
    The Snapshot ID for the root volume (for EBS-backed AMIs)
    """
    sriov_net_support: pulumi.Output[str]
    """
    When set to "simple" (the default), enables enhanced networking
    for created instances. No other value is supported at this time.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    virtualization_type: pulumi.Output[str]
    """
    Keyword to choose what virtualization mode created instances
    will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
    changes the set of further arguments that are required, as described below.
    """
    def __init__(__self__, resource_name, opts=None, architecture=None, description=None, ebs_block_devices=None, ena_support=None, ephemeral_block_devices=None, image_location=None, kernel_id=None, name=None, ramdisk_id=None, root_device_name=None, sriov_net_support=None, tags=None, virtualization_type=None, __props__=None, __name__=None, __opts__=None):
        """
        The AMI resource allows the creation and management of a completely-custom
        *Amazon Machine Image* (AMI).

        If you just want to duplicate an existing AMI, possibly copying it to another
        region, it's better to use `ec2.AmiCopy` instead.

        If you just want to share an existing AMI with another AWS account,
        it's better to use `ec2.AmiLaunchPermission` instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create an AMI that will start a machine whose root device is backed by
        # an EBS volume populated from a snapshot. It is assumed that such a snapshot
        # already exists with the id "snap-xxxxxxxx".
        example = aws.ec2.Ami("example",
            ebs_block_devices=[{
                "device_name": "/dev/xvda",
                "snapshot_id": "snap-xxxxxxxx",
                "volume_size": 8,
            }],
            root_device_name="/dev/xvda",
            virtualization_type="hvm")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] architecture: Machine architecture for created instances. Defaults to "x86_64".
        :param pulumi.Input[str] description: A longer, human-readable description for the AMI.
        :param pulumi.Input[list] ebs_block_devices: Nested block describing an EBS block device that should be
               attached to created instances. The structure of this block is described below.
        :param pulumi.Input[bool] ena_support: Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        :param pulumi.Input[list] ephemeral_block_devices: Nested block describing an ephemeral block device that
               should be attached to created instances. The structure of this block is described below.
        :param pulumi.Input[str] image_location: Path to an S3 object containing an image manifest, e.g. created
               by the `ec2-upload-bundle` command in the EC2 command line tools.
        :param pulumi.Input[str] kernel_id: The id of the kernel image (AKI) that will be used as the paravirtual
               kernel in created instances.
        :param pulumi.Input[str] name: A region-unique name for the AMI.
        :param pulumi.Input[str] ramdisk_id: The id of an initrd image (ARI) that will be used when booting the
               created instances.
        :param pulumi.Input[str] root_device_name: The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        :param pulumi.Input[str] sriov_net_support: When set to "simple" (the default), enables enhanced networking
               for created instances. No other value is supported at this time.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] virtualization_type: Keyword to choose what virtualization mode created instances
               will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
               changes the set of further arguments that are required, as described below.

        The **ebs_block_devices** object supports the following:

          * `deleteOnTermination` (`pulumi.Input[bool]`) - Boolean controlling whether the EBS volumes created to
            support each created instance will be deleted once that instance is terminated.
          * `device_name` (`pulumi.Input[str]`) - The path at which the device is exposed to created instances.
          * `encrypted` (`pulumi.Input[bool]`) - Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
          * `iops` (`pulumi.Input[float]`) - Number of I/O operations per second the
            created volumes will support.
          * `snapshot_id` (`pulumi.Input[str]`) - The id of an EBS snapshot that will be used to initialize the created
            EBS volumes. If set, the `volume_size` attribute must be at least as large as the referenced
            snapshot.
          * `volume_size` (`pulumi.Input[float]`) - The size of created volumes in GiB.
            If `snapshot_id` is set and `volume_size` is omitted then the volume will have the same size
            as the selected snapshot.
          * `volumeType` (`pulumi.Input[str]`) - The type of EBS volume to create. Can be one of "standard" (the
            default), "io1" or "gp2".

        The **ephemeral_block_devices** object supports the following:

          * `device_name` (`pulumi.Input[str]`) - The path at which the device is exposed to created instances.
          * `virtualName` (`pulumi.Input[str]`) - A name for the ephemeral device, of the form "ephemeralN" where
            *N* is a volume number starting from zero.
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

            __props__['architecture'] = architecture
            __props__['description'] = description
            __props__['ebs_block_devices'] = ebs_block_devices
            __props__['ena_support'] = ena_support
            __props__['ephemeral_block_devices'] = ephemeral_block_devices
            __props__['image_location'] = image_location
            __props__['kernel_id'] = kernel_id
            __props__['name'] = name
            __props__['ramdisk_id'] = ramdisk_id
            __props__['root_device_name'] = root_device_name
            __props__['sriov_net_support'] = sriov_net_support
            __props__['tags'] = tags
            __props__['virtualization_type'] = virtualization_type
            __props__['arn'] = None
            __props__['manage_ebs_snapshots'] = None
            __props__['root_snapshot_id'] = None
        super(Ami, __self__).__init__(
            'aws:ec2/ami:Ami',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, architecture=None, arn=None, description=None, ebs_block_devices=None, ena_support=None, ephemeral_block_devices=None, image_location=None, kernel_id=None, manage_ebs_snapshots=None, name=None, ramdisk_id=None, root_device_name=None, root_snapshot_id=None, sriov_net_support=None, tags=None, virtualization_type=None):
        """
        Get an existing Ami resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] architecture: Machine architecture for created instances. Defaults to "x86_64".
        :param pulumi.Input[str] arn: The ARN of the AMI.
        :param pulumi.Input[str] description: A longer, human-readable description for the AMI.
        :param pulumi.Input[list] ebs_block_devices: Nested block describing an EBS block device that should be
               attached to created instances. The structure of this block is described below.
        :param pulumi.Input[bool] ena_support: Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        :param pulumi.Input[list] ephemeral_block_devices: Nested block describing an ephemeral block device that
               should be attached to created instances. The structure of this block is described below.
        :param pulumi.Input[str] image_location: Path to an S3 object containing an image manifest, e.g. created
               by the `ec2-upload-bundle` command in the EC2 command line tools.
        :param pulumi.Input[str] kernel_id: The id of the kernel image (AKI) that will be used as the paravirtual
               kernel in created instances.
        :param pulumi.Input[str] name: A region-unique name for the AMI.
        :param pulumi.Input[str] ramdisk_id: The id of an initrd image (ARI) that will be used when booting the
               created instances.
        :param pulumi.Input[str] root_device_name: The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        :param pulumi.Input[str] root_snapshot_id: The Snapshot ID for the root volume (for EBS-backed AMIs)
        :param pulumi.Input[str] sriov_net_support: When set to "simple" (the default), enables enhanced networking
               for created instances. No other value is supported at this time.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] virtualization_type: Keyword to choose what virtualization mode created instances
               will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
               changes the set of further arguments that are required, as described below.

        The **ebs_block_devices** object supports the following:

          * `deleteOnTermination` (`pulumi.Input[bool]`) - Boolean controlling whether the EBS volumes created to
            support each created instance will be deleted once that instance is terminated.
          * `device_name` (`pulumi.Input[str]`) - The path at which the device is exposed to created instances.
          * `encrypted` (`pulumi.Input[bool]`) - Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
          * `iops` (`pulumi.Input[float]`) - Number of I/O operations per second the
            created volumes will support.
          * `snapshot_id` (`pulumi.Input[str]`) - The id of an EBS snapshot that will be used to initialize the created
            EBS volumes. If set, the `volume_size` attribute must be at least as large as the referenced
            snapshot.
          * `volume_size` (`pulumi.Input[float]`) - The size of created volumes in GiB.
            If `snapshot_id` is set and `volume_size` is omitted then the volume will have the same size
            as the selected snapshot.
          * `volumeType` (`pulumi.Input[str]`) - The type of EBS volume to create. Can be one of "standard" (the
            default), "io1" or "gp2".

        The **ephemeral_block_devices** object supports the following:

          * `device_name` (`pulumi.Input[str]`) - The path at which the device is exposed to created instances.
          * `virtualName` (`pulumi.Input[str]`) - A name for the ephemeral device, of the form "ephemeralN" where
            *N* is a volume number starting from zero.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["architecture"] = architecture
        __props__["arn"] = arn
        __props__["description"] = description
        __props__["ebs_block_devices"] = ebs_block_devices
        __props__["ena_support"] = ena_support
        __props__["ephemeral_block_devices"] = ephemeral_block_devices
        __props__["image_location"] = image_location
        __props__["kernel_id"] = kernel_id
        __props__["manage_ebs_snapshots"] = manage_ebs_snapshots
        __props__["name"] = name
        __props__["ramdisk_id"] = ramdisk_id
        __props__["root_device_name"] = root_device_name
        __props__["root_snapshot_id"] = root_snapshot_id
        __props__["sriov_net_support"] = sriov_net_support
        __props__["tags"] = tags
        __props__["virtualization_type"] = virtualization_type
        return Ami(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
