# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class IntegrationResponse(pulumi.CustomResource):
    api_id: pulumi.Output[str]
    """
    The API identifier.
    """
    content_handling_strategy: pulumi.Output[str]
    """
    How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
    """
    integration_id: pulumi.Output[str]
    """
    The identifier of the `apigatewayv2.Integration`.
    """
    integration_response_key: pulumi.Output[str]
    """
    The integration response key.
    """
    response_templates: pulumi.Output[dict]
    """
    A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
    """
    template_selection_expression: pulumi.Output[str]
    """
    The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
    """
    def __init__(__self__, resource_name, opts=None, api_id=None, content_handling_strategy=None, integration_id=None, integration_response_key=None, response_templates=None, template_selection_expression=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 integration response.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.IntegrationResponse("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            integration_id=aws_apigatewayv2_integration["example"]["id"],
            integration_response_key="/200/")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[str] content_handling_strategy: How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
        :param pulumi.Input[str] integration_id: The identifier of the `apigatewayv2.Integration`.
        :param pulumi.Input[str] integration_response_key: The integration response key.
        :param pulumi.Input[dict] response_templates: A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
        :param pulumi.Input[str] template_selection_expression: The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
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

            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['apiId'] = api_id
            __props__['contentHandlingStrategy'] = content_handling_strategy
            if integration_id is None:
                raise TypeError("Missing required property 'integration_id'")
            __props__['integrationId'] = integration_id
            if integration_response_key is None:
                raise TypeError("Missing required property 'integration_response_key'")
            __props__['integrationResponseKey'] = integration_response_key
            __props__['responseTemplates'] = response_templates
            __props__['templateSelectionExpression'] = template_selection_expression
        super(IntegrationResponse, __self__).__init__(
            'aws:apigatewayv2/integrationResponse:IntegrationResponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_id=None, content_handling_strategy=None, integration_id=None, integration_response_key=None, response_templates=None, template_selection_expression=None):
        """
        Get an existing IntegrationResponse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[str] content_handling_strategy: How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
        :param pulumi.Input[str] integration_id: The identifier of the `apigatewayv2.Integration`.
        :param pulumi.Input[str] integration_response_key: The integration response key.
        :param pulumi.Input[dict] response_templates: A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
        :param pulumi.Input[str] template_selection_expression: The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_id"] = api_id
        __props__["content_handling_strategy"] = content_handling_strategy
        __props__["integration_id"] = integration_id
        __props__["integration_response_key"] = integration_response_key
        __props__["response_templates"] = response_templates
        __props__["template_selection_expression"] = template_selection_expression
        return IntegrationResponse(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
