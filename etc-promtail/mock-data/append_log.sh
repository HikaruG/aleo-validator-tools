#!/bin/bash

old_log="mock-data.log"
new_log="mock-aleo.log"

while IFS= read -r line
do
  echo "$line" >> "$new_log"
  #sleep 0.01  # Sleep for 1 second (or more) between lines
done < "$old_log"
