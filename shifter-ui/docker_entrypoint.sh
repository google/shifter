#!/bin/bash
export FILE="/code/shifter.server.endpoint"
export EXISTING_VARS=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,); 

if [ -f "$FILE" ]; then
  PREVIOUS_ENDPOINT=$(cat ${FILE})
  if [ "$PREVIOUS_ENDPOINT" = "$SHIFTER_SERVER_ENDPOINT" ]; then
    echo "Note: shifter.server.endpoint is Unchanged."
  else
    echo "Note: shifter.server.endpoint has Changed."
    echo "---> Previous Endpoint: ${PREVIOUS_ENDPOINT}"
    echo "---> Updated  Endpoint: ${SHIFTER_SERVER_ENDPOINT}"
    rm -rf /code/assets
    cp -r /code/assets_original /code/assets
    rm -rf $FILE
  fi
else
  echo "Note: This is an Initial Run."
fi

if [ ! -e "$FILE" ]; then
    ########################################
    # Replace all References to Environment Variables  
    # within JS Files.
    ########################################
    for jsFile in $JSFOLDER;
    do
      
      cat $jsFile | envsubst $EXISTING_VARS | tee $jsFile

    done
    echo $SHIFTER_SERVER_ENDPOINT > $FILE
    echo "------> Initial Boot."
    cat $FILE
    echo $SHIFTER_SERVER_ENDPOINT
else 
    echo "------> Secondry Boot."
    cat $FILE 
    echo $SHIFTER_SERVER_ENDPOINT
fi

########################################
# Start NGINX Server
########################################
nginx