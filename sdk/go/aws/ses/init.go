// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"fmt"

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
	case "aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet":
		r = &ActiveReceiptRuleSet{}
	case "aws:ses/confgurationSet:ConfgurationSet":
		r = &ConfgurationSet{}
	case "aws:ses/configurationSet:ConfigurationSet":
		r = &ConfigurationSet{}
	case "aws:ses/domainDkim:DomainDkim":
		r = &DomainDkim{}
	case "aws:ses/domainIdentity:DomainIdentity":
		r = &DomainIdentity{}
	case "aws:ses/domainIdentityVerification:DomainIdentityVerification":
		r = &DomainIdentityVerification{}
	case "aws:ses/emailIdentity:EmailIdentity":
		r = &EmailIdentity{}
	case "aws:ses/eventDestination:EventDestination":
		r = &EventDestination{}
	case "aws:ses/identityNotificationTopic:IdentityNotificationTopic":
		r = &IdentityNotificationTopic{}
	case "aws:ses/identityPolicy:IdentityPolicy":
		r = &IdentityPolicy{}
	case "aws:ses/mailFrom:MailFrom":
		r = &MailFrom{}
	case "aws:ses/receiptFilter:ReceiptFilter":
		r = &ReceiptFilter{}
	case "aws:ses/receiptRule:ReceiptRule":
		r = &ReceiptRule{}
	case "aws:ses/receiptRuleSet:ReceiptRuleSet":
		r = &ReceiptRuleSet{}
	case "aws:ses/template:Template":
		r = &Template{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version := aws.PkgVersion()
	pulumi.RegisterResourceModule(
		"aws",
		"ses/activeReceiptRuleSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/confgurationSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/configurationSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/domainDkim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/domainIdentity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/domainIdentityVerification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/emailIdentity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/eventDestination",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/identityNotificationTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/identityPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/mailFrom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/receiptFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/receiptRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/receiptRuleSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ses/template",
		&module{version},
	)
}
