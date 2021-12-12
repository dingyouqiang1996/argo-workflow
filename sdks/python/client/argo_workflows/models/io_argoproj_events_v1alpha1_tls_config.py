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


class IoArgoprojEventsV1alpha1TLSConfig(object):
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
        'ca_cert_secret': 'SecretKeySelector',
        'client_cert_secret': 'SecretKeySelector',
        'client_key_secret': 'SecretKeySelector'
    }

    attribute_map = {
        'ca_cert_secret': 'caCertSecret',
        'client_cert_secret': 'clientCertSecret',
        'client_key_secret': 'clientKeySecret'
    }

    def __init__(self, ca_cert_secret=None, client_cert_secret=None, client_key_secret=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojEventsV1alpha1TLSConfig - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._ca_cert_secret = None
        self._client_cert_secret = None
        self._client_key_secret = None
        self.discriminator = None

        if ca_cert_secret is not None:
            self.ca_cert_secret = ca_cert_secret
        if client_cert_secret is not None:
            self.client_cert_secret = client_cert_secret
        if client_key_secret is not None:
            self.client_key_secret = client_key_secret

    @property
    def ca_cert_secret(self):
        """Gets the ca_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501


        :return: The ca_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :rtype: SecretKeySelector
        """
        return self._ca_cert_secret

    @ca_cert_secret.setter
    def ca_cert_secret(self, ca_cert_secret):
        """Sets the ca_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.


        :param ca_cert_secret: The ca_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :type: SecretKeySelector
        """

        self._ca_cert_secret = ca_cert_secret

    @property
    def client_cert_secret(self):
        """Gets the client_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501


        :return: The client_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :rtype: SecretKeySelector
        """
        return self._client_cert_secret

    @client_cert_secret.setter
    def client_cert_secret(self, client_cert_secret):
        """Sets the client_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.


        :param client_cert_secret: The client_cert_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :type: SecretKeySelector
        """

        self._client_cert_secret = client_cert_secret

    @property
    def client_key_secret(self):
        """Gets the client_key_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501


        :return: The client_key_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :rtype: SecretKeySelector
        """
        return self._client_key_secret

    @client_key_secret.setter
    def client_key_secret(self, client_key_secret):
        """Sets the client_key_secret of this IoArgoprojEventsV1alpha1TLSConfig.


        :param client_key_secret: The client_key_secret of this IoArgoprojEventsV1alpha1TLSConfig.  # noqa: E501
        :type: SecretKeySelector
        """

        self._client_key_secret = client_key_secret

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
        if not isinstance(other, IoArgoprojEventsV1alpha1TLSConfig):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojEventsV1alpha1TLSConfig):
            return True

        return self.to_dict() != other.to_dict()
