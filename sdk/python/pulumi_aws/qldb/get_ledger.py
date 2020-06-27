# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLedgerResult:
    """
    A collection of values returned by getLedger.
    """
    def __init__(__self__, arn=None, deletion_protection=None, id=None, name=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Amazon Resource Name (ARN) of the ledger.
        """
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        __self__.deletion_protection = deletion_protection
        """
        Deletion protection on the QLDB Ledger instance. Set to `true` by default.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetLedgerResult(GetLedgerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLedgerResult(
            arn=self.arn,
            deletion_protection=self.deletion_protection,
            id=self.id,
            name=self.name)

def get_ledger(name=None,opts=None):
    """
    Use this data source to fetch information about a Quantum Ledger Database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.qldb.get_ledger(name="an_example_ledger")
    ```


    :param str name: The friendly name of the ledger to match.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:qldb/getLedger:getLedger', __args__, opts=opts).value

    return AwaitableGetLedgerResult(
        arn=__ret__.get('arn'),
        deletion_protection=__ret__.get('deletionProtection'),
        id=__ret__.get('id'),
        name=__ret__.get('name'))
