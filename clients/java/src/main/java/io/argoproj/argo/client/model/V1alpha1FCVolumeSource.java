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
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1FCVolumeSource
 */

public class V1alpha1FCVolumeSource {
  public static final String SERIALIZED_NAME_FS_TYPE = "fsType";
  @SerializedName(SERIALIZED_NAME_FS_TYPE)
  private String fsType;

  public static final String SERIALIZED_NAME_LUN = "lun";
  @SerializedName(SERIALIZED_NAME_LUN)
  private Integer lun;

  public static final String SERIALIZED_NAME_READ_ONLY = "readOnly";
  @SerializedName(SERIALIZED_NAME_READ_ONLY)
  private Boolean readOnly;

  public static final String SERIALIZED_NAME_TARGET_W_W_NS = "targetWWNs";
  @SerializedName(SERIALIZED_NAME_TARGET_W_W_NS)
  private List<String> targetWWNs = null;

  public static final String SERIALIZED_NAME_WWIDS = "wwids";
  @SerializedName(SERIALIZED_NAME_WWIDS)
  private List<String> wwids = null;


  public V1alpha1FCVolumeSource fsType(String fsType) {
    
    this.fsType = fsType;
    return this;
  }

   /**
   * Get fsType
   * @return fsType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFsType() {
    return fsType;
  }


  public void setFsType(String fsType) {
    this.fsType = fsType;
  }


  public V1alpha1FCVolumeSource lun(Integer lun) {
    
    this.lun = lun;
    return this;
  }

   /**
   * Get lun
   * @return lun
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getLun() {
    return lun;
  }


  public void setLun(Integer lun) {
    this.lun = lun;
  }


  public V1alpha1FCVolumeSource readOnly(Boolean readOnly) {
    
    this.readOnly = readOnly;
    return this;
  }

   /**
   * Get readOnly
   * @return readOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getReadOnly() {
    return readOnly;
  }


  public void setReadOnly(Boolean readOnly) {
    this.readOnly = readOnly;
  }


  public V1alpha1FCVolumeSource targetWWNs(List<String> targetWWNs) {
    
    this.targetWWNs = targetWWNs;
    return this;
  }

  public V1alpha1FCVolumeSource addTargetWWNsItem(String targetWWNsItem) {
    if (this.targetWWNs == null) {
      this.targetWWNs = new ArrayList<String>();
    }
    this.targetWWNs.add(targetWWNsItem);
    return this;
  }

   /**
   * Get targetWWNs
   * @return targetWWNs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getTargetWWNs() {
    return targetWWNs;
  }


  public void setTargetWWNs(List<String> targetWWNs) {
    this.targetWWNs = targetWWNs;
  }


  public V1alpha1FCVolumeSource wwids(List<String> wwids) {
    
    this.wwids = wwids;
    return this;
  }

  public V1alpha1FCVolumeSource addWwidsItem(String wwidsItem) {
    if (this.wwids == null) {
      this.wwids = new ArrayList<String>();
    }
    this.wwids.add(wwidsItem);
    return this;
  }

   /**
   * Get wwids
   * @return wwids
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getWwids() {
    return wwids;
  }


  public void setWwids(List<String> wwids) {
    this.wwids = wwids;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1FCVolumeSource v1alpha1FCVolumeSource = (V1alpha1FCVolumeSource) o;
    return Objects.equals(this.fsType, v1alpha1FCVolumeSource.fsType) &&
        Objects.equals(this.lun, v1alpha1FCVolumeSource.lun) &&
        Objects.equals(this.readOnly, v1alpha1FCVolumeSource.readOnly) &&
        Objects.equals(this.targetWWNs, v1alpha1FCVolumeSource.targetWWNs) &&
        Objects.equals(this.wwids, v1alpha1FCVolumeSource.wwids);
  }

  @Override
  public int hashCode() {
    return Objects.hash(fsType, lun, readOnly, targetWWNs, wwids);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1FCVolumeSource {\n");
    sb.append("    fsType: ").append(toIndentedString(fsType)).append("\n");
    sb.append("    lun: ").append(toIndentedString(lun)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
    sb.append("    targetWWNs: ").append(toIndentedString(targetWWNs)).append("\n");
    sb.append("    wwids: ").append(toIndentedString(wwids)).append("\n");
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

