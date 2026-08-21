# contracts-automation

Protobuf/gRPC contract interfaces for media acquisition automation modules in MuxCore.

## Proto Service

- **`AutomationService`** (`proto/muxcore/automation/v1/automation.proto`)
  - `SearchItem` — search indexers for a wanted movie or episode
  - `Dispatch` — send a selected release to a downloader
  - `AddToQueue` / `GetQueue` / `RemoveFromQueue` — manage the wanted queue
  - `GetHistory` — download dispatch history
  - `SearchNow` — trigger an immediate search pass
  - `ListBlocklist` / `ClearBlocklist` / `BlocklistRelease` — release blocklist management
  - `ListDelayProfiles` / `UpsertDelayProfile` — per-protocol delay profiles
  - `ListCutoffUnmet` — items below quality cutoff
  - `RetryImport` — retry failed post-download imports

Proto source was migrated from `media-automation/proto/automationv1/`.

## Go import

Generated stubs (after `make proto`): `github.com/Muxcore-Media/contracts-automation/muxcore/automation/v1`

Generated Go is not checked in yet; run `make proto` when `protoc` and the Go plugins are available.

## Implementing Modules

- [media-automation](https://github.com/Muxcore-Media/media-automation) — wanted queue, search, and dispatch orchestration
