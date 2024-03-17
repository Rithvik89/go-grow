**so what are these i/o operators and stdout,stdin operators.**


*os.Stdin, os.Stdout, and os.Stderr are predefined variables in the os package in Go. They represent the standard input, standard output, and standard error streams, respectively.*

*On the other hand, io.Reader and io.Writer are interfaces in the io package. They represent types that can read from or write to a stream of bytes, respectively. These interfaces are more generic and can be implemented by various types, including files, network connections, buffers, etc.*

*In summary, while os.Stdin, os.Stdout, and os.Stderr are predefined streams for standard I/O, io.Reader and io.Writer are interfaces representing any type that can read from or write to a stream of bytes.*

**Then why bufio over i/o package.**

*bufio package is built on the top of io package.*

1. Buffering: bufio introduces buffering, which means it reads or writes data in larger chunks instead of performing individual read or write operations for each byte. This reduces the number of system calls, which can be expensive, especially for small reads or writes.

2. Performance: Buffered I/O operations are often more efficient than unbuffered operations, particularly when dealing with small reads or writes, or when interacting with slow I/O devices like network connections or disks.

3. Convenience: bufio provides additional methods and abstractions that make working with buffered I/O easier and more convenient. For example, it provides a Scanner type for scanning input and breaking it into lines or tokens, as well as methods like ReadLine, ReadString, and ReadBytes for reading data in various formats.

4. Flexible Buffer Sizes: With bufio, you have control over the buffer size, allowing you to tune performance based on your specific requirements. You can adjust the buffer size according to the expected data size and the efficiency of I/O operations.

For more details visit functions **io_read** & **buf_read** under main package.