/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/gocodemodel.js';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatLiteralValue, getParentImport, recursiveUnwrapMapSlice } from './helpers.js';
import { ImportManager } from './imports.js';

// Creates the content in polymorphic_helpers.go
export async function generatePolymorphicHelpers(codeModel: go.CodeModel, fakeServerPkg?: string): Promise<string> {
  if (codeModel.interfaceTypes.length === 0) {
    // no polymorphic types
    return '';
  }

  let text = contentPreamble(codeModel, fakeServerPkg);
  const imports = new ImportManager();
  imports.add('encoding/json');
  if (fakeServerPkg) {
    // content is being generated into a separate package, add the necessary import
    imports.add(getParentImport(codeModel));
  }

  text += imports.text();

  const scalars = new Set<string>();
  const arrays = new Set<string>();
  const maps = new Set<string>();

  // we know there are polymorphic types but we don't know how they're used.
  // i.e. are they vanilla fields, elements in a slice, or values in a map.
  // polymorphic types within maps/slices will also need the scalar helpers.
  const trackDisciminator = function(type: go.PossibleType) {
    if (go.isInterfaceType(type)) {
      scalars.add(type.name);
    } else if (go.isSliceType(type)) {
      const leafType = recursiveUnwrapMapSlice(type);
      if (go.isInterfaceType(leafType)) {
        scalars.add(leafType.name);
        arrays.add(leafType.name);
      }
    } else if (go.isMapType(type)) {
      const leafType = recursiveUnwrapMapSlice(type);
      if (go.isInterfaceType(leafType)) {
        scalars.add(leafType.name);
        maps.add(leafType.name);
      }
    }
  };

  // calculate which discriminator helpers we actually need to generate
  
  if (fakeServerPkg) {
    // when generating for the fakes server, we must look at operation parameters instead of return values
    for (const client of values(codeModel.clients)) {
      for (const method of values(client.methods)) {
        for (const param of values(method.parameters)) {
          trackDisciminator(param.type);
        }
      }
    }
  } else {
    for (const model of codeModel.models) {
      for (const field of model.fields) {
        trackDisciminator(field.type);
      }
    }

    for (const respEnv of values(codeModel.responseEnvelopes)) {
      if (!respEnv.result) {
        continue;
      }

      if (go.isMonomorphicResult(respEnv.result)) {
        if (go.isMapType(respEnv.result.monomorphicType)) {
          trackDisciminator(respEnv.result.monomorphicType.valueType);
        } else if (go.isSliceType(respEnv.result.monomorphicType)) {
          trackDisciminator(respEnv.result.monomorphicType.elementType);
        }
      } else if (go.isPolymorphicResult(respEnv.result)) {
        trackDisciminator(respEnv.result.interfaceType);
      }
    }
  }

  if (scalars.size === 0 && arrays.size === 0 && maps.size === 0) {
    // this is a corner-case that can happen when all the discriminated types
    // are error types.  there's a bug in M4 that incorrectly annotates such
    // types as 'output', 'exception' in the usage however it's really just
    // 'exception'.  until this is fixed, we can wind up here.
    return '';
  }

  let prefix = '';
  if (fakeServerPkg) {
    // content is being generated into a separate package, set the type name prefix
    prefix = `${codeModel.packageName}.`;
  }

  for (const interfaceType of codeModel.interfaceTypes) {
    // generate unmarshallers for each discriminator

    // scalar unmarshaller
    if (scalars.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}(rawMsg json.RawMessage) (${prefix}${interfaceType.name}, error) {\n`;
      text += '\tif rawMsg == nil || string(rawMsg) == "null" {\n';
      text += '\t\treturn nil, nil\n';
      text += '\t}\n';
      text += '\tvar m map[string]any\n';
      text += '\tif err := json.Unmarshal(rawMsg, &m); err != nil {\n';
      text += '\t\treturn nil, err\n';
      text += '\t}\n';
      text += `\tvar b ${prefix}${interfaceType.name}\n`;
      text += `\tswitch m["${interfaceType.discriminatorField}"] {\n`;
      for (const possibleType of interfaceType.possibleTypes) {
        let disc = formatLiteralValue(possibleType.discriminatorValue!, true);
        // when the discriminator value is an enum, cast the const as a string
        if (go.isConstantType(possibleType.discriminatorValue!.type)) {
          disc = `string(${prefix}${disc})`;
        }
        text += `\tcase ${disc}:\n`;
        text += `\t\tb = &${prefix}${possibleType.name}{}\n`;
      }
      text += '\tdefault:\n';
      text += `\t\tb = &${prefix}${interfaceType.rootType.name}{}\n`;
      text += '\t}\n';
      text += '\tif err := json.Unmarshal(rawMsg, b); err != nil {\n\t\treturn nil, err\n\t}\n';
      text += '\treturn b, nil\n';
      text += '}\n\n';
    }

    // array unmarshaller
    if (arrays.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}Array(rawMsg json.RawMessage) ([]${prefix}${interfaceType.name}, error) {\n`;
      text += '\tif rawMsg == nil || string(rawMsg) == "null" {\n';
      text += '\t\treturn nil, nil\n';
      text += '\t}\n';
      text += '\tvar rawMessages []json.RawMessage\n';
      text += '\tif err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n';
      text += '\t\treturn nil, err\n';
      text += '\t}\n';
      text += `\tfArray := make([]${prefix}${interfaceType.name}, len(rawMessages))\n`;
      text += '\tfor index, rawMessage := range rawMessages {\n';
      text += `\t\tf, err := unmarshal${interfaceType.name}(rawMessage)\n`;
      text += '\t\tif err != nil {\n';
      text += '\t\t\treturn nil, err\n';
      text += '\t\t}\n';
      text += '\t\tfArray[index] = f\n';
      text += '\t}\n';
      text += '\treturn fArray, nil\n';
      text += '}\n\n';
    }

    // map unmarshaller
    if (maps.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}Map(rawMsg json.RawMessage) (map[string]${prefix}${interfaceType.name}, error) {\n`;
      text += '\tif rawMsg == nil || string(rawMsg) == "null" {\n';
      text += '\t\treturn nil, nil\n';
      text += '\t}\n';
      text += '\tvar rawMessages map[string]json.RawMessage\n';
      text += '\tif err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n';
      text += '\t\treturn nil, err\n';
      text += '\t}\n';
      text += `\tfMap := make(map[string]${prefix}${interfaceType.name}, len(rawMessages))\n`;
      text += '\tfor key, rawMessage := range rawMessages {\n';
      text += `\t\tf, err := unmarshal${interfaceType.name}(rawMessage)\n`;
      text += '\t\tif err != nil {\n';
      text += '\t\t\treturn nil, err\n';
      text += '\t\t}\n';
      text += '\t\tfMap[key] = f\n';
      text += '\t}\n';
      text += '\treturn fMap, nil\n';
      text += '}\n\n';
    }
  }
  return text;
}
