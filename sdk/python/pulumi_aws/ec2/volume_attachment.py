# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VolumeAttachment']


class VolumeAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 force_detach: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 skip_destroy: Optional[pulumi.Input[bool]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS EBS Volume Attachment as a top level resource, to attach and
        detach volumes from AWS Instances.

        > **NOTE on EBS block devices:** If you use `ebs_block_device` on an `ec2.Instance`, this provider will assume management over the full set of non-root EBS block devices for the instance, and treats additional block devices as drift. For this reason, `ebs_block_device` cannot be mixed with external `ebs.Volume` + `aws_ebs_volume_attachment` resources for a given instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        web = aws.ec2.Instance("web",
            ami="ami-21f78e11",
            availability_zone="us-west-2a",
            instance_type="t1.micro",
            tags={
                "Name": "HelloWorld",
            })
        example = aws.ebs.Volume("example",
            availability_zone="us-west-2a",
            size=1)
        ebs_att = aws.ec2.VolumeAttachment("ebsAtt",
            device_name="/dev/sdh",
            volume_id=example.id,
            instance_id=web.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: The device name to expose to the instance (for
               example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
        :param pulumi.Input[bool] force_detach: Set to `true` if you want to force the
               volume to detach. Useful if previous attempts failed, but use this option only
               as a last resort, as this can result in **data loss**. See
               [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to
        :param pulumi.Input[bool] skip_destroy: Set this to true if you do not wish
               to detach the volume from the instance to which it is attached at destroy
               time, and instead just remove the attachment from this provider state. This is
               useful when destroying an instance which has volumes created by some other
               means attached.
        :param pulumi.Input[str] volume_id: ID of the Volume to be attached
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

            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            __props__['force_detach'] = force_detach
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['skip_destroy'] = skip_destroy
            if volume_id is None:
                raise TypeError("Missing required property 'volume_id'")
            __props__['volume_id'] = volume_id
        super(VolumeAttachment, __self__).__init__(
            'aws:ec2/volumeAttachment:VolumeAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_name: Optional[pulumi.Input[str]] = None,
            force_detach: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            skip_destroy: Optional[pulumi.Input[bool]] = None,
            volume_id: Optional[pulumi.Input[str]] = None) -> 'VolumeAttachment':
        """
        Get an existing VolumeAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: The device name to expose to the instance (for
               example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
        :param pulumi.Input[bool] force_detach: Set to `true` if you want to force the
               volume to detach. Useful if previous attempts failed, but use this option only
               as a last resort, as this can result in **data loss**. See
               [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to
        :param pulumi.Input[bool] skip_destroy: Set this to true if you do not wish
               to detach the volume from the instance to which it is attached at destroy
               time, and instead just remove the attachment from this provider state. This is
               useful when destroying an instance which has volumes created by some other
               means attached.
        :param pulumi.Input[str] volume_id: ID of the Volume to be attached
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["device_name"] = device_name
        __props__["force_detach"] = force_detach
        __props__["instance_id"] = instance_id
        __props__["skip_destroy"] = skip_destroy
        __props__["volume_id"] = volume_id
        return VolumeAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Output[str]:
        """
        The device name to expose to the instance (for
        example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter(name="forceDetach")
    def force_detach(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` if you want to force the
        volume to detach. Useful if previous attempts failed, but use this option only
        as a last resort, as this can result in **data loss**. See
        [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
        """
        return pulumi.get(self, "force_detach")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the Instance to attach to
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="skipDestroy")
    def skip_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Set this to true if you do not wish
        to detach the volume from the instance to which it is attached at destroy
        time, and instead just remove the attachment from this provider state. This is
        useful when destroying an instance which has volumes created by some other
        means attached.
        """
        return pulumi.get(self, "skip_destroy")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        """
        ID of the Volume to be attached
        """
        return pulumi.get(self, "volume_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

