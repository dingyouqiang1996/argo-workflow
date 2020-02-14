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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * InfoInfoResponse
 */

public class InfoInfoResponse {
  public static final String SERIALIZED_NAME_MANAGED_NAMESPACE = "managedNamespace";
  @SerializedName(SERIALIZED_NAME_MANAGED_NAMESPACE)
  private String managedNamespace;


  public InfoInfoResponse managedNamespace(String managedNamespace) {
    
    this.managedNamespace = managedNamespace;
    return this;
  }

   /**
   * Get managedNamespace
   * @return managedNamespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getManagedNamespace() {
    return managedNamespace;
  }


  public void setManagedNamespace(String managedNamespace) {
    this.managedNamespace = managedNamespace;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InfoInfoResponse infoInfoResponse = (InfoInfoResponse) o;
    return Objects.equals(this.managedNamespace, infoInfoResponse.managedNamespace);
  }

  @Override
  public int hashCode() {
    return Objects.hash(managedNamespace);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InfoInfoResponse {\n");
    sb.append("    managedNamespace: ").append(toIndentedString(managedNamespace)).append("\n");
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

