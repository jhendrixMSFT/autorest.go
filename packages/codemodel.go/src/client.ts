/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import { CodeModelError } from './errors.js';
import { MethodExample } from './examples.js';
import * as method from './method.js';
import * as pkg from './package.js';
import * as param from './param.js';
import * as result from './result.js';
import * as type from './type.js';

/** an SDK client */
export interface Client {
  /** the name of the client */
  name: string;

  /** any docs for the client */
  docs: type.Docs;

  /** the client options type */
  options: ClientOptionsType;

  /**
   * fields contains the ctor parameters that are
   * persisted as fields on the client type and might
   * also contain other fields that don't originate
   * from ctor params (e.g. the pipeline).
   * by convention, all fields that have their values
   * populated from ctor params (required or optional)
   * will have the same name as their originating param.
   */
  fields: Array<type.StructField>;

  /** all the constructors for this client, can be empty */
  constructors: Array<Constructor>;

  /** contains client methods. can be empty */
  methods: Array<MethodType>;

  /** contains any client accessor methods. can be empty */
  clientAccessors: Array<ClientAccessor>;

  /** the API version parameter, if applicable */
  apiVersion?: param.HeaderScalarParameter | param.PathScalarParameter | param.QueryScalarParameter | param.URIParameter;

  /**
   * templatedHost indicates that there's one or more URIParameters
   * required to construct the complete host. the parameters can
   * be solely on the client or span client and method params.
   */
  templatedHost?: string;

  /** the parent client in a hierarchical client */
  parent?: Client;
}

/** the possible types used for the client options type */
export type ClientOptionsType = type.ArmClientOptions | ClientOptionsParameter;

/** the possible client parameter types */
export type ClientParameter = ClientCredentialParameter | ClientEndpointParameter | param.MethodParameter;

/** a client method that returns a sub-client instance */
export interface ClientAccessor {
  /** the name of the client accessor method */
  name: string;

  /** the client returned by the accessor method */
  subClient: Client;
}

/** a credential constructor parameter */
export interface ClientCredentialParameter extends method.Parameter {
  kind: 'credentialParam';

  /** the parameter's type */
  type: type.TokenCredential;
}

/** an endpoint constructor parameter */
export interface ClientEndpointParameter extends method.Parameter {
  kind: 'endpointParam';

  /** endpoints are always strings */
  type: type.String;
}

/** the client options parameter type */
export interface ClientOptionsParameter {
  kind: 'clientOptions';

  /** the name of the type */
  name: string;

  /** any docs for the type */
  docs: type.Docs;

  /** the parameters that belong to this options */
  parameters: Array<param.MethodParameter>;
}

/** a client constructor function */
export interface Constructor {
  /** the name of the constructor function */
  name: string;

  /** the modeled parameters. can be empty */
  parameters: Array<ClientParameter>;

  /** any docs for the constructor */
  docs: type.Docs;
}

/** the possible values defining the "final state via" behavior for LROs */
export type FinalStateVia = 'azure-async-operation' | 'location' | 'operation-location' | 'original-uri';

/** the supported HTTP verbs */
export type HTTPMethod = 'delete' | 'get' | 'head' | 'patch' | 'post' | 'put';

/** the possible method types */
export type MethodType = LROMethod | LROPageableMethod | SyncMethod | PageableMethod;

/** a long-running operation method */
export interface LROMethod extends LROMethodBase {
  kind: 'lroMethod';
}

/** a long-running operation method that returns pages of responses */
export interface LROPageableMethod extends LROMethodBase, PageableMethodBase {
  kind: 'lroPageableMethod';
}

/** a synchronous method */
export interface SyncMethod extends HttpMethodBase {
  kind: 'method';
}

/** contains the names of the helper methods used to create a complete method implementation */
export interface MethodNaming {
  /** the name of the internal method for consumption by LROs/paging methods */
  internalMethod: string;

  /** the name of the internal method that creates the HTTP request */
  requestMethod: string;

  /** the name of the internal method that handles the HTTP response */
  responseMethod: string;
}

/**
 * the internal method used for fetching the next page for a PageableMethod.
 * It's unique from a regular Method as it's not exported and has no optional params/response envelope.
 * thus, it's not included in the array of methods for a client.
 */
export interface NextPageMethod {
  kind: 'nextPageMethod';

  /** the name of the next page method */
  name: string;
 
  /** the HTTP path used when creating the request */
  httpPath: string;

  /** the HTTP verb used when creating the request */
  httpMethod: HTTPMethod;

  /** any modeled parameters */
  parameters: Array<param.MethodParameter>;

  /** the complete list of successful HTTP status codes */
  httpStatusCodes: Array<number>;

  /** the method's receiver parameter */
  receiver: method.Receiver<Client>;

  apiVersions: Array<string>;
}

/** a synchronous method that returns pages of responses */
export interface PageableMethod extends PageableMethodBase {
  kind: 'pageableMethod';
}

/** narrows method to a LRO method type within the conditional block */
export function isLROMethod(method: MethodType): method is LROMethod | LROPageableMethod {
  return method.kind === 'lroMethod' || method.kind === 'lroPageableMethod';
}

/** narrows method to a pageable method type within the conditional block */
export function isPageableMethod(method: MethodType): method is LROPageableMethod | PageableMethod {
  return method.kind === 'lroPageableMethod' || method.kind === 'pageableMethod';
}

/** creates the ClientOptions type from the specified input */
export function newClientOptions(modelType: pkg.CodeModelType, clientName: string): ClientOptionsType {
  if (modelType === 'azure-arm') {
    return new type.ArmClientOptions();
  }

  const optionsTypeName = `${clientName}Options`;
  let options = new ClientOptionsParameter(optionsTypeName);
  options.docs.summary = `${optionsTypeName} contains the optional values for creating a [${clientName}]`;
  return options;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

interface LROMethodBase extends HttpMethodBase {
  finalStateVia?: FinalStateVia;

  operationLocationResultPath?: string;
}

interface HttpMethodBase extends method.Method<Client, result.ResponseEnvelope> {
  /** the HTTP path used when creating the request */
  httpPath: string;

  /** the HTTP verb used when creating the request */
  httpMethod: HTTPMethod;

  /** any modeled parameters. the ones we add to the generated code (context.Context etc) aren't included here */
  parameters: Array<param.MethodParameter>;

  /** the method options type for this method */
  optionalParamsGroup: param.ParameterGroup;

  /** the method's return type */
  returns: result.ResponseEnvelope;

  /** the complete list of successful HTTP status codes */
  httpStatusCodes: Array<number>;

  /** naming info for the internal hepler methods for which this method depends */
  naming: MethodNaming;

  /** the API version(s) used during ingestion. can be empty */
  apiVersions: Array<string>;

  /** any examples for this method */
  examples: Array<MethodExample>;
}

interface PageableMethodBase extends HttpMethodBase {
  nextLinkName?: string;

  nextPageMethod?: NextPageMethod;
}

class HttpMethodBase extends method.Method<Client, result.ResponseEnvelope> implements HttpMethodBase {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new CodeModelError('statusCodes cannot be empty');
    }
    super(name, new method.Receiver('client', client, false));
    this.apiVersions = new Array<string>();
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.name = name;
    this.naming = naming;
    this.parameters = new Array<param.MethodParameter>();
    this.examples = [];
    this.docs = {};
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Client implements Client {
  constructor(name: string, docs: type.Docs, options: ClientOptionsType) {
    this.name = name;
    this.constructors = new Array<Constructor>();
    this.docs = docs;
    this.methods = new Array<MethodType>();
    this.clientAccessors = new Array<ClientAccessor>();
    this.fields = new Array<type.StructField>();
    this.options = options;
  }
}

export class ClientAccessor implements ClientAccessor {
  constructor(name: string, subClient: Client) {
    this.name = name;
    this.subClient = subClient;
  }
}

export class ClientCredentialParameter extends method.Parameter implements ClientCredentialParameter {
  constructor(name: string, type: type.TokenCredential) {
    super(name, type, true);
    this.kind = 'credentialParam';
  }
}

export class ClientEndpointParameter extends method.Parameter implements ClientEndpointParameter {
  constructor(name: string) {
    super(name, new type.String(), true);
    this.kind = 'endpointParam';
  }
}

export class ClientOptionsParameter implements ClientOptionsParameter {
  constructor(name: string) {
    this.kind = 'clientOptions';
    this.name = name;
    this.docs = {};
    this.parameters = new Array<param.MethodParameter>();
  }
}

export class Constructor implements Constructor {
  constructor(name: string) {
    this.name = name;
    this.docs = {};
    this.parameters = new Array<ClientParameter>();
  }
}

export class LROMethod extends HttpMethodBase implements LROMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'lroMethod';
  }
}

export class LROPageableMethod extends HttpMethodBase implements LROPageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'lroPageableMethod';
  }
}

export class SyncMethod extends HttpMethodBase implements SyncMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'method';
  }
}

export class MethodNaming implements MethodNaming {
  constructor(internalMethod: string, requestMethod: string, responseMethod: string) {
    this.internalMethod = internalMethod;
    this.requestMethod = requestMethod;
    this.responseMethod = responseMethod;
  }
}

export class NextPageMethod implements NextPageMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>) {
    if (statusCodes.length === 0) {
      throw new CodeModelError('statusCodes cannot be empty');
    }
    this.kind = 'nextPageMethod';
    this.apiVersions = new Array<string>();
    this.receiver = new method.Receiver('client', client, false);
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.name = name;
    this.parameters = new Array<param.MethodParameter>();
  }
}

export class PageableMethod extends HttpMethodBase implements PageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'pageableMethod';
  }
}
