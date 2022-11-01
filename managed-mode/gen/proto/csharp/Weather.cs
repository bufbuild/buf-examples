// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: acme/weather/v1/weather.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Acme.Weather.V1 {

  /// <summary>Holder for reflection information generated from acme/weather/v1/weather.proto</summary>
  public static partial class WeatherReflection {

    #region Descriptor
    /// <summary>File descriptor for acme/weather/v1/weather.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static WeatherReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ch1hY21lL3dlYXRoZXIvdjEvd2VhdGhlci5wcm90bxIPYWNtZS53ZWF0aGVy",
            "LnYxIkQKCExvY2F0aW9uEhoKCGxhdGl0dWRlGAEgASgCUghsYXRpdHVkZRIc",
            "Cglsb25naXR1ZGUYAiABKAJSCWxvbmdpdHVkZSK7AQoHV2VhdGhlchI1Cghs",
            "b2NhdGlvbhgBIAEoCzIZLmFjbWUud2VhdGhlci52MS5Mb2NhdGlvblIIbG9j",
            "YXRpb24SIAoLdGVtcGVyYXR1cmUYAiABKAJSC3RlbXBlcmF0dXJlEh0KCndp",
            "bmRfc3BlZWQYAyABKAJSCXdpbmRTcGVlZBI4Cgljb25kaXRpb24YBCABKA4y",
            "Gi5hY21lLndlYXRoZXIudjEuQ29uZGl0aW9uUgljb25kaXRpb24qYgoJQ29u",
            "ZGl0aW9uEhUKEUNPTkRJVElPTl9VTktOT1dOEAASEwoPQ09ORElUSU9OX1JB",
            "SU5ZEAESEwoPQ09ORElUSU9OX1NVTk5ZEAISFAoQQ09ORElUSU9OX0NMT1VE",
            "WRADQrYBChJpby5hY21lLndlYXRoZXIudjFCDFdlYXRoZXJQcm90b0gCWjRn",
            "aXRodWIuY29tL2FjbWUvd2VhdGhlci92MS9hY21lL3dlYXRoZXIvdjE7d2Vh",
            "dGhlcnYxogIDQVdYqgIPQWNtZS5XZWF0aGVyLlYxygIPQWNtZVxXZWF0aGVy",
            "XFYx4gIbQWNtZVxXZWF0aGVyXFYxXEdQQk1ldGFkYXRh6gIRQWNtZTo6V2Vh",
            "dGhlcjo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Acme.Weather.V1.Condition), }, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Acme.Weather.V1.Location), global::Acme.Weather.V1.Location.Parser, new[]{ "Latitude", "Longitude" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Acme.Weather.V1.Weather), global::Acme.Weather.V1.Weather.Parser, new[]{ "Location", "Temperature", "WindSpeed", "Condition" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum Condition {
    [pbr::OriginalName("CONDITION_UNKNOWN")] Unknown = 0,
    [pbr::OriginalName("CONDITION_RAINY")] Rainy = 1,
    [pbr::OriginalName("CONDITION_SUNNY")] Sunny = 2,
    [pbr::OriginalName("CONDITION_CLOUDY")] Cloudy = 3,
  }

  #endregion

  #region Messages
  public sealed partial class Location : pb::IMessage<Location>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Location> _parser = new pb::MessageParser<Location>(() => new Location());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Location> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Acme.Weather.V1.WeatherReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Location() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Location(Location other) : this() {
      latitude_ = other.latitude_;
      longitude_ = other.longitude_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Location Clone() {
      return new Location(this);
    }

    /// <summary>Field number for the "latitude" field.</summary>
    public const int LatitudeFieldNumber = 1;
    private float latitude_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public float Latitude {
      get { return latitude_; }
      set {
        latitude_ = value;
      }
    }

    /// <summary>Field number for the "longitude" field.</summary>
    public const int LongitudeFieldNumber = 2;
    private float longitude_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public float Longitude {
      get { return longitude_; }
      set {
        longitude_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Location);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Location other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Latitude, other.Latitude)) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Longitude, other.Longitude)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (Latitude != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Latitude);
      if (Longitude != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Longitude);
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Latitude != 0F) {
        output.WriteRawTag(13);
        output.WriteFloat(Latitude);
      }
      if (Longitude != 0F) {
        output.WriteRawTag(21);
        output.WriteFloat(Longitude);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Latitude != 0F) {
        output.WriteRawTag(13);
        output.WriteFloat(Latitude);
      }
      if (Longitude != 0F) {
        output.WriteRawTag(21);
        output.WriteFloat(Longitude);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (Latitude != 0F) {
        size += 1 + 4;
      }
      if (Longitude != 0F) {
        size += 1 + 4;
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Location other) {
      if (other == null) {
        return;
      }
      if (other.Latitude != 0F) {
        Latitude = other.Latitude;
      }
      if (other.Longitude != 0F) {
        Longitude = other.Longitude;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 13: {
            Latitude = input.ReadFloat();
            break;
          }
          case 21: {
            Longitude = input.ReadFloat();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 13: {
            Latitude = input.ReadFloat();
            break;
          }
          case 21: {
            Longitude = input.ReadFloat();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class Weather : pb::IMessage<Weather>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Weather> _parser = new pb::MessageParser<Weather>(() => new Weather());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Weather> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Acme.Weather.V1.WeatherReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Weather() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Weather(Weather other) : this() {
      location_ = other.location_ != null ? other.location_.Clone() : null;
      temperature_ = other.temperature_;
      windSpeed_ = other.windSpeed_;
      condition_ = other.condition_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Weather Clone() {
      return new Weather(this);
    }

    /// <summary>Field number for the "location" field.</summary>
    public const int LocationFieldNumber = 1;
    private global::Acme.Weather.V1.Location location_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Acme.Weather.V1.Location Location {
      get { return location_; }
      set {
        location_ = value;
      }
    }

    /// <summary>Field number for the "temperature" field.</summary>
    public const int TemperatureFieldNumber = 2;
    private float temperature_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public float Temperature {
      get { return temperature_; }
      set {
        temperature_ = value;
      }
    }

    /// <summary>Field number for the "wind_speed" field.</summary>
    public const int WindSpeedFieldNumber = 3;
    private float windSpeed_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public float WindSpeed {
      get { return windSpeed_; }
      set {
        windSpeed_ = value;
      }
    }

    /// <summary>Field number for the "condition" field.</summary>
    public const int ConditionFieldNumber = 4;
    private global::Acme.Weather.V1.Condition condition_ = global::Acme.Weather.V1.Condition.Unknown;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Acme.Weather.V1.Condition Condition {
      get { return condition_; }
      set {
        condition_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Weather);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Weather other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Location, other.Location)) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Temperature, other.Temperature)) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(WindSpeed, other.WindSpeed)) return false;
      if (Condition != other.Condition) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (location_ != null) hash ^= Location.GetHashCode();
      if (Temperature != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Temperature);
      if (WindSpeed != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(WindSpeed);
      if (Condition != global::Acme.Weather.V1.Condition.Unknown) hash ^= Condition.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (location_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Location);
      }
      if (Temperature != 0F) {
        output.WriteRawTag(21);
        output.WriteFloat(Temperature);
      }
      if (WindSpeed != 0F) {
        output.WriteRawTag(29);
        output.WriteFloat(WindSpeed);
      }
      if (Condition != global::Acme.Weather.V1.Condition.Unknown) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Condition);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (location_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Location);
      }
      if (Temperature != 0F) {
        output.WriteRawTag(21);
        output.WriteFloat(Temperature);
      }
      if (WindSpeed != 0F) {
        output.WriteRawTag(29);
        output.WriteFloat(WindSpeed);
      }
      if (Condition != global::Acme.Weather.V1.Condition.Unknown) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Condition);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (location_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Location);
      }
      if (Temperature != 0F) {
        size += 1 + 4;
      }
      if (WindSpeed != 0F) {
        size += 1 + 4;
      }
      if (Condition != global::Acme.Weather.V1.Condition.Unknown) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Condition);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Weather other) {
      if (other == null) {
        return;
      }
      if (other.location_ != null) {
        if (location_ == null) {
          Location = new global::Acme.Weather.V1.Location();
        }
        Location.MergeFrom(other.Location);
      }
      if (other.Temperature != 0F) {
        Temperature = other.Temperature;
      }
      if (other.WindSpeed != 0F) {
        WindSpeed = other.WindSpeed;
      }
      if (other.Condition != global::Acme.Weather.V1.Condition.Unknown) {
        Condition = other.Condition;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (location_ == null) {
              Location = new global::Acme.Weather.V1.Location();
            }
            input.ReadMessage(Location);
            break;
          }
          case 21: {
            Temperature = input.ReadFloat();
            break;
          }
          case 29: {
            WindSpeed = input.ReadFloat();
            break;
          }
          case 32: {
            Condition = (global::Acme.Weather.V1.Condition) input.ReadEnum();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (location_ == null) {
              Location = new global::Acme.Weather.V1.Location();
            }
            input.ReadMessage(Location);
            break;
          }
          case 21: {
            Temperature = input.ReadFloat();
            break;
          }
          case 29: {
            WindSpeed = input.ReadFloat();
            break;
          }
          case 32: {
            Condition = (global::Acme.Weather.V1.Condition) input.ReadEnum();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
