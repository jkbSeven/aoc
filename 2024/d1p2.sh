#!/usr/bin/bash
declare -a arr
declare -A hashmap

while read -r line; do
    arr+=($(awk '{print $1}' <<< $line))

    key=$(awk '{print $2}' <<< $line)
    current_count="${hashmap[$key]}"

    if [[ -z $current_count ]]; then
        current_count=0
    fi

    hashmap["$key"]=$(( $current_count + 1 ))

done < inputs.txt

IFS=$'\n'
arr1_uniq=($(sort -n <<< "${arr1[*]}" | uniq))
unset IFS

similarity_score=0
for id in "${arr[@]}"; do
    similarity_score=$(( $similarity_score + $id * hashmap[$id] ))
done

echo "Result: $similarity_score"
