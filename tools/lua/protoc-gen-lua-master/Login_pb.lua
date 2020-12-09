-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('Login_pb')


local USERREGISTERREQ = protobuf.Descriptor();
local USERREGISTERREQ_ACCOUT_FIELD = protobuf.FieldDescriptor();
local USERREGISTERREQ_PASSWORD_FIELD = protobuf.FieldDescriptor();
local USERREGISTERRES = protobuf.Descriptor();
local USERREGISTERRES_RESULT_FIELD = protobuf.FieldDescriptor();
local USERLOGINREQ = protobuf.Descriptor();
local USERLOGINREQ_ACCOUT_FIELD = protobuf.FieldDescriptor();
local USERLOGINREQ_PASSWORD_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES = protobuf.Descriptor();
local USERLOGINRES_RESULT_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES_NICKNAME_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES_USERID_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES_HEADID_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES_GENDER_FIELD = protobuf.FieldDescriptor();
local USERLOGINRES_COIN_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINREQ = protobuf.Descriptor();
local GUESTLOGINRES = protobuf.Descriptor();
local GUESTLOGINRES_RESULT_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINRES_NICKNAME_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINRES_USERID_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINRES_HEADID_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINRES_GENDER_FIELD = protobuf.FieldDescriptor();
local GUESTLOGINRES_COIN_FIELD = protobuf.FieldDescriptor();

USERREGISTERREQ_ACCOUT_FIELD.name = "Accout"
USERREGISTERREQ_ACCOUT_FIELD.full_name = ".ProtoMsg.UserRegisterReq.Accout"
USERREGISTERREQ_ACCOUT_FIELD.number = 1
USERREGISTERREQ_ACCOUT_FIELD.index = 0
USERREGISTERREQ_ACCOUT_FIELD.label = 1
USERREGISTERREQ_ACCOUT_FIELD.has_default_value = false
USERREGISTERREQ_ACCOUT_FIELD.default_value = ""
USERREGISTERREQ_ACCOUT_FIELD.type = 9
USERREGISTERREQ_ACCOUT_FIELD.cpp_type = 9

USERREGISTERREQ_PASSWORD_FIELD.name = "PassWord"
USERREGISTERREQ_PASSWORD_FIELD.full_name = ".ProtoMsg.UserRegisterReq.PassWord"
USERREGISTERREQ_PASSWORD_FIELD.number = 2
USERREGISTERREQ_PASSWORD_FIELD.index = 1
USERREGISTERREQ_PASSWORD_FIELD.label = 1
USERREGISTERREQ_PASSWORD_FIELD.has_default_value = false
USERREGISTERREQ_PASSWORD_FIELD.default_value = ""
USERREGISTERREQ_PASSWORD_FIELD.type = 9
USERREGISTERREQ_PASSWORD_FIELD.cpp_type = 9

USERREGISTERREQ.name = "UserRegisterReq"
USERREGISTERREQ.full_name = ".ProtoMsg.UserRegisterReq"
USERREGISTERREQ.nested_types = {}
USERREGISTERREQ.enum_types = {}
USERREGISTERREQ.fields = {USERREGISTERREQ_ACCOUT_FIELD, USERREGISTERREQ_PASSWORD_FIELD}
USERREGISTERREQ.is_extendable = false
USERREGISTERREQ.extensions = {}
USERREGISTERRES_RESULT_FIELD.name = "Result"
USERREGISTERRES_RESULT_FIELD.full_name = ".ProtoMsg.UserRegisterRes.Result"
USERREGISTERRES_RESULT_FIELD.number = 1
USERREGISTERRES_RESULT_FIELD.index = 0
USERREGISTERRES_RESULT_FIELD.label = 1
USERREGISTERRES_RESULT_FIELD.has_default_value = false
USERREGISTERRES_RESULT_FIELD.default_value = 0
USERREGISTERRES_RESULT_FIELD.type = 5
USERREGISTERRES_RESULT_FIELD.cpp_type = 1

USERREGISTERRES.name = "UserRegisterRes"
USERREGISTERRES.full_name = ".ProtoMsg.UserRegisterRes"
USERREGISTERRES.nested_types = {}
USERREGISTERRES.enum_types = {}
USERREGISTERRES.fields = {USERREGISTERRES_RESULT_FIELD}
USERREGISTERRES.is_extendable = false
USERREGISTERRES.extensions = {}
USERLOGINREQ_ACCOUT_FIELD.name = "Accout"
USERLOGINREQ_ACCOUT_FIELD.full_name = ".ProtoMsg.UserLoginReq.Accout"
USERLOGINREQ_ACCOUT_FIELD.number = 1
USERLOGINREQ_ACCOUT_FIELD.index = 0
USERLOGINREQ_ACCOUT_FIELD.label = 1
USERLOGINREQ_ACCOUT_FIELD.has_default_value = false
USERLOGINREQ_ACCOUT_FIELD.default_value = ""
USERLOGINREQ_ACCOUT_FIELD.type = 9
USERLOGINREQ_ACCOUT_FIELD.cpp_type = 9

USERLOGINREQ_PASSWORD_FIELD.name = "PassWord"
USERLOGINREQ_PASSWORD_FIELD.full_name = ".ProtoMsg.UserLoginReq.PassWord"
USERLOGINREQ_PASSWORD_FIELD.number = 2
USERLOGINREQ_PASSWORD_FIELD.index = 1
USERLOGINREQ_PASSWORD_FIELD.label = 1
USERLOGINREQ_PASSWORD_FIELD.has_default_value = false
USERLOGINREQ_PASSWORD_FIELD.default_value = ""
USERLOGINREQ_PASSWORD_FIELD.type = 9
USERLOGINREQ_PASSWORD_FIELD.cpp_type = 9

USERLOGINREQ.name = "UserLoginReq"
USERLOGINREQ.full_name = ".ProtoMsg.UserLoginReq"
USERLOGINREQ.nested_types = {}
USERLOGINREQ.enum_types = {}
USERLOGINREQ.fields = {USERLOGINREQ_ACCOUT_FIELD, USERLOGINREQ_PASSWORD_FIELD}
USERLOGINREQ.is_extendable = false
USERLOGINREQ.extensions = {}
USERLOGINRES_RESULT_FIELD.name = "Result"
USERLOGINRES_RESULT_FIELD.full_name = ".ProtoMsg.UserLoginRes.Result"
USERLOGINRES_RESULT_FIELD.number = 1
USERLOGINRES_RESULT_FIELD.index = 0
USERLOGINRES_RESULT_FIELD.label = 1
USERLOGINRES_RESULT_FIELD.has_default_value = false
USERLOGINRES_RESULT_FIELD.default_value = 0
USERLOGINRES_RESULT_FIELD.type = 5
USERLOGINRES_RESULT_FIELD.cpp_type = 1

USERLOGINRES_NICKNAME_FIELD.name = "Nickname"
USERLOGINRES_NICKNAME_FIELD.full_name = ".ProtoMsg.UserLoginRes.Nickname"
USERLOGINRES_NICKNAME_FIELD.number = 2
USERLOGINRES_NICKNAME_FIELD.index = 1
USERLOGINRES_NICKNAME_FIELD.label = 1
USERLOGINRES_NICKNAME_FIELD.has_default_value = false
USERLOGINRES_NICKNAME_FIELD.default_value = ""
USERLOGINRES_NICKNAME_FIELD.type = 9
USERLOGINRES_NICKNAME_FIELD.cpp_type = 9

USERLOGINRES_USERID_FIELD.name = "UserId"
USERLOGINRES_USERID_FIELD.full_name = ".ProtoMsg.UserLoginRes.UserId"
USERLOGINRES_USERID_FIELD.number = 3
USERLOGINRES_USERID_FIELD.index = 2
USERLOGINRES_USERID_FIELD.label = 1
USERLOGINRES_USERID_FIELD.has_default_value = false
USERLOGINRES_USERID_FIELD.default_value = 0
USERLOGINRES_USERID_FIELD.type = 5
USERLOGINRES_USERID_FIELD.cpp_type = 1

USERLOGINRES_HEADID_FIELD.name = "HeadId"
USERLOGINRES_HEADID_FIELD.full_name = ".ProtoMsg.UserLoginRes.HeadId"
USERLOGINRES_HEADID_FIELD.number = 4
USERLOGINRES_HEADID_FIELD.index = 3
USERLOGINRES_HEADID_FIELD.label = 1
USERLOGINRES_HEADID_FIELD.has_default_value = false
USERLOGINRES_HEADID_FIELD.default_value = 0
USERLOGINRES_HEADID_FIELD.type = 5
USERLOGINRES_HEADID_FIELD.cpp_type = 1

USERLOGINRES_GENDER_FIELD.name = "Gender"
USERLOGINRES_GENDER_FIELD.full_name = ".ProtoMsg.UserLoginRes.Gender"
USERLOGINRES_GENDER_FIELD.number = 5
USERLOGINRES_GENDER_FIELD.index = 4
USERLOGINRES_GENDER_FIELD.label = 1
USERLOGINRES_GENDER_FIELD.has_default_value = false
USERLOGINRES_GENDER_FIELD.default_value = 0
USERLOGINRES_GENDER_FIELD.type = 5
USERLOGINRES_GENDER_FIELD.cpp_type = 1

USERLOGINRES_COIN_FIELD.name = "Coin"
USERLOGINRES_COIN_FIELD.full_name = ".ProtoMsg.UserLoginRes.Coin"
USERLOGINRES_COIN_FIELD.number = 6
USERLOGINRES_COIN_FIELD.index = 5
USERLOGINRES_COIN_FIELD.label = 1
USERLOGINRES_COIN_FIELD.has_default_value = false
USERLOGINRES_COIN_FIELD.default_value = 0
USERLOGINRES_COIN_FIELD.type = 5
USERLOGINRES_COIN_FIELD.cpp_type = 1

USERLOGINRES.name = "UserLoginRes"
USERLOGINRES.full_name = ".ProtoMsg.UserLoginRes"
USERLOGINRES.nested_types = {}
USERLOGINRES.enum_types = {}
USERLOGINRES.fields = {USERLOGINRES_RESULT_FIELD, USERLOGINRES_NICKNAME_FIELD, USERLOGINRES_USERID_FIELD, USERLOGINRES_HEADID_FIELD, USERLOGINRES_GENDER_FIELD, USERLOGINRES_COIN_FIELD}
USERLOGINRES.is_extendable = false
USERLOGINRES.extensions = {}
GUESTLOGINREQ.name = "GuestLoginReq"
GUESTLOGINREQ.full_name = ".ProtoMsg.GuestLoginReq"
GUESTLOGINREQ.nested_types = {}
GUESTLOGINREQ.enum_types = {}
GUESTLOGINREQ.fields = {}
GUESTLOGINREQ.is_extendable = false
GUESTLOGINREQ.extensions = {}
GUESTLOGINRES_RESULT_FIELD.name = "Result"
GUESTLOGINRES_RESULT_FIELD.full_name = ".ProtoMsg.GuestLoginRes.Result"
GUESTLOGINRES_RESULT_FIELD.number = 1
GUESTLOGINRES_RESULT_FIELD.index = 0
GUESTLOGINRES_RESULT_FIELD.label = 1
GUESTLOGINRES_RESULT_FIELD.has_default_value = false
GUESTLOGINRES_RESULT_FIELD.default_value = 0
GUESTLOGINRES_RESULT_FIELD.type = 5
GUESTLOGINRES_RESULT_FIELD.cpp_type = 1

GUESTLOGINRES_NICKNAME_FIELD.name = "Nickname"
GUESTLOGINRES_NICKNAME_FIELD.full_name = ".ProtoMsg.GuestLoginRes.Nickname"
GUESTLOGINRES_NICKNAME_FIELD.number = 2
GUESTLOGINRES_NICKNAME_FIELD.index = 1
GUESTLOGINRES_NICKNAME_FIELD.label = 1
GUESTLOGINRES_NICKNAME_FIELD.has_default_value = false
GUESTLOGINRES_NICKNAME_FIELD.default_value = ""
GUESTLOGINRES_NICKNAME_FIELD.type = 9
GUESTLOGINRES_NICKNAME_FIELD.cpp_type = 9

GUESTLOGINRES_USERID_FIELD.name = "UserId"
GUESTLOGINRES_USERID_FIELD.full_name = ".ProtoMsg.GuestLoginRes.UserId"
GUESTLOGINRES_USERID_FIELD.number = 3
GUESTLOGINRES_USERID_FIELD.index = 2
GUESTLOGINRES_USERID_FIELD.label = 1
GUESTLOGINRES_USERID_FIELD.has_default_value = false
GUESTLOGINRES_USERID_FIELD.default_value = 0
GUESTLOGINRES_USERID_FIELD.type = 5
GUESTLOGINRES_USERID_FIELD.cpp_type = 1

GUESTLOGINRES_HEADID_FIELD.name = "HeadId"
GUESTLOGINRES_HEADID_FIELD.full_name = ".ProtoMsg.GuestLoginRes.HeadId"
GUESTLOGINRES_HEADID_FIELD.number = 4
GUESTLOGINRES_HEADID_FIELD.index = 3
GUESTLOGINRES_HEADID_FIELD.label = 1
GUESTLOGINRES_HEADID_FIELD.has_default_value = false
GUESTLOGINRES_HEADID_FIELD.default_value = 0
GUESTLOGINRES_HEADID_FIELD.type = 5
GUESTLOGINRES_HEADID_FIELD.cpp_type = 1

GUESTLOGINRES_GENDER_FIELD.name = "Gender"
GUESTLOGINRES_GENDER_FIELD.full_name = ".ProtoMsg.GuestLoginRes.Gender"
GUESTLOGINRES_GENDER_FIELD.number = 5
GUESTLOGINRES_GENDER_FIELD.index = 4
GUESTLOGINRES_GENDER_FIELD.label = 1
GUESTLOGINRES_GENDER_FIELD.has_default_value = false
GUESTLOGINRES_GENDER_FIELD.default_value = 0
GUESTLOGINRES_GENDER_FIELD.type = 5
GUESTLOGINRES_GENDER_FIELD.cpp_type = 1

GUESTLOGINRES_COIN_FIELD.name = "Coin"
GUESTLOGINRES_COIN_FIELD.full_name = ".ProtoMsg.GuestLoginRes.Coin"
GUESTLOGINRES_COIN_FIELD.number = 6
GUESTLOGINRES_COIN_FIELD.index = 5
GUESTLOGINRES_COIN_FIELD.label = 1
GUESTLOGINRES_COIN_FIELD.has_default_value = false
GUESTLOGINRES_COIN_FIELD.default_value = 0
GUESTLOGINRES_COIN_FIELD.type = 5
GUESTLOGINRES_COIN_FIELD.cpp_type = 1

GUESTLOGINRES.name = "GuestLoginRes"
GUESTLOGINRES.full_name = ".ProtoMsg.GuestLoginRes"
GUESTLOGINRES.nested_types = {}
GUESTLOGINRES.enum_types = {}
GUESTLOGINRES.fields = {GUESTLOGINRES_RESULT_FIELD, GUESTLOGINRES_NICKNAME_FIELD, GUESTLOGINRES_USERID_FIELD, GUESTLOGINRES_HEADID_FIELD, GUESTLOGINRES_GENDER_FIELD, GUESTLOGINRES_COIN_FIELD}
GUESTLOGINRES.is_extendable = false
GUESTLOGINRES.extensions = {}

GuestLoginReq = protobuf.Message(GUESTLOGINREQ)
GuestLoginRes = protobuf.Message(GUESTLOGINRES)
UserLoginReq = protobuf.Message(USERLOGINREQ)
UserLoginRes = protobuf.Message(USERLOGINRES)
UserRegisterReq = protobuf.Message(USERREGISTERREQ)
UserRegisterRes = protobuf.Message(USERREGISTERRES)

