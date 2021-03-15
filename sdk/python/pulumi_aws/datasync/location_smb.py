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

__all__ = ['LocationSmbArgs', 'LocationSmb']

@pulumi.input_type
class LocationSmbArgs:
    def __init__(__self__, *,
                 agent_arns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 password: pulumi.Input[str],
                 server_hostname: pulumi.Input[str],
                 subdirectory: pulumi.Input[str],
                 user: pulumi.Input[str],
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input['LocationSmbMountOptionsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LocationSmb resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: A list of DataSync Agent ARNs with which this location will be associated.
        :param pulumi.Input[str] password: The password of the user who can mount the share and has file permissions in the SMB.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[str] user: The user who can mount the share and has file and folder permissions in the SMB share.
        :param pulumi.Input[str] domain: The name of the Windows domain the SMB server belongs to.
        :param pulumi.Input['LocationSmbMountOptionsArgs'] mount_options: Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        """
        pulumi.set(__self__, "agent_arns", agent_arns)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "server_hostname", server_hostname)
        pulumi.set(__self__, "subdirectory", subdirectory)
        pulumi.set(__self__, "user", user)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if mount_options is not None:
            pulumi.set(__self__, "mount_options", mount_options)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of DataSync Agent ARNs with which this location will be associated.
        """
        return pulumi.get(self, "agent_arns")

    @agent_arns.setter
    def agent_arns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "agent_arns", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password of the user who can mount the share and has file permissions in the SMB.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> pulumi.Input[str]:
        """
        Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        """
        return pulumi.get(self, "server_hostname")

    @server_hostname.setter
    def server_hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_hostname", value)

    @property
    @pulumi.getter
    def subdirectory(self) -> pulumi.Input[str]:
        """
        Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        """
        return pulumi.get(self, "subdirectory")

    @subdirectory.setter
    def subdirectory(self, value: pulumi.Input[str]):
        pulumi.set(self, "subdirectory", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The user who can mount the share and has file and folder permissions in the SMB share.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Windows domain the SMB server belongs to.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> Optional[pulumi.Input['LocationSmbMountOptionsArgs']]:
        """
        Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        """
        return pulumi.get(self, "mount_options")

    @mount_options.setter
    def mount_options(self, value: Optional[pulumi.Input['LocationSmbMountOptionsArgs']]):
        pulumi.set(self, "mount_options", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value pairs of resource tags to assign to the DataSync Location.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class LocationSmb(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocationSmbArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a SMB Location within AWS DataSync.

        > **NOTE:** The DataSync Agents must be available before creating this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.datasync.LocationSmb("example",
            server_hostname="smb.example.com",
            subdirectory="/exported/path",
            user="Guest",
            password="ANotGreatPassword",
            agent_arns=[aws_datasync_agent["example"]["arn"]])
        ```

        ## Import

        `aws_datasync_location_smb` can be imported by using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:datasync/locationSmb:LocationSmb example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
        ```

        :param str resource_name: The name of the resource.
        :param LocationSmbArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input[pulumi.InputType['LocationSmbMountOptionsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a SMB Location within AWS DataSync.

        > **NOTE:** The DataSync Agents must be available before creating this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.datasync.LocationSmb("example",
            server_hostname="smb.example.com",
            subdirectory="/exported/path",
            user="Guest",
            password="ANotGreatPassword",
            agent_arns=[aws_datasync_agent["example"]["arn"]])
        ```

        ## Import

        `aws_datasync_location_smb` can be imported by using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:datasync/locationSmb:LocationSmb example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: A list of DataSync Agent ARNs with which this location will be associated.
        :param pulumi.Input[str] domain: The name of the Windows domain the SMB server belongs to.
        :param pulumi.Input[pulumi.InputType['LocationSmbMountOptionsArgs']] mount_options: Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        :param pulumi.Input[str] password: The password of the user who can mount the share and has file permissions in the SMB.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        :param pulumi.Input[str] user: The user who can mount the share and has file and folder permissions in the SMB share.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocationSmbArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input[pulumi.InputType['LocationSmbMountOptionsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
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

            if agent_arns is None and not opts.urn:
                raise TypeError("Missing required property 'agent_arns'")
            __props__['agent_arns'] = agent_arns
            __props__['domain'] = domain
            __props__['mount_options'] = mount_options
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            if server_hostname is None and not opts.urn:
                raise TypeError("Missing required property 'server_hostname'")
            __props__['server_hostname'] = server_hostname
            if subdirectory is None and not opts.urn:
                raise TypeError("Missing required property 'subdirectory'")
            __props__['subdirectory'] = subdirectory
            __props__['tags'] = tags
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            __props__['arn'] = None
            __props__['uri'] = None
        super(LocationSmb, __self__).__init__(
            'aws:datasync/locationSmb:LocationSmb',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            mount_options: Optional[pulumi.Input[pulumi.InputType['LocationSmbMountOptionsArgs']]] = None,
            password: Optional[pulumi.Input[str]] = None,
            server_hostname: Optional[pulumi.Input[str]] = None,
            subdirectory: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'LocationSmb':
        """
        Get an existing LocationSmb resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: A list of DataSync Agent ARNs with which this location will be associated.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DataSync Location.
        :param pulumi.Input[str] domain: The name of the Windows domain the SMB server belongs to.
        :param pulumi.Input[pulumi.InputType['LocationSmbMountOptionsArgs']] mount_options: Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        :param pulumi.Input[str] password: The password of the user who can mount the share and has file permissions in the SMB.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        :param pulumi.Input[str] user: The user who can mount the share and has file and folder permissions in the SMB share.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["agent_arns"] = agent_arns
        __props__["arn"] = arn
        __props__["domain"] = domain
        __props__["mount_options"] = mount_options
        __props__["password"] = password
        __props__["server_hostname"] = server_hostname
        __props__["subdirectory"] = subdirectory
        __props__["tags"] = tags
        __props__["uri"] = uri
        __props__["user"] = user
        return LocationSmb(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of DataSync Agent ARNs with which this location will be associated.
        """
        return pulumi.get(self, "agent_arns")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the DataSync Location.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The name of the Windows domain the SMB server belongs to.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> pulumi.Output[Optional['outputs.LocationSmbMountOptions']]:
        """
        Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        """
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of the user who can mount the share and has file permissions in the SMB.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> pulumi.Output[str]:
        """
        Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        """
        return pulumi.get(self, "server_hostname")

    @property
    @pulumi.getter
    def subdirectory(self) -> pulumi.Output[str]:
        """
        Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        """
        return pulumi.get(self, "subdirectory")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value pairs of resource tags to assign to the DataSync Location.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The user who can mount the share and has file and folder permissions in the SMB share.
        """
        return pulumi.get(self, "user")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

