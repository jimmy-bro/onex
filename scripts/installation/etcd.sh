#!/bin/bash

# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/superproj/onex.
#


# The root of the build/dist directory.
ONEX_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
# If common.sh has already been sourced, it will not be sourced again here.
[[ -z ${COMMON_SOURCED} ]] && source ${ONEX_ROOT}/scripts/installation/common.sh
# Set some environment variables.
ONEX_ETCD_HOST=${ONEX_ETCD_HOST:-127.0.0.1}
ONEX_ETCD_CLIENT_PORT=${ONEX_ETCD_CLIENT_PORT:-2379}
ONEX_ETCD_PEER_PORT=${ONEX_ETCD_PEER_PORT:-2380}
ONEX_ETCD_VERSION=${ETCD_VERSION:-3.5.11}
ONEX_ETCD_DIR=${ONEX_ETCD_DIR:-${ONEX_THIRDPARTY_INSTALL_DIR}/etcd}

# Install etcd using containerization.
function onex::etcd::docker::install()
{
  onex::common::network
  docker run -d --name onex-etcd \
    --network onex \
    -v ${ONEX_ETCD_DIR}:/etcd-data \
    -p ${ONEX_ACCESS_HOST}:${ONEX_ETCD_CLIENT_PORT}:2379 \
    -p ${ONEX_ACCESS_HOST}:${ONEX_ETCD_PEER_PORT}:2380 \
    quay.io/coreos/etcd:latest \
    /usr/local/bin/etcd \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-client-urls http://0.0.0.0:2379 \
    --data-dir /etcd-data

  sleep 10
  onex::etcd::status || return 1
  onex::etcd::info
  onex::log::info "install etcd successfully"
}

# Uninstall the docker container.
function onex::etcd::docker::uninstall()
{
  docker rm -f onex-etcd &>/dev/null
  onex::util::sudo "rm -rf ${ONEX_ETCD_DIR}"
  onex::log::info "uninstall etcd successfully"
}

# Install the etcd step by step.
# sbs is the abbreviation for "step by step".
function onex::etcd::sbs::install()
{
  local os
  local arch

  os=$(onex::util::host_os)
  arch=$(onex::util::host_arch)

  download_file_name="etcd-v${ONEX_ETCD_VERSION}-${os}-${arch}"
  download_file="/tmp/${download_file_name}.tar.gz"
  url="https://github.com/coreos/etcd/releases/download/v${ONEX_ETCD_VERSION}/${download_file_name}.tar.gz"
  onex::util::download_file "${url}" "${download_file}"
  tar -xvzf "${download_file}" -C /tmp/
  echo ${LINUX_PASSWORD} | sudo -S cp /tmp/${download_file_name}/{etcd,etcdctl,etcdutl} /usr/bin/
  rm "${download_file}"
  rm -rf /tmp/${download_file_name}

  # 创建 Etcd 配置文件
  onex::util::sudo "mkdir -p /etc/etcd"

  echo ${LINUX_PASSWORD} | sudo -S cat << EOF | sudo tee /etc/etcd/config.yaml
name: onex
data-dir: ${ONEX_ETCD_DIR}
advertise-client-urls: http://0.0.0.0:${ONEX_ETCD_CLIENT_PORT}
listen-client-urls: http://0.0.0.0:${ONEX_ETCD_CLIENT_PORT}
initial-advertise-peer-urls: http://0.0.0.0:${ONEX_ETCD_PEER_PORT}
initial-cluster: onex=http://0.0.0.0:${ONEX_ETCD_PEER_PORT}
log-outputs: [${ONEX_LOG_DIR}/etcd.log]
log-level: debug
EOF

  echo ${LINUX_PASSWORD} | sudo -S cat << EOF | sudo tee /lib/systemd/system/etcd.service
# Etcd systemd unit template from
# https://github.com/etcd-io/etcd/blob/main/contrib/systemd/etcd.service
[Unit]
Description=etcd key-value store
Documentation=https://github.com/etcd-io/etcd
After=network-online.target local-fs.target remote-fs.target time-sync.target
Wants=network-online.target local-fs.target remote-fs.target time-sync.target

[Service]
Type=notify
ExecStart=/usr/bin/etcd --config-file=/etc/etcd/config.yaml
Restart=always
RestartSec=10s
LimitNOFILE=40000

[Install]
WantedBy=multi-user.target
EOF

  onex::util::sudo "systemctl daemon-reload"
  onex::util::sudo "systemctl start etcd"
  onex::util::sudo "systemctl enable etcd"

  onex::etcd::status || return 1
  onex::etcd::info
  onex::log::info "install etcd v${ONEX_ETCD_VERSION} successfully"
}

# Uninstall the etcd step by step.
function onex::etcd::sbs::uninstall()
{
  #set +o errexit
  #etcd_pids=$(pgrep -f $HOME/bin/etcd)
  #set -o errexit
  #if [[ ${etcd_pids} != "" ]];then
    #echo ${LINUX_PASSWORD} | sudo -S kill -9 ${etcd_pids} || true
  #fi

  # etcd, etcdctl, etcdutl 3 个文件这里不需要删除，因为下次安装时，文件会被覆盖，仍然是幂等安装
  onex::util::sudo "systemctl stop etcd"
  onex::util::sudo "rm -rf ${ONEX_ETCD_DIR}"
  onex::log::info "uninstall etcd successfully"
}

# Print necessary information after docker or sbs installation.
function onex::etcd::info()
{
  echo -e ${C_GREEN}etcd has been installed, here are some useful information:${C_NORMAL}
  cat << EOF | sed 's/^/  /'
Etcd endpoint is: ${ONEX_ETCD_HOST}:${ONEX_ETCD_CLIENT_PORT}
EOF
}

# Status check after docker or sbs installation.
function onex::etcd::status()
{
  onex::util::telnet ${ONEX_ETCD_HOST} ${ONEX_ETCD_CLIENT_PORT} || return 1

  #echo "Waiting for etcd to come up."
  #onex::util::wait_for_url "http://${ONEX_ETCD_HOST}:${ONEX_ETCD_CLIENT_PORT}/health" "etcd: " 0.25 80
}

if [[ "$*" =~ onex::etcd:: ]];then
  eval $*
fi
