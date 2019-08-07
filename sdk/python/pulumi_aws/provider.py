# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, access_key=None, allowed_account_ids=None, assume_role=None, endpoints=None, forbidden_account_ids=None, insecure=None, max_retries=None, profile=None, region=None, s3_force_path_style=None, secret_key=None, shared_credentials_file=None, skip_credentials_validation=None, skip_get_ec2_platforms=None, skip_metadata_api_check=None, skip_region_validation=None, skip_requesting_account_id=None, token=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the aws package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/index.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['access_key'] = access_key
            __props__['allowed_account_ids'] = pulumi.Output.from_input(allowed_account_ids).apply(json.dumps) if allowed_account_ids is not None else None
            __props__['assume_role'] = pulumi.Output.from_input(assume_role).apply(json.dumps) if assume_role is not None else None
            __props__['endpoints'] = pulumi.Output.from_input(endpoints).apply(json.dumps) if endpoints is not None else None
            __props__['forbidden_account_ids'] = pulumi.Output.from_input(forbidden_account_ids).apply(json.dumps) if forbidden_account_ids is not None else None
            __props__['insecure'] = pulumi.Output.from_input(insecure).apply(json.dumps) if insecure is not None else None
            __props__['max_retries'] = pulumi.Output.from_input(max_retries).apply(json.dumps) if max_retries is not None else None
            if profile is None:
                profile = utilities.get_env('AWS_PROFILE')
            __props__['profile'] = profile
            if region is None:
                region = utilities.get_env('AWS_REGION', 'AWS_DEFAULT_REGION')
            __props__['region'] = region
            __props__['s3_force_path_style'] = pulumi.Output.from_input(s3_force_path_style).apply(json.dumps) if s3_force_path_style is not None else None
            __props__['secret_key'] = secret_key
            __props__['shared_credentials_file'] = shared_credentials_file
            __props__['skip_credentials_validation'] = pulumi.Output.from_input(skip_credentials_validation).apply(json.dumps) if skip_credentials_validation is not None else None
            __props__['skip_get_ec2_platforms'] = pulumi.Output.from_input(skip_get_ec2_platforms).apply(json.dumps) if skip_get_ec2_platforms is not None else None
            __props__['skip_metadata_api_check'] = pulumi.Output.from_input(skip_metadata_api_check).apply(json.dumps) if skip_metadata_api_check is not None else None
            __props__['skip_region_validation'] = pulumi.Output.from_input(skip_region_validation).apply(json.dumps) if skip_region_validation is not None else None
            __props__['skip_requesting_account_id'] = pulumi.Output.from_input(skip_requesting_account_id).apply(json.dumps) if skip_requesting_account_id is not None else None
            __props__['token'] = token
        super(Provider, __self__).__init__(
            'aws',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Provider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/index.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        return Provider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

