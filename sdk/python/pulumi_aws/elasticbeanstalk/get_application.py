# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetApplicationResult:
    """
    A collection of values returned by getApplication.
    """
    def __init__(__self__, appversion_lifecycle=None, arn=None, description=None, id=None, name=None):
        if appversion_lifecycle and not isinstance(appversion_lifecycle, dict):
            raise TypeError("Expected argument 'appversion_lifecycle' to be a dict")
        __self__.appversion_lifecycle = appversion_lifecycle
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the application.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Short description of the application
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            appversion_lifecycle=self.appversion_lifecycle,
            arn=self.arn,
            description=self.description,
            id=self.id,
            name=self.name)

def get_application(name=None,opts=None):
    """
    Retrieve information about an Elastic Beanstalk Application.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.elasticbeanstalk.get_application(name="example")
    pulumi.export("arn", example.arn)
    pulumi.export("description", example.description)
    ```



    :param str name: The name of the application
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticbeanstalk/getApplication:getApplication', __args__, opts=opts).value

    return AwaitableGetApplicationResult(
        appversion_lifecycle=__ret__.get('appversionLifecycle'),
        arn=__ret__.get('arn'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'))
