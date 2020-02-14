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


class RuntimeStreamError(object):
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
        'details': 'list[ProtobufAny]',
        'grpc_code': 'int',
        'http_code': 'int',
        'http_status': 'str',
        'message': 'str'
    }

    attribute_map = {
        'details': 'details',
        'grpc_code': 'grpc_code',
        'http_code': 'http_code',
        'http_status': 'http_status',
        'message': 'message'
    }

    def __init__(self, details=None, grpc_code=None, http_code=None, http_status=None, message=None, local_vars_configuration=None):  # noqa: E501
        """RuntimeStreamError - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._details = None
        self._grpc_code = None
        self._http_code = None
        self._http_status = None
        self._message = None
        self.discriminator = None

        if details is not None:
            self.details = details
        if grpc_code is not None:
            self.grpc_code = grpc_code
        if http_code is not None:
            self.http_code = http_code
        if http_status is not None:
            self.http_status = http_status
        if message is not None:
            self.message = message

    @property
    def details(self):
        """Gets the details of this RuntimeStreamError.  # noqa: E501


        :return: The details of this RuntimeStreamError.  # noqa: E501
        :rtype: list[ProtobufAny]
        """
        return self._details

    @details.setter
    def details(self, details):
        """Sets the details of this RuntimeStreamError.


        :param details: The details of this RuntimeStreamError.  # noqa: E501
        :type: list[ProtobufAny]
        """

        self._details = details

    @property
    def grpc_code(self):
        """Gets the grpc_code of this RuntimeStreamError.  # noqa: E501


        :return: The grpc_code of this RuntimeStreamError.  # noqa: E501
        :rtype: int
        """
        return self._grpc_code

    @grpc_code.setter
    def grpc_code(self, grpc_code):
        """Sets the grpc_code of this RuntimeStreamError.


        :param grpc_code: The grpc_code of this RuntimeStreamError.  # noqa: E501
        :type: int
        """

        self._grpc_code = grpc_code

    @property
    def http_code(self):
        """Gets the http_code of this RuntimeStreamError.  # noqa: E501


        :return: The http_code of this RuntimeStreamError.  # noqa: E501
        :rtype: int
        """
        return self._http_code

    @http_code.setter
    def http_code(self, http_code):
        """Sets the http_code of this RuntimeStreamError.


        :param http_code: The http_code of this RuntimeStreamError.  # noqa: E501
        :type: int
        """

        self._http_code = http_code

    @property
    def http_status(self):
        """Gets the http_status of this RuntimeStreamError.  # noqa: E501


        :return: The http_status of this RuntimeStreamError.  # noqa: E501
        :rtype: str
        """
        return self._http_status

    @http_status.setter
    def http_status(self, http_status):
        """Sets the http_status of this RuntimeStreamError.


        :param http_status: The http_status of this RuntimeStreamError.  # noqa: E501
        :type: str
        """

        self._http_status = http_status

    @property
    def message(self):
        """Gets the message of this RuntimeStreamError.  # noqa: E501


        :return: The message of this RuntimeStreamError.  # noqa: E501
        :rtype: str
        """
        return self._message

    @message.setter
    def message(self, message):
        """Sets the message of this RuntimeStreamError.


        :param message: The message of this RuntimeStreamError.  # noqa: E501
        :type: str
        """

        self._message = message

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
        if not isinstance(other, RuntimeStreamError):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, RuntimeStreamError):
            return True

        return self.to_dict() != other.to_dict()
