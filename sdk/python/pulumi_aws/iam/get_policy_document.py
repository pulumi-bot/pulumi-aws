# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetPolicyDocumentResult:
    """
    A collection of values returned by getPolicyDocument.
    """
    def __init__(__self__, json=None, override_json=None, policy_id=None, source_json=None, statements=None, version=None, id=None):
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        __self__.json = json
        """
        The above arguments serialized as a standard JSON policy document.
        """
        if override_json and not isinstance(override_json, str):
            raise TypeError("Expected argument 'override_json' to be a str")
        __self__.override_json = override_json
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        __self__.policy_id = policy_id
        if source_json and not isinstance(source_json, str):
            raise TypeError("Expected argument 'source_json' to be a str")
        __self__.source_json = source_json
        if statements and not isinstance(statements, list):
            raise TypeError("Expected argument 'statements' to be a list")
        __self__.statements = statements
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_policy_document(override_json=None,policy_id=None,source_json=None,statements=None,version=None,opts=None):
    """
    Generates an IAM policy document in JSON format.
    
    This is a data source which can be used to construct a JSON representation of
    an IAM policy document, for use with resources which expect policy documents,
    such as the `iam.Policy` resource.
    
    
    Using this data source to generate policy documents is *optional*. It is also
    valid to use literal JSON strings within your configuration, or to use the
    `file` interpolation function to read a raw JSON policy document from a file.
    
    ## Context Variable Interpolation
    
    The IAM policy document format allows context variables to be interpolated
    into various strings within a statement. The native IAM policy document format
    uses `${...}`-style syntax that is in conflict with interpolation
    syntax, so this data source instead uses `&{...}` syntax for interpolations that
    should be processed by AWS rather than by this provider.
    
    ## Wildcard Principal
    
    In order to define wildcard principal (a.k.a. anonymous user) use `type = "*"` and
    `identifiers = ["*"]`. In that case the rendered json will contain `"Principal": "*"`.
    Note, that even though the [IAM Documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html)
    states that `"Principal": "*"` and `"Principal": {"AWS": "*"}` are equivalent,
    those principals have different behavior for IAM Role Trust Policy. Therefore
    this provider will normalize the principal field only in above-mentioned case and principals
    like `type = "AWS"` and `identifiers = ["*"]` will be rendered as `"Principal": {"AWS": "*"}`.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_policy_document.html.markdown.
    """
    __args__ = dict()

    __args__['overrideJson'] = override_json
    __args__['policyId'] = policy_id
    __args__['sourceJson'] = source_json
    __args__['statements'] = statements
    __args__['version'] = version
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iam/getPolicyDocument:getPolicyDocument', __args__, opts=opts).value

    return GetPolicyDocumentResult(
        json=__ret__.get('json'),
        override_json=__ret__.get('overrideJson'),
        policy_id=__ret__.get('policyId'),
        source_json=__ret__.get('sourceJson'),
        statements=__ret__.get('statements'),
        version=__ret__.get('version'),
        id=__ret__.get('id'))
