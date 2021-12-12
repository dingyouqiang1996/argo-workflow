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


class IoArgoprojWorkflowV1alpha1Memoize(object):
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
        'cache': 'IoArgoprojWorkflowV1alpha1Cache',
        'key': 'str',
        'max_age': 'str'
    }

    attribute_map = {
        'cache': 'cache',
        'key': 'key',
        'max_age': 'maxAge'
    }

    def __init__(self, cache=None, key=None, max_age=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojWorkflowV1alpha1Memoize - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._cache = None
        self._key = None
        self._max_age = None
        self.discriminator = None

        self.cache = cache
        self.key = key
        self.max_age = max_age

    @property
    def cache(self):
        """Gets the cache of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501


        :return: The cache of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :rtype: IoArgoprojWorkflowV1alpha1Cache
        """
        return self._cache

    @cache.setter
    def cache(self, cache):
        """Sets the cache of this IoArgoprojWorkflowV1alpha1Memoize.


        :param cache: The cache of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :type: IoArgoprojWorkflowV1alpha1Cache
        """
        if self.local_vars_configuration.client_side_validation and cache is None:  # noqa: E501
            raise ValueError("Invalid value for `cache`, must not be `None`")  # noqa: E501

        self._cache = cache

    @property
    def key(self):
        """Gets the key of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501

        Key is the key to use as the caching key  # noqa: E501

        :return: The key of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :rtype: str
        """
        return self._key

    @key.setter
    def key(self, key):
        """Sets the key of this IoArgoprojWorkflowV1alpha1Memoize.

        Key is the key to use as the caching key  # noqa: E501

        :param key: The key of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and key is None:  # noqa: E501
            raise ValueError("Invalid value for `key`, must not be `None`")  # noqa: E501

        self._key = key

    @property
    def max_age(self):
        """Gets the max_age of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501

        MaxAge is the maximum age (e.g. \"180s\", \"24h\") of an entry that is still considered valid. If an entry is older than the MaxAge, it will be ignored.  # noqa: E501

        :return: The max_age of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :rtype: str
        """
        return self._max_age

    @max_age.setter
    def max_age(self, max_age):
        """Sets the max_age of this IoArgoprojWorkflowV1alpha1Memoize.

        MaxAge is the maximum age (e.g. \"180s\", \"24h\") of an entry that is still considered valid. If an entry is older than the MaxAge, it will be ignored.  # noqa: E501

        :param max_age: The max_age of this IoArgoprojWorkflowV1alpha1Memoize.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and max_age is None:  # noqa: E501
            raise ValueError("Invalid value for `max_age`, must not be `None`")  # noqa: E501

        self._max_age = max_age

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
        if not isinstance(other, IoArgoprojWorkflowV1alpha1Memoize):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojWorkflowV1alpha1Memoize):
            return True

        return self.to_dict() != other.to_dict()
