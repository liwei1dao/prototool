// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: comm.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Pb {

  /// <summary>Holder for reflection information generated from comm.proto</summary>
  public static partial class CommReflection {

    #region Descriptor
    /// <summary>File descriptor for comm.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CommReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cgpjb21tLnByb3RvEgJwYhoUZ29nb3Byb3RvL2dvZ28ucHJvdG8iZAoIVXNl",
            "ckluZm8SKAoGVXNlcklkGAEgASgNQhjq3h8GVXNlcklk8t4fCmJzb246Il9p",
            "ZCISEAoITmljZU5hbWUYAiABKAkSCwoDU2V4GAMgASgNEg8KB0hlYWRVcmwY",
            "BCABKAlCBPjhHgFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Gogoproto.GogoReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Pb.UserInfo), global::Pb.UserInfo.Parser, new[]{ "UserId", "NiceName", "Sex", "HeadUrl" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  ///公用消息结构代码
  /// </summary>
  public sealed partial class UserInfo : pb::IMessage<UserInfo> {
    private static readonly pb::MessageParser<UserInfo> _parser = new pb::MessageParser<UserInfo>(() => new UserInfo());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<UserInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Pb.CommReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public UserInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public UserInfo(UserInfo other) : this() {
      userId_ = other.userId_;
      niceName_ = other.niceName_;
      sex_ = other.sex_;
      headUrl_ = other.headUrl_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public UserInfo Clone() {
      return new UserInfo(this);
    }

    /// <summary>Field number for the "UserId" field.</summary>
    public const int UserIdFieldNumber = 1;
    private uint userId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public uint UserId {
      get { return userId_; }
      set {
        userId_ = value;
      }
    }

    /// <summary>Field number for the "NiceName" field.</summary>
    public const int NiceNameFieldNumber = 2;
    private string niceName_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string NiceName {
      get { return niceName_; }
      set {
        niceName_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Sex" field.</summary>
    public const int SexFieldNumber = 3;
    private uint sex_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public uint Sex {
      get { return sex_; }
      set {
        sex_ = value;
      }
    }

    /// <summary>Field number for the "HeadUrl" field.</summary>
    public const int HeadUrlFieldNumber = 4;
    private string headUrl_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string HeadUrl {
      get { return headUrl_; }
      set {
        headUrl_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as UserInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(UserInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (UserId != other.UserId) return false;
      if (NiceName != other.NiceName) return false;
      if (Sex != other.Sex) return false;
      if (HeadUrl != other.HeadUrl) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (UserId != 0) hash ^= UserId.GetHashCode();
      if (NiceName.Length != 0) hash ^= NiceName.GetHashCode();
      if (Sex != 0) hash ^= Sex.GetHashCode();
      if (HeadUrl.Length != 0) hash ^= HeadUrl.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (UserId != 0) {
        output.WriteRawTag(8);
        output.WriteUInt32(UserId);
      }
      if (NiceName.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(NiceName);
      }
      if (Sex != 0) {
        output.WriteRawTag(24);
        output.WriteUInt32(Sex);
      }
      if (HeadUrl.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(HeadUrl);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (UserId != 0) {
        size += 1 + pb::CodedOutputStream.ComputeUInt32Size(UserId);
      }
      if (NiceName.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(NiceName);
      }
      if (Sex != 0) {
        size += 1 + pb::CodedOutputStream.ComputeUInt32Size(Sex);
      }
      if (HeadUrl.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(HeadUrl);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(UserInfo other) {
      if (other == null) {
        return;
      }
      if (other.UserId != 0) {
        UserId = other.UserId;
      }
      if (other.NiceName.Length != 0) {
        NiceName = other.NiceName;
      }
      if (other.Sex != 0) {
        Sex = other.Sex;
      }
      if (other.HeadUrl.Length != 0) {
        HeadUrl = other.HeadUrl;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            UserId = input.ReadUInt32();
            break;
          }
          case 18: {
            NiceName = input.ReadString();
            break;
          }
          case 24: {
            Sex = input.ReadUInt32();
            break;
          }
          case 34: {
            HeadUrl = input.ReadString();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code