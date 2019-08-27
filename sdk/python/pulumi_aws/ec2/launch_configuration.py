# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LaunchConfiguration(pulumi.CustomResource):
    associate_public_ip_address: pulumi.Output[bool]
    """
    Associate a public ip address with an instance in a VPC.
    """
    ebs_block_devices: pulumi.Output[list]
    """
    Additional EBS block devices to attach to the
    instance.  See Block Devices below for details.
    
      * `delete_on_termination` (`bool`)
      * `device_name` (`str`)
      * `encrypted` (`bool`)
      * `iops` (`float`)
      * `no_device` (`bool`)
      * `snapshot_id` (`str`)
      * `volume_size` (`float`)
      * `volume_type` (`str`)
    """
    ebs_optimized: pulumi.Output[bool]
    """
    If true, the launched EC2 instance will be EBS-optimized.
    """
    enable_monitoring: pulumi.Output[bool]
    """
    Enables/disables detailed monitoring. This is enabled by default.
    """
    ephemeral_block_devices: pulumi.Output[list]
    """
    Customize Ephemeral (also known as
    "Instance Store") volumes on the instance. See Block Devices below for details.
    
      * `device_name` (`str`)
      * `virtual_name` (`str`)
    """
    iam_instance_profile: pulumi.Output[str]
    """
    The name attribute of the IAM instance profile to associate
    with launched instances.
    """
    image_id: pulumi.Output[str]
    """
    The EC2 image ID to launch.
    """
    instance_type: pulumi.Output[str]
    """
    The size of instance to launch.
    """
    key_name: pulumi.Output[str]
    """
    The key name that should be used for the instance.
    """
    name: pulumi.Output[str]
    """
    The name of the launch configuration. If you leave
    this blank, this provider will auto-generate a unique name.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified
    prefix. Conflicts with `name`.
    """
    placement_tenancy: pulumi.Output[str]
    """
    The tenancy of the instance. Valid values are
    `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
    for more details
    """
    root_block_device: pulumi.Output[dict]
    """
    Customize details about the root block
    device of the instance. See Block Devices below for details.
    
      * `delete_on_termination` (`bool`)
      * `encrypted` (`bool`)
      * `iops` (`float`)
      * `volume_size` (`float`)
      * `volume_type` (`str`)
    """
    security_groups: pulumi.Output[list]
    """
    A list of associated security group IDS.
    """
    spot_price: pulumi.Output[str]
    """
    The maximum price to use for reserving spot instances.
    """
    user_data: pulumi.Output[str]
    """
    The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
    """
    user_data_base64: pulumi.Output[str]
    """
    Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
    """
    vpc_classic_link_id: pulumi.Output[str]
    """
    The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
    """
    vpc_classic_link_security_groups: pulumi.Output[list]
    """
    The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
    """
    def __init__(__self__, resource_name, opts=None, associate_public_ip_address=None, ebs_block_devices=None, ebs_optimized=None, enable_monitoring=None, ephemeral_block_devices=None, iam_instance_profile=None, image_id=None, instance_type=None, key_name=None, name=None, name_prefix=None, placement_tenancy=None, root_block_device=None, security_groups=None, spot_price=None, user_data=None, user_data_base64=None, vpc_classic_link_id=None, vpc_classic_link_security_groups=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a new launch configuration, used for autoscaling groups.
        
        ## Block devices
        
        Each of the `*_block_device` attributes controls a portion of the AWS
        Launch Configuration's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
        Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
        to understand the implications of using these attributes.
        
        The `root_block_device` mapping supports the following:
        
        * `volume_type` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
          or `"io1"`. (Default: `"standard"`).
        * `volume_size` - (Optional) The size of the volume in gigabytes.
        * `iops` - (Optional) The amount of provisioned
          [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
          This must be set with a `volume_type` of `"io1"`.
        * `delete_on_termination` - (Optional) Whether the volume should be destroyed
          on instance termination (Default: `true`).
        * `encrypted` - (Optional) Whether the volume should be encrypted or not. (Default: `false`).
        
        Modifying any of the `root_block_device` settings requires resource
        replacement.
        
        Each `ebs_block_device` supports the following:
        
        * `device_name` - (Required) The name of the device to mount.
        * `snapshot_id` - (Optional) The Snapshot ID to mount.
        * `volume_type` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
          or `"io1"`. (Default: `"standard"`).
        * `volume_size` - (Optional) The size of the volume in gigabytes.
        * `iops` - (Optional) The amount of provisioned
          [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
          This must be set with a `volume_type` of `"io1"`.
        * `delete_on_termination` - (Optional) Whether the volume should be destroyed
          on instance termination (Default: `true`).
        * `encrypted` - (Optional) Whether the volume should be encrypted or not. Do not use this option if you are using `snapshot_id` as the encrypted flag will be determined by the snapshot. (Default: `false`).
        
        Modifying any `ebs_block_device` currently requires resource replacement.
        
        Each `ephemeral_block_device` supports the following:
        
        * `device_name` - The name of the block device to mount on the instance.
        * `virtual_name` - The [Instance Store Device
          Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
          (e.g. `"ephemeral0"`)
        
        Each AWS Instance type has a different set of Instance Store block devices
        available for attachment. AWS [publishes a
        list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
        of which ephemeral devices are available on each type. The devices are always
        identified by the `virtual_name` in the format `"ephemeral{0..N}"`.
        
        > **NOTE:** Changes to `*_block_device` configuration of _existing_ resources
        cannot currently be detected by this provider. After updating to block device
        configuration, resource recreation can be manually triggered by using the
        [`taint` command](https://www.terraform.io/docs/commands/taint.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] associate_public_ip_address: Associate a public ip address with an instance in a VPC.
        :param pulumi.Input[list] ebs_block_devices: Additional EBS block devices to attach to the
               instance.  See Block Devices below for details.
        :param pulumi.Input[bool] ebs_optimized: If true, the launched EC2 instance will be EBS-optimized.
        :param pulumi.Input[bool] enable_monitoring: Enables/disables detailed monitoring. This is enabled by default.
        :param pulumi.Input[list] ephemeral_block_devices: Customize Ephemeral (also known as
               "Instance Store") volumes on the instance. See Block Devices below for details.
        :param pulumi.Input[str] iam_instance_profile: The name attribute of the IAM instance profile to associate
               with launched instances.
        :param pulumi.Input[str] image_id: The EC2 image ID to launch.
        :param pulumi.Input[str] instance_type: The size of instance to launch.
        :param pulumi.Input[str] key_name: The key name that should be used for the instance.
        :param pulumi.Input[str] name: The name of the launch configuration. If you leave
               this blank, this provider will auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[str] placement_tenancy: The tenancy of the instance. Valid values are
               `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
               for more details
        :param pulumi.Input[dict] root_block_device: Customize details about the root block
               device of the instance. See Block Devices below for details.
        :param pulumi.Input[list] security_groups: A list of associated security group IDS.
        :param pulumi.Input[str] spot_price: The maximum price to use for reserving spot instances.
        :param pulumi.Input[str] user_data: The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
        :param pulumi.Input[str] user_data_base64: Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
        :param pulumi.Input[str] vpc_classic_link_id: The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
        :param pulumi.Input[list] vpc_classic_link_security_groups: The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
        
        The **ebs_block_devices** object supports the following:
        
          * `delete_on_termination` (`pulumi.Input[bool]`)
          * `device_name` (`pulumi.Input[str]`)
          * `encrypted` (`pulumi.Input[bool]`)
          * `iops` (`pulumi.Input[float]`)
          * `no_device` (`pulumi.Input[bool]`)
          * `snapshot_id` (`pulumi.Input[str]`)
          * `volume_size` (`pulumi.Input[float]`)
          * `volume_type` (`pulumi.Input[str]`)
        
        The **ephemeral_block_devices** object supports the following:
        
          * `device_name` (`pulumi.Input[str]`)
          * `virtual_name` (`pulumi.Input[str]`)
        
        The **root_block_device** object supports the following:
        
          * `delete_on_termination` (`pulumi.Input[bool]`)
          * `encrypted` (`pulumi.Input[bool]`)
          * `iops` (`pulumi.Input[float]`)
          * `volume_size` (`pulumi.Input[float]`)
          * `volume_type` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/launch_configuration.html.markdown.
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

            __props__['associate_public_ip_address'] = associate_public_ip_address
            __props__['ebs_block_devices'] = ebs_block_devices
            __props__['ebs_optimized'] = ebs_optimized
            __props__['enable_monitoring'] = enable_monitoring
            __props__['ephemeral_block_devices'] = ephemeral_block_devices
            __props__['iam_instance_profile'] = iam_instance_profile
            if image_id is None:
                raise TypeError("Missing required property 'image_id'")
            __props__['image_id'] = image_id
            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['key_name'] = key_name
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['placement_tenancy'] = placement_tenancy
            __props__['root_block_device'] = root_block_device
            __props__['security_groups'] = security_groups
            __props__['spot_price'] = spot_price
            __props__['user_data'] = user_data
            __props__['user_data_base64'] = user_data_base64
            __props__['vpc_classic_link_id'] = vpc_classic_link_id
            __props__['vpc_classic_link_security_groups'] = vpc_classic_link_security_groups
        super(LaunchConfiguration, __self__).__init__(
            'aws:ec2/launchConfiguration:LaunchConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, associate_public_ip_address=None, ebs_block_devices=None, ebs_optimized=None, enable_monitoring=None, ephemeral_block_devices=None, iam_instance_profile=None, image_id=None, instance_type=None, key_name=None, name=None, name_prefix=None, placement_tenancy=None, root_block_device=None, security_groups=None, spot_price=None, user_data=None, user_data_base64=None, vpc_classic_link_id=None, vpc_classic_link_security_groups=None):
        """
        Get an existing LaunchConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] associate_public_ip_address: Associate a public ip address with an instance in a VPC.
        :param pulumi.Input[list] ebs_block_devices: Additional EBS block devices to attach to the
               instance.  See Block Devices below for details.
        :param pulumi.Input[bool] ebs_optimized: If true, the launched EC2 instance will be EBS-optimized.
        :param pulumi.Input[bool] enable_monitoring: Enables/disables detailed monitoring. This is enabled by default.
        :param pulumi.Input[list] ephemeral_block_devices: Customize Ephemeral (also known as
               "Instance Store") volumes on the instance. See Block Devices below for details.
        :param pulumi.Input[str] iam_instance_profile: The name attribute of the IAM instance profile to associate
               with launched instances.
        :param pulumi.Input[str] image_id: The EC2 image ID to launch.
        :param pulumi.Input[str] instance_type: The size of instance to launch.
        :param pulumi.Input[str] key_name: The key name that should be used for the instance.
        :param pulumi.Input[str] name: The name of the launch configuration. If you leave
               this blank, this provider will auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[str] placement_tenancy: The tenancy of the instance. Valid values are
               `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
               for more details
        :param pulumi.Input[dict] root_block_device: Customize details about the root block
               device of the instance. See Block Devices below for details.
        :param pulumi.Input[list] security_groups: A list of associated security group IDS.
        :param pulumi.Input[str] spot_price: The maximum price to use for reserving spot instances.
        :param pulumi.Input[str] user_data: The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
        :param pulumi.Input[str] user_data_base64: Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
        :param pulumi.Input[str] vpc_classic_link_id: The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
        :param pulumi.Input[list] vpc_classic_link_security_groups: The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
        
        The **ebs_block_devices** object supports the following:
        
          * `delete_on_termination` (`pulumi.Input[bool]`)
          * `device_name` (`pulumi.Input[str]`)
          * `encrypted` (`pulumi.Input[bool]`)
          * `iops` (`pulumi.Input[float]`)
          * `no_device` (`pulumi.Input[bool]`)
          * `snapshot_id` (`pulumi.Input[str]`)
          * `volume_size` (`pulumi.Input[float]`)
          * `volume_type` (`pulumi.Input[str]`)
        
        The **ephemeral_block_devices** object supports the following:
        
          * `device_name` (`pulumi.Input[str]`)
          * `virtual_name` (`pulumi.Input[str]`)
        
        The **root_block_device** object supports the following:
        
          * `delete_on_termination` (`pulumi.Input[bool]`)
          * `encrypted` (`pulumi.Input[bool]`)
          * `iops` (`pulumi.Input[float]`)
          * `volume_size` (`pulumi.Input[float]`)
          * `volume_type` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/launch_configuration.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["associate_public_ip_address"] = associate_public_ip_address
        __props__["ebs_block_devices"] = ebs_block_devices
        __props__["ebs_optimized"] = ebs_optimized
        __props__["enable_monitoring"] = enable_monitoring
        __props__["ephemeral_block_devices"] = ephemeral_block_devices
        __props__["iam_instance_profile"] = iam_instance_profile
        __props__["image_id"] = image_id
        __props__["instance_type"] = instance_type
        __props__["key_name"] = key_name
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["placement_tenancy"] = placement_tenancy
        __props__["root_block_device"] = root_block_device
        __props__["security_groups"] = security_groups
        __props__["spot_price"] = spot_price
        __props__["user_data"] = user_data
        __props__["user_data_base64"] = user_data_base64
        __props__["vpc_classic_link_id"] = vpc_classic_link_id
        __props__["vpc_classic_link_security_groups"] = vpc_classic_link_security_groups
        return LaunchConfiguration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

