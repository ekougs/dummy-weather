#!/bin/bash

cur_dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
gopaths=$GOPATH
delimiter=":"
relevant_gopath=""

if [[ "$gopaths" != *"$delimiter"* ]]
then
	echo $gopaths
	exit 0
fi

while [[ "$gopaths" == *"$delimiter"* ]]
do
	current_gopath=${gopaths#*${delimiter}}
	if [[ "$cur_dir" == *"$current_gopath"* ]]
	then
		relevant_gopath=$current_gopath
	fi
	gopaths=${gopaths%${delimiter}*}
done

if [[ "$cur_dir" == *"gopaths"* ]]
then
	relevant_gopath=$gopaths
fi

if [[ $relevant_gopath == "" ]]
then
	echo "No relevant go path found. Check if GOPATH '$GOPATH' contains ancestor of the current repository '$cur_dir'."
    exit 1
fi

echo $relevant_gopath
