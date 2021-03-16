# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SshKeyArgs', 'SshKey']

@pulumi.input_type
class SshKeyArgs:
    def __init__(__self__, *,
                 body: pulumi.Input[str],
                 server_id: pulumi.Input[str],
                 user_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a SshKey resource.
        :param pulumi.Input[str] body: The public key portion of an SSH key pair.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[str] user_name: The name of the user account that is assigned to one or more servers.
        """
        pulumi.set(__self__, "body", body)
        pulumi.set(__self__, "server_id", server_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def body(self) -> pulumi.Input[str]:
        """
        The public key portion of an SSH key pair.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: pulumi.Input[str]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The Server ID of the Transfer Server (e.g. `s-12345678`)
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The name of the user account that is assigned to one or more servers.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)


class SshKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a AWS Transfer User SSH Key resource.

        ```python
        import pulumi
        import pulumi_aws as aws

        example_server = aws.transfer.Server("exampleServer",
            identity_provider_type="SERVICE_MANAGED",
            tags={
                "NAME": "tf-acc-test-transfer-server",
            })
        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
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
        \"\"\")
        example_user = aws.transfer.User("exampleUser",
            server_id=example_server.id,
            user_name="tftestuser",
            role=example_role.arn,
            tags={
                "NAME": "tftestuser",
            })
        example_ssh_key = aws.transfer.SshKey("exampleSshKey",
            server_id=example_server.id,
            user_name=example_user.user_name,
            body="... SSH key ...")
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
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
        \"\"\")
        ```

        ## Import

        Transfer SSH Public Key can be imported using the `server_id` and `user_name` and `ssh_public_key_id` separated by `/`.

        ```sh
         $ pulumi import aws:transfer/sshKey:SshKey bar s-12345678/test-username/key-12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The public key portion of an SSH key pair.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[str] user_name: The name of the user account that is assigned to one or more servers.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SshKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a AWS Transfer User SSH Key resource.

        ```python
        import pulumi
        import pulumi_aws as aws

        example_server = aws.transfer.Server("exampleServer",
            identity_provider_type="SERVICE_MANAGED",
            tags={
                "NAME": "tf-acc-test-transfer-server",
            })
        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
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
        \"\"\")
        example_user = aws.transfer.User("exampleUser",
            server_id=example_server.id,
            user_name="tftestuser",
            role=example_role.arn,
            tags={
                "NAME": "tftestuser",
            })
        example_ssh_key = aws.transfer.SshKey("exampleSshKey",
            server_id=example_server.id,
            user_name=example_user.user_name,
            body="... SSH key ...")
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
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
        \"\"\")
        ```

        ## Import

        Transfer SSH Public Key can be imported using the `server_id` and `user_name` and `ssh_public_key_id` separated by `/`.

        ```sh
         $ pulumi import aws:transfer/sshKey:SshKey bar s-12345678/test-username/key-12345
        ```

        :param str resource_name: The name of the resource.
        :param SshKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SshKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
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

            if body is None and not opts.urn:
                raise TypeError("Missing required property 'body'")
            __props__['body'] = body
            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__['server_id'] = server_id
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__['user_name'] = user_name
        super(SshKey, __self__).__init__(
            'aws:transfer/sshKey:SshKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'SshKey':
        """
        Get an existing SshKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The public key portion of an SSH key pair.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[str] user_name: The name of the user account that is assigned to one or more servers.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["body"] = body
        __props__["server_id"] = server_id
        __props__["user_name"] = user_name
        return SshKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def body(self) -> pulumi.Output[str]:
        """
        The public key portion of an SSH key pair.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The Server ID of the Transfer Server (e.g. `s-12345678`)
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The name of the user account that is assigned to one or more servers.
        """
        return pulumi.get(self, "user_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

