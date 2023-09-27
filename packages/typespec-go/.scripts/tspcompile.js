// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { execSync } from 'child_process';

import { existsSync, opendirSync, unlinkSync } from 'fs';

import { semaphore } from './semaphore.js';

// limit to 8 concurrent builds
const sem = semaphore(8);

const pkgRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/';

const tspRoot = pkgRoot + 'node_modules/@azure-tools/cadl-ranch-specs/http/';

// the format is as follows
// 'moduleName': [ 'moduleVersion', 'outputDir', 'additional args' ]
const tsps = {
  'arraygroup': ['0.1.1', 'type/array', '--remove-unreferenced-types'],
  'dictionarygroup': ['0.1.1', 'type/dictionary', '--remove-unreferenced-types'],
  'extensibleenumgroup': ['0.1.1', 'type/enum/extensible', '--remove-unreferenced-types'],
  //'singlediscriminatorgroup': ['0.1.1', 'type/model/inheritance/single-discriminator', '--remove-unreferenced-types'],
  'visibilitygroup': ['0.1.1', 'type/model/visibility', '--remove-unreferenced-types']
};

// any new args must also be added to autorest.go\common\config\rush\command-line.json
const args = process.argv.slice(2);
var filter = undefined;
const switches = [];
for (var i = 0 ; i < args.length; i += 1) {
  switch (args[i]) {
    case '--filter':
    case '-f':
      filter = args[i + 1]
      i += 1
      break;
    case '--verbose':
    case '-v':
      switches.push('--verbose');
      break;
    default:
      break;
  }
}

if (filter !== undefined) {
  console.log("Using filter: " + filter)
}

function should_generate(name) {
  if (filter !== undefined) {
    const re = new RegExp(filter);
    return re.test(name)
  }
  return true
}

for (const module in tsps) {
  const values = tsps[module];
  generate(module, values[0], tspRoot + values[1], 'test/' + values[1]);
}

function generate(moduleName, moduleVersion, inputDir, outputDir, additionalArgs) {
  if (!should_generate(moduleName)) {
    return
  }
  if (additionalArgs === undefined) {
    additionalArgs = '';
  }
  sem.take(function() {
    console.log('generating ' + inputDir);
    const fullOutputDir = pkgRoot + outputDir;
    try {
      const command = `tsp compile ${inputDir}/main.tsp --emit=${pkgRoot} --option="@azure-tools/typespec-go.module=${moduleName}" --option="@azure-tools/typespec-go.module-version=${moduleVersion}" --option="@azure-tools/typespec-go.emitter-output-dir=${fullOutputDir}" --option="@azure-tools/typespec-go.file-prefix=zz_"`;
      if (switches.includes('--verbose')) {
        console.log(command);
      }
      cleanGeneratedFiles(fullOutputDir);
      execSync(command);
      execSync('gofmt -w .', { cwd: fullOutputDir});
      execSync('go mod tidy', { cwd: fullOutputDir});
    } catch (err) {
      console.error(err.output.toString());
    }
  });
}

function cleanGeneratedFiles(outputDir) {
  if (!existsSync(outputDir)) {
      return;
  }
  const dir = opendirSync(outputDir);
  while (true) {
      const dirEnt = dir.readSync()
      if (dirEnt === null) {
          break;
      }
      if (dirEnt.isFile() && dirEnt.name.startsWith('zz_')) {
          unlinkSync(dir.path + '/' + dirEnt.name);
      }
  }
  dir.close();
  cleanGeneratedFiles(outputDir + '/fake');
}
