#!/bin/bash
set -euo pipefail

TEMPLATE_PATH=$1
CMD_DIRECTORY=$2
PYTHON_INTERPRETER=$3

read -p "What is the puzzle name? " PUZZLE_NAME
read -p "What is the day number? " DAY_NUMERAL
DAY_NUMBER=$($PYTHON_INTERPRETER -c "import inflect; p = inflect.engine(); print(p.number_to_words('$DAY_NUMERAL').replace('-', ''))")

NEW_FILE=$CMD_DIRECTORY/$DAY_NUMERAL.go
cp $TEMPLATE_PATH $NEW_FILE
sed -i "s/{{DAY_NUMBER}}/$DAY_NUMBER/g" $NEW_FILE
sed -i "s/{{DAY_NUMERAL}}/$DAY_NUMERAL/g" $NEW_FILE
sed -i "s/{{PUZZLE_NAME}}/$PUZZLE_NAME/g" $NEW_FILE