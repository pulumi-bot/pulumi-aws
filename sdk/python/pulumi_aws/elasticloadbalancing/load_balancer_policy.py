# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables

warnings.warn("aws.elasticloadbalancing.LoadBalancerPolicy has been deprecated in favor of aws.elb.LoadBalancerPolicy", DeprecationWarning)


class LoadBalancerPolicy(pulumi.CustomResource):
    load_balancer_name: pulumi.Output[str]
    """
    The load balancer on which the policy is defined.
    """
    policy_attributes: pulumi.Output[list]
    """
    Policy attribute to apply to the policy.

      * `name` (`str`)
      * `value` (`str`)
    """
    policy_name: pulumi.Output[str]
    """
    The name of the load balancer policy.
    """
    policy_type_name: pulumi.Output[str]
    """
    The policy type.
    """
    warnings.warn("aws.elasticloadbalancing.LoadBalancerPolicy has been deprecated in favor of aws.elb.LoadBalancerPolicy", DeprecationWarning)

    def __init__(__self__, resource_name, opts=None, load_balancer_name=None, policy_attributes=None, policy_name=None, policy_type_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a load balancer policy, which can be attached to an ELB listener or backend server.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        wu_tang = aws.elb.LoadBalancer("wu-tang",
            availability_zones=["us-east-1a"],
            listeners=[{
                "instance_port": 443,
                "instanceProtocol": "http",
                "lb_port": 443,
                "lbProtocol": "https",
                "sslCertificateId": "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
            }],
            tags={
                "Name": "wu-tang",
            })
        wu_tang_ca_pubkey_policy = aws.elb.LoadBalancerPolicy("wu-tang-ca-pubkey-policy",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-ca-pubkey-policy",
            policy_type_name="PublicKeyPolicyType",
            policy_attributes=[{
                "name": "PublicKey",
                "value": (lambda path: open(path).read())("wu-tang-pubkey"),
            }])
        wu_tang_root_ca_backend_auth_policy = aws.elb.LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-root-ca-backend-auth-policy",
            policy_type_name="BackendServerAuthenticationPolicyType",
            policy_attributes=[{
                "name": "PublicKeyPolicyName",
                "value": aws_load_balancer_policy["wu-tang-root-ca-pubkey-policy"]["policy_name"],
            }])
        wu_tang_ssl = aws.elb.LoadBalancerPolicy("wu-tang-ssl",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-ssl",
            policy_type_name="SSLNegotiationPolicyType",
            policy_attributes=[
                {
                    "name": "ECDHE-ECDSA-AES128-GCM-SHA256",
                    "value": "true",
                },
                {
                    "name": "Protocol-TLSv1.2",
                    "value": "true",
                },
            ])
        wu_tang_ssl_tls_1_1 = aws.elb.LoadBalancerPolicy("wu-tang-ssl-tls-1-1",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-ssl",
            policy_type_name="SSLNegotiationPolicyType",
            policy_attributes=[{
                "name": "Reference-Security-Policy",
                "value": "ELBSecurityPolicy-TLS-1-1-2017-01",
            }])
        wu_tang_backend_auth_policies_443 = aws.elb.LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443",
            load_balancer_name=wu_tang.name,
            instance_port=443,
            policy_names=[wu_tang_root_ca_backend_auth_policy.policy_name])
        wu_tang_listener_policies_443 = aws.elb.ListenerPolicy("wu-tang-listener-policies-443",
            load_balancer_name=wu_tang.name,
            load_balancer_port=443,
            policy_names=[wu_tang_ssl.policy_name])
        ```

        Where the file `pubkey` in the current directory contains only the _public key_ of the certificate.

        ```python
        import pulumi
        ```

        This example shows how to enable backend authentication for an ELB as well as customize the TLS settings.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] load_balancer_name: The load balancer on which the policy is defined.
        :param pulumi.Input[list] policy_attributes: Policy attribute to apply to the policy.
        :param pulumi.Input[str] policy_name: The name of the load balancer policy.
        :param pulumi.Input[str] policy_type_name: The policy type.

        The **policy_attributes** object supports the following:

          * `name` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
        """
        pulumi.log.warn("LoadBalancerPolicy is deprecated: aws.elasticloadbalancing.LoadBalancerPolicy has been deprecated in favor of aws.elb.LoadBalancerPolicy")
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

            if load_balancer_name is None:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__['loadBalancerName'] = load_balancer_name
            __props__['policyAttributes'] = policy_attributes
            if policy_name is None:
                raise TypeError("Missing required property 'policy_name'")
            __props__['policyName'] = policy_name
            if policy_type_name is None:
                raise TypeError("Missing required property 'policy_type_name'")
            __props__['policyTypeName'] = policy_type_name
        super(LoadBalancerPolicy, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, load_balancer_name=None, policy_attributes=None, policy_name=None, policy_type_name=None):
        """
        Get an existing LoadBalancerPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] load_balancer_name: The load balancer on which the policy is defined.
        :param pulumi.Input[list] policy_attributes: Policy attribute to apply to the policy.
        :param pulumi.Input[str] policy_name: The name of the load balancer policy.
        :param pulumi.Input[str] policy_type_name: The policy type.

        The **policy_attributes** object supports the following:

          * `name` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["load_balancer_name"] = load_balancer_name
        __props__["policy_attributes"] = policy_attributes
        __props__["policy_name"] = policy_name
        __props__["policy_type_name"] = policy_type_name
        return LoadBalancerPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
