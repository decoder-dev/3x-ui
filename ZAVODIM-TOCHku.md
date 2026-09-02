# Заводим точку: 3хуя + Xray variant B

**По понятиям зэков.** Эталонная расстановка: **CDN → nginx → XHTTP packet-up**, stream-down keepalive включён, без «поставил upstream и удивился, почему не едет».

| Что ставим | Репозиторий | Тег |
| --- | --- | --- |
| Панель **3хуя** | [decoder-dev/3x-ui](https://github.com/decoder-dev/3x-ui) | [v3.7.0-decoder](https://github.com/decoder-dev/3x-ui/releases/tag/v3.7.0-decoder) |
| Ядро Xray | [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core) | [v26.7.28-decoder](https://github.com/decoder-dev/Xray-core/releases/tag/v26.7.28-decoder) |

---

## 1. Чистый VPS — поднимаем контору

**На входе:** root, Ubuntu/Debian 22.04+, порт 80 открыт (Let's Encrypt, если SSL на самой машине).

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Скрипт сам:
- тянет **decoder-dev/3x-ui**, не MHSanaei;
- по умолчанию ставит **`v3.7.0-decoder`**;
- заводит `x-ui` в systemd, кидает случайные логин / пароль / путь.

### Конкретный релиз — без лотереи

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) v3.7.0-decoder
```

### Dev — свежак с `main` (`dev-latest`)

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) dev-latest
```

### Без диалогов (cloud-init, когда руками некогда)

```bash
export XUI_NONINTERACTIVE=1
export XUI_USERNAME=admin
export XUI_PASSWORD='YourStrongPassword'
export XUI_PANEL_PORT=54321
export XUI_WEB_BASE_PATH='your-secret-path'
export XUI_SSL_MODE=none          # SSL на nginx/CDN, не на панели
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Креды — в `/etc/x-ui/install-result.env`. **Сохрани сразу**, потом не копай логи.

---

## 2. Панель уже есть — перекатываем только Xray

Upstream или старый бинарник — выкидываем, ставим decoder-dev **variant B**:

```bash
XRAY_ZIP=https://github.com/decoder-dev/Xray-core/releases/download/v26.7.28-decoder/Xray-linux-64.zip
curl -fsSL -o /tmp/xray-decoder.zip "$XRAY_ZIP"
systemctl stop x-ui

cp -a /usr/local/x-ui/bin/xray-linux-amd64 /usr/local/x-ui/bin/xray-linux-amd64.bak.$(date +%F)
unzip -o /tmp/xray-decoder.zip -d /tmp/xray-decoder
install -m 755 /tmp/xray-decoder/xray /usr/local/x-ui/bin/xray-linux-amd64

/usr/local/x-ui/bin/xray-linux-amd64 version
# Ждём commit вроде b3da218 / 3819e92 — не vanilla XTLS

x-ui restart
```

Или одной командой из меню:

```bash
x-ui install-xray
```

**Проверка, что variant B на месте:**

```bash
strings /usr/local/x-ui/bin/xray-linux-amd64 | grep -E 'scStreamDown|downFrame' | head -3
```

Нет строк — значит не тот Xray, катать так нельзя.

---

## 3. Апдейт панели — не ломай базу

**`install.sh` второй раз не гоняй** — переустановит с нуля, база полетит.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/update.sh)
```

Или: `x-ui` → Update.

Dev-канал:

```bash
XUI_UPDATE_TAG=dev-latest bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/update.sh)
```

После update Xray decoder-dev подтянется сам — но если ковырял руками, добей `x-ui install-xray`.

---

## 4. Inbound variant B — расстановка XHTTP

В панели **3хуя**: inbound **VLESS + XHTTP**:

| Параметр | Значение |
| --- | --- |
| listen | `127.0.0.1` (за nginx, не светить в мир) |
| port | `8080` |
| mode | `packet-up` |
| scStreamDownServerSecs | `15-45` |
| scStreamUpServerSecs | `20-80` |
| downFrame | ✅ (поля есть в UI форка) |

Снаружи: nginx/CDN бьёт на `127.0.0.1:8080`, домен типа `cdn.example.com:443`.

Скрипты эталона (если клонировал deploy):

```bash
python3 deploy/etalon/setup-inbound.py      # inbound из snapshot.json
python3 deploy/etalon/fix-inbound-sockopt.py # sockopt trustedXForwardedFor
```

---

## 5. ML-KEM — post-quantum, по желанию

Нужен **decoder-dev/Xray-core** с VLESS Encryption.

- **Telegram-бот:** ⚙️ Ядро → 🔐 ML-KEM → включить → Применить
- **Панель:** inbound VLESS → сгенерить **ML-KEM-768** (vlessenc)

После включения — **обновить подписку** у клиентов (Happ и прочие). Старые ключи без PQ не прокатят.

> **По понятиям:** при ML-KEM не оставляй пустой `"fallbacks": []` в inbound — Xray ляжет. Форк это чинит в боте/core_config; вручную — следи за конфигом.

---

## 6. Что уже вшито в 3хуя (форк)

- **scStreamDownServerSecs** / **downFrame** — UI, подписка, Clash
- Адаптив под мобилу
- Fix **tgId** (число / строка в JSON)
- install/update → **decoder-dev/3x-ui**, Xray → **decoder-dev/Xray-core**
- Брендинг **3хуя** в морде

---

## 7. Maintainer — собрать релиз

### Панель

```bash
gh workflow run "Release 3X-UI" --repo decoder-dev/3x-ui --ref main
# push в main → dev-latest; тег v3.7.x-decoder → stable
```

### Xray

```bash
gh workflow run "Build and Release" --repo decoder-dev/Xray-core --ref main
# amd64: Xray-linux-64.zip в v26.7.28-decoder
```

---

## 8. Шпаргалка одной строкой

```bash
# Вся контора с нуля
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)

# Только Xray amd64
curl -fsSL -o /tmp/x.zip https://github.com/decoder-dev/Xray-core/releases/download/v26.7.28-decoder/Xray-linux-64.zip
```

---

<p align="center"><strong>Точка заведена — variant B катает, nginx не ноет, клиенты sub тянут. Всё по понятиям.</strong></p>
