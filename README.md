# logger-go

Very minimal logger with stdout, stderr and debug logging.

To show debug output, use build tag `debug`.
The package will otherwise use an empty io.Writer for logger.Debug
