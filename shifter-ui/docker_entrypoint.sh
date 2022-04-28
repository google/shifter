#!/bin/bash
export EXISTING_VARS=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,); 

########################################
# Load Fixed Original Build Assets
########################################
rm -rf /code/assets
cp -r /code/assets_original /code/assets

########################################
# Update Dynamic Environment Varaibles
########################################
for jsFile in $JSFOLDER;
do
    cat $jsFile | envsubst $EXISTING_VARS | tee "${jsFile}.tmp"
    rm $jsFile
    cp "${jsFile}.tmp" $jsFile
    rm "${jsFile}.tmp"
done

########################################
# Start NGINX Server
########################################
nginx