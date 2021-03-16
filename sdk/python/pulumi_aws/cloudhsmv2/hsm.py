# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['HsmArgs', 'Hsm']

@pulumi.input_type
class HsmArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Hsm resource.
        :param pulumi.Input[str] cluster_id: The ID of Cloud HSM v2 cluster to which HSM will be added.
        :param pulumi.Input[str] availability_zone: The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
        :param pulumi.Input[str] ip_address: The IP address of HSM module. Must be within the CIDR of selected subnet.
        :param pulumi.Input[str] subnet_id: The ID of subnet in which HSM module will be located.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ID of Cloud HSM v2 cluster to which HSM will be added.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of HSM module. Must be within the CIDR of selected subnet.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of subnet in which HSM module will be located.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class Hsm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an HSM module in Amazon CloudHSM v2 cluster.

        ## Example Usage

        The following example below creates an HSM module in CloudHSM cluster.

        ```python
        import pulumi
        import pulumi_aws as aws

        cluster = aws.cloudhsmv2.get_cluster(cluster_id=var["cloudhsm_cluster_id"])
        cloudhsm_v2_hsm = aws.cloudhsmv2.Hsm("cloudhsmV2Hsm",
            subnet_id=cluster.subnet_ids[0],
            cluster_id=cluster.cluster_id)
        ```

        ## Import

        HSM modules can be imported using their HSM ID, e.g.

        ```sh
         $ pulumi import aws:cloudhsmv2/hsm:Hsm bar hsm-quo8dahtaca
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
        :param pulumi.Input[str] cluster_id: The ID of Cloud HSM v2 cluster to which HSM will be added.
        :param pulumi.Input[str] ip_address: The IP address of HSM module. Must be within the CIDR of selected subnet.
        :param pulumi.Input[str] subnet_id: The ID of subnet in which HSM module will be located.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HsmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an HSM module in Amazon CloudHSM v2 cluster.

        ## Example Usage

        The following example below creates an HSM module in CloudHSM cluster.

        ```python
        import pulumi
        import pulumi_aws as aws

        cluster = aws.cloudhsmv2.get_cluster(cluster_id=var["cloudhsm_cluster_id"])
        cloudhsm_v2_hsm = aws.cloudhsmv2.Hsm("cloudhsmV2Hsm",
            subnet_id=cluster.subnet_ids[0],
            cluster_id=cluster.cluster_id)
        ```

        ## Import

        HSM modules can be imported using their HSM ID, e.g.

        ```sh
         $ pulumi import aws:cloudhsmv2/hsm:Hsm bar hsm-quo8dahtaca
        ```

        :param str resource_name: The name of the resource.
        :param HsmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HsmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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

            __props__['availability_zone'] = availability_zone
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['ip_address'] = ip_address
            __props__['subnet_id'] = subnet_id
            __props__['hsm_eni_id'] = None
            __props__['hsm_id'] = None
            __props__['hsm_state'] = None
        super(Hsm, __self__).__init__(
            'aws:cloudhsmv2/hsm:Hsm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            hsm_eni_id: Optional[pulumi.Input[str]] = None,
            hsm_id: Optional[pulumi.Input[str]] = None,
            hsm_state: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'Hsm':
        """
        Get an existing Hsm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
        :param pulumi.Input[str] cluster_id: The ID of Cloud HSM v2 cluster to which HSM will be added.
        :param pulumi.Input[str] hsm_eni_id: The id of the ENI interface allocated for HSM module.
        :param pulumi.Input[str] hsm_id: The id of the HSM module.
        :param pulumi.Input[str] hsm_state: The state of the HSM module.
        :param pulumi.Input[str] ip_address: The IP address of HSM module. Must be within the CIDR of selected subnet.
        :param pulumi.Input[str] subnet_id: The ID of subnet in which HSM module will be located.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["availability_zone"] = availability_zone
        __props__["cluster_id"] = cluster_id
        __props__["hsm_eni_id"] = hsm_eni_id
        __props__["hsm_id"] = hsm_id
        __props__["hsm_state"] = hsm_state
        __props__["ip_address"] = ip_address
        __props__["subnet_id"] = subnet_id
        return Hsm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of Cloud HSM v2 cluster to which HSM will be added.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="hsmEniId")
    def hsm_eni_id(self) -> pulumi.Output[str]:
        """
        The id of the ENI interface allocated for HSM module.
        """
        return pulumi.get(self, "hsm_eni_id")

    @property
    @pulumi.getter(name="hsmId")
    def hsm_id(self) -> pulumi.Output[str]:
        """
        The id of the HSM module.
        """
        return pulumi.get(self, "hsm_id")

    @property
    @pulumi.getter(name="hsmState")
    def hsm_state(self) -> pulumi.Output[str]:
        """
        The state of the HSM module.
        """
        return pulumi.get(self, "hsm_state")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP address of HSM module. Must be within the CIDR of selected subnet.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of subnet in which HSM module will be located.
        """
        return pulumi.get(self, "subnet_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

