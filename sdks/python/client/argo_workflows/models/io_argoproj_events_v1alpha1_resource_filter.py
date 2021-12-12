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


class IoArgoprojEventsV1alpha1ResourceFilter(object):
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
        'after_start': 'bool',
        'created_by': 'datetime',
        'fields': 'list[IoArgoprojEventsV1alpha1Selector]',
        'labels': 'list[IoArgoprojEventsV1alpha1Selector]',
        'prefix': 'str'
    }

    attribute_map = {
        'after_start': 'afterStart',
        'created_by': 'createdBy',
        'fields': 'fields',
        'labels': 'labels',
        'prefix': 'prefix'
    }

    def __init__(self, after_start=None, created_by=None, fields=None, labels=None, prefix=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojEventsV1alpha1ResourceFilter - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._after_start = None
        self._created_by = None
        self._fields = None
        self._labels = None
        self._prefix = None
        self.discriminator = None

        if after_start is not None:
            self.after_start = after_start
        if created_by is not None:
            self.created_by = created_by
        if fields is not None:
            self.fields = fields
        if labels is not None:
            self.labels = labels
        if prefix is not None:
            self.prefix = prefix

    @property
    def after_start(self):
        """Gets the after_start of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501


        :return: The after_start of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :rtype: bool
        """
        return self._after_start

    @after_start.setter
    def after_start(self, after_start):
        """Sets the after_start of this IoArgoprojEventsV1alpha1ResourceFilter.


        :param after_start: The after_start of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :type: bool
        """

        self._after_start = after_start

    @property
    def created_by(self):
        """Gets the created_by of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501

        Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.  # noqa: E501

        :return: The created_by of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :rtype: datetime
        """
        return self._created_by

    @created_by.setter
    def created_by(self, created_by):
        """Sets the created_by of this IoArgoprojEventsV1alpha1ResourceFilter.

        Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.  # noqa: E501

        :param created_by: The created_by of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :type: datetime
        """

        self._created_by = created_by

    @property
    def fields(self):
        """Gets the fields of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501


        :return: The fields of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :rtype: list[IoArgoprojEventsV1alpha1Selector]
        """
        return self._fields

    @fields.setter
    def fields(self, fields):
        """Sets the fields of this IoArgoprojEventsV1alpha1ResourceFilter.


        :param fields: The fields of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :type: list[IoArgoprojEventsV1alpha1Selector]
        """

        self._fields = fields

    @property
    def labels(self):
        """Gets the labels of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501


        :return: The labels of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :rtype: list[IoArgoprojEventsV1alpha1Selector]
        """
        return self._labels

    @labels.setter
    def labels(self, labels):
        """Sets the labels of this IoArgoprojEventsV1alpha1ResourceFilter.


        :param labels: The labels of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :type: list[IoArgoprojEventsV1alpha1Selector]
        """

        self._labels = labels

    @property
    def prefix(self):
        """Gets the prefix of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501


        :return: The prefix of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :rtype: str
        """
        return self._prefix

    @prefix.setter
    def prefix(self, prefix):
        """Sets the prefix of this IoArgoprojEventsV1alpha1ResourceFilter.


        :param prefix: The prefix of this IoArgoprojEventsV1alpha1ResourceFilter.  # noqa: E501
        :type: str
        """

        self._prefix = prefix

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
        if not isinstance(other, IoArgoprojEventsV1alpha1ResourceFilter):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojEventsV1alpha1ResourceFilter):
            return True

        return self.to_dict() != other.to_dict()
