// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:wafregional/byteMatchSet:ByteMatchSet":
		r, err = NewByteMatchSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/geoMatchSet:GeoMatchSet":
		r, err = NewGeoMatchSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/ipSet:IpSet":
		r, err = NewIpSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/rateBasedRule:RateBasedRule":
		r, err = NewRateBasedRule(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/regexMatchSet:RegexMatchSet":
		r, err = NewRegexMatchSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/regexPatternSet:RegexPatternSet":
		r, err = NewRegexPatternSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/rule:Rule":
		r, err = NewRule(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/ruleGroup:RuleGroup":
		r, err = NewRuleGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/sizeConstraintSet:SizeConstraintSet":
		r, err = NewSizeConstraintSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet":
		r, err = NewSqlInjectionMatchSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/webAcl:WebAcl":
		r, err = NewWebAcl(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/webAclAssociation:WebAclAssociation":
		r, err = NewWebAclAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafregional/xssMatchSet:XssMatchSet":
		r, err = NewXssMatchSet(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/byteMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/geoMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/ipSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/rateBasedRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/regexMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/regexPatternSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/ruleGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/sizeConstraintSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/sqlInjectionMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/webAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/webAclAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/xssMatchSet",
		&module{version},
	)
}
