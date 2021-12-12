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


class IoArgoprojEventsV1alpha1RateLimit(object):
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
        'requests_per_unit': 'int',
        'unit': 'str'
    }

    attribute_map = {
        'requests_per_unit': 'requestsPerUnit',
        'unit': 'unit'
    }

    def __init__(self, requests_per_unit=None, unit=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojEventsV1alpha1RateLimit - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._requests_per_unit = None
        self._unit = None
        self.discriminator = None

        if requests_per_unit is not None:
            self.requests_per_unit = requests_per_unit
        if unit is not None:
            self.unit = unit

    @property
    def requests_per_unit(self):
        """Gets the requests_per_unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501


        :return: The requests_per_unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501
        :rtype: int
        """
        return self._requests_per_unit

    @requests_per_unit.setter
    def requests_per_unit(self, requests_per_unit):
        """Sets the requests_per_unit of this IoArgoprojEventsV1alpha1RateLimit.


        :param requests_per_unit: The requests_per_unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501
        :type: int
        """

        self._requests_per_unit = requests_per_unit

    @property
    def unit(self):
        """Gets the unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501


        :return: The unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501
        :rtype: str
        """
        return self._unit

    @unit.setter
    def unit(self, unit):
        """Sets the unit of this IoArgoprojEventsV1alpha1RateLimit.


        :param unit: The unit of this IoArgoprojEventsV1alpha1RateLimit.  # noqa: E501
        :type: str
        """

        self._unit = unit

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
        if not isinstance(other, IoArgoprojEventsV1alpha1RateLimit):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojEventsV1alpha1RateLimit):
            return True

        return self.to_dict() != other.to_dict()
