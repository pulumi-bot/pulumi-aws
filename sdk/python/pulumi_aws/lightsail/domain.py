# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] domain_name: The name of the Lightsail domain to manage
        """
        pulumi.set(__self__, "domain_name", domain_name)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The name of the Lightsail domain to manage
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)


@pulumi.input_type
class _DomainState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Domain resources.
        :param pulumi.Input[str] arn: The ARN of the Lightsail domain
        :param pulumi.Input[str] domain_name: The name of the Lightsail domain to manage
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Lightsail domain
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Lightsail domain to manage
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a domain resource for the specified domain (e.g., example.com).
        You cannot register a new domain name using Lightsail. You must register
        a domain name using Amazon Route 53 or another domain name registrar.
        If you have already registered your domain, you can enter its name in
        this parameter to manage the DNS records for that domain.

        > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        domain_test = aws.lightsail.Domain("domainTest", domain_name="mydomain.com")
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: The name of the Lightsail domain to manage
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a domain resource for the specified domain (e.g., example.com).
        You cannot register a new domain name using Lightsail. You must register
        a domain name using Amazon Route 53 or another domain name registrar.
        If you have already registered your domain, you can enter its name in
        this parameter to manage the DNS records for that domain.

        > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        domain_test = aws.lightsail.Domain("domainTest", domain_name="mydomain.com")
        ```

        :param str resource_name_: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = DomainArgs.__new__(DomainArgs)

            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["arn"] = None
        super(Domain, __self__).__init__(
            'aws:lightsail/domain:Domain',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Lightsail domain
        :param pulumi.Input[str] domain_name: The name of the Lightsail domain to manage
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainState.__new__(_DomainState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["domain_name"] = domain_name
        return Domain(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Lightsail domain
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The name of the Lightsail domain to manage
        """
        return pulumi.get(self, "domain_name")

