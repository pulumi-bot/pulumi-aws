# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

__config__ = pulumi.Config('aws')

access_key = __config__.get('accessKey')
"""
The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
"""

allowed_account_ids = __config__.get('allowedAccountIds')

assume_role = __config__.get('assumeRole')

dynamodb_endpoint = __config__.get('dynamodbEndpoint')
"""
Use this to override the default endpoint URL constructed from the `region`. It's typically used to connect to
dynamodb-local.
"""

endpoints = __config__.get('endpoints')

forbidden_account_ids = __config__.get('forbiddenAccountIds')

insecure = __config__.get('insecure')
"""
Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
"""

kinesis_endpoint = __config__.get('kinesisEndpoint')
"""
Use this to override the default endpoint URL constructed from the `region`. It's typically used to connect to
kinesalite.
"""

max_retries = __config__.get('maxRetries')
"""
The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
"""

profile = __config__.get('profile')
"""
The profile for API operations. If not set, the default profile created with `aws configure` will be used.
"""

region = utilities.require_with_default(lambda: __config__.require('region'), utilities.get_env('AWS_REGION', 'AWS_DEFAULT_REGION'))
"""
The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
"""

s3_force_path_style = __config__.get('s3ForcePathStyle')
"""
Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By
default, the S3 client will use virtual hosted bucket addressing when possible (http://BUCKET.s3.amazonaws.com/KEY).
Specific to the Amazon S3 service.
"""

secret_key = __config__.get('secretKey')
"""
The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
"""

shared_credentials_file = __config__.get('sharedCredentialsFile')
"""
The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
"""

skip_credentials_validation = __config__.get('skipCredentialsValidation')
"""
Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
available/implemented.
"""

skip_get_ec2_platforms = __config__.get('skipGetEc2Platforms')
"""
Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
"""

skip_metadata_api_check = __config__.get('skipMetadataApiCheck')

skip_region_validation = __config__.get('skipRegionValidation')
"""
Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
not public (yet).
"""

skip_requesting_account_id = __config__.get('skipRequestingAccountId')
"""
Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
"""

token = __config__.get('token')
"""
session token. A session token is only required if you are using temporary security credentials.
"""

