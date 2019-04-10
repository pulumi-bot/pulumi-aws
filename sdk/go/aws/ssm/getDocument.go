// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets the contents of the specified Systems Manager document.
func LookupDocument(ctx *pulumi.Context, args *GetDocumentArgs) (*GetDocumentResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["documentFormat"] = args.DocumentFormat
		inputs["documentVersion"] = args.DocumentVersion
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:ssm/getDocument:getDocument", inputs)
	if err != nil {
		return nil, err
	}
	return &GetDocumentResult{
		Arn: outputs["arn"],
		Content: outputs["content"],
		DocumentFormat: outputs["documentFormat"],
		DocumentType: outputs["documentType"],
		DocumentVersion: outputs["documentVersion"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getDocument.
type GetDocumentArgs struct {
	// Returns the document in the specified format. The document format can be either JSON or YAML. JSON is the default format.
	DocumentFormat interface{}
	// The document version for which you want information.
	DocumentVersion interface{}
	// The name of the Systems Manager document.
	Name interface{}
}

// A collection of values returned by getDocument.
type GetDocumentResult struct {
	// The ARN of the document.
	Arn interface{}
	// The contents of the document.
	Content interface{}
	DocumentFormat interface{}
	// The type of the document.
	DocumentType interface{}
	DocumentVersion interface{}
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
