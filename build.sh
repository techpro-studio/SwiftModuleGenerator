# shellcheck disable=SC2035
go build *.go
mv main swiftmodulegen
mv swiftmodulegen /usr/local/bin/swiftmodulegen
echo "finished"
