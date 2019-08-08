# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSolutionStackResult:
    """
    A collection of values returned by getSolutionStack.
    """
    def __init__(__self__, most_recent=None, name=None, name_regex=None, id=None):
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the solution stack.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_solution_stack(most_recent=None,name_regex=None,opts=None):
    """
    Use this data source to get the name of a elastic beanstalk solution stack.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_solution_stack.html.markdown.
    """
    __args__ = dict()

    __args__['mostRecent'] = most_recent
    __args__['nameRegex'] = name_regex
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticbeanstalk/getSolutionStack:getSolutionStack', __args__, opts=opts).value

    return GetSolutionStackResult(
        most_recent=__ret__.get('mostRecent'),
        name=__ret__.get('name'),
        name_regex=__ret__.get('nameRegex'),
        id=__ret__.get('id'))
