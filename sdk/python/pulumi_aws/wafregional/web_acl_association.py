# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class WebAclAssociation(pulumi.CustomResource):
    resource_arn: pulumi.Output[str]
    """
    ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
    """
    web_acl_id: pulumi.Output[str]
    """
    The ID of the WAF Regional WebACL to create an association.
    """
    def __init__(__self__, resource_name, opts=None, resource_arn=None, web_acl_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an association with WAF Regional Web ACL.

        > **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.

        ## Application Load Balancer Association Example
        {{% example %}}

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[{
            "type": "IPV4",
            "value": "192.0.7.0/24",
        }])
        foo_rule = aws.wafregional.Rule("fooRule",
            metric_name="tfWAFRule",
            predicates=[{
                "dataId": ipset.id,
                "negated": False,
                "type": "IPMatch",
            }])
        foo_web_acl = aws.wafregional.WebAcl("fooWebAcl",
            default_action={
                "type": "ALLOW",
            },
            metric_name="foo",
            rules=[{
                "action": {
                    "type": "BLOCK",
                },
                "priority": 1,
                "rule_id": foo_rule.id,
            }])
        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
        available = aws.get_availability_zones()
        foo_subnet = aws.ec2.Subnet("fooSubnet",
            availability_zone=available.names[0],
            cidr_block="10.1.1.0/24",
            vpc_id=foo_vpc.id)
        bar = aws.ec2.Subnet("bar",
            availability_zone=available.names[1],
            cidr_block="10.1.2.0/24",
            vpc_id=foo_vpc.id)
        foo_load_balancer = aws.alb.LoadBalancer("fooLoadBalancer",
            internal=True,
            subnets=[
                foo_subnet.id,
                bar.id,
            ])
        foo_web_acl_association = aws.wafregional.WebAclAssociation("fooWebAclAssociation",
            resource_arn=foo_load_balancer.arn,
            web_acl_id=foo_web_acl.id)
        ```

        {{% /example %}}
        ## API Gateway Association Example
        {{% example %}}

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[{
            "type": "IPV4",
            "value": "192.0.7.0/24",
        }])
        foo_rule = aws.wafregional.Rule("fooRule",
            metric_name="tfWAFRule",
            predicates=[{
                "dataId": ipset.id,
                "negated": False,
                "type": "IPMatch",
            }])
        foo_web_acl = aws.wafregional.WebAcl("fooWebAcl",
            default_action={
                "type": "ALLOW",
            },
            metric_name="foo",
            rules=[{
                "action": {
                    "type": "BLOCK",
                },
                "priority": 1,
                "rule_id": foo_rule.id,
            }])
        test_rest_api = aws.apigateway.RestApi("testRestApi")
        test_resource = aws.apigateway.Resource("testResource",
            parent_id=test_rest_api.root_resource_id,
            path_part="test",
            rest_api=test_rest_api.id)
        test_method = aws.apigateway.Method("testMethod",
            authorization="NONE",
            http_method="GET",
            resource_id=test_resource.id,
            rest_api=test_rest_api.id)
        test_method_response = aws.apigateway.MethodResponse("testMethodResponse",
            http_method=test_method.http_method,
            resource_id=test_resource.id,
            rest_api=test_rest_api.id,
            status_code="400")
        test_integration = aws.apigateway.Integration("testIntegration",
            http_method=test_method.http_method,
            integration_http_method="GET",
            resource_id=test_resource.id,
            rest_api=test_rest_api.id,
            type="HTTP",
            uri="http://www.example.com")
        test_integration_response = aws.apigateway.IntegrationResponse("testIntegrationResponse",
            http_method=test_integration.http_method,
            resource_id=test_resource.id,
            rest_api=test_rest_api.id,
            status_code=test_method_response.status_code)
        test_deployment = aws.apigateway.Deployment("testDeployment", rest_api=test_rest_api.id)
        test_stage = aws.apigateway.Stage("testStage",
            deployment=test_deployment.id,
            rest_api=test_rest_api.id,
            stage_name="test")
        association = aws.wafregional.WebAclAssociation("association",
            resource_arn=test_stage.arn,
            web_acl_id=foo_web_acl.id)
        ```

        {{% /example %}}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        :param pulumi.Input[str] web_acl_id: The ID of the WAF Regional WebACL to create an association.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if resource_arn is None:
                raise TypeError("Missing required property 'resource_arn'")
            __props__['resource_arn'] = resource_arn
            if web_acl_id is None:
                raise TypeError("Missing required property 'web_acl_id'")
            __props__['web_acl_id'] = web_acl_id
        super(WebAclAssociation, __self__).__init__(
            'aws:wafregional/webAclAssociation:WebAclAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, resource_arn=None, web_acl_id=None):
        """
        Get an existing WebAclAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        :param pulumi.Input[str] web_acl_id: The ID of the WAF Regional WebACL to create an association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["resource_arn"] = resource_arn
        __props__["web_acl_id"] = web_acl_id
        return WebAclAssociation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

