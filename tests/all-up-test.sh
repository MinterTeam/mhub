#!/bin/bash
# the directory of this script, useful for allowing this script
# to be run with any PWD
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

result=$( docker images -q peggy-base )

# builds the container containing various system deps
# also builds Peggy once in order to cache Go deps, this container
# must be rebuilt every time you run this test if you want a faster
# solution use start chains and then run tests
bash $DIR/build-container.sh

# Remove existing container instance
set +e
docker rm -f peggy_all_up_test_instance
set -e

NODES=3
TEST_TYPE=$1

# Run new test container instance
docker run --name peggy_all_up_test_instance --cap-add=NET_ADMIN -t peggy-base /bin/bash /peggy/tests/container-scripts/all-up-test-internal.sh $NODES $TEST_TYPE
