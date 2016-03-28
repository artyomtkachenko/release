#!/usr/bin/env bash
SELF=`basename $0`
curl -X POST -F "data=@${HOME}/Downloads/apache-activemq-5.12.2-bin.zip" \
             -F "config=@without_scripts.json" \
             -F "post=@${SELF}" \
             -F "preun=@${SELF}" \
             -F 'meta={"version": "5.12.2", "revision": "de1b85fbf4f855f48c6362ea26401e46" }' \
             localhost:8080/release/v1/build/rpm
