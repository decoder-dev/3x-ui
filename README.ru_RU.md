[English](/README.md) | [Русский](/README.ru_RU.md)

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="3хуя" src="./media/3x-ui-light.png">
  </picture>
</p>

<p align="center">
  <a href="https://github.com/decoder-dev/3x-ui/releases"><img src="https://img.shields.io/github/v/release/decoder-dev/3x-ui?label=релиз" alt="Release"></a>
  <a href="https://github.com/decoder-dev/3x-ui/actions"><img src="https://img.shields.io/github/actions/workflow/status/decoder-dev/3x-ui/release.yml.svg" alt="Build"></a>
  <a href="https://www.gnu.org/licenses/gpl-3.0.en.html"><img src="https://img.shields.io/badge/license-GPL%20V3-blue.svg" alt="License"></a>
</p>

# 3хуя

**Контора на районе.** Не «ещё одна панелька» — а свой стол, свои понятия, свой Xray под **variant B**.

Форк [3X-UI](https://github.com/MHSanaei/3x-ui), только без понтов ради понтов: сюда залит **[decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)** — XHTTP packet-up, `scStreamDown`, `downFrame`, всё что надо, когда схема **CDN → nginx → XHTTP**, а не «поставил и молись».

> [!IMPORTANT]
> **По понятиям:** только для своих, лаборатории, домашнего VPS. Не для продакшена и не для серых схем — мы тут инфраструктуру крутим, а не законы переписываем.
>
> Бренд **3хуя** — в морде панели. Под капотом всё по-старому: `x-ui`, `/usr/local/x-ui`, systemd. Чтоб апдейты и скрипты не ломались.

---

## Что тут за расстановка

| Понятие | Что по факту |
| --- | --- |
| **Ядро** | Xray с decoder-dev — variant B из коробки, `x-ui install-xray` если надо перекатить |
| **XHTTP** | В UI есть `scStreamDownServerSecs` и `downFrame` — не наугад в конфиге, а как люди |
| **Подписка / Clash** | Поля variant B уезжают в sub и Clash/Meta — клиент видит то, что на сервере |
| **Морда** | Переименовано в **3хуя**, на телефоне не страдай — адаптив есть |
| **База** | `tgId` не падает от кривого JSON — старые записи в SQLite не трогают нервы |
| **Доставка** | Релизы `decoder-dev/3x-ui`, тег `dev-latest` — свежак с `main` без церемоний |

Полный разбор установки под наш эталон: [`INSTALL-decoder-dev.ru.md`](INSTALL-decoder-dev.ru.md).

---

## Что умеет (наследие 3X-UI — не выкидывали)

- **Протоколы:** VLESS, VMess, Trojan, SS, WireGuard, Hysteria2, HTTP, SOCKS, Tunnel, TUN, MTProto
- **Транспорт:** TCP, mKCP, WS, gRPC, HTTPUpgrade, **XHTTP** + TLS / REALITY / Vision / ML-KEM
- **Клиенты:** лимиты трафика, срок, IP, HWID, онлайн, QR, подписки
- **Ноды, routing, WARP**, балансировщики, свой sub-сервер (Clash / JSON / raw)
- **Telegram-бот**, API + Swagger, SQLite или Postgres, 13 языков, Fail2ban

---

## Быстрый старт — завести точку

Один VPS, root, Ubuntu/Debian — по классике:

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

**Стабильный релиз** (рекомендуем, без сюрпризов):

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) v3.7.0-decoder
```

**Dev-сборка** — что в `main`, то и на сервере (`dev-latest`):

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) dev-latest
```

После установки — `x-ui`: логин, SSL, рестарт, обновление. Панель сама накидает логин/пароль/путь, если не указал.

### Панель уже стоит — только Xray перекатить

```bash
x-ui install-xray
```

Или руками — в [`INSTALL-decoder-dev.ru.md`](INSTALL-decoder-dev.ru.md), там без воды.

### Cloud-init / без диалогов

Когда терминала нет, а VPS должен подняться сам:

```bash
export XUI_NONINTERACTIVE=1
export XUI_USERNAME=admin
export XUI_PASSWORD='YourStrongPassword'
export XUI_PANEL_PORT=54321
export XUI_WEB_BASE_PATH='your-secret-path'
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Креды лежат в `/etc/x-ui/install-result.env` — сохрани, потом не ищи по логам.

---

## Два репозитория — не перепутай

| Что | Куда | Тег для примера |
| --- | --- | --- |
| Панель **3хуя** | [decoder-dev/3x-ui](https://github.com/decoder-dev/3x-ui) | `v3.7.0-decoder`, `dev-latest` |
| Ядро Xray | [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core) | `v26.7.28-decoder` |

Upstream-вики по общим темам: [MHSanaei/3x-ui Wiki](https://github.com/MHSanaei/3x-ui/wiki) — половина статей всё ещё актуальна.

---

## Платформы и база

**ОС:** Linux (Ubuntu, Debian, RHEL-семейство, Arch, Alpine…) и Windows — как у upstream.

**Архи:** amd64, arm64, armv7 и остальное из списка upstream.

**БД:** SQLite из коробки (`/etc/x-ui/x-ui.db`) или Postgres через `XUI_DB_TYPE` / `XUI_DB_DSN` в `/etc/default/x-ui`.

---

## Обновление — не ломай расстановку

```bash
x-ui update          # панель с релизов decoder-dev/3x-ui
x-ui install-xray    # свежий decoder-dev Xray в bin/
```

После апдейта панели Xray подтянется сам, если ставил через наш install/update — но `install-xray` не помешает, если руками ковырял.

---

## Откуда ноги растут

- [MHSanaei/3x-ui](https://github.com/MHSanaei/3x-ui) — базовая панель, уважение
- [XTLS/Xray-core](https://github.com/XTLS/Xray-core) — ядро; наши патчи — [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core)
- [alireza0](https://github.com/alireza0/) — lineage X-UI

**GPL-3.0** — как у предков. Форк open-source, не «закрытая контора».

---

<p align="center"><strong>3хуя — когда панель должна знать, что такое variant B, а не притворяться обычным 3X-UI.</strong></p>
