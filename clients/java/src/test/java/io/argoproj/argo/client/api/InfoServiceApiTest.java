/*
 * Argo
 * Workflow Service API performs CRUD actions against application resources
 *
 * The version of the OpenAPI document: latest
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.argoproj.argo.client.api;

import io.argoproj.argo.client.ApiException;
import io.argoproj.argo.client.model.InfoInfoResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for InfoServiceApi
 */
@Ignore
public class InfoServiceApiTest {

    private final InfoServiceApi api = new InfoServiceApi();

    
    /**
     * 
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInfoTest() throws ApiException {
        InfoInfoResponse response = api.getInfo();

        // TODO: test validations
    }
    
}
