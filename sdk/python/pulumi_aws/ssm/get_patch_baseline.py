# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetPatchBaselineResult:
    """
    A collection of values returned by getPatchBaseline.
    """
    def __init__(__self__, default_baseline=None, description=None, id=None, name=None, name_prefix=None, operating_system=None, owner=None):
        if default_baseline and not isinstance(default_baseline, bool):
            raise TypeError("Expected argument 'default_baseline' to be a bool")
        __self__.default_baseline = default_baseline
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the baseline.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the baseline.
        """
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        __self__.name_prefix = name_prefix
        if operating_system and not isinstance(operating_system, str):
            raise TypeError("Expected argument 'operating_system' to be a str")
        __self__.operating_system = operating_system
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        __self__.owner = owner
class AwaitableGetPatchBaselineResult(GetPatchBaselineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPatchBaselineResult(
            default_baseline=self.default_baseline,
            description=self.description,
            id=self.id,
            name=self.name,
            name_prefix=self.name_prefix,
            operating_system=self.operating_system,
            owner=self.owner)

def get_patch_baseline(default_baseline=None,name_prefix=None,operating_system=None,owner=None,opts=None):
    """
    Provides an SSM Patch Baseline data source. Useful if you wish to reuse the default baselines provided.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ssm_patch_baseline.html.markdown.


    :param bool default_baseline: Filters the results against the baselines default_baseline field.
    :param str name_prefix: Filter results by the baseline name prefix.
    :param str operating_system: The specified OS for the baseline.
    :param str owner: The owner of the baseline. Valid values: `All`, `AWS`, `Self` (the current account).
    """
    __args__ = dict()


    __args__['defaultBaseline'] = default_baseline
    __args__['namePrefix'] = name_prefix
    __args__['operatingSystem'] = operating_system
    __args__['owner'] = owner
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ssm/getPatchBaseline:getPatchBaseline', __args__, opts=opts).value

    return AwaitableGetPatchBaselineResult(
        default_baseline=__ret__.get('defaultBaseline'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        name_prefix=__ret__.get('namePrefix'),
        operating_system=__ret__.get('operatingSystem'),
        owner=__ret__.get('owner'))
