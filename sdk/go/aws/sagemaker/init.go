// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

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
	case "aws:sagemaker/app:App":
		r = &App{}
	case "aws:sagemaker/appImageConfig:AppImageConfig":
		r = &AppImageConfig{}
	case "aws:sagemaker/codeRepository:CodeRepository":
		r = &CodeRepository{}
	case "aws:sagemaker/domain:Domain":
		r = &Domain{}
	case "aws:sagemaker/endpoint:Endpoint":
		r = &Endpoint{}
	case "aws:sagemaker/endpointConfiguration:EndpointConfiguration":
		r = &EndpointConfiguration{}
	case "aws:sagemaker/featureGroup:FeatureGroup":
		r = &FeatureGroup{}
	case "aws:sagemaker/image:Image":
		r = &Image{}
	case "aws:sagemaker/imageVersion:ImageVersion":
		r = &ImageVersion{}
	case "aws:sagemaker/model:Model":
		r = &Model{}
	case "aws:sagemaker/modelPackageGroup:ModelPackageGroup":
		r = &ModelPackageGroup{}
	case "aws:sagemaker/notebookInstance:NotebookInstance":
		r = &NotebookInstance{}
	case "aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration":
		r = &NotebookInstanceLifecycleConfiguration{}
	case "aws:sagemaker/userProfile:UserProfile":
		r = &UserProfile{}
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
		"sagemaker/app",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/appImageConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/codeRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/endpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/endpointConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/featureGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/image",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/imageVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/model",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/modelPackageGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/notebookInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/notebookInstanceLifecycleConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/userProfile",
		&module{version},
	)
}
