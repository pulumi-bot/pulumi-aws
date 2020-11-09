# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Bucket']


class Bucket(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceleration_status: Optional[pulumi.Input[str]] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 bucket_prefix: Optional[pulumi.Input[str]] = None,
                 cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsRuleArgs']]]]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 grants: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketGrantArgs']]]]] = None,
                 hosted_zone_id: Optional[pulumi.Input[str]] = None,
                 lifecycle_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]]] = None,
                 loggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]]]] = None,
                 object_lock_configuration: Optional[pulumi.Input[pulumi.InputType['BucketObjectLockConfigurationArgs']]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 replication_configuration: Optional[pulumi.Input[pulumi.InputType['BucketReplicationConfigurationArgs']]] = None,
                 request_payer: Optional[pulumi.Input[str]] = None,
                 server_side_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['BucketServerSideEncryptionConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 versioning: Optional[pulumi.Input[pulumi.InputType['BucketVersioningArgs']]] = None,
                 website: Optional[pulumi.Input[pulumi.InputType['BucketWebsiteArgs']]] = None,
                 website_domain: Optional[pulumi.Input[str]] = None,
                 website_endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a S3 bucket resource.

        > This functionality is for managing S3 in an AWS Partition. To manage [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html), see the [`s3control.Bucket` resource](https://www.terraform.io/docs/providers/aws/r/s3control_bucket.html).

        ## Example Usage
        ### Private Bucket w/ Tags

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket",
            acl="private",
            tags={
                "Environment": "Dev",
                "Name": "My bucket",
            })
        ```
        ### Static Website Hosting

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket",
            acl="public-read",
            policy=(lambda path: open(path).read())("policy.json"),
            website=aws.s3.BucketWebsiteArgs(
                index_document="index.html",
                error_document="error.html",
                routing_rules=\"\"\"[{
            "Condition": {
                "KeyPrefixEquals": "docs/"
            },
            "Redirect": {
                "ReplaceKeyPrefixWith": "documents/"
            }
        }]
        \"\"\",
            ))
        ```
        ### Using CORS

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket",
            acl="public-read",
            cors_rules=[aws.s3.BucketCorsRuleArgs(
                allowed_headers=["*"],
                allowed_methods=[
                    "PUT",
                    "POST",
                ],
                allowed_origins=["https://s3-website-test.mydomain.com"],
                expose_headers=["ETag"],
                max_age_seconds=3000,
            )])
        ```
        ### Using versioning

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket",
            acl="private",
            versioning=aws.s3.BucketVersioningArgs(
                enabled=True,
            ))
        ```
        ### Enable Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        log_bucket = aws.s3.Bucket("logBucket", acl="log-delivery-write")
        bucket = aws.s3.Bucket("bucket",
            acl="private",
            loggings=[aws.s3.BucketLoggingArgs(
                target_bucket=log_bucket.id,
                target_prefix="log/",
            )])
        ```
        ### Using object lifecycle

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket",
            acl="private",
            lifecycle_rules=[
                aws.s3.BucketLifecycleRuleArgs(
                    enabled=True,
                    expiration=aws.s3.BucketLifecycleRuleExpirationArgs(
                        days=90,
                    ),
                    id="log",
                    prefix="log/",
                    tags={
                        "autoclean": "true",
                        "rule": "log",
                    },
                    transitions=[
                        aws.s3.BucketLifecycleRuleTransitionArgs(
                            days=30,
                            storage_class="STANDARD_IA",
                        ),
                        aws.s3.BucketLifecycleRuleTransitionArgs(
                            days=60,
                            storage_class="GLACIER",
                        ),
                    ],
                ),
                aws.s3.BucketLifecycleRuleArgs(
                    enabled=True,
                    expiration=aws.s3.BucketLifecycleRuleExpirationArgs(
                        date="2016-01-12",
                    ),
                    id="tmp",
                    prefix="tmp/",
                ),
            ])
        versioning_bucket = aws.s3.Bucket("versioningBucket",
            acl="private",
            lifecycle_rules=[aws.s3.BucketLifecycleRuleArgs(
                enabled=True,
                noncurrent_version_expiration=aws.s3.BucketLifecycleRuleNoncurrentVersionExpirationArgs(
                    days=90,
                ),
                noncurrent_version_transitions=[
                    aws.s3.BucketLifecycleRuleNoncurrentVersionTransitionArgs(
                        days=30,
                        storage_class="STANDARD_IA",
                    ),
                    aws.s3.BucketLifecycleRuleNoncurrentVersionTransitionArgs(
                        days=60,
                        storage_class="GLACIER",
                    ),
                ],
                prefix="config/",
            )],
            versioning=aws.s3.BucketVersioningArgs(
                enabled=True,
            ))
        ```
        ### Using replication configuration

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        central = pulumi.providers.Aws("central", region="eu-central-1")
        replication_role = aws.iam.Role("replicationRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "s3.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        destination = aws.s3.Bucket("destination", versioning=aws.s3.BucketVersioningArgs(
            enabled=True,
        ))
        bucket = aws.s3.Bucket("bucket",
            acl="private",
            versioning=aws.s3.BucketVersioningArgs(
                enabled=True,
            ),
            replication_configuration=aws.s3.BucketReplicationConfigurationArgs(
                role=replication_role.arn,
                rules=[aws.s3.BucketReplicationConfigurationRuleArgs(
                    id="foobar",
                    prefix="foo",
                    status="Enabled",
                    destination=aws.s3.BucketReplicationConfigurationRuleDestinationArgs(
                        bucket=destination.arn,
                        storage_class="STANDARD",
                    ),
                )],
            ),
            opts=ResourceOptions(provider=aws["central"]))
        replication_policy = aws.iam.Policy("replicationPolicy", policy=pulumi.Output.all(bucket.arn, bucket.arn, destination.arn).apply(lambda bucketArn, bucketArn1, destinationArn: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Action": [
                "s3:GetReplicationConfiguration",
                "s3:ListBucket"
              ],
              "Effect": "Allow",
              "Resource": [
                "{bucket_arn}"
              ]
            }},
            {{
              "Action": [
                "s3:GetObjectVersion",
                "s3:GetObjectVersionAcl"
              ],
              "Effect": "Allow",
              "Resource": [
                "{bucket_arn1}/*"
              ]
            }},
            {{
              "Action": [
                "s3:ReplicateObject",
                "s3:ReplicateDelete"
              ],
              "Effect": "Allow",
              "Resource": "{destination_arn}/*"
            }}
          ]
        }}
        \"\"\"))
        replication_role_policy_attachment = aws.iam.RolePolicyAttachment("replicationRolePolicyAttachment",
            role=replication_role.name,
            policy_arn=replication_policy.arn)
        ```
        ### Enable Default Server Side Encryption

        ```python
        import pulumi
        import pulumi_aws as aws

        mykey = aws.kms.Key("mykey",
            description="This key is used to encrypt bucket objects",
            deletion_window_in_days=10)
        mybucket = aws.s3.Bucket("mybucket", server_side_encryption_configuration=aws.s3.BucketServerSideEncryptionConfigurationArgs(
            rule=aws.s3.BucketServerSideEncryptionConfigurationRuleArgs(
                apply_server_side_encryption_by_default=aws.s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs(
                    kms_master_key_id=mykey.arn,
                    sse_algorithm="aws:kms",
                ),
            ),
        ))
        ```
        ### Using ACL policy grants

        ```python
        import pulumi
        import pulumi_aws as aws

        current_user = aws.get_canonical_user_id()
        bucket = aws.s3.Bucket("bucket", grants=[
            aws.s3.BucketGrantArgs(
                id=current_user.id,
                type="CanonicalUser",
                permissions=["FULL_CONTROL"],
            ),
            aws.s3.BucketGrantArgs(
                type="Group",
                permissions=[
                    "READ",
                    "WRITE",
                ],
                uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
            ),
        ])
        ```

        ## Import

        S3 bucket can be imported using the `bucket`, e.g.

        ```sh
         $ pulumi import aws:s3/bucket:Bucket bucket bucket-name
        ```

         The `policy` argument is not imported and will be deprecated in a future version 3.x of the Terraform AWS Provider for removal in version 4.0. Use the [`aws_s3_bucket_policy` resource](/docs/providers/aws/r/s3_bucket_policy.html) to manage the S3 Bucket Policy instead.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acceleration_status: Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
        :param pulumi.Input[str] acl: The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
        :param pulumi.Input[str] arn: The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        :param pulumi.Input[str] bucket: The name of the bucket. If omitted, this provider will assign a random, unique name. Must be less than or equal to 63 characters in length.
        :param pulumi.Input[str] bucket_prefix: Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be less than or equal to 37 characters in length.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsRuleArgs']]]] cors_rules: A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketGrantArgs']]]] grants: An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
        :param pulumi.Input[str] hosted_zone_id: The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]] lifecycle_rules: A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]]] loggings: A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
        :param pulumi.Input[pulumi.InputType['BucketObjectLockConfigurationArgs']] object_lock_configuration: A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
        :param pulumi.Input[str] policy: A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
        :param pulumi.Input[pulumi.InputType['BucketReplicationConfigurationArgs']] replication_configuration: A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
        :param pulumi.Input[str] request_payer: Specifies who should bear the cost of Amazon S3 data transfer.
               Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
               the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
               developer guide for more information.
        :param pulumi.Input[pulumi.InputType['BucketServerSideEncryptionConfigurationArgs']] server_side_encryption_configuration: A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the bucket.
        :param pulumi.Input[pulumi.InputType['BucketVersioningArgs']] versioning: A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        :param pulumi.Input[pulumi.InputType['BucketWebsiteArgs']] website: A website object (documented below).
        :param pulumi.Input[str] website_domain: The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        :param pulumi.Input[str] website_endpoint: The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
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

            __props__['acceleration_status'] = acceleration_status
            __props__['acl'] = acl
            __props__['arn'] = arn
            __props__['bucket'] = bucket
            __props__['bucket_prefix'] = bucket_prefix
            __props__['cors_rules'] = cors_rules
            __props__['force_destroy'] = force_destroy
            __props__['grants'] = grants
            __props__['hosted_zone_id'] = hosted_zone_id
            __props__['lifecycle_rules'] = lifecycle_rules
            __props__['loggings'] = loggings
            __props__['object_lock_configuration'] = object_lock_configuration
            __props__['policy'] = policy
            __props__['replication_configuration'] = replication_configuration
            __props__['request_payer'] = request_payer
            __props__['server_side_encryption_configuration'] = server_side_encryption_configuration
            __props__['tags'] = tags
            __props__['versioning'] = versioning
            __props__['website'] = website
            __props__['website_domain'] = website_domain
            __props__['website_endpoint'] = website_endpoint
            __props__['bucket_domain_name'] = None
            __props__['bucket_regional_domain_name'] = None
            __props__['region'] = None
        super(Bucket, __self__).__init__(
            'aws:s3/bucket:Bucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acceleration_status: Optional[pulumi.Input[str]] = None,
            acl: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            bucket_domain_name: Optional[pulumi.Input[str]] = None,
            bucket_prefix: Optional[pulumi.Input[str]] = None,
            bucket_regional_domain_name: Optional[pulumi.Input[str]] = None,
            cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsRuleArgs']]]]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            grants: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketGrantArgs']]]]] = None,
            hosted_zone_id: Optional[pulumi.Input[str]] = None,
            lifecycle_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]]] = None,
            loggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]]]] = None,
            object_lock_configuration: Optional[pulumi.Input[pulumi.InputType['BucketObjectLockConfigurationArgs']]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            replication_configuration: Optional[pulumi.Input[pulumi.InputType['BucketReplicationConfigurationArgs']]] = None,
            request_payer: Optional[pulumi.Input[str]] = None,
            server_side_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['BucketServerSideEncryptionConfigurationArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            versioning: Optional[pulumi.Input[pulumi.InputType['BucketVersioningArgs']]] = None,
            website: Optional[pulumi.Input[pulumi.InputType['BucketWebsiteArgs']]] = None,
            website_domain: Optional[pulumi.Input[str]] = None,
            website_endpoint: Optional[pulumi.Input[str]] = None) -> 'Bucket':
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acceleration_status: Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
        :param pulumi.Input[str] acl: The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
        :param pulumi.Input[str] arn: The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        :param pulumi.Input[str] bucket: The name of the bucket. If omitted, this provider will assign a random, unique name. Must be less than or equal to 63 characters in length.
        :param pulumi.Input[str] bucket_domain_name: The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
        :param pulumi.Input[str] bucket_prefix: Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be less than or equal to 37 characters in length.
        :param pulumi.Input[str] bucket_regional_domain_name: The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsRuleArgs']]]] cors_rules: A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketGrantArgs']]]] grants: An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
        :param pulumi.Input[str] hosted_zone_id: The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]] lifecycle_rules: A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]]] loggings: A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
        :param pulumi.Input[pulumi.InputType['BucketObjectLockConfigurationArgs']] object_lock_configuration: A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
        :param pulumi.Input[str] policy: A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
        :param pulumi.Input[str] region: The AWS region this bucket resides in.
        :param pulumi.Input[pulumi.InputType['BucketReplicationConfigurationArgs']] replication_configuration: A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
        :param pulumi.Input[str] request_payer: Specifies who should bear the cost of Amazon S3 data transfer.
               Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
               the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
               developer guide for more information.
        :param pulumi.Input[pulumi.InputType['BucketServerSideEncryptionConfigurationArgs']] server_side_encryption_configuration: A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the bucket.
        :param pulumi.Input[pulumi.InputType['BucketVersioningArgs']] versioning: A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        :param pulumi.Input[pulumi.InputType['BucketWebsiteArgs']] website: A website object (documented below).
        :param pulumi.Input[str] website_domain: The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        :param pulumi.Input[str] website_endpoint: The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acceleration_status"] = acceleration_status
        __props__["acl"] = acl
        __props__["arn"] = arn
        __props__["bucket"] = bucket
        __props__["bucket_domain_name"] = bucket_domain_name
        __props__["bucket_prefix"] = bucket_prefix
        __props__["bucket_regional_domain_name"] = bucket_regional_domain_name
        __props__["cors_rules"] = cors_rules
        __props__["force_destroy"] = force_destroy
        __props__["grants"] = grants
        __props__["hosted_zone_id"] = hosted_zone_id
        __props__["lifecycle_rules"] = lifecycle_rules
        __props__["loggings"] = loggings
        __props__["object_lock_configuration"] = object_lock_configuration
        __props__["policy"] = policy
        __props__["region"] = region
        __props__["replication_configuration"] = replication_configuration
        __props__["request_payer"] = request_payer
        __props__["server_side_encryption_configuration"] = server_side_encryption_configuration
        __props__["tags"] = tags
        __props__["versioning"] = versioning
        __props__["website"] = website
        __props__["website_domain"] = website_domain
        __props__["website_endpoint"] = website_endpoint
        return Bucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accelerationStatus")
    def acceleration_status(self) -> pulumi.Output[str]:
        """
        Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
        """
        return pulumi.get(self, "acceleration_status")

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output[Optional[str]]:
        """
        The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket. If omitted, this provider will assign a random, unique name. Must be less than or equal to 63 characters in length.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="bucketDomainName")
    def bucket_domain_name(self) -> pulumi.Output[str]:
        """
        The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
        """
        return pulumi.get(self, "bucket_domain_name")

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be less than or equal to 37 characters in length.
        """
        return pulumi.get(self, "bucket_prefix")

    @property
    @pulumi.getter(name="bucketRegionalDomainName")
    def bucket_regional_domain_name(self) -> pulumi.Output[str]:
        """
        The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
        """
        return pulumi.get(self, "bucket_regional_domain_name")

    @property
    @pulumi.getter(name="corsRules")
    def cors_rules(self) -> pulumi.Output[Optional[Sequence['outputs.BucketCorsRule']]]:
        """
        A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        """
        return pulumi.get(self, "cors_rules")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def grants(self) -> pulumi.Output[Optional[Sequence['outputs.BucketGrant']]]:
        """
        An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
        """
        return pulumi.get(self, "grants")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Output[str]:
        """
        The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        """
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="lifecycleRules")
    def lifecycle_rules(self) -> pulumi.Output[Optional[Sequence['outputs.BucketLifecycleRule']]]:
        """
        A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
        """
        return pulumi.get(self, "lifecycle_rules")

    @property
    @pulumi.getter
    def loggings(self) -> pulumi.Output[Optional[Sequence['outputs.BucketLogging']]]:
        """
        A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
        """
        return pulumi.get(self, "loggings")

    @property
    @pulumi.getter(name="objectLockConfiguration")
    def object_lock_configuration(self) -> pulumi.Output[Optional['outputs.BucketObjectLockConfiguration']]:
        """
        A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
        """
        return pulumi.get(self, "object_lock_configuration")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Optional[str]]:
        """
        A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The AWS region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="replicationConfiguration")
    def replication_configuration(self) -> pulumi.Output[Optional['outputs.BucketReplicationConfiguration']]:
        """
        A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
        """
        return pulumi.get(self, "replication_configuration")

    @property
    @pulumi.getter(name="requestPayer")
    def request_payer(self) -> pulumi.Output[str]:
        """
        Specifies who should bear the cost of Amazon S3 data transfer.
        Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
        the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
        developer guide for more information.
        """
        return pulumi.get(self, "request_payer")

    @property
    @pulumi.getter(name="serverSideEncryptionConfiguration")
    def server_side_encryption_configuration(self) -> pulumi.Output[Optional['outputs.BucketServerSideEncryptionConfiguration']]:
        """
        A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
        """
        return pulumi.get(self, "server_side_encryption_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the bucket.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def versioning(self) -> pulumi.Output['outputs.BucketVersioning']:
        """
        A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        """
        return pulumi.get(self, "versioning")

    @property
    @pulumi.getter
    def website(self) -> pulumi.Output[Optional['outputs.BucketWebsite']]:
        """
        A website object (documented below).
        """
        return pulumi.get(self, "website")

    @property
    @pulumi.getter(name="websiteDomain")
    def website_domain(self) -> pulumi.Output[str]:
        """
        The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        """
        return pulumi.get(self, "website_domain")

    @property
    @pulumi.getter(name="websiteEndpoint")
    def website_endpoint(self) -> pulumi.Output[str]:
        """
        The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
        """
        return pulumi.get(self, "website_endpoint")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

