# trim

trim output to a terminal height and width

The size of the terminal is fixed to default when the file goes from stdin, for some reason term.GetSize returns error, but when a file is specified as an argument it works as expected.
