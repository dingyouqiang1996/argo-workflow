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
import io.argoproj.argo.client.model.Workflowv1alpha1NodeSelectorRequirement;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Workflowv1alpha1NodeSelectorTerm
 */

public class Workflowv1alpha1NodeSelectorTerm {
  public static final String SERIALIZED_NAME_MATCH_EXPRESSIONS = "matchExpressions";
  @SerializedName(SERIALIZED_NAME_MATCH_EXPRESSIONS)
  private List<Workflowv1alpha1NodeSelectorRequirement> matchExpressions = null;

  public static final String SERIALIZED_NAME_MATCH_FIELDS = "matchFields";
  @SerializedName(SERIALIZED_NAME_MATCH_FIELDS)
  private List<Workflowv1alpha1NodeSelectorRequirement> matchFields = null;


  public Workflowv1alpha1NodeSelectorTerm matchExpressions(List<Workflowv1alpha1NodeSelectorRequirement> matchExpressions) {
    
    this.matchExpressions = matchExpressions;
    return this;
  }

  public Workflowv1alpha1NodeSelectorTerm addMatchExpressionsItem(Workflowv1alpha1NodeSelectorRequirement matchExpressionsItem) {
    if (this.matchExpressions == null) {
      this.matchExpressions = new ArrayList<Workflowv1alpha1NodeSelectorRequirement>();
    }
    this.matchExpressions.add(matchExpressionsItem);
    return this;
  }

   /**
   * Get matchExpressions
   * @return matchExpressions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Workflowv1alpha1NodeSelectorRequirement> getMatchExpressions() {
    return matchExpressions;
  }


  public void setMatchExpressions(List<Workflowv1alpha1NodeSelectorRequirement> matchExpressions) {
    this.matchExpressions = matchExpressions;
  }


  public Workflowv1alpha1NodeSelectorTerm matchFields(List<Workflowv1alpha1NodeSelectorRequirement> matchFields) {
    
    this.matchFields = matchFields;
    return this;
  }

  public Workflowv1alpha1NodeSelectorTerm addMatchFieldsItem(Workflowv1alpha1NodeSelectorRequirement matchFieldsItem) {
    if (this.matchFields == null) {
      this.matchFields = new ArrayList<Workflowv1alpha1NodeSelectorRequirement>();
    }
    this.matchFields.add(matchFieldsItem);
    return this;
  }

   /**
   * Get matchFields
   * @return matchFields
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Workflowv1alpha1NodeSelectorRequirement> getMatchFields() {
    return matchFields;
  }


  public void setMatchFields(List<Workflowv1alpha1NodeSelectorRequirement> matchFields) {
    this.matchFields = matchFields;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Workflowv1alpha1NodeSelectorTerm workflowv1alpha1NodeSelectorTerm = (Workflowv1alpha1NodeSelectorTerm) o;
    return Objects.equals(this.matchExpressions, workflowv1alpha1NodeSelectorTerm.matchExpressions) &&
        Objects.equals(this.matchFields, workflowv1alpha1NodeSelectorTerm.matchFields);
  }

  @Override
  public int hashCode() {
    return Objects.hash(matchExpressions, matchFields);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Workflowv1alpha1NodeSelectorTerm {\n");
    sb.append("    matchExpressions: ").append(toIndentedString(matchExpressions)).append("\n");
    sb.append("    matchFields: ").append(toIndentedString(matchFields)).append("\n");
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

