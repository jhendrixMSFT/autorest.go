/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize } from './namer.js';
import * as go from '../../../codemodel.go/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';

// cache of previously created types
const types = new Map<string, go.PossibleType>();
const constValues = new Map<string, go.ConstantValue>();

export function adaptConstantType(enumType: tcgc.SdkEnumType): go.ConstantType {
  let constType = types.get(enumType.name);
  if (constType) {
    return <go.ConstantType>constType;
  }
  constType = new go.ConstantType(enumType.name, adaptPrimitiveType(enumType.valueType.kind), `Possible${enumType.name}Values`);
  constType.values = adaptConstantValue(constType, enumType.values);
  constType.description = enumType.description;
  types.set(enumType.name, constType);
  return constType;
}

export function adaptModel(model: tcgc.SdkModelType): go.ModelType | go.PolymorphicType {
  let modelType = types.get(model.name);
  if (modelType) {
    return <go.ModelType | go.PolymorphicType>modelType;
  }

  // TODO: what's the extension equivalent in TS?
  const annotations = new go.ModelAnnotations(false);
  if (model.discriminatedSubtypes || model.discriminatorValue) {
    throw new Error('discriminators nyi');
  } else {
    // TODO: hard-coded format
    modelType = new go.ModelType(model.name, 'json', annotations);
    // polymorphic types don't have XMLInfo
    // TODO: XMLInfo
  }
  modelType.description = model.description;
  types.set(model.name, modelType);
  return modelType;
}

export function adaptModelField(prop: tcgc.SdkBodyModelPropertyType, obj: tcgc.SdkModelType): go.ModelField {
  // TODO: hard-coded values
  const annotations = new go.ModelFieldAnnotations(prop.optional == false, false, false, false);
  const field = new go.ModelField(capitalize(prop.nameInClient), adaptPossibleType(prop.type), isTypePassedByValue(prop.type), prop.serializedName, annotations);
  field.description = prop.description;
  if (prop.discriminator && obj.discriminatorValue) {
    const keyName = `discriminator-value-${obj.discriminatorValue}`;
    let discriminatorLiteral = <go.LiteralValue>types.get(keyName);
    if (!discriminatorLiteral) {
      // the discriminatorValue is either a quoted string or a constant (i.e. enum) value
      if (obj.discriminatorValue[0] === '"') {
        discriminatorLiteral = new go.LiteralValue(new go.PrimitiveType('string'), obj.discriminatorValue);
      } else {
        // find the corresponding constant value
        const value = constValues.get(obj.discriminatorValue);
        if (!value) {
          throw new Error(`didn't find a constant value for discriminator value ${obj.discriminatorValue}`);
        }
        discriminatorLiteral = new go.LiteralValue(value.type, value);
      }
    }
    types.set(keyName, discriminatorLiteral);
    field.defaultValue = discriminatorLiteral;
  } /*else if (prop.clientDefaultValue) {
    if (!go.isLiteralValueType(field.type)) {
      throw new Error(`unsupported default value type ${go.getTypeDeclaration(field.type)} for field ${field.fieldName}`);
    }
    if (go.isConstantType(field.type)) {
      // find the corresponding ConstantValue
      const constType = types.get(field.type.name);
      if (!constType) {
        throw new Error(`didn't find ConstantType for ${field.type.name}`);
      }
      let found = false;
      for (const val of values((<go.ConstantType>constType).values)) {
        if (val.value === prop.clientDefaultValue) {
          const keyName = `literal-${val.valueName}`;
          let literalValue = types.get(keyName);
          if (!literalValue) {
            literalValue = new go.LiteralValue(field.type, val);
            types.set(keyName, literalValue);
          }
          field.defaultValue = <go.LiteralValue>literalValue;
          found = true;
          break;
        }
      }
      if (!found) {
        throw new Error(`didn't find ConstantValue for ${prop.clientDefaultValue}`);
      }
    } else {
      const keyName = `literal-${go.getTypeDeclaration(field.type)}-${prop.clientDefaultValue}`;
      let literalValue = types.get(keyName);
      if (!literalValue) {
        literalValue = new go.LiteralValue(field.type, prop.clientDefaultValue);
        types.set(keyName, literalValue);
      }
      field.defaultValue = <go.LiteralValue>literalValue;
    }
  }*/

  // TODO: XMLInfo
  //field.xml = adaptXMLInfo(prop.schema);

  return field;
}

// converts an tcgc type to a Go code model type
export function adaptPossibleType(type: tcgc.SdkType, elementTypeByValue?: boolean): go.PossibleType {
  const rawJSONAsBytes = true;// TODO: <boolean>schema.language.go!.rawJSONAsBytes;
  switch (type.kind) {
    case 'any': {
      if (rawJSONAsBytes) {
        const anyRawJSONKey = 'any-raw-json';
        let anyRawJSON = types.get(anyRawJSONKey);
        if (anyRawJSON) {
          return anyRawJSON;
        }
        anyRawJSON = new go.SliceType(new go.PrimitiveType('byte'), true);
        anyRawJSON.rawJSONAsBytes = true;
        types.set(anyRawJSONKey, anyRawJSON);
        return anyRawJSON;
      }
      let anyType = types.get('any');
      if (anyType) {
        return anyType;
      }
      anyType = new go.PrimitiveType('any');
      types.set('any', anyType);
      return anyType;
    }
    /*case m4.SchemaType.AnyObject: {
      if (rawJSONAsBytes) {
        const anyObjectRawJSONKey = 'any-object-raw-json';
        let anyObjectRawJSON = types.get(anyObjectRawJSONKey);
        if (anyObjectRawJSON) {
          return anyObjectRawJSON;
        }
        anyObjectRawJSON = new go.SliceType(new go.PrimitiveType('byte'), true);
        anyObjectRawJSON.rawJSONAsBytes = true;
        types.set(anyObjectRawJSONKey, anyObjectRawJSON);
        return anyObjectRawJSON;
      }
      const anyObjKey = 'any-object';
      let anyObject = types.get(anyObjKey);
      if (anyObject) {
        return anyObject;
      }
      anyObject = new go.MapType(new go.PrimitiveType('any'), true);
      types.set(anyObjKey, anyObject);
      return anyObject;
    }*/
    case 'array': {
      const myElementTypeByValue = isTypePassedByValue(type.valueType);
      const keyName = recursiveKeyName(`array-${myElementTypeByValue}`, type.valueType);
      let arrayType = types.get(keyName);
      if (arrayType) {
        return arrayType;
      }
      arrayType = new go.SliceType(adaptPossibleType(type.valueType, elementTypeByValue), myElementTypeByValue);
      types.set(keyName, arrayType);
      return arrayType;
    }
    /*case m4.SchemaType.Binary: {
      let binaryType = types.get(m4.SchemaType.Binary);
      if (binaryType) {
        return binaryType;
      }
      binaryType = new go.StandardType('io.ReadSeekCloser', 'io');
      types.set(m4.SchemaType.Binary, binaryType);
      return binaryType;
    }*/
    case 'boolean': {
      const boolKey = 'boolean';
      let primitiveBool = types.get(boolKey);
      if (primitiveBool) {
        return primitiveBool;
      }
      primitiveBool = new go.PrimitiveType('bool');
      types.set(boolKey, primitiveBool);
      return primitiveBool;
    }
    case 'bytes':
      return adaptBytesType(type);
    /*case m4.SchemaType.Char: {
      let rune = types.get(m4.SchemaType.Char);
      if (rune) {
        return rune;
      }
      rune = new go.PrimitiveType('rune');
      types.set(m4.SchemaType.Char, rune);
      return rune;
    }*/
    case 'enum':
      return adaptConstantType(type);
    case 'constant':
      return adaptLiteralValue(type);
    /*case m4.SchemaType.Credential: {
      let credType = types.get(m4.SchemaType.Credential);
      if (credType) {
        return credType;
      }
      credType = new go.PrimitiveType('string');
      types.set(m4.SchemaType.Credential, credType);
      return credType;
    }*/
    case 'date': {
      if (type.encode !== 'rfc3339') {
        throw new Error(`unsupported date encoding ${type.encode}`);
      }
      const dateKey = `date-${type.encode}`;
      let date = types.get(dateKey);
      if (date) {
        return date;
      }
      date = new go.TimeType('dateType');
      types.set(dateKey, date);
      return date;
    }
    case 'datetime': {
      const encoding = adaptDateTimeEncoding(type.encode);
      let datetime = types.get(encoding);
      if (datetime) {
        return datetime;
      }
      datetime = new go.TimeType(encoding);
      types.set(encoding, datetime);
      return datetime;
    }
    case 'time': {
      if (type.encode !== 'rfc3339') {
        throw new Error(`unsupported time encoding ${type.encode}`);
      }
      const encoding = 'timeRFC3339';
      let time = types.get(encoding);
      if (time) {
        return time;
      }
      time = new go.TimeType(encoding);
      types.set(encoding, time);
      return time;
    }
    case 'dict': {
      const valueTypeByValue = isTypePassedByValue(type.valueType);
      const keyName = recursiveKeyName(`dict-${valueTypeByValue}`, type.valueType);
      let mapType = types.get(keyName);
      if (mapType) {
        return mapType;
      }
      mapType = new go.MapType(adaptPossibleType(type.valueType, elementTypeByValue), valueTypeByValue);
      types.set(keyName, mapType);
      return mapType;
    }
    case 'int32': {
      const int32Key = 'int32';
      let int32 = types.get(int32Key);
      if (int32) {
        return int32;
      }
      int32 = new go.PrimitiveType(int32Key);
      types.set(int32Key, int32);
      return int32;
    }
    case 'int64': {
      const int64Key = 'int64';
      let int64 = types.get(int64Key);
      if (int64) {
        return int64;
      }
      int64 = new go.PrimitiveType(int64Key);
      types.set(int64Key, int64);
      return int64;
    }
    case 'float32': {
      const float32Key = 'float32';
      let float32 = types.get(float32Key);
      if (float32) {
        return float32;
      }
      float32 = new go.PrimitiveType(float32Key);
      types.set(float32Key, float32);
      return float32;
    }
    case 'float64': {
      const float64Key = 'float64';
      let float64 = types.get(float64Key);
      if (float64) {
        return float64;
      }
      float64 = new go.PrimitiveType(float64Key);
      types.set(float64Key, float64);
      return float64;
    }
    case 'model':
      return adaptModel(type);
    /*case m4.SchemaType.ODataQuery: {
      let odataType = types.get(m4.SchemaType.ODataQuery);
      if (odataType) {
        return odataType;
      }
      odataType = new go.PrimitiveType('string');
      types.set(m4.SchemaType.ODataQuery, odataType);
      return odataType;
    }*/
    case 'armId':
    case 'duration':
    case 'string': {
      const stringKey = 'string';
      let stringType = types.get(stringKey);
      if (stringType) {
        return stringType;
      }
      stringType = new go.PrimitiveType('string');
      types.set(stringKey, stringType);
      return stringType;
    }
    case 'url': {
      const urlKey = 'url';
      let uriType = types.get(urlKey);
      if (uriType) {
        return uriType;
      }
      uriType = new go.PrimitiveType('string');
      types.set(urlKey, uriType);
      return uriType;
    }
    case 'uuid':
    case 'guid': {
      const uuidKey = 'uuid';
      let uuid = types.get(uuidKey);
      if (uuid) {
        return uuid;
      }
      uuid = new go.PrimitiveType('string');
      types.set(uuidKey, uuid);
      return uuid;
    }
    default:
      throw new Error(`unhandled property schema type ${type.kind}`);
  }
}

function adaptConstantValue(type: go.ConstantType, valueTypes: Array<tcgc.SdkEnumValueType>): Array<go.ConstantValue> {
  const values = new Array<go.ConstantValue>();
  for (const valueType of valueTypes) {
    let value = constValues.get(valueType.name);
    if (!value) {
      value = new go.ConstantValue(`${type.name}${valueType.name}`, type, valueType.value);
      value.description = valueType.description;
      constValues.set(valueType.name, value);
    }
    values.push(value);
  }
  return values;
}

function adaptPrimitiveType(kind: tcgc.SdkBuiltInKinds): 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string' {
  switch (kind) {
    case 'boolean':
      return 'bool';
    case 'float32':
    case 'float64':
    case 'int32':
    case 'int64':
    case 'string':
      return kind;
    default:
      throw new Error(`unhandled tcgc.SdkBuiltInKinds: ${kind}`);
  }
}

function adaptBytesType(schema: tcgc.SdkBuiltInType): go.BytesType {
  let format: go.BytesEncoding = 'Std';
  if (schema.encode === 'base64url') {
    format = 'URL';
  }
  const keyName = `bytes-${format}`;
  let bytesType = types.get(keyName);
  if (bytesType) {
    return <go.BytesType>bytesType;
  }
  bytesType = new go.BytesType(format);
  types.set(keyName, bytesType);
  return bytesType;
}

function adaptDateTimeEncoding(encoding: tsp.DateTimeKnownEncoding): go.DateTimeFormat {
  switch (encoding) {
    case 'rfc3339':
      return 'dateTimeRFC3339';
    case 'rfc7231':
      return 'dateTimeRFC1123';
    case 'unixTimestamp':
      return 'timeUnix';
  }
}

function adaptLiteralValue(constSchema: tcgc.SdkConstantType): go.LiteralValue {
  switch (constSchema.valueType.kind) {
    case 'boolean': {
      const keyName = `literal-boolean-${constSchema.value}`;
      let literalBool = types.get(keyName);
      if (literalBool) {
        return <go.LiteralValue>literalBool;
      }
      literalBool = new go.LiteralValue(new go.PrimitiveType('bool'), constSchema.value);
      types.set(keyName, literalBool);
      return literalBool;
    }
    /*case m4.SchemaType.ByteArray: {
      const keyName = `literal-${m4.SchemaType.ByteArray}-${constSchema.value.value}`;
      let literalByteArray = types.get(keyName);
      if (literalByteArray) {
        return <go.LiteralValue>literalByteArray;
      }
      literalByteArray = new go.LiteralValue(adaptBytesType(<m4.ByteArraySchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalByteArray);
      return literalByteArray;
    }
    case m4.SchemaType.Choice:
    case m4.SchemaType.SealedChoice: {
      const keyName = `literal-choice-${constSchema.value.value}`;
      let literalConst = types.get(keyName);
      if (literalConst) {
        return <go.LiteralValue>literalConst;
      }
      literalConst = new go.LiteralValue(adaptConstantType(<m4.ChoiceSchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalConst);
      return literalConst;
    }
    case m4.SchemaType.Date:
    case m4.SchemaType.DateTime:
    case m4.SchemaType.UnixTime: {
      const keyName = `literal-${constSchema.valueType.language.go!.internalTimeType}-${constSchema.value.value}`;
      let literalTime = types.get(keyName);
      if (literalTime) {
        return <go.LiteralValue>literalTime;
      }
      literalTime = new go.LiteralValue(new go.TimeType(constSchema.valueType.language.go!.internalTimeType), constSchema.value.value);
      types.set(keyName, literalTime);
      return literalTime;
    }*/
    case 'int32':
    case 'int64': {
      const keyName = `literal-${constSchema.valueType.kind}-${constSchema.value}`;
      let literalInt = types.get(keyName);
      if (literalInt) {
        return <go.LiteralValue>literalInt;
      }
      literalInt = new go.LiteralValue(new go.PrimitiveType(constSchema.valueType.kind), constSchema.value);
      types.set(keyName, literalInt);
      return literalInt;
    }
    case 'float32':
    case 'float64': {
      const keyName = `literal-${constSchema.valueType.kind}-${constSchema.value}`;
      let literalFloat = types.get(keyName);
      if (literalFloat) {
        return <go.LiteralValue>literalFloat;
      }
      literalFloat = new go.LiteralValue(new go.PrimitiveType(constSchema.valueType.kind), constSchema.value);
      types.set(keyName, literalFloat);
      return literalFloat;
    }
    case 'string':
    case 'guid':
    case 'uuid': {
      const keyName = `literal-string-${constSchema.value}`;
      let literalString = types.get(keyName);
      if (literalString) {
        return <go.LiteralValue>literalString;
      }
      literalString = new go.LiteralValue(new go.PrimitiveType('string'), constSchema.value);
      types.set(keyName, literalString);
      return literalString;
    }
    default:
      throw new Error(`unsupported scheam type ${constSchema.valueType.kind} for LiteralValue`);
  }

  // TODO: case m4.SchemaType.Duration:
}

function recursiveKeyName(root: string, obj: tcgc.SdkType): string {
  switch (obj.kind) {
    case 'array':
      return recursiveKeyName(`${root}-array`, obj.valueType);
    case 'dict':
      return recursiveKeyName(`${root}-dict`, obj.valueType);
    case 'date':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported date encoding ${obj.encode}`);
      }
      return `${root}-dateRFC3339`;
    case 'datetime':
      return `${root}-${adaptDateTimeEncoding(obj.encode)}`;
    case 'time':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported time encoding ${obj.encode}`);
      }
      return `${root}-timeRFC3339`;
    default:
      return `${root}-${obj.kind}`;
  }
}

export function isTypePassedByValue(type: tcgc.SdkType): boolean {
  return type.kind === 'any' || type.kind === 'array' ||
  type.kind === 'bytes' || type.kind === 'dict' ||
    (type.kind === 'model' && !!type.discriminatedSubtypes);
}
