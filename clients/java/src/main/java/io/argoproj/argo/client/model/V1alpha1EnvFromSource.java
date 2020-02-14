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
import io.argoproj.argo.client.model.V1alpha1ConfigMapEnvSource;
import io.argoproj.argo.client.model.V1alpha1SecretEnvSource;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * V1alpha1EnvFromSource
 */

public class V1alpha1EnvFromSource {
  public static final String SERIALIZED_NAME_CONFIG_MAP_REF = "configMapRef";
  @SerializedName(SERIALIZED_NAME_CONFIG_MAP_REF)
  private V1alpha1ConfigMapEnvSource configMapRef;

  public static final String SERIALIZED_NAME_PREFIX = "prefix";
  @SerializedName(SERIALIZED_NAME_PREFIX)
  private String prefix;

  public static final String SERIALIZED_NAME_SECRET_REF = "secretRef";
  @SerializedName(SERIALIZED_NAME_SECRET_REF)
  private V1alpha1SecretEnvSource secretRef;


  public V1alpha1EnvFromSource configMapRef(V1alpha1ConfigMapEnvSource configMapRef) {
    
    this.configMapRef = configMapRef;
    return this;
  }

   /**
   * Get configMapRef
   * @return configMapRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1ConfigMapEnvSource getConfigMapRef() {
    return configMapRef;
  }


  public void setConfigMapRef(V1alpha1ConfigMapEnvSource configMapRef) {
    this.configMapRef = configMapRef;
  }


  public V1alpha1EnvFromSource prefix(String prefix) {
    
    this.prefix = prefix;
    return this;
  }

   /**
   * Get prefix
   * @return prefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPrefix() {
    return prefix;
  }


  public void setPrefix(String prefix) {
    this.prefix = prefix;
  }


  public V1alpha1EnvFromSource secretRef(V1alpha1SecretEnvSource secretRef) {
    
    this.secretRef = secretRef;
    return this;
  }

   /**
   * Get secretRef
   * @return secretRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1SecretEnvSource getSecretRef() {
    return secretRef;
  }


  public void setSecretRef(V1alpha1SecretEnvSource secretRef) {
    this.secretRef = secretRef;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1EnvFromSource v1alpha1EnvFromSource = (V1alpha1EnvFromSource) o;
    return Objects.equals(this.configMapRef, v1alpha1EnvFromSource.configMapRef) &&
        Objects.equals(this.prefix, v1alpha1EnvFromSource.prefix) &&
        Objects.equals(this.secretRef, v1alpha1EnvFromSource.secretRef);
  }

  @Override
  public int hashCode() {
    return Objects.hash(configMapRef, prefix, secretRef);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1EnvFromSource {\n");
    sb.append("    configMapRef: ").append(toIndentedString(configMapRef)).append("\n");
    sb.append("    prefix: ").append(toIndentedString(prefix)).append("\n");
    sb.append("    secretRef: ").append(toIndentedString(secretRef)).append("\n");
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

