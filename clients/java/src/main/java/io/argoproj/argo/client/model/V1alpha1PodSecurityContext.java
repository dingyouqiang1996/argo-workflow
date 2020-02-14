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
import io.argoproj.argo.client.model.V1alpha1SELinuxOptions;
import io.argoproj.argo.client.model.V1alpha1Sysctl;
import io.argoproj.argo.client.model.V1alpha1WindowsSecurityContextOptions;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1alpha1PodSecurityContext
 */

public class V1alpha1PodSecurityContext {
  public static final String SERIALIZED_NAME_FS_GROUP = "fsGroup";
  @SerializedName(SERIALIZED_NAME_FS_GROUP)
  private String fsGroup;

  public static final String SERIALIZED_NAME_RUN_AS_GROUP = "runAsGroup";
  @SerializedName(SERIALIZED_NAME_RUN_AS_GROUP)
  private String runAsGroup;

  public static final String SERIALIZED_NAME_RUN_AS_NON_ROOT = "runAsNonRoot";
  @SerializedName(SERIALIZED_NAME_RUN_AS_NON_ROOT)
  private Boolean runAsNonRoot;

  public static final String SERIALIZED_NAME_RUN_AS_USER = "runAsUser";
  @SerializedName(SERIALIZED_NAME_RUN_AS_USER)
  private String runAsUser;

  public static final String SERIALIZED_NAME_SE_LINUX_OPTIONS = "seLinuxOptions";
  @SerializedName(SERIALIZED_NAME_SE_LINUX_OPTIONS)
  private V1alpha1SELinuxOptions seLinuxOptions;

  public static final String SERIALIZED_NAME_SUPPLEMENTAL_GROUPS = "supplementalGroups";
  @SerializedName(SERIALIZED_NAME_SUPPLEMENTAL_GROUPS)
  private List<String> supplementalGroups = null;

  public static final String SERIALIZED_NAME_SYSCTLS = "sysctls";
  @SerializedName(SERIALIZED_NAME_SYSCTLS)
  private List<V1alpha1Sysctl> sysctls = null;

  public static final String SERIALIZED_NAME_WINDOWS_OPTIONS = "windowsOptions";
  @SerializedName(SERIALIZED_NAME_WINDOWS_OPTIONS)
  private V1alpha1WindowsSecurityContextOptions windowsOptions;


  public V1alpha1PodSecurityContext fsGroup(String fsGroup) {
    
    this.fsGroup = fsGroup;
    return this;
  }

   /**
   * 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR&#39;d with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. +optional
   * @return fsGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. +optional")

  public String getFsGroup() {
    return fsGroup;
  }


  public void setFsGroup(String fsGroup) {
    this.fsGroup = fsGroup;
  }


  public V1alpha1PodSecurityContext runAsGroup(String runAsGroup) {
    
    this.runAsGroup = runAsGroup;
    return this;
  }

   /**
   * Get runAsGroup
   * @return runAsGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRunAsGroup() {
    return runAsGroup;
  }


  public void setRunAsGroup(String runAsGroup) {
    this.runAsGroup = runAsGroup;
  }


  public V1alpha1PodSecurityContext runAsNonRoot(Boolean runAsNonRoot) {
    
    this.runAsNonRoot = runAsNonRoot;
    return this;
  }

   /**
   * Get runAsNonRoot
   * @return runAsNonRoot
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRunAsNonRoot() {
    return runAsNonRoot;
  }


  public void setRunAsNonRoot(Boolean runAsNonRoot) {
    this.runAsNonRoot = runAsNonRoot;
  }


  public V1alpha1PodSecurityContext runAsUser(String runAsUser) {
    
    this.runAsUser = runAsUser;
    return this;
  }

   /**
   * Get runAsUser
   * @return runAsUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRunAsUser() {
    return runAsUser;
  }


  public void setRunAsUser(String runAsUser) {
    this.runAsUser = runAsUser;
  }


  public V1alpha1PodSecurityContext seLinuxOptions(V1alpha1SELinuxOptions seLinuxOptions) {
    
    this.seLinuxOptions = seLinuxOptions;
    return this;
  }

   /**
   * Get seLinuxOptions
   * @return seLinuxOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1SELinuxOptions getSeLinuxOptions() {
    return seLinuxOptions;
  }


  public void setSeLinuxOptions(V1alpha1SELinuxOptions seLinuxOptions) {
    this.seLinuxOptions = seLinuxOptions;
  }


  public V1alpha1PodSecurityContext supplementalGroups(List<String> supplementalGroups) {
    
    this.supplementalGroups = supplementalGroups;
    return this;
  }

  public V1alpha1PodSecurityContext addSupplementalGroupsItem(String supplementalGroupsItem) {
    if (this.supplementalGroups == null) {
      this.supplementalGroups = new ArrayList<String>();
    }
    this.supplementalGroups.add(supplementalGroupsItem);
    return this;
  }

   /**
   * Get supplementalGroups
   * @return supplementalGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getSupplementalGroups() {
    return supplementalGroups;
  }


  public void setSupplementalGroups(List<String> supplementalGroups) {
    this.supplementalGroups = supplementalGroups;
  }


  public V1alpha1PodSecurityContext sysctls(List<V1alpha1Sysctl> sysctls) {
    
    this.sysctls = sysctls;
    return this;
  }

  public V1alpha1PodSecurityContext addSysctlsItem(V1alpha1Sysctl sysctlsItem) {
    if (this.sysctls == null) {
      this.sysctls = new ArrayList<V1alpha1Sysctl>();
    }
    this.sysctls.add(sysctlsItem);
    return this;
  }

   /**
   * Get sysctls
   * @return sysctls
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1alpha1Sysctl> getSysctls() {
    return sysctls;
  }


  public void setSysctls(List<V1alpha1Sysctl> sysctls) {
    this.sysctls = sysctls;
  }


  public V1alpha1PodSecurityContext windowsOptions(V1alpha1WindowsSecurityContextOptions windowsOptions) {
    
    this.windowsOptions = windowsOptions;
    return this;
  }

   /**
   * Get windowsOptions
   * @return windowsOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1alpha1WindowsSecurityContextOptions getWindowsOptions() {
    return windowsOptions;
  }


  public void setWindowsOptions(V1alpha1WindowsSecurityContextOptions windowsOptions) {
    this.windowsOptions = windowsOptions;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1alpha1PodSecurityContext v1alpha1PodSecurityContext = (V1alpha1PodSecurityContext) o;
    return Objects.equals(this.fsGroup, v1alpha1PodSecurityContext.fsGroup) &&
        Objects.equals(this.runAsGroup, v1alpha1PodSecurityContext.runAsGroup) &&
        Objects.equals(this.runAsNonRoot, v1alpha1PodSecurityContext.runAsNonRoot) &&
        Objects.equals(this.runAsUser, v1alpha1PodSecurityContext.runAsUser) &&
        Objects.equals(this.seLinuxOptions, v1alpha1PodSecurityContext.seLinuxOptions) &&
        Objects.equals(this.supplementalGroups, v1alpha1PodSecurityContext.supplementalGroups) &&
        Objects.equals(this.sysctls, v1alpha1PodSecurityContext.sysctls) &&
        Objects.equals(this.windowsOptions, v1alpha1PodSecurityContext.windowsOptions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(fsGroup, runAsGroup, runAsNonRoot, runAsUser, seLinuxOptions, supplementalGroups, sysctls, windowsOptions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1alpha1PodSecurityContext {\n");
    sb.append("    fsGroup: ").append(toIndentedString(fsGroup)).append("\n");
    sb.append("    runAsGroup: ").append(toIndentedString(runAsGroup)).append("\n");
    sb.append("    runAsNonRoot: ").append(toIndentedString(runAsNonRoot)).append("\n");
    sb.append("    runAsUser: ").append(toIndentedString(runAsUser)).append("\n");
    sb.append("    seLinuxOptions: ").append(toIndentedString(seLinuxOptions)).append("\n");
    sb.append("    supplementalGroups: ").append(toIndentedString(supplementalGroups)).append("\n");
    sb.append("    sysctls: ").append(toIndentedString(sysctls)).append("\n");
    sb.append("    windowsOptions: ").append(toIndentedString(windowsOptions)).append("\n");
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

