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


class GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron(object):
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
        'layout': 'str',
        'schedule': 'str'
    }

    attribute_map = {
        'layout': 'layout',
        'schedule': 'schedule'
    }

    def __init__(self, layout=None, schedule=None, local_vars_configuration=None):  # noqa: E501
        """GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._layout = None
        self._schedule = None
        self.discriminator = None

        if layout is not None:
            self.layout = layout
        if schedule is not None:
            self.schedule = schedule

    @property
    def layout(self):
        """Gets the layout of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501


        :return: The layout of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501
        :rtype: str
        """
        return self._layout

    @layout.setter
    def layout(self, layout):
        """Sets the layout of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.


        :param layout: The layout of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501
        :type: str
        """

        self._layout = layout

    @property
    def schedule(self):
        """Gets the schedule of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501


        :return: The schedule of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501
        :rtype: str
        """
        return self._schedule

    @schedule.setter
    def schedule(self, schedule):
        """Sets the schedule of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.


        :param schedule: The schedule of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.  # noqa: E501
        :type: str
        """

        self._schedule = schedule

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
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron):
            return True

        return self.to_dict() != other.to_dict()
