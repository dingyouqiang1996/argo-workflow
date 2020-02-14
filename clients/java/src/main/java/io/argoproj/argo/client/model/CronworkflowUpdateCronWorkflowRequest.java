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


package io.argoproj.argo.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.argoproj.argo.client.model.V1alpha1CronWorkflow;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * CronworkflowUpdateCronWorkflowRequest
 */

public class CronworkflowUpdateCronWorkflowRequest {
  public static final String SERIALIZED_NAME_CRON_WORKFLOW = "cronWorkflow";
  @SerializedName(SERIALIZED_NAME_CRON_WORKFLOW)
  private V1alpha1CronWorkflow cronWorkflow;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NAMESPACE = "namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;


  public CronworkflowUpdateCronWorkflowRequest cronWorkflow(V1alpha1CronWorkflow cronWorkflow) {
    
    this.cronWorkflow = cronWorkflow;
    return this;
  }

   /**
   * Get cronWorkflow
   * @return cronWorkflow
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1CronWorkflow getCronWorkflow() {
    return cronWorkflow;
  }


  public void setCronWorkflow(V1alpha1CronWorkflow cronWorkflow) {
    this.cronWorkflow = cronWorkflow;
  }


  public CronworkflowUpdateCronWorkflowRequest name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public CronworkflowUpdateCronWorkflowRequest namespace(String namespace) {
    
    this.namespace = namespace;
    return this;
  }

   /**
   * Get namespace
   * @return namespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamespace() {
    return namespace;
  }


  public void setNamespace(String namespace) {
    this.namespace = namespace;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CronworkflowUpdateCronWorkflowRequest cronworkflowUpdateCronWorkflowRequest = (CronworkflowUpdateCronWorkflowRequest) o;
    return Objects.equals(this.cronWorkflow, cronworkflowUpdateCronWorkflowRequest.cronWorkflow) &&
        Objects.equals(this.name, cronworkflowUpdateCronWorkflowRequest.name) &&
        Objects.equals(this.namespace, cronworkflowUpdateCronWorkflowRequest.namespace);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cronWorkflow, name, namespace);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CronworkflowUpdateCronWorkflowRequest {\n");
    sb.append("    cronWorkflow: ").append(toIndentedString(cronWorkflow)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

