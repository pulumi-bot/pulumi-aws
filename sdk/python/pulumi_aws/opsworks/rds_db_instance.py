# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RdsDbInstance(pulumi.CustomResource):
    db_password: pulumi.Output[str]
    """
    A db password
    """
    db_user: pulumi.Output[str]
    """
    A db username
    """
    rds_db_instance_arn: pulumi.Output[str]
    """
    The db instance to register for this stack. Changing this will force a new resource.
    """
    stack_id: pulumi.Output[str]
    """
    The stack to register a db instance for. Changing this will force a new resource.
    """
    def __init__(__self__, resource_name, opts=None, db_password=None, db_user=None, rds_db_instance_arn=None, stack_id=None, __name__=None, __opts__=None):
        """
        Provides an OpsWorks RDS DB Instance resource.
        
        > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_password: A db password
        :param pulumi.Input[str] db_user: A db username
        :param pulumi.Input[str] rds_db_instance_arn: The db instance to register for this stack. Changing this will force a new resource.
        :param pulumi.Input[str] stack_id: The stack to register a db instance for. Changing this will force a new resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_rds_db_instance.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if db_password is None:
            raise TypeError("Missing required property 'db_password'")
        __props__['db_password'] = db_password

        if db_user is None:
            raise TypeError("Missing required property 'db_user'")
        __props__['db_user'] = db_user

        if rds_db_instance_arn is None:
            raise TypeError("Missing required property 'rds_db_instance_arn'")
        __props__['rds_db_instance_arn'] = rds_db_instance_arn

        if stack_id is None:
            raise TypeError("Missing required property 'stack_id'")
        __props__['stack_id'] = stack_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RdsDbInstance, __self__).__init__(
            'aws:opsworks/rdsDbInstance:RdsDbInstance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

