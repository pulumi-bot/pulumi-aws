// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about an Elastic File System (EFS) File System.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		fileSystemId := ""
// 		if param := cfg.Get("fileSystemId"); param != "" {
// 			fileSystemId = param
// 		}
// 		opt0 := fileSystemId
// 		_, err := efs.LookupFileSystem(ctx, &efs.LookupFileSystemArgs{
// 			FileSystemId: _opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFileSystem(ctx *pulumi.Context, args *LookupFileSystemArgs, opts ...pulumi.InvokeOption) (*LookupFileSystemResult, error) {
	var rv LookupFileSystemResult
	err := ctx.Invoke("aws:efs/getFileSystem:getFileSystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFileSystem.
type LookupFileSystemArgs struct {
	// Restricts the list to the file system with this creation token.
	CreationToken *string `pulumi:"creationToken"`
	// The ID that identifies the file system (e.g. fs-ccfc0d65).
	FileSystemId *string           `pulumi:"fileSystemId"`
	Tags         map[string]string `pulumi:"tags"`
}

// A collection of values returned by getFileSystem.
type LookupFileSystemResult struct {
	// Amazon Resource Name of the file system.
	Arn           string `pulumi:"arn"`
	CreationToken string `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName string `pulumi:"dnsName"`
	// Whether EFS is encrypted.
	Encrypted    bool   `pulumi:"encrypted"`
	FileSystemId string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ARN for the KMS encryption key.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object.
	LifecyclePolicy GetFileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	// The file system performance mode.
	PerformanceMode string `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system.
	// * `tags` -A map of tags to assign to the file system.
	ProvisionedThroughputInMibps float64 `pulumi:"provisionedThroughputInMibps"`
	// The current byte count used by the file system.
	SizeInBytes int               `pulumi:"sizeInBytes"`
	Tags        map[string]string `pulumi:"tags"`
	// Throughput mode for the file system.
	ThroughputMode string `pulumi:"throughputMode"`
}
