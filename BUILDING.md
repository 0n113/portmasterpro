# portmasterpro — Build & Test Guide

Diese Anleitung beschreibt wie du portmasterpro lokal kompilierst, testest
und die `no-spn-no-telemetry`-Änderungen verifizierst.

---

## Voraussetzungen

| Tool | Mindestversion | Installieren |
|---|---|---|
| Go | 1.22 | https://go.dev/dl/ |
| Git | beliebig | https://git-scm.com |
| Node.js | 20 LTS | https://nodejs.org (nur für Desktop-Tests) |
| pnpm / npm | beliebig | `npm i -g pnpm` (nur für Desktop-Tests) |
| gcc / build-essential | beliebig | `sudo apt install build-essential` (Linux) |

> **Windows:** Nutze WSL2 (Ubuntu 22.04+) oder installiere
> [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) für cgo-Support.

---

## 1 · Repository klonen

```bash
git clone https://github.com/0n113/portmasterpro.git
cd portmasterpro

# Branch mit den Telemetrie-/SPN-Bereinigungen auschecken
git checkout no-spn-no-telemetry
```

---

## 2 · Go-Abhängigkeiten laden & aufräumen

```bash
# Abhängigkeiten herunterladen
go mod download

# Nicht mehr benötigte Einträge aus go.mod / go.sum entfernen
# (nach dem Entfernen der SPN-Pakete wichtig)
go mod tidy
```

> `go mod tidy` aktualisiert `go.sum` automatisch.
> Veränderte Dateien danach committen:
> ```bash
> git add go.mod go.sum
> git commit -m "chore: go mod tidy after SPN removal"
> ```

---

## 3 · Go-Code kompilieren

### Nur prüfen ob alles kompiliert (kein Binary)

```bash
go build ./...
```

Erwartete Ausgabe: keine Fehler, kein Output.

### Binary bauen (Linux)

```bash
mkdir -p build
go build -o build/portmasterpro ./cmds/portmaster-core/
```

### Binary bauen (Windows, in WSL2)

```bash
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 \
  CC=x86_64-w64-mingw32-gcc \
  go build -o build/portmasterpro.exe ./cmds/portmaster-core/
```

---

## 4 · Go-Tests ausführen

### Alle Tests (komplett)

```bash
go test ./...
```

### Nur die Telemetrie-/SPN-Stub-Tests

```bash
# service/sync — Stub startet/stoppt ohne Fehler, kein Manager
go test -v ./service/sync/...

# spn/access — kein Login, alle Features offen, kein Manager
go test -v ./spn/access/...

# service/netquery — Endpoints ungated, kein SPN-Import
go test -v ./service/netquery/...
```

### Erwartete Ausgabe (Beispiel `spn/access`)

```
=== RUN   TestAccessStubNoOp
--- PASS: TestAccessStubNoOp (0.00s)
=== RUN   TestAccessStubNotLoggedIn
--- PASS: TestAccessStubNotLoggedIn (0.00s)
=== RUN   TestAccessStubAllFeaturesUnlocked
--- PASS: TestAccessStubAllFeaturesUnlocked (0.00s)
=== RUN   TestAccessStubManagerIsNil
--- PASS: TestAccessStubManagerIsNil (0.00s)
PASS
ok  	github.com/safing/portmaster/spn/access
```

### Mit Race-Detector (empfohlen vor einem Merge)

```bash
go test -race ./service/sync/... ./spn/access/... ./service/netquery/...
```

### go vet (statische Analyse)

```bash
go vet ./...
```

---

## 5 · Manuelle Telemetrie-Prüfung

Sicherstellen dass keine bekannten Safing-Telemetrie-URLs mehr im Code sind:

```bash
# Sollte keine Ausgabe liefern:
grep -r 'account.safing.io\|updates.safing.io\|telemetry.safing.io\|api.safing.io' \
  ./service/sync/ ./spn/access/

echo "Exit $?: 0 = sauber, alles andere = Fund"
```

Sicherstellen dass `service/netquery` keine SPN-Pakete importiert:

```bash
# Sollte keine Ausgabe liefern:
grep -r 'safing/portmaster/spn' ./service/netquery/

echo "Exit $?: 0 = sauber"
```

---

## 6 · Desktop-Frontend testen (TypeScript / Vitest)

```bash
cd desktop

# Abhängigkeiten installieren
pnpm install        # oder: npm install

# Nur die SPN-Shim-Tests laufen lassen
pnpm vitest run src/lib/spn-removed.spec.ts

# Oder alle Frontend-Tests
pnpm test
```

### Erwartete Ausgabe

```
✓ SPN/account removal shim > SPN_REMOVED flag is true
✓ SPN/account removal shim > isLoggedIn() always returns false
✓ SPN/account removal shim > isSPNActive() always returns false
✓ SPN/account removal shim > hasFeature() always returns true for any feature string
✓ SPN/account removal shim > hasFeature() returns true even for empty string

Test Files  1 passed (1)
Tests       5 passed (5)
```

---

## 7 · Verbleibende SPN-Referenzen im Frontend finden

Nach dem Build alle Stellen suchen die noch auf alten SPN/Login-Code zeigen:

```bash
grep -rn 'spn\|SPN\|loginModal\|isLoggedIn\|isSPNActive\|hasFeature' \
  desktop/src \
  --include='*.ts' \
  --include='*.svelte' \
  --include='*.vue' \
  --include='*.tsx' \
  --exclude='spn-removed*'
```

Jeden Treffer durch den entsprechenden Import aus
`desktop/src/lib/spn-removed.ts` ersetzen:

```ts
import { hasFeature, isLoggedIn, isSPNActive } from '$lib/spn-removed';
```

---

## 8 · Vollständiger Check-Ablauf vor dem Merge

Diesen Block einmalig von oben nach unten durchlaufen:

```bash
# 1. Sauber starten
git checkout no-spn-no-telemetry
git pull

# 2. Deps aufräumen
go mod tidy

# 3. Kompilieren
go build ./...

# 4. Vet
go vet ./...

# 5. Tests (mit Race-Detector)
go test -race ./...

# 6. Telemetrie-URLs prüfen
grep -r 'account.safing.io\|updates.safing.io\|api.safing.io' \
  ./service/sync/ ./spn/access/ && echo "FUND!" || echo "Sauber."

# 7. SPN-Import in netquery prüfen
grep -r 'safing/portmaster/spn' ./service/netquery/ && echo "FUND!" || echo "Sauber."

# 8. Frontend
cd desktop && pnpm install && pnpm vitest run src/lib/spn-removed.spec.ts
```

Wenn alle Schritte ohne Fehler und ohne "FUND!" durchlaufen → Branch ist merge-ready.

---

## Troubleshooting

### `undefined: mgr.Manager` beim Kompilieren
Die Stub-Dateien importieren `service/mgr`. Stelle sicher dass du auf dem
richtigen Branch bist und `go mod tidy` ausgeführt hast.

```bash
git branch          # muss 'no-spn-no-telemetry' zeigen
go mod tidy
go build ./...
```

### `cgo: C compiler not found`

```bash
# Ubuntu/Debian
sudo apt install build-essential

# Fedora/Rocky
sudo dnf groupinstall 'Development Tools'
```

### `pnpm: command not found`

```bash
npm install -g pnpm
```

### Test schlägt fehl mit `only one instance allowed`
Das `shimLoaded`-Flag in `netquery` ist prozessweit. Tests einzeln ausführen:

```bash
go test -v -count=1 -run TestNetqueryEndpointPathsPresent ./service/netquery/
```
