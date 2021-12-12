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


class IoArgoprojEventsV1alpha1GenericEventSource(object):
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
        'auth_secret': 'SecretKeySelector',
        'config': 'str',
        'insecure': 'bool',
        'json_body': 'bool',
        'metadata': 'dict(str, str)',
        'url': 'str'
    }

    attribute_map = {
        'auth_secret': 'authSecret',
        'config': 'config',
        'insecure': 'insecure',
        'json_body': 'jsonBody',
        'metadata': 'metadata',
        'url': 'url'
    }

    def __init__(self, auth_secret=None, config=None, insecure=None, json_body=None, metadata=None, url=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojEventsV1alpha1GenericEventSource - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._auth_secret = None
        self._config = None
        self._insecure = None
        self._json_body = None
        self._metadata = None
        self._url = None
        self.discriminator = None

        if auth_secret is not None:
            self.auth_secret = auth_secret
        if config is not None:
            self.config = config
        if insecure is not None:
            self.insecure = insecure
        if json_body is not None:
            self.json_body = json_body
        if metadata is not None:
            self.metadata = metadata
        if url is not None:
            self.url = url

    @property
    def auth_secret(self):
        """Gets the auth_secret of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501


        :return: The auth_secret of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: SecretKeySelector
        """
        return self._auth_secret

    @auth_secret.setter
    def auth_secret(self, auth_secret):
        """Sets the auth_secret of this IoArgoprojEventsV1alpha1GenericEventSource.


        :param auth_secret: The auth_secret of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: SecretKeySelector
        """

        self._auth_secret = auth_secret

    @property
    def config(self):
        """Gets the config of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501


        :return: The config of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: str
        """
        return self._config

    @config.setter
    def config(self, config):
        """Sets the config of this IoArgoprojEventsV1alpha1GenericEventSource.


        :param config: The config of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: str
        """

        self._config = config

    @property
    def insecure(self):
        """Gets the insecure of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501

        Insecure determines the type of connection.  # noqa: E501

        :return: The insecure of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: bool
        """
        return self._insecure

    @insecure.setter
    def insecure(self, insecure):
        """Sets the insecure of this IoArgoprojEventsV1alpha1GenericEventSource.

        Insecure determines the type of connection.  # noqa: E501

        :param insecure: The insecure of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: bool
        """

        self._insecure = insecure

    @property
    def json_body(self):
        """Gets the json_body of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501


        :return: The json_body of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: bool
        """
        return self._json_body

    @json_body.setter
    def json_body(self, json_body):
        """Sets the json_body of this IoArgoprojEventsV1alpha1GenericEventSource.


        :param json_body: The json_body of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: bool
        """

        self._json_body = json_body

    @property
    def metadata(self):
        """Gets the metadata of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501


        :return: The metadata of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this IoArgoprojEventsV1alpha1GenericEventSource.


        :param metadata: The metadata of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: dict(str, str)
        """

        self._metadata = metadata

    @property
    def url(self):
        """Gets the url of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501

        URL of the gRPC server that implements the event source.  # noqa: E501

        :return: The url of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :rtype: str
        """
        return self._url

    @url.setter
    def url(self, url):
        """Sets the url of this IoArgoprojEventsV1alpha1GenericEventSource.

        URL of the gRPC server that implements the event source.  # noqa: E501

        :param url: The url of this IoArgoprojEventsV1alpha1GenericEventSource.  # noqa: E501
        :type: str
        """

        self._url = url

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
        if not isinstance(other, IoArgoprojEventsV1alpha1GenericEventSource):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojEventsV1alpha1GenericEventSource):
            return True

        return self.to_dict() != other.to_dict()
