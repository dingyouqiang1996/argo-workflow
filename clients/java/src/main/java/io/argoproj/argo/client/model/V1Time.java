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
 * Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.  +protobuf.options.marshal&#x3D;false +protobuf.as&#x3D;Timestamp +protobuf.options.(gogoproto.goproto_stringer)&#x3D;false
 */
@ApiModel(description = "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.  +protobuf.options.marshal=false +protobuf.as=Timestamp +protobuf.options.(gogoproto.goproto_stringer)=false")

public class V1Time {
  public static final String SERIALIZED_NAME_NANOS = "nanos";
  @SerializedName(SERIALIZED_NAME_NANOS)
  private Integer nanos;

  public static final String SERIALIZED_NAME_SECONDS = "seconds";
  @SerializedName(SERIALIZED_NAME_SECONDS)
  private String seconds;


  public V1Time nanos(Integer nanos) {
    
    this.nanos = nanos;
    return this;
  }

   /**
   * Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.
   * @return nanos
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.")

  public Integer getNanos() {
    return nanos;
  }


  public void setNanos(Integer nanos) {
    this.nanos = nanos;
  }


  public V1Time seconds(String seconds) {
    
    this.seconds = seconds;
    return this;
  }

   /**
   * Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.
   * @return seconds
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.")

  public String getSeconds() {
    return seconds;
  }


  public void setSeconds(String seconds) {
    this.seconds = seconds;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Time v1Time = (V1Time) o;
    return Objects.equals(this.nanos, v1Time.nanos) &&
        Objects.equals(this.seconds, v1Time.seconds);
  }

  @Override
  public int hashCode() {
    return Objects.hash(nanos, seconds);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Time {\n");
    sb.append("    nanos: ").append(toIndentedString(nanos)).append("\n");
    sb.append("    seconds: ").append(toIndentedString(seconds)).append("\n");
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

