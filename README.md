# contracts-automation

Protobuf/gRPC contract interfaces for media acquisition automation modules in MuxCore.

## Proto Service

- **`AutomationService`** (`proto/muxcore/automation/v1/automation.proto`)
  - `SearchItem` — search indexers for a wanted item (`item_type`: movie|tv|music|book|comic|audiobook)
  - `Dispatch` — send a selected release to a downloader
  - `AddToQueue` / `GetQueue` / `UpdateQueueItem` / `RemoveFromQueue` — manage the wanted queue
  - `GetHistory` — download dispatch history (filter by `status`, `wanted_item_id`)
  - `SearchNow` — trigger search (full queue or one item via `queue_id` / `item_type`+`item_id`)
  - `ListBlocklist` / `ClearBlocklist` / `BlocklistRelease` — release blocklist management
  - `ListDelayProfiles` / `UpsertDelayProfile` — per-protocol delay profiles
  - `ListCutoffUnmet` — items below quality cutoff
  - `ListSeriesOverrides` / `UpsertSeriesOverride` / `DeleteSeriesOverride` — per-series delay/group overrides
  - `RetryImport` — retry failed post-download imports
  - `GetCapabilities` — discover automation features (item types, delay/cutoff/blocklist, anime absolute, protocols)

Proto source was migrated from `media-automation/proto/automationv1/`.

## Go import

Generated stubs (after `make proto`): `github.com/Muxcore-Media/contracts-automation/muxcore/automation/v1`

Generated stubs: `github.com/Muxcore-Media/contracts-automation/muxcore/automation/v1` (checked in; run `make proto` after `.proto` edits).

## Implementing Modules

- [media-automation](https://github.com/Muxcore-Media/media-automation) — wanted queue, search, and dispatch orchestration
