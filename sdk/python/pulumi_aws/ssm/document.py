# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Document(pulumi.CustomResource):
    arn: pulumi.Output[str]
    content: pulumi.Output[str]
    """
    The JSON or YAML content of the document.
    """
    created_date: pulumi.Output[str]
    """
    The date the document was created.
    """
    default_version: pulumi.Output[str]
    """
    The default version of the document.
    """
    description: pulumi.Output[str]
    """
    The description of the document.
    """
    document_format: pulumi.Output[str]
    """
    The format of the document. Valid document types include: `JSON` and `YAML`
    """
    document_type: pulumi.Output[str]
    """
    The type of the document. Valid document types include: `Command`, `Policy`, `Automation` and `Session`
    """
    hash: pulumi.Output[str]
    """
    The sha1 or sha256 of the document content
    """
    hash_type: pulumi.Output[str]
    """
    "Sha1" "Sha256". The hashing algorithm used when hashing the content.
    """
    latest_version: pulumi.Output[str]
    """
    The latest version of the document.
    """
    name: pulumi.Output[str]
    """
    The name of the document.
    """
    owner: pulumi.Output[str]
    """
    The AWS user account of the person who created the document.
    """
    parameters: pulumi.Output[list]
    """
    The parameters that are available to this document.
    """
    permissions: pulumi.Output[dict]
    """
    Additional Permissions to attach to the document. See Permissions below for details.
    """
    platform_types: pulumi.Output[list]
    """
    A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
    """
    schema_version: pulumi.Output[str]
    """
    The schema version of the document.
    """
    status: pulumi.Output[str]
    """
    "Creating", "Active" or "Deleting". The current status of the document.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the object.
    """
    def __init__(__self__, resource_name, opts=None, content=None, document_format=None, document_type=None, name=None, permissions=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an SSM Document resource
        
        > **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
        or greater can update their content once created, see [SSM Schema Features][1]. To update a document with an older
        schema version you must recreate the resource.
        
        ## Permissions
        
        The permissions attribute specifies how you want to share the document. If you share a document privately,
        you must specify the AWS user account IDs for those people who can use the document. If you share a document
        publicly, you must specify All as the account ID.
        
        The permissions mapping supports the following:
        
        * `type` - The permission type for the document. The permission type can be `Share`.
        * `account_ids` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: The JSON or YAML content of the document.
        :param pulumi.Input[str] document_format: The format of the document. Valid document types include: `JSON` and `YAML`
        :param pulumi.Input[str] document_type: The type of the document. Valid document types include: `Command`, `Policy`, `Automation` and `Session`
        :param pulumi.Input[str] name: The name of the document.
        :param pulumi.Input[dict] permissions: Additional Permissions to attach to the document. See Permissions below for details.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the object.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not content:
            raise TypeError('Missing required property content')
        __props__['content'] = content

        __props__['document_format'] = document_format

        if not document_type:
            raise TypeError('Missing required property document_type')
        __props__['document_type'] = document_type

        __props__['name'] = name

        __props__['permissions'] = permissions

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['created_date'] = None
        __props__['default_version'] = None
        __props__['description'] = None
        __props__['hash'] = None
        __props__['hash_type'] = None
        __props__['latest_version'] = None
        __props__['owner'] = None
        __props__['parameters'] = None
        __props__['platform_types'] = None
        __props__['schema_version'] = None
        __props__['status'] = None

        super(Document, __self__).__init__(
            'aws:ssm/document:Document',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

