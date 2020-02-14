# coding: utf-8

"""
    Argo

    Workflow Service API performs CRUD actions against application resources  # noqa: E501

    The version of the OpenAPI document: latest
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from openapi_client.configuration import Configuration


class V1EnvVar(object):
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
        'name': 'str',
        'value': 'str',
        'value_from': 'V1EnvVarSource'
    }

    attribute_map = {
        'name': 'name',
        'value': 'value',
        'value_from': 'valueFrom'
    }

    def __init__(self, name=None, value=None, value_from=None, local_vars_configuration=None):  # noqa: E501
        """V1EnvVar - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._name = None
        self._value = None
        self._value_from = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if value is not None:
            self.value = value
        if value_from is not None:
            self.value_from = value_from

    @property
    def name(self):
        """Gets the name of this V1EnvVar.  # noqa: E501

        Name of the environment variable. Must be a C_IDENTIFIER.  # noqa: E501

        :return: The name of this V1EnvVar.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1EnvVar.

        Name of the environment variable. Must be a C_IDENTIFIER.  # noqa: E501

        :param name: The name of this V1EnvVar.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def value(self):
        """Gets the value of this V1EnvVar.  # noqa: E501


        :return: The value of this V1EnvVar.  # noqa: E501
        :rtype: str
        """
        return self._value

    @value.setter
    def value(self, value):
        """Sets the value of this V1EnvVar.


        :param value: The value of this V1EnvVar.  # noqa: E501
        :type: str
        """

        self._value = value

    @property
    def value_from(self):
        """Gets the value_from of this V1EnvVar.  # noqa: E501


        :return: The value_from of this V1EnvVar.  # noqa: E501
        :rtype: V1EnvVarSource
        """
        return self._value_from

    @value_from.setter
    def value_from(self, value_from):
        """Sets the value_from of this V1EnvVar.


        :param value_from: The value_from of this V1EnvVar.  # noqa: E501
        :type: V1EnvVarSource
        """

        self._value_from = value_from

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
        if not isinstance(other, V1EnvVar):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1EnvVar):
            return True

        return self.to_dict() != other.to_dict()
