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
import io.argoproj.argo.client.model.V1alpha1Artifact;
import io.argoproj.argo.client.model.V1alpha1Parameter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1Arguments
 */

public class V1alpha1Arguments {
  public static final String SERIALIZED_NAME_ARTIFACTS = "artifacts";
  @SerializedName(SERIALIZED_NAME_ARTIFACTS)
  private List<V1alpha1Artifact> artifacts = null;

  public static final String SERIALIZED_NAME_PARAMETERS = "parameters";
  @SerializedName(SERIALIZED_NAME_PARAMETERS)
  private List<V1alpha1Parameter> parameters = null;


  public V1alpha1Arguments artifacts(List<V1alpha1Artifact> artifacts) {
    
    this.artifacts = artifacts;
    return this;
  }

  public V1alpha1Arguments addArtifactsItem(V1alpha1Artifact artifactsItem) {
    if (this.artifacts == null) {
      this.artifacts = new ArrayList<V1alpha1Artifact>();
    }
    this.artifacts.add(artifactsItem);
    return this;
  }

   /**
   * Get artifacts
   * @return artifacts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1alpha1Artifact> getArtifacts() {
    return artifacts;
  }


  public void setArtifacts(List<V1alpha1Artifact> artifacts) {
    this.artifacts = artifacts;
  }


  public V1alpha1Arguments parameters(List<V1alpha1Parameter> parameters) {
    
    this.parameters = parameters;
    return this;
  }

  public V1alpha1Arguments addParametersItem(V1alpha1Parameter parametersItem) {
    if (this.parameters == null) {
      this.parameters = new ArrayList<V1alpha1Parameter>();
    }
    this.parameters.add(parametersItem);
    return this;
  }

   /**
   * Get parameters
   * @return parameters
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1alpha1Parameter> getParameters() {
    return parameters;
  }


  public void setParameters(List<V1alpha1Parameter> parameters) {
    this.parameters = parameters;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1Arguments v1alpha1Arguments = (V1alpha1Arguments) o;
    return Objects.equals(this.artifacts, v1alpha1Arguments.artifacts) &&
        Objects.equals(this.parameters, v1alpha1Arguments.parameters);
  }

  @Override
  public int hashCode() {
    return Objects.hash(artifacts, parameters);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1Arguments {\n");
    sb.append("    artifacts: ").append(toIndentedString(artifacts)).append("\n");
    sb.append("    parameters: ").append(toIndentedString(parameters)).append("\n");
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

