# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DocumentationVersion']


class DocumentationVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage an API Gateway Documentation Version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rest_api = aws.apigateway.RestApi("exampleRestApi")
        example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
            location={
                "type": "API",
            },
            properties="{\"description\":\"Example\"}",
            rest_api_id=example_rest_api.id)
        example_documentation_version = aws.apigateway.DocumentationVersion("exampleDocumentationVersion",
            version="example_version",
            rest_api_id=example_rest_api.id,
            description="Example description",
            opts=ResourceOptions(depends_on=[example_documentation_part]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the API documentation version.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        :param pulumi.Input[str] version: The version identifier of the API documentation snapshot.
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

            __props__['description'] = description
            if rest_api_id is None:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__['rest_api_id'] = rest_api_id
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
        super(DocumentationVersion, __self__).__init__(
            'aws:apigateway/documentationVersion:DocumentationVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            rest_api_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'DocumentationVersion':
        """
        Get an existing DocumentationVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the API documentation version.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        :param pulumi.Input[str] version: The version identifier of the API documentation snapshot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["rest_api_id"] = rest_api_id
        __props__["version"] = version
        return DocumentationVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the API documentation version.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> str:
        """
        The ID of the associated Rest API
        """
        return pulumi.get(self, "rest_api_id")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version identifier of the API documentation snapshot.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

