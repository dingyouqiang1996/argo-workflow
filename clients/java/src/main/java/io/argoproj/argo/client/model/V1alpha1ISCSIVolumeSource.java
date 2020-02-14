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
import io.argoproj.argo.client.model.V1alpha1LocalObjectReference;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1ISCSIVolumeSource
 */

public class V1alpha1ISCSIVolumeSource {
  public static final String SERIALIZED_NAME_CHAP_AUTH_DISCOVERY = "chapAuthDiscovery";
  @SerializedName(SERIALIZED_NAME_CHAP_AUTH_DISCOVERY)
  private Boolean chapAuthDiscovery;

  public static final String SERIALIZED_NAME_CHAP_AUTH_SESSION = "chapAuthSession";
  @SerializedName(SERIALIZED_NAME_CHAP_AUTH_SESSION)
  private Boolean chapAuthSession;

  public static final String SERIALIZED_NAME_FS_TYPE = "fsType";
  @SerializedName(SERIALIZED_NAME_FS_TYPE)
  private String fsType;

  public static final String SERIALIZED_NAME_INITIATOR_NAME = "initiatorName";
  @SerializedName(SERIALIZED_NAME_INITIATOR_NAME)
  private String initiatorName;

  public static final String SERIALIZED_NAME_IQN = "iqn";
  @SerializedName(SERIALIZED_NAME_IQN)
  private String iqn;

  public static final String SERIALIZED_NAME_ISCSI_INTERFACE = "iscsiInterface";
  @SerializedName(SERIALIZED_NAME_ISCSI_INTERFACE)
  private String iscsiInterface;

  public static final String SERIALIZED_NAME_LUN = "lun";
  @SerializedName(SERIALIZED_NAME_LUN)
  private Integer lun;

  public static final String SERIALIZED_NAME_PORTALS = "portals";
  @SerializedName(SERIALIZED_NAME_PORTALS)
  private List<String> portals = null;

  public static final String SERIALIZED_NAME_READ_ONLY = "readOnly";
  @SerializedName(SERIALIZED_NAME_READ_ONLY)
  private Boolean readOnly;

  public static final String SERIALIZED_NAME_SECRET_REF = "secretRef";
  @SerializedName(SERIALIZED_NAME_SECRET_REF)
  private V1alpha1LocalObjectReference secretRef;

  public static final String SERIALIZED_NAME_TARGET_PORTAL = "targetPortal";
  @SerializedName(SERIALIZED_NAME_TARGET_PORTAL)
  private String targetPortal;


  public V1alpha1ISCSIVolumeSource chapAuthDiscovery(Boolean chapAuthDiscovery) {
    
    this.chapAuthDiscovery = chapAuthDiscovery;
    return this;
  }

   /**
   * Get chapAuthDiscovery
   * @return chapAuthDiscovery
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getChapAuthDiscovery() {
    return chapAuthDiscovery;
  }


  public void setChapAuthDiscovery(Boolean chapAuthDiscovery) {
    this.chapAuthDiscovery = chapAuthDiscovery;
  }


  public V1alpha1ISCSIVolumeSource chapAuthSession(Boolean chapAuthSession) {
    
    this.chapAuthSession = chapAuthSession;
    return this;
  }

   /**
   * Get chapAuthSession
   * @return chapAuthSession
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getChapAuthSession() {
    return chapAuthSession;
  }


  public void setChapAuthSession(Boolean chapAuthSession) {
    this.chapAuthSession = chapAuthSession;
  }


  public V1alpha1ISCSIVolumeSource fsType(String fsType) {
    
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


  public V1alpha1ISCSIVolumeSource initiatorName(String initiatorName) {
    
    this.initiatorName = initiatorName;
    return this;
  }

   /**
   * Get initiatorName
   * @return initiatorName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInitiatorName() {
    return initiatorName;
  }


  public void setInitiatorName(String initiatorName) {
    this.initiatorName = initiatorName;
  }


  public V1alpha1ISCSIVolumeSource iqn(String iqn) {
    
    this.iqn = iqn;
    return this;
  }

   /**
   * Target iSCSI Qualified Name.
   * @return iqn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Target iSCSI Qualified Name.")

  public String getIqn() {
    return iqn;
  }


  public void setIqn(String iqn) {
    this.iqn = iqn;
  }


  public V1alpha1ISCSIVolumeSource iscsiInterface(String iscsiInterface) {
    
    this.iscsiInterface = iscsiInterface;
    return this;
  }

   /**
   * Get iscsiInterface
   * @return iscsiInterface
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIscsiInterface() {
    return iscsiInterface;
  }


  public void setIscsiInterface(String iscsiInterface) {
    this.iscsiInterface = iscsiInterface;
  }


  public V1alpha1ISCSIVolumeSource lun(Integer lun) {
    
    this.lun = lun;
    return this;
  }

   /**
   * iSCSI Target Lun number.
   * @return lun
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "iSCSI Target Lun number.")

  public Integer getLun() {
    return lun;
  }


  public void setLun(Integer lun) {
    this.lun = lun;
  }


  public V1alpha1ISCSIVolumeSource portals(List<String> portals) {
    
    this.portals = portals;
    return this;
  }

  public V1alpha1ISCSIVolumeSource addPortalsItem(String portalsItem) {
    if (this.portals == null) {
      this.portals = new ArrayList<String>();
    }
    this.portals.add(portalsItem);
    return this;
  }

   /**
   * Get portals
   * @return portals
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getPortals() {
    return portals;
  }


  public void setPortals(List<String> portals) {
    this.portals = portals;
  }


  public V1alpha1ISCSIVolumeSource readOnly(Boolean readOnly) {
    
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


  public V1alpha1ISCSIVolumeSource secretRef(V1alpha1LocalObjectReference secretRef) {
    
    this.secretRef = secretRef;
    return this;
  }

   /**
   * Get secretRef
   * @return secretRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1LocalObjectReference getSecretRef() {
    return secretRef;
  }


  public void setSecretRef(V1alpha1LocalObjectReference secretRef) {
    this.secretRef = secretRef;
  }


  public V1alpha1ISCSIVolumeSource targetPortal(String targetPortal) {
    
    this.targetPortal = targetPortal;
    return this;
  }

   /**
   * iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
   * @return targetPortal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).")

  public String getTargetPortal() {
    return targetPortal;
  }


  public void setTargetPortal(String targetPortal) {
    this.targetPortal = targetPortal;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1ISCSIVolumeSource v1alpha1ISCSIVolumeSource = (V1alpha1ISCSIVolumeSource) o;
    return Objects.equals(this.chapAuthDiscovery, v1alpha1ISCSIVolumeSource.chapAuthDiscovery) &&
        Objects.equals(this.chapAuthSession, v1alpha1ISCSIVolumeSource.chapAuthSession) &&
        Objects.equals(this.fsType, v1alpha1ISCSIVolumeSource.fsType) &&
        Objects.equals(this.initiatorName, v1alpha1ISCSIVolumeSource.initiatorName) &&
        Objects.equals(this.iqn, v1alpha1ISCSIVolumeSource.iqn) &&
        Objects.equals(this.iscsiInterface, v1alpha1ISCSIVolumeSource.iscsiInterface) &&
        Objects.equals(this.lun, v1alpha1ISCSIVolumeSource.lun) &&
        Objects.equals(this.portals, v1alpha1ISCSIVolumeSource.portals) &&
        Objects.equals(this.readOnly, v1alpha1ISCSIVolumeSource.readOnly) &&
        Objects.equals(this.secretRef, v1alpha1ISCSIVolumeSource.secretRef) &&
        Objects.equals(this.targetPortal, v1alpha1ISCSIVolumeSource.targetPortal);
  }

  @Override
  public int hashCode() {
    return Objects.hash(chapAuthDiscovery, chapAuthSession, fsType, initiatorName, iqn, iscsiInterface, lun, portals, readOnly, secretRef, targetPortal);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1ISCSIVolumeSource {\n");
    sb.append("    chapAuthDiscovery: ").append(toIndentedString(chapAuthDiscovery)).append("\n");
    sb.append("    chapAuthSession: ").append(toIndentedString(chapAuthSession)).append("\n");
    sb.append("    fsType: ").append(toIndentedString(fsType)).append("\n");
    sb.append("    initiatorName: ").append(toIndentedString(initiatorName)).append("\n");
    sb.append("    iqn: ").append(toIndentedString(iqn)).append("\n");
    sb.append("    iscsiInterface: ").append(toIndentedString(iscsiInterface)).append("\n");
    sb.append("    lun: ").append(toIndentedString(lun)).append("\n");
    sb.append("    portals: ").append(toIndentedString(portals)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
    sb.append("    secretRef: ").append(toIndentedString(secretRef)).append("\n");
    sb.append("    targetPortal: ").append(toIndentedString(targetPortal)).append("\n");
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

