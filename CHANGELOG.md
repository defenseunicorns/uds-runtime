# Changelog

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
