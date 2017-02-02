#!/bin/bash

readonly ME=$(whoami)
readonly DIST='vivid'

function update {
    echo '~~ Updating apt-get ~~'
    sudo apt-add-repository "deb http://download.virtualbox.org/virtualbox/debian $DIST contrib"
    wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -
    sudo apt-get update
}

function install_deps {
    echo '~~ Installing via apt-get ~~'
    sudo apt-get install -yq \
        git \
        mysql-client \
        curl \
        wget \
        zip \
        unzip \
        vim-nox \
        mercurial \
        npm
}

function make_gopath {
    echo '~~ Preparing the gopath ~~'
    mkdir -p "/opt/gopath/src" || 0
    mkdir -p "/opt/gopath/bin" || 0
    mkdir -p "/opt/gopath/pkg" || 0
}

function install_go {

    local readonly GO_BIN_VERSION="go1.8beta2"
    local readonly URL="storage.googleapis.com/golang/${GO_BIN_VERSION}.linux-amd64.tar.gz"

    echo "~~ Installing $GO_BIN_VERSION ~~"
    if [ ! -d /usr/local/go ]; then
        # curl -O -J -L "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
        wget -q "https://$URL"
        sudo tar -C /usr/local -xvf "${GO_BIN_VERSION}.linux-amd64.tar.gz"
        rm "${GO_BIN_VERSION}.linux-amd64.tar.gz"
    fi
    echo "~~ $GO_BIN_VERSION Installation complete ~~"
}

function install_packer {
    local readonly VERSION="0.12.1"
    local readonly PACKER_VERSION="packer_${VERSION}"
    local readonly URL="releases.hashicorp.com/packer/${VERSION}/${PACKER_VERSION}_linux_amd64.zip"
    echo "~~ Installing ${PACKER_VERSION} ~~"

    if [ ! -d /usr/local/packer ]; then
        wget -q "https://$URL"
        sudo unzip -d /usr/local/bin "${PACKER_VERSION}_linux_amd64.zip"
        rm "${PACKER_VERSION}_linux_amd64.zip"
    fi
    echo "~~ ${PACKER_VERSION} installed ~~"
}

function install_dredd {

    echo '~~ Updating NPM to latest ~~'
    sudo npm install -g npm@latest
    echo '~~ Linking nodejs to /usr/bin/node ~~'
    sudo ln -s $(which nodejs) /usr/bin/node
    echo '~~ Installing dredd ~~'
    sudo npm install -g dredd@2.2.5
    echo '~~ dredd installed successfully ~~'
    sudo npm install -g swagger2blueprint
    echo '~~ swagger2blueprint installed successfully ~~'
}

function install_kubectl {
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    sudo mv ./kubectl /usr/local/bin/kubectl
}

function main {
    update
    install_deps
    make_gopath
    install_go
    install_packer
    install_dredd
    install_kubectl
}

set -e
set -x
set -u
main
echo '~~ Root Provisioning Complete ~~'
