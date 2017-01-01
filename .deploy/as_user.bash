#!/bin/bash

readonly ME=$(whoami)
touch $HOME/.profile

function set_env {
    export PATH="/usr/local/go/bin":$PATH
    export GOROOT="/usr/local/go"
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-backend"
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
    cd $APPDIR
    glide i
    cd -
}

function install_autoenv {
    git clone git://github.com/kennethreitz/autoenv.git ~/.autoenv
    echo 'source ~/.autoenv/activate.sh' >> ~/.profile
    source ~/.profile
}

function install_spacecowboy {
    wget https://gist.githubusercontent.com/seanlinsley/3e49c09dcdfc37f05eb4/raw/2e61c48fef7adfdf7267199bc39d2ab5bf38e9b7/seeyouspacecowboy.rb
    mv seeyouspacecowboy.rb $HOME/.spacecowboy.rb
    echo 'ruby ~/.spacecowboy.rb' >> $HOME/.bash_logout
}

function make_storage {
    mkdir $HOME/storage
}

function main {
    set_env
    install_autoenv
    install_spacecowboy
    install_godeps
    make_storage
}

main