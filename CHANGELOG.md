# honeymarker changelog

## [0.2.12] - 2024-04-09

### Maintenance

- Document use of configuration key instead of writekey (#79) | [@danielcompton](https://github.com/danielcompton)

## [0.2.11] - 2023-10-23

### Added

- Adding env support for write key (#76) | [@elliotpope](https://github.com/elliottpope)

### Fixes

- add missing permissions for label action (#71) | [@jsoref](https://github.com/jsoref)
- strip leading 'v' from version for Debian (#77) | [@jharley](https://github.com/jharley)

### Maintenance

- chore: integrate with internal Asana (#66, #67, #68, #73)
- maint: bump github.com/jessevdk/go-flags (#69)
- maint: bump golang.org/x/sys from 0.0.0-20210320140829-1e4c9ba3b0c4 to 0.1.0 (#70)
- chore: spelling fixes (#72) | [@jsoref](https://github.com/jsoref)
- chore: switch to temp credentials for CI (#74) | [@NLincoln](https://github.com/nlincoln)

## [0.2.10] - 2022-07-20

### Maintenance

- Release to fix OpenSSL CVE

## [0.2.9] - 2022-06-02

### Fixes
- Change --version flag to version subcommand (#59) | [@kentquirk](https://github.com/kentquirk)

## [0.2.8] - 2022-05-25

### Maintenance

- Strip the leading v from the version number for Debian packages (#55) | [@kentquirk](https://github.com/kentquirk)

### Added
- Add --version flag for printing the version (#56) | [@kentquirk](https://github.com/kentquirk)

## [0.2.7] - 2022-05-12

### Maintenance

- A few more changes to CI pipeline to fix 64-bit RPM issues (#52) | [@kentquirk](https://github.com/kentquirk)

## [0.2.6] - 2022-05-10

### Maintenance

- Rework and modernize builds (#49) | [@kentquirk](https://github.com/kentquirk)

## [0.2.5] - 2022-05-02

- Resolve CVE by upgrading to a newer docker version (#47) | [@kentquirk](https://github.com/kentquirk)

## [0.2.4] - 2022-03-29

### Maintenance

- maint: update release job (#45) | [@vreynolds](https://github.com/vreynolds)
- docs: how to use with environments (#44) | [@vreynolds](https://github.com/vreynolds)

## [0.2.3] - 2022-02-10

### Fixes

- honeymarker: omit empty datetime fields on req (#41) | [@bradolson-virta](https://github.com/bradolson-virta)

## [0.2.2] - 2022-01-19

- update go to v1.17 (#38) | [@MikeGoldSmith](https://github.com/MikeGoldsmith)
- ci: allow forked builds (#37) | [@vreynolds](https://github.com/vreynolds)
- gh: add re-triage workflow (#36) | [@vreynolds](https://github.com/vreynolds)
- empower apply-labels action to apply labels (#35) | [@robbkidd](https://github.com/robbkidd)

## [0.2.1] - 2021-10-13

### Added

- Build and publish multi-arch docker images on tag (#27) | [@MikeGoldSmith](https://github.com/MikeGoldsmith)

### Maintenance

- Publish to docker hub on release (#22) | [@MikeGoldSmith](https://github.com/MikeGoldsmith)
- Update installation command (#29) | [@bixu](https://github.com/bixu)
- Change maintenance badge to maintained (#25)
- Adds Stalebot (#26)
- Add NOTICE (#23)
- Add issue and PR templates (#21)
- Add OSS lifecycle badge (#20)
- Add community health files (#19)
- Updates GitHub Action Workflows (#18)
- Adds Dependabot (#17)
- Switches CODEOWNERS to telemetry-team (#16)
