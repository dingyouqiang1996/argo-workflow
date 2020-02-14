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
import io.argoproj.argo.client.model.ResourceQuantity;
import io.argoproj.argo.client.model.V1PersistentVolumeClaimCondition;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * PersistentVolumeClaimStatus is the current status of a persistent volume claim.
 */
@ApiModel(description = "PersistentVolumeClaimStatus is the current status of a persistent volume claim.")

public class V1PersistentVolumeClaimStatus {
  public static final String SERIALIZED_NAME_ACCESS_MODES = "accessModes";
  @SerializedName(SERIALIZED_NAME_ACCESS_MODES)
  private List<String> accessModes = null;

  public static final String SERIALIZED_NAME_CAPACITY = "capacity";
  @SerializedName(SERIALIZED_NAME_CAPACITY)
  private Map<String, ResourceQuantity> capacity = null;

  public static final String SERIALIZED_NAME_CONDITIONS = "conditions";
  @SerializedName(SERIALIZED_NAME_CONDITIONS)
  private List<V1PersistentVolumeClaimCondition> conditions = null;

  public static final String SERIALIZED_NAME_PHASE = "phase";
  @SerializedName(SERIALIZED_NAME_PHASE)
  private String phase;


  public V1PersistentVolumeClaimStatus accessModes(List<String> accessModes) {
    
    this.accessModes = accessModes;
    return this;
  }

  public V1PersistentVolumeClaimStatus addAccessModesItem(String accessModesItem) {
    if (this.accessModes == null) {
      this.accessModes = new ArrayList<String>();
    }
    this.accessModes.add(accessModesItem);
    return this;
  }

   /**
   * Get accessModes
   * @return accessModes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getAccessModes() {
    return accessModes;
  }


  public void setAccessModes(List<String> accessModes) {
    this.accessModes = accessModes;
  }


  public V1PersistentVolumeClaimStatus capacity(Map<String, ResourceQuantity> capacity) {
    
    this.capacity = capacity;
    return this;
  }

  public V1PersistentVolumeClaimStatus putCapacityItem(String key, ResourceQuantity capacityItem) {
    if (this.capacity == null) {
      this.capacity = new HashMap<String, ResourceQuantity>();
    }
    this.capacity.put(key, capacityItem);
    return this;
  }

   /**
   * Get capacity
   * @return capacity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, ResourceQuantity> getCapacity() {
    return capacity;
  }


  public void setCapacity(Map<String, ResourceQuantity> capacity) {
    this.capacity = capacity;
  }


  public V1PersistentVolumeClaimStatus conditions(List<V1PersistentVolumeClaimCondition> conditions) {
    
    this.conditions = conditions;
    return this;
  }

  public V1PersistentVolumeClaimStatus addConditionsItem(V1PersistentVolumeClaimCondition conditionsItem) {
    if (this.conditions == null) {
      this.conditions = new ArrayList<V1PersistentVolumeClaimCondition>();
    }
    this.conditions.add(conditionsItem);
    return this;
  }

   /**
   * Get conditions
   * @return conditions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1PersistentVolumeClaimCondition> getConditions() {
    return conditions;
  }


  public void setConditions(List<V1PersistentVolumeClaimCondition> conditions) {
    this.conditions = conditions;
  }


  public V1PersistentVolumeClaimStatus phase(String phase) {
    
    this.phase = phase;
    return this;
  }

   /**
   * Get phase
   * @return phase
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPhase() {
    return phase;
  }


  public void setPhase(String phase) {
    this.phase = phase;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1PersistentVolumeClaimStatus v1PersistentVolumeClaimStatus = (V1PersistentVolumeClaimStatus) o;
    return Objects.equals(this.accessModes, v1PersistentVolumeClaimStatus.accessModes) &&
        Objects.equals(this.capacity, v1PersistentVolumeClaimStatus.capacity) &&
        Objects.equals(this.conditions, v1PersistentVolumeClaimStatus.conditions) &&
        Objects.equals(this.phase, v1PersistentVolumeClaimStatus.phase);
  }

  @Override
  public int hashCode() {
    return Objects.hash(accessModes, capacity, conditions, phase);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1PersistentVolumeClaimStatus {\n");
    sb.append("    accessModes: ").append(toIndentedString(accessModes)).append("\n");
    sb.append("    capacity: ").append(toIndentedString(capacity)).append("\n");
    sb.append("    conditions: ").append(toIndentedString(conditions)).append("\n");
    sb.append("    phase: ").append(toIndentedString(phase)).append("\n");
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

