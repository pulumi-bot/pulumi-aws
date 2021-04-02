# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ServerArgs', 'Server']

@pulumi.input_type
class ServerArgs:
    def __init__(__self__, *,
                 endpoint_details: Optional[pulumi.Input['ServerEndpointDetailsArgs']] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 host_key: Optional[pulumi.Input[str]] = None,
                 identity_provider_type: Optional[pulumi.Input[str]] = None,
                 invocation_role: Optional[pulumi.Input[str]] = None,
                 logging_role: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Server resource.
        :param pulumi.Input['ServerEndpointDetailsArgs'] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] host_key: RSA private key (e.g. as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        """
        if endpoint_details is not None:
            pulumi.set(__self__, "endpoint_details", endpoint_details)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if host_key is not None:
            pulumi.set(__self__, "host_key", host_key)
        if identity_provider_type is not None:
            pulumi.set(__self__, "identity_provider_type", identity_provider_type)
        if invocation_role is not None:
            pulumi.set(__self__, "invocation_role", invocation_role)
        if logging_role is not None:
            pulumi.set(__self__, "logging_role", logging_role)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="endpointDetails")
    def endpoint_details(self) -> Optional[pulumi.Input['ServerEndpointDetailsArgs']]:
        """
        The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        """
        return pulumi.get(self, "endpoint_details")

    @endpoint_details.setter
    def endpoint_details(self, value: Optional[pulumi.Input['ServerEndpointDetailsArgs']]):
        pulumi.set(self, "endpoint_details", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter(name="hostKey")
    def host_key(self) -> Optional[pulumi.Input[str]]:
        """
        RSA private key (e.g. as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        """
        return pulumi.get(self, "host_key")

    @host_key.setter
    def host_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_key", value)

    @property
    @pulumi.getter(name="identityProviderType")
    def identity_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        """
        return pulumi.get(self, "identity_provider_type")

    @identity_provider_type.setter
    def identity_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_provider_type", value)

    @property
    @pulumi.getter(name="invocationRole")
    def invocation_role(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        """
        return pulumi.get(self, "invocation_role")

    @invocation_role.setter
    def invocation_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invocation_role", value)

    @property
    @pulumi.getter(name="loggingRole")
    def logging_role(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        """
        return pulumi.get(self, "logging_role")

    @logging_role.setter
    def logging_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logging_role", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Server(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_details: Optional[pulumi.Input[pulumi.InputType['ServerEndpointDetailsArgs']]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 host_key: Optional[pulumi.Input[str]] = None,
                 identity_provider_type: Optional[pulumi.Input[str]] = None,
                 invocation_role: Optional[pulumi.Input[str]] = None,
                 logging_role: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a AWS Transfer Server resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

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
        example_server = aws.transfer.Server("exampleServer",
            identity_provider_type="SERVICE_MANAGED",
            logging_role=example_role.arn,
            tags={
                "NAME": "tf-acc-test-transfer-server",
                "ENV": "test",
            })
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
        	"Version": "2012-10-17",
        	"Statement": [
        		{
        		"Sid": "AllowFullAccesstoCloudWatchLogs",
        		"Effect": "Allow",
        		"Action": [
        			"logs:*"
        		],
        		"Resource": "*"
        		}
        	]
        }
        \"\"\")
        ```

        ## Import

        Transfer Servers can be imported using the `server id`, e.g.

        ```sh
         $ pulumi import aws:transfer/server:Server bar s-12345678
        ```

         Certain resource arguments, such as `host_key`, cannot be read via the API and imported into the provider. This provider will display a difference for these arguments the first run after import if declared in the provider configuration for an imported resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServerEndpointDetailsArgs']] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] host_key: RSA private key (e.g. as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a AWS Transfer Server resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

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
        example_server = aws.transfer.Server("exampleServer",
            identity_provider_type="SERVICE_MANAGED",
            logging_role=example_role.arn,
            tags={
                "NAME": "tf-acc-test-transfer-server",
                "ENV": "test",
            })
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
        	"Version": "2012-10-17",
        	"Statement": [
        		{
        		"Sid": "AllowFullAccesstoCloudWatchLogs",
        		"Effect": "Allow",
        		"Action": [
        			"logs:*"
        		],
        		"Resource": "*"
        		}
        	]
        }
        \"\"\")
        ```

        ## Import

        Transfer Servers can be imported using the `server id`, e.g.

        ```sh
         $ pulumi import aws:transfer/server:Server bar s-12345678
        ```

         Certain resource arguments, such as `host_key`, cannot be read via the API and imported into the provider. This provider will display a difference for these arguments the first run after import if declared in the provider configuration for an imported resource.

        :param str resource_name: The name of the resource.
        :param ServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_details: Optional[pulumi.Input[pulumi.InputType['ServerEndpointDetailsArgs']]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 host_key: Optional[pulumi.Input[str]] = None,
                 identity_provider_type: Optional[pulumi.Input[str]] = None,
                 invocation_role: Optional[pulumi.Input[str]] = None,
                 logging_role: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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

            __props__['endpoint_details'] = endpoint_details
            __props__['endpoint_type'] = endpoint_type
            __props__['force_destroy'] = force_destroy
            __props__['host_key'] = host_key
            __props__['identity_provider_type'] = identity_provider_type
            __props__['invocation_role'] = invocation_role
            __props__['logging_role'] = logging_role
            __props__['tags'] = tags
            __props__['url'] = url
            __props__['arn'] = None
            __props__['endpoint'] = None
            __props__['host_key_fingerprint'] = None
        super(Server, __self__).__init__(
            'aws:transfer/server:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            endpoint_details: Optional[pulumi.Input[pulumi.InputType['ServerEndpointDetailsArgs']]] = None,
            endpoint_type: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            host_key: Optional[pulumi.Input[str]] = None,
            host_key_fingerprint: Optional[pulumi.Input[str]] = None,
            identity_provider_type: Optional[pulumi.Input[str]] = None,
            invocation_role: Optional[pulumi.Input[str]] = None,
            logging_role: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of Transfer Server
        :param pulumi.Input[str] endpoint: The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`)
        :param pulumi.Input[pulumi.InputType['ServerEndpointDetailsArgs']] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] host_key: RSA private key (e.g. as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        :param pulumi.Input[str] host_key_fingerprint: This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["endpoint"] = endpoint
        __props__["endpoint_details"] = endpoint_details
        __props__["endpoint_type"] = endpoint_type
        __props__["force_destroy"] = force_destroy
        __props__["host_key"] = host_key
        __props__["host_key_fingerprint"] = host_key_fingerprint
        __props__["identity_provider_type"] = identity_provider_type
        __props__["invocation_role"] = invocation_role
        __props__["logging_role"] = logging_role
        __props__["tags"] = tags
        __props__["url"] = url
        return Server(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of Transfer Server
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`)
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="endpointDetails")
    def endpoint_details(self) -> pulumi.Output[Optional['outputs.ServerEndpointDetails']]:
        """
        The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        """
        return pulumi.get(self, "endpoint_details")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter(name="hostKey")
    def host_key(self) -> pulumi.Output[Optional[str]]:
        """
        RSA private key (e.g. as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        """
        return pulumi.get(self, "host_key")

    @property
    @pulumi.getter(name="hostKeyFingerprint")
    def host_key_fingerprint(self) -> pulumi.Output[str]:
        """
        This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
        """
        return pulumi.get(self, "host_key_fingerprint")

    @property
    @pulumi.getter(name="identityProviderType")
    def identity_provider_type(self) -> pulumi.Output[Optional[str]]:
        """
        The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        """
        return pulumi.get(self, "identity_provider_type")

    @property
    @pulumi.getter(name="invocationRole")
    def invocation_role(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        """
        return pulumi.get(self, "invocation_role")

    @property
    @pulumi.getter(name="loggingRole")
    def logging_role(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        """
        return pulumi.get(self, "logging_role")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

