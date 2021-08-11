// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package engine

import (
	"errors"
	"strings"
)

// This file repeats the same content many times in order to make access
// efficient for seclang and transactions

const VARIABLES_COUNT = 88
const (
	// Single valua variables
	VARIABLE_URLENCODED_ERROR                 = 0x00
	VARIABLE_RESPONSE_CONTENT_TYPE            = 0x01
	VARIABLE_UNIQUE_ID                        = 0x02
	VARIABLE_ARGS_COMBINED_SIZE               = 0x03
	VARIABLE_AUTH_TYPE                        = 0x04
	VARIABLE_FILES_COMBINED_SIZE              = 0x05
	VARIABLE_FULL_REQUEST                     = 0x06
	VARIABLE_FULL_REQUEST_LENGTH              = 0x07
	VARIABLE_INBOUND_DATA_ERROR               = 0x08
	VARIABLE_MATCHED_VAR                      = 0x09
	VARIABLE_MATCHED_VAR_NAME                 = 0x0A
	VARIABLE_MULTIPART_BOUNDARY_QUOTED        = 0x0B
	VARIABLE_MULTIPART_BOUNDARY_WHITESPACE    = 0x0C
	VARIABLE_MULTIPART_CRLF_LF_LINES          = 0x0D
	VARIABLE_MULTIPART_DATA_AFTER             = 0x0E
	VARIABLE_MULTIPART_DATA_BEFORE            = 0x0F
	VARIABLE_MULTIPART_FILE_LIMIT_EXCEEDED    = 0x10
	VARIABLE_MULTIPART_HEADER_FOLDING         = 0x11
	VARIABLE_MULTIPART_INVALID_HEADER_FOLDING = 0x12
	VARIABLE_MULTIPART_INVALID_PART           = 0x13
	VARIABLE_MULTIPART_INVALID_QUOTING        = 0x14
	VARIABLE_MULTIPART_LF_LINE                = 0x15
	VARIABLE_MULTIPART_MISSING_SEMICOLON      = 0x16
	VARIABLE_MULTIPART_STRICT_ERROR           = 0x17
	VARIABLE_MULTIPART_UNMATCHED_BOUNDARY     = 0x18
	VARIABLE_OUTBOUND_DATA_ERROR              = 0x19
	VARIABLE_PATH_INFO                        = 0x1A
	VARIABLE_QUERY_STRING                     = 0x1B
	VARIABLE_REMOTE_ADDR                      = 0x1C
	VARIABLE_REMOTE_HOST                      = 0x1D
	VARIABLE_REMOTE_PORT                      = 0x1E
	VARIABLE_REQBODY_ERROR                    = 0x1F
	VARIABLE_REQBODY_ERROR_MSG                = 0x20
	VARIABLE_REQBODY_PROCESSOR_ERROR          = 0x21
	VARIABLE_REQBODY_PROCESSOR_ERROR_MSG      = 0x22
	VARIABLE_REQBODY_PROCESSOR                = 0x23
	VARIABLE_REQUEST_BASENAME                 = 0x24
	VARIABLE_REQUEST_BODY                     = 0x25
	VARIABLE_REQUEST_BODY_LENGTH              = 0x26
	VARIABLE_REQUEST_FILENAME                 = 0x27
	VARIABLE_REQUEST_LINE                     = 0x28
	VARIABLE_REQUEST_METHOD                   = 0x29
	VARIABLE_REQUEST_PROTOCOL                 = 0x2A
	VARIABLE_REQUEST_URI                      = 0x2B
	VARIABLE_REQUEST_URI_RAW                  = 0x2C
	VARIABLE_RESPONSE_BODY                    = 0x2D
	VARIABLE_RESPONSE_CONTENT_LENGTH          = 0x2E
	VARIABLE_RESPONSE_PROTOCOL                = 0x2F
	VARIABLE_RESPONSE_STATUS                  = 0x30
	VARIABLE_SERVER_ADDR                      = 0x31
	VARIABLE_SERVER_NAME                      = 0x32
	VARIABLE_SERVER_PORT                      = 0x33
	VARIABLE_SESSIONID                        = 0x34

	// Set Variables
	VARIABLE_RESPONSE_HEADERS_NAMES = 0x35
	VARIABLE_REQUEST_HEADERS_NAMES  = 0x36
	VARIABLE_USERID                 = 0x37
	VARIABLE_ARGS                   = 0x38
	VARIABLE_ARGS_GET               = 0x39
	VARIABLE_ARGS_POST              = 0x3A
	VARIABLE_FILES_SIZES            = 0x3B
	VARIABLE_FILES_NAMES            = 0x3C
	VARIABLE_FILES_TMP_CONTENT      = 0x3D
	VARIABLE_MULTIPART_FILENAME     = 0x3E
	VARIABLE_MULTIPART_NAME         = 0x3F
	VARIABLE_MATCHED_VARS_NAMES     = 0x40
	VARIABLE_MATCHED_VARS           = 0x41
	VARIABLE_FILES                  = 0x42
	VARIABLE_REQUEST_COOKIES        = 0x43
	VARIABLE_REQUEST_HEADERS        = 0x44
	VARIABLE_RESPONSE_HEADERS       = 0x45
	VARIABLE_GEO                    = 0x46
	VARIABLE_REQUEST_COOKIES_NAMES  = 0x47
	VARIABLE_FILES_TMPNAMES         = 0x48
	VARIABLE_ARGS_NAMES             = 0x49
	VARIABLE_ARGS_GET_NAMES         = 0x4A
	VARIABLE_ARGS_POST_NAMES        = 0x4B
	VARIABLE_TX                     = 0x4C

	VARIABLE_RULE               = 0x52 //TODO FIX
	VARIABLE_XML                = 0x53 //TODO FIX
	VARIABLE_JSON               = 0x54 //TODO FIX
	VARIABLE_INBOUND_ERROR_DATA = 0x55 //TODO FIX
	VARIABLE_DURATION           = 0x56 //TODO FIX

	// Persistent collections
	VARIABLE_GLOBAL   = 0x4D
	VARIABLE_IP       = 0x4E
	VARIABLE_SESSION  = 0x4F
	VARIABLE_USER     = 0x50
	VARIABLE_RESOURCE = 0x51
)

// NameToVariable returns the byte interpretation
// of a variable from a string
// Returns error if there is no representation
func NameToVariable(name string) (byte, error) {
	name = strings.ToUpper(name)
	switch name {
	case "URLENCODED_ERROR":
		return VARIABLE_URLENCODED_ERROR, nil
	case "RESPONSE_CONTENT_TYPE":
		return VARIABLE_RESPONSE_CONTENT_TYPE, nil
	case "UNIQUE_ID":
		return VARIABLE_UNIQUE_ID, nil
	case "ARGS_COMBINED_SIZE":
		return VARIABLE_ARGS_COMBINED_SIZE, nil
	case "AUTH_TYPE":
		return VARIABLE_AUTH_TYPE, nil
	case "FILES_COMBINED_SIZE":
		return VARIABLE_FILES_COMBINED_SIZE, nil
	case "FULL_REQUEST":
		return VARIABLE_FULL_REQUEST, nil
	case "FULL_REQUEST_LENGTH":
		return VARIABLE_FULL_REQUEST_LENGTH, nil
	case "INBOUND_DATA_ERROR":
		return VARIABLE_INBOUND_DATA_ERROR, nil
	case "MATCHED_VAR":
		return VARIABLE_MATCHED_VAR, nil
	case "MATCHED_VAR_NAME":
		return VARIABLE_MATCHED_VAR_NAME, nil
	case "MULTIPART_BOUNDARY_QUOTED":
		return VARIABLE_MULTIPART_BOUNDARY_QUOTED, nil
	case "MULTIPART_BOUNDARY_WHITESPACE":
		return VARIABLE_MULTIPART_BOUNDARY_WHITESPACE, nil
	case "MULTIPART_CRLF_LF_LINES":
		return VARIABLE_MULTIPART_CRLF_LF_LINES, nil
	case "MULTIPART_DATA_AFTER":
		return VARIABLE_MULTIPART_DATA_AFTER, nil
	case "MULTIPART_DATA_BEFORE":
		return VARIABLE_MULTIPART_DATA_BEFORE, nil
	case "MULTIPART_FILE_LIMIT_EXCEEDED":
		return VARIABLE_MULTIPART_FILE_LIMIT_EXCEEDED, nil
	case "MULTIPART_HEADER_FOLDING":
		return VARIABLE_MULTIPART_HEADER_FOLDING, nil
	case "MULTIPART_INVALID_HEADER_FOLDING":
		return VARIABLE_MULTIPART_INVALID_HEADER_FOLDING, nil
	case "MULTIPART_INVALID_PART":
		return VARIABLE_MULTIPART_INVALID_PART, nil
	case "MULTIPART_INVALID_QUOTING":
		return VARIABLE_MULTIPART_INVALID_QUOTING, nil
	case "MULTIPART_LF_LINE":
		return VARIABLE_MULTIPART_LF_LINE, nil
	case "MULTIPART_MISSING_SEMICOLON":
		return VARIABLE_MULTIPART_MISSING_SEMICOLON, nil
	case "MULTIPART_STRICT_ERROR":
		return VARIABLE_MULTIPART_STRICT_ERROR, nil
	case "MULTIPART_UNMATCHED_BOUNDARY":
		return VARIABLE_MULTIPART_UNMATCHED_BOUNDARY, nil
	case "OUTBOUND_DATA_ERROR":
		return VARIABLE_OUTBOUND_DATA_ERROR, nil
	case "PATH_INFO":
		return VARIABLE_PATH_INFO, nil
	case "QUERY_STRING":
		return VARIABLE_QUERY_STRING, nil
	case "REMOTE_ADDR":
		return VARIABLE_REMOTE_ADDR, nil
	case "REMOTE_HOST":
		return VARIABLE_REMOTE_HOST, nil
	case "REMOTE_PORT":
		return VARIABLE_REMOTE_PORT, nil
	case "REQBODY_ERROR":
		return VARIABLE_REQBODY_ERROR, nil
	case "REQBODY_ERROR_MSG":
		return VARIABLE_REQBODY_ERROR_MSG, nil
	case "REQBODY_PROCESSOR_ERROR":
		return VARIABLE_REQBODY_PROCESSOR_ERROR, nil
	case "REQBODY_PROCESSOR_ERROR_MSG":
		return VARIABLE_REQBODY_PROCESSOR_ERROR_MSG, nil
	case "REQBODY_PROCESSOR":
		return VARIABLE_REQBODY_PROCESSOR, nil
	case "REQUEST_BASENAME":
		return VARIABLE_REQUEST_BASENAME, nil
	case "REQUEST_BODY":
		return VARIABLE_REQUEST_BODY, nil
	case "REQUEST_BODY_LENGTH":
		return VARIABLE_REQUEST_BODY_LENGTH, nil
	case "REQUEST_FILENAME":
		return VARIABLE_REQUEST_FILENAME, nil
	case "REQUEST_LINE":
		return VARIABLE_REQUEST_LINE, nil
	case "REQUEST_METHOD":
		return VARIABLE_REQUEST_METHOD, nil
	case "REQUEST_PROTOCOL":
		return VARIABLE_REQUEST_PROTOCOL, nil
	case "REQUEST_URI":
		return VARIABLE_REQUEST_URI, nil
	case "REQUEST_URI_RAW":
		return VARIABLE_REQUEST_URI_RAW, nil
	case "RESOURCE":
		return VARIABLE_RESOURCE, nil
	case "RESPONSE_BODY":
		return VARIABLE_RESPONSE_BODY, nil
	case "RESPONSE_CONTENT_LENGTH":
		return VARIABLE_RESPONSE_CONTENT_LENGTH, nil
	case "RESPONSE_PROTOCOL":
		return VARIABLE_RESPONSE_PROTOCOL, nil
	case "RESPONSE_STATUS":
		return VARIABLE_RESPONSE_STATUS, nil
	case "SERVER_ADDR":
		return VARIABLE_SERVER_ADDR, nil
	case "SERVER_NAME":
		return VARIABLE_SERVER_NAME, nil
	case "SERVER_PORT":
		return VARIABLE_SERVER_PORT, nil
	case "SESSIONID":
		return VARIABLE_SESSIONID, nil
	case "RESPONSE_HEADERS_NAMES":
		return VARIABLE_RESPONSE_HEADERS_NAMES, nil
	case "REQUEST_HEADERS_NAMES":
		return VARIABLE_REQUEST_HEADERS_NAMES, nil
	case "USERID":
		return VARIABLE_USERID, nil
	case "ARGS":
		return VARIABLE_ARGS, nil
	case "ARGS_GET":
		return VARIABLE_ARGS_GET, nil
	case "ARGS_POST":
		return VARIABLE_ARGS_POST, nil
	case "FILES_SIZES":
		return VARIABLE_FILES_SIZES, nil
	case "FILES_NAMES":
		return VARIABLE_FILES_NAMES, nil
	case "FILES_TMP_CONTENT":
		return VARIABLE_FILES_TMP_CONTENT, nil
	case "MULTIPART_FILENAME":
		return VARIABLE_MULTIPART_FILENAME, nil
	case "MULTIPART_NAME":
		return VARIABLE_MULTIPART_NAME, nil
	case "MATCHED_VARS_NAMES":
		return VARIABLE_MATCHED_VARS_NAMES, nil
	case "MATCHED_VARS":
		return VARIABLE_MATCHED_VARS, nil
	case "FILES":
		return VARIABLE_FILES, nil
	case "REQUEST_COOKIES":
		return VARIABLE_REQUEST_COOKIES, nil
	case "REQUEST_HEADERS":
		return VARIABLE_REQUEST_HEADERS, nil
	case "RESPONSE_HEADERS":
		return VARIABLE_RESPONSE_HEADERS, nil
	case "GEO":
		return VARIABLE_GEO, nil
	case "REQUEST_COOKIES_NAMES":
		return VARIABLE_REQUEST_COOKIES_NAMES, nil
	case "FILES_TMPNAMES":
		return VARIABLE_FILES_TMPNAMES, nil
	case "ARGS_NAMES":
		return VARIABLE_ARGS_NAMES, nil
	case "ARGS_GET_NAMES":
		return VARIABLE_ARGS_GET_NAMES, nil
	case "ARGS_POST_NAMES":
		return VARIABLE_ARGS_POST_NAMES, nil
	case "GLOBAL":
		return VARIABLE_GLOBAL, nil
	case "IP":
		return VARIABLE_IP, nil
	case "SESSION":
		return VARIABLE_SESSION, nil
	case "USER":
		return VARIABLE_USER, nil
	case "RULE":
		return VARIABLE_RULE, nil
	case "XML":
		return VARIABLE_XML, nil
	case "TX":
		return VARIABLE_TX, nil
	case "DURATION":
		return VARIABLE_DURATION, nil
	}
	return 0, errors.New("Invalid variable " + name)
}

// VariableToName transforms a VARIABLE representation
// into a string, it's used for audit and logging
func VariableToName(v byte) string {
	switch v {
	case VARIABLE_URLENCODED_ERROR:
		return "URLENCODED_ERROR"
	case VARIABLE_RESPONSE_CONTENT_TYPE:
		return "RESPONSE_CONTENT_TYPE"
	case VARIABLE_UNIQUE_ID:
		return "UNIQUE_ID"
	case VARIABLE_ARGS_COMBINED_SIZE:
		return "ARGS_COMBINED_SIZE"
	case VARIABLE_AUTH_TYPE:
		return "AUTH_TYPE"
	case VARIABLE_FILES_COMBINED_SIZE:
		return "FILES_COMBINED_SIZE"
	case VARIABLE_FULL_REQUEST:
		return "FULL_REQUEST"
	case VARIABLE_FULL_REQUEST_LENGTH:
		return "FULL_REQUEST_LENGTH"
	case VARIABLE_INBOUND_DATA_ERROR:
		return "INBOUND_DATA_ERROR"
	case VARIABLE_MATCHED_VAR:
		return "MATCHED_VAR"
	case VARIABLE_MATCHED_VAR_NAME:
		return "MATCHED_VAR_NAME"
	case VARIABLE_MULTIPART_BOUNDARY_QUOTED:
		return "MULTIPART_BOUNDARY_QUOTED"
	case VARIABLE_MULTIPART_BOUNDARY_WHITESPACE:
		return "MULTIPART_BOUNDARY_WHITESPACE"
	case VARIABLE_MULTIPART_CRLF_LF_LINES:
		return "MULTIPART_CRLF_LF_LINES"
	case VARIABLE_MULTIPART_DATA_AFTER:
		return "MULTIPART_DATA_AFTER"
	case VARIABLE_MULTIPART_DATA_BEFORE:
		return "MULTIPART_DATA_BEFORE"
	case VARIABLE_MULTIPART_FILE_LIMIT_EXCEEDED:
		return "MULTIPART_FILE_LIMIT_EXCEEDED"
	case VARIABLE_MULTIPART_HEADER_FOLDING:
		return "MULTIPART_HEADER_FOLDING"
	case VARIABLE_MULTIPART_INVALID_HEADER_FOLDING:
		return "MULTIPART_INVALID_HEADER_FOLDING"
	case VARIABLE_MULTIPART_INVALID_PART:
		return "MULTIPART_INVALID_PART"
	case VARIABLE_MULTIPART_INVALID_QUOTING:
		return "MULTIPART_INVALID_QUOTING"
	case VARIABLE_MULTIPART_LF_LINE:
		return "MULTIPART_LF_LINE"
	case VARIABLE_MULTIPART_MISSING_SEMICOLON:
		return "MULTIPART_MISSING_SEMICOLON"
	case VARIABLE_MULTIPART_STRICT_ERROR:
		return "MULTIPART_STRICT_ERROR"
	case VARIABLE_MULTIPART_UNMATCHED_BOUNDARY:
		return "MULTIPART_UNMATCHED_BOUNDARY"
	case VARIABLE_OUTBOUND_DATA_ERROR:
		return "OUTBOUND_DATA_ERROR"
	case VARIABLE_PATH_INFO:
		return "PATH_INFO"
	case VARIABLE_QUERY_STRING:
		return "QUERY_STRING"
	case VARIABLE_REMOTE_ADDR:
		return "REMOTE_ADDR"
	case VARIABLE_REMOTE_HOST:
		return "REMOTE_HOST"
	case VARIABLE_REMOTE_PORT:
		return "REMOTE_PORT"
	case VARIABLE_REQBODY_ERROR:
		return "REQBODY_ERROR"
	case VARIABLE_REQBODY_ERROR_MSG:
		return "REQBODY_ERROR_MSG"
	case VARIABLE_REQBODY_PROCESSOR_ERROR:
		return "REQBODY_PROCESSOR_ERROR"
	case VARIABLE_REQBODY_PROCESSOR_ERROR_MSG:
		return "REQBODY_PROCESSOR_ERROR_MSG"
	case VARIABLE_REQBODY_PROCESSOR:
		return "REQBODY_PROCESSOR"
	case VARIABLE_REQUEST_BASENAME:
		return "REQUEST_BASENAME"
	case VARIABLE_REQUEST_BODY:
		return "REQUEST_BODY"
	case VARIABLE_REQUEST_BODY_LENGTH:
		return "REQUEST_BODY_LENGTH"
	case VARIABLE_REQUEST_FILENAME:
		return "REQUEST_FILENAME"
	case VARIABLE_REQUEST_LINE:
		return "REQUEST_LINE"
	case VARIABLE_REQUEST_METHOD:
		return "REQUEST_METHOD"
	case VARIABLE_REQUEST_PROTOCOL:
		return "REQUEST_PROTOCOL"
	case VARIABLE_REQUEST_URI:
		return "REQUEST_URI"
	case VARIABLE_REQUEST_URI_RAW:
		return "REQUEST_URI_RAW"
	case VARIABLE_RESOURCE:
		return "RESOURCE"
	case VARIABLE_RESPONSE_BODY:
		return "RESPONSE_BODY"
	case VARIABLE_RESPONSE_CONTENT_LENGTH:
		return "RESPONSE_CONTENT_LENGTH"
	case VARIABLE_RESPONSE_PROTOCOL:
		return "RESPONSE_PROTOCOL"
	case VARIABLE_RESPONSE_STATUS:
		return "RESPONSE_STATUS"
	case VARIABLE_SERVER_ADDR:
		return "SERVER_ADDR"
	case VARIABLE_SERVER_NAME:
		return "SERVER_NAME"
	case VARIABLE_SERVER_PORT:
		return "SERVER_PORT"
	case VARIABLE_SESSIONID:
		return "SESSIONID"
	case VARIABLE_RESPONSE_HEADERS_NAMES:
		return "RESPONSE_HEADERS_NAMES"
	case VARIABLE_REQUEST_HEADERS_NAMES:
		return "REQUEST_HEADERS_NAMES"
	case VARIABLE_USERID:
		return "USERID"
	case VARIABLE_ARGS:
		return "ARGS"
	case VARIABLE_ARGS_GET:
		return "ARGS_GET"
	case VARIABLE_ARGS_POST:
		return "ARGS_POST"
	case VARIABLE_FILES_SIZES:
		return "FILES_SIZES"
	case VARIABLE_FILES_NAMES:
		return "FILES_NAMES"
	case VARIABLE_FILES_TMP_CONTENT:
		return "FILES_TMP_CONTENT"
	case VARIABLE_MULTIPART_FILENAME:
		return "MULTIPART_FILENAME"
	case VARIABLE_MULTIPART_NAME:
		return "MULTIPART_NAME"
	case VARIABLE_MATCHED_VARS_NAMES:
		return "MATCHED_VARS_NAMES"
	case VARIABLE_MATCHED_VARS:
		return "MATCHED_VARS"
	case VARIABLE_FILES:
		return "FILES"
	case VARIABLE_REQUEST_COOKIES:
		return "REQUEST_COOKIES"
	case VARIABLE_REQUEST_HEADERS:
		return "REQUEST_HEADERS"
	case VARIABLE_RESPONSE_HEADERS:
		return "RESPONSE_HEADERS"
	case VARIABLE_GEO:
		return "GEO"
	case VARIABLE_REQUEST_COOKIES_NAMES:
		return "REQUEST_COOKIES_NAMES"
	case VARIABLE_FILES_TMPNAMES:
		return "FILES_TMPNAMES"
	case VARIABLE_ARGS_NAMES:
		return "ARGS_NAMES"
	case VARIABLE_ARGS_GET_NAMES:
		return "ARGS_GET_NAMES"
	case VARIABLE_ARGS_POST_NAMES:
		return "ARGS_POST_NAMES"
	case VARIABLE_TX:
		return "TX"
	case VARIABLE_GLOBAL:
		return "GLOBAL"
	case VARIABLE_IP:
		return "IP"
	case VARIABLE_SESSION:
		return "SESSION"
	case VARIABLE_USER:
		return "USER"
	case VARIABLE_DURATION:
		return "DURATION"
	case VARIABLE_RULE:
		return "RULE"
	}
	return ""
}
