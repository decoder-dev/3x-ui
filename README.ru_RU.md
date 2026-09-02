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

**3хуя** — поддерживаемый форк [3X-UI](https://github.com/MHSanaei/3x-ui), веб-панель для [Xray-core](https://github.com/XTLS/Xray-core). В комплекте **[decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)** с XHTTP variant B (`scStreamDown`, `downFrame`) и доработками панели под схему CDN → nginx → XHTTP.

> [!IMPORTANT]
> Только для личного/лабораторного использования. Брендинг **3хуя** — только в интерфейсе; пути, systemd и бинарник остаются `x-ui` / `/usr/local/x-ui`.

## Зачем этот форк

| Область | Что даёт 3хуя |
| --- | --- |
| **Xray** | Автоустановка/обновление сборок **decoder-dev** (`*-decoder`), команда `x-ui install-xray` |
| **XHTTP variant B** | Поля `scStreamDownServerSecs`, `downFrame` в UI, подписке и Clash |
| **Интерфейс** | Брендинг **3хуя**, адаптивная вёрстка под мобильные |
| **Стабильность** | `tgId` принимает число или строку в JSON (старые записи в БД) |
| **Релизы** | Тарболы `decoder-dev/3x-ui`, тег `dev-latest` на каждый push в `main` |

Подробная установка: [`INSTALL-decoder-dev.ru.md`](INSTALL-decoder-dev.ru.md).

## Возможности (от 3X-UI)

- **Протоколы** — VLESS, VMess, Trojan, Shadowsocks, WireGuard, Hysteria2, HTTP, SOCKS, Tunnel, TUN, MTProto
- **Транспорты** — TCP, mKCP, WebSocket, gRPC, HTTPUpgrade, **XHTTP** (TLS / REALITY / Vision / ML-KEM)
- **Клиенты** — квоты, срок, лимит IP, лимит HWID, онлайн, QR, подписки
- **Мульти-ноды**, routing, WARP, балансировщики, сервер подписок (Clash / JSON / raw)
- **Telegram-бот**, REST API + Swagger, SQLite или PostgreSQL, 13 языков, Fail2ban

## Быстрый старт

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Конкретный релиз:

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) v3.7.0-decoder
```

Dev-сборка с `main`:

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) dev-latest
```

После установки: `x-ui` — меню (логин, SSL, перезапуск, обновление).

### Только заменить Xray (панель уже стоит)

```bash
x-ui install-xray
```

### Без диалогов (cloud-init / VPS)

```bash
export XUI_NONINTERACTIVE=1
export XUI_USERNAME=admin
export XUI_PASSWORD='YourStrongPassword'
export XUI_PANEL_PORT=54321
export XUI_WEB_BASE_PATH='your-secret-path'
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Учётные данные: `/etc/x-ui/install-result.env`.

## Репозитории

| Компонент | Репозиторий | Пример тега |
| --- | --- | --- |
| Панель | [decoder-dev/3x-ui](https://github.com/decoder-dev/3x-ui) | `v3.7.0-decoder`, `dev-latest` |
| Xray | [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core) | `v26.7.28-decoder` |

Общая документация upstream: [Wiki 3X-UI](https://github.com/MHSanaei/3x-ui/wiki).

## Платформы и БД

Как у upstream: Linux/Windows, архитектуры `amd64` … `s390x`; SQLite по умолчанию или PostgreSQL через `XUI_DB_TYPE` / `XUI_DB_DSN`.

## Обновление

```bash
x-ui update          # панель с релизов decoder-dev/3x-ui
x-ui install-xray    # обновить бинарник decoder-dev Xray
```

## Благодарности

- [MHSanaei/3x-ui](https://github.com/MHSanaei/3x-ui) — базовая панель
- [XTLS/Xray-core](https://github.com/XTLS/Xray-core) — ядро; патчи в [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)
- [alireza0](https://github.com/alireza0/) — lineage X-UI

Лицензия **GPL-3.0** (как у upstream).
