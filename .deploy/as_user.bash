#!/bin/bash

readonly ME=$(whoami)
touch $HOME/.profile

function set_env {
    export PATH="/usr/local/go/bin":$PATH
    export GOROOT="/usr/local/go"
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-api"
    export GOPATH="/opt/gopath"
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    echo "export PATH=/usr/local/go/bin:\$PATH"
    echo "export GOROOT=/usr/local/go"
    echo "export GOPATH=/opt/gopath" >> $HOME/.profile
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
    echo "cd $APPDIR" >> $HOME/.bashrc
    source $HOME/.profile
}

function install_node_modules {
    cd $GOPATH/src/github.com/alligrader/gradebook
    npm install
    cd -
}

function install_godeps {
    sudo chown -R $ME:$ME /opt  # Work around bug in Vagrant
    go get github.com/snikch/goodman
    go get bitbucket.org/liamstask/goose/cmd/goose
    go get github.com/Masterminds/glide
    go get -u github.com/goadesign/goa/...
    cd $APPDIR
    glide i
    cd -
}

function install_autoenv {
    echo '~~ installing autoenv ~~'
    git clone git://github.com/kennethreitz/autoenv.git $HOME/.autoenv
    echo '~~ cloning autoenv completed ~~'
    echo '~~ persisting autoenv to the .profile ~~'
    echo 'source ~/.autoenv/activate.sh' >> $HOME/.profile
    echo '~~ finshed installing autoenv ~~'
}

function install_spacecowboy {
    wget https://gist.githubusercontent.com/seanlinsley/3e49c09dcdfc37f05eb4/raw/2e61c48fef7adfdf7267199bc39d2ab5bf38e9b7/seeyouspacecowboy.rb
    mv seeyouspacecowboy.rb $HOME/.spacecowboy.rb
    echo 'ruby ~/.spacecowboy.rb' >> $HOME/.bash_logout
}

function make_storage {
    mkdir $HOME/storage
}

function configure_kubectl {
    echo "export KUBECONFIG=${KUBECONFIG}:$APPDIR/.deploy/kubeconfig" >> $HOME/.profile
    echo "kubectl config use-context vagrant-single"                  >> $HOME/.profile
}

function main {
    set_env
    install_node_modules
    install_autoenv

    # This command breaks Docker. God love us.
    # install_spacecowboy

    install_godeps
    make_storage
    configure_kubectl
}

set -e
set -x
set -u
main
echo '~~ Provisioning User-Mode Complete ~~'
