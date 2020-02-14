# coding: utf-8

"""
    Argo

    Workflow Service API performs CRUD actions against application resources  # noqa: E501

    The version of the OpenAPI document: latest
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.io.argoproj.argo.client.model.workflowv1alpha1_volume_source import Workflowv1alpha1VolumeSource  # noqa: E501
from openapi_client.rest import ApiException

class TestWorkflowv1alpha1VolumeSource(unittest.TestCase):
    """Workflowv1alpha1VolumeSource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test Workflowv1alpha1VolumeSource
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.workflowv1alpha1_volume_source.Workflowv1alpha1VolumeSource()  # noqa: E501
        if include_optional :
            return Workflowv1alpha1VolumeSource(
                aws_elastic_block_store = openapi_client.models.workflowv1alpha1_aws_elastic_block_store_volume_source.workflowv1alpha1AWSElasticBlockStoreVolumeSource(
                    fs_type = '0', 
                    partition = 56, 
                    read_only = True, 
                    volume_id = '0', ), 
                azure_disk = openapi_client.models.workflowv1alpha1_azure_disk_volume_source.workflowv1alpha1AzureDiskVolumeSource(
                    caching_mode = '0', 
                    disk_name = '0', 
                    disk_uri = '0', 
                    fs_type = '0', 
                    kind = '0', 
                    read_only = True, ), 
                azure_file = openapi_client.models.workflowv1alpha1_azure_file_volume_source.workflowv1alpha1AzureFileVolumeSource(
                    read_only = True, 
                    secret_name = '0', 
                    share_name = '0', ), 
                cephfs = openapi_client.models.workflowv1alpha1_ceph_fs_volume_source.workflowv1alpha1CephFSVolumeSource(
                    monitors = [
                        '0'
                        ], 
                    path = '0', 
                    read_only = True, 
                    secret_file = '0', 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    user = '0', ), 
                cinder = openapi_client.models.workflowv1alpha1_cinder_volume_source.workflowv1alpha1CinderVolumeSource(
                    fs_type = '0', 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    volume_id = '0', ), 
                config_map = openapi_client.models.workflowv1alpha1_config_map_volume_source.workflowv1alpha1ConfigMapVolumeSource(
                    default_mode = 56, 
                    items = [
                        openapi_client.models.workflowv1alpha1_key_to_path.workflowv1alpha1KeyToPath(
                            key = '0', 
                            mode = 56, 
                            path = '0', )
                        ], 
                    local_object_reference = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    optional = True, ), 
                csi = openapi_client.models.workflowv1alpha1_csi_volume_source.workflowv1alpha1CSIVolumeSource(
                    driver = '0', 
                    fs_type = '0', 
                    node_publish_secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    read_only = True, 
                    volume_attributes = {
                        'key' : '0'
                        }, ), 
                downward_api = openapi_client.models.workflowv1alpha1_downward_api_volume_source.workflowv1alpha1DownwardAPIVolumeSource(
                    default_mode = 56, 
                    items = [
                        openapi_client.models.workflowv1alpha1_downward_api_volume_file.workflowv1alpha1DownwardAPIVolumeFile(
                            field_ref = openapi_client.models.workflowv1alpha1_object_field_selector.workflowv1alpha1ObjectFieldSelector(
                                api_version = '0', 
                                field_path = '0', ), 
                            mode = 56, 
                            path = '0', 
                            resource_field_ref = openapi_client.models.workflowv1alpha1_resource_field_selector.workflowv1alpha1ResourceFieldSelector(
                                container_name = '0', 
                                divisor = openapi_client.models.resource_quantity.resourceQuantity(
                                    string = '0', ), 
                                resource = '0', ), )
                        ], ), 
                empty_dir = openapi_client.models.workflowv1alpha1_empty_dir_volume_source.workflowv1alpha1EmptyDirVolumeSource(
                    medium = '0', 
                    size_limit = openapi_client.models.resource_quantity.resourceQuantity(
                        string = '0', ), ), 
                fc = openapi_client.models.workflowv1alpha1_fc_volume_source.workflowv1alpha1FCVolumeSource(
                    fs_type = '0', 
                    lun = 56, 
                    read_only = True, 
                    target_ww_ns = [
                        '0'
                        ], 
                    wwids = [
                        '0'
                        ], ), 
                flex_volume = openapi_client.models.workflowv1alpha1_flex_volume_source.workflowv1alpha1FlexVolumeSource(
                    driver = '0', 
                    fs_type = '0', 
                    options = {
                        'key' : '0'
                        }, 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), ), 
                flocker = openapi_client.models.workflowv1alpha1_flocker_volume_source.workflowv1alpha1FlockerVolumeSource(
                    dataset_name = '0', 
                    dataset_uuid = '0', ), 
                gce_persistent_disk = openapi_client.models.workflowv1alpha1_gce_persistent_disk_volume_source.workflowv1alpha1GCEPersistentDiskVolumeSource(
                    fs_type = '0', 
                    partition = 56, 
                    pd_name = '0', 
                    read_only = True, ), 
                git_repo = openapi_client.models.workflowv1alpha1_git_repo_volume_source.workflowv1alpha1GitRepoVolumeSource(
                    directory = '0', 
                    repository = '0', 
                    revision = '0', ), 
                glusterfs = openapi_client.models.workflowv1alpha1_glusterfs_volume_source.workflowv1alpha1GlusterfsVolumeSource(
                    endpoints = '0', 
                    path = '0', 
                    read_only = True, ), 
                host_path = openapi_client.models.workflowv1alpha1_host_path_volume_source.workflowv1alpha1HostPathVolumeSource(
                    path = '0', 
                    type = '0', ), 
                iscsi = openapi_client.models.workflowv1alpha1_iscsi_volume_source.workflowv1alpha1ISCSIVolumeSource(
                    chap_auth_discovery = True, 
                    chap_auth_session = True, 
                    fs_type = '0', 
                    initiator_name = '0', 
                    iqn = '0', 
                    iscsi_interface = '0', 
                    lun = 56, 
                    portals = [
                        '0'
                        ], 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    target_portal = '0', ), 
                nfs = openapi_client.models.workflowv1alpha1_nfs_volume_source.workflowv1alpha1NFSVolumeSource(
                    path = '0', 
                    read_only = True, 
                    server = '0', ), 
                persistent_volume_claim = openapi_client.models.workflowv1alpha1_persistent_volume_claim_volume_source.workflowv1alpha1PersistentVolumeClaimVolumeSource(
                    claim_name = '0', 
                    read_only = True, ), 
                photon_persistent_disk = openapi_client.models.workflowv1alpha1_photon_persistent_disk_volume_source.workflowv1alpha1PhotonPersistentDiskVolumeSource(
                    fs_type = '0', 
                    pd_id = '0', ), 
                portworx_volume = openapi_client.models.workflowv1alpha1_portworx_volume_source.workflowv1alpha1PortworxVolumeSource(
                    fs_type = '0', 
                    read_only = True, 
                    volume_id = '0', ), 
                projected = openapi_client.models.workflowv1alpha1_projected_volume_source.workflowv1alpha1ProjectedVolumeSource(
                    default_mode = 56, 
                    sources = [
                        openapi_client.models.workflowv1alpha1_volume_projection.workflowv1alpha1VolumeProjection(
                            config_map = openapi_client.models.workflowv1alpha1_config_map_projection.workflowv1alpha1ConfigMapProjection(
                                items = [
                                    openapi_client.models.workflowv1alpha1_key_to_path.workflowv1alpha1KeyToPath(
                                        key = '0', 
                                        mode = 56, 
                                        path = '0', )
                                    ], 
                                local_object_reference = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                                    name = '0', ), 
                                optional = True, ), 
                            downward_api = openapi_client.models.workflowv1alpha1_downward_api_projection.workflowv1alpha1DownwardAPIProjection(
                                items = [
                                    openapi_client.models.workflowv1alpha1_downward_api_volume_file.workflowv1alpha1DownwardAPIVolumeFile(
                                        field_ref = openapi_client.models.workflowv1alpha1_object_field_selector.workflowv1alpha1ObjectFieldSelector(
                                            api_version = '0', 
                                            field_path = '0', ), 
                                        mode = 56, 
                                        path = '0', 
                                        resource_field_ref = openapi_client.models.workflowv1alpha1_resource_field_selector.workflowv1alpha1ResourceFieldSelector(
                                            container_name = '0', 
                                            divisor = openapi_client.models.resource_quantity.resourceQuantity(
                                                string = '0', ), 
                                            resource = '0', ), )
                                    ], ), 
                            secret = openapi_client.models.workflowv1alpha1_secret_projection.workflowv1alpha1SecretProjection(
                                items = [
                                    openapi_client.models.workflowv1alpha1_key_to_path.workflowv1alpha1KeyToPath(
                                        key = '0', 
                                        mode = 56, 
                                        path = '0', )
                                    ], 
                                optional = True, ), 
                            service_account_token = openapi_client.models.workflowv1alpha1_service_account_token_projection.workflowv1alpha1ServiceAccountTokenProjection(
                                audience = '0', 
                                expiration_seconds = '0', 
                                path = '0', ), )
                        ], ), 
                quobyte = openapi_client.models.workflowv1alpha1_quobyte_volume_source.workflowv1alpha1QuobyteVolumeSource(
                    group = '0', 
                    read_only = True, 
                    registry = '0', 
                    tenant = '0', 
                    user = '0', 
                    volume = '0', ), 
                rbd = openapi_client.models.workflowv1alpha1_rbd_volume_source.workflowv1alpha1RBDVolumeSource(
                    fs_type = '0', 
                    image = '0', 
                    keyring = '0', 
                    monitors = [
                        '0'
                        ], 
                    pool = '0', 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    user = '0', ), 
                scale_io = openapi_client.models.workflowv1alpha1_scale_io_volume_source.workflowv1alpha1ScaleIOVolumeSource(
                    fs_type = '0', 
                    gateway = '0', 
                    protection_domain = '0', 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    ssl_enabled = True, 
                    storage_mode = '0', 
                    storage_pool = '0', 
                    system = '0', 
                    volume_name = '0', ), 
                secret = openapi_client.models.workflowv1alpha1_secret_volume_source.workflowv1alpha1SecretVolumeSource(
                    default_mode = 56, 
                    items = [
                        openapi_client.models.workflowv1alpha1_key_to_path.workflowv1alpha1KeyToPath(
                            key = '0', 
                            mode = 56, 
                            path = '0', )
                        ], 
                    optional = True, 
                    secret_name = '0', ), 
                storageos = openapi_client.models.workflowv1alpha1_storage_os_volume_source.workflowv1alpha1StorageOSVolumeSource(
                    fs_type = '0', 
                    read_only = True, 
                    secret_ref = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    volume_name = '0', 
                    volume_namespace = '0', ), 
                vsphere_volume = openapi_client.models.workflowv1alpha1_vsphere_virtual_disk_volume_source.workflowv1alpha1VsphereVirtualDiskVolumeSource(
                    fs_type = '0', 
                    storage_policy_id = '0', 
                    storage_policy_name = '0', 
                    volume_path = '0', )
            )
        else :
            return Workflowv1alpha1VolumeSource(
        )

    def testWorkflowv1alpha1VolumeSource(self):
        """Test Workflowv1alpha1VolumeSource"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
