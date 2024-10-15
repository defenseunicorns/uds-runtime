// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

export interface DeployedPackage {
  name: string
  data: ZarfPackage
  cliVersion: string
  generation: number
  deployedComponents: DeployedComponent[]
  componentWebhooks?: { [key: string]: { [key: string]: Webhook } }
  connectStrings?: { [key: string]: ConnectString }
}

export interface ConnectString {
  description: string
  url: string
}

export interface DeployedComponent {
  name: string
  installedCharts: InstalledChart[]
  status: string
  observedGeneration: number
}

export interface Webhook {
  name: string
  waitDurationSeconds?: number
  status: string
  observedGeneration: number
}

export interface InstalledChart {
  namespace: string
  chartName: string
}

export interface ZarfPackage {
  kind: 'ZarfPackageConfig' | 'ZarfInitConfig'
  metadata?: ZarfMetadata
  build?: ZarfBuildData
  components: ZarfComponent[]
  constants?: Constant[]
  variables?: InteractiveVariable[]
}

export interface ZarfMetadata {
  name: string
  description?: string
  version?: string
  url?: string
  image?: string
  uncompressed?: boolean
  architecture?: string
  yolo?: boolean
  authors?: string
  documentation?: string
  source?: string
  vendor?: string
  aggregateChecksum?: string
}

export interface ZarfBuildData {
  terminal: string
  user: string
  architecture: string
  timestamp: string
  version: string
  migrations?: string[]
  registryOverrides?: { [key: string]: string }
  differential?: boolean
  differentialPackageVersion?: string
  differentialMissing?: string[]
  lastNonBreakingVersion?: string
  flavor?: string
}

export interface ZarfComponent {
  name: string
  description?: string
  default?: boolean
  required?: boolean
  only?: ZarfComponentOnlyTarget
  group?: string
  cosignKeyPath?: string
  import?: ZarfComponentImport
  manifests?: ZarfManifest[]
  charts?: ZarfChart[]
  dataInjections?: ZarfDataInjection[]
  files?: ZarfFile[]
  images?: string[]
  repos?: string[]
  scripts?: DeprecatedZarfComponentScripts
  actions?: ZarfComponentActions
}

export interface ZarfComponentOnlyTarget {
  localOS?: string
  cluster?: ZarfComponentOnlyCluster
  flavor?: string
}

export interface ZarfComponentOnlyCluster {
  architecture?: string
  distros?: string[]
}

export interface ZarfFile {
  source: string
  shasum?: string
  target: string
  executable?: boolean
  symlinks?: string[]
  extractPath?: string
}

export interface ZarfChart {
  name: string
  version?: string
  url?: string
  repoName?: string
  gitPath?: string
  localPath?: string
  namespace?: string
  releaseName?: string
  noWait?: boolean
  valuesFiles?: string[]
  variables?: ZarfChartVariable[]
}

export interface ZarfChartVariable {
  name: string
  description: string
  path: string
}

export interface ZarfManifest {
  name: string
  namespace?: string
  files?: string[]
  kustomizeAllowAnyDirectory?: boolean
  kustomizations?: string[]
  noWait?: boolean
}

export interface DeprecatedZarfComponentScripts {
  showOutput?: boolean
  timeoutSeconds?: number
  retry?: boolean
  prepare?: string[]
  before?: string[]
  after?: string[]
}

export interface ZarfComponentActions {
  onCreate?: ZarfComponentActionSet
  onDeploy?: ZarfComponentActionSet
  onRemove?: ZarfComponentActionSet
}

export interface ZarfComponentActionSet {
  defaults?: ZarfComponentActionDefaults
  before?: ZarfComponentAction[]
  after?: ZarfComponentAction[]
  onSuccess?: ZarfComponentAction[]
  onFailure?: ZarfComponentAction[]
}

export interface ZarfComponentActionDefaults {
  mute?: boolean
  maxTotalSeconds?: number
  maxRetries?: number
  dir?: string
  env?: string[]
  shell?: string
}

export interface ZarfComponentAction {
  mute?: boolean
  maxTotalSeconds?: number
  maxRetries?: number
  dir?: string
  env?: string[]
  cmd?: string
  shell?: string
  setVariable?: string
  setVariables?: Variable[]
  description?: string
  wait?: ZarfComponentActionWait
}

export interface ZarfComponentActionWait {
  cluster?: ZarfComponentActionWaitCluster
  network?: ZarfComponentActionWaitNetwork
}

export interface ZarfComponentActionWaitCluster {
  kind: string
  name: string
  namespace?: string
  condition?: string
}

export interface ZarfComponentActionWaitNetwork {
  protocol: string
  address: string
  code?: number
}

export interface ZarfContainerTarget {
  namespace: string
  selector: string
  container: string
  path: string
}

export interface ZarfDataInjection {
  source: string
  target: ZarfContainerTarget
  compress?: boolean
}

export interface ZarfComponentImport {
  name?: string
  path?: string
  url?: string
}

export interface Variable {
  name: string
  sensitive?: boolean
  autoIndent?: boolean
  pattern?: string
  type?: string
}

export interface InteractiveVariable {
  description?: string
  default?: string
  prompt?: boolean
}

export interface Constant {
  name: string
  value: string
  description?: string
  autoIndent?: boolean
  pattern?: string
}

export interface SetVariable {
  value: string
}
