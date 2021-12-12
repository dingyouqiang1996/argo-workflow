# coding: utf-8

"""
    Argo Server API

    You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`  # noqa: E501

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from argo_workflows.configuration import Configuration


class IoArgoprojWorkflowV1alpha1RetryStrategy(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'affinity': 'IoArgoprojWorkflowV1alpha1RetryAffinity',
        'backoff': 'IoArgoprojWorkflowV1alpha1Backoff',
        'expression': 'str',
        'limit': 'str',
        'retry_policy': 'str'
    }

    attribute_map = {
        'affinity': 'affinity',
        'backoff': 'backoff',
        'expression': 'expression',
        'limit': 'limit',
        'retry_policy': 'retryPolicy'
    }

    def __init__(self, affinity=None, backoff=None, expression=None, limit=None, retry_policy=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojWorkflowV1alpha1RetryStrategy - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._affinity = None
        self._backoff = None
        self._expression = None
        self._limit = None
        self._retry_policy = None
        self.discriminator = None

        if affinity is not None:
            self.affinity = affinity
        if backoff is not None:
            self.backoff = backoff
        if expression is not None:
            self.expression = expression
        if limit is not None:
            self.limit = limit
        if retry_policy is not None:
            self.retry_policy = retry_policy

    @property
    def affinity(self):
        """Gets the affinity of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501


        :return: The affinity of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :rtype: IoArgoprojWorkflowV1alpha1RetryAffinity
        """
        return self._affinity

    @affinity.setter
    def affinity(self, affinity):
        """Sets the affinity of this IoArgoprojWorkflowV1alpha1RetryStrategy.


        :param affinity: The affinity of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :type: IoArgoprojWorkflowV1alpha1RetryAffinity
        """

        self._affinity = affinity

    @property
    def backoff(self):
        """Gets the backoff of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501


        :return: The backoff of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :rtype: IoArgoprojWorkflowV1alpha1Backoff
        """
        return self._backoff

    @backoff.setter
    def backoff(self, backoff):
        """Sets the backoff of this IoArgoprojWorkflowV1alpha1RetryStrategy.


        :param backoff: The backoff of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :type: IoArgoprojWorkflowV1alpha1Backoff
        """

        self._backoff = backoff

    @property
    def expression(self):
        """Gets the expression of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501

        Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored  # noqa: E501

        :return: The expression of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :rtype: str
        """
        return self._expression

    @expression.setter
    def expression(self, expression):
        """Sets the expression of this IoArgoprojWorkflowV1alpha1RetryStrategy.

        Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored  # noqa: E501

        :param expression: The expression of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :type: str
        """

        self._expression = expression

    @property
    def limit(self):
        """Gets the limit of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501


        :return: The limit of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :rtype: str
        """
        return self._limit

    @limit.setter
    def limit(self, limit):
        """Sets the limit of this IoArgoprojWorkflowV1alpha1RetryStrategy.


        :param limit: The limit of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :type: str
        """

        self._limit = limit

    @property
    def retry_policy(self):
        """Gets the retry_policy of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501

        RetryPolicy is a policy of NodePhase statuses that will be retried  # noqa: E501

        :return: The retry_policy of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :rtype: str
        """
        return self._retry_policy

    @retry_policy.setter
    def retry_policy(self, retry_policy):
        """Sets the retry_policy of this IoArgoprojWorkflowV1alpha1RetryStrategy.

        RetryPolicy is a policy of NodePhase statuses that will be retried  # noqa: E501

        :param retry_policy: The retry_policy of this IoArgoprojWorkflowV1alpha1RetryStrategy.  # noqa: E501
        :type: str
        """

        self._retry_policy = retry_policy

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, IoArgoprojWorkflowV1alpha1RetryStrategy):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojWorkflowV1alpha1RetryStrategy):
            return True

        return self.to_dict() != other.to_dict()
