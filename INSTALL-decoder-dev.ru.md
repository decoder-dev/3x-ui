# Установка decoder-dev: 3x-ui + Xray (variant B)

Эталонный стек для CDN → nginx → XHTTP packet-up с **stream-down keepalive**.

| Компонент | Репозиторий | Релиз |
|-----------|-------------|-------|
| Панель | [decoder-dev/3x-ui](https://github.com/decoder-dev/3x-ui) | [v3.7.0-decoder](https://github.com/decoder-dev/3x-ui/releases/tag/v3.7.0-decoder) |
| Xray | [decoder-dev/Xray-core](https://github.com/decoder-dev/Xray-core) | [v26.7.28-decoder](https://github.com/decoder-dev/Xray-core/releases/tag/v26.7.28-decoder) |

---

## 1. Новый сервер — установка 3x-ui (форк)

**Требования:** root, Ubuntu/Debian 22.04+, порт 80 открыт (для Let's Encrypt при первой настройке).

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Скрипт:
- скачает **decoder-dev/3x-ui** (не upstream MHSanaei);
- по умолчанию ставит latest release **`v3.7.0-decoder`**;
- создаст systemd-сервис `x-ui`, команду `x-ui`, случайные логин/пароль/порт.

### Конкретная версия

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) v3.7.0-decoder
```

### Dev-сборка (rolling, каждый push в main)

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh) dev-latest
```

### Non-interactive (cloud-init / VPS)

```bash
export XUI_NONINTERACTIVE=1
export XUI_USERNAME=admin
export XUI_PASSWORD='YourStrongPassword'
export XUI_PANEL_PORT=54321
export XUI_WEB_BASE_PATH='your-secret-path'
export XUI_SSL_MODE=none          # за reverse proxy (nginx/Caddy)
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)
```

Учётные данные сохраняются в `/etc/x-ui/install-result.env`.

---

## 2. Заменить Xray на decoder-dev (variant B)

Если панель уже установлена (в т.ч. upstream 3x-ui), замените бинарник Xray:

```bash
XRAY_ZIP=https://github.com/decoder-dev/Xray-core/releases/download/v26.7.28-decoder/Xray-linux-64.zip
curl -fsSL -o /tmp/xray-decoder.zip "$XRAY_ZIP"
systemctl stop x-ui

cp -a /usr/local/x-ui/bin/xray-linux-amd64 /usr/local/x-ui/bin/xray-linux-amd64.bak.$(date +%F)
unzip -o /tmp/xray-decoder.zip -d /tmp/xray-decoder
install -m 755 /tmp/xray-decoder/xray /usr/local/x-ui/bin/xray-linux-amd64

/usr/local/x-ui/bin/xray-linux-amd64 version
# Ожидается: ... 3819e92 ... и наличие scStreamDown в strings

x-ui restart
```

Проверка variant B:

```bash
strings /usr/local/x-ui/bin/xray-linux-amd64 | grep -E 'scStreamDown|downFrame' | head -3
```

---

## 3. Обновление форка 3x-ui на существующем сервере

**Не используйте `install.sh` повторно** — он переустановит панель с нуля.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/update.sh)
```

Или из меню: `x-ui` → пункт Update.

Dev-канал:

```bash
XUI_UPDATE_TAG=dev-latest bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/update.sh)
```

---

## 4. Inbound variant B (XHTTP packet-up + stream-down)

После установки панели настройте inbound **VLESS + XHTTP**:

| Параметр | Значение |
|----------|----------|
| listen | `127.0.0.1` (за nginx) |
| port | `8080` |
| mode | `packet-up` |
| scStreamDownServerSecs | `15-45` |
| scStreamUpServerSecs | `20-80` |
| downFrame | ✅ (в форке 3x-ui — поля в панели) |

Перед nginx/CDN: `127.0.0.1:8080` → `cdn.example.com:443`.

Скрипты эталона (из репозитория deploy):

```bash
# Восстановить inbound из snapshot.json
python3 deploy/etalon/setup-inbound.py

# Применить sockopt trustedXForwardedFor
python3 deploy/etalon/fix-inbound-sockopt.py
```

---

## 5. ML-KEM (post-quantum, опционально)

Требует Xray с VLESS Encryption (есть в **decoder-dev/Xray-core**).

В Telegram-боте (если задеплоен): **⚙️ Ядро → 🔐 ML-KEM → включить → Применить**.

Или в панели 3x-ui: inbound VLESS → сгенерировать **ML-KEM-768** (кнопка vlessenc).

После включения — обновить подписку у клиентов (Happ).

---

## 6. Что входит в форк 3x-ui (v3.7.0-decoder)

- Поля **scStreamDownServerSecs** / **downFrame** в UI и подписке
- Адаптивная вёрстка (mobile)
- Fix **tgId** JSON unmarshal
- `install.sh` / `update.sh` → **decoder-dev/3x-ui**

---

## 7. Сборка релизов (для maintainer)

### 3x-ui

```bash
gh workflow run "Release 3X-UI" --repo decoder-dev/3x-ui --ref main
# или push в main → dev-latest; тег v3.7.x-decoder → stable release
```

### Xray-core

```bash
gh workflow run "Build and Release" --repo decoder-dev/Xray-core --ref main
# linux-amd64: Xray-linux-64.zip в release v26.7.28-decoder
```

---

## 8. Быстрые ссылки

```bash
# Установка панели
bash <(curl -Ls https://raw.githubusercontent.com/decoder-dev/3x-ui/main/install.sh)

# Только Xray amd64
curl -fsSL -o /tmp/x.zip https://github.com/decoder-dev/Xray-core/releases/download/v26.7.28-decoder/Xray-linux-64.zip
```
