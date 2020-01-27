// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eks

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get an authentication token to communicate with an EKS cluster.
// 
// Uses IAM credentials from the AWS provider to generate a temporary token that is compatible with
// [AWS IAM Authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator) authentication.
// This can be used to authenticate to an EKS cluster or to a cluster that has the AWS IAM Authenticator
// server configured.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/eks_cluster_auth.html.markdown.
func GetClusterAuth(ctx *pulumi.Context, args *GetClusterAuthArgs, opts ...pulumi.InvokeOption) (*GetClusterAuthResult, error) {
	var rv GetClusterAuthResult
	err := ctx.Invoke("aws:eks/getClusterAuth:getClusterAuth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterAuth.
type GetClusterAuthArgs struct {
	// The name of the cluster
	Name string `pulumi:"name"`
}


// A collection of values returned by getClusterAuth.
type GetClusterAuthResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The token to use to authenticate with the cluster.
	Token string `pulumi:"token"`
}

