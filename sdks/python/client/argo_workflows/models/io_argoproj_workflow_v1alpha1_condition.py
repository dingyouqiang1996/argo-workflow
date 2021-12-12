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


class IoArgoprojWorkflowV1alpha1Condition(object):
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
        'message': 'str',
        'status': 'str',
        'type': 'str'
    }

    attribute_map = {
        'message': 'message',
        'status': 'status',
        'type': 'type'
    }

    def __init__(self, message=None, status=None, type=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojWorkflowV1alpha1Condition - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._message = None
        self._status = None
        self._type = None
        self.discriminator = None

        if message is not None:
            self.message = message
        if status is not None:
            self.status = status
        if type is not None:
            self.type = type

    @property
    def message(self):
        """Gets the message of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501

        Message is the condition message  # noqa: E501

        :return: The message of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :rtype: str
        """
        return self._message

    @message.setter
    def message(self, message):
        """Sets the message of this IoArgoprojWorkflowV1alpha1Condition.

        Message is the condition message  # noqa: E501

        :param message: The message of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :type: str
        """

        self._message = message

    @property
    def status(self):
        """Gets the status of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501

        Status is the status of the condition  # noqa: E501

        :return: The status of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :rtype: str
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this IoArgoprojWorkflowV1alpha1Condition.

        Status is the status of the condition  # noqa: E501

        :param status: The status of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :type: str
        """

        self._status = status

    @property
    def type(self):
        """Gets the type of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501

        Type is the type of condition  # noqa: E501

        :return: The type of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this IoArgoprojWorkflowV1alpha1Condition.

        Type is the type of condition  # noqa: E501

        :param type: The type of this IoArgoprojWorkflowV1alpha1Condition.  # noqa: E501
        :type: str
        """

        self._type = type

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
        if not isinstance(other, IoArgoprojWorkflowV1alpha1Condition):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojWorkflowV1alpha1Condition):
            return True

        return self.to_dict() != other.to_dict()
