#!/usr/bin/bash
declare -a arr1
declare -a arr2

while read -r line; do
    arr1+=($(awk '{print $1}' <<< $line))
    arr2+=($(awk '{print $2}' <<< $line))
done < inputs.txt

IFS=$'\n'
arr1_sorted=($(sort -n <<< "${arr1[*]}"))
arr2_sorted=($(sort -n <<< "${arr2[*]}"))
unset IFS

difference=0
for i in ${!arr1_sorted[@]}; do
    temp=$(( ${arr1_sorted[i]} - ${arr2_sorted[i]}))
    difference=$(( $difference + ${temp#-} ))
done

echo "Result: $difference"
