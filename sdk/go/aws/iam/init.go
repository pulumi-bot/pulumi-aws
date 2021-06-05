// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:iam/accessKey:AccessKey":
		r = &AccessKey{}
	case "aws:iam/accountAlias:AccountAlias":
		r = &AccountAlias{}
	case "aws:iam/accountPasswordPolicy:AccountPasswordPolicy":
		r = &AccountPasswordPolicy{}
	case "aws:iam/group:Group":
		r = &Group{}
	case "aws:iam/groupMembership:GroupMembership":
		r = &GroupMembership{}
	case "aws:iam/groupPolicy:GroupPolicy":
		r = &GroupPolicy{}
	case "aws:iam/groupPolicyAttachment:GroupPolicyAttachment":
		r = &GroupPolicyAttachment{}
	case "aws:iam/instanceProfile:InstanceProfile":
		r = &InstanceProfile{}
	case "aws:iam/openIdConnectProvider:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "aws:iam/policy:Policy":
		r = &Policy{}
	case "aws:iam/policyAttachment:PolicyAttachment":
		r = &PolicyAttachment{}
	case "aws:iam/role:Role":
		r = &Role{}
	case "aws:iam/rolePolicy:RolePolicy":
		r = &RolePolicy{}
	case "aws:iam/rolePolicyAttachment:RolePolicyAttachment":
		r = &RolePolicyAttachment{}
	case "aws:iam/samlProvider:SamlProvider":
		r = &SamlProvider{}
	case "aws:iam/serverCertificate:ServerCertificate":
		r = &ServerCertificate{}
	case "aws:iam/serviceLinkedRole:ServiceLinkedRole":
		r = &ServiceLinkedRole{}
	case "aws:iam/sshKey:SshKey":
		r = &SshKey{}
	case "aws:iam/user:User":
		r = &User{}
	case "aws:iam/userGroupMembership:UserGroupMembership":
		r = &UserGroupMembership{}
	case "aws:iam/userLoginProfile:UserLoginProfile":
		r = &UserLoginProfile{}
	case "aws:iam/userPolicy:UserPolicy":
		r = &UserPolicy{}
	case "aws:iam/userPolicyAttachment:UserPolicyAttachment":
		r = &UserPolicyAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accessKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accountAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accountPasswordPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/instanceProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/openIdConnectProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/policyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/rolePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/rolePolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/samlProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/serverCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/serviceLinkedRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/sshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userGroupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userLoginProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userPolicyAttachment",
		&module{version},
	)
}
