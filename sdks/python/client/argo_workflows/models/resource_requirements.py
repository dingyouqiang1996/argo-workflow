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


class ResourceRequirements(object):
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
        'limits': 'dict(str, str)',
        'requests': 'dict(str, str)'
    }

    attribute_map = {
        'limits': 'limits',
        'requests': 'requests'
    }

    def __init__(self, limits=None, requests=None, local_vars_configuration=None):  # noqa: E501
        """ResourceRequirements - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._limits = None
        self._requests = None
        self.discriminator = None

        if limits is not None:
            self.limits = limits
        if requests is not None:
            self.requests = requests

    @property
    def limits(self):
        """Gets the limits of this ResourceRequirements.  # noqa: E501

        Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/  # noqa: E501

        :return: The limits of this ResourceRequirements.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._limits

    @limits.setter
    def limits(self, limits):
        """Sets the limits of this ResourceRequirements.

        Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/  # noqa: E501

        :param limits: The limits of this ResourceRequirements.  # noqa: E501
        :type: dict(str, str)
        """

        self._limits = limits

    @property
    def requests(self):
        """Gets the requests of this ResourceRequirements.  # noqa: E501

        Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/  # noqa: E501

        :return: The requests of this ResourceRequirements.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._requests

    @requests.setter
    def requests(self, requests):
        """Sets the requests of this ResourceRequirements.

        Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/  # noqa: E501

        :param requests: The requests of this ResourceRequirements.  # noqa: E501
        :type: dict(str, str)
        """

        self._requests = requests

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
        if not isinstance(other, ResourceRequirements):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ResourceRequirements):
            return True

        return self.to_dict() != other.to_dict()
