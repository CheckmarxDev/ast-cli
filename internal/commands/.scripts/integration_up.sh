#!/bin/bash

# Step 1: Check if the failedTests file exists
FAILED_TESTS_FILE="failedTests"

# Step 2: Create the failedTests file
echo "Creating $FAILED_TESTS_FILE..."
touch "$FAILED_TESTS_FILE"

# Step 3: Run all tests and write failed test names to failedTests file
echo "Running all tests..."
go test \
    -tags integration \
    -v \
    -timeout 210m \
    -coverpkg github.com/checkmarx/ast-cli/internal/commands,github.com/checkmarx/ast-cli/internal/services,github.com/checkmarx/ast-cli/internal/wrappers \
    -coverprofile cover.out \
    github.com/checkmarx/ast-cli/test/integration 2>&1 | tee test_output.log

grep -E "^--- FAIL: " test_output.log | awk '{print $3}' > "$FAILED_TESTS_FILE"

status=$?
echo "status value after tests $status"
if [ $status -ne 0 ]; then
    echo "Integration tests failed"
fi

# Step 4: Check if failedTests file is empty
if [ ! -s "$FAILED_TESTS_FILE" ]; then
    # If empty, exit with success
    echo "All tests passed."
    rm -f "$FAILED_TESTS_FILE" test_output.log
    exit 0
else
    # If not empty, rerun the failed tests
    echo "Rerunning failed tests..."
    rerun_status=0
    while IFS= read -r testName; do
        go test \
            -tags integration \
            -v \
            -timeout 210m \
            -coverpkg github.com/checkmarx/ast-cli/internal/commands,github.com/checkmarx/ast-cli/internal/services,github.com/checkmarx/ast-cli/internal/wrappers \
            -coverprofile cover.out \
            -run "^$testName$" \
            github.com/checkmarx/ast-cli/test/integration || rerun_status=1
    done < "$FAILED_TESTS_FILE"

    go tool cover -html=cover.out -o coverage.html

    # Check if any tests failed again
    if [ $rerun_status -eq 1 ]; then
        echo "Some tests are still failing."
        rm -f "$FAILED_TESTS_FILE" test_output.log
        exit 1
    else
        echo "All failed tests passed on rerun."
        rm -f "$FAILED_TESTS_FILE" test_output.log
        exit 0
    fi
fi
