# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['VpcEndpointService']


class VpcEndpointService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptance_required: Optional[pulumi.Input[bool]] = None,
                 allowed_principals: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 network_load_balancer_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC Endpoint Service resource.
        Service consumers can create an _Interface_ VPC Endpoint to connect to the service.

        > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
        both a standalone VPC Endpoint Service Allowed Principal resource
        and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
        a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
        and will overwrite the association.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["example"]["arn"]])
        ```
        ### Basic w/ Tags

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["example"]["arn"]],
            tags={
                "Environment": "test",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[List[pulumi.Input[str]]] network_load_balancer_arns: The ARNs of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

            if acceptance_required is None:
                raise TypeError("Missing required property 'acceptance_required'")
            __props__['acceptance_required'] = acceptance_required
            __props__['allowed_principals'] = allowed_principals
            if network_load_balancer_arns is None:
                raise TypeError("Missing required property 'network_load_balancer_arns'")
            __props__['network_load_balancer_arns'] = network_load_balancer_arns
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['availability_zones'] = None
            __props__['base_endpoint_dns_names'] = None
            __props__['manages_vpc_endpoints'] = None
            __props__['private_dns_name'] = None
            __props__['service_name'] = None
            __props__['service_type'] = None
            __props__['state'] = None
        super(VpcEndpointService, __self__).__init__(
            'aws:ec2/vpcEndpointService:VpcEndpointService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acceptance_required: Optional[pulumi.Input[bool]] = None,
            allowed_principals: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            availability_zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            base_endpoint_dns_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            manages_vpc_endpoints: Optional[pulumi.Input[bool]] = None,
            network_load_balancer_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            private_dns_name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            service_type: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VpcEndpointService':
        """
        Get an existing VpcEndpointService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the VPC endpoint service.
        :param pulumi.Input[List[pulumi.Input[str]]] availability_zones: The Availability Zones in which the service is available.
        :param pulumi.Input[List[pulumi.Input[str]]] base_endpoint_dns_names: The DNS names for the service.
        :param pulumi.Input[bool] manages_vpc_endpoints: Whether or not the service manages its VPC endpoints - `true` or `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_load_balancer_arns: The ARNs of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[str] private_dns_name: The private DNS name for the service.
        :param pulumi.Input[str] service_name: The service name.
        :param pulumi.Input[str] service_type: The service type, `Gateway` or `Interface`.
        :param pulumi.Input[str] state: The state of the VPC endpoint service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acceptance_required"] = acceptance_required
        __props__["allowed_principals"] = allowed_principals
        __props__["arn"] = arn
        __props__["availability_zones"] = availability_zones
        __props__["base_endpoint_dns_names"] = base_endpoint_dns_names
        __props__["manages_vpc_endpoints"] = manages_vpc_endpoints
        __props__["network_load_balancer_arns"] = network_load_balancer_arns
        __props__["private_dns_name"] = private_dns_name
        __props__["service_name"] = service_name
        __props__["service_type"] = service_type
        __props__["state"] = state
        __props__["tags"] = tags
        return VpcEndpointService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> bool:
        """
        Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        """
        return pulumi.get(self, "acceptance_required")

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> List[str]:
        """
        The ARNs of one or more principals allowed to discover the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the VPC endpoint service.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> List[str]:
        """
        The Availability Zones in which the service is available.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="baseEndpointDnsNames")
    def base_endpoint_dns_names(self) -> List[str]:
        """
        The DNS names for the service.
        """
        return pulumi.get(self, "base_endpoint_dns_names")

    @property
    @pulumi.getter(name="managesVpcEndpoints")
    def manages_vpc_endpoints(self) -> bool:
        """
        Whether or not the service manages its VPC endpoints - `true` or `false`.
        """
        return pulumi.get(self, "manages_vpc_endpoints")

    @property
    @pulumi.getter(name="networkLoadBalancerArns")
    def network_load_balancer_arns(self) -> List[str]:
        """
        The ARNs of one or more Network Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "network_load_balancer_arns")

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> str:
        """
        The private DNS name for the service.
        """
        return pulumi.get(self, "private_dns_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The service name.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        The service type, `Gateway` or `Interface`.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the VPC endpoint service.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

