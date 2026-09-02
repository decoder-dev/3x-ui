[English](/README.md) | [Русский](/README.ru_RU.md)

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="3хуя" src="./media/3x-ui-light.png">
  </picture>
</p>

<p align="center">
  <a href="https://github.com/decoder-dev/3x-ui/releases"><img src="https://img.shields.io/github/v/release/decoder-dev/3x-ui?label=release" alt="Release"></a>
  <a href="https://github.com/decoder-dev/3x-ui/actions"><img src="https://img.shields.io/github/actions/workflow/status/decoder-dev/3x-ui/release.yml.svg" alt="Build"></a>
  <a href="#"><img src="https://img.shields.io/github/go-mod/go-version/decoder-dev/3x-ui.svg" alt="GO Version"></a>
  <a href="https://www.gnu.org/licenses/gpl-3.0.en.html"><img src="https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true" alt="License"></a>
</p>

**3хуя** is a maintained fork of [3X-UI](https://github.com/MHSanaei/3x-ui) — an open-source web panel for [Xray-core](https://github.com/XTLS/Xray-core). It ships with **[decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)** (XHTTP variant B, `scStreamDown`, `downFrame`) and panel changes needed for CDN → nginx → XHTTP deployments.

> [!IMPORTANT]
> For personal/lab use only. Fork branding is display-only: paths, service name, and binary stay `x-ui` / `/usr/local/x-ui`.

## Why this fork

| Area | What 3хуя adds |
| --- | --- |
| **Xray** | Auto-install/update **decoder-dev** builds (`*-decoder` tags), `x-ui install-xray` |
| **XHTTP variant B** | UI fields `scStreamDownServerSecs`, `downFrame`; Clash/sub/link export |
| **Panel UX** | Rebrand to **3хуя**, mobile-friendly layout |
| **Stability** | `tgId` accepts JSON number or numeric string (legacy DB rows) |
| **Releases** | `decoder-dev/3x-ui` tarballs + tag `dev-latest` on every `main` push |

Full Russian install guide: [`INSTALL-decoder-dev.ru.md`](INSTALL-decoder-dev.ru.md).

## Features (inherited from 3X-UI)

- **Multi-protocol inbounds** — VLESS, VMess, Trojan, Shadowsocks, WireGuard, Hysteria2, HTTP, SOCKS, Tunnel, TUN, MTProto
- **Transports** — TCP, mKCP, WebSocket, gRPC, HTTPUpgrade, **XHTTP** (TLS / REALITY / Vision / ML-KEM)
- **Clients** — quotas, expiry, IP limits, HWID limits, online status, QR, subscriptions
- **Multi-node**, outbound routing, WARP, balancers, subscription server (Clash / JSON / raw)
- **Telegram bot**, REST API + Swagger, SQLite or PostgreSQL, 13 UI languages, Fail2ban

## Quick start

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Specific release (recommended):

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) v3.7.0-decoder
```

Rolling build from `main`:

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) dev-latest
```

After install, run `x-ui` for the management menu (credentials, SSL, restart, update).

### Replace Xray only (existing panel)

```bash
x-ui install-xray
# or see INSTALL-decoder-dev.ru.md for manual steps
```

### Unattended install

```bash
export XUI_NONINTERACTIVE=1
export XUI_USERNAME=admin
export XUI_PASSWORD='YourStrongPassword'
export XUI_PANEL_PORT=54321
export XUI_WEB_BASE_PATH='your-secret-path'
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Credentials: `/etc/x-ui/install-result.env`.

## Repositories

| Component | Repository | Example tag |
| --- | --- | --- |
| Panel | [decoder-dev/3x-ui](https://github.com/decoder-dev/3x-ui) | `v3.7.0-decoder`, `dev-latest` |
| Xray | [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core) | `v26.7.28-decoder` |

Upstream docs (most topics still apply): [MHSanaei/3x-ui Wiki](https://github.com/MHSanaei/3x-ui/wiki).

## Supported platforms

**OS:** Ubuntu, Debian, Fedora, CentOS, RHEL, Alma/Rocky, Arch, Alpine, Windows, and others supported by upstream installer.

**Arch:** `amd64` · `386` · `arm64` · `armv7` · `armv6` · `armv5` · `s390x`.

## Database

Same as upstream: SQLite (default) or PostgreSQL via `XUI_DB_TYPE` / `XUI_DB_DSN` in `/etc/default/x-ui`.

## Environment variables

Same as [upstream README](https://github.com/MHSanaei/3x-ui#environment-variables), plus installer reads `XUI_GITHUB_REPO=decoder-dev/3x-ui` by default in this fork.

## Updating

```bash
x-ui update          # panel from decoder-dev/3x-ui releases
x-ui install-xray    # refresh decoder-dev Xray binary
```

## Acknowledgments

- [MHSanaei/3x-ui](https://github.com/MHSanaei/3x-ui) — base panel
- [XTLS/Xray-core](https://github.com/XTLS/Xray-core) — core; decoder patches in [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)
- [alireza0](https://github.com/alireza0/) — original X-UI lineage

Licensed under **GPL-3.0** (same as upstream).
