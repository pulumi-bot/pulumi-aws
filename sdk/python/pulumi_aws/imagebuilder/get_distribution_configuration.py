# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDistributionConfigurationResult',
    'AwaitableGetDistributionConfigurationResult',
    'get_distribution_configuration',
]

@pulumi.output_type
class GetDistributionConfigurationResult:
    """
    A collection of values returned by getDistributionConfiguration.
    """
    def __init__(__self__, arn=None, date_created=None, date_updated=None, description=None, distributions=None, id=None, name=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if date_updated and not isinstance(date_updated, str):
            raise TypeError("Expected argument 'date_updated' to be a str")
        pulumi.set(__self__, "date_updated", date_updated)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if distributions and not isinstance(distributions, list):
            raise TypeError("Expected argument 'distributions' to be a list")
        pulumi.set(__self__, "distributions", distributions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        Date the distribution configuration was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> str:
        """
        Date the distribution configuration was updated.
        """
        return pulumi.get(self, "date_updated")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description to apply to distributed AMI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distributions(self) -> Sequence['outputs.GetDistributionConfigurationDistributionResult']:
        """
        Set of distributions.
        """
        return pulumi.get(self, "distributions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the distribution configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value map of resource tags for the distribution configuration.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDistributionConfigurationResult(GetDistributionConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributionConfigurationResult(
            arn=self.arn,
            date_created=self.date_created,
            date_updated=self.date_updated,
            description=self.description,
            distributions=self.distributions,
            id=self.id,
            name=self.name,
            tags=self.tags)


def get_distribution_configuration(arn: Optional[str] = None,
                                   tags: Optional[Mapping[str, str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDistributionConfigurationResult:
    """
    Provides details about an Image Builder Distribution Configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.imagebuilder.get_distribution_configuration(arn="arn:aws:imagebuilder:us-west-2:aws:distribution-configuration/example")
    ```


    :param str arn: Amazon Resource Name (ARN) of the distribution configuration.
    :param Mapping[str, str] tags: Key-value map of resource tags for the distribution configuration.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:imagebuilder/getDistributionConfiguration:getDistributionConfiguration', __args__, opts=opts, typ=GetDistributionConfigurationResult).value

    return AwaitableGetDistributionConfigurationResult(
        arn=__ret__.arn,
        date_created=__ret__.date_created,
        date_updated=__ret__.date_updated,
        description=__ret__.description,
        distributions=__ret__.distributions,
        id=__ret__.id,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_distribution_configuration)
def get_distribution_configuration_apply(arn: Optional[pulumi.Input[str]] = None,
                                         tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDistributionConfigurationResult]:
    ...
