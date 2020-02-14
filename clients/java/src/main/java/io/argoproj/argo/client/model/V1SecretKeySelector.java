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
import io.argoproj.argo.client.model.V1LocalObjectReference;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * SecretKeySelector selects a key of a Secret.
 */
@ApiModel(description = "SecretKeySelector selects a key of a Secret.")

public class V1SecretKeySelector {
  public static final String SERIALIZED_NAME_KEY = "key";
  @SerializedName(SERIALIZED_NAME_KEY)
  private String key;

  public static final String SERIALIZED_NAME_LOCAL_OBJECT_REFERENCE = "localObjectReference";
  @SerializedName(SERIALIZED_NAME_LOCAL_OBJECT_REFERENCE)
  private V1LocalObjectReference localObjectReference;

  public static final String SERIALIZED_NAME_OPTIONAL = "optional";
  @SerializedName(SERIALIZED_NAME_OPTIONAL)
  private Boolean optional;


  public V1SecretKeySelector key(String key) {
    
    this.key = key;
    return this;
  }

   /**
   * The key of the secret to select from.  Must be a valid secret key.
   * @return key
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The key of the secret to select from.  Must be a valid secret key.")

  public String getKey() {
    return key;
  }


  public void setKey(String key) {
    this.key = key;
  }


  public V1SecretKeySelector localObjectReference(V1LocalObjectReference localObjectReference) {
    
    this.localObjectReference = localObjectReference;
    return this;
  }

   /**
   * Get localObjectReference
   * @return localObjectReference
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1LocalObjectReference getLocalObjectReference() {
    return localObjectReference;
  }


  public void setLocalObjectReference(V1LocalObjectReference localObjectReference) {
    this.localObjectReference = localObjectReference;
  }


  public V1SecretKeySelector optional(Boolean optional) {
    
    this.optional = optional;
    return this;
  }

   /**
   * Get optional
   * @return optional
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getOptional() {
    return optional;
  }


  public void setOptional(Boolean optional) {
    this.optional = optional;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1SecretKeySelector v1SecretKeySelector = (V1SecretKeySelector) o;
    return Objects.equals(this.key, v1SecretKeySelector.key) &&
        Objects.equals(this.localObjectReference, v1SecretKeySelector.localObjectReference) &&
        Objects.equals(this.optional, v1SecretKeySelector.optional);
  }

  @Override
  public int hashCode() {
    return Objects.hash(key, localObjectReference, optional);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1SecretKeySelector {\n");
    sb.append("    key: ").append(toIndentedString(key)).append("\n");
    sb.append("    localObjectReference: ").append(toIndentedString(localObjectReference)).append("\n");
    sb.append("    optional: ").append(toIndentedString(optional)).append("\n");
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

