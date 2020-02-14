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
from openapi_client.io.argoproj.argo.client.model.workflowv1alpha1_persistent_volume_claim import Workflowv1alpha1PersistentVolumeClaim  # noqa: E501
from openapi_client.rest import ApiException

class TestWorkflowv1alpha1PersistentVolumeClaim(unittest.TestCase):
    """Workflowv1alpha1PersistentVolumeClaim unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test Workflowv1alpha1PersistentVolumeClaim
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.workflowv1alpha1_persistent_volume_claim.Workflowv1alpha1PersistentVolumeClaim()  # noqa: E501
        if include_optional :
            return Workflowv1alpha1PersistentVolumeClaim(
                metadata = openapi_client.models.v1_object_meta.v1ObjectMeta(
                    annotations = {
                        'key' : '0'
                        }, 
                    cluster_name = '0', 
                    creation_timestamp = openapi_client.models.v1_time.v1Time(
                        nanos = 56, 
                        seconds = '0', ), 
                    deletion_grace_period_seconds = '0', 
                    deletion_timestamp = openapi_client.models.v1_time.v1Time(
                        nanos = 56, 
                        seconds = '0', ), 
                    finalizers = [
                        '0'
                        ], 
                    generate_name = '0', 
                    generation = '0', 
                    labels = {
                        'key' : '0'
                        }, 
                    managed_fields = [
                        openapi_client.models.v1_managed_fields_entry.v1ManagedFieldsEntry(
                            api_version = '0', 
                            fields_type = '0', 
                            fields_v1 = openapi_client.models.v1_fields_v1.v1FieldsV1(
                                raw = 'YQ==', ), 
                            manager = '0', 
                            operation = '0', 
                            time = openapi_client.models.v1_time.v1Time(
                                nanos = 56, 
                                seconds = '0', ), )
                        ], 
                    name = '0', 
                    namespace = '0', 
                    owner_references = [
                        openapi_client.models.v1_owner_reference.v1OwnerReference(
                            api_version = '0', 
                            block_owner_deletion = True, 
                            controller = True, 
                            kind = '0', 
                            name = '0', 
                            uid = '0', )
                        ], 
                    resource_version = '0', 
                    self_link = '0', 
                    uid = '0', ), 
                spec = openapi_client.models.workflowv1alpha1_persistent_volume_claim_spec.workflowv1alpha1PersistentVolumeClaimSpec(
                    access_modes = [
                        '0'
                        ], 
                    data_source = openapi_client.models.workflowv1alpha1_typed_local_object_reference.workflowv1alpha1TypedLocalObjectReference(
                        api_group = '0', 
                        kind = '0', 
                        name = '0', ), 
                    resources = openapi_client.models.workflowv1alpha1_resource_requirements.workflowv1alpha1ResourceRequirements(
                        limits = {
                            'key' : openapi_client.models.resource_quantity.resourceQuantity(
                                string = '0', )
                            }, 
                        requests = {
                            'key' : openapi_client.models.resource_quantity.resourceQuantity(
                                string = '0', )
                            }, ), 
                    selector = openapi_client.models.v1_label_selector.v1LabelSelector(
                        match_expressions = [
                            openapi_client.models.v1_label_selector_requirement.v1LabelSelectorRequirement(
                                key = '0', 
                                operator = '0', 
                                values = [
                                    '0'
                                    ], )
                            ], 
                        match_labels = {
                            'key' : '0'
                            }, ), 
                    storage_class_name = '0', 
                    volume_mode = '0', 
                    volume_name = '0', ), 
                status = openapi_client.models.workflowv1alpha1_persistent_volume_claim_status.workflowv1alpha1PersistentVolumeClaimStatus(
                    access_modes = [
                        '0'
                        ], 
                    capacity = {
                        'key' : openapi_client.models.resource_quantity.resourceQuantity(
                            string = '0', )
                        }, 
                    conditions = [
                        openapi_client.models.workflowv1alpha1_persistent_volume_claim_condition.workflowv1alpha1PersistentVolumeClaimCondition(
                            last_probe_time = openapi_client.models.v1_time.v1Time(
                                nanos = 56, 
                                seconds = '0', ), 
                            last_transition_time = openapi_client.models.v1_time.v1Time(
                                nanos = 56, 
                                seconds = '0', ), 
                            message = '0', 
                            reason = '0', 
                            type = '0', )
                        ], 
                    phase = '0', ), 
                type_meta = openapi_client.models.metav1_type_meta.metav1TypeMeta(
                    api_version = '0', 
                    kind = '0', )
            )
        else :
            return Workflowv1alpha1PersistentVolumeClaim(
        )

    def testWorkflowv1alpha1PersistentVolumeClaim(self):
        """Test Workflowv1alpha1PersistentVolumeClaim"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
