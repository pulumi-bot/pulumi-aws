// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides details about a specific Amazon Lex Intent.
func LookupIntent(ctx *pulumi.Context, args *LookupIntentArgs, opts ...pulumi.InvokeOption) (*LookupIntentResult, error) {
	var rv LookupIntentResult
	err := ctx.Invoke("aws:lex/getIntent:getIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntent.
type LookupIntentArgs struct {
	// The name of the intent. The name is case sensitive.
	Name string `pulumi:"name"`
	// The version of the intent.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getIntent.
type LookupIntentResult struct {
	// The ARN of the Lex intent.
	Arn string `pulumi:"arn"`
	// Checksum identifying the version of the intent that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the intent.
	Checksum string `pulumi:"checksum"`
	// The date when the intent version was created.
	CreatedDate string `pulumi:"createdDate"`
	// A description of the intent.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The date when the $LATEST version of this intent was updated.
	LastUpdatedDate string `pulumi:"lastUpdatedDate"`
	// The name of the intent, not case sensitive.
	Name string `pulumi:"name"`
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature string `pulumi:"parentIntentSignature"`
	// The version of the bot.
	Version *string `pulumi:"version"`
}
