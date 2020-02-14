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
import io.argoproj.argo.client.model.V1PodDNSConfigOption;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
 */
@ApiModel(description = "PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.")

public class V1PodDNSConfig {
  public static final String SERIALIZED_NAME_NAMESERVERS = "nameservers";
  @SerializedName(SERIALIZED_NAME_NAMESERVERS)
  private List<String> nameservers = null;

  public static final String SERIALIZED_NAME_OPTIONS = "options";
  @SerializedName(SERIALIZED_NAME_OPTIONS)
  private List<V1PodDNSConfigOption> options = null;

  public static final String SERIALIZED_NAME_SEARCHES = "searches";
  @SerializedName(SERIALIZED_NAME_SEARCHES)
  private List<String> searches = null;


  public V1PodDNSConfig nameservers(List<String> nameservers) {
    
    this.nameservers = nameservers;
    return this;
  }

  public V1PodDNSConfig addNameserversItem(String nameserversItem) {
    if (this.nameservers == null) {
      this.nameservers = new ArrayList<String>();
    }
    this.nameservers.add(nameserversItem);
    return this;
  }

   /**
   * Get nameservers
   * @return nameservers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getNameservers() {
    return nameservers;
  }


  public void setNameservers(List<String> nameservers) {
    this.nameservers = nameservers;
  }


  public V1PodDNSConfig options(List<V1PodDNSConfigOption> options) {
    
    this.options = options;
    return this;
  }

  public V1PodDNSConfig addOptionsItem(V1PodDNSConfigOption optionsItem) {
    if (this.options == null) {
      this.options = new ArrayList<V1PodDNSConfigOption>();
    }
    this.options.add(optionsItem);
    return this;
  }

   /**
   * Get options
   * @return options
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1PodDNSConfigOption> getOptions() {
    return options;
  }


  public void setOptions(List<V1PodDNSConfigOption> options) {
    this.options = options;
  }


  public V1PodDNSConfig searches(List<String> searches) {
    
    this.searches = searches;
    return this;
  }

  public V1PodDNSConfig addSearchesItem(String searchesItem) {
    if (this.searches == null) {
      this.searches = new ArrayList<String>();
    }
    this.searches.add(searchesItem);
    return this;
  }

   /**
   * Get searches
   * @return searches
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getSearches() {
    return searches;
  }


  public void setSearches(List<String> searches) {
    this.searches = searches;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1PodDNSConfig v1PodDNSConfig = (V1PodDNSConfig) o;
    return Objects.equals(this.nameservers, v1PodDNSConfig.nameservers) &&
        Objects.equals(this.options, v1PodDNSConfig.options) &&
        Objects.equals(this.searches, v1PodDNSConfig.searches);
  }

  @Override
  public int hashCode() {
    return Objects.hash(nameservers, options, searches);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1PodDNSConfig {\n");
    sb.append("    nameservers: ").append(toIndentedString(nameservers)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
    sb.append("    searches: ").append(toIndentedString(searches)).append("\n");
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

