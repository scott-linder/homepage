#!/bin/bash

#
# Deploy script
# Authors:
#   Scott Linder
#

set -e

cd `dirname $0`

SRC=site/

function deploy {
    USER=$1
    HOST=$2
    DEST=$3
    PUBLINK=$4
    # (a)rchive with (z) compression; print (h)uman-readable; delete extraneous
    rsync -azh --delete --progress $SRC $USER@$HOST:$DEST
    # Update public/ symlink on remote host
    ssh -q $USER@$HOST <<- EOL
        [[ -h $PUBLINK ]] && rm $PUBLINK
        ln -s ${DEST}public $PUBLINK
EOL
}

case $1 in
    cclub)
        deploy 'majorstringy' 'yakko.cs.wmich.edu' '~/homepage/' '~/www'
        ;;
    cs)
        deploy 'smn2028' 'login.cs.wmich.edu' '~/homepage/' '~/www'
        ;;
    *)
        echo 'Unknown target.'
esac

