/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/gocodemodel.js';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as helpers from '../helpers.js';
import { fixUpMethodName } from '../operations.js';
import { ImportManager } from '../imports.js';
import { generateServerInternal, RequiredHelpers } from './internal.js';

// contains the generated content for all servers and the required helpers
export class ServerContent {
  readonly servers: Array<OperationGroupContent>;
  readonly internals: string;

  constructor(servers: Array<OperationGroupContent>, internals: string) {
    this.servers = servers;
    this.internals = internals;
  }
}

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// used to track the helpers we need to emit. they're all false by default.
const requiredHelpers = new RequiredHelpers();

export function getServerName(client: go.Client): string {
  // for the fake server, we use the suffix Server instead of Client
  return capitalize(client.name.replace(/[C|c]lient$/, 'Server'));
}

function isMethodInternal(method: go.Method): boolean {
  return !!method.name.match(/^[a-z]{1}/);
}

export async function generateServers(codeModel: go.CodeModel): Promise<ServerContent> {
  const operations = new Array<OperationGroupContent>();
  const clientPkg = codeModel.packageName;
  for (const client of values(codeModel.clients)) {
    if (client.clientAccessors.length === 0 && values(client.methods).all(method => { return isMethodInternal(method)})) {
      // client has no client accessors and no exported methods, skip it
      continue;
    }

    // the list of packages to import
    const imports = new ImportManager();

    // add standard imports
    imports.add('errors');
    imports.add('fmt');
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

    const serverName = getServerName(client);

    let content: string;
    content = `// ${serverName} is a fake server for instances of the ${clientPkg}.${client.name} type.\n`;
    content += `type ${serverName} struct{\n`;

    // we might remove some operations from the list
    const finalMethods = new Array<go.Method>();
    let countLROs = 0;
    let countPagers = 0;

    // add server transports for client accessors
    // we might remove some clients from the list
    const finalSubClients = new Array<go.Client>();
    for (const clientAccessor of client.clientAccessors) {
      if (values(clientAccessor.subClient.methods).all(method => { return isMethodInternal(method)})) {
        // client has no exported methods, skip it
        continue;
      }
      const serverName = getServerName(clientAccessor.subClient);
      content += `\t// ${serverName} contains the fakes for client ${clientAccessor.subClient.name}\n`;
      content += `\t${serverName} ${serverName}\n\n`;
      finalSubClients.push(clientAccessor.subClient);
    }

    for (const method of values(client.methods)) {
      if (isMethodInternal(method)) {
        // method isn't exported, don't create a fake for it
        continue;
      }

      let serverResponse: string;

      if (go.isLROMethod(method)) {
        let respType = `${clientPkg}.${method.responseEnvelope.name}`;
        if (go.isPageableMethod(method)) {
          respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
        }
        serverResponse = `resp azfake.PollerResponder[${respType}], errResp azfake.ErrorResponder`;
      } else if (go.isPageableMethod(method)) {
        serverResponse = `resp azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
      } else {
        serverResponse = `resp azfake.Responder[${clientPkg}.${method.responseEnvelope.name}], errResp azfake.ErrorResponder`;
      }

      const operationName = fixUpMethodName(method);
      content += `\t// ${operationName} is the fake for method ${client.name}.${operationName}\n`;
      const successCodes = new Array<string>();
      if (method.responseEnvelope.result && go.isAnyResult(method.responseEnvelope.result)) {
        for (const httpStatus of values(method.httpStatusCodes)) {
          const result = method.responseEnvelope.result.httpStatusCodeType[httpStatus];
          if (!result) {
            // the operation contains a mix of schemas and non-schema responses
            successCodes.push(`${helpers.formatStatusCode(httpStatus)} (no return type)`);
            continue;
          }
          successCodes.push(`${helpers.formatStatusCode(httpStatus)} (returns ${go.getTypeDeclaration(result, clientPkg)})`);
        }
        content += '\t// HTTP status codes to indicate success:\n';
        for (const successCode of successCodes) {
          content += `\t//   - ${successCode}\n`;
        }
      } else {
        for (const statusCode of values(method.httpStatusCodes)) {
          successCodes.push(`${helpers.formatStatusCode(statusCode)}`);
        }
        content += `\t// HTTP status codes to indicate success: ${successCodes.join(', ')}\n`;
      }
      content += `\t${operationName} func(${getAPIParametersSig(method, imports, clientPkg)}) (${serverResponse})\n\n`;
      finalMethods.push(method);
      if (go.isLROMethod(method)) {
        ++countLROs;
      } else if (go.isPageableMethod(method)) {
        ++countPagers;
      }
    }

    content += '}\n\n';

    ///////////////////////////////////////////////////////////////////////////

    const serverTransport = `${serverName}Transport`;

    content += `// New${serverTransport} creates a new instance of ${serverTransport} with the provided implementation.\n`;
    content += `// The returned ${serverTransport} instance is connected to an instance of ${clientPkg}.${client.name} via the\n`;
    content += '// azcore.ClientOptions.Transporter field in the client\'s constructor parameters.\n';
    content += `func New${serverTransport}(srv *${serverName}) *${serverTransport} {\n`;
    if (countLROs === 0 && countPagers === 0) {
      content += `\treturn &${serverTransport}{srv: srv}\n}\n\n`;
    } else {
      content += `\treturn &${serverTransport}{\n\t\tsrv: srv,\n`;
      for (const method of values(finalMethods)) {
        let respType = `${clientPkg}.${method.responseEnvelope.name}`;
        if (go.isLROMethod(method)) {
          if (go.isPageableMethod(method)) {
            respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
          }
          requiredHelpers.tracker = true;
          content += `\t\t${uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PollerResponder[${respType}]](),\n`;
        } else if (go.isPageableMethod(method)) {
          requiredHelpers.tracker = true;
          content += `\t\t${uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PagerResponder[${respType}]](),\n`;
        }
      }
      content += '\t}\n}\n\n';
    }

    content += `// ${serverTransport} connects instances of ${clientPkg}.${client.name} to instances of ${serverName}.\n`;
    content += `// Don't use this type directly, use New${serverTransport} instead.\n`;
    content += `type ${serverTransport} struct {\n`;
    content += `\tsrv *${serverName}\n`;

    // add server transports for client accessors
    if (finalSubClients.length > 0) {
      requiredHelpers.initServer = true;
      imports.add('sync');
      content += '\ttrMu sync.Mutex\n';
      for (const subClient of finalSubClients) {
        const serverName = getServerName(subClient);
        content += `\ttr${serverName} *${serverName}Transport\n`;
      }
    }

    for (const method of values(finalMethods)) {
      // create state machines for any pager/poller operations
      let respType = `${clientPkg}.${method.responseEnvelope.name}`;
      if (go.isLROMethod(method)) {
        if (go.isPageableMethod(method)) {
          respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
        }
        requiredHelpers.tracker = true;
        content +=`\t${uncapitalize(fixUpMethodName(method))} *tracker[azfake.PollerResponder[${respType}]]\n`;
      } else if (go.isPageableMethod(method)) {
        requiredHelpers.tracker = true;
        content += `\t${uncapitalize(fixUpMethodName(method))} *tracker[azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]]\n`;
      }
    }
    content += '}\n\n';

    content += generateServerTransportDo(serverTransport, client, finalSubClients, finalMethods);
    content += generateServerTransportClientDispatch(serverTransport, finalSubClients, imports);
    content += generateServerTransportMethodDispatch(serverTransport, client, finalMethods);
    content += generateServerTransportMethods(codeModel, serverTransport, finalMethods, imports);

    ///////////////////////////////////////////////////////////////////////////

    // stitch everything together
    let text = helpers.contentPreamble(codeModel, 'fake');
    text += imports.text();
    text += content;
    operations.push(new OperationGroupContent(serverName, text));
  }
  return new ServerContent(operations, generateServerInternal(codeModel, requiredHelpers));
}

// method names for fakes dispatching
const dispatchMethodFake = 'dispatchToMethodFake';
const dispatchToClientFake = 'dispatchToClientFake';

function generateServerTransportDo(serverTransport: string, client: go.Client, finalSubClients: Array<go.Client>, finalMethods: Array<go.Method>): string {
  const receiverName = serverTransport[0].toLowerCase();
  let content = `// Do implements the policy.Transporter interface for ${serverTransport}.\n`;
  content += `func (${receiverName} *${serverTransport}) Do(req *http.Request) (*http.Response, error) {\n`;
  content += '\trawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n';
  content += '\tmethod, ok := rawMethod.(string)\n';
  content += '\tif !ok {\n\t\treturn nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n\t}\n\n';

  if (finalSubClients.length > 0 && finalMethods.length > 0) {
    // client contains client accessors and methods.
    // if the method isn't for this client, dispatch to the correct client
    content += `\tif client := method[:strings.Index(method, ".")]; client != "${client.name}" {\n`;
    content += `\t\treturn ${receiverName}.${dispatchToClientFake}(req, client)\n\t}\n`;
    // else dispatch to our method fakes
    content += `\treturn ${receiverName}.${dispatchMethodFake}(req, method)\n`;
  } else if (finalSubClients.length > 0) {
    content += `\treturn ${receiverName}.${dispatchToClientFake}(req, method[:strings.Index(method, ".")])\n`;
  } else {
    content += `\treturn ${receiverName}.${dispatchMethodFake}(req, method)\n`;
  }
  content += '}\n\n'; // end Do
  return content;
}

function generateServerTransportClientDispatch(serverTransport: string, subClients: Array<go.Client>, imports: ImportManager): string {
  if (subClients.length === 0) {
    return '';
  }

  const receiverName = serverTransport[0].toLowerCase();
  imports.add('strings');
  let content = `func (${receiverName} *${serverTransport}) ${dispatchToClientFake}(req *http.Request, client string) (*http.Response, error) {\n`;
  content += '\tvar resp *http.Response\n\tvar err error\n\n';
  content += '\tswitch client {\n';
  for (const subClient of subClients) {
    content += `\tcase "${subClient.name}":\n`;
    const serverName = getServerName(subClient);
    content += `\t\tinitServer(&${receiverName}.trMu, &${receiverName}.tr${serverName}, func() *${serverName}Transport {\n\t\treturn New${serverName}Transport(&${receiverName}.srv.${serverName}) })\n`;
    content += `\t\tresp, err = ${receiverName}.tr${serverName}.Do(req)\n`;
  }
  content += '\tdefault:\n\t\terr = fmt.Errorf("unhandled client %s", client)\n';
  content += '\t}\n\n'; // end switch
  content += '\treturn resp, err\n}\n\n';
  return content;
}

function generateServerTransportMethodDispatch(serverTransport: string, client: go.Client, finalMethods: Array<go.Method>): string {
  if (finalMethods.length === 0) {
    return '';
  }

  const receiverName = serverTransport[0].toLowerCase();
  let content = `func (${receiverName} *${serverTransport}) ${dispatchMethodFake}(req *http.Request, method string) (*http.Response, error) {\n`;
  content += '\tvar resp *http.Response\n';
  content += '\tvar err error\n\n';
  content += '\tswitch method {\n';

  for (const method of values(finalMethods)) {
    const operationName = fixUpMethodName(method);
    content += `\tcase "${client.name}.${operationName}":\n`;
    content += `\t\tresp, err = ${receiverName}.dispatch${operationName}(req)\n`;
  }

  content += '\tdefault:\n\t\terr = fmt.Errorf("unhandled API %s", method)\n';
  content += '\t}\n\n'; // end switch
  content += '\treturn resp, err\n}\n\n';
  return content;
}

function generateServerTransportMethods(codeModel: go.CodeModel, serverTransport: string, finalMethods: Array<go.Method>, imports: ImportManager): string {
  if (finalMethods.length === 0) {
    return '';
  }

  imports.add(helpers.getParentImport(codeModel));
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');

  const receiverName = serverTransport[0].toLowerCase();
  let content = '';
  for (const method of values(finalMethods)) {
    content += `func (${receiverName} *${serverTransport}) dispatch${fixUpMethodName(method)}(req *http.Request) (*http.Response, error) {\n`;
    content += `\tif ${receiverName}.srv.${fixUpMethodName(method)} == nil {\n`;
    content += `\t\treturn nil, &nonRetriableError{errors.New("fake for method ${fixUpMethodName(method)} not implemented")}\n\t}\n`;

    if (go.isLROMethod(method)) {
      // must check LRO before pager as you can have paged LROs
      content += dispatchForLROBody(codeModel.packageName, receiverName, method, imports);
    } else if (go.isPageableMethod(method)) {
      content += dispatchForPagerBody(codeModel.packageName, receiverName, method, imports);
    } else {
      content += dispatchForOperationBody(codeModel.packageName, receiverName, method, imports);
      content += '\trespContent := server.GetResponseContent(respr)\n';
      const formattedStatusCodes = helpers.formatStatusCodes(method.httpStatusCodes);
      content += `\tif !contains([]int{${formattedStatusCodes}}, respContent.HTTPStatus) {\n`;
      content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", respContent.HTTPStatus)}\n\t}\n`;
      if (!method.responseEnvelope.result || go.isHeadAsBooleanResult(method.responseEnvelope.result)) {
        content += '\tresp, err := server.NewResponse(respContent, req, nil)\n';
      } else if (go.isAnyResult(method.responseEnvelope.result)) {
        content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)}, req)\n`;
      } else if (go.isBinaryResult(method.responseEnvelope.result)) {
        content += '\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n';
        content += `\t\tBody: server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)},\n`;
        content += '\t\tContentType: req.Header.Get("Content-Type"),\n';
        content += '\t})\n';
      } else if (go.isMonomorphicResult(method.responseEnvelope.result)) {
        if (go.isBytesType(method.responseEnvelope.result.monomorphicType)) {
          const encoding = method.responseEnvelope.result.monomorphicType.encoding;
          content += `\tresp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)}, runtime.Base64${encoding}Format, req)\n`;
        } else if (go.isSliceType(method.responseEnvelope.result.monomorphicType) && method.responseEnvelope.result.monomorphicType.rawJSONAsBytes) {
          imports.add('bytes');
          imports.add('io');
          content += '\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n';
          content += '\t\tBody: io.NopCloser(bytes.NewReader(server.GetResponse(respr).RawJSON)),\n';
          content += '\t\tContentType: "application/json",\n\t})\n';
        } else {
          let respField = `.${getResultFieldName(method.responseEnvelope.result)}`;
          if (method.responseEnvelope.result.format === 'XML' && go.isSliceType(method.responseEnvelope.result.monomorphicType)) {
            // for XML array responses we use the response type directly as it has the necessary XML tag for proper marshalling
            respField = '';
          }
          let responseField = `server.GetResponse(respr)${respField}`;
          if (go.isTimeType(method.responseEnvelope.result.monomorphicType)) {
            responseField = `(*${method.responseEnvelope.result.monomorphicType.dateTimeFormat})(${responseField})`;
          }
          content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, ${responseField}, req)\n`;
        }
      } else if (go.isModelResult(method.responseEnvelope.result) || go.isPolymorphicResult(method.responseEnvelope.result)) {
        const respField = `.${getResultFieldName(method.responseEnvelope.result)}`;
        const responseField = `server.GetResponse(respr)${respField}`;
        content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, ${responseField}, req)\n`;
      }

      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';

      // propagate any header response values into the *http.Response
      for (const header of values(method.responseEnvelope.headers)) {
        if (go.isHeaderMapResponse(header)) {
          content += `\tfor k, v := range server.GetResponse(respr).${header.fieldName} {\n`;
          content += '\t\tif v != nil {\n';
          content += `\t\t\tresp.Header.Set("${header.collectionPrefix}"+k, *v)\n`;
          content += '\t\t}\n';
          content += '\t}\n';
        } else {
          content += `\tif val := server.GetResponse(respr).${header.fieldName}; val != nil {\n`;
          content += `\t\tresp.Header.Set("${header.headerName}", ${helpers.formatValue('val', header.type, imports, true)})\n\t}\n`;
        }
      }

      content += '\treturn resp, nil\n';
    }
    content += '}\n\n';
  }

  return content;
}

function dispatchForOperationBody(clientPkg: string, receiverName: string, method: go.Method, imports: ImportManager): string {
  const numPathParams = values(method.parameters).where((each: go.Parameter) => { return go.isPathParameter(each) && !helpers.isLiteralParameter(each); }).count();
  let content = '';
  if (numPathParams > 0) {
    imports.add('regexp');
    content += `\tconst regexStr = \`${createPathParamsRegex(method)}\`\n`;
    content += '\tregex := regexp.MustCompile(regexStr)\n';
    content += '\tmatches := regex.FindStringSubmatch(req.URL.EscapedPath())\n';
    content += `\tif matches == nil || len(matches) < ${numPathParams} {\n`;
    content += '\t\treturn nil, fmt.Errorf("failed to parse path %s", req.URL.Path)\n\t}\n';
  }
  if (values(method.parameters).where((each: go.Parameter) => { return go.isQueryParameter(each) && each.location === 'method' && !helpers.isLiteralParameter(each); }).any()) {
    content += '\tqp := req.URL.Query()\n';
  }
  const bodyParam = <go.BodyParameter | undefined>values(method.parameters).where((each: go.Parameter) => { return go.isBodyParameter(each) || go.isFormBodyParameter(each) || go.isMultipartFormBodyParameter(each); }).first();
  if (!bodyParam) {
    // no body, just headers and/or query params
  } else if (go.isMultipartFormBodyParameter(bodyParam)) {
    imports.add('io');
    imports.add('mime');
    imports.add('mime/multipart');
    content += '\t_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))\n';
    content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    content += '\treader := multipart.NewReader(req.Body, params["boundary"])\n';
    for (const param of values(method.parameters)) {
      if (go.isMultipartFormBodyParameter(param)) {
        let pkgPrefix = '';
        if (go.isConstantType(param.type) || go.isModelType(param.type)) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.name} ${pkgPrefix}${go.getTypeDeclaration(param.type)}\n`;
      }
    }

    content += '\tfor {\n';
    content += '\t\tvar part *multipart.Part\n';
    content += '\t\tpart, err = reader.NextPart()\n';
    content += '\t\tif errors.Is(err, io.EOF) {\n\t\t\tbreak\n';
    content += '\t\t} else if err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
    content += '\t\tvar content []byte\n';
    content += '\t\tswitch fn := part.FormName(); fn {\n';

    // specify boolTarget if parsing bools happens in place.
    // in this case, the err check is omitted and assumed to happen elsewhere.
    // the parsed value is in a local var named parsed.
    const parsePrimitiveType = function(typeName: go.PrimitiveTypeName, boolTarget?: string): string {
      const parseResults = 'parsed, parseErr';
      let parsingCode = '';
      imports.add('strconv');
      switch (typeName) {
        case 'bool':
          if (boolTarget) {
            parsingCode = `\t\t\t${boolTarget} = strconv.ParseBool(string(content))\n`;
          } else {
            parsingCode = `\t\t\t${parseResults} := strconv.ParseBool(string(content))\n`;
          }
          break;
        case 'float32':
        case 'float64':
          parsingCode = `\t\t\t${parseResults} := strconv.ParseFloat(string(content), ${helpers.getBitSizeForNumber(typeName)})\n`;
          break;
        case 'int8':
        case 'int16':
        case 'int32':
        case 'int64':
          parsingCode = `\t\t\t${parseResults} := strconv.ParseInt(string(content), 10, ${helpers.getBitSizeForNumber(typeName)})\n`;
          break;
        default:
          throw new Error(`unhandled multipart parameter primitive type ${typeName}`);
      }
      if (!boolTarget) {
        parsingCode += '\t\t\tif parseErr != nil {\n\t\t\t\treturn nil, parseErr\n\t\t\t}\n';
      }
      return parsingCode;
    };

    const isMultipartContentType = function(type: go.PossibleType): type is go.QualifiedType {
      type = helpers.recursiveUnwrapMapSlice(type);
      return (go.isQualifiedType(type) && type.typeName === 'MultipartContent');
    };

    const emitCase = function(caseValue: string, paramVar: string, type: go.PossibleType): string {
      let caseContent = `\t\tcase "${caseValue}":\n`;
      caseContent += '\t\t\tcontent, err = io.ReadAll(part)\n';
      caseContent += '\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n';
      let assignedValue: string | undefined;
      if (go.isModelType(helpers.recursiveUnwrapMapSlice(type))) {
        imports.add('encoding/json');
        caseContent += `\t\t\tif err = json.Unmarshal(content, &${paramVar}); err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n`;
      } else if (go.isQualifiedType(type) && type.typeName === 'ReadSeekCloser') {
        imports.add('bytes');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        assignedValue = 'streaming.NopCloser(bytes.NewReader(content))';
      } else if (go.isConstantType(type)) {
        let from: string;
        switch (type.type) {
          case 'bool':
          case 'float32':
          case 'float64':
          case 'int32':
          case 'int64':
            caseContent += parsePrimitiveType(type.type);
            from = 'parsed';
            break;
          case 'string':
            from = 'content';
            break;
        }
        assignedValue = `${clientPkg}.${type.name}(${from})`;
      } else if (go.isPrimitiveType(type)) {
        switch (type.typeName) {
          case 'bool':
            imports.add('strconv');
            // ParseBool happens in place, so no need to set assignedValue
            caseContent += parsePrimitiveType(type.typeName, `${paramVar}, err`);
            break;
          case 'float32':
          case 'float64':
          case 'int8':
          case 'int16':
          case 'int32':
          case 'int64':
            caseContent += parsePrimitiveType(type.typeName);
            assignedValue = `${type.typeName}(parsed)`;
            break;
          case 'string':
            assignedValue = 'string(content)';
            break;
          default:
            throw new Error(`unhandled multipart parameter primitive type ${type.typeName}`);
        }
      } else if (isMultipartContentType(type)) {
        imports.add('bytes');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        const bodyContent = 'streaming.NopCloser(bytes.NewReader(content))';
        const contentType = 'part.Header.Get("Content-Type")';
        const filename = 'part.FileName()';
        if (go.isSliceType(type)) {
          caseContent += `\t\t\t${paramVar} = append(${paramVar}, streaming.MultipartContent{\n`;
          caseContent += `\t\t\t\tBody: ${bodyContent},\n`;
          caseContent += `\t\t\t\tContentType: ${contentType},\n`;
          caseContent += `\t\t\t\tFilename: ${filename},\n`;
          caseContent += '\t\t\t})\n';
        } else {
          caseContent += `\t\t\t${paramVar}.Body = ${bodyContent}\n`;
          caseContent += `\t\t\t${paramVar}.ContentType = ${contentType}\n`;
          caseContent += `\t\t\t${paramVar}.Filename = ${filename}\n`;
        }
      } else if (go.isSliceType(type)) {
        if (go.isQualifiedType(type.elementType) && type.elementType.typeName === 'ReadSeekCloser') {
          imports.add('bytes');
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
          assignedValue = `append(${paramVar}, streaming.NopCloser(bytes.NewReader(content)))`;
        } else {
          throw new Error(`uhandled multipart parameter array element type ${go.getTypeDeclaration(type.elementType)}`);
        }
      } else {
        throw new Error(`uhandled multipart parameter type ${go.getTypeDeclaration(type)}`);
      }
      if (assignedValue) {
        caseContent += `\t\t\t${paramVar} = ${assignedValue}\n`;
      }
      return caseContent;
    };

    for (const param of values(method.parameters)) {
      if (go.isMultipartFormBodyParameter(param)) {
        if (go.isModelType(param.type)) {
          for (const field of param.type.fields) {
            content += emitCase(field.serializedName, `${param.name}.${field.name}`, field.type);
          }
        } else {
          content += emitCase(param.name, param.name, param.type);
        }
      }
    }

    content += '\t\tdefault:\n\t\t\treturn nil, fmt.Errorf("unexpected part %s", fn)\n';
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (go.isFormBodyParameter(bodyParam)) {
    for (const param of values(method.parameters)) {
      if (go.isFormBodyParameter(param)) {
        let pkgPrefix = '';
        if (go.isConstantType(param.type)) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.name} ${pkgPrefix}${go.getTypeDeclaration(param.type)}\n`;
      }
    }
    content += '\tif err := req.ParseForm(); err != nil {\n\t\treturn nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}\n\t}\n';
    content += '\tfor key := range req.Form {\n';
    content += '\t\tswitch key {\n';
    for (const param of values(method.parameters)) {
      if (go.isFormBodyParameter(param)) {
        content += `\t\tcase "${param.formDataName}":\n`;
        let assignedValue: string;
        if (go.isConstantType(param.type)) {
          assignedValue = `${go.getTypeDeclaration(param.type, clientPkg)}(req.FormValue(key))`;
        } else if (go.isPrimitiveType(param.type) && param.type.typeName === 'string') {
          assignedValue = 'req.FormValue(key)';
        } else {
          throw new Error(`uhandled form parameter type ${go.getTypeDeclaration(param.type)}`);
        }
        content += `\t\t\t${param.name} = ${assignedValue}\n`;
      }
    }
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (bodyParam.bodyFormat === 'binary') {
    // nothing to do for binary media type
  } else if (bodyParam.bodyFormat === 'Text') {
    if (bodyParam && !helpers.isLiteralParameter(bodyParam)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      content += '\tbody, err := server.UnmarshalRequestAsText(req)\n';
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    }
  } else if (bodyParam.bodyFormat === 'JSON' || bodyParam.bodyFormat === 'XML') {
    if (bodyParam && !helpers.isLiteralParameter(bodyParam)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      if (go.isBytesType(bodyParam.type)) {
        content += `\tbody, err := server.UnmarshalRequestAsByteArray(req, runtime.Base64${bodyParam.type.encoding}Format)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else if (go.isSliceType(bodyParam.type) && bodyParam.type.rawJSONAsBytes) {
        content += '\tbody, err := io.ReadAll(req.Body)\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += '\treq.Body.Close()\n';
      } else if (go.isInterfaceType(bodyParam.type)) {
        requiredHelpers.readRequestBody = true;
        content += '\traw, err := readRequestBody(req)\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += `\tbody, err := unmarshal${bodyParam.type.name}(raw)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else {
        let bodyTypeName = go.getTypeDeclaration(bodyParam.type, clientPkg);
        if (go.isTimeType(bodyParam.type)) {
          bodyTypeName = bodyParam.type.dateTimeFormat;
        }
        content += `\tbody, err := server.UnmarshalRequestAs${bodyParam.bodyFormat}[${bodyTypeName}](req)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      }
    }
  }

  const result = parseHeaderPathQueryParams(clientPkg, method, imports);
  content += result.content;

  const apiCall = `:= ${receiverName}.srv.${fixUpMethodName(method)}(${populateApiParams(clientPkg, method, result.params, imports)})`;
  if (go.isPageableMethod(method) && !go.isLROMethod(method)) {
    content += `resp ${apiCall}\n`;
    return content;
  }
  content += `\trespr, errRespr ${apiCall}\n`;
  content += '\tif respErr := server.GetError(errRespr, req); respErr != nil {\n';
  content += '\t\treturn nil, respErr\n\t}\n';
  return content;
}

function dispatchForLROBody(clientPkg: string, receiverName: string, method: go.Method, imports: ImportManager): string {
  const operationName = fixUpMethodName(method);
  const localVarName = uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\t${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `\tif ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, method, imports);
  content += `\t\t${localVarName} = &respr\n`;
  content += `\t\t${operationStateMachine}.add(req, ${localVarName})\n`;
  content += '\t}\n\n';

  content += `\tresp, err := server.PollerResponderNext(${localVarName}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';

  const formattedStatusCodes = helpers.formatStatusCodes(method.httpStatusCodes);
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PollerResponderMore(${localVarName}) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n\t}\n\n`;
  content += '\treturn resp, nil\n';
  return content;
}

function dispatchForPagerBody(clientPkg: string, receiverName: string, method: go.PageableMethod, imports: ImportManager): string {
  const operationName = fixUpMethodName(method);
  const localVarName = uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\t${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `\tif ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, method, imports);
  content += `\t\t${localVarName} = &resp\n`;
  content += `\t\t${operationStateMachine}.add(req, ${localVarName})\n`;
  if (method.nextLinkName) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    content += `\t\tserver.PagerResponderInjectNextLinks(${localVarName}, req, func(page *${clientPkg}.${method.responseEnvelope.name}, createLink func() string) {\n`;
    content += `\t\t\tpage.${method.nextLinkName} = to.Ptr(createLink())\n`;
    content += '\t\t})\n';
  }
  content += '\t}\n'; // end if
  content += `\tresp, err := server.PagerResponderNext(${localVarName}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';

  const formattedStatusCodes = helpers.formatStatusCodes(method.httpStatusCodes);
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PagerResponderMore(${localVarName}) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n\t}\n`;
  content += '\treturn resp, nil\n';
  return content;
}

function sanitizeRegexpCaptureGroupName(name: string): string {
  // dash '-' characters are not allowed so replace them with '_'
  return name.replace('-', '_');
}

function createPathParamsRegex(method: go.Method): string {
  // "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"
  // each path param will replaced with a regex capture.
  // note that some path params are optional.
  let urlPath = method.httpPath;
  // escape any characters in the path that could be interpreted as regex tokens
  // per RFC3986, these are the pchars that also double as regex tokens
  // . $ * + ()
  urlPath = urlPath.replace(/([.$*+()])/g, '\\$1');
  for (const param of values(method.parameters)) {
    if (!go.isPathParameter(param)) {
      continue;
    }
    const toReplace = `{${param.pathSegment}}`;
    let replaceWith = `(?P<${sanitizeRegexpCaptureGroupName(param.pathSegment)}>[!#&$-;=?-\\[\\]_a-zA-Z0-9~%@]+)`;
    if (param.paramType === 'optional' || param.paramType === 'flag') {
      replaceWith += '?';
    }
    urlPath = urlPath.replace(toReplace, replaceWith);
  }
  return urlPath;
}

interface parseResult {
  // contains the param parsing code
  content: string;

  // maps a param name to the var containing the "final" value.
  // only params that required parsing/casting will have an entry.
  params: Map<string, string>;
}

// parses header/path/query params as required.
// returns the parsing code and the params that contain the parsed values.
function parseHeaderPathQueryParams(clientPkg: string, method: go.Method, imports: ImportManager): parseResult {
  let content = '';
  const paramValues = new Map<string, string>();

  const createLocalVariableName = function(param: go.Parameter, suffix: string): string {
    const paramName = `${uncapitalize(param.name)}${suffix}`;
    paramValues.set(param.name, paramName);
    return paramName;
  };

  const emitNumericConversion = function(src: string, type: 'float32' | 'float64' | 'int32' | 'int64'): string {
    imports.add('strconv');
    let precision: '32' | '64' = '32';
    if (type === 'float64' || type === 'int64') {
      precision = '64';
    }
    let parseType: 'Int' | 'Float' = 'Int';
    let base = '10, ';
    if (type === 'float32' || type === 'float64') {
      parseType = 'Float';
      base = '';
    }
    return `strconv.Parse${parseType}(${src}, ${base}${precision})`;
  };

  // track the param groups that need to be instantiated/populated.
  // we track the params separately as it might be a subset of ParameterGroup.params
  const paramGroups = new Map<go.ParameterGroup, Array<go.Parameter>>();

  for (const param of values(consolidateHostParams(method.parameters))) {
    if (param.location === 'client' || helpers.isLiteralParameter(param)) {
      // client params and parameter literals aren't passed to APIs
      continue;
    }
    if (go.isResumeTokenParameter(param)) {
      // skip the ResumeToken param as we don't send that back to the caller
      continue;
    }

    // NOTE: param group check must happen before skipping body params.
    // this is to handle the case where the body param is grouped/optional
    if (param.group) {
      if (!paramGroups.has(param.group)) {
        paramGroups.set(param.group, new Array<go.Parameter>());
      }
      const params = paramGroups.get(param.group);
      params!.push(param);
    }

    if (go.isBodyParameter(param) || go.isFormBodyParameter(param) || go.isMultipartFormBodyParameter(param)) {
      // body params will be unmarshalled, no need for parsing.
      continue;
    }

    // paramValue is initialized with the "raw" source value.
    // e.g. getHeaderValue(...), qp.Get("foo") etc
    // since path/query params need to be unescaped, the value
    // of paramValue will be updated with the var name that
    // contains the unescaped value.
    let paramValue = getRawParamValue(param);

    // path/query params might be escaped, so we need to unescape them first.
    // must handle query collections first as it's a superset of query param.
    if (go.isQueryCollectionParameter(param) && param.collectionFormat === 'multi') {
      imports.add('net/url');
      const escapedParam = createLocalVariableName(param, 'Escaped');
      content += `\t${escapedParam} := ${paramValue}\n`;
      let paramVar = createLocalVariableName(param, 'Unescaped');
      if (go.isPrimitiveType(param.type.elementType) && param.type.elementType.typeName === 'string') {
        // by convention, if the value is in its "final form" (i.e. no parsing required)
        // then its var is to have the "Param" suffix. the only case is string, everything
        // else requires some amount of parsing/conversion.
        paramVar = createLocalVariableName(param, 'Param');
      }
      content += `\t${paramVar} := make([]string, len(${escapedParam}))\n`;
      content += `\tfor i, v := range ${escapedParam} {\n`;
      content += '\t\tu, unescapeErr := url.QueryUnescape(v)\n';
      content += '\t\tif unescapeErr != nil {\n\t\t\treturn nil, unescapeErr\n\t\t}\n';
      content += `\t\t${paramVar}[i] = u\n\t}\n`;
      paramValue = paramVar;
    } else if (go.isPathParameter(param) || go.isQueryParameter(param)) {
      imports.add('net/url');
      let where: string;
      if (go.isPathParameter(param)) {
        where = 'Path';
      } else {
        where = 'Query';
      }
      let paramVar = createLocalVariableName(param, 'Unescaped');
      if (helpers.isRequiredParameter(param) && go.isConstantType(param.type) && param.type.type === 'string') {
        // for string-based enums, we perform the conversion as part of unescaping
        requiredHelpers.parseWithCast = true;
        paramVar = createLocalVariableName(param, 'Param');
        content += `\t${paramVar}, err := parseWithCast(${paramValue}, func (v string) (${go.getTypeDeclaration(param.type, clientPkg)}, error) {\n`;
        content += `\t\tp, unescapeErr := url.${where}Unescape(v)\n`;
        content += '\t\tif unescapeErr != nil {\n\t\t\treturn "", unescapeErr\n\t\t}\n';
        content += `\t\treturn ${go.getTypeDeclaration(param.type, clientPkg)}(p), nil\n\t})\n`;
      } else {
        if (helpers.isRequiredParameter(param) &&
        ((go.isPrimitiveType(param.type) && param.type.typeName === 'string') ||
          (go.isSliceType(param.type) && go.isPrimitiveType(param.type.elementType) && param.type.elementType.typeName === 'string'))) {
          // by convention, if the value is in its "final form" (i.e. no parsing required)
          // then its var is to have the "Param" suffix. the only case is string, everything
          // else requires some amount of parsing/conversion.
          paramVar = createLocalVariableName(param, 'Param');
        }
        content += `\t${paramVar}, err := url.${where}Unescape(${paramValue})\n`;
      }
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      paramValue = paramVar;
    }

    // parse params as required
    if (go.isHeaderCollectionParameter(param) || go.isPathCollectionParameter(param) || go.isQueryCollectionParameter(param)) {
      // any element type other than string will require some form of conversion/parsing
      if (!(go.isPrimitiveType(param.type.elementType) && param.type.elementType.typeName === 'string')) {
        if (param.collectionFormat !== 'multi') {
          requiredHelpers.splitHelper = true;
          const elementsParam = createLocalVariableName(param, 'Elements');
          content += `\t${elementsParam} := splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")\n`;
          paramValue = elementsParam;
        }

        const paramVar = createLocalVariableName(param, 'Param');
        let elementFormat: go.PrimitiveTypeName | go.DateTimeFormat | go.BytesEncoding;
        if (go.isConstantType(param.type.elementType)) {
          elementFormat = param.type.elementType.type;
        } else if (go.isPrimitiveType(param.type.elementType)) {
          elementFormat = param.type.elementType.typeName;
        } else if (go.isBytesType(param.type.elementType)) {
          elementFormat = param.type.elementType.encoding;
        } else if (go.isTimeType(param.type.elementType)) {
          elementFormat = param.type.elementType.dateTimeFormat;
        } else {
          throw new Error(`unhandled element type ${go.getTypeDeclaration(param.type.elementType)}`);
        }

        let toType = go.getTypeDeclaration(param.type.elementType);
        if (go.isConstantType(param.type.elementType)) {
          toType = `${clientPkg}.${toType}`;
        }
        content += `\t${paramVar} := make([]${toType}, len(${paramValue}))\n`;
        content += `\tfor i := 0; i < len(${paramValue}); i++ {\n`;
        let fromVar: string;

        // TODO: consolidate with non-collection parsing code
        if (elementFormat === 'bool') {
          imports.add('strconv');
          fromVar = 'parsedBool';
          content += `\t\t${fromVar}, parseErr := strconv.ParseBool(${paramValue}[i])\n`;
          content += '\t\tif parseErr != nil {\n\t\t\treturn nil, parseErr\n\t\t}\n';
        } else if (elementFormat === 'float32' || elementFormat === 'float64' || elementFormat === 'int32' || elementFormat === 'int64') {
          fromVar = `parsed${capitalize(elementFormat)}`;
          content += `\t\t${fromVar}, parseErr := ${emitNumericConversion(`${paramValue}[i]`, elementFormat)}\n`;
          content += '\t\tif parseErr != nil {\n\t\t\treturn nil, parseErr\n\t\t}\n';
        } else if (elementFormat === 'string') {
          // we're casting an enum string value to its const type
          // TODO: what about enums that aren't strings?
          fromVar = `${paramValue}[i]`;
        } else if (elementFormat === 'Std' || elementFormat === 'URL') {
          imports.add('encoding/base64');
          fromVar = `parsed${capitalize(elementFormat)}`;
          content += `\t\t${fromVar}, parseErr := base64.${elementFormat}Encoding.DecodeString(${paramValue}[i])\n`;
          content += '\t\tif parseErr != nil {\n\t\t\treturn nil, parseErr\n\t\t}\n';
        } else if (elementFormat === 'dateTimeRFC1123' || elementFormat === 'dateTimeRFC3339' || elementFormat === 'timeUnix') {
          imports.add('time');
          fromVar = `parsed${capitalize(elementFormat)}`;
          if (elementFormat === 'timeUnix') {
            imports.add('strconv');
            content += `\t\tp, parseErr := strconv.ParseInt(${paramValue}[i], 10, 64)\n`;
            content += '\t\tif parseErr != nil {\n\t\t\treturn nil, parseErr\n\t\t}\n';
            content += `\t\t${fromVar} := time.Unix(p, 0).UTC()\n`;
          } else {
            let format = 'time.RFC3339Nano';
            if (elementFormat === 'dateTimeRFC1123') {
              format = 'time.RFC1123';
            }
            content += `\t\t${fromVar}, parseErr := time.Parse(${format}, ${paramValue}[i])\n`;
            content += '\t\tif parseErr != nil {\n\t\t\treturn nil, parseErr\n\t\t}\n';
          }
        } else {
          throw new Error(`unhandled element format ${elementFormat}`);
        }
        // TODO: remove cast in some cases
        content += `\t\t${paramVar}[i] = ${toType}(${fromVar})\n\t}\n`;
      } else if (!helpers.isRequiredParameter(param) && param.collectionFormat !== 'multi') {
        // for slices of strings that are required, the call to splitHelper(...) is inlined into
        // the invocation of the fake e.g. srv.FakeFunc(splitHelper...). but if it's optional, we
        // need to create a local first which will later be copied into the optional param group.
        requiredHelpers.splitHelper = true;
        content += `\t${createLocalVariableName(param, 'Param')} := splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")\n`;
      }
    } else if (go.isPrimitiveType(param.type) && param.type.typeName === 'bool') {
      imports.add('strconv');
      let from = `strconv.ParseBool(${paramValue})`;
      if (!helpers.isRequiredParameter(param)) {
        requiredHelpers.parseOptional = true;
        from = `parseOptional(${paramValue}, strconv.ParseBool)`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (go.isBytesType(param.type)) {
      imports.add('encoding/base64');
      content += `\t${createLocalVariableName(param, 'Param')}, err := base64.${param.type.encoding}Encoding.DecodeString(${paramValue})\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (go.isTimeType(param.type)) {
      if (param.type.dateTimeFormat === 'dateType' || param.type.dateTimeFormat === 'timeRFC3339') {
        imports.add('time');
        let format = helpers.dateFormat;
        if (param.type.dateTimeFormat === 'timeRFC3339') {
          format = helpers.timeRFC3339Format;
        }
        let from = `time.Parse("${format}", ${paramValue})`;
        if (!helpers.isRequiredParameter(param)) {
          requiredHelpers.parseOptional = true;
          from = `parseOptional(${paramValue}, func(v string) (time.Time, error) { return time.Parse("${format}", v) })`;
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else if (param.type.dateTimeFormat === 'dateTimeRFC1123' || param.type.dateTimeFormat === 'dateTimeRFC3339') {
        imports.add('time');
        let format = 'time.RFC3339Nano';
        if (param.type.dateTimeFormat === 'dateTimeRFC1123') {
          format = 'time.RFC1123';
        }
        let from = `time.Parse(${format}, ${paramValue})`;
        if (!helpers.isRequiredParameter(param)) {
          requiredHelpers.parseOptional = true;
          from = `parseOptional(${paramValue}, func(v string) (time.Time, error) { return time.Parse(${format}, v) })`;
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else {
        imports.add('strconv');
        let parser: string;
        if (!helpers.isRequiredParameter(param)) {
          requiredHelpers.parseOptional = true;
          parser = 'parseOptional';
        } else {
          requiredHelpers.parseWithCast = true;
          parser = 'parseWithCast';
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${parser}(${paramValue}, func (v string) (time.Time, error) {\n`;
        content += '\t\tp, parseErr := strconv.ParseInt(v, 10, 64)\n';
        content += '\t\tif parseErr != nil {\n\t\t\treturn time.Time{}, parseErr\n\t\t}\n';
        content += '\t\treturn time.Unix(p, 0).UTC(), nil\n\t})\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      }
    } else if (go.isPrimitiveType(param.type) && (param.type.typeName === 'float32' || param.type.typeName === 'float64' || param.type.typeName === 'int32' || param.type.typeName === 'int64')) {
      let parser: string;
      if (!helpers.isRequiredParameter(param)) {
        requiredHelpers.parseOptional = true;
        parser = 'parseOptional';
      } else {
        requiredHelpers.parseWithCast = true;
        parser = 'parseWithCast';
      }
      if ((param.type.typeName === 'float32' || param.type.typeName === 'int32') || !helpers.isRequiredParameter(param)) {
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${parser}(${paramValue}, func(v string) (${param.type.typeName}, error) {\n`;
        content += `\t\tp, parseErr := ${emitNumericConversion('v', param.type.typeName)}\n`;
        content += '\t\tif parseErr != nil {\n\t\t\treturn 0, parseErr\n\t\t}\n';
        let result = 'p';
        if (param.type.typeName === 'float32' || param.type.typeName === 'int32') {
          result = `${param.type.typeName}(${result})`;
        }
        content += `\t\treturn ${result}, nil\n\t})\n`;
      } else {
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${emitNumericConversion(paramValue, param.type.typeName)}\n`;
      }
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (go.isHeaderMapParameter(param)) {
      imports.add('strings');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      const localVar = createLocalVariableName(param, 'Param');
      content += `\tvar ${localVar} map[string]*string\n`;
      content += `\tfor hh := range ${paramValue} {\n`;
      const headerPrefix = param.collectionPrefix;
      requiredHelpers.getHeaderValue = true;
      content += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("x-ms-meta-")], "${headerPrefix}") {\n`;
      content += `\t\t\tif ${localVar} == nil {\n\t\t\t\t${localVar} = map[string]*string{}\n\t\t\t}\n`;
      content += `\t\t\t${localVar}[hh[len("${headerPrefix}"):]] = to.Ptr(getHeaderValue(req.Header, hh))\n`;
      content += '\t\t}\n\t}\n';
    } else if (go.isConstantType(param.type) && param.type.type !== 'string') {
      let parseHelper: string;
      if (!helpers.isRequiredParameter(param)) {
        requiredHelpers.parseOptional = true;
        parseHelper = 'parseOptional';
      } else {
        requiredHelpers.parseWithCast = true;
        parseHelper = 'parseWithCast';
      }
      let parse: string;
      let zeroValue: string;
      if (param.type.type === 'bool') {
        imports.add('strconv');
        parse = 'strconv.ParseBool(v)';
        zeroValue = 'false';
      } else {
        // emitNumericConversion adds the necessary import of strconv
        parse = emitNumericConversion('v', param.type.type);
        zeroValue = '0';
      }
      const toConstType = go.getTypeDeclaration(param.type, clientPkg);
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${parseHelper}(${paramValue}, func(v string) (${toConstType}, error) {\n`;
      content += `\t\tp, parseErr := ${parse}\n`;
      content += `\t\tif parseErr != nil {\n\t\t\treturn ${zeroValue}, parseErr\n\t\t}\n`;
      content += `\t\treturn ${toConstType}(p), nil\n\t})\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (!helpers.isRequiredParameter(param)) {
      // we check this last as it's a superset of the previous conditions
      requiredHelpers.getOptional = true;
      if (go.isConstantType(param.type)) {
        paramValue = `${go.getTypeDeclaration(param.type, clientPkg)}(${paramValue})`;
      }
      content += `\t${createLocalVariableName(param, 'Param')} := getOptional(${paramValue})\n`;
    }
  }

  // create the param groups and populate their values
  for (const paramGroup of values(paramGroups.keys())) {
    if (paramGroup.required) {
      content += `\t${uncapitalize(paramGroup.name)} := ${clientPkg}.${paramGroup.groupName}{\n`;
      for (const param of values(paramGroups.get(paramGroup))) {
        content += `\t\t${capitalize(param.name)}: ${getFinalParamValue(clientPkg, param, paramValues)},\n`;
      }
      content += '\t}\n';
    } else {
      content += `\tvar ${uncapitalize(paramGroup.name)} *${clientPkg}.${paramGroup.groupName}\n`;
      const params = paramGroups.get(paramGroup);
      const paramNilCheck = new Array<string>();
      for (const param of values(params)) {
        // check array before body in case the body is just an array
        if (go.isSliceType(param.type)) {
          paramNilCheck.push(`len(${getFinalParamValue(clientPkg, param, paramValues)}) > 0`);
        } else if (go.isBodyParameter(param)) {
          if (param.bodyFormat === 'binary') {
            imports.add('io');
            paramNilCheck.push('req.Body != nil');
          } else {
            imports.add('reflect');
            paramNilCheck.push('!reflect.ValueOf(body).IsZero()');
          }
        } else if (go.isFormBodyParameter(param) || go.isMultipartFormBodyParameter(param)) {
          imports.add('reflect');
          paramNilCheck.push(`!reflect.ValueOf(${param.name}).IsZero()`);
        } else {
          paramNilCheck.push(`${getFinalParamValue(clientPkg, param, paramValues)} != nil`);
        }
      }
      content += `\tif ${paramNilCheck.join(' || ')} {\n`;
      content += `\t\t${uncapitalize(paramGroup.name)} = &${clientPkg}.${paramGroup.groupName}{\n`;
      for (const param of values(params)) {
        let byRef = '&';
        if (param.byValue || (!helpers.isRequiredParameter(param) && !go.isBodyParameter(param) && !go.isFormBodyParameter(param) && !go.isMultipartFormBodyParameter(param))) {
          byRef = '';
        }
        content += `\t\t\t${capitalize(param.name)}: ${byRef}${getFinalParamValue(clientPkg, param, paramValues)},\n`;
      }
      content += '\t\t}\n';
      content += '\t}\n';
    }
  }

  return {
    content: content,
    params: paramValues
  };
}

// works in conjunction with parseHeaderPathQueryParams
function populateApiParams(clientPkg: string, method: go.Method, paramValues: Map<string, string>, imports: ImportManager): string {
  // FooOperation(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], qp.Get("api-version"), nil)
  // this assumes that our caller has created matches and qp as required
  const params = new Array<string>();

  // for non-paged APIs, first param is always the context. use the one
  // from the HTTP request. be careful to properly handle paged LROs
  if (go.isLROMethod(method) || !go.isPageableMethod(method)) {
    params.push('req.Context()');
  }

  // now create the API call sig
  for (const param of values(helpers.getMethodParameters(method, consolidateHostParams))) {
    if (helpers.isParameterGroup(param)) {
      if (param.groupName === method.optionalParamsGroup.groupName) {
        // this is the optional params type. in some cases we just pass nil
        const countParams = values(param.params).where((each: go.Parameter) => { return !go.isResumeTokenParameter(each); }).count();
        if (countParams === 0) {
          // if the options param is empty or only contains the resume token param just pass nil
          params.push('nil');
          continue;
        }
      }
      // by convention, for param groups, the param parsing code
      // creates a local var with the name of the param
      params.push(uncapitalize(param.name));
      continue;
    }
    imports.addImportForType(param.type);
    params.push(getFinalParamValue(clientPkg, param, paramValues));
  }

  return params.join(', ');
}

// getRawParamValue returns the "raw" value for the specified parameter.
// depending on the type, the value might require parsing before it can be passed to the fake.
function getRawParamValue(param: go.Parameter): string {
  if (go.isFormBodyParameter(param) || go.isMultipartFormBodyParameter(param)) {
    // multipart form data values have been read and assigned
    // to local params with the same name. must check this first
    // as it's a superset of other cases that follow.
    return param.name;
  } else if (go.isPathParameter(param)) {
    // path params are in the matches slice
    return `matches[regex.SubexpIndex("${sanitizeRegexpCaptureGroupName(param.pathSegment)}")]`;
  } else if (go.isQueryParameter(param)) {
    // use qp
    if (go.isQueryCollectionParameter(param) && param.collectionFormat === 'multi') {
      return `qp["${param.queryParameter}"]`;
    }
    return `qp.Get("${param.queryParameter}")`;
  } else if (go.isHeaderParameter(param)) {
    if (go.isHeaderMapParameter(param) ) {
      return 'req.Header';
    }
    // use req
    requiredHelpers.getHeaderValue = true;
    return `getHeaderValue(req.Header, "${param.headerName}")`;
  } else if (go.isBodyParameter(param)) {
    if (param.bodyFormat === 'binary') {
      return 'req.Body.(io.ReadSeekCloser)';
    }
    // JSON/XML/text bodies have been deserialized into a local named body
    return 'body';
  } else if (go.isURIParameter(param)) {
    return 'req.URL.Host';
  } else {
    throw new Error(`unhandled parameter ${param.name}`);
  }
}

// getFinalParamValue returns the "final" value of param to be passed to the fake.
function getFinalParamValue(clientPkg: string, param: go.Parameter, paramValues: Map<string, string>): string {
  let paramValue = paramValues.get(param.name);
  if (!paramValue) {
    // the param didn't require parsing so the "raw" value can be used
    paramValue = getRawParamValue(param);
  }

  // there are a few corner-cases that require some fix-ups

  if ((go.isBodyParameter(param) || go.isFormBodyParameter(param) || go.isFormBodyCollectionParameter(param) || go.isMultipartFormBodyParameter(param)) && go.isTimeType(param.type)) {
    // time types in the body have been unmarshalled into our time helpers thus require a cast to time.Time
    return `time.Time(${paramValue})`;
  } else if (helpers.isRequiredParameter(param)) {
    // optional params are always in their "final" form
    if (go.isHeaderCollectionParameter(param) || go.isPathCollectionParameter(param) || go.isQueryCollectionParameter(param)) {
      // for required params that are collections of strings, we split them inline.
      // not necessary for optional params as they're already in slice format.
      if (param.collectionFormat !== 'multi' && go.isPrimitiveType(param.type.elementType) && param.type.elementType.typeName === 'string') {
        requiredHelpers.splitHelper = true;
        return `splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")`;
      }
    } else if (go.isHeaderParameter(param) && go.isConstantType(param.type) && param.type.type === 'string') {
      // since headers aren't escaped, we cast required, string-based enums inline
      return `${go.getTypeDeclaration(param.type, clientPkg)}(${paramValue})`;
    }
  }

  return paramValue;
}

// takes multiple host parameters and consolidates them into a single "host" parameter.
// this is necessary as there's no way to rehydrate multiple host parameters.
// e.g. host := "{vault}{secret}{dnsSuffix}" becomes http://contososecret.com
// there's no way to reliably split the host back up into its constituent parameters.
// so we just pass the full value as a single host parameter.
function consolidateHostParams(params: Array<go.Parameter>): Array<go.Parameter> {
  if (!values(params).where((each: go.Parameter) => { return go.isURIParameter(each); }).any()) {
    // no host params
    return params;
  }

  // consolidate multiple host params into a single "host" param
  const consolidatedParams = new Array<go.Parameter>();
  let hostParamAdded = false;
  for (const param of values(params)) {
    if (!go.isURIParameter(param)) {
      consolidatedParams.push(param);
    } else if (!hostParamAdded) {
      consolidatedParams.push(param);
      hostParamAdded = true;
    }
  }

  return consolidatedParams;
}

// copied from generator/operations.ts but with a slight tweak to consolidate host parameters
function getAPIParametersSig(method: go.Method, imports: ImportManager, pkgName?: string): string {
  const methodParams = helpers.getMethodParameters(method, consolidateHostParams);
  const params = new Array<string>();
  if (!go.isPageableMethod(method) || go.isLROMethod(method)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    let paramName = uncapitalize(methodParam.name);
    if (helpers.isParameter(methodParam) && go.isURIParameter(methodParam)) {
      paramName = 'host';
    }
    params.push(`${paramName} ${helpers.formatParameterTypeName(methodParam, pkgName)}`);
  }
  return params.join(', ');
}

// copied from generator/helpers.ts but without the XML-specific stuff
function getResultFieldName(result: go.AnyResult | go.BinaryResult | go.MonomorphicResult | go.PolymorphicResult | go.ModelResult): string {
  if (go.isAnyResult(result)) {
    return result.fieldName;
  } else if (go.isModelResult(result)) {
    return result.modelType.name;
  } else if (go.isPolymorphicResult(result)) {
    return result.interfaceType.name;
  }
  return result.fieldName;
}
