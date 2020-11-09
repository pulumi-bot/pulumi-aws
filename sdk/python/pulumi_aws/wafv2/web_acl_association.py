# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['WebAclAssociation']


class WebAclAssociation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 web_acl_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a WAFv2 Web ACL Association.

        > **NOTE on associating a WAFv2 Web ACL with a Cloudfront distribution:** Do not use this resource to associate a WAFv2 Web ACL with a Cloudfront Distribution. The [AWS API call backing this resource](https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html) notes that you should use the [`web_acl_id`][2] property on the [`cloudfront_distribution`][2] instead.

        [1]: https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html
        [2]: https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#web_acl_id

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rest_api = aws.apigateway.RestApi("exampleRestApi")
        example_resource = aws.apigateway.Resource("exampleResource",
            rest_api=example_rest_api.id,
            parent_id=example_rest_api.root_resource_id,
            path_part="mytestresource")
        example_method = aws.apigateway.Method("exampleMethod",
            rest_api=example_rest_api.id,
            resource_id=example_resource.id,
            http_method="GET",
            authorization="NONE")
        example_integration = aws.apigateway.Integration("exampleIntegration",
            rest_api=example_rest_api.id,
            resource_id=example_resource.id,
            http_method=example_method.http_method,
            type="MOCK")
        example_deployment = aws.apigateway.Deployment("exampleDeployment", rest_api=example_rest_api.id,
        opts=ResourceOptions(depends_on=[example_integration]))
        example_stage = aws.apigateway.Stage("exampleStage",
            stage_name="test",
            rest_api=example_rest_api.id,
            deployment=example_deployment.id)
        example_web_acl = aws.wafv2.WebAcl("exampleWebAcl",
            scope="REGIONAL",
            default_action=aws.wafv2.WebAclDefaultActionArgs(
                allow=aws.wafv2.WebAclDefaultActionAllowArgs(),
            ),
            visibility_config=aws.wafv2.WebAclVisibilityConfigArgs(
                cloudwatch_metrics_enabled=False,
                metric_name="friendly-metric-name",
                sampled_requests_enabled=False,
            ))
        example_web_acl_association = aws.wafv2.WebAclAssociation("exampleWebAclAssociation",
            resource_arn=example_stage.arn,
            web_acl_arn=example_web_acl.arn)
        ```

        ## Import

        WAFv2 Web ACL Association can be imported using `WEB_ACL_ARN,RESOURCE_ARN` e.g.

        ```sh
         $ pulumi import aws:wafv2/webAclAssociation:WebAclAssociation example arn:aws:wafv2:...7ce849ea,arn:aws:apigateway:...ages/name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
        :param pulumi.Input[str] web_acl_arn: The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
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

            if resource_arn is None:
                raise TypeError("Missing required property 'resource_arn'")
            __props__['resource_arn'] = resource_arn
            if web_acl_arn is None:
                raise TypeError("Missing required property 'web_acl_arn'")
            __props__['web_acl_arn'] = web_acl_arn
        super(WebAclAssociation, __self__).__init__(
            'aws:wafv2/webAclAssociation:WebAclAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            resource_arn: Optional[pulumi.Input[str]] = None,
            web_acl_arn: Optional[pulumi.Input[str]] = None) -> 'WebAclAssociation':
        """
        Get an existing WebAclAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
        :param pulumi.Input[str] web_acl_arn: The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["resource_arn"] = resource_arn
        __props__["web_acl_arn"] = web_acl_arn
        return WebAclAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="webAclArn")
    def web_acl_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
        """
        return pulumi.get(self, "web_acl_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

