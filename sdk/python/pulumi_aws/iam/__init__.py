# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
# Export this package's modules as members:
from ._enums import *
from .access_key import *
from .account_alias import *
from .account_password_policy import *
from .get_account_alias import *
from .get_group import *
from .get_instance_profile import *
from .get_policy import *
from .get_policy_document import *
from .get_role import *
from .get_server_certificate import *
from .get_user import *
from .group import *
from .group_membership import *
from .group_policy import *
from .group_policy_attachment import *
from .instance_profile import *
from .open_id_connect_provider import *
from .policy import *
from .policy_attachment import *
from .role import *
from .role_policy import *
from .role_policy_attachment import *
from .saml_provider import *
from .server_certificate import *
from .service_linked_role import *
from .ssh_key import *
from .user import *
from .user_group_membership import *
from .user_login_profile import *
from .user_policy import *
from .user_policy_attachment import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:iam/accessKey:AccessKey":
                return AccessKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/accountAlias:AccountAlias":
                return AccountAlias(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/accountPasswordPolicy:AccountPasswordPolicy":
                return AccountPasswordPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/group:Group":
                return Group(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/groupMembership:GroupMembership":
                return GroupMembership(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/groupPolicy:GroupPolicy":
                return GroupPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/groupPolicyAttachment:GroupPolicyAttachment":
                return GroupPolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/instanceProfile:InstanceProfile":
                return InstanceProfile(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/openIdConnectProvider:OpenIdConnectProvider":
                return OpenIdConnectProvider(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/policy:Policy":
                return Policy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/policyAttachment:PolicyAttachment":
                return PolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/role:Role":
                return Role(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/rolePolicy:RolePolicy":
                return RolePolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/rolePolicyAttachment:RolePolicyAttachment":
                return RolePolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/samlProvider:SamlProvider":
                return SamlProvider(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/serverCertificate:ServerCertificate":
                return ServerCertificate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/serviceLinkedRole:ServiceLinkedRole":
                return ServiceLinkedRole(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/sshKey:SshKey":
                return SshKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/user:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/userGroupMembership:UserGroupMembership":
                return UserGroupMembership(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/userLoginProfile:UserLoginProfile":
                return UserLoginProfile(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/userPolicy:UserPolicy":
                return UserPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:iam/userPolicyAttachment:UserPolicyAttachment":
                return UserPolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "iam/accessKey", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/accountAlias", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/accountPasswordPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/group", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/groupMembership", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/groupPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/groupPolicyAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/instanceProfile", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/openIdConnectProvider", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/policy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/policyAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/role", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/rolePolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/rolePolicyAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/samlProvider", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/serverCertificate", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/serviceLinkedRole", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/sshKey", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/user", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/userGroupMembership", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/userLoginProfile", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/userPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "iam/userPolicyAttachment", _module_instance)

_register_module()
