# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetDocumentResult:
    """
    A collection of values returned by getDocument.
    """
    def __init__(__self__, arn=None, content=None, document_format=None, document_type=None, document_version=None, name=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN of the document.
        """
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        __self__.content = content
        """
        The contents of the document.
        """
        if document_format and not isinstance(document_format, str):
            raise TypeError("Expected argument 'document_format' to be a str")
        __self__.document_format = document_format
        if document_type and not isinstance(document_type, str):
            raise TypeError("Expected argument 'document_type' to be a str")
        __self__.document_type = document_type
        """
        The type of the document.
        """
        if document_version and not isinstance(document_version, str):
            raise TypeError("Expected argument 'document_version' to be a str")
        __self__.document_version = document_version
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
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

def get_document(document_format=None,document_version=None,name=None,opts=None):
    """
    Gets the contents of the specified Systems Manager document.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ssm_document.html.markdown.
    """
    __args__ = dict()

    __args__['documentFormat'] = document_format
    __args__['documentVersion'] = document_version
    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ssm/getDocument:getDocument', __args__, opts=opts).value

    return GetDocumentResult(
        arn=__ret__.get('arn'),
        content=__ret__.get('content'),
        document_format=__ret__.get('documentFormat'),
        document_type=__ret__.get('documentType'),
        document_version=__ret__.get('documentVersion'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))
