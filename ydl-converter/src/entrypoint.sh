#!/bin/bash

echo "> curl -s ${REST_URL%/}/pop?type=cnv | jq '.'"
JSON=`curl -s "${REST_URL%/}/pop?type=cnv" | jq '.'`
[[ $JSON == "" ]] && exit 0

FILENAME=`echo "${JSON}" | jq '.url' | sed 's/"//g'`
FILENAME=`echo ${FILENAME} | base64 --decode`
FORMAT=`echo "${JSON}" | jq '.fmt' | sed 's/"//g'`
CRF=`echo "${JSON}" | jq '.crf' | sed 's/"//g'`
echo "FILENAME=${FILENAME}"
[[ "${FILENAME}" == "null" ]] && exit 0

echo "------ task received ------"
[[ -e /convert/"${FILENAME}.${FORMAT}" ]] && rm -v "/convert/${FILENAME}.${FORMAT}"

echo "> ffmpeg -i /convert/${FILENAME}.video -i /convert/${FILENAME}.audio -c:v copy -c:a copy -map 0:v:0 -map 1:a:0 -f ${FORMAT} /convert/${FILENAME}.${FORMAT}"
ffmpeg -i "/convert/${FILENAME}.video" -i "/convert/${FILENAME}.audio" -c:v copy -c:a copy -map 0:v:0 -map 1:a:0 -f ${FORMAT} "/convert/${FILENAME}.${FORMAT}"
[[ ! -e "/convert/${FILENAME}.${FORMAT}" ]] && exit 1

[[ -e /convert/"${FILENAME}.video" ]] && rm -v "/convert/${FILENAME}.video"
[[ -e /convert/"${FILENAME}.audio" ]] && rm -v "/convert/${FILENAME}.audio"

OWNER=`ls -ld /convert/ | awk '{ print $3 }'`
chown ${OWNER}:${OWNER} /convert/"${FILENAME}.${FORMAT}"

echo "> ls -lh /convert/ | grep -F ${FILENAME}"
ls -lh /convert/ | grep -F "${FILENAME}"

[[ $? -ne 0 ]] && exit 1
echo "------ task completed ------"
exit 0