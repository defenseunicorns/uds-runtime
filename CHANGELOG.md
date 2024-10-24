# Changelog

## [0.7.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.6.1...v0.7.0) (2024-10-17)


### Features

* adds security context to Helm chart ([#468](https://github.com/defenseunicorns/uds-runtime/issues/468)) ([665fde1](https://github.com/defenseunicorns/uds-runtime/commit/665fde1472427bce5f985e82ec1ff96e261f04ae))
* **ui:** view all cluster CRDs ([#423](https://github.com/defenseunicorns/uds-runtime/issues/423)) ([010a320](https://github.com/defenseunicorns/uds-runtime/commit/010a32024d252fdf419654823fe5961e624948ef))


### Miscellaneous

* **ci:** e2e test for cluster disconnection and reconnection ([#447](https://github.com/defenseunicorns/uds-runtime/issues/447)) ([b4847d4](https://github.com/defenseunicorns/uds-runtime/commit/b4847d4daab19eecf33098b5445af27b8501a21a))
* more auth refactors, move Go code into src, remove build tags ([#461](https://github.com/defenseunicorns/uds-runtime/issues/461)) ([f14dcce](https://github.com/defenseunicorns/uds-runtime/commit/f14dcce11ea1ba2639dcd482bc1c5bc2030da839))
* **ui:** re-enable svelte/no-at-html-tags rule ([#464](https://github.com/defenseunicorns/uds-runtime/issues/464)) ([0137419](https://github.com/defenseunicorns/uds-runtime/commit/013741996aa525dd0490684cfd35ce62d4d683c7))
* watch error handler test ([#462](https://github.com/defenseunicorns/uds-runtime/issues/462)) ([a4fc5a2](https://github.com/defenseunicorns/uds-runtime/commit/a4fc5a239dfc46aa0f014075eb670986800bdeaa))

## [0.6.1](https://github.com/defenseunicorns/uds-runtime/compare/v0.6.0...v0.6.1) (2024-10-15)


### Miscellaneous

* bump uds-types in package lock ([#455](https://github.com/defenseunicorns/uds-runtime/issues/455)) ([c17a583](https://github.com/defenseunicorns/uds-runtime/commit/c17a58343c94c3c4ea08717ffbaa2f5c10917a66))
* bumps uds-types version ([#454](https://github.com/defenseunicorns/uds-runtime/issues/454)) ([4c767d5](https://github.com/defenseunicorns/uds-runtime/commit/4c767d5e55a0f698b33ee50335cf11cfa5b5d155))
* **deps:** update dependency kubernetes-fluent-client to v3.1.1 ([#448](https://github.com/defenseunicorns/uds-runtime/issues/448)) ([abe8583](https://github.com/defenseunicorns/uds-runtime/commit/abe85835dce4e57681cdff3f3a4ec9ec1c787b04))
* **deps:** update devdependencies ([#446](https://github.com/defenseunicorns/uds-runtime/issues/446)) ([a0f6163](https://github.com/defenseunicorns/uds-runtime/commit/a0f616335d824069ad511c79ebe7d61bb30c6739))
* license to AGPLv3 and update codeowners ([#444](https://github.com/defenseunicorns/uds-runtime/issues/444)) ([2f0682a](https://github.com/defenseunicorns/uds-runtime/commit/2f0682a57a8b889515096e8c765582b07ae41c12))
* refactor auth logic ([#426](https://github.com/defenseunicorns/uds-runtime/issues/426)) ([caa7e8f](https://github.com/defenseunicorns/uds-runtime/commit/caa7e8f77dce756a5fae4cfb81eb26d85dadd41d))
* update license headers ([#450](https://github.com/defenseunicorns/uds-runtime/issues/450)) ([5e3a614](https://github.com/defenseunicorns/uds-runtime/commit/5e3a6148df1acb556acd39914c36ef06a6f6fb72))
* updates codeowners for both license files ([#453](https://github.com/defenseunicorns/uds-runtime/issues/453)) ([6c4076a](https://github.com/defenseunicorns/uds-runtime/commit/6c4076a086fe4d9b8e67749a60151f17803e5448))

## [0.6.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.5.0...v0.6.0) (2024-10-10)


### Features

* **api:** adds caching layer to pepr endpoints ([#402](https://github.com/defenseunicorns/uds-runtime/issues/402)) ([782d93b](https://github.com/defenseunicorns/uds-runtime/commit/782d93b9b030a0a16191de58ec7c6920982f248a))
* **api:** use tls when running locally ([#405](https://github.com/defenseunicorns/uds-runtime/issues/405)) ([d4764eb](https://github.com/defenseunicorns/uds-runtime/commit/d4764ebcff0ad5c3b4d3dfaee563429ce07d4e87))
* **ci:** publish the runtime ui build as a release artifact ([#418](https://github.com/defenseunicorns/uds-runtime/issues/418)) ([ce4a592](https://github.com/defenseunicorns/uds-runtime/commit/ce4a592c615b02e149f2a8b7c0698fb85157b3b6))
* **ui:** 409 overview dashboard events logs widget ([#415](https://github.com/defenseunicorns/uds-runtime/issues/415)) ([ea47a72](https://github.com/defenseunicorns/uds-runtime/commit/ea47a72b47d60c4b51a9b6ef63cd5853f57d6fd6))
* **ui:** create stat widget ([#386](https://github.com/defenseunicorns/uds-runtime/issues/386)) ([f1dce2e](https://github.com/defenseunicorns/uds-runtime/commit/f1dce2ebf81a2701267cc76a92a34d9b59293d9a))


### Bug Fixes

* closes pepr goroutines when navigating away ([#400](https://github.com/defenseunicorns/uds-runtime/issues/400)) ([95fb845](https://github.com/defenseunicorns/uds-runtime/commit/95fb8453779042f3b695351127d0097162a856cb))
* reconnection handling after introducing TLS ([#412](https://github.com/defenseunicorns/uds-runtime/issues/412)) ([b89cf16](https://github.com/defenseunicorns/uds-runtime/commit/b89cf16ddba0ec37891aab72876a26f6d5f0b402))
* storageclass view data ([#425](https://github.com/defenseunicorns/uds-runtime/issues/425)) ([f228e7f](https://github.com/defenseunicorns/uds-runtime/commit/f228e7f2eb619b0a599deff6b519113c51d7a7bd))


### Miscellaneous

* **api:** refactor organization of reconnection logic ([#406](https://github.com/defenseunicorns/uds-runtime/issues/406)) ([1dd4a06](https://github.com/defenseunicorns/uds-runtime/commit/1dd4a06153c79301a6ae23001bcd0aef5002e3db))
* **api:** updates for handling unavailable metrics server ([#421](https://github.com/defenseunicorns/uds-runtime/issues/421)) ([6bb9728](https://github.com/defenseunicorns/uds-runtime/commit/6bb9728d2e01de73b6462b420c384c64580524e7))
* **deps:** update dependency kubernetes-fluent-client to v3.0.4 ([#403](https://github.com/defenseunicorns/uds-runtime/issues/403)) ([ef35113](https://github.com/defenseunicorns/uds-runtime/commit/ef35113c05ecf881c5077fdea4c9c322febbff66))
* **deps:** update devdependencies ([#214](https://github.com/defenseunicorns/uds-runtime/issues/214)) ([16f75ed](https://github.com/defenseunicorns/uds-runtime/commit/16f75ed767af5fd2f128044723e6c010b46be576))
* **deps:** update github actions ([#385](https://github.com/defenseunicorns/uds-runtime/issues/385)) ([ac2797f](https://github.com/defenseunicorns/uds-runtime/commit/ac2797f73abf450e0ba784b269752f61fb322599))
* **deps:** update module github.com/zarf-dev/zarf to v0.41.0 ([#414](https://github.com/defenseunicorns/uds-runtime/issues/414)) ([aaa2fed](https://github.com/defenseunicorns/uds-runtime/commit/aaa2fed5529f3e78b7b70d500eb4cac371907c06))
* fix typo in README ([#396](https://github.com/defenseunicorns/uds-runtime/issues/396)) ([ceddd9b](https://github.com/defenseunicorns/uds-runtime/commit/ceddd9ba36ff86454a9cbdba00d40bf8ceb57a8f))
* fix typo, clean URL, de-dup toast ([#407](https://github.com/defenseunicorns/uds-runtime/issues/407)) ([b8616c6](https://github.com/defenseunicorns/uds-runtime/commit/b8616c619b11feef6a6e3376f7bf8a965c84ff5e))
* implementation for handling crds that dont exist ([#408](https://github.com/defenseunicorns/uds-runtime/issues/408)) ([d5d7d14](https://github.com/defenseunicorns/uds-runtime/commit/d5d7d14962304b955d8b9a7eff4a78eb9559583d))
* swap svelte-chartjs for chartjs lib ([#246](https://github.com/defenseunicorns/uds-runtime/issues/246)) ([cedaa43](https://github.com/defenseunicorns/uds-runtime/commit/cedaa43b1b065be77add427777f03a7837a61b15))
* **ui:** add status colors for event type in events table ([#422](https://github.com/defenseunicorns/uds-runtime/issues/422)) ([e1111e9](https://github.com/defenseunicorns/uds-runtime/commit/e1111e96163eaa67c2f36e9b2dc8cb0501cfbbd5))
* **ui:** adding convention for figma to front end component naming ([#401](https://github.com/defenseunicorns/uds-runtime/issues/401)) ([f65ec88](https://github.com/defenseunicorns/uds-runtime/commit/f65ec88d95656f7217a5af1320b009816bd99f64))

## [0.5.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.4.0...v0.5.0) (2024-09-25)


### Features

* **api:** adds jwt group validation when in-cluster ([#387](https://github.com/defenseunicorns/uds-runtime/issues/387)) ([8a53f76](https://github.com/defenseunicorns/uds-runtime/commit/8a53f76fc684f3e2562ab3076b67d7a68571589d))
* make deployment resources configurable and up default memory ([#372](https://github.com/defenseunicorns/uds-runtime/issues/372)) ([2981598](https://github.com/defenseunicorns/uds-runtime/commit/29815984b28b793087ede78a5de59e5376903477))
* **ui:** adding events tab ([#342](https://github.com/defenseunicorns/uds-runtime/issues/342)) ([cb9b43a](https://github.com/defenseunicorns/uds-runtime/commit/cb9b43a84a3d157b9759a063788e1cecf9e9e868))
* **ui:** fixing issue with progress bar ([#375](https://github.com/defenseunicorns/uds-runtime/issues/375)) ([3f8f204](https://github.com/defenseunicorns/uds-runtime/commit/3f8f20442b89f04af04aa7cd62655fe98d716063))
* **ui:** updating pods table column name ([#369](https://github.com/defenseunicorns/uds-runtime/issues/369)) ([25aaaf9](https://github.com/defenseunicorns/uds-runtime/commit/25aaaf9a630b785a4b8b1b123d29a4df148e0288))


### Bug Fixes

* ensure pod counts are consistent ([#363](https://github.com/defenseunicorns/uds-runtime/issues/363)) ([a5c837b](https://github.com/defenseunicorns/uds-runtime/commit/a5c837b767a316a89a01cdcbebf36a6e407a4883))
* fixing issue with firefox not supporting text-wrap: pretty ([#382](https://github.com/defenseunicorns/uds-runtime/issues/382)) ([4465617](https://github.com/defenseunicorns/uds-runtime/commit/44656170b50ce308cc61eb6b1aaa3ca325926080))
* fixing memory bar for overview page ([#381](https://github.com/defenseunicorns/uds-runtime/issues/381)) ([4f321e3](https://github.com/defenseunicorns/uds-runtime/commit/4f321e348f57fd5d8d15fc3f383a3f64b403d9fd))
* use tls 1.2 in canary deployment ([#383](https://github.com/defenseunicorns/uds-runtime/issues/383)) ([08b5bdc](https://github.com/defenseunicorns/uds-runtime/commit/08b5bdc6e9df063ce06466276bf78312f9e14aaf))


### Miscellaneous

* api route test helper with retries and exponential backoff ([#357](https://github.com/defenseunicorns/uds-runtime/issues/357)) ([674f37e](https://github.com/defenseunicorns/uds-runtime/commit/674f37ed94dfcf1722f8f23bbcd2434cf1adde66))
* **ci:** update core demo bundle version for ephemeral env ([#360](https://github.com/defenseunicorns/uds-runtime/issues/360)) ([161fa7e](https://github.com/defenseunicorns/uds-runtime/commit/161fa7ee3685387a9ce4469b6e07dfa3082aec17))
* **deps:** update dependency vite to v5.3.6 [security] ([#341](https://github.com/defenseunicorns/uds-runtime/issues/341)) ([6dc4ad2](https://github.com/defenseunicorns/uds-runtime/commit/6dc4ad22400ee48f71e2197441fa40253c6bf9c6))
* **deps:** update github actions ([#334](https://github.com/defenseunicorns/uds-runtime/issues/334)) ([117410e](https://github.com/defenseunicorns/uds-runtime/commit/117410e92c7a6d7d5324d50338f34b11484e9818))
* **deps:** update github actions ([#371](https://github.com/defenseunicorns/uds-runtime/issues/371)) ([78bb3a1](https://github.com/defenseunicorns/uds-runtime/commit/78bb3a1aa73de7d5f6145d6a7310d3460dbbdd46))
* **deps:** update module github.com/zarf-dev/zarf to v0.40.1 ([#359](https://github.com/defenseunicorns/uds-runtime/issues/359)) ([12a0f5e](https://github.com/defenseunicorns/uds-runtime/commit/12a0f5ed8c64c6862db43408b7e878a996795673))
* **deps:** update uds-core-types digest to df4d2da ([#362](https://github.com/defenseunicorns/uds-runtime/issues/362)) ([0a67913](https://github.com/defenseunicorns/uds-runtime/commit/0a679136bda31b46e66008bfbf30f8f8ddc61df8))
* refactors uds tasks and adds smoke test ([#344](https://github.com/defenseunicorns/uds-runtime/issues/344)) ([2fec985](https://github.com/defenseunicorns/uds-runtime/commit/2fec985eecf3026a2cd05dda08aca1845ac0479c))
* update description for UDS Package ([#356](https://github.com/defenseunicorns/uds-runtime/issues/356)) ([d360d38](https://github.com/defenseunicorns/uds-runtime/commit/d360d38f8e50648c6e6e167c70a12233ff2eef23))
* update slim core with authsvc deployment and bump k3d version ([#389](https://github.com/defenseunicorns/uds-runtime/issues/389)) ([216bab6](https://github.com/defenseunicorns/uds-runtime/commit/216bab6a2735a1d8d2bbf7b55d6a8369579e81a3))

## [0.4.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.3.0...v0.4.0) (2024-09-19)


### Features

* alert users to cluster disconnection and reconnect ([#270](https://github.com/defenseunicorns/uds-runtime/issues/270)) ([d121bdc](https://github.com/defenseunicorns/uds-runtime/commit/d121bdc0a9632df96a2fc68a447089538e21c05a))
* **api:** implement authsvc in uds package and use jwt in api ([#335](https://github.com/defenseunicorns/uds-runtime/issues/335)) ([def3b83](https://github.com/defenseunicorns/uds-runtime/commit/def3b83fd36f59093460d1f0a52289055a097087))
* **ui:** add functionality to Pepr table ([#298](https://github.com/defenseunicorns/uds-runtime/issues/298)) ([a11d397](https://github.com/defenseunicorns/uds-runtime/commit/a11d39736a203e4f4240b6997332f97e234d43aa))
* **ui:** adding functionality to cards ([#327](https://github.com/defenseunicorns/uds-runtime/issues/327)) ([15dbd05](https://github.com/defenseunicorns/uds-runtime/commit/15dbd05e9049c2533f7cd9d5c8134ef3da872354))
* **ui:** adding progress bar component ([#337](https://github.com/defenseunicorns/uds-runtime/issues/337)) ([a36a1ef](https://github.com/defenseunicorns/uds-runtime/commit/a36a1ef44451129d43ce104995a87bad413295c8))
* **ui:** adding truncate and line clamp option ([#279](https://github.com/defenseunicorns/uds-runtime/issues/279)) ([de760fb](https://github.com/defenseunicorns/uds-runtime/commit/de760fb1217b635847a651c98658da732f7a6d6e))
* **ui:** adjusting drawer margins ([#303](https://github.com/defenseunicorns/uds-runtime/issues/303)) ([c23c809](https://github.com/defenseunicorns/uds-runtime/commit/c23c80933e2d5f978b7106aba8a7ab3bcefaeecc))
* **ui:** confirm status coloring ([#318](https://github.com/defenseunicorns/uds-runtime/issues/318)) ([eb31544](https://github.com/defenseunicorns/uds-runtime/commit/eb31544de52f72574be97cd5604601dffe09c3c0))
* **ui:** reload data view when cluster reconnection is successful ([#336](https://github.com/defenseunicorns/uds-runtime/issues/336)) ([95137bc](https://github.com/defenseunicorns/uds-runtime/commit/95137bcec835031a82282a4a115d29817d34d5bd))
* **ui:** update prettier plugin organize imports ([#324](https://github.com/defenseunicorns/uds-runtime/issues/324)) ([115cc6b](https://github.com/defenseunicorns/uds-runtime/commit/115cc6b8fe99a867c78b42e4e092861478de9014))


### Bug Fixes

* **api:** errors from trying to get kubecontext while in cluster ([#339](https://github.com/defenseunicorns/uds-runtime/issues/339)) ([23488f8](https://github.com/defenseunicorns/uds-runtime/commit/23488f8a69a40b00cf815d14f618a437c5ee6357))
* **api:** flaky api tests ([#350](https://github.com/defenseunicorns/uds-runtime/issues/350)) ([1ef0a77](https://github.com/defenseunicorns/uds-runtime/commit/1ef0a771b61e8bda364a2c2782eb8a1cc2250958))
* field selectors broken when metadata not present ([#338](https://github.com/defenseunicorns/uds-runtime/issues/338)) ([1a7443a](https://github.com/defenseunicorns/uds-runtime/commit/1a7443a053c8b12884354d67681a609bd0e32b0c))
* historical usage loading ([#340](https://github.com/defenseunicorns/uds-runtime/issues/340)) ([f236ad9](https://github.com/defenseunicorns/uds-runtime/commit/f236ad923960c1170ffc72b29b373bd0640af1de))
* **ui:** fixing issue with table data being null ([#320](https://github.com/defenseunicorns/uds-runtime/issues/320)) ([08c8c40](https://github.com/defenseunicorns/uds-runtime/commit/08c8c40526e61b42dcaa72e90b974ddae6f72a02))


### Miscellaneous

* api auth state design doc ([#331](https://github.com/defenseunicorns/uds-runtime/issues/331)) ([59ee0da](https://github.com/defenseunicorns/uds-runtime/commit/59ee0da8be37796115e5a9da9a40828e14ebefbb))
* **api:** complete api tests for all routes ([#319](https://github.com/defenseunicorns/uds-runtime/issues/319)) ([734cc33](https://github.com/defenseunicorns/uds-runtime/commit/734cc336973b8e458f82ee4fea0ce0b545e17a6a))
* **deps:** update dependency kubernetes-fluent-client to v3.0.3 ([#329](https://github.com/defenseunicorns/uds-runtime/issues/329)) ([467dcab](https://github.com/defenseunicorns/uds-runtime/commit/467dcabdd45d5ec3f07b5004682af7b846159a18))
* **deps:** update github actions ([#266](https://github.com/defenseunicorns/uds-runtime/issues/266)) ([898acdb](https://github.com/defenseunicorns/uds-runtime/commit/898acdbf8de6c6af7b3e90b46a41c3119ba4e90e))
* **deps:** update kubernetes packages to v0.31.1 ([#321](https://github.com/defenseunicorns/uds-runtime/issues/321)) ([0865775](https://github.com/defenseunicorns/uds-runtime/commit/086577577580b120fa67dcb220c107d27708b319))
* move api token auth state management to backend using cookies ([#343](https://github.com/defenseunicorns/uds-runtime/issues/343)) ([45a4e76](https://github.com/defenseunicorns/uds-runtime/commit/45a4e76683b57954820019d24e8670e361727d1d))
* undo release-as 0.3.0 ([#291](https://github.com/defenseunicorns/uds-runtime/issues/291)) ([3ee2cac](https://github.com/defenseunicorns/uds-runtime/commit/3ee2cac0fe3bf4866bc46bff8f337e00799eb027))
* update readme with badges ([#225](https://github.com/defenseunicorns/uds-runtime/issues/225)) ([38df19a](https://github.com/defenseunicorns/uds-runtime/commit/38df19a3ba92594766ee0929f8e5dbe03fb577fe))

## [0.3.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.3.0...v0.3.0) (2024-09-06)


### Features

* add missing routes & route filter ([4235c60](https://github.com/defenseunicorns/uds-runtime/commit/4235c60b5243a606c3c02129c3508655cfb6d51a))
* add pod disruption budgets view ([#79](https://github.com/defenseunicorns/uds-runtime/issues/79)) ([a2f17b3](https://github.com/defenseunicorns/uds-runtime/commit/a2f17b39cd0b0304c5bfb9e1cea9b2320f62ef76)), closes [#48](https://github.com/defenseunicorns/uds-runtime/issues/48)
* add sparse streams, compression, single resource api endpoints ([#28](https://github.com/defenseunicorns/uds-runtime/issues/28)) ([c5c4c9c](https://github.com/defenseunicorns/uds-runtime/commit/c5c4c9c5e98cd81e4ea5a0e295983ced1ce05832))
* add storageclasses view ([#83](https://github.com/defenseunicorns/uds-runtime/issues/83)) ([941fff2](https://github.com/defenseunicorns/uds-runtime/commit/941fff2715ee85d2dfda0fb26fbb59fc10f2708f)), closes [#60](https://github.com/defenseunicorns/uds-runtime/issues/60)
* add UDS Exemptions view ([ac0d43e](https://github.com/defenseunicorns/uds-runtime/commit/ac0d43ec947d47341aa4e233aebeefe2e51aeb78))
* add uds package view ([7be0628](https://github.com/defenseunicorns/uds-runtime/commit/7be06284014e0f4b0337c345873b1e0951c09951))
* adding adr for data-test ids ([#200](https://github.com/defenseunicorns/uds-runtime/issues/200)) ([4936722](https://github.com/defenseunicorns/uds-runtime/commit/4936722b96202d97294049c4bc839f3a2231d083))
* adding cluster ops limit ranges view ([#78](https://github.com/defenseunicorns/uds-runtime/issues/78)) ([6907151](https://github.com/defenseunicorns/uds-runtime/commit/69071515bdefa8d50eb04465e29b96b1fa6b2c5f)), closes [#47](https://github.com/defenseunicorns/uds-runtime/issues/47)
* adding drawer e2e tests ([#164](https://github.com/defenseunicorns/uds-runtime/issues/164)) ([f19ca04](https://github.com/defenseunicorns/uds-runtime/commit/f19ca04c90dd03f9b96639beb26e787f3f3923f6))
* adding e2e tests for DataTable ([#188](https://github.com/defenseunicorns/uds-runtime/issues/188)) ([a7c6256](https://github.com/defenseunicorns/uds-runtime/commit/a7c625694fa088103c2a0ee7c7b3f75281ad19bb))
* adding e2e tests for search and dropdown selections ([#193](https://github.com/defenseunicorns/uds-runtime/issues/193)) ([e1a6c20](https://github.com/defenseunicorns/uds-runtime/commit/e1a6c207ea466853003886ae77c7bfcdd6de1995))
* adds pods to pvc view, make swagger more discoverable ([#271](https://github.com/defenseunicorns/uds-runtime/issues/271)) ([d6e5eba](https://github.com/defenseunicorns/uds-runtime/commit/d6e5eba45e4cb5649977cfdbed78cf65c3c1eb23))
* adds pv view ([#71](https://github.com/defenseunicorns/uds-runtime/issues/71)) ([cd013d5](https://github.com/defenseunicorns/uds-runtime/commit/cd013d5271e049310b2d3a6aaf9a0ae2a5b5f5dc))
* **api:** add fields query parameter for fine-grained data filtering ([#94](https://github.com/defenseunicorns/uds-runtime/issues/94)) ([0021790](https://github.com/defenseunicorns/uds-runtime/commit/002179074a4b2ba90eb3e7b62f1e355c6c0717f4))
* **api:** add remaining k8s resource endpoints ([7fa24ee](https://github.com/defenseunicorns/uds-runtime/commit/7fa24ee72ea7ba9ee38f079aa5ed4863f497650b))
* **api:** debounce SSE messages ([#100](https://github.com/defenseunicorns/uds-runtime/issues/100)) ([27d7433](https://github.com/defenseunicorns/uds-runtime/commit/27d7433d37bbae977bc9f0269f717056fdcb0d6c))
* **api:** strip annotations from sparse resources ([#250](https://github.com/defenseunicorns/uds-runtime/issues/250)) ([f43533a](https://github.com/defenseunicorns/uds-runtime/commit/f43533aa75c1f2c8b697b87bab8feaa8147b9f14))
* application views + dashboard ([#144](https://github.com/defenseunicorns/uds-runtime/issues/144)) ([a4c5110](https://github.com/defenseunicorns/uds-runtime/commit/a4c5110025241d38d241f17e10c236b53e00c313))
* **ci:** add nightly releases ([#162](https://github.com/defenseunicorns/uds-runtime/issues/162)) ([d523a64](https://github.com/defenseunicorns/uds-runtime/commit/d523a647d54c144455f4a4808a4770ae05ac91b8))
* enables uds-runtime releases ([#93](https://github.com/defenseunicorns/uds-runtime/issues/93)) ([6af3a95](https://github.com/defenseunicorns/uds-runtime/commit/6af3a95e574f95f1df8ad39a696b311f2c63d16e))
* export pepr data stream ([#24](https://github.com/defenseunicorns/uds-runtime/issues/24)) ([b05976f](https://github.com/defenseunicorns/uds-runtime/commit/b05976fad98a3eaa543f21db8e4f24a3b99805de)), closes [#23](https://github.com/defenseunicorns/uds-runtime/issues/23)
* local mode token protection ([#245](https://github.com/defenseunicorns/uds-runtime/issues/245)) ([d356940](https://github.com/defenseunicorns/uds-runtime/commit/d356940af5e63244381ab408eca854d0ea669761))
* make root URL the overview page ([#212](https://github.com/defenseunicorns/uds-runtime/issues/212)) ([d273a1e](https://github.com/defenseunicorns/uds-runtime/commit/d273a1e8e3ed842e7cb8d1583389fe0f1e25a102))
* swagger ([#114](https://github.com/defenseunicorns/uds-runtime/issues/114)) ([c4016ca](https://github.com/defenseunicorns/uds-runtime/commit/c4016caaef1d1e5e0b6114d351521b0f80e99663))
* **ui:** 112 stylized and colorized yaml view ([#121](https://github.com/defenseunicorns/uds-runtime/issues/121)) ([1ea6532](https://github.com/defenseunicorns/uds-runtime/commit/1ea65325b1e9bcec758e9aed091e186f9d3c257b)), closes [#112](https://github.com/defenseunicorns/uds-runtime/issues/112)
* **ui:** add cluster ops hpa k8s view ([#97](https://github.com/defenseunicorns/uds-runtime/issues/97)) ([a55696d](https://github.com/defenseunicorns/uds-runtime/commit/a55696defbdf7215f752fc36a8d666b6995476f0)), closes [#46](https://github.com/defenseunicorns/uds-runtime/issues/46)
* **ui:** add cluster-ops/mutating-webhooks view ([#42](https://github.com/defenseunicorns/uds-runtime/issues/42)) ([b29c434](https://github.com/defenseunicorns/uds-runtime/commit/b29c434e0c5300194d57b9b831b0528ca3ffd743)), closes [#32](https://github.com/defenseunicorns/uds-runtime/issues/32)
* **ui:** add cronjobs view ([9a80579](https://github.com/defenseunicorns/uds-runtime/commit/9a80579fdcfa14af2426314dfdfb11f89089b72c))
* **ui:** add header to tables ([#106](https://github.com/defenseunicorns/uds-runtime/issues/106)) ([19cdcfe](https://github.com/defenseunicorns/uds-runtime/commit/19cdcfe4ef3c602bada5edc1e560dec311bcda60))
* **ui:** add jobs view ([125f635](https://github.com/defenseunicorns/uds-runtime/commit/125f63517ad9dc69993b7b178cd10a287896b7c9))
* **ui:** add nodes view ([f1cf0b0](https://github.com/defenseunicorns/uds-runtime/commit/f1cf0b0725a6deb4fd33eb74086a31de8059c875))
* **ui:** add pepr details ([#107](https://github.com/defenseunicorns/uds-runtime/issues/107)) ([a4ed8cd](https://github.com/defenseunicorns/uds-runtime/commit/a4ed8cd591ab1c691209fd55bd1bb764f3756952))
* **ui:** add pvc view ([#72](https://github.com/defenseunicorns/uds-runtime/issues/72)) ([5580dd4](https://github.com/defenseunicorns/uds-runtime/commit/5580dd4ece5725c248f8369607a67a31f0e20cdf))
* **ui:** add toast for status messages ([#116](https://github.com/defenseunicorns/uds-runtime/issues/116)) ([983ae6a](https://github.com/defenseunicorns/uds-runtime/commit/983ae6af3b9f9b55d29926b178ebcc228bb7e83f)), closes [#115](https://github.com/defenseunicorns/uds-runtime/issues/115)
* **ui:** adding 404 ([#98](https://github.com/defenseunicorns/uds-runtime/issues/98)) ([24b24b0](https://github.com/defenseunicorns/uds-runtime/commit/24b24b0516cbe563b4e6e7e2ebe7495173c4da5b)), closes [#38](https://github.com/defenseunicorns/uds-runtime/issues/38)
* **ui:** adding cluster ops priority classes view ([#62](https://github.com/defenseunicorns/uds-runtime/issues/62)) ([f52ea3d](https://github.com/defenseunicorns/uds-runtime/commit/f52ea3d47aabd868507e563885fd9133b03bd0af))
* **ui:** adding cluster ops resource quotas view ([#81](https://github.com/defenseunicorns/uds-runtime/issues/81)) ([5370bf9](https://github.com/defenseunicorns/uds-runtime/commit/5370bf98e6e7f65987933de13de863a7a0dd3cd4))
* **ui:** adding cluster ops runtime classes view ([#68](https://github.com/defenseunicorns/uds-runtime/issues/68)) ([084fc83](https://github.com/defenseunicorns/uds-runtime/commit/084fc831fb75a6a0bd70427888cea5f2bd6f57fa))
* **ui:** adding status color mapping for k8s ([#264](https://github.com/defenseunicorns/uds-runtime/issues/264)) ([15fc75f](https://github.com/defenseunicorns/uds-runtime/commit/15fc75f35c529c70ffd3cd0e53a22802ac1691e8))
* **ui:** adds validatingwebhooks view ([#110](https://github.com/defenseunicorns/uds-runtime/issues/110)) ([cad4f35](https://github.com/defenseunicorns/uds-runtime/commit/cad4f3508fe301532a71e4bd04986f7113df406d))
* **ui:** endpoints view ([#80](https://github.com/defenseunicorns/uds-runtime/issues/80)) ([e9bf61e](https://github.com/defenseunicorns/uds-runtime/commit/e9bf61e3cb284769e6402a4a69e97bf99dd8e347))
* **ui:** format age with more detail ([#111](https://github.com/defenseunicorns/uds-runtime/issues/111)) ([98916d6](https://github.com/defenseunicorns/uds-runtime/commit/98916d6328d9c131aa5394058eaeb77b10ec8bf3))
* **ui:** make package endpoints links ([#136](https://github.com/defenseunicorns/uds-runtime/issues/136)) ([572f385](https://github.com/defenseunicorns/uds-runtime/commit/572f3859b0ba5901182fe335eb5e9ac679635716))
* **ui:** networkpolicies view ([#75](https://github.com/defenseunicorns/uds-runtime/issues/75)) ([d04d50b](https://github.com/defenseunicorns/uds-runtime/commit/d04d50bab0c23ba15952a7977a3046e280e39f4f))
* **ui:** networks-services view ([#61](https://github.com/defenseunicorns/uds-runtime/issues/61)) ([1d0b68a](https://github.com/defenseunicorns/uds-runtime/commit/1d0b68aa39ad59e79107ddb61d650ae427cc9598))
* **ui:** resource detail view ([#95](https://github.com/defenseunicorns/uds-runtime/issues/95)) ([34e7e6c](https://github.com/defenseunicorns/uds-runtime/commit/34e7e6ca168fd3baa7b12dae7d4edd66f6343a8a)), closes [#34](https://github.com/defenseunicorns/uds-runtime/issues/34)
* **ui:** show results even if not filtering ([#131](https://github.com/defenseunicorns/uds-runtime/issues/131)) ([4603b65](https://github.com/defenseunicorns/uds-runtime/commit/4603b656b41e3a94883a4694a4a881f27512424b))
* **ui:** updating drawer header ([#142](https://github.com/defenseunicorns/uds-runtime/issues/142)) ([98dad19](https://github.com/defenseunicorns/uds-runtime/commit/98dad199800d06a0fb2690673dc35785784400a0)), closes [#102](https://github.com/defenseunicorns/uds-runtime/issues/102)
* **ui:** virtualservices view ([#73](https://github.com/defenseunicorns/uds-runtime/issues/73)) ([4d01cb8](https://github.com/defenseunicorns/uds-runtime/commit/4d01cb8db3125c747f5e26b835edc9fd0cf61131))


### Bug Fixes

* add missing routes ([#6](https://github.com/defenseunicorns/uds-runtime/issues/6)) ([3e61370](https://github.com/defenseunicorns/uds-runtime/commit/3e61370bb0443beefcb7b69f298570baca8837b2))
* add tasks file to release please config and add docker login ([#139](https://github.com/defenseunicorns/uds-runtime/issues/139)) ([f61680f](https://github.com/defenseunicorns/uds-runtime/commit/f61680f2941bb8e092b8da69d5c9cbcdcbb12d43))
* api token auth details view ([#276](https://github.com/defenseunicorns/uds-runtime/issues/276)) ([871cdf5](https://github.com/defenseunicorns/uds-runtime/commit/871cdf5f9561f1241a77f4b6271d94cb0dfabc7b))
* **api:** dockerfile for releasing go binaries ([#258](https://github.com/defenseunicorns/uds-runtime/issues/258)) ([e35987c](https://github.com/defenseunicorns/uds-runtime/commit/e35987c01acc82fbab2023e691483a281e1673ac))
* **ci:** more permissions for release ([#135](https://github.com/defenseunicorns/uds-runtime/issues/135)) ([3695630](https://github.com/defenseunicorns/uds-runtime/commit/36956305edfbcf2fa94befb20d900d1c7503d6b3))
* **ci:** swag init task output directory ([#248](https://github.com/defenseunicorns/uds-runtime/issues/248)) ([0d65c2e](https://github.com/defenseunicorns/uds-runtime/commit/0d65c2e97c8202729bdae9519892e47455e66ddf))
* correct nightly working dir ([#196](https://github.com/defenseunicorns/uds-runtime/issues/196)) ([f61dea2](https://github.com/defenseunicorns/uds-runtime/commit/f61dea22ee20a41579d31dbd583e6e63ae7a8975))
* docker container run failure ([#267](https://github.com/defenseunicorns/uds-runtime/issues/267)) ([573a605](https://github.com/defenseunicorns/uds-runtime/commit/573a6053820cc8f1ff55589e38b45c44f59788ab))
* empty endpoints are shown as Pending ([#158](https://github.com/defenseunicorns/uds-runtime/issues/158)) ([08b96ad](https://github.com/defenseunicorns/uds-runtime/commit/08b96ad9334f76af99856e0c9e9198aea2614c8e))
* enable release-please PRs to run workflows ([#146](https://github.com/defenseunicorns/uds-runtime/issues/146)) ([10d7858](https://github.com/defenseunicorns/uds-runtime/commit/10d7858adcffd2f977c6101a25b81696dd266f00))
* ensure nightly artifacts use proper tag ([#190](https://github.com/defenseunicorns/uds-runtime/issues/190)) ([bb44ae3](https://github.com/defenseunicorns/uds-runtime/commit/bb44ae33ddd2d220cb14410b17f828823ad03b4d))
* fixing bug for issue [#147](https://github.com/defenseunicorns/uds-runtime/issues/147) ([#149](https://github.com/defenseunicorns/uds-runtime/issues/149)) ([4d6122d](https://github.com/defenseunicorns/uds-runtime/commit/4d6122dede264208238c2e3e26747db39489a58c))
* fixing FOUC issue ([#236](https://github.com/defenseunicorns/uds-runtime/issues/236)) ([99886ea](https://github.com/defenseunicorns/uds-runtime/commit/99886eab9806faed4c269f3bea4d99ca67248d36))
* fixing issue with chart not being destroyed onDestroy ([#243](https://github.com/defenseunicorns/uds-runtime/issues/243)) ([4bd9baf](https://github.com/defenseunicorns/uds-runtime/commit/4bd9bafc67cd39570aa94b1dccf9072933cb360e))
* fixing tooltip behavior ([#26](https://github.com/defenseunicorns/uds-runtime/issues/26)) ([dd8b092](https://github.com/defenseunicorns/uds-runtime/commit/dd8b0928be7e0aa4b9c656a8ac83e8656fa4f530)), closes [#25](https://github.com/defenseunicorns/uds-runtime/issues/25)
* give ssm-user more permissions ([#228](https://github.com/defenseunicorns/uds-runtime/issues/228)) ([a55626c](https://github.com/defenseunicorns/uds-runtime/commit/a55626c7a0645de7bd3c6743c27f6113cad49343))
* nightly release ([#260](https://github.com/defenseunicorns/uds-runtime/issues/260)) ([a99bc23](https://github.com/defenseunicorns/uds-runtime/commit/a99bc23473a31ee7b4761bf50b0b162c388eed51))
* nightly release artifact fix ([#259](https://github.com/defenseunicorns/uds-runtime/issues/259)) ([648600d](https://github.com/defenseunicorns/uds-runtime/commit/648600dcf811ef2cf060389b9682a8ee8ff82dbc))
* pod status color ([#176](https://github.com/defenseunicorns/uds-runtime/issues/176)) ([2fe0c99](https://github.com/defenseunicorns/uds-runtime/commit/2fe0c99e51dea34bbec0e5383ef2890e563d9f8d))
* remove alpha from release ([#145](https://github.com/defenseunicorns/uds-runtime/issues/145)) ([70344d8](https://github.com/defenseunicorns/uds-runtime/commit/70344d8a07a33df57d90cc848974e08fde327ac7))
* remove breadcrumb links, closes [#9](https://github.com/defenseunicorns/uds-runtime/issues/9) ([e2d8cfe](https://github.com/defenseunicorns/uds-runtime/commit/e2d8cfecb5c1ce1c863435eabf42937f57100401))
* removes empty docs page, adds favicon ([#234](https://github.com/defenseunicorns/uds-runtime/issues/234)) ([db3212c](https://github.com/defenseunicorns/uds-runtime/commit/db3212c9ad0a33df1576ec9866917eba20b0ccf4))
* sidebar expand toggle icon rotation ([e1aa7cb](https://github.com/defenseunicorns/uds-runtime/commit/e1aa7cbe97faae1044aa895ee338427ab3b55180))
* sidebar ux issues ([589a970](https://github.com/defenseunicorns/uds-runtime/commit/589a9700794e10332e953f63682c4b3d970797cd))
* turn off autocomplete for datatable ([#22](https://github.com/defenseunicorns/uds-runtime/issues/22)) ([f71eccf](https://github.com/defenseunicorns/uds-runtime/commit/f71eccfbe084f2c9808b970fff16c3c8d12c4ad1)), closes [#21](https://github.com/defenseunicorns/uds-runtime/issues/21)
* **ui:** deployments table not showing accurate available count ([#192](https://github.com/defenseunicorns/uds-runtime/issues/192)) ([5b0f394](https://github.com/defenseunicorns/uds-runtime/commit/5b0f394ef80ea6286a017093851b4f42ae16009c))
* **ui:** details reactivity on mouse/keyboard navigate ([#130](https://github.com/defenseunicorns/uds-runtime/issues/130)) ([02da735](https://github.com/defenseunicorns/uds-runtime/commit/02da735dd73c56d01ac1ea67baa09f8192c8abed))
* **ui:** fixing issue with graph menu color ([#195](https://github.com/defenseunicorns/uds-runtime/issues/195)) ([b10f95f](https://github.com/defenseunicorns/uds-runtime/commit/b10f95f5595a4341b93356f671dca6b818abc31b))
* **ui:** pepr monitor unsubscriber ([#67](https://github.com/defenseunicorns/uds-runtime/issues/67)) ([2995497](https://github.com/defenseunicorns/uds-runtime/commit/2995497a02fb1b6b6a261aad614c4b4bfe84bc98)), closes [#37](https://github.com/defenseunicorns/uds-runtime/issues/37)
* **ui:** prevent error on null fields in DataTable component ([#43](https://github.com/defenseunicorns/uds-runtime/issues/43)) ([d51636e](https://github.com/defenseunicorns/uds-runtime/commit/d51636e5ee9154c1863d06222d927c6e61f26cb1))
* **ui:** restart and node columns in pod table ([#194](https://github.com/defenseunicorns/uds-runtime/issues/194)) ([88a41c4](https://github.com/defenseunicorns/uds-runtime/commit/88a41c4fc4838945361d8d3117086664273b3a50))
* use docker buildx in github actions ([#133](https://github.com/defenseunicorns/uds-runtime/issues/133)) ([d52cb6c](https://github.com/defenseunicorns/uds-runtime/commit/d52cb6cebb3eaab211d0087726b7c3f2a4d18b4b))
* use middleware for compression ([#86](https://github.com/defenseunicorns/uds-runtime/issues/86)) ([f16a18c](https://github.com/defenseunicorns/uds-runtime/commit/f16a18c755f67cad6ffdd3a21d1708af4f99813e)), closes [#85](https://github.com/defenseunicorns/uds-runtime/issues/85)


### Miscellaneous

* add contributing guide ([#17](https://github.com/defenseunicorns/uds-runtime/issues/17)) ([f03c08f](https://github.com/defenseunicorns/uds-runtime/commit/f03c08f73abccdc7dd69744fdbd5a25ad4d6371f)), closes [#18](https://github.com/defenseunicorns/uds-runtime/issues/18)
* add empty release please manifest to bootstrap ([#120](https://github.com/defenseunicorns/uds-runtime/issues/120)) ([139fe0a](https://github.com/defenseunicorns/uds-runtime/commit/139fe0a0dc32d04c6728b0de3f82a102ef7638f9))
* add env for gh token release workflow ([#286](https://github.com/defenseunicorns/uds-runtime/issues/286)) ([0398329](https://github.com/defenseunicorns/uds-runtime/commit/039832932923eb4c62f21dca01adf3b788e393f6))
* add go binaries as artifacts of releases ([#240](https://github.com/defenseunicorns/uds-runtime/issues/240)) ([422deea](https://github.com/defenseunicorns/uds-runtime/commit/422deea7ef61f48e228c397aeb4489d5026c09df))
* add initial ci ([#11](https://github.com/defenseunicorns/uds-runtime/issues/11)) ([57248fb](https://github.com/defenseunicorns/uds-runtime/commit/57248fb1c2bfe03131b7bc8ce90f6ef419078838))
* add release ci ([#118](https://github.com/defenseunicorns/uds-runtime/issues/118)) ([70c0984](https://github.com/defenseunicorns/uds-runtime/commit/70c09845408ec8e4e62916bef7ae5d4cd4ebc87e))
* add resource store tests ([#74](https://github.com/defenseunicorns/uds-runtime/issues/74)) ([33ad7bc](https://github.com/defenseunicorns/uds-runtime/commit/33ad7bcee1e2231ab2570865932bd29b6194d997))
* add test infra for dummy cluster ([#157](https://github.com/defenseunicorns/uds-runtime/issues/157)) ([2ba9e26](https://github.com/defenseunicorns/uds-runtime/commit/2ba9e261255c57b1631249e3aa1864eb5c7b34f1))
* add uds-cli taks & zarf schema configs for vscode ([344fd7e](https://github.com/defenseunicorns/uds-runtime/commit/344fd7e6ab217e7355e58073e07a8972986249d9))
* add v to release workflow ([#289](https://github.com/defenseunicorns/uds-runtime/issues/289)) ([94d627a](https://github.com/defenseunicorns/uds-runtime/commit/94d627abb256a4cb8e86db529b2ad07b0c02d943))
* add vscode go debug config ([c811f47](https://github.com/defenseunicorns/uds-runtime/commit/c811f47f804cd21d08300d4cc906bc5e6c9f8536))
* adding all routes for navigation ([#16](https://github.com/defenseunicorns/uds-runtime/issues/16)) ([eb518c3](https://github.com/defenseunicorns/uds-runtime/commit/eb518c39a9d45e74fc76f15abd2346ae304478a7)), closes [#15](https://github.com/defenseunicorns/uds-runtime/issues/15)
* adding contributing instructions for front end testing locators ([#242](https://github.com/defenseunicorns/uds-runtime/issues/242)) ([4a253cf](https://github.com/defenseunicorns/uds-runtime/commit/4a253cf78039cc0ce450eb5a00baa6e31e1723b0))
* adds basic e2e tests ([#175](https://github.com/defenseunicorns/uds-runtime/issues/175)) ([770c4a3](https://github.com/defenseunicorns/uds-runtime/commit/770c4a32a0ac1aae4e58cc8af89c6d998513a688))
* adds navbar tests ([#189](https://github.com/defenseunicorns/uds-runtime/issues/189)) ([247afc7](https://github.com/defenseunicorns/uds-runtime/commit/247afc7813fe560648453764f1b891845483a9c4))
* adds ssm to runtime-canary ([#221](https://github.com/defenseunicorns/uds-runtime/issues/221)) ([13d2350](https://github.com/defenseunicorns/uds-runtime/commit/13d2350e0e6f2c2ce6a6d06c9c3e54855a762990))
* api testing adr ([#198](https://github.com/defenseunicorns/uds-runtime/issues/198)) ([f50bd77](https://github.com/defenseunicorns/uds-runtime/commit/f50bd773740300745e19f7821a1a9143a3c12679))
* api token auth documentation ([#269](https://github.com/defenseunicorns/uds-runtime/issues/269)) ([20cbd02](https://github.com/defenseunicorns/uds-runtime/commit/20cbd0265f56b5dbb17c1ca3538cf9f13f1a4987))
* **api:** add api tests ([#231](https://github.com/defenseunicorns/uds-runtime/issues/231)) ([3f9eb2c](https://github.com/defenseunicorns/uds-runtime/commit/3f9eb2caa6cbca61bd7b99b73c482868cb061a20))
* **api:** add field selector api tests ([#262](https://github.com/defenseunicorns/uds-runtime/issues/262)) ([6ee6720](https://github.com/defenseunicorns/uds-runtime/commit/6ee672046101cd965da57111c091d8233489c6ee))
* bump spdx header to 2024-Present ([2102add](https://github.com/defenseunicorns/uds-runtime/commit/2102addd1864b67a82b59e4ac11c4f7fb51f0824))
* **ci:** add swagger docs gen to CI ([#167](https://github.com/defenseunicorns/uds-runtime/issues/167)) ([aed5b83](https://github.com/defenseunicorns/uds-runtime/commit/aed5b836288fb5592b254302355d05cb7e7ee3bc))
* **ci:** add type checking ([#187](https://github.com/defenseunicorns/uds-runtime/issues/187)) ([3357983](https://github.com/defenseunicorns/uds-runtime/commit/3357983a961bc565aca55464dc19b8c6e346cd81))
* codeowners ([#174](https://github.com/defenseunicorns/uds-runtime/issues/174)) ([b4cabcf](https://github.com/defenseunicorns/uds-runtime/commit/b4cabcf34c90d40fc2c6467932ac42c2f34e9fc6))
* configure renovate ([#209](https://github.com/defenseunicorns/uds-runtime/issues/209)) ([4712c6f](https://github.com/defenseunicorns/uds-runtime/commit/4712c6f71d04b37e6f55dc57b2b910ba3c86429c))
* **deps:** update all dependencies ([#13](https://github.com/defenseunicorns/uds-runtime/issues/13)) ([ec95b7e](https://github.com/defenseunicorns/uds-runtime/commit/ec95b7e67e07246489220e1d6783ad9ef195ca74))
* **deps:** update dependency jsdom to v25 ([#247](https://github.com/defenseunicorns/uds-runtime/issues/247)) ([772930b](https://github.com/defenseunicorns/uds-runtime/commit/772930ba4677fcee0f4b242c9f0f4a9f66960a66))
* **deps:** update dependency kubernetes-fluent-client to v3 ([#218](https://github.com/defenseunicorns/uds-runtime/issues/218)) ([245c256](https://github.com/defenseunicorns/uds-runtime/commit/245c25633a5019dc8a52649d2830173c1e95da5c))
* **deps:** update dependency kubernetes-fluent-client to v3.0.2 ([#256](https://github.com/defenseunicorns/uds-runtime/issues/256)) ([4c075a8](https://github.com/defenseunicorns/uds-runtime/commit/4c075a8e108c71c31b0e0354a04272dbd1efd6e2))
* **deps:** update github actions ([#213](https://github.com/defenseunicorns/uds-runtime/issues/213)) ([b51692f](https://github.com/defenseunicorns/uds-runtime/commit/b51692fff7c6f9f7f423ac0e3fbf8188cedd51bd))
* **deps:** update github actions ([#227](https://github.com/defenseunicorns/uds-runtime/issues/227)) ([22b18b0](https://github.com/defenseunicorns/uds-runtime/commit/22b18b0c0d1a42e3f30d5d12849cdc478b500482))
* **deps:** update github actions ([#257](https://github.com/defenseunicorns/uds-runtime/issues/257)) ([6c366cc](https://github.com/defenseunicorns/uds-runtime/commit/6c366ccbe0264a0c95362051ec072c59320c489e))
* **deps:** update kubernetes packages to v0.31.0 ([#215](https://github.com/defenseunicorns/uds-runtime/issues/215)) ([9875e62](https://github.com/defenseunicorns/uds-runtime/commit/9875e62315d17814c861888f55c4d2f88ce0d92a))
* **deps:** update module github.com/charmbracelet/lipgloss to v0.13.0 ([#232](https://github.com/defenseunicorns/uds-runtime/issues/232)) ([1cf315a](https://github.com/defenseunicorns/uds-runtime/commit/1cf315af910c026469f331fb338a03a012270716))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.2 ([#216](https://github.com/defenseunicorns/uds-runtime/issues/216)) ([21c7457](https://github.com/defenseunicorns/uds-runtime/commit/21c745736397f461e9611da42f8b24ff1c89c493))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.3 ([#244](https://github.com/defenseunicorns/uds-runtime/issues/244)) ([adf6d82](https://github.com/defenseunicorns/uds-runtime/commit/adf6d82d175e139f84ec93f7e77f824a2537df5f))
* **deps:** update module github.com/zarf-dev/zarf to v0.39.0 ([#283](https://github.com/defenseunicorns/uds-runtime/issues/283)) ([e3719a2](https://github.com/defenseunicorns/uds-runtime/commit/e3719a2721fe48fe1f6bf208cecf996da6ca6cd4))
* ensure lang=ts for all svelte script tags ([b40ecbb](https://github.com/defenseunicorns/uds-runtime/commit/b40ecbba365fecc7eacd312ee37765a3a297f3b6))
* fix lint errors ([68df75f](https://github.com/defenseunicorns/uds-runtime/commit/68df75fc7f8c6d3693d85c30663fff2cb9e9fe2d))
* go unit tests ([#191](https://github.com/defenseunicorns/uds-runtime/issues/191)) ([b850316](https://github.com/defenseunicorns/uds-runtime/commit/b85031664c6e2cc68cf28379961d30ccf62fabd0))
* install Istio CRDs for minimal e2e tests ([#168](https://github.com/defenseunicorns/uds-runtime/issues/168)) ([8832dcd](https://github.com/defenseunicorns/uds-runtime/commit/8832dcd203cfc3086fa2280d76e4dd901b74f2c9))
* lint ([ef14348](https://github.com/defenseunicorns/uds-runtime/commit/ef1434877c8feb39ab3b1601c02e24269abd43f8))
* **main:** release 0.1.0 ([#143](https://github.com/defenseunicorns/uds-runtime/issues/143)) ([ef1fa4c](https://github.com/defenseunicorns/uds-runtime/commit/ef1fa4ca67007db38984daa42cc0dbef208cb9f9))
* **main:** release 0.1.0-alpha.3 ([#134](https://github.com/defenseunicorns/uds-runtime/issues/134)) ([410ce9b](https://github.com/defenseunicorns/uds-runtime/commit/410ce9bdf54acbef0f2972e78987d05787196063))
* **main:** release 0.1.0-alpha.4 ([#137](https://github.com/defenseunicorns/uds-runtime/issues/137)) ([2907c75](https://github.com/defenseunicorns/uds-runtime/commit/2907c75d313e985fa2f3dd1c7a7a09c06d78efad))
* **main:** release 0.1.0-alpha.5 ([#138](https://github.com/defenseunicorns/uds-runtime/issues/138)) ([12ff0c1](https://github.com/defenseunicorns/uds-runtime/commit/12ff0c138b606bc5abbc4201061c2e341b5deb1f))
* **main:** release 0.2.0 ([#150](https://github.com/defenseunicorns/uds-runtime/issues/150)) ([fb2456d](https://github.com/defenseunicorns/uds-runtime/commit/fb2456d1f4e87e2799836216a2f45e41783f7f59))
* **main:** release 0.3.0 ([#224](https://github.com/defenseunicorns/uds-runtime/issues/224)) ([2d6d21c](https://github.com/defenseunicorns/uds-runtime/commit/2d6d21c5ad762d58ad0ad0be597b0cd82be77b19))
* **main:** release 0.3.0 ([#288](https://github.com/defenseunicorns/uds-runtime/issues/288)) ([ccfc914](https://github.com/defenseunicorns/uds-runtime/commit/ccfc9147a84b635fd1f798d10cfcd3bfb0d2f7b0))
* make uds package host configurable for nightly ([#201](https://github.com/defenseunicorns/uds-runtime/issues/201)) ([bb34173](https://github.com/defenseunicorns/uds-runtime/commit/bb34173b6418d17c9aedc4876941a80927a7645e))
* refactor controller to controlled by ([#122](https://github.com/defenseunicorns/uds-runtime/issues/122)) ([839502d](https://github.com/defenseunicorns/uds-runtime/commit/839502df42618c09b6d28451248bf70aa73453a1))
* remove message.Fatal for zarf upgrade ([8438b37](https://github.com/defenseunicorns/uds-runtime/commit/8438b37532367f160487c77b9dbdce500d6b9e8e))
* **setting-up-release-please:** release 0.1.0-alpha.2 ([#127](https://github.com/defenseunicorns/uds-runtime/issues/127)) ([52bfd99](https://github.com/defenseunicorns/uds-runtime/commit/52bfd99dc58ef17d4bd3b14426f07afc13372f41))
* some go unit tests ([#119](https://github.com/defenseunicorns/uds-runtime/issues/119)) ([a0a7433](https://github.com/defenseunicorns/uds-runtime/commit/a0a74335249ac076e6ac40872abaeeffc0b2e5c6))
* speed up e2e tests with minimal cluster ([#155](https://github.com/defenseunicorns/uds-runtime/issues/155)) ([b8aae8a](https://github.com/defenseunicorns/uds-runtime/commit/b8aae8a9fb31be7be57394817769feeb19ee5747))
* swap to admin gateway ([#124](https://github.com/defenseunicorns/uds-runtime/issues/124)) ([776927b](https://github.com/defenseunicorns/uds-runtime/commit/776927baf5e255b5e80e490240318a0a6a3e34dd))
* tag uds cli install for renovate ([#226](https://github.com/defenseunicorns/uds-runtime/issues/226)) ([e3f5f71](https://github.com/defenseunicorns/uds-runtime/commit/e3f5f71c36a1c7fb3d1823e359ab93e0169d7c91))
* uds-engine -&gt; uds-runtime ([e553195](https://github.com/defenseunicorns/uds-runtime/commit/e553195ffc0cef8dd0b4fca0b6cf350301ecdcfc))
* **ui:** redirect to auth page when unauthenticated ([#278](https://github.com/defenseunicorns/uds-runtime/issues/278)) ([432002f](https://github.com/defenseunicorns/uds-runtime/commit/432002f19d041e371170d886d1124ae6bb4c2c7b))
* **ui:** remove breadcrumb ([#132](https://github.com/defenseunicorns/uds-runtime/issues/132)) ([982c278](https://github.com/defenseunicorns/uds-runtime/commit/982c278cba5226515f9aef5942e5fcfa7bb3e76e))
* update config sidebar icon ([9321de2](https://github.com/defenseunicorns/uds-runtime/commit/9321de2f6ef67f3267c639e22657536e510a4039))
* update pullPolicy ([#220](https://github.com/defenseunicorns/uds-runtime/issues/220)) ([4320468](https://github.com/defenseunicorns/uds-runtime/commit/4320468956c1e44e425e0492bf1faaa1d1b44d20))
* update test iac ([#223](https://github.com/defenseunicorns/uds-runtime/issues/223)) ([9f21f91](https://github.com/defenseunicorns/uds-runtime/commit/9f21f91b5bd8691efa6a9824c5a3c13eacd625e9))
* update zarf ([#99](https://github.com/defenseunicorns/uds-runtime/issues/99)) ([d7c0bee](https://github.com/defenseunicorns/uds-runtime/commit/d7c0bee206272b7fff36e545725145eea315187d))
* updating README ([#204](https://github.com/defenseunicorns/uds-runtime/issues/204)) ([69aab8e](https://github.com/defenseunicorns/uds-runtime/commit/69aab8e21eb3b76318f5f52da8458f8597370d8a))

## [0.3.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.3.0...v0.3.0) (2024-09-06)


### Features

* add missing routes & route filter ([4235c60](https://github.com/defenseunicorns/uds-runtime/commit/4235c60b5243a606c3c02129c3508655cfb6d51a))
* add pod disruption budgets view ([#79](https://github.com/defenseunicorns/uds-runtime/issues/79)) ([a2f17b3](https://github.com/defenseunicorns/uds-runtime/commit/a2f17b39cd0b0304c5bfb9e1cea9b2320f62ef76)), closes [#48](https://github.com/defenseunicorns/uds-runtime/issues/48)
* add sparse streams, compression, single resource api endpoints ([#28](https://github.com/defenseunicorns/uds-runtime/issues/28)) ([c5c4c9c](https://github.com/defenseunicorns/uds-runtime/commit/c5c4c9c5e98cd81e4ea5a0e295983ced1ce05832))
* add storageclasses view ([#83](https://github.com/defenseunicorns/uds-runtime/issues/83)) ([941fff2](https://github.com/defenseunicorns/uds-runtime/commit/941fff2715ee85d2dfda0fb26fbb59fc10f2708f)), closes [#60](https://github.com/defenseunicorns/uds-runtime/issues/60)
* add UDS Exemptions view ([ac0d43e](https://github.com/defenseunicorns/uds-runtime/commit/ac0d43ec947d47341aa4e233aebeefe2e51aeb78))
* add uds package view ([7be0628](https://github.com/defenseunicorns/uds-runtime/commit/7be06284014e0f4b0337c345873b1e0951c09951))
* adding adr for data-test ids ([#200](https://github.com/defenseunicorns/uds-runtime/issues/200)) ([4936722](https://github.com/defenseunicorns/uds-runtime/commit/4936722b96202d97294049c4bc839f3a2231d083))
* adding cluster ops limit ranges view ([#78](https://github.com/defenseunicorns/uds-runtime/issues/78)) ([6907151](https://github.com/defenseunicorns/uds-runtime/commit/69071515bdefa8d50eb04465e29b96b1fa6b2c5f)), closes [#47](https://github.com/defenseunicorns/uds-runtime/issues/47)
* adding drawer e2e tests ([#164](https://github.com/defenseunicorns/uds-runtime/issues/164)) ([f19ca04](https://github.com/defenseunicorns/uds-runtime/commit/f19ca04c90dd03f9b96639beb26e787f3f3923f6))
* adding e2e tests for DataTable ([#188](https://github.com/defenseunicorns/uds-runtime/issues/188)) ([a7c6256](https://github.com/defenseunicorns/uds-runtime/commit/a7c625694fa088103c2a0ee7c7b3f75281ad19bb))
* adding e2e tests for search and dropdown selections ([#193](https://github.com/defenseunicorns/uds-runtime/issues/193)) ([e1a6c20](https://github.com/defenseunicorns/uds-runtime/commit/e1a6c207ea466853003886ae77c7bfcdd6de1995))
* adds pods to pvc view, make swagger more discoverable ([#271](https://github.com/defenseunicorns/uds-runtime/issues/271)) ([d6e5eba](https://github.com/defenseunicorns/uds-runtime/commit/d6e5eba45e4cb5649977cfdbed78cf65c3c1eb23))
* adds pv view ([#71](https://github.com/defenseunicorns/uds-runtime/issues/71)) ([cd013d5](https://github.com/defenseunicorns/uds-runtime/commit/cd013d5271e049310b2d3a6aaf9a0ae2a5b5f5dc))
* **api:** add fields query parameter for fine-grained data filtering ([#94](https://github.com/defenseunicorns/uds-runtime/issues/94)) ([0021790](https://github.com/defenseunicorns/uds-runtime/commit/002179074a4b2ba90eb3e7b62f1e355c6c0717f4))
* **api:** add remaining k8s resource endpoints ([7fa24ee](https://github.com/defenseunicorns/uds-runtime/commit/7fa24ee72ea7ba9ee38f079aa5ed4863f497650b))
* **api:** debounce SSE messages ([#100](https://github.com/defenseunicorns/uds-runtime/issues/100)) ([27d7433](https://github.com/defenseunicorns/uds-runtime/commit/27d7433d37bbae977bc9f0269f717056fdcb0d6c))
* **api:** strip annotations from sparse resources ([#250](https://github.com/defenseunicorns/uds-runtime/issues/250)) ([f43533a](https://github.com/defenseunicorns/uds-runtime/commit/f43533aa75c1f2c8b697b87bab8feaa8147b9f14))
* application views + dashboard ([#144](https://github.com/defenseunicorns/uds-runtime/issues/144)) ([a4c5110](https://github.com/defenseunicorns/uds-runtime/commit/a4c5110025241d38d241f17e10c236b53e00c313))
* **ci:** add nightly releases ([#162](https://github.com/defenseunicorns/uds-runtime/issues/162)) ([d523a64](https://github.com/defenseunicorns/uds-runtime/commit/d523a647d54c144455f4a4808a4770ae05ac91b8))
* enables uds-runtime releases ([#93](https://github.com/defenseunicorns/uds-runtime/issues/93)) ([6af3a95](https://github.com/defenseunicorns/uds-runtime/commit/6af3a95e574f95f1df8ad39a696b311f2c63d16e))
* export pepr data stream ([#24](https://github.com/defenseunicorns/uds-runtime/issues/24)) ([b05976f](https://github.com/defenseunicorns/uds-runtime/commit/b05976fad98a3eaa543f21db8e4f24a3b99805de)), closes [#23](https://github.com/defenseunicorns/uds-runtime/issues/23)
* local mode token protection ([#245](https://github.com/defenseunicorns/uds-runtime/issues/245)) ([d356940](https://github.com/defenseunicorns/uds-runtime/commit/d356940af5e63244381ab408eca854d0ea669761))
* make root URL the overview page ([#212](https://github.com/defenseunicorns/uds-runtime/issues/212)) ([d273a1e](https://github.com/defenseunicorns/uds-runtime/commit/d273a1e8e3ed842e7cb8d1583389fe0f1e25a102))
* swagger ([#114](https://github.com/defenseunicorns/uds-runtime/issues/114)) ([c4016ca](https://github.com/defenseunicorns/uds-runtime/commit/c4016caaef1d1e5e0b6114d351521b0f80e99663))
* **ui:** 112 stylized and colorized yaml view ([#121](https://github.com/defenseunicorns/uds-runtime/issues/121)) ([1ea6532](https://github.com/defenseunicorns/uds-runtime/commit/1ea65325b1e9bcec758e9aed091e186f9d3c257b)), closes [#112](https://github.com/defenseunicorns/uds-runtime/issues/112)
* **ui:** add cluster ops hpa k8s view ([#97](https://github.com/defenseunicorns/uds-runtime/issues/97)) ([a55696d](https://github.com/defenseunicorns/uds-runtime/commit/a55696defbdf7215f752fc36a8d666b6995476f0)), closes [#46](https://github.com/defenseunicorns/uds-runtime/issues/46)
* **ui:** add cluster-ops/mutating-webhooks view ([#42](https://github.com/defenseunicorns/uds-runtime/issues/42)) ([b29c434](https://github.com/defenseunicorns/uds-runtime/commit/b29c434e0c5300194d57b9b831b0528ca3ffd743)), closes [#32](https://github.com/defenseunicorns/uds-runtime/issues/32)
* **ui:** add cronjobs view ([9a80579](https://github.com/defenseunicorns/uds-runtime/commit/9a80579fdcfa14af2426314dfdfb11f89089b72c))
* **ui:** add header to tables ([#106](https://github.com/defenseunicorns/uds-runtime/issues/106)) ([19cdcfe](https://github.com/defenseunicorns/uds-runtime/commit/19cdcfe4ef3c602bada5edc1e560dec311bcda60))
* **ui:** add jobs view ([125f635](https://github.com/defenseunicorns/uds-runtime/commit/125f63517ad9dc69993b7b178cd10a287896b7c9))
* **ui:** add nodes view ([f1cf0b0](https://github.com/defenseunicorns/uds-runtime/commit/f1cf0b0725a6deb4fd33eb74086a31de8059c875))
* **ui:** add pepr details ([#107](https://github.com/defenseunicorns/uds-runtime/issues/107)) ([a4ed8cd](https://github.com/defenseunicorns/uds-runtime/commit/a4ed8cd591ab1c691209fd55bd1bb764f3756952))
* **ui:** add pvc view ([#72](https://github.com/defenseunicorns/uds-runtime/issues/72)) ([5580dd4](https://github.com/defenseunicorns/uds-runtime/commit/5580dd4ece5725c248f8369607a67a31f0e20cdf))
* **ui:** add toast for status messages ([#116](https://github.com/defenseunicorns/uds-runtime/issues/116)) ([983ae6a](https://github.com/defenseunicorns/uds-runtime/commit/983ae6af3b9f9b55d29926b178ebcc228bb7e83f)), closes [#115](https://github.com/defenseunicorns/uds-runtime/issues/115)
* **ui:** adding 404 ([#98](https://github.com/defenseunicorns/uds-runtime/issues/98)) ([24b24b0](https://github.com/defenseunicorns/uds-runtime/commit/24b24b0516cbe563b4e6e7e2ebe7495173c4da5b)), closes [#38](https://github.com/defenseunicorns/uds-runtime/issues/38)
* **ui:** adding cluster ops priority classes view ([#62](https://github.com/defenseunicorns/uds-runtime/issues/62)) ([f52ea3d](https://github.com/defenseunicorns/uds-runtime/commit/f52ea3d47aabd868507e563885fd9133b03bd0af))
* **ui:** adding cluster ops resource quotas view ([#81](https://github.com/defenseunicorns/uds-runtime/issues/81)) ([5370bf9](https://github.com/defenseunicorns/uds-runtime/commit/5370bf98e6e7f65987933de13de863a7a0dd3cd4))
* **ui:** adding cluster ops runtime classes view ([#68](https://github.com/defenseunicorns/uds-runtime/issues/68)) ([084fc83](https://github.com/defenseunicorns/uds-runtime/commit/084fc831fb75a6a0bd70427888cea5f2bd6f57fa))
* **ui:** adding status color mapping for k8s ([#264](https://github.com/defenseunicorns/uds-runtime/issues/264)) ([15fc75f](https://github.com/defenseunicorns/uds-runtime/commit/15fc75f35c529c70ffd3cd0e53a22802ac1691e8))
* **ui:** adds validatingwebhooks view ([#110](https://github.com/defenseunicorns/uds-runtime/issues/110)) ([cad4f35](https://github.com/defenseunicorns/uds-runtime/commit/cad4f3508fe301532a71e4bd04986f7113df406d))
* **ui:** endpoints view ([#80](https://github.com/defenseunicorns/uds-runtime/issues/80)) ([e9bf61e](https://github.com/defenseunicorns/uds-runtime/commit/e9bf61e3cb284769e6402a4a69e97bf99dd8e347))
* **ui:** format age with more detail ([#111](https://github.com/defenseunicorns/uds-runtime/issues/111)) ([98916d6](https://github.com/defenseunicorns/uds-runtime/commit/98916d6328d9c131aa5394058eaeb77b10ec8bf3))
* **ui:** make package endpoints links ([#136](https://github.com/defenseunicorns/uds-runtime/issues/136)) ([572f385](https://github.com/defenseunicorns/uds-runtime/commit/572f3859b0ba5901182fe335eb5e9ac679635716))
* **ui:** networkpolicies view ([#75](https://github.com/defenseunicorns/uds-runtime/issues/75)) ([d04d50b](https://github.com/defenseunicorns/uds-runtime/commit/d04d50bab0c23ba15952a7977a3046e280e39f4f))
* **ui:** networks-services view ([#61](https://github.com/defenseunicorns/uds-runtime/issues/61)) ([1d0b68a](https://github.com/defenseunicorns/uds-runtime/commit/1d0b68aa39ad59e79107ddb61d650ae427cc9598))
* **ui:** resource detail view ([#95](https://github.com/defenseunicorns/uds-runtime/issues/95)) ([34e7e6c](https://github.com/defenseunicorns/uds-runtime/commit/34e7e6ca168fd3baa7b12dae7d4edd66f6343a8a)), closes [#34](https://github.com/defenseunicorns/uds-runtime/issues/34)
* **ui:** show results even if not filtering ([#131](https://github.com/defenseunicorns/uds-runtime/issues/131)) ([4603b65](https://github.com/defenseunicorns/uds-runtime/commit/4603b656b41e3a94883a4694a4a881f27512424b))
* **ui:** updating drawer header ([#142](https://github.com/defenseunicorns/uds-runtime/issues/142)) ([98dad19](https://github.com/defenseunicorns/uds-runtime/commit/98dad199800d06a0fb2690673dc35785784400a0)), closes [#102](https://github.com/defenseunicorns/uds-runtime/issues/102)
* **ui:** virtualservices view ([#73](https://github.com/defenseunicorns/uds-runtime/issues/73)) ([4d01cb8](https://github.com/defenseunicorns/uds-runtime/commit/4d01cb8db3125c747f5e26b835edc9fd0cf61131))


### Bug Fixes

* add missing routes ([#6](https://github.com/defenseunicorns/uds-runtime/issues/6)) ([3e61370](https://github.com/defenseunicorns/uds-runtime/commit/3e61370bb0443beefcb7b69f298570baca8837b2))
* add tasks file to release please config and add docker login ([#139](https://github.com/defenseunicorns/uds-runtime/issues/139)) ([f61680f](https://github.com/defenseunicorns/uds-runtime/commit/f61680f2941bb8e092b8da69d5c9cbcdcbb12d43))
* api token auth details view ([#276](https://github.com/defenseunicorns/uds-runtime/issues/276)) ([871cdf5](https://github.com/defenseunicorns/uds-runtime/commit/871cdf5f9561f1241a77f4b6271d94cb0dfabc7b))
* **api:** dockerfile for releasing go binaries ([#258](https://github.com/defenseunicorns/uds-runtime/issues/258)) ([e35987c](https://github.com/defenseunicorns/uds-runtime/commit/e35987c01acc82fbab2023e691483a281e1673ac))
* **ci:** more permissions for release ([#135](https://github.com/defenseunicorns/uds-runtime/issues/135)) ([3695630](https://github.com/defenseunicorns/uds-runtime/commit/36956305edfbcf2fa94befb20d900d1c7503d6b3))
* **ci:** swag init task output directory ([#248](https://github.com/defenseunicorns/uds-runtime/issues/248)) ([0d65c2e](https://github.com/defenseunicorns/uds-runtime/commit/0d65c2e97c8202729bdae9519892e47455e66ddf))
* correct nightly working dir ([#196](https://github.com/defenseunicorns/uds-runtime/issues/196)) ([f61dea2](https://github.com/defenseunicorns/uds-runtime/commit/f61dea22ee20a41579d31dbd583e6e63ae7a8975))
* docker container run failure ([#267](https://github.com/defenseunicorns/uds-runtime/issues/267)) ([573a605](https://github.com/defenseunicorns/uds-runtime/commit/573a6053820cc8f1ff55589e38b45c44f59788ab))
* empty endpoints are shown as Pending ([#158](https://github.com/defenseunicorns/uds-runtime/issues/158)) ([08b96ad](https://github.com/defenseunicorns/uds-runtime/commit/08b96ad9334f76af99856e0c9e9198aea2614c8e))
* enable release-please PRs to run workflows ([#146](https://github.com/defenseunicorns/uds-runtime/issues/146)) ([10d7858](https://github.com/defenseunicorns/uds-runtime/commit/10d7858adcffd2f977c6101a25b81696dd266f00))
* ensure nightly artifacts use proper tag ([#190](https://github.com/defenseunicorns/uds-runtime/issues/190)) ([bb44ae3](https://github.com/defenseunicorns/uds-runtime/commit/bb44ae33ddd2d220cb14410b17f828823ad03b4d))
* fixing bug for issue [#147](https://github.com/defenseunicorns/uds-runtime/issues/147) ([#149](https://github.com/defenseunicorns/uds-runtime/issues/149)) ([4d6122d](https://github.com/defenseunicorns/uds-runtime/commit/4d6122dede264208238c2e3e26747db39489a58c))
* fixing FOUC issue ([#236](https://github.com/defenseunicorns/uds-runtime/issues/236)) ([99886ea](https://github.com/defenseunicorns/uds-runtime/commit/99886eab9806faed4c269f3bea4d99ca67248d36))
* fixing issue with chart not being destroyed onDestroy ([#243](https://github.com/defenseunicorns/uds-runtime/issues/243)) ([4bd9baf](https://github.com/defenseunicorns/uds-runtime/commit/4bd9bafc67cd39570aa94b1dccf9072933cb360e))
* fixing tooltip behavior ([#26](https://github.com/defenseunicorns/uds-runtime/issues/26)) ([dd8b092](https://github.com/defenseunicorns/uds-runtime/commit/dd8b0928be7e0aa4b9c656a8ac83e8656fa4f530)), closes [#25](https://github.com/defenseunicorns/uds-runtime/issues/25)
* give ssm-user more permissions ([#228](https://github.com/defenseunicorns/uds-runtime/issues/228)) ([a55626c](https://github.com/defenseunicorns/uds-runtime/commit/a55626c7a0645de7bd3c6743c27f6113cad49343))
* nightly release ([#260](https://github.com/defenseunicorns/uds-runtime/issues/260)) ([a99bc23](https://github.com/defenseunicorns/uds-runtime/commit/a99bc23473a31ee7b4761bf50b0b162c388eed51))
* nightly release artifact fix ([#259](https://github.com/defenseunicorns/uds-runtime/issues/259)) ([648600d](https://github.com/defenseunicorns/uds-runtime/commit/648600dcf811ef2cf060389b9682a8ee8ff82dbc))
* pod status color ([#176](https://github.com/defenseunicorns/uds-runtime/issues/176)) ([2fe0c99](https://github.com/defenseunicorns/uds-runtime/commit/2fe0c99e51dea34bbec0e5383ef2890e563d9f8d))
* remove alpha from release ([#145](https://github.com/defenseunicorns/uds-runtime/issues/145)) ([70344d8](https://github.com/defenseunicorns/uds-runtime/commit/70344d8a07a33df57d90cc848974e08fde327ac7))
* remove breadcrumb links, closes [#9](https://github.com/defenseunicorns/uds-runtime/issues/9) ([e2d8cfe](https://github.com/defenseunicorns/uds-runtime/commit/e2d8cfecb5c1ce1c863435eabf42937f57100401))
* removes empty docs page, adds favicon ([#234](https://github.com/defenseunicorns/uds-runtime/issues/234)) ([db3212c](https://github.com/defenseunicorns/uds-runtime/commit/db3212c9ad0a33df1576ec9866917eba20b0ccf4))
* sidebar expand toggle icon rotation ([e1aa7cb](https://github.com/defenseunicorns/uds-runtime/commit/e1aa7cbe97faae1044aa895ee338427ab3b55180))
* sidebar ux issues ([589a970](https://github.com/defenseunicorns/uds-runtime/commit/589a9700794e10332e953f63682c4b3d970797cd))
* turn off autocomplete for datatable ([#22](https://github.com/defenseunicorns/uds-runtime/issues/22)) ([f71eccf](https://github.com/defenseunicorns/uds-runtime/commit/f71eccfbe084f2c9808b970fff16c3c8d12c4ad1)), closes [#21](https://github.com/defenseunicorns/uds-runtime/issues/21)
* **ui:** deployments table not showing accurate available count ([#192](https://github.com/defenseunicorns/uds-runtime/issues/192)) ([5b0f394](https://github.com/defenseunicorns/uds-runtime/commit/5b0f394ef80ea6286a017093851b4f42ae16009c))
* **ui:** details reactivity on mouse/keyboard navigate ([#130](https://github.com/defenseunicorns/uds-runtime/issues/130)) ([02da735](https://github.com/defenseunicorns/uds-runtime/commit/02da735dd73c56d01ac1ea67baa09f8192c8abed))
* **ui:** fixing issue with graph menu color ([#195](https://github.com/defenseunicorns/uds-runtime/issues/195)) ([b10f95f](https://github.com/defenseunicorns/uds-runtime/commit/b10f95f5595a4341b93356f671dca6b818abc31b))
* **ui:** pepr monitor unsubscriber ([#67](https://github.com/defenseunicorns/uds-runtime/issues/67)) ([2995497](https://github.com/defenseunicorns/uds-runtime/commit/2995497a02fb1b6b6a261aad614c4b4bfe84bc98)), closes [#37](https://github.com/defenseunicorns/uds-runtime/issues/37)
* **ui:** prevent error on null fields in DataTable component ([#43](https://github.com/defenseunicorns/uds-runtime/issues/43)) ([d51636e](https://github.com/defenseunicorns/uds-runtime/commit/d51636e5ee9154c1863d06222d927c6e61f26cb1))
* **ui:** restart and node columns in pod table ([#194](https://github.com/defenseunicorns/uds-runtime/issues/194)) ([88a41c4](https://github.com/defenseunicorns/uds-runtime/commit/88a41c4fc4838945361d8d3117086664273b3a50))
* use docker buildx in github actions ([#133](https://github.com/defenseunicorns/uds-runtime/issues/133)) ([d52cb6c](https://github.com/defenseunicorns/uds-runtime/commit/d52cb6cebb3eaab211d0087726b7c3f2a4d18b4b))
* use middleware for compression ([#86](https://github.com/defenseunicorns/uds-runtime/issues/86)) ([f16a18c](https://github.com/defenseunicorns/uds-runtime/commit/f16a18c755f67cad6ffdd3a21d1708af4f99813e)), closes [#85](https://github.com/defenseunicorns/uds-runtime/issues/85)


### Miscellaneous

* add contributing guide ([#17](https://github.com/defenseunicorns/uds-runtime/issues/17)) ([f03c08f](https://github.com/defenseunicorns/uds-runtime/commit/f03c08f73abccdc7dd69744fdbd5a25ad4d6371f)), closes [#18](https://github.com/defenseunicorns/uds-runtime/issues/18)
* add empty release please manifest to bootstrap ([#120](https://github.com/defenseunicorns/uds-runtime/issues/120)) ([139fe0a](https://github.com/defenseunicorns/uds-runtime/commit/139fe0a0dc32d04c6728b0de3f82a102ef7638f9))
* add env for gh token release workflow ([#286](https://github.com/defenseunicorns/uds-runtime/issues/286)) ([0398329](https://github.com/defenseunicorns/uds-runtime/commit/039832932923eb4c62f21dca01adf3b788e393f6))
* add go binaries as artifacts of releases ([#240](https://github.com/defenseunicorns/uds-runtime/issues/240)) ([422deea](https://github.com/defenseunicorns/uds-runtime/commit/422deea7ef61f48e228c397aeb4489d5026c09df))
* add initial ci ([#11](https://github.com/defenseunicorns/uds-runtime/issues/11)) ([57248fb](https://github.com/defenseunicorns/uds-runtime/commit/57248fb1c2bfe03131b7bc8ce90f6ef419078838))
* add release ci ([#118](https://github.com/defenseunicorns/uds-runtime/issues/118)) ([70c0984](https://github.com/defenseunicorns/uds-runtime/commit/70c09845408ec8e4e62916bef7ae5d4cd4ebc87e))
* add resource store tests ([#74](https://github.com/defenseunicorns/uds-runtime/issues/74)) ([33ad7bc](https://github.com/defenseunicorns/uds-runtime/commit/33ad7bcee1e2231ab2570865932bd29b6194d997))
* add test infra for dummy cluster ([#157](https://github.com/defenseunicorns/uds-runtime/issues/157)) ([2ba9e26](https://github.com/defenseunicorns/uds-runtime/commit/2ba9e261255c57b1631249e3aa1864eb5c7b34f1))
* add uds-cli taks & zarf schema configs for vscode ([344fd7e](https://github.com/defenseunicorns/uds-runtime/commit/344fd7e6ab217e7355e58073e07a8972986249d9))
* add vscode go debug config ([c811f47](https://github.com/defenseunicorns/uds-runtime/commit/c811f47f804cd21d08300d4cc906bc5e6c9f8536))
* adding all routes for navigation ([#16](https://github.com/defenseunicorns/uds-runtime/issues/16)) ([eb518c3](https://github.com/defenseunicorns/uds-runtime/commit/eb518c39a9d45e74fc76f15abd2346ae304478a7)), closes [#15](https://github.com/defenseunicorns/uds-runtime/issues/15)
* adding contributing instructions for front end testing locators ([#242](https://github.com/defenseunicorns/uds-runtime/issues/242)) ([4a253cf](https://github.com/defenseunicorns/uds-runtime/commit/4a253cf78039cc0ce450eb5a00baa6e31e1723b0))
* adds basic e2e tests ([#175](https://github.com/defenseunicorns/uds-runtime/issues/175)) ([770c4a3](https://github.com/defenseunicorns/uds-runtime/commit/770c4a32a0ac1aae4e58cc8af89c6d998513a688))
* adds navbar tests ([#189](https://github.com/defenseunicorns/uds-runtime/issues/189)) ([247afc7](https://github.com/defenseunicorns/uds-runtime/commit/247afc7813fe560648453764f1b891845483a9c4))
* adds ssm to runtime-canary ([#221](https://github.com/defenseunicorns/uds-runtime/issues/221)) ([13d2350](https://github.com/defenseunicorns/uds-runtime/commit/13d2350e0e6f2c2ce6a6d06c9c3e54855a762990))
* api testing adr ([#198](https://github.com/defenseunicorns/uds-runtime/issues/198)) ([f50bd77](https://github.com/defenseunicorns/uds-runtime/commit/f50bd773740300745e19f7821a1a9143a3c12679))
* api token auth documentation ([#269](https://github.com/defenseunicorns/uds-runtime/issues/269)) ([20cbd02](https://github.com/defenseunicorns/uds-runtime/commit/20cbd0265f56b5dbb17c1ca3538cf9f13f1a4987))
* **api:** add api tests ([#231](https://github.com/defenseunicorns/uds-runtime/issues/231)) ([3f9eb2c](https://github.com/defenseunicorns/uds-runtime/commit/3f9eb2caa6cbca61bd7b99b73c482868cb061a20))
* **api:** add field selector api tests ([#262](https://github.com/defenseunicorns/uds-runtime/issues/262)) ([6ee6720](https://github.com/defenseunicorns/uds-runtime/commit/6ee672046101cd965da57111c091d8233489c6ee))
* bump spdx header to 2024-Present ([2102add](https://github.com/defenseunicorns/uds-runtime/commit/2102addd1864b67a82b59e4ac11c4f7fb51f0824))
* **ci:** add swagger docs gen to CI ([#167](https://github.com/defenseunicorns/uds-runtime/issues/167)) ([aed5b83](https://github.com/defenseunicorns/uds-runtime/commit/aed5b836288fb5592b254302355d05cb7e7ee3bc))
* **ci:** add type checking ([#187](https://github.com/defenseunicorns/uds-runtime/issues/187)) ([3357983](https://github.com/defenseunicorns/uds-runtime/commit/3357983a961bc565aca55464dc19b8c6e346cd81))
* codeowners ([#174](https://github.com/defenseunicorns/uds-runtime/issues/174)) ([b4cabcf](https://github.com/defenseunicorns/uds-runtime/commit/b4cabcf34c90d40fc2c6467932ac42c2f34e9fc6))
* configure renovate ([#209](https://github.com/defenseunicorns/uds-runtime/issues/209)) ([4712c6f](https://github.com/defenseunicorns/uds-runtime/commit/4712c6f71d04b37e6f55dc57b2b910ba3c86429c))
* **deps:** update all dependencies ([#13](https://github.com/defenseunicorns/uds-runtime/issues/13)) ([ec95b7e](https://github.com/defenseunicorns/uds-runtime/commit/ec95b7e67e07246489220e1d6783ad9ef195ca74))
* **deps:** update dependency jsdom to v25 ([#247](https://github.com/defenseunicorns/uds-runtime/issues/247)) ([772930b](https://github.com/defenseunicorns/uds-runtime/commit/772930ba4677fcee0f4b242c9f0f4a9f66960a66))
* **deps:** update dependency kubernetes-fluent-client to v3 ([#218](https://github.com/defenseunicorns/uds-runtime/issues/218)) ([245c256](https://github.com/defenseunicorns/uds-runtime/commit/245c25633a5019dc8a52649d2830173c1e95da5c))
* **deps:** update dependency kubernetes-fluent-client to v3.0.2 ([#256](https://github.com/defenseunicorns/uds-runtime/issues/256)) ([4c075a8](https://github.com/defenseunicorns/uds-runtime/commit/4c075a8e108c71c31b0e0354a04272dbd1efd6e2))
* **deps:** update github actions ([#213](https://github.com/defenseunicorns/uds-runtime/issues/213)) ([b51692f](https://github.com/defenseunicorns/uds-runtime/commit/b51692fff7c6f9f7f423ac0e3fbf8188cedd51bd))
* **deps:** update github actions ([#227](https://github.com/defenseunicorns/uds-runtime/issues/227)) ([22b18b0](https://github.com/defenseunicorns/uds-runtime/commit/22b18b0c0d1a42e3f30d5d12849cdc478b500482))
* **deps:** update github actions ([#257](https://github.com/defenseunicorns/uds-runtime/issues/257)) ([6c366cc](https://github.com/defenseunicorns/uds-runtime/commit/6c366ccbe0264a0c95362051ec072c59320c489e))
* **deps:** update kubernetes packages to v0.31.0 ([#215](https://github.com/defenseunicorns/uds-runtime/issues/215)) ([9875e62](https://github.com/defenseunicorns/uds-runtime/commit/9875e62315d17814c861888f55c4d2f88ce0d92a))
* **deps:** update module github.com/charmbracelet/lipgloss to v0.13.0 ([#232](https://github.com/defenseunicorns/uds-runtime/issues/232)) ([1cf315a](https://github.com/defenseunicorns/uds-runtime/commit/1cf315af910c026469f331fb338a03a012270716))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.2 ([#216](https://github.com/defenseunicorns/uds-runtime/issues/216)) ([21c7457](https://github.com/defenseunicorns/uds-runtime/commit/21c745736397f461e9611da42f8b24ff1c89c493))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.3 ([#244](https://github.com/defenseunicorns/uds-runtime/issues/244)) ([adf6d82](https://github.com/defenseunicorns/uds-runtime/commit/adf6d82d175e139f84ec93f7e77f824a2537df5f))
* **deps:** update module github.com/zarf-dev/zarf to v0.39.0 ([#283](https://github.com/defenseunicorns/uds-runtime/issues/283)) ([e3719a2](https://github.com/defenseunicorns/uds-runtime/commit/e3719a2721fe48fe1f6bf208cecf996da6ca6cd4))
* ensure lang=ts for all svelte script tags ([b40ecbb](https://github.com/defenseunicorns/uds-runtime/commit/b40ecbba365fecc7eacd312ee37765a3a297f3b6))
* fix lint errors ([68df75f](https://github.com/defenseunicorns/uds-runtime/commit/68df75fc7f8c6d3693d85c30663fff2cb9e9fe2d))
* go unit tests ([#191](https://github.com/defenseunicorns/uds-runtime/issues/191)) ([b850316](https://github.com/defenseunicorns/uds-runtime/commit/b85031664c6e2cc68cf28379961d30ccf62fabd0))
* install Istio CRDs for minimal e2e tests ([#168](https://github.com/defenseunicorns/uds-runtime/issues/168)) ([8832dcd](https://github.com/defenseunicorns/uds-runtime/commit/8832dcd203cfc3086fa2280d76e4dd901b74f2c9))
* lint ([ef14348](https://github.com/defenseunicorns/uds-runtime/commit/ef1434877c8feb39ab3b1601c02e24269abd43f8))
* **main:** release 0.1.0 ([#143](https://github.com/defenseunicorns/uds-runtime/issues/143)) ([ef1fa4c](https://github.com/defenseunicorns/uds-runtime/commit/ef1fa4ca67007db38984daa42cc0dbef208cb9f9))
* **main:** release 0.1.0-alpha.3 ([#134](https://github.com/defenseunicorns/uds-runtime/issues/134)) ([410ce9b](https://github.com/defenseunicorns/uds-runtime/commit/410ce9bdf54acbef0f2972e78987d05787196063))
* **main:** release 0.1.0-alpha.4 ([#137](https://github.com/defenseunicorns/uds-runtime/issues/137)) ([2907c75](https://github.com/defenseunicorns/uds-runtime/commit/2907c75d313e985fa2f3dd1c7a7a09c06d78efad))
* **main:** release 0.1.0-alpha.5 ([#138](https://github.com/defenseunicorns/uds-runtime/issues/138)) ([12ff0c1](https://github.com/defenseunicorns/uds-runtime/commit/12ff0c138b606bc5abbc4201061c2e341b5deb1f))
* **main:** release 0.2.0 ([#150](https://github.com/defenseunicorns/uds-runtime/issues/150)) ([fb2456d](https://github.com/defenseunicorns/uds-runtime/commit/fb2456d1f4e87e2799836216a2f45e41783f7f59))
* **main:** release 0.3.0 ([#224](https://github.com/defenseunicorns/uds-runtime/issues/224)) ([2d6d21c](https://github.com/defenseunicorns/uds-runtime/commit/2d6d21c5ad762d58ad0ad0be597b0cd82be77b19))
* make uds package host configurable for nightly ([#201](https://github.com/defenseunicorns/uds-runtime/issues/201)) ([bb34173](https://github.com/defenseunicorns/uds-runtime/commit/bb34173b6418d17c9aedc4876941a80927a7645e))
* refactor controller to controlled by ([#122](https://github.com/defenseunicorns/uds-runtime/issues/122)) ([839502d](https://github.com/defenseunicorns/uds-runtime/commit/839502df42618c09b6d28451248bf70aa73453a1))
* remove message.Fatal for zarf upgrade ([8438b37](https://github.com/defenseunicorns/uds-runtime/commit/8438b37532367f160487c77b9dbdce500d6b9e8e))
* **setting-up-release-please:** release 0.1.0-alpha.2 ([#127](https://github.com/defenseunicorns/uds-runtime/issues/127)) ([52bfd99](https://github.com/defenseunicorns/uds-runtime/commit/52bfd99dc58ef17d4bd3b14426f07afc13372f41))
* some go unit tests ([#119](https://github.com/defenseunicorns/uds-runtime/issues/119)) ([a0a7433](https://github.com/defenseunicorns/uds-runtime/commit/a0a74335249ac076e6ac40872abaeeffc0b2e5c6))
* speed up e2e tests with minimal cluster ([#155](https://github.com/defenseunicorns/uds-runtime/issues/155)) ([b8aae8a](https://github.com/defenseunicorns/uds-runtime/commit/b8aae8a9fb31be7be57394817769feeb19ee5747))
* swap to admin gateway ([#124](https://github.com/defenseunicorns/uds-runtime/issues/124)) ([776927b](https://github.com/defenseunicorns/uds-runtime/commit/776927baf5e255b5e80e490240318a0a6a3e34dd))
* tag uds cli install for renovate ([#226](https://github.com/defenseunicorns/uds-runtime/issues/226)) ([e3f5f71](https://github.com/defenseunicorns/uds-runtime/commit/e3f5f71c36a1c7fb3d1823e359ab93e0169d7c91))
* uds-engine -&gt; uds-runtime ([e553195](https://github.com/defenseunicorns/uds-runtime/commit/e553195ffc0cef8dd0b4fca0b6cf350301ecdcfc))
* **ui:** redirect to auth page when unauthenticated ([#278](https://github.com/defenseunicorns/uds-runtime/issues/278)) ([432002f](https://github.com/defenseunicorns/uds-runtime/commit/432002f19d041e371170d886d1124ae6bb4c2c7b))
* **ui:** remove breadcrumb ([#132](https://github.com/defenseunicorns/uds-runtime/issues/132)) ([982c278](https://github.com/defenseunicorns/uds-runtime/commit/982c278cba5226515f9aef5942e5fcfa7bb3e76e))
* update config sidebar icon ([9321de2](https://github.com/defenseunicorns/uds-runtime/commit/9321de2f6ef67f3267c639e22657536e510a4039))
* update pullPolicy ([#220](https://github.com/defenseunicorns/uds-runtime/issues/220)) ([4320468](https://github.com/defenseunicorns/uds-runtime/commit/4320468956c1e44e425e0492bf1faaa1d1b44d20))
* update test iac ([#223](https://github.com/defenseunicorns/uds-runtime/issues/223)) ([9f21f91](https://github.com/defenseunicorns/uds-runtime/commit/9f21f91b5bd8691efa6a9824c5a3c13eacd625e9))
* update zarf ([#99](https://github.com/defenseunicorns/uds-runtime/issues/99)) ([d7c0bee](https://github.com/defenseunicorns/uds-runtime/commit/d7c0bee206272b7fff36e545725145eea315187d))
* updating README ([#204](https://github.com/defenseunicorns/uds-runtime/issues/204)) ([69aab8e](https://github.com/defenseunicorns/uds-runtime/commit/69aab8e21eb3b76318f5f52da8458f8597370d8a))

## [0.3.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.2.0...v0.3.0) (2024-09-06)


### Features

* adding adr for data-test ids ([#200](https://github.com/defenseunicorns/uds-runtime/issues/200)) ([4936722](https://github.com/defenseunicorns/uds-runtime/commit/4936722b96202d97294049c4bc839f3a2231d083))
* adds pods to pvc view, make swagger more discoverable ([#271](https://github.com/defenseunicorns/uds-runtime/issues/271)) ([d6e5eba](https://github.com/defenseunicorns/uds-runtime/commit/d6e5eba45e4cb5649977cfdbed78cf65c3c1eb23))
* **api:** add fields query parameter for fine-grained data filtering ([#94](https://github.com/defenseunicorns/uds-runtime/issues/94)) ([0021790](https://github.com/defenseunicorns/uds-runtime/commit/002179074a4b2ba90eb3e7b62f1e355c6c0717f4))
* **api:** strip annotations from sparse resources ([#250](https://github.com/defenseunicorns/uds-runtime/issues/250)) ([f43533a](https://github.com/defenseunicorns/uds-runtime/commit/f43533aa75c1f2c8b697b87bab8feaa8147b9f14))
* local mode token protection ([#245](https://github.com/defenseunicorns/uds-runtime/issues/245)) ([d356940](https://github.com/defenseunicorns/uds-runtime/commit/d356940af5e63244381ab408eca854d0ea669761))
* **ui:** adding status color mapping for k8s ([#264](https://github.com/defenseunicorns/uds-runtime/issues/264)) ([15fc75f](https://github.com/defenseunicorns/uds-runtime/commit/15fc75f35c529c70ffd3cd0e53a22802ac1691e8))


### Bug Fixes

* api token auth details view ([#276](https://github.com/defenseunicorns/uds-runtime/issues/276)) ([871cdf5](https://github.com/defenseunicorns/uds-runtime/commit/871cdf5f9561f1241a77f4b6271d94cb0dfabc7b))
* **api:** dockerfile for releasing go binaries ([#258](https://github.com/defenseunicorns/uds-runtime/issues/258)) ([e35987c](https://github.com/defenseunicorns/uds-runtime/commit/e35987c01acc82fbab2023e691483a281e1673ac))
* **ci:** swag init task output directory ([#248](https://github.com/defenseunicorns/uds-runtime/issues/248)) ([0d65c2e](https://github.com/defenseunicorns/uds-runtime/commit/0d65c2e97c8202729bdae9519892e47455e66ddf))
* docker container run failure ([#267](https://github.com/defenseunicorns/uds-runtime/issues/267)) ([573a605](https://github.com/defenseunicorns/uds-runtime/commit/573a6053820cc8f1ff55589e38b45c44f59788ab))
* fixing FOUC issue ([#236](https://github.com/defenseunicorns/uds-runtime/issues/236)) ([99886ea](https://github.com/defenseunicorns/uds-runtime/commit/99886eab9806faed4c269f3bea4d99ca67248d36))
* fixing issue with chart not being destroyed onDestroy ([#243](https://github.com/defenseunicorns/uds-runtime/issues/243)) ([4bd9baf](https://github.com/defenseunicorns/uds-runtime/commit/4bd9bafc67cd39570aa94b1dccf9072933cb360e))
* give ssm-user more permissions ([#228](https://github.com/defenseunicorns/uds-runtime/issues/228)) ([a55626c](https://github.com/defenseunicorns/uds-runtime/commit/a55626c7a0645de7bd3c6743c27f6113cad49343))
* nightly release ([#260](https://github.com/defenseunicorns/uds-runtime/issues/260)) ([a99bc23](https://github.com/defenseunicorns/uds-runtime/commit/a99bc23473a31ee7b4761bf50b0b162c388eed51))
* nightly release artifact fix ([#259](https://github.com/defenseunicorns/uds-runtime/issues/259)) ([648600d](https://github.com/defenseunicorns/uds-runtime/commit/648600dcf811ef2cf060389b9682a8ee8ff82dbc))
* removes empty docs page, adds favicon ([#234](https://github.com/defenseunicorns/uds-runtime/issues/234)) ([db3212c](https://github.com/defenseunicorns/uds-runtime/commit/db3212c9ad0a33df1576ec9866917eba20b0ccf4))
* **ui:** fixing issue with graph menu color ([#195](https://github.com/defenseunicorns/uds-runtime/issues/195)) ([b10f95f](https://github.com/defenseunicorns/uds-runtime/commit/b10f95f5595a4341b93356f671dca6b818abc31b))


### Miscellaneous

* add go binaries as artifacts of releases ([#240](https://github.com/defenseunicorns/uds-runtime/issues/240)) ([422deea](https://github.com/defenseunicorns/uds-runtime/commit/422deea7ef61f48e228c397aeb4489d5026c09df))
* adding contributing instructions for front end testing locators ([#242](https://github.com/defenseunicorns/uds-runtime/issues/242)) ([4a253cf](https://github.com/defenseunicorns/uds-runtime/commit/4a253cf78039cc0ce450eb5a00baa6e31e1723b0))
* api token auth documentation ([#269](https://github.com/defenseunicorns/uds-runtime/issues/269)) ([20cbd02](https://github.com/defenseunicorns/uds-runtime/commit/20cbd0265f56b5dbb17c1ca3538cf9f13f1a4987))
* **api:** add api tests ([#231](https://github.com/defenseunicorns/uds-runtime/issues/231)) ([3f9eb2c](https://github.com/defenseunicorns/uds-runtime/commit/3f9eb2caa6cbca61bd7b99b73c482868cb061a20))
* **api:** add field selector api tests ([#262](https://github.com/defenseunicorns/uds-runtime/issues/262)) ([6ee6720](https://github.com/defenseunicorns/uds-runtime/commit/6ee672046101cd965da57111c091d8233489c6ee))
* **deps:** update dependency jsdom to v25 ([#247](https://github.com/defenseunicorns/uds-runtime/issues/247)) ([772930b](https://github.com/defenseunicorns/uds-runtime/commit/772930ba4677fcee0f4b242c9f0f4a9f66960a66))
* **deps:** update dependency kubernetes-fluent-client to v3.0.2 ([#256](https://github.com/defenseunicorns/uds-runtime/issues/256)) ([4c075a8](https://github.com/defenseunicorns/uds-runtime/commit/4c075a8e108c71c31b0e0354a04272dbd1efd6e2))
* **deps:** update github actions ([#213](https://github.com/defenseunicorns/uds-runtime/issues/213)) ([b51692f](https://github.com/defenseunicorns/uds-runtime/commit/b51692fff7c6f9f7f423ac0e3fbf8188cedd51bd))
* **deps:** update github actions ([#227](https://github.com/defenseunicorns/uds-runtime/issues/227)) ([22b18b0](https://github.com/defenseunicorns/uds-runtime/commit/22b18b0c0d1a42e3f30d5d12849cdc478b500482))
* **deps:** update github actions ([#257](https://github.com/defenseunicorns/uds-runtime/issues/257)) ([6c366cc](https://github.com/defenseunicorns/uds-runtime/commit/6c366ccbe0264a0c95362051ec072c59320c489e))
* **deps:** update kubernetes packages to v0.31.0 ([#215](https://github.com/defenseunicorns/uds-runtime/issues/215)) ([9875e62](https://github.com/defenseunicorns/uds-runtime/commit/9875e62315d17814c861888f55c4d2f88ce0d92a))
* **deps:** update module github.com/charmbracelet/lipgloss to v0.13.0 ([#232](https://github.com/defenseunicorns/uds-runtime/issues/232)) ([1cf315a](https://github.com/defenseunicorns/uds-runtime/commit/1cf315af910c026469f331fb338a03a012270716))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.3 ([#244](https://github.com/defenseunicorns/uds-runtime/issues/244)) ([adf6d82](https://github.com/defenseunicorns/uds-runtime/commit/adf6d82d175e139f84ec93f7e77f824a2537df5f))
* **deps:** update module github.com/zarf-dev/zarf to v0.39.0 ([#283](https://github.com/defenseunicorns/uds-runtime/issues/283)) ([e3719a2](https://github.com/defenseunicorns/uds-runtime/commit/e3719a2721fe48fe1f6bf208cecf996da6ca6cd4))
* tag uds cli install for renovate ([#226](https://github.com/defenseunicorns/uds-runtime/issues/226)) ([e3f5f71](https://github.com/defenseunicorns/uds-runtime/commit/e3f5f71c36a1c7fb3d1823e359ab93e0169d7c91))
* **ui:** redirect to auth page when unauthenticated ([#278](https://github.com/defenseunicorns/uds-runtime/issues/278)) ([432002f](https://github.com/defenseunicorns/uds-runtime/commit/432002f19d041e371170d886d1124ae6bb4c2c7b))
* update test iac ([#223](https://github.com/defenseunicorns/uds-runtime/issues/223)) ([9f21f91](https://github.com/defenseunicorns/uds-runtime/commit/9f21f91b5bd8691efa6a9824c5a3c13eacd625e9))

## [0.2.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0...v0.2.0) (2024-08-19)


### Features

* adding drawer e2e tests ([#164](https://github.com/defenseunicorns/uds-runtime/issues/164)) ([f19ca04](https://github.com/defenseunicorns/uds-runtime/commit/f19ca04c90dd03f9b96639beb26e787f3f3923f6))
* adding e2e tests for DataTable ([#188](https://github.com/defenseunicorns/uds-runtime/issues/188)) ([a7c6256](https://github.com/defenseunicorns/uds-runtime/commit/a7c625694fa088103c2a0ee7c7b3f75281ad19bb))
* adding e2e tests for search and dropdown selections ([#193](https://github.com/defenseunicorns/uds-runtime/issues/193)) ([e1a6c20](https://github.com/defenseunicorns/uds-runtime/commit/e1a6c207ea466853003886ae77c7bfcdd6de1995))
* application views + dashboard ([#144](https://github.com/defenseunicorns/uds-runtime/issues/144)) ([a4c5110](https://github.com/defenseunicorns/uds-runtime/commit/a4c5110025241d38d241f17e10c236b53e00c313))
* **ci:** add nightly releases ([#162](https://github.com/defenseunicorns/uds-runtime/issues/162)) ([d523a64](https://github.com/defenseunicorns/uds-runtime/commit/d523a647d54c144455f4a4808a4770ae05ac91b8))
* make root URL the overview page ([#212](https://github.com/defenseunicorns/uds-runtime/issues/212)) ([d273a1e](https://github.com/defenseunicorns/uds-runtime/commit/d273a1e8e3ed842e7cb8d1583389fe0f1e25a102))


### Bug Fixes

* correct nightly working dir ([#196](https://github.com/defenseunicorns/uds-runtime/issues/196)) ([f61dea2](https://github.com/defenseunicorns/uds-runtime/commit/f61dea22ee20a41579d31dbd583e6e63ae7a8975))
* empty endpoints are shown as Pending ([#158](https://github.com/defenseunicorns/uds-runtime/issues/158)) ([08b96ad](https://github.com/defenseunicorns/uds-runtime/commit/08b96ad9334f76af99856e0c9e9198aea2614c8e))
* ensure nightly artifacts use proper tag ([#190](https://github.com/defenseunicorns/uds-runtime/issues/190)) ([bb44ae3](https://github.com/defenseunicorns/uds-runtime/commit/bb44ae33ddd2d220cb14410b17f828823ad03b4d))
* fixing bug for issue [#147](https://github.com/defenseunicorns/uds-runtime/issues/147) ([#149](https://github.com/defenseunicorns/uds-runtime/issues/149)) ([4d6122d](https://github.com/defenseunicorns/uds-runtime/commit/4d6122dede264208238c2e3e26747db39489a58c))
* pod status color ([#176](https://github.com/defenseunicorns/uds-runtime/issues/176)) ([2fe0c99](https://github.com/defenseunicorns/uds-runtime/commit/2fe0c99e51dea34bbec0e5383ef2890e563d9f8d))
* **ui:** deployments table not showing accurate available count ([#192](https://github.com/defenseunicorns/uds-runtime/issues/192)) ([5b0f394](https://github.com/defenseunicorns/uds-runtime/commit/5b0f394ef80ea6286a017093851b4f42ae16009c))
* **ui:** restart and node columns in pod table ([#194](https://github.com/defenseunicorns/uds-runtime/issues/194)) ([88a41c4](https://github.com/defenseunicorns/uds-runtime/commit/88a41c4fc4838945361d8d3117086664273b3a50))


### Miscellaneous

* add resource store tests ([#74](https://github.com/defenseunicorns/uds-runtime/issues/74)) ([33ad7bc](https://github.com/defenseunicorns/uds-runtime/commit/33ad7bcee1e2231ab2570865932bd29b6194d997))
* add test infra for dummy cluster ([#157](https://github.com/defenseunicorns/uds-runtime/issues/157)) ([2ba9e26](https://github.com/defenseunicorns/uds-runtime/commit/2ba9e261255c57b1631249e3aa1864eb5c7b34f1))
* adds basic e2e tests ([#175](https://github.com/defenseunicorns/uds-runtime/issues/175)) ([770c4a3](https://github.com/defenseunicorns/uds-runtime/commit/770c4a32a0ac1aae4e58cc8af89c6d998513a688))
* adds navbar tests ([#189](https://github.com/defenseunicorns/uds-runtime/issues/189)) ([247afc7](https://github.com/defenseunicorns/uds-runtime/commit/247afc7813fe560648453764f1b891845483a9c4))
* adds ssm to runtime-canary ([#221](https://github.com/defenseunicorns/uds-runtime/issues/221)) ([13d2350](https://github.com/defenseunicorns/uds-runtime/commit/13d2350e0e6f2c2ce6a6d06c9c3e54855a762990))
* api testing adr ([#198](https://github.com/defenseunicorns/uds-runtime/issues/198)) ([f50bd77](https://github.com/defenseunicorns/uds-runtime/commit/f50bd773740300745e19f7821a1a9143a3c12679))
* **ci:** add swagger docs gen to CI ([#167](https://github.com/defenseunicorns/uds-runtime/issues/167)) ([aed5b83](https://github.com/defenseunicorns/uds-runtime/commit/aed5b836288fb5592b254302355d05cb7e7ee3bc))
* **ci:** add type checking ([#187](https://github.com/defenseunicorns/uds-runtime/issues/187)) ([3357983](https://github.com/defenseunicorns/uds-runtime/commit/3357983a961bc565aca55464dc19b8c6e346cd81))
* codeowners ([#174](https://github.com/defenseunicorns/uds-runtime/issues/174)) ([b4cabcf](https://github.com/defenseunicorns/uds-runtime/commit/b4cabcf34c90d40fc2c6467932ac42c2f34e9fc6))
* configure renovate ([#209](https://github.com/defenseunicorns/uds-runtime/issues/209)) ([4712c6f](https://github.com/defenseunicorns/uds-runtime/commit/4712c6f71d04b37e6f55dc57b2b910ba3c86429c))
* **deps:** update dependency kubernetes-fluent-client to v3 ([#218](https://github.com/defenseunicorns/uds-runtime/issues/218)) ([245c256](https://github.com/defenseunicorns/uds-runtime/commit/245c25633a5019dc8a52649d2830173c1e95da5c))
* **deps:** update module github.com/zarf-dev/zarf to v0.38.2 ([#216](https://github.com/defenseunicorns/uds-runtime/issues/216)) ([21c7457](https://github.com/defenseunicorns/uds-runtime/commit/21c745736397f461e9611da42f8b24ff1c89c493))
* go unit tests ([#191](https://github.com/defenseunicorns/uds-runtime/issues/191)) ([b850316](https://github.com/defenseunicorns/uds-runtime/commit/b85031664c6e2cc68cf28379961d30ccf62fabd0))
* install Istio CRDs for minimal e2e tests ([#168](https://github.com/defenseunicorns/uds-runtime/issues/168)) ([8832dcd](https://github.com/defenseunicorns/uds-runtime/commit/8832dcd203cfc3086fa2280d76e4dd901b74f2c9))
* make uds package host configurable for nightly ([#201](https://github.com/defenseunicorns/uds-runtime/issues/201)) ([bb34173](https://github.com/defenseunicorns/uds-runtime/commit/bb34173b6418d17c9aedc4876941a80927a7645e))
* speed up e2e tests with minimal cluster ([#155](https://github.com/defenseunicorns/uds-runtime/issues/155)) ([b8aae8a](https://github.com/defenseunicorns/uds-runtime/commit/b8aae8a9fb31be7be57394817769feeb19ee5747))
* update pullPolicy ([#220](https://github.com/defenseunicorns/uds-runtime/issues/220)) ([4320468](https://github.com/defenseunicorns/uds-runtime/commit/4320468956c1e44e425e0492bf1faaa1d1b44d20))
* updating README ([#204](https://github.com/defenseunicorns/uds-runtime/issues/204)) ([69aab8e](https://github.com/defenseunicorns/uds-runtime/commit/69aab8e21eb3b76318f5f52da8458f8597370d8a))

## [0.1.0](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0-alpha.5...v0.1.0) (2024-08-02)


### Features

* **ui:** updating drawer header ([#142](https://github.com/defenseunicorns/uds-runtime/issues/142)) ([98dad19](https://github.com/defenseunicorns/uds-runtime/commit/98dad199800d06a0fb2690673dc35785784400a0)), closes [#102](https://github.com/defenseunicorns/uds-runtime/issues/102)


### Bug Fixes

* enable release-please PRs to run workflows ([#146](https://github.com/defenseunicorns/uds-runtime/issues/146)) ([10d7858](https://github.com/defenseunicorns/uds-runtime/commit/10d7858adcffd2f977c6101a25b81696dd266f00))
* remove alpha from release ([#145](https://github.com/defenseunicorns/uds-runtime/issues/145)) ([70344d8](https://github.com/defenseunicorns/uds-runtime/commit/70344d8a07a33df57d90cc848974e08fde327ac7))


### Miscellaneous

* **ui:** remove breadcrumb ([#132](https://github.com/defenseunicorns/uds-runtime/issues/132)) ([982c278](https://github.com/defenseunicorns/uds-runtime/commit/982c278cba5226515f9aef5942e5fcfa7bb3e76e))

## [0.1.0-alpha.5](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0-alpha.4...v0.1.0-alpha.5) (2024-08-02)


### Features

* **ui:** make package endpoints links ([#136](https://github.com/defenseunicorns/uds-runtime/issues/136)) ([572f385](https://github.com/defenseunicorns/uds-runtime/commit/572f3859b0ba5901182fe335eb5e9ac679635716))


### Bug Fixes

* add tasks file to release please config and add docker login ([#139](https://github.com/defenseunicorns/uds-runtime/issues/139)) ([f61680f](https://github.com/defenseunicorns/uds-runtime/commit/f61680f2941bb8e092b8da69d5c9cbcdcbb12d43))

## [0.1.0-alpha.4](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0-alpha.3...v0.1.0-alpha.4) (2024-08-01)


### Bug Fixes

* **ci:** more permissions for release ([#135](https://github.com/defenseunicorns/uds-runtime/issues/135)) ([3695630](https://github.com/defenseunicorns/uds-runtime/commit/36956305edfbcf2fa94befb20d900d1c7503d6b3))

## [0.1.0-alpha.3](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0-alpha.2...v0.1.0-alpha.3) (2024-08-01)


### Features

* **ui:** show results even if not filtering ([#131](https://github.com/defenseunicorns/uds-runtime/issues/131)) ([4603b65](https://github.com/defenseunicorns/uds-runtime/commit/4603b656b41e3a94883a4694a4a881f27512424b))


### Bug Fixes

* use docker buildx in github actions ([#133](https://github.com/defenseunicorns/uds-runtime/issues/133)) ([d52cb6c](https://github.com/defenseunicorns/uds-runtime/commit/d52cb6cebb3eaab211d0087726b7c3f2a4d18b4b))

## [0.1.0-alpha.2](https://github.com/defenseunicorns/uds-runtime/compare/v0.1.0-alpha.1...v0.1.0-alpha.2) (2024-08-01)


### Features

* add missing routes & route filter ([4235c60](https://github.com/defenseunicorns/uds-runtime/commit/4235c60b5243a606c3c02129c3508655cfb6d51a))
* add pod disruption budgets view ([#79](https://github.com/defenseunicorns/uds-runtime/issues/79)) ([a2f17b3](https://github.com/defenseunicorns/uds-runtime/commit/a2f17b39cd0b0304c5bfb9e1cea9b2320f62ef76)), closes [#48](https://github.com/defenseunicorns/uds-runtime/issues/48)
* add sparse streams, compression, single resource api endpoints ([#28](https://github.com/defenseunicorns/uds-runtime/issues/28)) ([c5c4c9c](https://github.com/defenseunicorns/uds-runtime/commit/c5c4c9c5e98cd81e4ea5a0e295983ced1ce05832))
* add storageclasses view ([#83](https://github.com/defenseunicorns/uds-runtime/issues/83)) ([941fff2](https://github.com/defenseunicorns/uds-runtime/commit/941fff2715ee85d2dfda0fb26fbb59fc10f2708f)), closes [#60](https://github.com/defenseunicorns/uds-runtime/issues/60)
* add UDS Exemptions view ([ac0d43e](https://github.com/defenseunicorns/uds-runtime/commit/ac0d43ec947d47341aa4e233aebeefe2e51aeb78))
* add uds package view ([7be0628](https://github.com/defenseunicorns/uds-runtime/commit/7be06284014e0f4b0337c345873b1e0951c09951))
* adding cluster ops limit ranges view ([#78](https://github.com/defenseunicorns/uds-runtime/issues/78)) ([6907151](https://github.com/defenseunicorns/uds-runtime/commit/69071515bdefa8d50eb04465e29b96b1fa6b2c5f)), closes [#47](https://github.com/defenseunicorns/uds-runtime/issues/47)
* adds pv view ([#71](https://github.com/defenseunicorns/uds-runtime/issues/71)) ([cd013d5](https://github.com/defenseunicorns/uds-runtime/commit/cd013d5271e049310b2d3a6aaf9a0ae2a5b5f5dc))
* **api:** add remaining k8s resource endpoints ([7fa24ee](https://github.com/defenseunicorns/uds-runtime/commit/7fa24ee72ea7ba9ee38f079aa5ed4863f497650b))
* **api:** debounce SSE messages ([#100](https://github.com/defenseunicorns/uds-runtime/issues/100)) ([27d7433](https://github.com/defenseunicorns/uds-runtime/commit/27d7433d37bbae977bc9f0269f717056fdcb0d6c))
* enables uds-runtime releases ([#93](https://github.com/defenseunicorns/uds-runtime/issues/93)) ([6af3a95](https://github.com/defenseunicorns/uds-runtime/commit/6af3a95e574f95f1df8ad39a696b311f2c63d16e))
* export pepr data stream ([#24](https://github.com/defenseunicorns/uds-runtime/issues/24)) ([b05976f](https://github.com/defenseunicorns/uds-runtime/commit/b05976fad98a3eaa543f21db8e4f24a3b99805de)), closes [#23](https://github.com/defenseunicorns/uds-runtime/issues/23)
* swagger ([#114](https://github.com/defenseunicorns/uds-runtime/issues/114)) ([c4016ca](https://github.com/defenseunicorns/uds-runtime/commit/c4016caaef1d1e5e0b6114d351521b0f80e99663))
* **ui:** add cluster ops hpa k8s view ([#97](https://github.com/defenseunicorns/uds-runtime/issues/97)) ([a55696d](https://github.com/defenseunicorns/uds-runtime/commit/a55696defbdf7215f752fc36a8d666b6995476f0)), closes [#46](https://github.com/defenseunicorns/uds-runtime/issues/46)
* **ui:** add cluster-ops/mutating-webhooks view ([#42](https://github.com/defenseunicorns/uds-runtime/issues/42)) ([b29c434](https://github.com/defenseunicorns/uds-runtime/commit/b29c434e0c5300194d57b9b831b0528ca3ffd743)), closes [#32](https://github.com/defenseunicorns/uds-runtime/issues/32)
* **ui:** add cronjobs view ([9a80579](https://github.com/defenseunicorns/uds-runtime/commit/9a80579fdcfa14af2426314dfdfb11f89089b72c))
* **ui:** add header to tables ([#106](https://github.com/defenseunicorns/uds-runtime/issues/106)) ([19cdcfe](https://github.com/defenseunicorns/uds-runtime/commit/19cdcfe4ef3c602bada5edc1e560dec311bcda60))
* **ui:** add jobs view ([125f635](https://github.com/defenseunicorns/uds-runtime/commit/125f63517ad9dc69993b7b178cd10a287896b7c9))
* **ui:** add nodes view ([f1cf0b0](https://github.com/defenseunicorns/uds-runtime/commit/f1cf0b0725a6deb4fd33eb74086a31de8059c875))
* **ui:** add pepr details ([#107](https://github.com/defenseunicorns/uds-runtime/issues/107)) ([a4ed8cd](https://github.com/defenseunicorns/uds-runtime/commit/a4ed8cd591ab1c691209fd55bd1bb764f3756952))
* **ui:** add pvc view ([#72](https://github.com/defenseunicorns/uds-runtime/issues/72)) ([5580dd4](https://github.com/defenseunicorns/uds-runtime/commit/5580dd4ece5725c248f8369607a67a31f0e20cdf))
* **ui:** add toast for status messages ([#116](https://github.com/defenseunicorns/uds-runtime/issues/116)) ([983ae6a](https://github.com/defenseunicorns/uds-runtime/commit/983ae6af3b9f9b55d29926b178ebcc228bb7e83f)), closes [#115](https://github.com/defenseunicorns/uds-runtime/issues/115)
* **ui:** adding 404 ([#98](https://github.com/defenseunicorns/uds-runtime/issues/98)) ([24b24b0](https://github.com/defenseunicorns/uds-runtime/commit/24b24b0516cbe563b4e6e7e2ebe7495173c4da5b)), closes [#38](https://github.com/defenseunicorns/uds-runtime/issues/38)
* **ui:** adding cluster ops priority classes view ([#62](https://github.com/defenseunicorns/uds-runtime/issues/62)) ([f52ea3d](https://github.com/defenseunicorns/uds-runtime/commit/f52ea3d47aabd868507e563885fd9133b03bd0af))
* **ui:** adding cluster ops resource quotas view ([#81](https://github.com/defenseunicorns/uds-runtime/issues/81)) ([5370bf9](https://github.com/defenseunicorns/uds-runtime/commit/5370bf98e6e7f65987933de13de863a7a0dd3cd4))
* **ui:** adding cluster ops runtime classes view ([#68](https://github.com/defenseunicorns/uds-runtime/issues/68)) ([084fc83](https://github.com/defenseunicorns/uds-runtime/commit/084fc831fb75a6a0bd70427888cea5f2bd6f57fa))
* **ui:** adds validatingwebhooks view ([#110](https://github.com/defenseunicorns/uds-runtime/issues/110)) ([cad4f35](https://github.com/defenseunicorns/uds-runtime/commit/cad4f3508fe301532a71e4bd04986f7113df406d))
* **ui:** endpoints view ([#80](https://github.com/defenseunicorns/uds-runtime/issues/80)) ([e9bf61e](https://github.com/defenseunicorns/uds-runtime/commit/e9bf61e3cb284769e6402a4a69e97bf99dd8e347))
* **ui:** format age with more detail ([#111](https://github.com/defenseunicorns/uds-runtime/issues/111)) ([98916d6](https://github.com/defenseunicorns/uds-runtime/commit/98916d6328d9c131aa5394058eaeb77b10ec8bf3))
* **ui:** networkpolicies view ([#75](https://github.com/defenseunicorns/uds-runtime/issues/75)) ([d04d50b](https://github.com/defenseunicorns/uds-runtime/commit/d04d50bab0c23ba15952a7977a3046e280e39f4f))
* **ui:** networks-services view ([#61](https://github.com/defenseunicorns/uds-runtime/issues/61)) ([1d0b68a](https://github.com/defenseunicorns/uds-runtime/commit/1d0b68aa39ad59e79107ddb61d650ae427cc9598))
* **ui:** resource detail view ([#95](https://github.com/defenseunicorns/uds-runtime/issues/95)) ([34e7e6c](https://github.com/defenseunicorns/uds-runtime/commit/34e7e6ca168fd3baa7b12dae7d4edd66f6343a8a)), closes [#34](https://github.com/defenseunicorns/uds-runtime/issues/34)
* **ui:** virtualservices view ([#73](https://github.com/defenseunicorns/uds-runtime/issues/73)) ([4d01cb8](https://github.com/defenseunicorns/uds-runtime/commit/4d01cb8db3125c747f5e26b835edc9fd0cf61131))


### Bug Fixes

* add missing routes ([#6](https://github.com/defenseunicorns/uds-runtime/issues/6)) ([3e61370](https://github.com/defenseunicorns/uds-runtime/commit/3e61370bb0443beefcb7b69f298570baca8837b2))
* fixing tooltip behavior ([#26](https://github.com/defenseunicorns/uds-runtime/issues/26)) ([dd8b092](https://github.com/defenseunicorns/uds-runtime/commit/dd8b0928be7e0aa4b9c656a8ac83e8656fa4f530)), closes [#25](https://github.com/defenseunicorns/uds-runtime/issues/25)
* remove breadcrumb links, closes [#9](https://github.com/defenseunicorns/uds-runtime/issues/9) ([e2d8cfe](https://github.com/defenseunicorns/uds-runtime/commit/e2d8cfecb5c1ce1c863435eabf42937f57100401))
* sidebar expand toggle icon rotation ([e1aa7cb](https://github.com/defenseunicorns/uds-runtime/commit/e1aa7cbe97faae1044aa895ee338427ab3b55180))
* sidebar ux issues ([589a970](https://github.com/defenseunicorns/uds-runtime/commit/589a9700794e10332e953f63682c4b3d970797cd))
* turn off autocomplete for datatable ([#22](https://github.com/defenseunicorns/uds-runtime/issues/22)) ([f71eccf](https://github.com/defenseunicorns/uds-runtime/commit/f71eccfbe084f2c9808b970fff16c3c8d12c4ad1)), closes [#21](https://github.com/defenseunicorns/uds-runtime/issues/21)
* **ui:** pepr monitor unsubscriber ([#67](https://github.com/defenseunicorns/uds-runtime/issues/67)) ([2995497](https://github.com/defenseunicorns/uds-runtime/commit/2995497a02fb1b6b6a261aad614c4b4bfe84bc98)), closes [#37](https://github.com/defenseunicorns/uds-runtime/issues/37)
* **ui:** prevent error on null fields in DataTable component ([#43](https://github.com/defenseunicorns/uds-runtime/issues/43)) ([d51636e](https://github.com/defenseunicorns/uds-runtime/commit/d51636e5ee9154c1863d06222d927c6e61f26cb1))
* use middleware for compression ([#86](https://github.com/defenseunicorns/uds-runtime/issues/86)) ([f16a18c](https://github.com/defenseunicorns/uds-runtime/commit/f16a18c755f67cad6ffdd3a21d1708af4f99813e)), closes [#85](https://github.com/defenseunicorns/uds-runtime/issues/85)


### Miscellaneous

* add contributing guide ([#17](https://github.com/defenseunicorns/uds-runtime/issues/17)) ([f03c08f](https://github.com/defenseunicorns/uds-runtime/commit/f03c08f73abccdc7dd69744fdbd5a25ad4d6371f)), closes [#18](https://github.com/defenseunicorns/uds-runtime/issues/18)
* add empty release please manifest to bootstrap ([#120](https://github.com/defenseunicorns/uds-runtime/issues/120)) ([139fe0a](https://github.com/defenseunicorns/uds-runtime/commit/139fe0a0dc32d04c6728b0de3f82a102ef7638f9))
* add initial ci ([#11](https://github.com/defenseunicorns/uds-runtime/issues/11)) ([57248fb](https://github.com/defenseunicorns/uds-runtime/commit/57248fb1c2bfe03131b7bc8ce90f6ef419078838))
* add release ci ([#118](https://github.com/defenseunicorns/uds-runtime/issues/118)) ([70c0984](https://github.com/defenseunicorns/uds-runtime/commit/70c09845408ec8e4e62916bef7ae5d4cd4ebc87e))
* add uds-cli taks & zarf schema configs for vscode ([344fd7e](https://github.com/defenseunicorns/uds-runtime/commit/344fd7e6ab217e7355e58073e07a8972986249d9))
* add vscode go debug config ([c811f47](https://github.com/defenseunicorns/uds-runtime/commit/c811f47f804cd21d08300d4cc906bc5e6c9f8536))
* adding all routes for navigation ([#16](https://github.com/defenseunicorns/uds-runtime/issues/16)) ([eb518c3](https://github.com/defenseunicorns/uds-runtime/commit/eb518c39a9d45e74fc76f15abd2346ae304478a7)), closes [#15](https://github.com/defenseunicorns/uds-runtime/issues/15)
* bump spdx header to 2024-Present ([2102add](https://github.com/defenseunicorns/uds-runtime/commit/2102addd1864b67a82b59e4ac11c4f7fb51f0824))
* **deps:** update all dependencies ([#13](https://github.com/defenseunicorns/uds-runtime/issues/13)) ([ec95b7e](https://github.com/defenseunicorns/uds-runtime/commit/ec95b7e67e07246489220e1d6783ad9ef195ca74))
* ensure lang=ts for all svelte script tags ([b40ecbb](https://github.com/defenseunicorns/uds-runtime/commit/b40ecbba365fecc7eacd312ee37765a3a297f3b6))
* fix lint errors ([68df75f](https://github.com/defenseunicorns/uds-runtime/commit/68df75fc7f8c6d3693d85c30663fff2cb9e9fe2d))
* lint ([ef14348](https://github.com/defenseunicorns/uds-runtime/commit/ef1434877c8feb39ab3b1601c02e24269abd43f8))
* remove message.Fatal for zarf upgrade ([8438b37](https://github.com/defenseunicorns/uds-runtime/commit/8438b37532367f160487c77b9dbdce500d6b9e8e))
* setting up release please ([16efc50](https://github.com/defenseunicorns/uds-runtime/commit/16efc50f1382448acabd4ffc28b650781dcb9118))
* uds-engine -&gt; uds-runtime ([e553195](https://github.com/defenseunicorns/uds-runtime/commit/e553195ffc0cef8dd0b4fca0b6cf350301ecdcfc))
* update config sidebar icon ([9321de2](https://github.com/defenseunicorns/uds-runtime/commit/9321de2f6ef67f3267c639e22657536e510a4039))
* update zarf ([#99](https://github.com/defenseunicorns/uds-runtime/issues/99)) ([d7c0bee](https://github.com/defenseunicorns/uds-runtime/commit/d7c0bee206272b7fff36e545725145eea315187d))
