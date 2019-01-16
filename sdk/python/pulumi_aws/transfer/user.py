# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class User(pulumi.CustomResource):
    """
    Provides a AWS Transfer User resource. Managing SSH keys can be accomplished with the [`aws_transfer_ssh_key` resource](https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html).
    
    
    ```hcl
    resource "aws_transfer_server" "foo" {
    	identity_provider_type = "SERVICE_MANAGED"
    
    	tags {
    		NAME     = "tf-acc-test-transfer-server"
    	}
    }
    
    resource "aws_iam_role" "foo" {
    	name = "tf-test-transfer-user-iam-role"
    
    	assume_role_policy = <<EOF
    {
    	"Version": "2012-10-17",
    	"Statement": [
    		{
    		"Effect": "Allow",
    		"Principal": {
    			"Service": "transfer.amazonaws.com"
    		},
    		"Action": "sts:AssumeRole"
    		}
    	]
    }
    EOF
    }
    
    resource "aws_iam_role_policy" "foo" {
    	name = "tf-test-transfer-user-iam-policy"
    	role = "${aws_iam_role.foo.id}"
    	policy = <<POLICY
    {
    	"Version": "2012-10-17",
    	"Statement": [
    		{
    			"Sid": "AllowFullAccesstoS3",
    			"Effect": "Allow",
    			"Action": [
    				"s3:*"
    			],
    			"Resource": "*"
    		}
    	]
    }
    POLICY
    }
    
    resource "aws_transfer_user" "foo" {
    	server_id      = "${aws_transfer_server.foo.id}"
    	user_name      = "tftestuser"
    	role           = "${aws_iam_role.foo.arn}"
    }
    
    ```
    """
    def __init__(__self__, __name__, __opts__=None, home_directory=None, policy=None, role=None, server_id=None, tags=None, user_name=None):
        """Create a User resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['home_directory'] = home_directory

        __props__['policy'] = policy

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        if not server_id:
            raise TypeError('Missing required property server_id')
        __props__['server_id'] = server_id

        __props__['tags'] = tags

        if not user_name:
            raise TypeError('Missing required property user_name')
        __props__['user_name'] = user_name

        __props__['arn'] = None

        super(User, __self__).__init__(
            'aws:transfer/user:User',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

