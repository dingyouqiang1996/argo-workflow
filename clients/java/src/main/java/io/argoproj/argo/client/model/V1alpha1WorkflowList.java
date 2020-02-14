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
import io.argoproj.argo.client.model.V1ListMeta;
import io.argoproj.argo.client.model.V1alpha1Workflow;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1WorkflowList
 */

public class V1alpha1WorkflowList {
  public static final String SERIALIZED_NAME_ITEMS = "items";
  @SerializedName(SERIALIZED_NAME_ITEMS)
  private List<V1alpha1Workflow> items = null;

  public static final String SERIALIZED_NAME_METADATA = "metadata";
  @SerializedName(SERIALIZED_NAME_METADATA)
  private V1ListMeta metadata;


  public V1alpha1WorkflowList items(List<V1alpha1Workflow> items) {
    
    this.items = items;
    return this;
  }

  public V1alpha1WorkflowList addItemsItem(V1alpha1Workflow itemsItem) {
    if (this.items == null) {
      this.items = new ArrayList<V1alpha1Workflow>();
    }
    this.items.add(itemsItem);
    return this;
  }

   /**
   * Get items
   * @return items
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1alpha1Workflow> getItems() {
    return items;
  }


  public void setItems(List<V1alpha1Workflow> items) {
    this.items = items;
  }


  public V1alpha1WorkflowList metadata(V1ListMeta metadata) {
    
    this.metadata = metadata;
    return this;
  }

   /**
   * Get metadata
   * @return metadata
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1ListMeta getMetadata() {
    return metadata;
  }


  public void setMetadata(V1ListMeta metadata) {
    this.metadata = metadata;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1WorkflowList v1alpha1WorkflowList = (V1alpha1WorkflowList) o;
    return Objects.equals(this.items, v1alpha1WorkflowList.items) &&
        Objects.equals(this.metadata, v1alpha1WorkflowList.metadata);
  }

  @Override
  public int hashCode() {
    return Objects.hash(items, metadata);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1WorkflowList {\n");
    sb.append("    items: ").append(toIndentedString(items)).append("\n");
    sb.append("    metadata: ").append(toIndentedString(metadata)).append("\n");
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

