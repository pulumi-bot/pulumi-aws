# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetDocumentResult(object):
    """
    A collection of values returned by getDocument.
    """
    def __init__(__self__, arn=None, content=None, document_type=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The ARN of the document.
        """
        if content and not isinstance(content, str):
            raise TypeError('Expected argument content to be a str')
        __self__.content = content
        """
        The contents of the document.
        """
        if document_type and not isinstance(document_type, str):
            raise TypeError('Expected argument document_type to be a str')
        __self__.document_type = document_type
        """
        The type of the document.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_document(document_format=None, document_version=None, name=None):
    """
    Gets the contents of the specified Systems Manager document.
    """
    __args__ = dict()

    __args__['documentFormat'] = document_format
    __args__['documentVersion'] = document_version
    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('aws:ssm/getDocument:getDocument', __args__)

    return GetDocumentResult(
        arn=__ret__.get('arn'),
        content=__ret__.get('content'),
        document_type=__ret__.get('documentType'),
        id=__ret__.get('id'))
