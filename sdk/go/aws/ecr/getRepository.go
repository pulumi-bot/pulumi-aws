// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecr.LookupRepository(ctx, &ecr.LookupRepositoryArgs{
// 			Name: "ecr-repository",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	var rv LookupRepositoryResult
	err := ctx.Invoke("aws:ecr/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryArgs struct {
	// The name of the ECR Repository.
	Name string `pulumi:"name"`
	// A map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getRepository.
type LookupRepositoryResult struct {
	// Full ARN of the repository.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId string `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// A map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
