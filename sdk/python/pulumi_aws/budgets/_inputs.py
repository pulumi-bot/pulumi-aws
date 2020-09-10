# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'BudgetCostTypesArgs',
    'BudgetNotificationArgs',
]

@pulumi.input_type
class BudgetCostTypesArgs:
    def __init__(__self__, *,
                 include_credit: Optional[pulumi.Input[bool]] = None,
                 include_discount: Optional[pulumi.Input[bool]] = None,
                 include_other_subscription: Optional[pulumi.Input[bool]] = None,
                 include_recurring: Optional[pulumi.Input[bool]] = None,
                 include_refund: Optional[pulumi.Input[bool]] = None,
                 include_subscription: Optional[pulumi.Input[bool]] = None,
                 include_support: Optional[pulumi.Input[bool]] = None,
                 include_tax: Optional[pulumi.Input[bool]] = None,
                 include_upfront: Optional[pulumi.Input[bool]] = None,
                 use_amortized: Optional[pulumi.Input[bool]] = None,
                 use_blended: Optional[pulumi.Input[bool]] = None):
        if include_credit is not None:
            pulumi.set(__self__, "include_credit", include_credit)
        if include_discount is not None:
            pulumi.set(__self__, "include_discount", include_discount)
        if include_other_subscription is not None:
            pulumi.set(__self__, "include_other_subscription", include_other_subscription)
        if include_recurring is not None:
            pulumi.set(__self__, "include_recurring", include_recurring)
        if include_refund is not None:
            pulumi.set(__self__, "include_refund", include_refund)
        if include_subscription is not None:
            pulumi.set(__self__, "include_subscription", include_subscription)
        if include_support is not None:
            pulumi.set(__self__, "include_support", include_support)
        if include_tax is not None:
            pulumi.set(__self__, "include_tax", include_tax)
        if include_upfront is not None:
            pulumi.set(__self__, "include_upfront", include_upfront)
        if use_amortized is not None:
            pulumi.set(__self__, "use_amortized", use_amortized)
        if use_blended is not None:
            pulumi.set(__self__, "use_blended", use_blended)

    @property
    @pulumi.getter(name="includeCredit")
    def include_credit(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_credit")

    @include_credit.setter
    def include_credit(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_credit", value)

    @property
    @pulumi.getter(name="includeDiscount")
    def include_discount(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_discount")

    @include_discount.setter
    def include_discount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_discount", value)

    @property
    @pulumi.getter(name="includeOtherSubscription")
    def include_other_subscription(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_other_subscription")

    @include_other_subscription.setter
    def include_other_subscription(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_other_subscription", value)

    @property
    @pulumi.getter(name="includeRecurring")
    def include_recurring(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_recurring")

    @include_recurring.setter
    def include_recurring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_recurring", value)

    @property
    @pulumi.getter(name="includeRefund")
    def include_refund(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_refund")

    @include_refund.setter
    def include_refund(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_refund", value)

    @property
    @pulumi.getter(name="includeSubscription")
    def include_subscription(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_subscription")

    @include_subscription.setter
    def include_subscription(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_subscription", value)

    @property
    @pulumi.getter(name="includeSupport")
    def include_support(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_support")

    @include_support.setter
    def include_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_support", value)

    @property
    @pulumi.getter(name="includeTax")
    def include_tax(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_tax")

    @include_tax.setter
    def include_tax(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_tax", value)

    @property
    @pulumi.getter(name="includeUpfront")
    def include_upfront(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_upfront")

    @include_upfront.setter
    def include_upfront(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_upfront", value)

    @property
    @pulumi.getter(name="useAmortized")
    def use_amortized(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_amortized")

    @use_amortized.setter
    def use_amortized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_amortized", value)

    @property
    @pulumi.getter(name="useBlended")
    def use_blended(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_blended")

    @use_blended.setter
    def use_blended(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_blended", value)


@pulumi.input_type
class BudgetNotificationArgs:
    def __init__(__self__, *,
                 comparison_operator: pulumi.Input[str],
                 notification_type: pulumi.Input[str],
                 threshold: pulumi.Input[float],
                 threshold_type: pulumi.Input[str],
                 subscriber_email_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 subscriber_sns_topic_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        pulumi.set(__self__, "notification_type", notification_type)
        pulumi.set(__self__, "threshold", threshold)
        pulumi.set(__self__, "threshold_type", threshold_type)
        if subscriber_email_addresses is not None:
            pulumi.set(__self__, "subscriber_email_addresses", subscriber_email_addresses)
        if subscriber_sns_topic_arns is not None:
            pulumi.set(__self__, "subscriber_sns_topic_arns", subscriber_sns_topic_arns)

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> pulumi.Input[str]:
        return pulumi.get(self, "comparison_operator")

    @comparison_operator.setter
    def comparison_operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison_operator", value)

    @property
    @pulumi.getter(name="notificationType")
    def notification_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "notification_type")

    @notification_type.setter
    def notification_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "notification_type", value)

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Input[float]:
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "threshold_type")

    @threshold_type.setter
    def threshold_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "threshold_type", value)

    @property
    @pulumi.getter(name="subscriberEmailAddresses")
    def subscriber_email_addresses(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        return pulumi.get(self, "subscriber_email_addresses")

    @subscriber_email_addresses.setter
    def subscriber_email_addresses(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "subscriber_email_addresses", value)

    @property
    @pulumi.getter(name="subscriberSnsTopicArns")
    def subscriber_sns_topic_arns(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        return pulumi.get(self, "subscriber_sns_topic_arns")

    @subscriber_sns_topic_arns.setter
    def subscriber_sns_topic_arns(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "subscriber_sns_topic_arns", value)


