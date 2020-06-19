# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKeyResult:
    """
    A collection of values returned by getKey.
    """
    def __init__(__self__, created_date=None, description=None, enabled=None, id=None, last_updated_date=None, name=None, tags=None, value=None):
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        __self__.created_date = created_date
        """
        The date and time when the API Key was created.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the API Key.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        Specifies whether the API Key is enabled.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        Set to the ID of the API Key.
        """
        if last_updated_date and not isinstance(last_updated_date, str):
            raise TypeError("Expected argument 'last_updated_date' to be a str")
        __self__.last_updated_date = last_updated_date
        """
        The date and time when the API Key was last updated.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Set to the name of the API Key.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A map of tags for the resource.
        """
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        __self__.value = value
        """
        Set to the value of the API Key.
        """
class AwaitableGetKeyResult(GetKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyResult(
            created_date=self.created_date,
            description=self.description,
            enabled=self.enabled,
            id=self.id,
            last_updated_date=self.last_updated_date,
            name=self.name,
            tags=self.tags,
            value=self.value)

def get_key(id=None,tags=None,opts=None):
    """
    Use this data source to get the name and value of a pre-existing API Key, for
    example to supply credentials for a dependency microservice.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    my_api_key = aws.apigateway.get_key(id="ru3mpjgse6")
    ```


    :param str id: The ID of the API Key to look up.
    :param dict tags: A map of tags for the resource.
    """
    __args__ = dict()


    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:apigateway/getKey:getKey', __args__, opts=opts).value

    return AwaitableGetKeyResult(
        created_date=__ret__.get('createdDate'),
        description=__ret__.get('description'),
        enabled=__ret__.get('enabled'),
        id=__ret__.get('id'),
        last_updated_date=__ret__.get('lastUpdatedDate'),
        name=__ret__.get('name'),
        tags=__ret__.get('tags'),
        value=__ret__.get('value'))
