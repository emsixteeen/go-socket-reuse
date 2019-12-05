# go-socket-reuse
Simple Proof of Concept that demonstrates SO_REUSEADDR in Golang to achieve a large number of client connections (outbound) to many servers, using the same local TCP tuple.

Ensure that `ulimit -a` is set to something reasonably high to test.
