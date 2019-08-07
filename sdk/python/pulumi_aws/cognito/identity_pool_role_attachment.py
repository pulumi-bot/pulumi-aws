# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IdentityPoolRoleAttachment(pulumi.CustomResource):
    identity_pool_id: pulumi.Output[str]
    """
    An identity pool ID in the format REGION:GUID.
    """
    role_mappings: pulumi.Output[list]
    """
    A List of Role Mapping.
    """
    roles: pulumi.Output[dict]
    """
    The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
    """
    def __init__(__self__, resource_name, opts=None, identity_pool_id=None, role_mappings=None, roles=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Cognito Identity Pool Roles Attachment.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_pool_id: An identity pool ID in the format REGION:GUID.
        :param pulumi.Input[list] role_mappings: A List of Role Mapping.
        :param pulumi.Input[dict] roles: The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_identity_pool_roles_attachment.html.markdown.
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

            if identity_pool_id is None:
                raise TypeError("Missing required property 'identity_pool_id'")
            __props__['identity_pool_id'] = identity_pool_id
            __props__['role_mappings'] = role_mappings
            if roles is None:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
        super(IdentityPoolRoleAttachment, __self__).__init__(
            'aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, identity_pool_id=None, role_mappings=None, roles=None):
        """
        Get an existing IdentityPoolRoleAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_pool_id: An identity pool ID in the format REGION:GUID.
        :param pulumi.Input[list] role_mappings: A List of Role Mapping.
        :param pulumi.Input[dict] roles: The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_identity_pool_roles_attachment.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["identity_pool_id"] = identity_pool_id
        __props__["role_mappings"] = role_mappings
        __props__["roles"] = roles
        return IdentityPoolRoleAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

