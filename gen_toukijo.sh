#!/bin/bash

# 入力ファイル
input="toukijo.csv"

# 出力ファイル
output="const.go"

# ヘッダー
echo "package main" > $output
echo "" >> $output
echo "var ToukijoCodes = []string{" >> $output

# 最初の行をスキップして、コードのみを配列に変換
tail -n +2 "$input" | cut -d',' -f1 | while read -r code; do
    echo "    \"$code\"," >> $output
done
# フッター
echo "}" >> $output
