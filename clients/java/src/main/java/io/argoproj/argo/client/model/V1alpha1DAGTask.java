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
import io.argoproj.argo.client.model.V1alpha1Arguments;
import io.argoproj.argo.client.model.V1alpha1ContinueOn;
import io.argoproj.argo.client.model.V1alpha1Item;
import io.argoproj.argo.client.model.V1alpha1Sequence;
import io.argoproj.argo.client.model.V1alpha1TemplateRef;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1DAGTask
 */

public class V1alpha1DAGTask {
  public static final String SERIALIZED_NAME_ARGUMENTS = "arguments";
  @SerializedName(SERIALIZED_NAME_ARGUMENTS)
  private V1alpha1Arguments arguments;

  public static final String SERIALIZED_NAME_CONTINUE_ON = "continueOn";
  @SerializedName(SERIALIZED_NAME_CONTINUE_ON)
  private V1alpha1ContinueOn continueOn;

  public static final String SERIALIZED_NAME_DEPENDENCIES = "dependencies";
  @SerializedName(SERIALIZED_NAME_DEPENDENCIES)
  private List<String> dependencies = null;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ON_EXIT = "onExit";
  @SerializedName(SERIALIZED_NAME_ON_EXIT)
  private String onExit;

  public static final String SERIALIZED_NAME_TEMPLATE = "template";
  @SerializedName(SERIALIZED_NAME_TEMPLATE)
  private String template;

  public static final String SERIALIZED_NAME_TEMPLATE_REF = "templateRef";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_REF)
  private V1alpha1TemplateRef templateRef;

  public static final String SERIALIZED_NAME_WHEN = "when";
  @SerializedName(SERIALIZED_NAME_WHEN)
  private String when;

  public static final String SERIALIZED_NAME_WITH_ITEMS = "withItems";
  @SerializedName(SERIALIZED_NAME_WITH_ITEMS)
  private List<V1alpha1Item> withItems = null;

  public static final String SERIALIZED_NAME_WITH_PARAM = "withParam";
  @SerializedName(SERIALIZED_NAME_WITH_PARAM)
  private String withParam;

  public static final String SERIALIZED_NAME_WITH_SEQUENCE = "withSequence";
  @SerializedName(SERIALIZED_NAME_WITH_SEQUENCE)
  private V1alpha1Sequence withSequence;


  public V1alpha1DAGTask arguments(V1alpha1Arguments arguments) {
    
    this.arguments = arguments;
    return this;
  }

   /**
   * Get arguments
   * @return arguments
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1Arguments getArguments() {
    return arguments;
  }


  public void setArguments(V1alpha1Arguments arguments) {
    this.arguments = arguments;
  }


  public V1alpha1DAGTask continueOn(V1alpha1ContinueOn continueOn) {
    
    this.continueOn = continueOn;
    return this;
  }

   /**
   * Get continueOn
   * @return continueOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1ContinueOn getContinueOn() {
    return continueOn;
  }


  public void setContinueOn(V1alpha1ContinueOn continueOn) {
    this.continueOn = continueOn;
  }


  public V1alpha1DAGTask dependencies(List<String> dependencies) {
    
    this.dependencies = dependencies;
    return this;
  }

  public V1alpha1DAGTask addDependenciesItem(String dependenciesItem) {
    if (this.dependencies == null) {
      this.dependencies = new ArrayList<String>();
    }
    this.dependencies.add(dependenciesItem);
    return this;
  }

   /**
   * Get dependencies
   * @return dependencies
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDependencies() {
    return dependencies;
  }


  public void setDependencies(List<String> dependencies) {
    this.dependencies = dependencies;
  }


  public V1alpha1DAGTask name(String name) {
    
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


  public V1alpha1DAGTask onExit(String onExit) {
    
    this.onExit = onExit;
    return this;
  }

   /**
   * OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template.
   * @return onExit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template.")

  public String getOnExit() {
    return onExit;
  }


  public void setOnExit(String onExit) {
    this.onExit = onExit;
  }


  public V1alpha1DAGTask template(String template) {
    
    this.template = template;
    return this;
  }

   /**
   * Get template
   * @return template
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTemplate() {
    return template;
  }


  public void setTemplate(String template) {
    this.template = template;
  }


  public V1alpha1DAGTask templateRef(V1alpha1TemplateRef templateRef) {
    
    this.templateRef = templateRef;
    return this;
  }

   /**
   * Get templateRef
   * @return templateRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1TemplateRef getTemplateRef() {
    return templateRef;
  }


  public void setTemplateRef(V1alpha1TemplateRef templateRef) {
    this.templateRef = templateRef;
  }


  public V1alpha1DAGTask when(String when) {
    
    this.when = when;
    return this;
  }

   /**
   * Get when
   * @return when
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getWhen() {
    return when;
  }


  public void setWhen(String when) {
    this.when = when;
  }


  public V1alpha1DAGTask withItems(List<V1alpha1Item> withItems) {
    
    this.withItems = withItems;
    return this;
  }

  public V1alpha1DAGTask addWithItemsItem(V1alpha1Item withItemsItem) {
    if (this.withItems == null) {
      this.withItems = new ArrayList<V1alpha1Item>();
    }
    this.withItems.add(withItemsItem);
    return this;
  }

   /**
   * Get withItems
   * @return withItems
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1alpha1Item> getWithItems() {
    return withItems;
  }


  public void setWithItems(List<V1alpha1Item> withItems) {
    this.withItems = withItems;
  }


  public V1alpha1DAGTask withParam(String withParam) {
    
    this.withParam = withParam;
    return this;
  }

   /**
   * WithParam expands a task into multiple parallel tasks from the value in the parameter, which is expected to be a JSON list.
   * @return withParam
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "WithParam expands a task into multiple parallel tasks from the value in the parameter, which is expected to be a JSON list.")

  public String getWithParam() {
    return withParam;
  }


  public void setWithParam(String withParam) {
    this.withParam = withParam;
  }


  public V1alpha1DAGTask withSequence(V1alpha1Sequence withSequence) {
    
    this.withSequence = withSequence;
    return this;
  }

   /**
   * Get withSequence
   * @return withSequence
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1Sequence getWithSequence() {
    return withSequence;
  }


  public void setWithSequence(V1alpha1Sequence withSequence) {
    this.withSequence = withSequence;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1DAGTask v1alpha1DAGTask = (V1alpha1DAGTask) o;
    return Objects.equals(this.arguments, v1alpha1DAGTask.arguments) &&
        Objects.equals(this.continueOn, v1alpha1DAGTask.continueOn) &&
        Objects.equals(this.dependencies, v1alpha1DAGTask.dependencies) &&
        Objects.equals(this.name, v1alpha1DAGTask.name) &&
        Objects.equals(this.onExit, v1alpha1DAGTask.onExit) &&
        Objects.equals(this.template, v1alpha1DAGTask.template) &&
        Objects.equals(this.templateRef, v1alpha1DAGTask.templateRef) &&
        Objects.equals(this.when, v1alpha1DAGTask.when) &&
        Objects.equals(this.withItems, v1alpha1DAGTask.withItems) &&
        Objects.equals(this.withParam, v1alpha1DAGTask.withParam) &&
        Objects.equals(this.withSequence, v1alpha1DAGTask.withSequence);
  }

  @Override
  public int hashCode() {
    return Objects.hash(arguments, continueOn, dependencies, name, onExit, template, templateRef, when, withItems, withParam, withSequence);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1DAGTask {\n");
    sb.append("    arguments: ").append(toIndentedString(arguments)).append("\n");
    sb.append("    continueOn: ").append(toIndentedString(continueOn)).append("\n");
    sb.append("    dependencies: ").append(toIndentedString(dependencies)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    onExit: ").append(toIndentedString(onExit)).append("\n");
    sb.append("    template: ").append(toIndentedString(template)).append("\n");
    sb.append("    templateRef: ").append(toIndentedString(templateRef)).append("\n");
    sb.append("    when: ").append(toIndentedString(when)).append("\n");
    sb.append("    withItems: ").append(toIndentedString(withItems)).append("\n");
    sb.append("    withParam: ").append(toIndentedString(withParam)).append("\n");
    sb.append("    withSequence: ").append(toIndentedString(withSequence)).append("\n");
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

