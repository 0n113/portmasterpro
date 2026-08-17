# Desktop UI Changes — portmasterpro

## SPN & Account Login Removal

The following UI elements have been **removed or stubbed** as part of the
`no-spn-no-telemetry` cleanup:

### Removed
- SPN enable/disable toggle in the navigation bar
- Account login / logout button
- SPN connection status indicator (globe icon with node count)
- "Upgrade to SPN" upsell banners and modals
- Account dashboard page (`/account`)
- SPN node map / exit-node selector
- Feature-gate overlays ("This feature requires a SPN subscription")

### Replaced with shim (`desktop/src/lib/spn-removed.ts`)
| Old call | New call | Behaviour |
|---|---|---|
| `isLoggedIn()` | `isLoggedIn()` from shim | Always `false` |
| `isSPNActive()` | `isSPNActive()` from shim | Always `false` |
| `hasFeature(x)` | `hasFeature(x)` from shim | Always `true` |
| `openLoginModal()` | `openLoginModal()` from shim | No-op + console.warn |
| `toggleSPN()` | `toggleSPN()` from shim | No-op + console.warn |

### Network History & Bandwidth — Always Visible
Bandwidth charts and network history are now always shown.
The previous `hasFeature('bandwidth-visibility')` guard has been removed;
all users see the full data view immediately.

### How to Find Remaining SPN References
```bash
# Find any remaining SPN/account UI references:
grep -r 'spn\|SPN\|loginModal\|isLoggedIn\|isSPNActive\|hasFeature' \
  desktop/src --include='*.ts' --include='*.svelte' --include='*.vue' \
  --exclude='spn-removed*'
```
Replace each hit with the appropriate shim function from `spn-removed.ts`.
