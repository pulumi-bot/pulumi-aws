# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['UserPoolDomainArgs', 'UserPoolDomain']

@pulumi.input_type
class UserPoolDomainArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 user_pool_id: pulumi.Input[str],
                 certificate_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserPoolDomain resource.
        :param pulumi.Input[str] domain: The domain string.
        :param pulumi.Input[str] user_pool_id: The user pool ID.
        :param pulumi.Input[str] certificate_arn: The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "user_pool_id", user_pool_id)
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain string.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Input[str]:
        """
        The user pool ID.
        """
        return pulumi.get(self, "user_pool_id")

    @user_pool_id.setter
    def user_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_pool_id", value)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
        """
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)


class UserPoolDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPoolDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cognito User Pool Domain resource.

        ## Example Usage
        ### Amazon Cognito domain

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cognito.UserPool("example")
        main = aws.cognito.UserPoolDomain("main",
            domain="example-domain",
            user_pool_id=example.id)
        ```
        ### Custom Cognito domain

        ```python
        import pulumi
        import pulumi_aws as aws

        example_user_pool = aws.cognito.UserPool("exampleUserPool")
        main = aws.cognito.UserPoolDomain("main",
            domain="example-domain.example.com",
            certificate_arn=aws_acm_certificate["cert"]["arn"],
            user_pool_id=example_user_pool.id)
        example_zone = aws.route53.get_zone(name="example.com")
        auth_cognito__a = aws.route53.Record("auth-cognito-A",
            name=main.domain,
            type="A",
            zone_id=example_zone.zone_id,
            aliases=[aws.route53.RecordAliasArgs(
                evaluate_target_health=False,
                name=main.cloudfront_distribution_arn,
                zone_id="Z2FDTNDATAQYW2",
            )])
        ```

        ## Import

        Cognito User Pool Domains can be imported using the `domain`, e.g.

        ```sh
         $ pulumi import aws:cognito/userPoolDomain:UserPoolDomain main <domain>
        ```

        :param str resource_name: The name of the resource.
        :param UserPoolDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Cognito User Pool Domain resource.

        ## Example Usage
        ### Amazon Cognito domain

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cognito.UserPool("example")
        main = aws.cognito.UserPoolDomain("main",
            domain="example-domain",
            user_pool_id=example.id)
        ```
        ### Custom Cognito domain

        ```python
        import pulumi
        import pulumi_aws as aws

        example_user_pool = aws.cognito.UserPool("exampleUserPool")
        main = aws.cognito.UserPoolDomain("main",
            domain="example-domain.example.com",
            certificate_arn=aws_acm_certificate["cert"]["arn"],
            user_pool_id=example_user_pool.id)
        example_zone = aws.route53.get_zone(name="example.com")
        auth_cognito__a = aws.route53.Record("auth-cognito-A",
            name=main.domain,
            type="A",
            zone_id=example_zone.zone_id,
            aliases=[aws.route53.RecordAliasArgs(
                evaluate_target_health=False,
                name=main.cloudfront_distribution_arn,
                zone_id="Z2FDTNDATAQYW2",
            )])
        ```

        ## Import

        Cognito User Pool Domains can be imported using the `domain`, e.g.

        ```sh
         $ pulumi import aws:cognito/userPoolDomain:UserPoolDomain main <domain>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
        :param pulumi.Input[str] domain: The domain string.
        :param pulumi.Input[str] user_pool_id: The user pool ID.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPoolDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['certificate_arn'] = certificate_arn
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            if user_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_pool_id'")
            __props__['user_pool_id'] = user_pool_id
            __props__['aws_account_id'] = None
            __props__['cloudfront_distribution_arn'] = None
            __props__['s3_bucket'] = None
            __props__['version'] = None
        super(UserPoolDomain, __self__).__init__(
            'aws:cognito/userPoolDomain:UserPoolDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_account_id: Optional[pulumi.Input[str]] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            cloudfront_distribution_arn: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            s3_bucket: Optional[pulumi.Input[str]] = None,
            user_pool_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'UserPoolDomain':
        """
        Get an existing UserPoolDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_account_id: The AWS account ID for the user pool owner.
        :param pulumi.Input[str] certificate_arn: The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
        :param pulumi.Input[str] cloudfront_distribution_arn: The URL of the CloudFront distribution. This is required to generate the ALIAS `route53.Record`
        :param pulumi.Input[str] domain: The domain string.
        :param pulumi.Input[str] s3_bucket: The S3 bucket where the static files for this domain are stored.
        :param pulumi.Input[str] user_pool_id: The user pool ID.
        :param pulumi.Input[str] version: The app version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["aws_account_id"] = aws_account_id
        __props__["certificate_arn"] = certificate_arn
        __props__["cloudfront_distribution_arn"] = cloudfront_distribution_arn
        __props__["domain"] = domain
        __props__["s3_bucket"] = s3_bucket
        __props__["user_pool_id"] = user_pool_id
        __props__["version"] = version
        return UserPoolDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID for the user pool owner.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
        """
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="cloudfrontDistributionArn")
    def cloudfront_distribution_arn(self) -> pulumi.Output[str]:
        """
        The URL of the CloudFront distribution. This is required to generate the ALIAS `route53.Record`
        """
        return pulumi.get(self, "cloudfront_distribution_arn")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain string.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Output[str]:
        """
        The S3 bucket where the static files for this domain are stored.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Output[str]:
        """
        The user pool ID.
        """
        return pulumi.get(self, "user_pool_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The app version.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

