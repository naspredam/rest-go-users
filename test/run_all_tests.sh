
echo "RUNNING UNIT TESTS..."
go test ./... -tags=unit_tests -v && echo "UNIT TEST PASSED!" || echo "UNIT TEST FAILURE(S) FOUND!"

echo "RUNNIN INTEGRATION TESTS..."
go test ./... -tags=integration_tests -v && echo "INTEGRATION TEST PASSED!" || echo "INTEGRATION TEST FAILURE(S) FOUND!"