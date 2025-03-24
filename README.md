# logger-go

Very minimal logger with stdout, stderr and debug logging.

To suppress debug output, use build tag `release`.
This will cause the package to use an empty io.Writer for logger.Debug
