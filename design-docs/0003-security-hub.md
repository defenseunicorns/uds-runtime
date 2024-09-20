# Security Hub Integration

Author(s): @decleaver and @TristanHoladay  
Date Created: Sept 20, 2024  
Status: DRAFT

### Problem Statement

Runtime needs to show security information, like SBOMs and CVEs, for each UDS package in the cluster. To do this, we'll need to integrate with the UDS Security Hub database and tooling for generating the needed information.

### Unknowns

1. how are we going to pull the db from s3?
   - runtime github action
1. what exactly is in this db?
   - looks like every thursday it does scans of core, SWF, and Leapfrog packages.
1. how do we "package" db with Runtime?
   - in image so Core doesn't have to also pull and package the db?...
   - is this db sensitive so that it actually does need to be pulled at deploy time so it doesn't live in any repo / image
1. how do we generate the data required for the security views?
   - there's already data for certain packages, what about other packages not scanned in this database?
     - use the [uds-security-hub application](https://github.com/defenseunicorns/uds-security-hub?tab=readme-ov-file) directly?
     - use the vendored version of the tool in UDS CLI

### Implementation Details

TBD

### Alternatives Considered
