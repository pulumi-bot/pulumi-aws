# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    agent_version: pulumi.Output[str]
    """
    The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
    """
    ami_id: pulumi.Output[str]
    """
    The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
    """
    architecture: pulumi.Output[str]
    """
    Machine architecture for created instances.  Can be either `"x86_64"` (the default) or `"i386"`
    """
    auto_scaling_type: pulumi.Output[str]
    """
    Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
    """
    availability_zone: pulumi.Output[str]
    """
    Name of the availability zone where instances will be created
    by default.
    """
    created_at: pulumi.Output[str]
    delete_ebs: pulumi.Output[bool]
    delete_eip: pulumi.Output[bool]
    ebs_block_devices: pulumi.Output[list]
    """
    Additional EBS block devices to attach to the
    instance.  See Block Devices below for details.
    """
    ebs_optimized: pulumi.Output[bool]
    """
    If true, the launched EC2 instance will be EBS-optimized.
    """
    ec2_instance_id: pulumi.Output[str]
    """
    EC2 instance ID
    """
    ecs_cluster_arn: pulumi.Output[str]
    elastic_ip: pulumi.Output[str]
    ephemeral_block_devices: pulumi.Output[list]
    """
    Customize Ephemeral (also known as
    "Instance Store") volumes on the instance. See Block Devices below for details.
    """
    hostname: pulumi.Output[str]
    """
    The instance's host name.
    """
    infrastructure_class: pulumi.Output[str]
    install_updates_on_boot: pulumi.Output[bool]
    """
    Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
    """
    instance_profile_arn: pulumi.Output[str]
    instance_type: pulumi.Output[str]
    """
    The type of instance to start
    """
    last_service_error_id: pulumi.Output[str]
    layer_ids: pulumi.Output[list]
    """
    The ids of the layers the instance will belong to.
    """
    os: pulumi.Output[str]
    """
    Name of operating system that will be installed.
    """
    platform: pulumi.Output[str]
    private_dns: pulumi.Output[str]
    """
    The private DNS name assigned to the instance. Can only be
    used inside the Amazon EC2, and only available if you've enabled DNS hostnames
    for your VPC
    """
    private_ip: pulumi.Output[str]
    """
    The private IP address assigned to the instance
    """
    public_dns: pulumi.Output[str]
    """
    The public DNS name assigned to the instance. For EC2-VPC, this
    is only available if you've enabled DNS hostnames for your VPC
    """
    public_ip: pulumi.Output[str]
    """
    The public IP address assigned to the instance, if applicable.
    """
    registered_by: pulumi.Output[str]
    reported_agent_version: pulumi.Output[str]
    reported_os_family: pulumi.Output[str]
    reported_os_name: pulumi.Output[str]
    reported_os_version: pulumi.Output[str]
    root_block_devices: pulumi.Output[list]
    """
    Customize details about the root block
    device of the instance. See Block Devices below for details.
    """
    root_device_type: pulumi.Output[str]
    """
    Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
    """
    root_device_volume_id: pulumi.Output[str]
    security_group_ids: pulumi.Output[list]
    """
    The associated security groups.
    """
    ssh_host_dsa_key_fingerprint: pulumi.Output[str]
    ssh_host_rsa_key_fingerprint: pulumi.Output[str]
    ssh_key_name: pulumi.Output[str]
    """
    Name of the SSH keypair that instances will have by default.
    """
    stack_id: pulumi.Output[str]
    """
    The id of the stack the instance will belong to.
    """
    state: pulumi.Output[str]
    """
    The desired state of the instance.  Can be either `"running"` or `"stopped"`.
    """
    status: pulumi.Output[str]
    subnet_id: pulumi.Output[str]
    """
    Subnet ID to attach to
    """
    tenancy: pulumi.Output[str]
    """
    Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
    """
    virtualization_type: pulumi.Output[str]
    """
    Keyword to choose what virtualization mode created instances
    will use. Can be either `"paravirtual"` or `"hvm"`.
    """
    def __init__(__self__, resource_name, opts=None, agent_version=None, ami_id=None, architecture=None, auto_scaling_type=None, availability_zone=None, created_at=None, delete_ebs=None, delete_eip=None, ebs_block_devices=None, ebs_optimized=None, ecs_cluster_arn=None, elastic_ip=None, ephemeral_block_devices=None, hostname=None, infrastructure_class=None, install_updates_on_boot=None, instance_profile_arn=None, instance_type=None, last_service_error_id=None, layer_ids=None, os=None, platform=None, private_dns=None, private_ip=None, public_dns=None, public_ip=None, registered_by=None, reported_agent_version=None, reported_os_family=None, reported_os_name=None, reported_os_version=None, root_block_devices=None, root_device_type=None, root_device_volume_id=None, security_group_ids=None, ssh_host_dsa_key_fingerprint=None, ssh_host_rsa_key_fingerprint=None, ssh_key_name=None, stack_id=None, state=None, status=None, subnet_id=None, tenancy=None, virtualization_type=None, __name__=None, __opts__=None):
        """
        Provides an OpsWorks instance resource.
        
        ## Block devices
        
        Each of the `*_block_device` attributes controls a portion of the AWS
        Instance's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
        Mapping docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
        to understand the implications of using these attributes.
        
        The `root_block_device` mapping supports the following:
        
        * `volume_type` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
          or `"io1"`. (Default: `"standard"`).
        * `volume_size` - (Optional) The size of the volume in gigabytes.
        * `iops` - (Optional) The amount of provisioned
          [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
          This must be set with a `volume_type` of `"io1"`.
        * `delete_on_termination` - (Optional) Whether the volume should be destroyed
          on instance termination (Default: `true`).
        
        Modifying any of the `root_block_device` settings requires resource
        replacement.
        
        Each `ebs_block_device` supports the following:
        
        * `device_name` - The name of the device to mount.
        * `snapshot_id` - (Optional) The Snapshot ID to mount.
        * `volume_type` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
          or `"io1"`. (Default: `"standard"`).
        * `volume_size` - (Optional) The size of the volume in gigabytes.
        * `iops` - (Optional) The amount of provisioned
          [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
          This must be set with a `volume_type` of `"io1"`.
        * `delete_on_termination` - (Optional) Whether the volume should be destroyed
          on instance termination (Default: `true`).
        
        Modifying any `ebs_block_device` currently requires resource replacement.
        
        Each `ephemeral_block_device` supports the following:
        
        * `device_name` - The name of the block device to mount on the instance.
        * `virtual_name` - The [Instance Store Device
          Name](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
          (e.g. `"ephemeral0"`)
        
        Each AWS Instance type has a different set of Instance Store block devices
        available for attachment. AWS [publishes a
        list](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
        of which ephemeral devices are available on each type. The devices are always
        identified by the `virtual_name` in the format `"ephemeral{0..N}"`.
        
        > **NOTE:** Currently, changes to `*_block_device` configuration of _existing_
        resources cannot be automatically detected by Terraform. After making updates
        to block device configuration, resource recreation can be manually triggered by
        using the [`taint` command](https://www.terraform.io/docs/commands/taint.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_version: The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
        :param pulumi.Input[str] ami_id: The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
        :param pulumi.Input[str] architecture: Machine architecture for created instances.  Can be either `"x86_64"` (the default) or `"i386"`
        :param pulumi.Input[str] auto_scaling_type: Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
        :param pulumi.Input[str] availability_zone: Name of the availability zone where instances will be created
               by default.
        :param pulumi.Input[list] ebs_block_devices: Additional EBS block devices to attach to the
               instance.  See Block Devices below for details.
        :param pulumi.Input[bool] ebs_optimized: If true, the launched EC2 instance will be EBS-optimized.
        :param pulumi.Input[list] ephemeral_block_devices: Customize Ephemeral (also known as
               "Instance Store") volumes on the instance. See Block Devices below for details.
        :param pulumi.Input[str] hostname: The instance's host name.
        :param pulumi.Input[bool] install_updates_on_boot: Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
        :param pulumi.Input[str] instance_type: The type of instance to start
        :param pulumi.Input[list] layer_ids: The ids of the layers the instance will belong to.
        :param pulumi.Input[str] os: Name of operating system that will be installed.
        :param pulumi.Input[str] private_dns: The private DNS name assigned to the instance. Can only be
               used inside the Amazon EC2, and only available if you've enabled DNS hostnames
               for your VPC
        :param pulumi.Input[str] private_ip: The private IP address assigned to the instance
        :param pulumi.Input[str] public_dns: The public DNS name assigned to the instance. For EC2-VPC, this
               is only available if you've enabled DNS hostnames for your VPC
        :param pulumi.Input[str] public_ip: The public IP address assigned to the instance, if applicable.
        :param pulumi.Input[list] root_block_devices: Customize details about the root block
               device of the instance. See Block Devices below for details.
        :param pulumi.Input[str] root_device_type: Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
        :param pulumi.Input[list] security_group_ids: The associated security groups.
        :param pulumi.Input[str] ssh_key_name: Name of the SSH keypair that instances will have by default.
        :param pulumi.Input[str] stack_id: The id of the stack the instance will belong to.
        :param pulumi.Input[str] state: The desired state of the instance.  Can be either `"running"` or `"stopped"`.
        :param pulumi.Input[str] subnet_id: Subnet ID to attach to
        :param pulumi.Input[str] tenancy: Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
        :param pulumi.Input[str] virtualization_type: Keyword to choose what virtualization mode created instances
               will use. Can be either `"paravirtual"` or `"hvm"`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['agent_version'] = agent_version

        __props__['ami_id'] = ami_id

        __props__['architecture'] = architecture

        __props__['auto_scaling_type'] = auto_scaling_type

        __props__['availability_zone'] = availability_zone

        __props__['created_at'] = created_at

        __props__['delete_ebs'] = delete_ebs

        __props__['delete_eip'] = delete_eip

        __props__['ebs_block_devices'] = ebs_block_devices

        __props__['ebs_optimized'] = ebs_optimized

        __props__['ecs_cluster_arn'] = ecs_cluster_arn

        __props__['elastic_ip'] = elastic_ip

        __props__['ephemeral_block_devices'] = ephemeral_block_devices

        __props__['hostname'] = hostname

        __props__['infrastructure_class'] = infrastructure_class

        __props__['install_updates_on_boot'] = install_updates_on_boot

        __props__['instance_profile_arn'] = instance_profile_arn

        __props__['instance_type'] = instance_type

        __props__['last_service_error_id'] = last_service_error_id

        if layer_ids is None:
            raise TypeError("Missing required property 'layer_ids'")
        __props__['layer_ids'] = layer_ids

        __props__['os'] = os

        __props__['platform'] = platform

        __props__['private_dns'] = private_dns

        __props__['private_ip'] = private_ip

        __props__['public_dns'] = public_dns

        __props__['public_ip'] = public_ip

        __props__['registered_by'] = registered_by

        __props__['reported_agent_version'] = reported_agent_version

        __props__['reported_os_family'] = reported_os_family

        __props__['reported_os_name'] = reported_os_name

        __props__['reported_os_version'] = reported_os_version

        __props__['root_block_devices'] = root_block_devices

        __props__['root_device_type'] = root_device_type

        __props__['root_device_volume_id'] = root_device_volume_id

        __props__['security_group_ids'] = security_group_ids

        __props__['ssh_host_dsa_key_fingerprint'] = ssh_host_dsa_key_fingerprint

        __props__['ssh_host_rsa_key_fingerprint'] = ssh_host_rsa_key_fingerprint

        __props__['ssh_key_name'] = ssh_key_name

        if stack_id is None:
            raise TypeError("Missing required property 'stack_id'")
        __props__['stack_id'] = stack_id

        __props__['state'] = state

        __props__['status'] = status

        __props__['subnet_id'] = subnet_id

        __props__['tenancy'] = tenancy

        __props__['virtualization_type'] = virtualization_type

        __props__['ec2_instance_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Instance, __self__).__init__(
            'aws:opsworks/instance:Instance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

