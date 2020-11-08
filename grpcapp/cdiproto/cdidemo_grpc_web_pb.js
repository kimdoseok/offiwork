/**
 * @fileoverview gRPC-Web generated client stub for cdiproto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.cdiproto = require('./cdidemo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cdiproto.CdidemoClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cdiproto.CdidemoPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.UserRequest,
 *   !proto.cdiproto.ValidResponse>}
 */
const methodDescriptor_Cdidemo_AuthUser = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/AuthUser',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.UserRequest,
  proto.cdiproto.ValidResponse,
  /**
   * @param {!proto.cdiproto.UserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.ValidResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.UserRequest,
 *   !proto.cdiproto.ValidResponse>}
 */
const methodInfo_Cdidemo_AuthUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.ValidResponse,
  /**
   * @param {!proto.cdiproto.UserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.ValidResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.UserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.ValidResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.ValidResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.authUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/AuthUser',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_AuthUser,
      callback);
};


/**
 * @param {!proto.cdiproto.UserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.ValidResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.authUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/AuthUser',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_AuthUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.UserRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetUser = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetUser',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.UserRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.UserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.UserRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.UserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.UserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetUser',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetUser,
      callback);
};


/**
 * @param {!proto.cdiproto.UserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetUser',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetKeyCode = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetKeyCode',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetKeyCode = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getKeyCode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCode',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCode,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getKeyCode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCode',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetKeyCodeDetail = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetKeyCodeDetail',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetKeyCodeDetail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getKeyCodeDetail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCodeDetail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCodeDetail,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getKeyCodeDetail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCodeDetail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCodeDetail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetKeyCodeRef = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetKeyCodeRef',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetKeyCodeRef = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getKeyCodeRef =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCodeRef',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCodeRef,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getKeyCodeRef =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetKeyCodeRef',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetKeyCodeRef);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetICD10 = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetICD10',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetICD10 = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getICD10 =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICD10',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICD10,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getICD10 =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICD10',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICD10);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetICD10Detail = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetICD10Detail',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetICD10Detail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getICD10Detail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICD10Detail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICD10Detail,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getICD10Detail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICD10Detail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICD10Detail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetICDref = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetICDref',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetICDref = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getICDref =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICDref',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICDref,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getICDref =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetICDref',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetICDref);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetDRG = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetDRG',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetDRG = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getDRG =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRG',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRG,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getDRG =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRG',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRG);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetDRGDetail = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetDRGDetail',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetDRGDetail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getDRGDetail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRGDetail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRGDetail,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getDRGDetail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRGDetail',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRGDetail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodDescriptor_Cdidemo_GetDRGref = new grpc.web.MethodDescriptor(
  '/cdiproto.Cdidemo/GetDRGref',
  grpc.web.MethodType.UNARY,
  proto.cdiproto.SearchRequest,
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cdiproto.SearchRequest,
 *   !proto.cdiproto.JsonResponse>}
 */
const methodInfo_Cdidemo_GetDRGref = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cdiproto.JsonResponse,
  /**
   * @param {!proto.cdiproto.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cdiproto.JsonResponse.deserializeBinary
);


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cdiproto.JsonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cdiproto.JsonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cdiproto.CdidemoClient.prototype.getDRGref =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRGref',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRGref,
      callback);
};


/**
 * @param {!proto.cdiproto.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cdiproto.JsonResponse>}
 *     A native promise that resolves to the response
 */
proto.cdiproto.CdidemoPromiseClient.prototype.getDRGref =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cdiproto.Cdidemo/GetDRGref',
      request,
      metadata || {},
      methodDescriptor_Cdidemo_GetDRGref);
};


module.exports = proto.cdiproto;

