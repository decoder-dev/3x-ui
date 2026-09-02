#!/bin/bash
# Install / refresh decoder-dev/Xray-core into 3x-ui bin/ (variant B + ML-KEM).
# Sourced from install.sh, update.sh, and x-ui.sh.

XRAY_GITHUB_REPO="${XRAY_GITHUB_REPO:-decoder-dev/Xray-core}"
XRAY_RELEASE_TAG="${XRAY_RELEASE_TAG:-}"

_decoder_xray_zip_name() {
    case "$1" in
        amd64) echo "Xray-linux-64.zip" ;;
        arm64) echo "Xray-linux-arm64-v8a.zip" ;;
        armv7) echo "Xray-linux-arm32-v7a.zip" ;;
        armv6) echo "Xray-linux-arm32-v6.zip" ;;
        armv5) echo "Xray-linux-arm32-v5.zip" ;;
        386) echo "Xray-linux-32.zip" ;;
        s390x) echo "Xray-linux-s390x.zip" ;;
        *) return 1 ;;
    esac
}

_decoder_xray_bin_name() {
    case "$1" in
        amd64) echo "xray-linux-amd64" ;;
        arm64) echo "xray-linux-arm64" ;;
        armv7 | armv6 | armv5) echo "xray-linux-arm32" ;;
        386) echo "xray-linux-386" ;;
        s390x) echo "xray-linux-s390x" ;;
        *) return 1 ;;
    esac
}

_decoder_xray_latest_tag() {
    local repo="$1"
    local tag=""
    if command -v curl >/dev/null 2>&1; then
        tag=$(
            curl -fsSL --retry 3 --connect-timeout 15 --max-time 60 \
                "https://api.github.com/repos/${repo}/releases?per_page=15" 2>/dev/null \
                | grep -oE '"tag_name"[[:space:]]*:[[:space:]]*"v[^"]+-decoder"' \
                | head -1 \
                | sed -E 's/.*"([^"]+)".*/\1/'
        )
    fi
    if [[ -z "$tag" ]]; then
        tag="v26.7.28-decoder"
    fi
    echo "$tag"
}

# install_decoder_xray [x-ui_dir] [arch]
# Downloads decoder-dev Xray and replaces bin/xray-linux-*.
install_decoder_xray() {
    local xui_dir="${1:-/usr/local/x-ui}"
    local cpu_arch="${2:-$(arch 2>/dev/null || true)}"
    if [[ -z "$cpu_arch" ]]; then
        case "$(uname -m)" in
            x86_64 | x64 | amd64) cpu_arch="amd64" ;;
            aarch64 | arm64) cpu_arch="arm64" ;;
            armv7*) cpu_arch="armv7" ;;
            armv6*) cpu_arch="armv6" ;;
            armv5*) cpu_arch="armv5" ;;
            i*86 | x86) cpu_arch="386" ;;
            s390x) cpu_arch="s390x" ;;
            *) echo "install_decoder_xray: unsupported arch $(uname -m)" >&2; return 1 ;;
        esac
    fi

    local zip_name bin_name tag repo url tmp
    zip_name=$(_decoder_xray_zip_name "$cpu_arch") || {
        echo "install_decoder_xray: no zip mapping for arch ${cpu_arch}" >&2
        return 1
    }
    bin_name=$(_decoder_xray_bin_name "$cpu_arch") || {
        echo "install_decoder_xray: no bin mapping for arch ${cpu_arch}" >&2
        return 1
    }

    repo="${XRAY_GITHUB_REPO}"
    tag="${XRAY_RELEASE_TAG:-$(_decoder_xray_latest_tag "$repo")}"
    url="https://github.com/${repo}/releases/download/${tag}/${zip_name}"

    if [[ ! -d "${xui_dir}/bin" ]]; then
        mkdir -p "${xui_dir}/bin"
    fi

    tmp=$(mktemp -d)
    trap 'rm -rf "$tmp"' RETURN

    echo -e "${green}Installing decoder-dev Xray ${tag} (${zip_name})...${plain}"
    if ! curl -fL --retry 5 --retry-delay 3 --connect-timeout 20 --max-time 600 \
        -o "${tmp}/xray.zip" "$url"; then
        echo -e "${red}Failed to download ${url}${plain}" >&2
        return 1
    fi

    if ! unzip -qo "${tmp}/xray.zip" -d "${tmp}/extract"; then
        echo -e "${red}Failed to unzip decoder Xray archive${plain}" >&2
        return 1
    fi

    if [[ ! -f "${tmp}/extract/xray" ]]; then
        echo -e "${red}Archive missing xray binary${plain}" >&2
        return 1
    fi

    local dest="${xui_dir}/bin/${bin_name}"
    if [[ -f "$dest" ]]; then
        cp -a "$dest" "${dest}.bak.$(date +%Y%m%d-%H%M%S)" 2>/dev/null || true
    fi
    install -m 755 "${tmp}/extract/xray" "$dest"
    echo -e "${green}decoder-dev Xray -> ${dest}${plain}"
    "$dest" version 2>/dev/null | head -1 || true
    return 0
}
