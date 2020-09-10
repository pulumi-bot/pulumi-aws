// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	var rv LookupDocumentResult
	err := ctx.Invoke("aws:ssm/getDocument:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocument.
type LookupDocumentArgs struct {
	DocumentFormat  *string `pulumi:"documentFormat"`
	DocumentVersion *string `pulumi:"documentVersion"`
	Name            string  `pulumi:"name"`
}

// A collection of values returned by getDocument.
type LookupDocumentResult struct {
	Arn             string  `pulumi:"arn"`
	Content         string  `pulumi:"content"`
	DocumentFormat  *string `pulumi:"documentFormat"`
	DocumentType    string  `pulumi:"documentType"`
	DocumentVersion *string `pulumi:"documentVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
