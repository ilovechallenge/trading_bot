#!/bin/bash
set -e

target=$1
bin_dir=bin
bin_type=bbgo-slim
host=bbgo
host_user=root
host_home=/root
host_systemd_service_dir=/etc/systemd/system
os=linux
arch=amd64
setup_host_systemd_service=no

# use the git describe as the binary version, you may override this with something else.
tag=$(git describe --tags)

RED='\033[1;31m'
GREEN='\033[1;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

function remote_test() {
  ssh $host "(test $* && echo yes)"
}

function remote_run() {
  ssh $host "$*"
}

function warn() {
  echo "${YELLOW}$*${NC}"
}

function error() {
  echo "${RED}$*${NC}"
}

function info() {
  echo "${GREEN}$@${NC}"
}

if [[ -n $BBGO_HOST ]]; then
  host=$BBGO_HOST
else
  warn "env var BBGO_HOST is not set, using \"bbgo\" host alias as the default host, you can add \"bbgo\" to your ~/.ssh/config file"
  host=bbgo
fi

if [[ -z $target ]]; then
  echo "Usage: $0 [target]"
  echo "target name is required"
  exit 1
fi

# initialize the remote environment
# create the directory for placing binaries
ssh $host "mkdir -p \$HOME/$bin_dir && mkdir -p \$HOME/$target"

if [[ $(remote_test "-e $host_systemd_service_dir/$target.service") != "yes" ]]; then
  if [[ "$setup_host_systemd_service" == "no" ]]; then
    error "The systemd $target.service on host $host is not configured, can not deploy"
    exit 1
  fi

  warn "$host_systemd_service_dir/$target.service does not exist, setting up..."

  cat <<END >".systemd.$target.service"
[Unit]
After=network-online.target
Wants=network-online.target

[Install]
WantedBy=multi-user.target

[Service]
WorkingDirectory=$host_home/$target
KillMode=process
ExecStart=$host_home/$target/bbgo run
User=$host_user
Restart=always
RestartSec=30
END

  info "uploading systemd service file..."
  scp ".systemd.$target.service" "$host:$host_systemd_service_dir/$target.service"

  info "reloading systemd daemon..."
  remote_run "systemctl daemon-reload && systemctl enable $target"
fi

info "building binary: $bin_type-$os-$arch..."
make $bin_type-$os-$arch

# copy the binary to the server
info "deploying..."
info "copying binary to host $host..."
scp build/bbgo/$bin_type-$os-$arch $host:$bin_dir/bbgo-$tag

# link binary and restart the systemd service
info "linking binary and restarting..."
ssh $host "(cd $target && ln -sf \$HOME/$bin_dir/bbgo-$tag bbgo && systemctl restart $target.service)"

info "deployed successfully!"
