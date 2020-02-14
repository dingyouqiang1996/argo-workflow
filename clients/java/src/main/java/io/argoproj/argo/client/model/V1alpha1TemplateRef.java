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
 * TemplateRef is a reference of template resource.
 */
@ApiModel(description = "TemplateRef is a reference of template resource.")

public class V1alpha1TemplateRef {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_RUNTIME_RESOLUTION = "runtimeResolution";
  @SerializedName(SERIALIZED_NAME_RUNTIME_RESOLUTION)
  private Boolean runtimeResolution;

  public static final String SERIALIZED_NAME_TEMPLATE = "template";
  @SerializedName(SERIALIZED_NAME_TEMPLATE)
  private String template;


  public V1alpha1TemplateRef name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name is the resource name of the template.
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Name is the resource name of the template.")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public V1alpha1TemplateRef runtimeResolution(Boolean runtimeResolution) {
    
    this.runtimeResolution = runtimeResolution;
    return this;
  }

   /**
   * RuntimeResolution skips validation at creation time. By enabling this option, you can create the referred workflow template before the actual runtime.
   * @return runtimeResolution
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "RuntimeResolution skips validation at creation time. By enabling this option, you can create the referred workflow template before the actual runtime.")

  public Boolean getRuntimeResolution() {
    return runtimeResolution;
  }


  public void setRuntimeResolution(Boolean runtimeResolution) {
    this.runtimeResolution = runtimeResolution;
  }


  public V1alpha1TemplateRef template(String template) {
    
    this.template = template;
    return this;
  }

   /**
   * Template is the name of referred template in the resource.
   * @return template
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Template is the name of referred template in the resource.")

  public String getTemplate() {
    return template;
  }


  public void setTemplate(String template) {
    this.template = template;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1TemplateRef v1alpha1TemplateRef = (V1alpha1TemplateRef) o;
    return Objects.equals(this.name, v1alpha1TemplateRef.name) &&
        Objects.equals(this.runtimeResolution, v1alpha1TemplateRef.runtimeResolution) &&
        Objects.equals(this.template, v1alpha1TemplateRef.template);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, runtimeResolution, template);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1TemplateRef {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    runtimeResolution: ").append(toIndentedString(runtimeResolution)).append("\n");
    sb.append("    template: ").append(toIndentedString(template)).append("\n");
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

