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


class VsphereVirtualDiskVolumeSource(object):
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
        'fs_type': 'str',
        'storage_policy_id': 'str',
        'storage_policy_name': 'str',
        'volume_path': 'str'
    }

    attribute_map = {
        'fs_type': 'fsType',
        'storage_policy_id': 'storagePolicyID',
        'storage_policy_name': 'storagePolicyName',
        'volume_path': 'volumePath'
    }

    def __init__(self, fs_type=None, storage_policy_id=None, storage_policy_name=None, volume_path=None, local_vars_configuration=None):  # noqa: E501
        """VsphereVirtualDiskVolumeSource - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._fs_type = None
        self._storage_policy_id = None
        self._storage_policy_name = None
        self._volume_path = None
        self.discriminator = None

        if fs_type is not None:
            self.fs_type = fs_type
        if storage_policy_id is not None:
            self.storage_policy_id = storage_policy_id
        if storage_policy_name is not None:
            self.storage_policy_name = storage_policy_name
        self.volume_path = volume_path

    @property
    def fs_type(self):
        """Gets the fs_type of this VsphereVirtualDiskVolumeSource.  # noqa: E501

        Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.  # noqa: E501

        :return: The fs_type of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :rtype: str
        """
        return self._fs_type

    @fs_type.setter
    def fs_type(self, fs_type):
        """Sets the fs_type of this VsphereVirtualDiskVolumeSource.

        Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.  # noqa: E501

        :param fs_type: The fs_type of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :type: str
        """

        self._fs_type = fs_type

    @property
    def storage_policy_id(self):
        """Gets the storage_policy_id of this VsphereVirtualDiskVolumeSource.  # noqa: E501

        Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.  # noqa: E501

        :return: The storage_policy_id of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :rtype: str
        """
        return self._storage_policy_id

    @storage_policy_id.setter
    def storage_policy_id(self, storage_policy_id):
        """Sets the storage_policy_id of this VsphereVirtualDiskVolumeSource.

        Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.  # noqa: E501

        :param storage_policy_id: The storage_policy_id of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :type: str
        """

        self._storage_policy_id = storage_policy_id

    @property
    def storage_policy_name(self):
        """Gets the storage_policy_name of this VsphereVirtualDiskVolumeSource.  # noqa: E501

        Storage Policy Based Management (SPBM) profile name.  # noqa: E501

        :return: The storage_policy_name of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :rtype: str
        """
        return self._storage_policy_name

    @storage_policy_name.setter
    def storage_policy_name(self, storage_policy_name):
        """Sets the storage_policy_name of this VsphereVirtualDiskVolumeSource.

        Storage Policy Based Management (SPBM) profile name.  # noqa: E501

        :param storage_policy_name: The storage_policy_name of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :type: str
        """

        self._storage_policy_name = storage_policy_name

    @property
    def volume_path(self):
        """Gets the volume_path of this VsphereVirtualDiskVolumeSource.  # noqa: E501

        Path that identifies vSphere volume vmdk  # noqa: E501

        :return: The volume_path of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :rtype: str
        """
        return self._volume_path

    @volume_path.setter
    def volume_path(self, volume_path):
        """Sets the volume_path of this VsphereVirtualDiskVolumeSource.

        Path that identifies vSphere volume vmdk  # noqa: E501

        :param volume_path: The volume_path of this VsphereVirtualDiskVolumeSource.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and volume_path is None:  # noqa: E501
            raise ValueError("Invalid value for `volume_path`, must not be `None`")  # noqa: E501

        self._volume_path = volume_path

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
        if not isinstance(other, VsphereVirtualDiskVolumeSource):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, VsphereVirtualDiskVolumeSource):
            return True

        return self.to_dict() != other.to_dict()
