# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: memezis.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='memezis.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\rmemezis.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x96\x01\n\x0e\x41\x64\x64PostRequest\x12\x15\n\x05media\x18\x01 \x03(\x0b\x32\x06.Media\x12\x0f\n\x07\x61\x64\x64\x65\x64\x42y\x18\x02 \x01(\t\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x0c\n\x04tags\x18\x04 \x03(\t\x12-\n\tcreatedAt\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tsourceUrl\x18\x06 \x01(\t\"B\n\x05Media\x12\x0b\n\x03URL\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x10\n\x08sourceID\x18\x03 \x01(\t\x12\x0c\n\x04SHA1\x18\x04 \x01(\t\"0\n\x04Vote\x12\n\n\x02up\x18\x01 \x01(\x03\x12\x0c\n\x04\x64own\x18\x02 \x01(\x03\x12\x0e\n\x06status\x18\x03 \x01(\t\"?\n\nDuplicates\x12\x10\n\x08\x63omplete\x18\x01 \x03(\x03\x12\x0e\n\x06likely\x18\x02 \x03(\x03\x12\x0f\n\x07similar\x18\x03 \x03(\x03\">\n\x0f\x41\x64\x64PostResponse\x12\n\n\x02ID\x18\x01 \x01(\x03\x12\x1f\n\nduplicates\x18\x02 \x01(\x0b\x32\x0b.Duplicates\"$\n\x12GetPostByIDRequest\x12\x0e\n\x06postID\x18\x01 \x01(\x03\"|\n\x04Post\x12\n\n\x02ID\x18\x01 \x01(\x03\x12\x15\n\x05media\x18\x02 \x03(\x0b\x32\x06.Media\x12\x0f\n\x07\x61\x64\x64\x65\x64\x42y\x18\x03 \x01(\t\x12\x0e\n\x06source\x18\x04 \x01(\t\x12\x14\n\x05votes\x18\x05 \x01(\x0b\x32\x05.Vote\x12\x0c\n\x04tags\x18\x06 \x03(\t\x12\x0c\n\x04text\x18\x07 \x01(\t\"w\n\x12PublishPostRequest\x12\x0e\n\x06postID\x18\x01 \x01(\x03\x12\x0b\n\x03URL\x18\x02 \x01(\t\x12\x13\n\x0bpublishedTo\x18\x03 \x01(\t\x12/\n\x0bpublishedAt\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"-\n\x0bVoteRequest\x12\x0e\n\x06userID\x18\x01 \x01(\t\x12\x0e\n\x06postID\x18\x02 \x01(\x03\"$\n\x13GetQueueInfoRequest\x12\r\n\x05queue\x18\x01 \x01(\t\"\x85\x01\n\x14GetQueueInfoResponse\x12\x0e\n\x06length\x18\x01 \x01(\x03\x12\x30\n\x0clastPostTime\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12+\n\x07\x64ueTime\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"M\n\rMediaMetadata\x12\x10\n\x08\x66ilename\x18\x01 \x01(\t\x12\x18\n\x04type\x18\x02 \x01(\x0e\x32\n.MediaType\x12\x10\n\x08\x66ilesize\x18\x03 \x01(\x03\"J\n\x12UploadMediaRequest\x12\x0f\n\x05image\x18\x01 \x01(\x0cH\x00\x12\x1e\n\x04meta\x18\x02 \x01(\x0b\x32\x0e.MediaMetadataH\x00\x42\x03\n\x01t\"\"\n\x13UploadMediaResponse\x12\x0b\n\x03URL\x18\x01 \x01(\t*&\n\tMediaType\x12\x07\n\x03PNG\x10\x00\x12\x07\n\x03JPG\x10\x01\x12\x07\n\x03GIF\x10\x02\x32\xcd\x04\n\x07Memezis\x12>\n\x07\x41\x64\x64Post\x12\x0f.AddPostRequest\x1a\x10.AddPostResponse\"\x10\x82\xd3\xe4\x93\x02\n\"\x05/post:\x01*\x12]\n\x0bPublishPost\x12\x13.PublishPostRequest\x1a\x16.google.protobuf.Empty\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/post/{postID}/publish:\x01*\x12\x41\n\x0bGetPostByID\x12\x13.GetPostByIDRequest\x1a\x05.Post\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/post/{postID}\x12\x44\n\rGetRandomPost\x12\x16.google.protobuf.Empty\x1a\x05.Post\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\x0c/post/random\x12?\n\x06UpVote\x12\x0c.VoteRequest\x1a\x05.Vote\" \x82\xd3\xe4\x93\x02\x1a\"\x15/post/{postID}/upvote:\x01*\x12\x43\n\x08\x44ownVote\x12\x0c.VoteRequest\x1a\x05.Vote\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/post/{postID}/downvote:\x01*\x12X\n\x0cGetQueueInfo\x12\x14.GetQueueInfoRequest\x1a\x15.GetQueueInfoResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/queue/{queue}/info\x12:\n\x0bUploadMedia\x12\x13.UploadMediaRequest\x1a\x14.UploadMediaResponse(\x01\x62\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_MEDIATYPE = _descriptor.EnumDescriptor(
  name='MediaType',
  full_name='MediaType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PNG', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JPG', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GIF', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1206,
  serialized_end=1244,
)
_sym_db.RegisterEnumDescriptor(_MEDIATYPE)

MediaType = enum_type_wrapper.EnumTypeWrapper(_MEDIATYPE)
PNG = 0
JPG = 1
GIF = 2



_ADDPOSTREQUEST = _descriptor.Descriptor(
  name='AddPostRequest',
  full_name='AddPostRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='media', full_name='AddPostRequest.media', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addedBy', full_name='AddPostRequest.addedBy', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='text', full_name='AddPostRequest.text', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='AddPostRequest.tags', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='createdAt', full_name='AddPostRequest.createdAt', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sourceUrl', full_name='AddPostRequest.sourceUrl', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=110,
  serialized_end=260,
)


_MEDIA = _descriptor.Descriptor(
  name='Media',
  full_name='Media',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='URL', full_name='Media.URL', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='Media.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sourceID', full_name='Media.sourceID', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='SHA1', full_name='Media.SHA1', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=262,
  serialized_end=328,
)


_VOTE = _descriptor.Descriptor(
  name='Vote',
  full_name='Vote',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='up', full_name='Vote.up', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='down', full_name='Vote.down', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='Vote.status', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=330,
  serialized_end=378,
)


_DUPLICATES = _descriptor.Descriptor(
  name='Duplicates',
  full_name='Duplicates',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='complete', full_name='Duplicates.complete', index=0,
      number=1, type=3, cpp_type=2, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='likely', full_name='Duplicates.likely', index=1,
      number=2, type=3, cpp_type=2, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='similar', full_name='Duplicates.similar', index=2,
      number=3, type=3, cpp_type=2, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=380,
  serialized_end=443,
)


_ADDPOSTRESPONSE = _descriptor.Descriptor(
  name='AddPostResponse',
  full_name='AddPostResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='AddPostResponse.ID', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='duplicates', full_name='AddPostResponse.duplicates', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=445,
  serialized_end=507,
)


_GETPOSTBYIDREQUEST = _descriptor.Descriptor(
  name='GetPostByIDRequest',
  full_name='GetPostByIDRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='postID', full_name='GetPostByIDRequest.postID', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=509,
  serialized_end=545,
)


_POST = _descriptor.Descriptor(
  name='Post',
  full_name='Post',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='Post.ID', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='media', full_name='Post.media', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addedBy', full_name='Post.addedBy', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='Post.source', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='votes', full_name='Post.votes', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='Post.tags', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='text', full_name='Post.text', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=547,
  serialized_end=671,
)


_PUBLISHPOSTREQUEST = _descriptor.Descriptor(
  name='PublishPostRequest',
  full_name='PublishPostRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='postID', full_name='PublishPostRequest.postID', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='URL', full_name='PublishPostRequest.URL', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='publishedTo', full_name='PublishPostRequest.publishedTo', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='publishedAt', full_name='PublishPostRequest.publishedAt', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=673,
  serialized_end=792,
)


_VOTEREQUEST = _descriptor.Descriptor(
  name='VoteRequest',
  full_name='VoteRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='userID', full_name='VoteRequest.userID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='postID', full_name='VoteRequest.postID', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=794,
  serialized_end=839,
)


_GETQUEUEINFOREQUEST = _descriptor.Descriptor(
  name='GetQueueInfoRequest',
  full_name='GetQueueInfoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='queue', full_name='GetQueueInfoRequest.queue', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=841,
  serialized_end=877,
)


_GETQUEUEINFORESPONSE = _descriptor.Descriptor(
  name='GetQueueInfoResponse',
  full_name='GetQueueInfoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='length', full_name='GetQueueInfoResponse.length', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lastPostTime', full_name='GetQueueInfoResponse.lastPostTime', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dueTime', full_name='GetQueueInfoResponse.dueTime', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=880,
  serialized_end=1013,
)


_MEDIAMETADATA = _descriptor.Descriptor(
  name='MediaMetadata',
  full_name='MediaMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='filename', full_name='MediaMetadata.filename', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='MediaMetadata.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='filesize', full_name='MediaMetadata.filesize', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1015,
  serialized_end=1092,
)


_UPLOADMEDIAREQUEST = _descriptor.Descriptor(
  name='UploadMediaRequest',
  full_name='UploadMediaRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='UploadMediaRequest.image', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='UploadMediaRequest.meta', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='t', full_name='UploadMediaRequest.t',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=1094,
  serialized_end=1168,
)


_UPLOADMEDIARESPONSE = _descriptor.Descriptor(
  name='UploadMediaResponse',
  full_name='UploadMediaResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='URL', full_name='UploadMediaResponse.URL', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1170,
  serialized_end=1204,
)

_ADDPOSTREQUEST.fields_by_name['media'].message_type = _MEDIA
_ADDPOSTREQUEST.fields_by_name['createdAt'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_ADDPOSTRESPONSE.fields_by_name['duplicates'].message_type = _DUPLICATES
_POST.fields_by_name['media'].message_type = _MEDIA
_POST.fields_by_name['votes'].message_type = _VOTE
_PUBLISHPOSTREQUEST.fields_by_name['publishedAt'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GETQUEUEINFORESPONSE.fields_by_name['lastPostTime'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GETQUEUEINFORESPONSE.fields_by_name['dueTime'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_MEDIAMETADATA.fields_by_name['type'].enum_type = _MEDIATYPE
_UPLOADMEDIAREQUEST.fields_by_name['meta'].message_type = _MEDIAMETADATA
_UPLOADMEDIAREQUEST.oneofs_by_name['t'].fields.append(
  _UPLOADMEDIAREQUEST.fields_by_name['image'])
_UPLOADMEDIAREQUEST.fields_by_name['image'].containing_oneof = _UPLOADMEDIAREQUEST.oneofs_by_name['t']
_UPLOADMEDIAREQUEST.oneofs_by_name['t'].fields.append(
  _UPLOADMEDIAREQUEST.fields_by_name['meta'])
_UPLOADMEDIAREQUEST.fields_by_name['meta'].containing_oneof = _UPLOADMEDIAREQUEST.oneofs_by_name['t']
DESCRIPTOR.message_types_by_name['AddPostRequest'] = _ADDPOSTREQUEST
DESCRIPTOR.message_types_by_name['Media'] = _MEDIA
DESCRIPTOR.message_types_by_name['Vote'] = _VOTE
DESCRIPTOR.message_types_by_name['Duplicates'] = _DUPLICATES
DESCRIPTOR.message_types_by_name['AddPostResponse'] = _ADDPOSTRESPONSE
DESCRIPTOR.message_types_by_name['GetPostByIDRequest'] = _GETPOSTBYIDREQUEST
DESCRIPTOR.message_types_by_name['Post'] = _POST
DESCRIPTOR.message_types_by_name['PublishPostRequest'] = _PUBLISHPOSTREQUEST
DESCRIPTOR.message_types_by_name['VoteRequest'] = _VOTEREQUEST
DESCRIPTOR.message_types_by_name['GetQueueInfoRequest'] = _GETQUEUEINFOREQUEST
DESCRIPTOR.message_types_by_name['GetQueueInfoResponse'] = _GETQUEUEINFORESPONSE
DESCRIPTOR.message_types_by_name['MediaMetadata'] = _MEDIAMETADATA
DESCRIPTOR.message_types_by_name['UploadMediaRequest'] = _UPLOADMEDIAREQUEST
DESCRIPTOR.message_types_by_name['UploadMediaResponse'] = _UPLOADMEDIARESPONSE
DESCRIPTOR.enum_types_by_name['MediaType'] = _MEDIATYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AddPostRequest = _reflection.GeneratedProtocolMessageType('AddPostRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDPOSTREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:AddPostRequest)
  })
_sym_db.RegisterMessage(AddPostRequest)

Media = _reflection.GeneratedProtocolMessageType('Media', (_message.Message,), {
  'DESCRIPTOR' : _MEDIA,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:Media)
  })
_sym_db.RegisterMessage(Media)

Vote = _reflection.GeneratedProtocolMessageType('Vote', (_message.Message,), {
  'DESCRIPTOR' : _VOTE,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:Vote)
  })
_sym_db.RegisterMessage(Vote)

Duplicates = _reflection.GeneratedProtocolMessageType('Duplicates', (_message.Message,), {
  'DESCRIPTOR' : _DUPLICATES,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:Duplicates)
  })
_sym_db.RegisterMessage(Duplicates)

AddPostResponse = _reflection.GeneratedProtocolMessageType('AddPostResponse', (_message.Message,), {
  'DESCRIPTOR' : _ADDPOSTRESPONSE,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:AddPostResponse)
  })
_sym_db.RegisterMessage(AddPostResponse)

GetPostByIDRequest = _reflection.GeneratedProtocolMessageType('GetPostByIDRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPOSTBYIDREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:GetPostByIDRequest)
  })
_sym_db.RegisterMessage(GetPostByIDRequest)

Post = _reflection.GeneratedProtocolMessageType('Post', (_message.Message,), {
  'DESCRIPTOR' : _POST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:Post)
  })
_sym_db.RegisterMessage(Post)

PublishPostRequest = _reflection.GeneratedProtocolMessageType('PublishPostRequest', (_message.Message,), {
  'DESCRIPTOR' : _PUBLISHPOSTREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:PublishPostRequest)
  })
_sym_db.RegisterMessage(PublishPostRequest)

VoteRequest = _reflection.GeneratedProtocolMessageType('VoteRequest', (_message.Message,), {
  'DESCRIPTOR' : _VOTEREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:VoteRequest)
  })
_sym_db.RegisterMessage(VoteRequest)

GetQueueInfoRequest = _reflection.GeneratedProtocolMessageType('GetQueueInfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETQUEUEINFOREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:GetQueueInfoRequest)
  })
_sym_db.RegisterMessage(GetQueueInfoRequest)

GetQueueInfoResponse = _reflection.GeneratedProtocolMessageType('GetQueueInfoResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETQUEUEINFORESPONSE,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:GetQueueInfoResponse)
  })
_sym_db.RegisterMessage(GetQueueInfoResponse)

MediaMetadata = _reflection.GeneratedProtocolMessageType('MediaMetadata', (_message.Message,), {
  'DESCRIPTOR' : _MEDIAMETADATA,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:MediaMetadata)
  })
_sym_db.RegisterMessage(MediaMetadata)

UploadMediaRequest = _reflection.GeneratedProtocolMessageType('UploadMediaRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADMEDIAREQUEST,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:UploadMediaRequest)
  })
_sym_db.RegisterMessage(UploadMediaRequest)

UploadMediaResponse = _reflection.GeneratedProtocolMessageType('UploadMediaResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADMEDIARESPONSE,
  '__module__' : 'memezis_pb2'
  # @@protoc_insertion_point(class_scope:UploadMediaResponse)
  })
_sym_db.RegisterMessage(UploadMediaResponse)



_MEMEZIS = _descriptor.ServiceDescriptor(
  name='Memezis',
  full_name='Memezis',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1247,
  serialized_end=1836,
  methods=[
  _descriptor.MethodDescriptor(
    name='AddPost',
    full_name='Memezis.AddPost',
    index=0,
    containing_service=None,
    input_type=_ADDPOSTREQUEST,
    output_type=_ADDPOSTRESPONSE,
    serialized_options=b'\202\323\344\223\002\n\"\005/post:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='PublishPost',
    full_name='Memezis.PublishPost',
    index=1,
    containing_service=None,
    input_type=_PUBLISHPOSTREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\033\"\026/post/{postID}/publish:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='GetPostByID',
    full_name='Memezis.GetPostByID',
    index=2,
    containing_service=None,
    input_type=_GETPOSTBYIDREQUEST,
    output_type=_POST,
    serialized_options=b'\202\323\344\223\002\020\022\016/post/{postID}',
  ),
  _descriptor.MethodDescriptor(
    name='GetRandomPost',
    full_name='Memezis.GetRandomPost',
    index=3,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_POST,
    serialized_options=b'\202\323\344\223\002\016\022\014/post/random',
  ),
  _descriptor.MethodDescriptor(
    name='UpVote',
    full_name='Memezis.UpVote',
    index=4,
    containing_service=None,
    input_type=_VOTEREQUEST,
    output_type=_VOTE,
    serialized_options=b'\202\323\344\223\002\032\"\025/post/{postID}/upvote:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='DownVote',
    full_name='Memezis.DownVote',
    index=5,
    containing_service=None,
    input_type=_VOTEREQUEST,
    output_type=_VOTE,
    serialized_options=b'\202\323\344\223\002\034\"\027/post/{postID}/downvote:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='GetQueueInfo',
    full_name='Memezis.GetQueueInfo',
    index=6,
    containing_service=None,
    input_type=_GETQUEUEINFOREQUEST,
    output_type=_GETQUEUEINFORESPONSE,
    serialized_options=b'\202\323\344\223\002\025\022\023/queue/{queue}/info',
  ),
  _descriptor.MethodDescriptor(
    name='UploadMedia',
    full_name='Memezis.UploadMedia',
    index=7,
    containing_service=None,
    input_type=_UPLOADMEDIAREQUEST,
    output_type=_UPLOADMEDIARESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MEMEZIS)

DESCRIPTOR.services_by_name['Memezis'] = _MEMEZIS

# @@protoc_insertion_point(module_scope)
