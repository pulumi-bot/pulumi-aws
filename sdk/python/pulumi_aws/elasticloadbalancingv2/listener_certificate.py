# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ListenerCertificate']

warnings.warn("""aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate""", DeprecationWarning)


class ListenerCertificate(pulumi.CustomResource):
    warnings.warn("""aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 listener_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Load Balancer Listener Certificate resource.

        This resource is for additional certificates and does not replace the default certificate on the listener.

        > **Note:** `alb.ListenerCertificate` is known as `lb.ListenerCertificate`. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_certificate = aws.acm.Certificate("exampleCertificate")
        # ...
        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener")
        # ...
        example_listener_certificate = aws.lb.ListenerCertificate("exampleListenerCertificate",
            listener_arn=front_end_listener.arn,
            certificate_arn=example_certificate.arn)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate to attach to the listener.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the certificate.
        """
        pulumi.log.warn("ListenerCertificate is deprecated: aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate")
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

            if certificate_arn is None:
                raise TypeError("Missing required property 'certificate_arn'")
            __props__['certificate_arn'] = certificate_arn
            if listener_arn is None:
                raise TypeError("Missing required property 'listener_arn'")
            __props__['listener_arn'] = listener_arn
        super(ListenerCertificate, __self__).__init__(
            'aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            listener_arn: Optional[pulumi.Input[str]] = None) -> 'ListenerCertificate':
        """
        Get an existing ListenerCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate to attach to the listener.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_arn"] = certificate_arn
        __props__["listener_arn"] = listener_arn
        return ListenerCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the certificate to attach to the listener.
        """
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="listenerArn")
    def listener_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the listener to which to attach the certificate.
        """
        return pulumi.get(self, "listener_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

