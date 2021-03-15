# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['StaticWebLayerArgs', 'StaticWebLayer']

@pulumi.input_type
class StaticWebLayerArgs:
    def __init__(__self__, *,
                 stack_id: pulumi.Input[str],
                 auto_assign_elastic_ips: Optional[pulumi.Input[bool]] = None,
                 auto_assign_public_ips: Optional[pulumi.Input[bool]] = None,
                 auto_healing: Optional[pulumi.Input[bool]] = None,
                 custom_configure_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_deploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_instance_profile_arn: Optional[pulumi.Input[str]] = None,
                 custom_json: Optional[pulumi.Input[str]] = None,
                 custom_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_setup_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_shutdown_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_undeploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 drain_elb_on_shutdown: Optional[pulumi.Input[bool]] = None,
                 ebs_volumes: Optional[pulumi.Input[Sequence[pulumi.Input['StaticWebLayerEbsVolumeArgs']]]] = None,
                 elastic_load_balancer: Optional[pulumi.Input[str]] = None,
                 install_updates_on_boot: Optional[pulumi.Input[bool]] = None,
                 instance_shutdown_timeout: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 system_packages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 use_ebs_optimized_instances: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a StaticWebLayer resource.
        :param pulumi.Input[str] stack_id: The id of the stack the layer will belong to.
        :param pulumi.Input[bool] auto_assign_elastic_ips: Whether to automatically assign an elastic IP address to the layer's instances.
        :param pulumi.Input[bool] auto_assign_public_ips: For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        :param pulumi.Input[bool] auto_healing: Whether to enable auto-healing for the layer.
        :param pulumi.Input[str] custom_instance_profile_arn: The ARN of an IAM profile that will be used for the layer's instances.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_security_group_ids: Ids for a set of security groups to apply to the layer's instances.
        :param pulumi.Input[bool] drain_elb_on_shutdown: Whether to enable Elastic Load Balancing connection draining.
        :param pulumi.Input[Sequence[pulumi.Input['StaticWebLayerEbsVolumeArgs']]] ebs_volumes: `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        :param pulumi.Input[str] elastic_load_balancer: Name of an Elastic Load Balancer to attach to this layer
        :param pulumi.Input[bool] install_updates_on_boot: Whether to install OS and package updates on each instance when it boots.
        :param pulumi.Input[int] instance_shutdown_timeout: The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        :param pulumi.Input[str] name: A human-readable name for the layer.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] system_packages: Names of a set of system packages to install on the layer's instances.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[bool] use_ebs_optimized_instances: Whether to use EBS-optimized instances.
        """
        pulumi.set(__self__, "stack_id", stack_id)
        if auto_assign_elastic_ips is not None:
            pulumi.set(__self__, "auto_assign_elastic_ips", auto_assign_elastic_ips)
        if auto_assign_public_ips is not None:
            pulumi.set(__self__, "auto_assign_public_ips", auto_assign_public_ips)
        if auto_healing is not None:
            pulumi.set(__self__, "auto_healing", auto_healing)
        if custom_configure_recipes is not None:
            pulumi.set(__self__, "custom_configure_recipes", custom_configure_recipes)
        if custom_deploy_recipes is not None:
            pulumi.set(__self__, "custom_deploy_recipes", custom_deploy_recipes)
        if custom_instance_profile_arn is not None:
            pulumi.set(__self__, "custom_instance_profile_arn", custom_instance_profile_arn)
        if custom_json is not None:
            pulumi.set(__self__, "custom_json", custom_json)
        if custom_security_group_ids is not None:
            pulumi.set(__self__, "custom_security_group_ids", custom_security_group_ids)
        if custom_setup_recipes is not None:
            pulumi.set(__self__, "custom_setup_recipes", custom_setup_recipes)
        if custom_shutdown_recipes is not None:
            pulumi.set(__self__, "custom_shutdown_recipes", custom_shutdown_recipes)
        if custom_undeploy_recipes is not None:
            pulumi.set(__self__, "custom_undeploy_recipes", custom_undeploy_recipes)
        if drain_elb_on_shutdown is not None:
            pulumi.set(__self__, "drain_elb_on_shutdown", drain_elb_on_shutdown)
        if ebs_volumes is not None:
            pulumi.set(__self__, "ebs_volumes", ebs_volumes)
        if elastic_load_balancer is not None:
            pulumi.set(__self__, "elastic_load_balancer", elastic_load_balancer)
        if install_updates_on_boot is not None:
            pulumi.set(__self__, "install_updates_on_boot", install_updates_on_boot)
        if instance_shutdown_timeout is not None:
            pulumi.set(__self__, "instance_shutdown_timeout", instance_shutdown_timeout)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if system_packages is not None:
            pulumi.set(__self__, "system_packages", system_packages)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if use_ebs_optimized_instances is not None:
            pulumi.set(__self__, "use_ebs_optimized_instances", use_ebs_optimized_instances)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        """
        The id of the stack the layer will belong to.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="autoAssignElasticIps")
    def auto_assign_elastic_ips(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically assign an elastic IP address to the layer's instances.
        """
        return pulumi.get(self, "auto_assign_elastic_ips")

    @auto_assign_elastic_ips.setter
    def auto_assign_elastic_ips(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_assign_elastic_ips", value)

    @property
    @pulumi.getter(name="autoAssignPublicIps")
    def auto_assign_public_ips(self) -> Optional[pulumi.Input[bool]]:
        """
        For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        """
        return pulumi.get(self, "auto_assign_public_ips")

    @auto_assign_public_ips.setter
    def auto_assign_public_ips(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_assign_public_ips", value)

    @property
    @pulumi.getter(name="autoHealing")
    def auto_healing(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable auto-healing for the layer.
        """
        return pulumi.get(self, "auto_healing")

    @auto_healing.setter
    def auto_healing(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_healing", value)

    @property
    @pulumi.getter(name="customConfigureRecipes")
    def custom_configure_recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_configure_recipes")

    @custom_configure_recipes.setter
    def custom_configure_recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_configure_recipes", value)

    @property
    @pulumi.getter(name="customDeployRecipes")
    def custom_deploy_recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_deploy_recipes")

    @custom_deploy_recipes.setter
    def custom_deploy_recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_deploy_recipes", value)

    @property
    @pulumi.getter(name="customInstanceProfileArn")
    def custom_instance_profile_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an IAM profile that will be used for the layer's instances.
        """
        return pulumi.get(self, "custom_instance_profile_arn")

    @custom_instance_profile_arn.setter
    def custom_instance_profile_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_instance_profile_arn", value)

    @property
    @pulumi.getter(name="customJson")
    def custom_json(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "custom_json")

    @custom_json.setter
    def custom_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_json", value)

    @property
    @pulumi.getter(name="customSecurityGroupIds")
    def custom_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Ids for a set of security groups to apply to the layer's instances.
        """
        return pulumi.get(self, "custom_security_group_ids")

    @custom_security_group_ids.setter
    def custom_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_security_group_ids", value)

    @property
    @pulumi.getter(name="customSetupRecipes")
    def custom_setup_recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_setup_recipes")

    @custom_setup_recipes.setter
    def custom_setup_recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_setup_recipes", value)

    @property
    @pulumi.getter(name="customShutdownRecipes")
    def custom_shutdown_recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_shutdown_recipes")

    @custom_shutdown_recipes.setter
    def custom_shutdown_recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_shutdown_recipes", value)

    @property
    @pulumi.getter(name="customUndeployRecipes")
    def custom_undeploy_recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_undeploy_recipes")

    @custom_undeploy_recipes.setter
    def custom_undeploy_recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_undeploy_recipes", value)

    @property
    @pulumi.getter(name="drainElbOnShutdown")
    def drain_elb_on_shutdown(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Elastic Load Balancing connection draining.
        """
        return pulumi.get(self, "drain_elb_on_shutdown")

    @drain_elb_on_shutdown.setter
    def drain_elb_on_shutdown(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drain_elb_on_shutdown", value)

    @property
    @pulumi.getter(name="ebsVolumes")
    def ebs_volumes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StaticWebLayerEbsVolumeArgs']]]]:
        """
        `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        """
        return pulumi.get(self, "ebs_volumes")

    @ebs_volumes.setter
    def ebs_volumes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StaticWebLayerEbsVolumeArgs']]]]):
        pulumi.set(self, "ebs_volumes", value)

    @property
    @pulumi.getter(name="elasticLoadBalancer")
    def elastic_load_balancer(self) -> Optional[pulumi.Input[str]]:
        """
        Name of an Elastic Load Balancer to attach to this layer
        """
        return pulumi.get(self, "elastic_load_balancer")

    @elastic_load_balancer.setter
    def elastic_load_balancer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elastic_load_balancer", value)

    @property
    @pulumi.getter(name="installUpdatesOnBoot")
    def install_updates_on_boot(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to install OS and package updates on each instance when it boots.
        """
        return pulumi.get(self, "install_updates_on_boot")

    @install_updates_on_boot.setter
    def install_updates_on_boot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "install_updates_on_boot", value)

    @property
    @pulumi.getter(name="instanceShutdownTimeout")
    def instance_shutdown_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        """
        return pulumi.get(self, "instance_shutdown_timeout")

    @instance_shutdown_timeout.setter
    def instance_shutdown_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_shutdown_timeout", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable name for the layer.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="systemPackages")
    def system_packages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Names of a set of system packages to install on the layer's instances.
        """
        return pulumi.get(self, "system_packages")

    @system_packages.setter
    def system_packages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "system_packages", value)

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

    @property
    @pulumi.getter(name="useEbsOptimizedInstances")
    def use_ebs_optimized_instances(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use EBS-optimized instances.
        """
        return pulumi.get(self, "use_ebs_optimized_instances")

    @use_ebs_optimized_instances.setter
    def use_ebs_optimized_instances(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ebs_optimized_instances", value)


class StaticWebLayer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StaticWebLayerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an OpsWorks static web server layer resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        web = aws.opsworks.StaticWebLayer("web", stack_id=aws_opsworks_stack["main"]["id"])
        ```

        ## Import

        OpsWorks static web server Layers can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:opsworks/staticWebLayer:StaticWebLayer bar 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param StaticWebLayerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_assign_elastic_ips: Optional[pulumi.Input[bool]] = None,
                 auto_assign_public_ips: Optional[pulumi.Input[bool]] = None,
                 auto_healing: Optional[pulumi.Input[bool]] = None,
                 custom_configure_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_deploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_instance_profile_arn: Optional[pulumi.Input[str]] = None,
                 custom_json: Optional[pulumi.Input[str]] = None,
                 custom_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_setup_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_shutdown_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_undeploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 drain_elb_on_shutdown: Optional[pulumi.Input[bool]] = None,
                 ebs_volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StaticWebLayerEbsVolumeArgs']]]]] = None,
                 elastic_load_balancer: Optional[pulumi.Input[str]] = None,
                 install_updates_on_boot: Optional[pulumi.Input[bool]] = None,
                 instance_shutdown_timeout: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 system_packages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 use_ebs_optimized_instances: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an OpsWorks static web server layer resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        web = aws.opsworks.StaticWebLayer("web", stack_id=aws_opsworks_stack["main"]["id"])
        ```

        ## Import

        OpsWorks static web server Layers can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:opsworks/staticWebLayer:StaticWebLayer bar 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_assign_elastic_ips: Whether to automatically assign an elastic IP address to the layer's instances.
        :param pulumi.Input[bool] auto_assign_public_ips: For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        :param pulumi.Input[bool] auto_healing: Whether to enable auto-healing for the layer.
        :param pulumi.Input[str] custom_instance_profile_arn: The ARN of an IAM profile that will be used for the layer's instances.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_security_group_ids: Ids for a set of security groups to apply to the layer's instances.
        :param pulumi.Input[bool] drain_elb_on_shutdown: Whether to enable Elastic Load Balancing connection draining.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StaticWebLayerEbsVolumeArgs']]]] ebs_volumes: `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        :param pulumi.Input[str] elastic_load_balancer: Name of an Elastic Load Balancer to attach to this layer
        :param pulumi.Input[bool] install_updates_on_boot: Whether to install OS and package updates on each instance when it boots.
        :param pulumi.Input[int] instance_shutdown_timeout: The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        :param pulumi.Input[str] name: A human-readable name for the layer.
        :param pulumi.Input[str] stack_id: The id of the stack the layer will belong to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] system_packages: Names of a set of system packages to install on the layer's instances.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[bool] use_ebs_optimized_instances: Whether to use EBS-optimized instances.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StaticWebLayerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_assign_elastic_ips: Optional[pulumi.Input[bool]] = None,
                 auto_assign_public_ips: Optional[pulumi.Input[bool]] = None,
                 auto_healing: Optional[pulumi.Input[bool]] = None,
                 custom_configure_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_deploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_instance_profile_arn: Optional[pulumi.Input[str]] = None,
                 custom_json: Optional[pulumi.Input[str]] = None,
                 custom_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_setup_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_shutdown_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_undeploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 drain_elb_on_shutdown: Optional[pulumi.Input[bool]] = None,
                 ebs_volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StaticWebLayerEbsVolumeArgs']]]]] = None,
                 elastic_load_balancer: Optional[pulumi.Input[str]] = None,
                 install_updates_on_boot: Optional[pulumi.Input[bool]] = None,
                 instance_shutdown_timeout: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 system_packages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 use_ebs_optimized_instances: Optional[pulumi.Input[bool]] = None,
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

            __props__['auto_assign_elastic_ips'] = auto_assign_elastic_ips
            __props__['auto_assign_public_ips'] = auto_assign_public_ips
            __props__['auto_healing'] = auto_healing
            __props__['custom_configure_recipes'] = custom_configure_recipes
            __props__['custom_deploy_recipes'] = custom_deploy_recipes
            __props__['custom_instance_profile_arn'] = custom_instance_profile_arn
            __props__['custom_json'] = custom_json
            __props__['custom_security_group_ids'] = custom_security_group_ids
            __props__['custom_setup_recipes'] = custom_setup_recipes
            __props__['custom_shutdown_recipes'] = custom_shutdown_recipes
            __props__['custom_undeploy_recipes'] = custom_undeploy_recipes
            __props__['drain_elb_on_shutdown'] = drain_elb_on_shutdown
            __props__['ebs_volumes'] = ebs_volumes
            __props__['elastic_load_balancer'] = elastic_load_balancer
            __props__['install_updates_on_boot'] = install_updates_on_boot
            __props__['instance_shutdown_timeout'] = instance_shutdown_timeout
            __props__['name'] = name
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__['stack_id'] = stack_id
            __props__['system_packages'] = system_packages
            __props__['tags'] = tags
            __props__['use_ebs_optimized_instances'] = use_ebs_optimized_instances
            __props__['arn'] = None
        super(StaticWebLayer, __self__).__init__(
            'aws:opsworks/staticWebLayer:StaticWebLayer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_assign_elastic_ips: Optional[pulumi.Input[bool]] = None,
            auto_assign_public_ips: Optional[pulumi.Input[bool]] = None,
            auto_healing: Optional[pulumi.Input[bool]] = None,
            custom_configure_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            custom_deploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            custom_instance_profile_arn: Optional[pulumi.Input[str]] = None,
            custom_json: Optional[pulumi.Input[str]] = None,
            custom_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            custom_setup_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            custom_shutdown_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            custom_undeploy_recipes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            drain_elb_on_shutdown: Optional[pulumi.Input[bool]] = None,
            ebs_volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StaticWebLayerEbsVolumeArgs']]]]] = None,
            elastic_load_balancer: Optional[pulumi.Input[str]] = None,
            install_updates_on_boot: Optional[pulumi.Input[bool]] = None,
            instance_shutdown_timeout: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            system_packages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            use_ebs_optimized_instances: Optional[pulumi.Input[bool]] = None) -> 'StaticWebLayer':
        """
        Get an existing StaticWebLayer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name(ARN) of the layer.
        :param pulumi.Input[bool] auto_assign_elastic_ips: Whether to automatically assign an elastic IP address to the layer's instances.
        :param pulumi.Input[bool] auto_assign_public_ips: For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        :param pulumi.Input[bool] auto_healing: Whether to enable auto-healing for the layer.
        :param pulumi.Input[str] custom_instance_profile_arn: The ARN of an IAM profile that will be used for the layer's instances.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_security_group_ids: Ids for a set of security groups to apply to the layer's instances.
        :param pulumi.Input[bool] drain_elb_on_shutdown: Whether to enable Elastic Load Balancing connection draining.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StaticWebLayerEbsVolumeArgs']]]] ebs_volumes: `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        :param pulumi.Input[str] elastic_load_balancer: Name of an Elastic Load Balancer to attach to this layer
        :param pulumi.Input[bool] install_updates_on_boot: Whether to install OS and package updates on each instance when it boots.
        :param pulumi.Input[int] instance_shutdown_timeout: The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        :param pulumi.Input[str] name: A human-readable name for the layer.
        :param pulumi.Input[str] stack_id: The id of the stack the layer will belong to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] system_packages: Names of a set of system packages to install on the layer's instances.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[bool] use_ebs_optimized_instances: Whether to use EBS-optimized instances.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["auto_assign_elastic_ips"] = auto_assign_elastic_ips
        __props__["auto_assign_public_ips"] = auto_assign_public_ips
        __props__["auto_healing"] = auto_healing
        __props__["custom_configure_recipes"] = custom_configure_recipes
        __props__["custom_deploy_recipes"] = custom_deploy_recipes
        __props__["custom_instance_profile_arn"] = custom_instance_profile_arn
        __props__["custom_json"] = custom_json
        __props__["custom_security_group_ids"] = custom_security_group_ids
        __props__["custom_setup_recipes"] = custom_setup_recipes
        __props__["custom_shutdown_recipes"] = custom_shutdown_recipes
        __props__["custom_undeploy_recipes"] = custom_undeploy_recipes
        __props__["drain_elb_on_shutdown"] = drain_elb_on_shutdown
        __props__["ebs_volumes"] = ebs_volumes
        __props__["elastic_load_balancer"] = elastic_load_balancer
        __props__["install_updates_on_boot"] = install_updates_on_boot
        __props__["instance_shutdown_timeout"] = instance_shutdown_timeout
        __props__["name"] = name
        __props__["stack_id"] = stack_id
        __props__["system_packages"] = system_packages
        __props__["tags"] = tags
        __props__["use_ebs_optimized_instances"] = use_ebs_optimized_instances
        return StaticWebLayer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name(ARN) of the layer.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoAssignElasticIps")
    def auto_assign_elastic_ips(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to automatically assign an elastic IP address to the layer's instances.
        """
        return pulumi.get(self, "auto_assign_elastic_ips")

    @property
    @pulumi.getter(name="autoAssignPublicIps")
    def auto_assign_public_ips(self) -> pulumi.Output[Optional[bool]]:
        """
        For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        """
        return pulumi.get(self, "auto_assign_public_ips")

    @property
    @pulumi.getter(name="autoHealing")
    def auto_healing(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable auto-healing for the layer.
        """
        return pulumi.get(self, "auto_healing")

    @property
    @pulumi.getter(name="customConfigureRecipes")
    def custom_configure_recipes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_configure_recipes")

    @property
    @pulumi.getter(name="customDeployRecipes")
    def custom_deploy_recipes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_deploy_recipes")

    @property
    @pulumi.getter(name="customInstanceProfileArn")
    def custom_instance_profile_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of an IAM profile that will be used for the layer's instances.
        """
        return pulumi.get(self, "custom_instance_profile_arn")

    @property
    @pulumi.getter(name="customJson")
    def custom_json(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "custom_json")

    @property
    @pulumi.getter(name="customSecurityGroupIds")
    def custom_security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Ids for a set of security groups to apply to the layer's instances.
        """
        return pulumi.get(self, "custom_security_group_ids")

    @property
    @pulumi.getter(name="customSetupRecipes")
    def custom_setup_recipes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_setup_recipes")

    @property
    @pulumi.getter(name="customShutdownRecipes")
    def custom_shutdown_recipes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_shutdown_recipes")

    @property
    @pulumi.getter(name="customUndeployRecipes")
    def custom_undeploy_recipes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_undeploy_recipes")

    @property
    @pulumi.getter(name="drainElbOnShutdown")
    def drain_elb_on_shutdown(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable Elastic Load Balancing connection draining.
        """
        return pulumi.get(self, "drain_elb_on_shutdown")

    @property
    @pulumi.getter(name="ebsVolumes")
    def ebs_volumes(self) -> pulumi.Output[Optional[Sequence['outputs.StaticWebLayerEbsVolume']]]:
        """
        `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        """
        return pulumi.get(self, "ebs_volumes")

    @property
    @pulumi.getter(name="elasticLoadBalancer")
    def elastic_load_balancer(self) -> pulumi.Output[Optional[str]]:
        """
        Name of an Elastic Load Balancer to attach to this layer
        """
        return pulumi.get(self, "elastic_load_balancer")

    @property
    @pulumi.getter(name="installUpdatesOnBoot")
    def install_updates_on_boot(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to install OS and package updates on each instance when it boots.
        """
        return pulumi.get(self, "install_updates_on_boot")

    @property
    @pulumi.getter(name="instanceShutdownTimeout")
    def instance_shutdown_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        """
        return pulumi.get(self, "instance_shutdown_timeout")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-readable name for the layer.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        The id of the stack the layer will belong to.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="systemPackages")
    def system_packages(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Names of a set of system packages to install on the layer's instances.
        """
        return pulumi.get(self, "system_packages")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="useEbsOptimizedInstances")
    def use_ebs_optimized_instances(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use EBS-optimized instances.
        """
        return pulumi.get(self, "use_ebs_optimized_instances")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

