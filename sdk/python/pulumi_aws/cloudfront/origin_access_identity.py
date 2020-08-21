# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['OriginAccessIdentity']


class OriginAccessIdentity(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Amazon CloudFront origin access identity.

        For information about CloudFront distributions, see the
        [Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For more information on generating
        origin access identities, see
        [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].

        ## Example Usage

        The following example below creates a CloudFront origin access identity.

        ```python
        import pulumi
        import pulumi_aws as aws

        origin_access_identity = aws.cloudfront.OriginAccessIdentity("originAccessIdentity", comment="Some comment")
        ```
        ## Using With CloudFront

        Normally, when referencing an origin access identity in CloudFront, you need to
        prefix the ID with the `origin-access-identity/cloudfront/` special path.
        The `cloudfront_access_identity_path` allows this to be circumvented.
        The below snippet demonstrates use with the `s3_origin_config` structure for the
        [`cloudfront.Distribution`][3] resource:

        ```python
        import pulumi
        ```

        ### Updating your bucket policy

        Note that the AWS API may translate the `s3_canonical_user_id` `CanonicalUser`
        principal into an `AWS` IAM ARN principal when supplied in an
        [`s3.Bucket`][4] bucket policy, causing spurious diffs. If
        you see this behaviour, use the `iam_arn` instead:

        ```python
        import pulumi
        import pulumi_aws as aws

        s3_policy = aws.iam.get_policy_document(statements=[{
            "actions": ["s3:GetObject"],
            "resources": [f"{aws_s3_bucket['example']['arn']}/*"],
            "principals": [{
                "type": "AWS",
                "identifiers": [aws_cloudfront_origin_access_identity["origin_access_identity"]["iam_arn"]],
            }],
        }])
        example = aws.s3.BucketPolicy("example",
            bucket=aws_s3_bucket["example"]["id"],
            policy=s3_policy.json)
        ```

        [1]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html
        [2]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html
        [3]: https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html
        [4]: https://www.terraform.io/docs/providers/aws/r/s3_bucket.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: An optional comment for the origin access identity.
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

            __props__['comment'] = comment
            __props__['caller_reference'] = None
            __props__['cloudfront_access_identity_path'] = None
            __props__['etag'] = None
            __props__['iam_arn'] = None
            __props__['s3_canonical_user_id'] = None
        super(OriginAccessIdentity, __self__).__init__(
            'aws:cloudfront/originAccessIdentity:OriginAccessIdentity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            caller_reference: Optional[pulumi.Input[str]] = None,
            cloudfront_access_identity_path: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            iam_arn: Optional[pulumi.Input[str]] = None,
            s3_canonical_user_id: Optional[pulumi.Input[str]] = None) -> 'OriginAccessIdentity':
        """
        Get an existing OriginAccessIdentity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] caller_reference: Internal value used by CloudFront to allow future
               updates to the origin access identity.
        :param pulumi.Input[str] cloudfront_access_identity_path: A shortcut to the full path for the
               origin access identity to use in CloudFront, see below.
        :param pulumi.Input[str] comment: An optional comment for the origin access identity.
        :param pulumi.Input[str] etag: The current version of the origin access identity's information.
               For example: `E2QWRUHAPOMQZL`.
        :param pulumi.Input[str] iam_arn: A pre-generated ARN for use in S3 bucket policies (see below).
               Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
               E2QWRUHAPOMQZL`.
        :param pulumi.Input[str] s3_canonical_user_id: The Amazon S3 canonical user ID for the origin
               access identity, which you use when giving the origin access identity read
               permission to an object in Amazon S3.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["caller_reference"] = caller_reference
        __props__["cloudfront_access_identity_path"] = cloudfront_access_identity_path
        __props__["comment"] = comment
        __props__["etag"] = etag
        __props__["iam_arn"] = iam_arn
        __props__["s3_canonical_user_id"] = s3_canonical_user_id
        return OriginAccessIdentity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="callerReference")
    def caller_reference(self) -> str:
        """
        Internal value used by CloudFront to allow future
        updates to the origin access identity.
        """
        return pulumi.get(self, "caller_reference")

    @property
    @pulumi.getter(name="cloudfrontAccessIdentityPath")
    def cloudfront_access_identity_path(self) -> str:
        """
        A shortcut to the full path for the
        origin access identity to use in CloudFront, see below.
        """
        return pulumi.get(self, "cloudfront_access_identity_path")

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        An optional comment for the origin access identity.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        The current version of the origin access identity's information.
        For example: `E2QWRUHAPOMQZL`.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="iamArn")
    def iam_arn(self) -> str:
        """
        A pre-generated ARN for use in S3 bucket policies (see below).
        Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
        E2QWRUHAPOMQZL`.
        """
        return pulumi.get(self, "iam_arn")

    @property
    @pulumi.getter(name="s3CanonicalUserId")
    def s3_canonical_user_id(self) -> str:
        """
        The Amazon S3 canonical user ID for the origin
        access identity, which you use when giving the origin access identity read
        permission to an object in Amazon S3.
        """
        return pulumi.get(self, "s3_canonical_user_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

